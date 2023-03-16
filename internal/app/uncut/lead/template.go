package lead

import (
	"fmt"
)

type Template struct {
	slotTemplate       []slotJson
	content            content
	adScheduleStrategy AdScheduleStrategy
}

func NewTemplate(schedule []byte, content []byte, adScheduleStrategy AdScheduleStrategy) Template {
	s, err := slotsFromJson(schedule)
	if err != nil {
		panic(fmt.Errorf("invalid json schedule for lead [%w]", err))
	}
	c, err := contentFromJson(content)
	if err != nil {
		panic(fmt.Errorf("invalid content json [%w]", err))
	}
	return Template{slotTemplate: s, content: c, adScheduleStrategy: adScheduleStrategy}
}

func (t *Template) GetTrailerCount() (count int) {
	count = 0
	for _, slot := range t.slotTemplate {
		if slot.SlotType == trailer {
			count++
		}
	}
	return count
}

func (t *Template) slotsFromTemplate() (slotMap slotMap, order []*slot) {
	slotMap = make(map[string]*slot, len(t.slotTemplate))
	order = make([]*slot, len(t.slotTemplate))
	for i, s := range t.slotTemplate {
		slot := slot{
			slotType:  s.SlotType,
			resources: make([]*Resource, 0),
		}
		slotMap[s.Id] = &slot
		order[i] = &slot
	}
	return slotMap, order
}
