package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iv-p/site-builder/pkg/fragment"
	"github.com/iv-p/site-builder/pkg/template"
	"github.com/iv-p/site-builder/pkg/page"
	"github.com/iv-p/site-builder/pkg/site"
)

var siteResolver *site.Resolver
var pageResolver *page.Resolver
var pageLoader page.ILoader
var renderer fragment.IRenderer
// var siteRenderer *render.SiteRenderer

func main() {
	domainMapLoader := site.NewDomainMapLoader()
	siteResolver = site.NewResolver(domainMapLoader)

	urlMapLoader := page.NewURLMapLoader()
	pageResolver = page.NewResolver(urlMapLoader)

	fragmentLoader := fragment.NewLoader()
	templateLoader := template.NewLoader("templates/")
	renderer = fragment.NewRenderer(fragmentLoader, templateLoader)
	
	pageLoader = page.NewLoader()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(fmt.Sprintf("http://%s%s", r.Host, r.RequestURI))
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
		return
	}

	siteContext, err := siteResolver.Resolve(u)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
		return
	}

	pageContext, err := pageResolver.Resolve(siteContext, u)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
		return
	}

	pag := pageLoader.Load(pageContext)

	html, err := renderer.Render(pag.FragmentID)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(html.HTML)
}
