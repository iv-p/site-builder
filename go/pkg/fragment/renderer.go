package fragment

import (
	"bytes"
	gtpl "html/template"
	"strings"

	"github.com/iv-p/site-builder/pkg/scheduler"

	"github.com/iv-p/site-builder/pkg/page"
	"github.com/iv-p/site-builder/pkg/site"

	"github.com/iv-p/site-builder/pkg/template"
)

type IRenderer interface {
	Render(ID, page.ID, site.ID) (Compiled, error)
}

type Renderer struct {
	templateLoader template.ILoader
	fragmentLoader ILoader
	scheduler      scheduler.IScheduler
}

func NewRenderer(fragmentLoader ILoader, templateLoader template.ILoader, scheduler scheduler.IScheduler) IRenderer {
	return &Renderer{
		templateLoader: templateLoader,
		fragmentLoader: fragmentLoader,
		scheduler:      scheduler,
	}
}

func (r *Renderer) Render(fragmentID ID, pageID page.ID, siteID site.ID) (Compiled, error) {
	return r.renderFragment(fragmentID, pageID, siteID)
}

func (r *Renderer) renderFragment(fragmentID ID, pageID page.ID, siteID site.ID) (result Compiled, err error) {
	fragment := r.fragmentLoader.Load(fragmentID)
	htmlTemplate := gtpl.New("")

	var compiledSection Compiled
	nested, err := r.scheduler.Get(fragmentID, pageID, siteID)
	if err != nil {
		return
	}

	for sectionName, sectionIDs := range nested.Children {
		compiledSection, err = r.renderSection(sectionIDs)
		htmlTemplate, err = htmlTemplate.New(sectionName).Parse(string(compiledSection.HTML))
		if err != nil {
			return
		}
		result.CSS = append(result.CSS, compiledSection.CSS...)
	}

	var tpl template.Template
	tpl, err = r.templateLoader.Load(fragment.Template)
	if err != nil {
		return
	}

	definedTemplates := htmlTemplate.DefinedTemplates()
	for slotName := range tpl.Config.Slots {
		if strings.Index(definedTemplates, slotName) == -1 {
			htmlTemplate, err = htmlTemplate.New(slotName).Parse("")
		}
	}

	htmlTemplate, err = htmlTemplate.New("css").Parse(string(result.CSS))
	if err != nil {
		return
	}
	htmlTemplate, err = htmlTemplate.New(fragment.Template).Parse(string(tpl.HTML))
	if err != nil {
		return
	}

	w := bytes.NewBuffer([]byte{})
	err = htmlTemplate.Execute(w, fragment.Data)
	if err != nil {
		return
	}
	result.HTML = w.Bytes()
	return
}

func (r *Renderer) renderSection(ids []ID) (Compiled, error) {
	var compiledFragments []Compiled
	for _, fragmentID := range ids {
		compiled, err := r.renderFragment(fragmentID)
		if err != nil {
			return Compiled{}, err
		}
		compiledFragments = append(compiledFragments, compiled)
	}
	return joinCompiled(compiledFragments), nil
}

func joinCompiled(list []Compiled) (result Compiled) {
	for _, compiled := range list {
		result.HTML = append(result.HTML, compiled.HTML...)
		result.CSS = append(result.CSS, compiled.CSS...)
	}
	return
}
