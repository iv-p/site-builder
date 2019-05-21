package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"

	"github.com/iv-p/site-builder/pkg/content"
	"github.com/iv-p/site-builder/pkg/fragment"
	"github.com/iv-p/site-builder/pkg/page"
	"github.com/iv-p/site-builder/pkg/site"
	"github.com/iv-p/site-builder/pkg/render"
	templateloader "github.com/iv-p/site-builder/pkg/render/template-loader"
)

var tpl *template.Template
var siteResolver *site.Resolver
var pageResolver *page.Resolver
var contentLoader *content.Loader
var siteRenderer *render.SiteRenderer

func main() {
	domainMapLoader := site.NewDomainMapLoader()
	siteResolver = site.NewResolver(domainMapLoader)

	urlMapLoader := page.NewURLMapLoader()
	pageResolver = page.NewResolver(urlMapLoader)

	rawFragmentLoader := fragment.NewRawLoader()
	deepFragmentLoader := fragment.NewDeepLoader(rawFragmentLoader)
	pageLoader := page.NewLoader()
	contentLoader = content.NewLoader(pageLoader, deepFragmentLoader)

	templateLoader := templateloader.NewLoader("templates/")
	contentRenderer := render.NewContentRenderer(templateLoader)
	cssRenderer := render.NewCSSRenderer(templateLoader)
	cssMinifier := render.NewCSSMinifier()
	siteRenderer = render.NewSiteRenderer(contentRenderer, cssRenderer, cssMinifier, templateLoader)

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

	content := contentLoader.Load(siteContext, pageContext)

	html, err := siteRenderer.Render(content)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(html)
}
