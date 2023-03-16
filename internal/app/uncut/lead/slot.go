package lead

import (
	"encoding/json"
)

type slotMap map[string]*slot

type slotType string

const (
	trailer slotType = "trailer"
	ad      slotType = "ad"
)

type slotJson struct {
	Id       string   `json:"id"`
	SlotType slotType `json:"type"`
}

type slot struct {
	slotType  slotType
	resources []*Resource
}

func (s *slot) addContent(resources ...*Resource) {
	s.resources = append(s.resources, resources...)
}

func (s *slot) GetWeight() (weight float64) {
	weight = 0
	for _, r := range s.resources {
		weight += r.duration
	}
	return weight
}

func slotsFromJson(slotsJson []byte) (slotTemplate []slotJson, err error) {
	slots := make([]slotJson, 0, 10)
	err = json.Unmarshal(slotsJson, &slots)
	return slots, err
}
