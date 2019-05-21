package page

// ILoader is page loaders interface
type ILoader interface {
	Load(Context) Raw
}

// Loader loads the data for a single page
type Loader struct {
}

// NewLoader creates and returns a new page loader
func NewLoader() *Loader {
	return &Loader{}
}

// Load fetches and returns a page by its id
func (l *Loader) Load(context Context) Raw {
	return Raw{
		ID:         ID(context.PageID),
		FragmentID: "layout",
	}
}
