package fragment

import "github.com/google/uuid"

// InstanceID is the id of a fragment
type InstanceID string

type SlotID string

// Instance is the database representation of a fragment
type Instance struct {
	ID         InstanceID
	TemplateID TemplateID
	Data       InstanceData
	Nested     map[string][]InstanceID
	// maps slot name to slot id
	Slots map[SlotID]string
}

func NewInstance(template Template) *Instance {
	return &Instance{
		ID:         InstanceID(uuid.New().String()),
		TemplateID: template.ID,
		Data:       NewInstanceDataFromParams(template.Config.Params),
		Slots:      MapSlotsToId(template.Config.Slots),
	}
}

type InstanceData struct {
	Default  map[string]interface{}
	Override map[string]interface{}
}

type Compiled struct {
	HTML []byte
	CSS  []byte
}

func NewInstanceDataFromParams(params []Param) InstanceData {
	data := InstanceData{
		make(map[string]interface{}),
		make(map[string]interface{}),
	}
	for _, param := range params {
		data.Default[param.Key] = param.Default
	}
	return data
}

func MapSlotsToId(slots []Slot) map[SlotID]string {
	m := make(map[SlotID]string)
	for _, slot := range slots {
		m[SlotID(uuid.New().String())] = slot.Name
	}
	return m
}
