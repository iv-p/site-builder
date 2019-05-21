package render

import (
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
)

// ICSSMinifier is the interface of CSSMinifier
type ICSSMinifier interface {
	Minify([]byte) ([]byte, error)
}

// CSSMinifier takes a byte array with the compiled css and minifies it
type CSSMinifier struct {
	m *minify.M
}

// NewCSSMinifier creates and returns a CSSMinifier object
func NewCSSMinifier() ICSSMinifier {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	// m.AddFunc("text/html", html.Minify)
	// m.AddFunc("image/svg+xml", svg.Minify)
	// m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	// m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	// m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)

	return &CSSMinifier{
		m: m,
	}
}

// Minify gets a css in a byte array and minifies it
func (m *CSSMinifier) Minify(raw []byte) ([]byte, error) {
	return m.m.Bytes("text/css", raw)
}
