package scheduler

import (
	"fmt"

	"github.com/iv-p/site-builder/pkg/request"

	"github.com/iv-p/site-builder/pkg/fragment"
)

var (
	// We have one of these for every site
	SiteData = map[string]NestedData{
		"1": {
			map[string][]fragment.ID{
				"slot1": {"2"},
				"slot2": {"4"},
			},
		},
		"2": {
			map[string][]fragment.ID{
				"slot1": {"3"},
			},
		},
		"3": {
			map[string][]fragment.ID{},
		},
		"4": {
			map[string][]fragment.ID{
				"slot1": {"CONTENT"},
			},
		},
	}

	// We have one of these for every page for every site
	PageData = map[string]NestedData{
		"11": {
			map[string][]fragment.ID{
				"slot1": {"12"},
				"slot2": {"14"},
			},
		},
		"12": {
			map[string][]fragment.ID{
				"slot1": {"13"},
			},
		},
		"13": {
			map[string][]fragment.ID{},
		},
		"14": {
			map[string][]fragment.ID{},
		},
	}
)

type NestedData struct {
	Children map[string][]fragment.ID
}

type IScheduler interface {
	Get(fragment.ID, request.Context) (NestedData, error)
}

type Scheduler struct{}

func NewScheduler() IScheduler {
	return &Scheduler{}
}

func (s *Scheduler) Get(fragmentID fragment.ID, ctx request.Context) (NestedData, error) {
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

func isPageRootSection(id fragment.ID) bool {
	return id == "CONTENT"
}
