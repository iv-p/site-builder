package content

import (
	fragmentloader "github.com/iv-p/site-builder/pkg/content/fragment-loader"
	pageloader "github.com/iv-p/site-builder/pkg/content/page-loader"
	"github.com/iv-p/site-builder/pkg/context/page"
	"github.com/iv-p/site-builder/pkg/context/site"
)

// Loader loads all the information about a page by it's id
type Loader struct {
	fragmentLoader fragmentloader.IDeepLoader
	pageLoader     pageloader.ILoader
}

// NewLoader creates a new loader object
func NewLoader(pageLoader pageloader.ILoader, fragmentLoader fragmentloader.IDeepLoader) *Loader {
	return &Loader{
		fragmentLoader: fragmentLoader,
		pageLoader:     pageLoader,
	}
}

// Load returns the page contents by id
func (l *Loader) Load(siteContext site.Context, pageContext page.Context) Content {
	page := l.pageLoader.Load(pageloader.IDFromPageContext(pageContext))

	fragment := l.fragmentLoader.Load(page.FragmentID)

	return Content{
		Fragment: fragment,
	}
}

func (l *Loader) loadFragment(id fragmentloader.ID) fragmentloader.Fragment {
	return l.loadFragment(id)
}
