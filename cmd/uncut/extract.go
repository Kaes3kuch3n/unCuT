package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"uncut/internal/app/analyze"
)

func Extract() {
	assets := []string{
		"BurgerMe 1",
		"BurgerMe 2",
		"Oettinger 1",
		"Oettinger 2",
		"Oettinger 3",
		"Fraunhofer",
		"Die Techniker",
		"Weißer Ring",
	}

	paths := []string{
		"assets/ads/clips/burgerme-1.mov",
		"assets/ads/clips/burgerme-2.mov",
		"assets/ads/clips/oettinger-1.mov",
		"assets/ads/clips/oettinger-2.mov",
		"assets/ads/clips/oettinger-3.mov",
		"assets/ads/clips/fraunhofer.mov",
		"assets/ads/clips/techniker.mov",
		"assets/ads/clips/weisser-ring.mov",
	}

	var err error
	videos := make([]analyze.Video, len(paths))
	for i, p := range paths {
		videos[i], err = analyze.NewVideo(p)
		if err != nil {
			panic(err)
		}
	}

	calculatePerceivedLightness(videos, assets)
	calculateFlashiness(videos, assets)
}

func calculatePerceivedLightness(videos []analyze.Video, assets []string) {
	length := 3.0
	startTimestamp := 0.0

	perceivedStartLightness := make([]float64, len(videos))
	perceivedEndLightness := make([]float64, len(videos))

	var wg sync.WaitGroup
	for i := 0; i < len(videos); i++ {
		wg.Add(2)

		go func(i int, v analyze.Video) {
			defer wg.Done()
			avgColor := analyze.GetAvgColorOverFrames(v, startTimestamp, length, v.FrameLength())
			perceivedStartLightness[i] = analyze.GetPerceivedLightness(avgColor)
		}(i, videos[i])

		go func(i int, v analyze.Video) {
			defer wg.Done()
			avgColor := analyze.GetAvgColorOverFrames(v, v.Duration()-length, length, v.FrameLength())
			perceivedEndLightness[i] = analyze.GetPerceivedLightness(avgColor)
		}(i, videos[i])
	}
	wg.Wait()

	for i := 0; i < len(videos); i++ {
		fmt.Printf("%s: %f - %f\n", assets[i], perceivedStartLightness[i], perceivedEndLightness[i])
	}
}

func calculateFlashiness(videos []analyze.Video, assets []string) {
	stepSizes := []float64{0.1, 0.15, 0.2, 0.3, 0.4, 0.5, 0.6, 0.8, 1, 1.5}
	resultsColor := make([][]float64, len(stepSizes))
	resultsImage := make([][]float64, len(stepSizes))

	for i := 0; i < len(stepSizes); i++ {
		resultsColor[i] = make([]float64, len(videos))
		resultsImage[i] = make([]float64, len(videos))
	}

	var wg sync.WaitGroup
	for j, stepSize := range stepSizes {
		for i, v := range videos {
			wg.Add(1)
			go func(i int, j int, v analyze.Video, s float64) {
				defer wg.Done()
				diff := analyze.GetAvgColorDifference(v, s)
				resultsColor[j][i] = diff
			}(i, j, v, stepSize)
		}

		for i, v := range videos {
			wg.Add(1)
			go func(i int, j int, v analyze.Video, s float64) {
				defer wg.Done()
				diff := analyze.GetAvgFrameDifference(v, s)
				resultsImage[j][i] = diff
			}(i, j, v, stepSize)
		}
	}
	wg.Wait()

	stepSizesStr := make([]string, len(stepSizes))
	for i := 0; i < len(stepSizes); i++ {
		stepSizesStr[i] = strconv.FormatFloat(stepSizes[i], 'f', -1, 64)
	}
	fmt.Printf("Step Size,%s\n", strings.Join(stepSizesStr, ","))
	printResults(assets, resultsColor)
	printResults(assets, resultsImage)
	print()
}

func printResults(assets []string, results [][]float64) {
	for j := 0; j < len(results[0]); j++ {
		fmt.Printf("%s,", assets[j])
		for i := 0; i < len(results); i++ {
			if i == len(results)-1 {
				fmt.Printf("%f\n", results[i][j])
			} else {
				fmt.Printf("%f,", results[i][j])
			}
		}
	}
}
