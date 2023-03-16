package lead

import (
	"encoding/json"
	"errors"
	"github.com/hashicorp/go-multierror"
	"github.com/samber/lo"
	"math"
	"sort"
	"sync"
	"uncut/internal/app/analyze"
	"uncut/internal/app/ffmpeg"
	"uncut/internal/app/uncut/bin"
	"uncut/internal/app/uncut/entities"
)

type AdScheduleStrategy interface {
	ScheduleAds(slots []*slot, ads []entities.Ad, content map[uint]adContent) (err error)
}

type AdSchedulingStrategyName = string

const (
	LightestBinStrategyName                      AdSchedulingStrategyName = "lightestBin"
	PerceivedLightnessStrategyName               AdSchedulingStrategyName = "perceivedLightness"
	LightestBinAndPerceivedLightnessStrategyName AdSchedulingStrategyName = "lightestBinAndPerceivedLightness"
	BalancedFlashinessStrategyName               AdSchedulingStrategyName = "balancedFlashiness"
)

type selectedStrategy struct {
	Strategy AdSchedulingStrategyName `json:"strategy"`
	Options  struct {
		Algorithm string  `json:"algorithm,omitempty"`
		StepSize  float64 `json:"stepSize,omitempty"`
	} `json:"options,omitempty"`
}

func CreateSelectedStrategy(strategyJson []byte) (strategy AdScheduleStrategy, err error) {
	var selected selectedStrategy
	err = json.Unmarshal(strategyJson, &selected)
	if err != nil {
		return nil, err
	}
	switch selected.Strategy {
	case LightestBinStrategyName:
		return LightestBinStrategy{}, nil
	case PerceivedLightnessStrategyName:
		return PerceivedLightnessStrategy{}, nil
	case LightestBinAndPerceivedLightnessStrategyName:
		return LightestBinAndPerceivedLightnessStrategy{}, nil
	case BalancedFlashinessStrategyName:
		if !lo.Contains(flashinessAlgorithms, selected.Options.Algorithm) {
			return nil, errors.New("unrecognized flashiness algorithm specified")
		}
		if selected.Options.StepSize <= 0 {
			return nil, errors.New("invalid step size: must be bigger than zero")
		}
		return BalancedFlashinessStrategy{
			Algorithm: selected.Options.Algorithm,
			StepSize:  selected.Options.StepSize,
		}, nil
	}
	return nil, errors.New("unrecognized ad scheduling strategy specified")
}

type LightestBinStrategy struct{}

func (l LightestBinStrategy) ScheduleAds(slots []*slot, ads []entities.Ad, content map[uint]adContent) (err error) {
	adGroups := mapValuesToArray(groupAdsByAdvertiser(ads))

	// Sort advertisers by amount of ads descending
	sort.Slice(adGroups, func(i, j int) bool {
		return len(adGroups[i]) > len(adGroups[j])
	})

	// Distribute ads into ad bins evenly
	for _, ads := range adGroups {
		for _, ad := range ads {
			// Find the lightest bin and add the ad to it
			b := bin.GetLightest(slots)
			err := addAdToSlot(ad, content[ad.ID], b)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type PerceivedLightnessStrategy struct{}

type adLightness struct {
	entities.Ad
	start float64
	end   float64
}

func (p PerceivedLightnessStrategy) ScheduleAds(slots []*slot, ads []entities.Ad, content map[uint]adContent) (err error) {
	if len(ads) == 0 || len(slots) == 0 {
		return nil
	}

	// Handle single ad: Skip calculations since there's nothing to sort here
	if len(ads) == 1 {
		ad := ads[0]
		centerSlot := slots[(len(slots)/2+1)%len(slots)]
		return addAdToSlot(ad, content[ad.ID], centerSlot)
	}

	l := getPerceivedLightness(ads)
	sortedAds := sortAdsByPerceivedLightness(l)

	adsPerSlot := len(sortedAds)/len(slots) + 1
	slotId := 0
	for i, ad := range sortedAds {
		if i != 0 && i%adsPerSlot == 0 {
			slotId++
		}
		err := addAdToSlot(ad, content[ad.ID], slots[slotId])
		if err != nil {
			return err
		}
	}
	return nil
}

type LightestBinAndPerceivedLightnessStrategy struct{}

type tempSlot struct {
	weight float64
	ads    []entities.Ad
}

func (t *tempSlot) Add(ad entities.Ad, weight float64) (err error) {
	t.ads = append(t.ads, ad)
	if weight != 0 {
		t.weight += weight
	} else {
		duration, err := ffmpeg.GetDuration(ad.FilePath)
		if err != nil {
			return err
		}
		t.weight += duration
	}
	return nil
}

func (t *tempSlot) GetWeight() (weight float64) {
	return t.weight
}

func (l LightestBinAndPerceivedLightnessStrategy) ScheduleAds(slots []*slot, ads []entities.Ad, content map[uint]adContent) (err error) {
	adGroups := mapValuesToArray(groupAdsByAdvertiser(ads))
	// Sort advertisers by amount of ads descending
	sort.Slice(adGroups, func(i, j int) bool {
		return len(adGroups[i]) > len(adGroups[j])
	})
	// Create temporary intermediate bins for distributing ads evenly before sorting them
	bins := make([]*tempSlot, len(slots))
	for i := range bins {
		bins[i] = &tempSlot{
			weight: 0,
			ads:    make([]entities.Ad, 0, 3),
		}
	}
	// Distribute ads into ad bins evenly
	for _, ads := range adGroups {
		for _, ad := range ads {
			content := content[ad.ID]
			// Find the lightest bin and add the ad to it
			b := bin.GetLightest(bins)
			err := b.Add(ad, content.Duration)
			if err != nil {
				return err
			}
		}
	}
	// Sort ads in bins by perceived lightness
	for i, b := range bins {
		l := getPerceivedLightness(b.ads)
		sortedAds := sortAdsByPerceivedLightness(l)

		for _, ad := range sortedAds {
			err := addAdToSlot(ad, content[ad.ID], slots[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type FlashinessAlgorithm = string

const (
	FlashinessAvgColor FlashinessAlgorithm = "avgColor"
	FlashinessAvgFrame FlashinessAlgorithm = "avgFrame"
	FlashinessCombined FlashinessAlgorithm = "combined"
)

var flashinessAlgorithms = []string{
	FlashinessAvgColor,
	FlashinessAvgFrame,
	FlashinessCombined,
}

type BalancedFlashinessStrategy struct {
	Algorithm FlashinessAlgorithm
	StepSize  float64
}

type adFlashiness struct {
	entities.Ad
	flashiness float64
}

func (b BalancedFlashinessStrategy) ScheduleAds(slots []*slot, ads []entities.Ad, content map[uint]adContent) (err error) {
	flashiness, err := b.calculateFlashiness(ads)
	if err != nil {
		return err
	}
	// Sort ads by flashiness
	sort.Slice(flashiness, func(i, j int) bool {
		return flashiness[i].flashiness < flashiness[j].flashiness
	})
	// Get ads with varying levels of flashiness from list offset by slot count
	// to distribute them equally based on their flashiness
	for i, s := range slots {
		for j := i; j < len(flashiness); j += len(slots) {
			ad := flashiness[j].Ad
			err := addAdToSlot(ad, content[ad.ID], s)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (b BalancedFlashinessStrategy) calculateFlashiness(ads []entities.Ad) (flashiness []adFlashiness, err error) {
	flashinessFunc, err := b.getFlashinessFunc()
	if err != nil {
		return nil, err
	}
	// Calculate flashiness for all ads
	flashiness = make([]adFlashiness, len(ads))
	var wg sync.WaitGroup
	errCh := make(chan error)
	defer close(errCh)
	for i, ad := range ads {
		// Flashiness for images is always 0
		if ad.Type == entities.Image {
			flashiness[i] = adFlashiness{Ad: ad, flashiness: 0}
			continue
		}
		wg.Add(1)
		// Calculate flashiness of ads concurrently
		go func(i int, errCh chan<- error) {
			defer wg.Done()
			v, err := analyze.NewVideo(ads[i].FilePath)
			if err != nil {
				errCh <- err
			}
			f := flashinessFunc(v, b.StepSize)
			flashiness[i] = adFlashiness{Ad: ads[i], flashiness: f}
		}(i, errCh)
	}
	wg.Wait()
	err = getErrors(errCh)
	if err != nil {
		return nil, err
	}
	return flashiness, nil
}

func getErrors(errCh <-chan error) (err error) {
	if len(errCh) > 0 {
		err = <-errCh
		for len(errCh) > 0 {
			err = multierror.Append(err, <-errCh)
		}
		return err
	}
	return nil
}

func (b BalancedFlashinessStrategy) getFlashinessFunc() (func(video analyze.Video, stepSize float64) (avgDiff float64), error) {
	var flashinessFunc func(video analyze.Video, stepSize float64) (avgDiff float64)
	switch b.Algorithm {
	case FlashinessAvgColor:
		flashinessFunc = analyze.GetAvgColorDifference
		break
	case FlashinessAvgFrame:
		flashinessFunc = analyze.GetAvgFrameDifference
		break
	case FlashinessCombined:
		flashinessFunc = func(v analyze.Video, s float64) (avgDiff float64) {
			resCh := make(chan float64)
			defer close(resCh)
			go func(ch chan<- float64) {
				ch <- analyze.GetAvgColorDifference(v, s)
			}(resCh)
			go func(ch chan<- float64) {
				ch <- analyze.GetAvgColorDifference(v, s)
			}(resCh)
			return <-resCh + <-resCh
		}
		break
	default:
		return nil, errors.New("unrecognized flashiness algorithm specified")
	}
	return flashinessFunc, nil
}

func getPerceivedLightness(adList []entities.Ad) (lightness []adLightness) {
	lightness = make([]adLightness, len(adList))
	var wg sync.WaitGroup
	wg.Add(len(adList))
	for i, ad := range adList {
		go func(i int, ad entities.Ad) {
			defer wg.Done()

			var (
				startColor analyze.RGBColor
				endColor   analyze.RGBColor
			)
			if ad.Type == entities.Video {
				v, err := analyze.NewVideo(ad.FilePath)
				if err != nil {
					panic(err)
				}
				startColor = analyze.GetAvgColorOverFrames(v, 0, 1, v.FrameLength())
				endColor = analyze.GetAvgColorOverFrames(v, v.Duration()-1, 1, v.FrameLength())
			} else {
				img, err := analyze.LoadImage(ad.FilePath)
				if err != nil {
					panic(err)
				}
				color := analyze.GetAvgImageColor(img)
				startColor = color
				endColor = color
			}

			lightness[i] = adLightness{
				Ad:    ad,
				start: analyze.GetPerceivedLightness(startColor),
				end:   analyze.GetPerceivedLightness(endColor),
			}
		}(i, ad)
	}
	wg.Wait()
	return lightness
}

func sortAdsByPerceivedLightness(lightness []adLightness) (sortedAds []entities.Ad) {
	sortedAds = make([]entities.Ad, 0, len(lightness))
	if len(lightness) == 1 {
		return append(sortedAds, lightness[0].Ad)
	} else if len(lightness) == 2 {
		// Simply compare these two ads to see which order is a better fit
		a := lightness[0]
		b := lightness[1]
		diffAB := math.Abs(a.end - b.start)
		diffBA := math.Abs(b.end - a.start)
		if diffAB < diffBA {
			sortedAds = append(sortedAds, a.Ad)
			sortedAds = append(sortedAds, b.Ad)
		} else {
			sortedAds = append(sortedAds, b.Ad)
			sortedAds = append(sortedAds, a.Ad)
		}
	} else {
		currentAd := lightness[0]
		sortedAds = append(sortedAds, currentAd.Ad)
		lightness = lightness[1:]
		// Append best fitting ad to sorted ads until all ads are sorted
		for len(lightness) > 0 {
			closestAdIndex := 0
			closestAdDiff := math.Abs(currentAd.end - lightness[closestAdIndex].start)
			for i := 1; i < len(lightness); i++ {
				ad := lightness[i]
				currentAdDiff := math.Abs(currentAd.end - ad.start)
				if currentAdDiff < closestAdDiff {
					closestAdIndex = i
					closestAdDiff = currentAdDiff
				}
			}
			currentAd = lightness[closestAdIndex]
			sortedAds = append(sortedAds, currentAd.Ad)
			lightness = remove(lightness, closestAdIndex)
		}
	}
	return sortedAds
}

func groupAdsByAdvertiser(ads []entities.Ad) (groupedAds map[string][]entities.Ad) {
	groupedAds = make(map[string][]entities.Ad)

	for _, ad := range ads {
		group, exists := groupedAds[ad.Advertiser.Name]
		if exists {
			groupedAds[ad.Advertiser.Name] = append(group, ad)
		} else {
			groupedAds[ad.Advertiser.Name] = []entities.Ad{ad}
		}
	}

	return groupedAds
}

func mapValuesToArray[K comparable, V any](m map[K]V) (values []V) {
	values = make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func remove[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
