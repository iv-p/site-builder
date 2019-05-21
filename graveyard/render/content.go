package render

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/iv-p/site-builder/pkg/content"
	templateloader "github.com/iv-p/site-builder/pkg/render/template-loader"
)

// IContentRenderer is ContentRenderer's interface
type IContentRenderer interface {
	Render(content.Content) ([]byte, error)
}

// ContentRenderer gets a page content and generates its html
type ContentRenderer struct {
	tpl *template.Template
}

// NewContentRenderer creates and returns a new content renderer
func NewContentRenderer(templateLoader templateloader.ILoader) *ContentRenderer {
	return &ContentRenderer{
		tpl: templateLoader.Load(),
	}
}

// Render takes a page content object and returns the generated HTML for that page
func (r *ContentRenderer) Render(c content.Content) ([]byte, error) {
	layout := r.tpl.Lookup(fmt.Sprintf("%s-html", c.Fragment.Template))
	if layout == nil {
		return []byte{}, fmt.Errorf("template %s not found", c.Fragment.Template)
	}
	w := bytes.NewBuffer([]byte{})
	err := layout.Execute(w, c.Fragment)
	return w.Bytes(), err
}
