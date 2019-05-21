package fragment

var dummyData = map[ID]Raw{
	"layout": Raw{
		Template: "layouts/simple",
		Fragments: map[string][]ID{
			"content": []ID{"1"},
		},
		Data: map[string]interface{}{
			"BodyWidth": "768px",
		},
	},
	"1": Raw{
		Template: "partials/heading",
		Fragments: map[string][]ID{
			"content": []ID{"2"},
		},
		Data: map[string]interface{}{
			"BodyWidth": "768px",
		},
	},
	"2": Raw{
		Template:  "partials/paragraph",
		Fragments: map[string][]ID{},
	},
}

// IRawLoader is Loader's interface
type IRawLoader interface {
	Load(ID) Raw
}

// RawLoader loads fragments from a datastore
type RawLoader struct {
}

// NewRawLoader creates and returns a new fragment loader
func NewRawLoader() *RawLoader {
	return &RawLoader{}
}

// Load retrieves one fragmet from the datastore by its id
func (l *RawLoader) Load(id ID) Raw {
	return dummyData[id]
}
