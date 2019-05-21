package content

import (
	"github.com/iv-p/site-builder/pkg/fragment"
	"github.com/iv-p/site-builder/pkg/page"
	"github.com/iv-p/site-builder/pkg/site"
)

// Loader loads all the information about a page by it's id
type Loader struct {
	fragmentLoader fragment.IDeepLoader
	pageLoader     page.ILoader
}

// NewLoader creates a new loader object
func NewLoader(pageLoader page.ILoader, fragmentLoader fragment.IDeepLoader) *Loader {
	return &Loader{
		fragmentLoader: fragmentLoader,
		pageLoader:     pageLoader,
	}
}

// Load returns the page contents by id
func (l *Loader) Load(siteContext site.Context, pageContext page.Context) Content {
	page := l.pageLoader.Load(page.IDFromPageContext(pageContext))

	fragment := l.fragmentLoader.Load(page.FragmentID)

	return Content{
		Fragment: fragment,
	}
}

func (l *Loader) loadFragment(id fragment.ID) fragment.Fragment {
	return l.loadFragment(id)
}
