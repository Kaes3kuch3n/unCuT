package lead

import (
	"fmt"
	"uncut/internal/app/ffmpeg"
	"uncut/internal/app/uncut/entities"
)

type Lead struct {
	slotMap            slotMap
	slots              []*slot
	content            content
	adScheduleStrategy AdScheduleStrategy
}

// New creates a new lead from the given template
func New(template Template) *Lead {
	slots, order := template.slotsFromTemplate()
	return &Lead{
		slotMap:            slots,
		slots:              order,
		content:            template.content,
		adScheduleStrategy: template.adScheduleStrategy,
	}
}

// AddScreens adds the scheduled screens to the lead
func (l *Lead) AddScreens(screens entities.ScreenMap) {
	for _, screenContent := range l.content.Screens {
		slot, exists := l.slotMap[screenContent.Slot]
		if !exists {
			panic(fmt.Errorf("invalid slot id: %s", screenContent.Slot))
		}
		screen, exists := screens[screenContent.ID]
		if !exists {
			panic(fmt.Errorf("invalid screen type id: %d", screenContent.ID))
		}
		resource, err := NewImageResource(screen.FilePath, screenContent.Duration)
		if err != nil {
			panic(err)
		}
		slot.addContent(resource)
	}
}

func (l *Lead) AddTrailers(upcomingMovies []*entities.Movie, movieScreens []*entities.MovieScreen, upcomingScreen *entities.Screen) {
	trailerSlots := l.getSlots(trailer)
	if len(upcomingMovies) < len(trailerSlots) {
		// TODO: Add logic for distributing less than `trailerCount` trailers
		panic("not implemented")
	}

	for i, trailerSlot := range trailerSlots {
		resources, err := GetTrailerResources(upcomingMovies[i], movieScreens[i], upcomingScreen)
		if err != nil {
			panic(fmt.Errorf("failed to add trailer [%w]", err))
		}
		trailerSlot.addContent(resources...)
	}
}

func (l *Lead) AddAds(ads entities.AdMap) {
	scheduled, unscheduled := l.splitAds(ads)

	err := l.addScheduledAds(scheduled)
	if err != nil {
		panic(fmt.Errorf("error while adding advertisement [%w]", err))
	}

	// Get map of unscheduled ads to add to the lead
	unscheduledAds := make(map[uint]adContent, len(l.content.Ads.Unscheduled))
	for _, a := range l.content.Ads.Unscheduled {
		// Map the ad id to the ad content for getting the duration by the id
		// (and maybe also other properties in the future)
		unscheduledAds[a.ID] = a
	}

	err = l.adScheduleStrategy.ScheduleAds(l.getSlots(ad), unscheduled, unscheduledAds)
	if err != nil {
		panic(fmt.Errorf("error while adding advertisement [%w]", err))
	}
}

func (l *Lead) Generate(outputPath string) (err error) {
	// Get all resources in the correct order
	resources := make([]*Resource, 0, 64)
	for _, s := range l.slots {
		resources = append(resources, s.resources...)
	}

	streams := make([]*ffmpeg.Formatted, 0, len(resources))
	// Apply video filters
	for _, r := range resources {
		stream := r.ToImport().Filter(ffmpeg.KVArgs{
			"scale":     "1920:1080",
			"setsar":    1,
			"framerate": 25,
		})
		streams = append(streams, stream)
	}
	// Concatenate all streams
	return ffmpeg.Concat(streams, []string{"-framerate", "25", outputPath})
}

func (l *Lead) addScheduledAds(ads entities.AdMap) (err error) {
	for _, ad := range l.content.Ads.Scheduled {
		slot, exists := l.slotMap[ad.Slot]
		if !exists {
			return fmt.Errorf("invalid slot id for ad: %s", ad.Slot)
		}
		err := addAdToSlot(ads[ad.ID], ad.adContent, slot)
		if err != nil {
			return fmt.Errorf("failed to add advertisement [%w]", err)
		}
	}
	return nil
}

func addAdToSlot(ad entities.Ad, adContent adContent, slot *slot) (err error) {
	var adResource *Resource
	if adContent.Duration != 0 {
		adResource, err = AdToResourceWithDuration(ad, adContent.Duration)
	} else {
		adResource, err = AdToResource(ad)
	}
	if err != nil {
		return err
	}
	slot.addContent(adResource)
	return nil
}

func (l *Lead) splitAds(ads entities.AdMap) (scheduled entities.AdMap, unscheduled []entities.Ad) {
	scheduled = make(entities.AdMap, len(l.content.Ads.Scheduled))
	for _, ad := range l.content.Ads.Scheduled {
		scheduled[ad.ID] = ads[ad.ID]
	}

	unscheduled = make([]entities.Ad, 0, len(l.content.Ads.Unscheduled))
	for _, ad := range l.content.Ads.Unscheduled {
		unscheduled = append(unscheduled, ads[ad.ID])
	}

	return scheduled, unscheduled
}

func (l *Lead) getSlots(t slotType) (slots []*slot) {
	slots = make([]*slot, 0, 3)
	for _, s := range l.slots {
		if s.slotType == t {
			slots = append(slots, s)
		}
	}
	return slots
}
