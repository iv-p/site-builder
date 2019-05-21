package fragment

// ID is the id of a fragment
type ID string

// Fragment is the database representation of a fragment
type Fragment struct {
	ID       ID
	Template string
	Data     map[string]interface{}
	Nested   map[string][]ID
}

type Compiled struct {
	HTML []byte
	CSS  []byte
}
