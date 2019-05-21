package template

type Template struct {
	ID     string
	HTML   []byte
	CSS    []byte
	Config Config
}

type Config struct {
	ID     string
	Params []Param
	Slots  Slots
}

type Param struct {
	Name    string
	Key     string
	Default string
}

type Slots map[string]Slot

type Slot struct {
	Multiple bool
}
