package lead

import "encoding/json"

type scheduledContent struct {
	Slot string `json:"slot"`
}

type screenContent struct {
	scheduledContent
	ID       uint    `json:"id"`
	Duration float64 `json:"duration"`
}

type adContent struct {
	ID       uint    `json:"id"`
	Duration float64 `json:"duration,omitempty"`
}

type scheduledAdContent struct {
	adContent
	scheduledContent
}

type content struct {
	Screens []screenContent `json:"screens"`
	Ads     struct {
		Scheduled   []scheduledAdContent `json:"scheduled"`
		Unscheduled []adContent          `json:"unscheduled"`
	} `json:"ads"`
}

func contentFromJson(contentJson []byte) (c content, err error) {
	err = json.Unmarshal(contentJson, &c)
	return c, err
}
