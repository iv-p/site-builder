package scheduler

import (
	"fmt"

	"github.com/iv-p/site-builder/pkg/request"

	"github.com/iv-p/site-builder/pkg/fragment"
)

var (
	// We have one of these for every site
	SiteData = map[string]NestedData{}

	// We have one of these for every page for every site
	PageData = map[string]NestedData{}
)

type NestedData struct {
	Children map[fragment.SlotID][]fragment.InstanceID
}

type IScheduler interface {
	Get(fragment.InstanceID, request.Context) (NestedData, error)
}

type Scheduler struct{}

func NewScheduler() IScheduler {
	return &Scheduler{}
}

func (s *Scheduler) Get(fragmentID fragment.InstanceID, ctx request.Context) (NestedData, error) {
	if isPageRootSection(fragmentID) {
		pageFragment, ok := PageData[string(pageID)]
		if ok {
			return pageFragment, nil
		}
	}

	pageFragment, ok := PageData[string(fragmentID)]
	if ok {
		return pageFragment, nil
	}

	siteFragment, ok := SiteData[string(fragmentID)]
	if ok {
		return siteFragment, nil
	}
	return NestedData{}, fmt.Errorf("could not find fragment with id %s for site %s and page %s", fragmentID, siteID, pageID)
}

func isPageRootSection(id fragment.InstanceID) bool {
	return id == "CONTENT"
}
