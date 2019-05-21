package template

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	htmlFileName   = "template.html"
	cssFileName    = "style.css"
	configFileName = "config.yml"
)

type ILoader interface {
	Load(id string) (Template, error)
}

type Loader struct {
	rootDir string
	ILoader
}

func NewLoader(rootDir string) ILoader {
	return &Loader{
		rootDir: rootDir,
	}
}

func (r *Loader) Load(path string) (t Template, err error) {
	t.HTML, err = ioutil.ReadFile(fmt.Sprintf("%s/%s/%s", r.rootDir, path, htmlFileName)) // just pass the file name
	if err != nil {
		return
	}

	t.CSS, err = ioutil.ReadFile(fmt.Sprintf("%s/%s/%s", r.rootDir, path, cssFileName)) // just pass the file name
	if err != nil {
		return
	}

	var config []byte
	config, err = ioutil.ReadFile(fmt.Sprintf("%s/%s/%s", r.rootDir, path, configFileName)) // just pass the file name
	if err != nil {
		return
	}

	err = yaml.Unmarshal(config, &t.Config)
	return
}
