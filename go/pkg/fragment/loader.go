package fragment

var dummyData = map[InstanceID]Instance{
	"layout": {
		Template: "root",
		Nested: map[string][]InstanceID{
			"content": {"1"},
		},
		Data: map[string]interface{}{
			"BodyWidth": "768px",
		},
	},
	"1": {
		Template: "partials/heading",
		Nested: map[string][]InstanceID{
			"content": {"2"},
		},
		Data: map[string]interface{}{
			"BodyWidth": "768px",
		},
	},
	"2": {
		Template: "partials/paragraph",
		Nested:   map[string][]InstanceID{},
	},
}

// ILoader is TemplateLoader's interface
type ILoader interface {
	Load(InstanceID) Instance
}

// TemplateLoader loads fragments from a datastore
type Loader struct {
}

// NewLoader creates and returns a new fragment loader
func NewLoader() *Loader {
	return &Loader{}
}

// Load retrieves one fragmet from the datastore by its id
func (l *Loader) Load(id InstanceID) Instance {
	return dummyData[id]
}
