package domain

import (
	"fmt"
	"reflect"
	"testing"
)

type DummyLoader struct {
	Domains map[string]Domain

	ILoader
}

func NewDummyLoader(domains map[string]Domain) ILoader {
	return &DummyLoader{
		Domains: domains,
	}
}

func (l *DummyLoader) GetByDomain(domain string) (Domain, error) {
	var d Domain
	var ok bool
	if d, ok = l.Domains[domain]; !ok {
		return d, fmt.Errorf("err")
	}
	return d, nil
}

func TestNewResolver(t *testing.T) {
	loader := NewLoader()
	type args struct {
		domainLoader ILoader
	}
	tests := []struct {
		name string
		args args
		want *Resolver
	}{
		{
			"should return new domain resolver",
			args{
				loader,
			},
			&Resolver{
				domainLoader: loader,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResolver(tt.args.domainLoader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResolver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolver_Resolve(t *testing.T) {
	type fields struct {
		domainLoader ILoader
	}
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Domain
		wantErr bool
	}{
		{
			"should resolve http domain",
			fields{
				domainLoader: NewDummyLoader(map[string]Domain{
					"localhost": {},
				}),
			},
			args{
				"http://localhost",
			},
			Domain{},
			false,
		},
		{
			"should resolve domain with path",
			fields{
				domainLoader: NewDummyLoader(map[string]Domain{
					"dummydomain.com": {},
				}),
			},
			args{
				"https://dummydomain.com/some/path",
			},
			Domain{},
			false,
		},
		{
			"should resolve subdomain without path",
			fields{
				domainLoader: NewDummyLoader(map[string]Domain{
					"subdomain.dummydomain.com": {},
				}),
			},
			args{
				"https://subdomain.dummydomain.com/",
			},
			Domain{},
			false,
		},
		{
			"should resolve subdomain with fragment",
			fields{
				domainLoader: NewDummyLoader(map[string]Domain{
					"subdomain.dummydomain.com": {},
				}),
			},
			args{
				"https://subdomain.dummydomain.com/#asdf",
			},
			Domain{},
			false,
		},
		{
			"should resolve domain with port",
			fields{
				domainLoader: NewDummyLoader(map[string]Domain{
					"dummydomain.com": {},
				}),
			},
			args{
				"https://dummydomain.com:8080/",
			},
			Domain{},
			false,
		},
		{
			"should not resolve non existing domain",
			fields{
				domainLoader: NewDummyLoader(map[string]Domain{}),
			},
			args{
				"https://nonexistingdomain.com/query123",
			},
			Domain{},
			true,
		},
		{
			"should not resolve url without scheme",
			fields{
				domainLoader: NewDummyLoader(map[string]Domain{}),
			},
			args{
				"subdomain.dummydomain.com",
			},
			Domain{},
			true,
		},
		{
			"should not resolve malformed url",
			fields{
				domainLoader: NewDummyLoader(map[string]Domain{}),
			},
			args{
				"notadomain",
			},
			Domain{},
			true,
		},
		{
			"should not resolve empty url",
			fields{
				domainLoader: NewDummyLoader(map[string]Domain{}),
			},
			args{
				"",
			},
			Domain{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Resolver{
				domainLoader: tt.fields.domainLoader,
			}
			got, err := r.Resolve(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resolver.Resolve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resolver.Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
