package render

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/iv-p/site-builder/pkg/content"
	templateloader "github.com/iv-p/site-builder/pkg/render/template-loader"
)

// IRenderer is the interface of a renderer
type IRenderer interface {
	Render(content.Content) ([]byte, error)
}

// SiteRenderer gets a page content and generates its html
type SiteRenderer struct {
	contentRenderer IContentRenderer
	cssRenderer     ICSSRenderer
	CSSMinifier     ICSSMinifier
	tpl             *template.Template
}

// NewSiteRenderer creates and returns a new content renderer
func NewSiteRenderer(contentRenderer IContentRenderer, cssRenderer ICSSRenderer, cssMinifier ICSSMinifier, templateLoader templateloader.ILoader) *SiteRenderer {
	return &SiteRenderer{
		contentRenderer: contentRenderer,
		cssRenderer:     cssRenderer,
		CSSMinifier:     cssMinifier,
		tpl:             templateLoader.Load(),
	}
}

// Render takes a page content object and returns the generated HTML for that page
func (r *SiteRenderer) Render(c content.Content) ([]byte, error) {
	templates, err := r.tpl.Clone()
	if err != nil {
		return []byte{}, err
	}

	renderedContent, err := r.contentRenderer.Render(c)
	if err != nil {
		return []byte{}, err
	}
	_, err = templates.New("content").Parse(string(renderedContent))
	if err != nil {
		return []byte{}, err
	}

	cssContent, err := r.cssRenderer.Render(c)
	if err != nil {
		return []byte{}, err
	}
	minifiedCSS, err := r.CSSMinifier.Minify(cssContent)
	if err != nil {
		return []byte{}, err
	}

	_, err = templates.New("styles").Parse(string(minifiedCSS))
	if err != nil {
		return []byte{}, err
	}

	layout := templates.Lookup("root-html")
	if layout == nil {
		return []byte{}, fmt.Errorf("template %s not found", c.Fragment.Template)
	}
	w := bytes.NewBuffer([]byte{})
	err = layout.Execute(w, c)
	return w.Bytes(), err
}
