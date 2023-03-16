package gui

import (
	"uncut/internal/app/gui/html"
	"uncut/internal/app/uncut/lead"
)

type strategyParameter struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Input       any    `json:"input"`
}

type adSchedulingStrategyDescriptor struct {
	ID          lead.AdSchedulingStrategyName `json:"id"`
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	Parameters  []strategyParameter           `json:"parameters,omitempty"`
}

var adSchedulingStrategies = []adSchedulingStrategyDescriptor{
	{
		ID:          lead.LightestBinStrategyName,
		Name:        "Lightest Bin",
		Description: "Distributes ads equally based on their length, so that every ad block has about the same length",
		Parameters:  nil,
	},
	{
		ID:   lead.PerceivedLightnessStrategyName,
		Name: "Perceived Lightness",
		Description: "Distributes ads so that in each transition between ads " +
			"the difference in perceived lightness is as low as possible",
		Parameters: nil,
	},
	{
		ID:          lead.LightestBinAndPerceivedLightnessStrategyName,
		Name:        "Lightest Bin & Perceived Lightness",
		Description: "A combination of the first two strategies",
		Parameters:  nil,
	},
	{
		ID:   lead.BalancedFlashinessStrategyName,
		Name: "Balanced Flashiness",
		Description: "Distributes ads so that each ad block consists of some ads with less \"action\" " +
			"and some ads with more \"action\". Action is determined by the chosen algorithm.",
		Parameters: []strategyParameter{
			{
				ID:          "algorithm",
				Name:        "Algorithm",
				Description: "The algorithm to use for determining the amount of action in the ad",
				Input: html.NewSelect(map[string]string{
					lead.FlashinessAvgColor: "Average Color",
					lead.FlashinessAvgFrame: "Average Frame",
					lead.FlashinessCombined: "Combined",
				}, lead.FlashinessAvgColor),
			},
			{
				ID:   "stepSize",
				Name: "Step size",
				Description: "The number of seconds between the samples used in the algorithm. " +
					"A lower number increases the accuracy, but also the compute time.",
				Input: html.NewNumber(html.NumberOptions{
					Min:  0.1,
					Max:  5,
					Step: 0.1,
				}, 0.7),
			},
		},
	},
}
