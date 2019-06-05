package request

import (
	"github.com/iv-p/site-builder/pkg/page"
	"github.com/iv-p/site-builder/pkg/site"
)

type Context struct {
	siteContext site.Context
	pageContext page.Context
}

func NewContext(siteContext site.Context, pageContext page.Context) *Context {
	return &Context{siteContext: siteContext, pageContext: pageContext}
}
