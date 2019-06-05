package domain

import "github.com/iv-p/site-builder/pkg/site"

type ID string

type Domain struct {
	ID     ID
	Domain string
	Site   site.ID
}
