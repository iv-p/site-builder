package fragment

var dummyData = map[ID]Fragment{
	"layout": Fragment{
		Template: "root",
		Nested: map[string][]ID{
			"content": []ID{"1"},
		},
		Data: map[string]interface{}{
			"BodyWidth": "768px",
		},
	},
	"1": Fragment{
		Template: "partials/heading",
		Nested: map[string][]ID{
			"content": []ID{"2"},
		},
		Data: map[string]interface{}{
			"BodyWidth": "768px",
		},
	},
	"2": Fragment{
		Template: "partials/paragraph",
		Nested:   map[string][]ID{},
	},
}

// ILoader is Loader's interface
type ILoader interface {
	Load(ID) Fragment
}

// Loader loads fragments from a datastore
type Loader struct {
}

// NewLoader creates and returns a new fragment loader
func NewLoader() *Loader {
	return &Loader{}
}

// Load retrieves one fragmet from the datastore by its id
func (l *Loader) Load(id ID) Fragment {
	return dummyData[id]
}
