package render

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/iv-p/site-builder/pkg/content/fragment-loader"

	"github.com/iv-p/site-builder/pkg/content"
	templateloader "github.com/iv-p/site-builder/pkg/render/template-loader"
)

// ICSSRenderer is the CSSRenderer interface
type ICSSRenderer interface {
	Render(content.Content) ([]byte, error)
}

// CSSRenderer gets a page content and generates its html
type CSSRenderer struct {
	tpl *template.Template
}

// NewCSSRenderer creates and returns a new content renderer
func NewCSSRenderer(templateLoader templateloader.ILoader) *CSSRenderer {
	return &CSSRenderer{
		tpl: templateLoader.Load(),
	}
}

// Render takes a page content object and returns the generated HTML for that page
func (r *CSSRenderer) Render(c content.Content) ([]byte, error) {
	return r.getFragmentCSS(c.Fragment)
}

func (r *CSSRenderer) getFragmentCSS(root fragmentloader.Fragment) (css []byte, err error) {
	css, err = r.compileFragmentCSS(root)
	if err != nil {
		return
	}
	for _, fragments := range root.Fragments {
		for _, fragment := range fragments {
			var res []byte
			res, err = r.compileFragmentCSS(fragment)
			if err != nil {
				return
			}
			css = append(css, res...)
		}
	}
	return
}

func (r *CSSRenderer) compileFragmentCSS(fragment fragmentloader.Fragment) (css []byte, err error) {
	layout := r.tpl.Lookup(fmt.Sprintf("%s-css", fragment.Template))
	if layout == nil {
		return []byte{}, fmt.Errorf("template %s not found", fragment.Template)
	}
	w := bytes.NewBuffer([]byte{})
	err = layout.Execute(w, fragment)
	if err != nil {
		return
	}
	css = w.Bytes()
	return
}
