package page

// Raw holds information for a single page
type Raw struct {
	ID   ID
	Data map[string]interface{}
}

// ID is the id of a page
type ID string
