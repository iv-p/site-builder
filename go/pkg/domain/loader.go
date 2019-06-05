package domain

import "fmt"

// DomainMap is a map of domains to site ids
type DomainMap []Domain

// DummyData holds a dummy map for local development
var DummyData = []Domain{
	{
		ID:     "1",
		Domain: "localhost",
		Site:   "1",
	},
	{
		ID:     "2",
		Domain: "dummydomain.com",
		Site:   "2",
	},
	{
		ID:     "3",
		Domain: "subdomain.dummydomain.com",
		Site:   "3",
	},
}

// ILoader is the interface of a domain map loader
type ILoader interface {
	Load(ID) (Domain, error)
	GetByDomain(string) (Domain, error)
}

// Loader resolves site url data to a dummy object
type Loader struct {
}

// NewDomainMapLoader creates and returns a new dummy loader object
func NewLoader() ILoader {
	return &Loader{}
}

// Load fetches the DomainMap from an appropriate data source and returns it
func (l *Loader) Load(id ID) (Domain, error) {
	for _, domain := range DummyData {
		if domain.ID == id {
			return domain, nil
		}
	}
	return Domain{}, fmt.Errorf("domain with id %s not found", id)
}

// GetByDomain returns the domain object for a specific domain string
func (l *Loader) GetByDomain(domain string) (Domain, error) {
	for _, d := range DummyData {
		if d.Domain == domain {
			return d, nil
		}
	}
	return Domain{}, fmt.Errorf("domain object for domain %s not found", domain)
}
