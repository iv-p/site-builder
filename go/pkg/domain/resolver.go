package domain

import (
	"net"
	"net/url"
	"strings"
)

type IResolver interface {
	Resolve(string) (Domain, error)
}

// Resolver extracts the site id from the request url
type Resolver struct {
	// Map of site domain name to site id
	domainLoader ILoader

	IResolver
}

// NewResolver creates a new site id resolver
func NewResolver(domainLoader ILoader) IResolver {
	return &Resolver{
		domainLoader: domainLoader,
	}
}

// Resolve takes a request url and finds the appropriate
// site id it is for
func (r *Resolver) Resolve(raw string) (Domain, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return Domain{}, err
	}
	domain := u.Host
	if strings.Index(domain, ":") != -1 {
		domain, _, err = net.SplitHostPort(u.Host)
		if err != nil {
			return Domain{}, err
		}
	}
	return r.domainLoader.GetByDomain(domain)
}
