package fragment

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	HTML_FILE_NAME   = "template.html"
	CSS_FILE_NAME    = "style.css"
	CONFIG_FILE_NAME = "config.yml"
)

type TemplateID string

type Type int

const (
	LAYOUT Type = 1
	HEADER Type = 2
	MENU Type = 3
	FOOTER Type = 4

)

type Template struct {
	ID     TemplateID
	Type   Type
	HTML   []byte
	CSS    []byte
	Config Config
}

type Config struct {
	ID     string
	Params []Param
	Slots  []Slot
}

type Param struct {
	Name    string
	Key     string
	Default string
}

type Slot struct {
	Name string
}

type ITemplateLoader interface {
	Load(id string) (Template, error)
}

type TemplateLoader struct {
	rootDir string
	ILoader
}

func NewTemplateLoader(rootDir string) ITemplateLoader {
	return &TemplateLoader{
		rootDir: rootDir,
	}
}

func (r *TemplateLoader) Load(path string) (t Template, err error) {
	t.HTML, err = ioutil.ReadFile(fmt.Sprintf("%s/%s/%s", r.rootDir, path, HTML_FILE_NAME))
	if err != nil {
		return
	}

	t.CSS, err = ioutil.ReadFile(fmt.Sprintf("%s/%s/%s", r.rootDir, path, CSS_FILE_NAME))
	if err != nil {
		return
	}

	var config []byte
	config, err = ioutil.ReadFile(fmt.Sprintf("%s/%s/%s", r.rootDir, path, CONFIG_FILE_NAME))
	if err != nil {
		return
	}

	err = yaml.Unmarshal(config, &t.Config)
	return
}
