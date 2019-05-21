package pageloader

import (
	"github.com/iv-p/site-builder/pkg/content/fragment-loader"
)

// Raw holds information for a single page
type Raw struct {
	ID         ID
	FragmentID fragmentloader.ID
}

// ID is the id of a page
type ID string
