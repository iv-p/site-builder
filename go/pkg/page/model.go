package page

import (
	"github.com/iv-p/site-builder/pkg/fragment"
)

// Raw holds information for a single page
type Raw struct {
	ID         ID
	FragmentID fragment.ID
}

// ID is the id of a page
type ID string
