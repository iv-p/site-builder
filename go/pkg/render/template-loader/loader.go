package templateloader

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	fragmentloader "github.com/iv-p/site-builder/pkg/content/fragment-loader"
)

// ILoader is the interface of Loader
type ILoader interface {
	Load() *template.Template
}

// Loader loads html template files
type Loader struct {
	dir string
}

// NewLoader creates and returns a new template loader
func NewLoader(dir string) *Loader {
	return &Loader{
		dir: dir,
	}
}

// Load traverses a directory and loads all the templates setting
// the names of the templates to the relative filepath to the directory
func (tl *Loader) Load() *template.Template {
	tpl := template.New("")

	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"CallTemplate": func(fragment fragmentloader.Fragment) (ret template.HTML, err error) {
			buf := bytes.NewBuffer([]byte{})
			err = tpl.ExecuteTemplate(buf, fmt.Sprintf("%s-html", fragment.Template), fragment)
			ret = template.HTML(buf.String())
			return
		},
	}

	files := []string{}
	err := filepath.Walk(tl.dir, func(path string, f os.FileInfo, err error) error {
		if strings.Index(path, ".yml") != -1 {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)

	}
	fmt.Println(files)

	for _, f := range files {
		folder := getFolder(f)
		htmlFileName, cssFileName := getTemplateFileLocations(folder)

		htmlContent, err := ioutil.ReadFile(htmlFileName)
		if err != nil {
			panic(err)
		}

		cssContent, err := ioutil.ReadFile(cssFileName)
		if err != nil {
			panic(err)
		}

		f = strings.Replace(f, tl.dir, "", -1)

		fmt.Printf("%s %s", htmlFileName, cssFileName)
		relativeFolder := strings.Replace(folder, tl.dir, "", -1)
		_, err = tpl.New(fmt.Sprintf("%s-html", relativeFolder)).Funcs(funcMap).Parse(string(htmlContent))
		_, err = tpl.New(fmt.Sprintf("%s-css", relativeFolder)).Funcs(funcMap).Parse(string(cssContent))
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(tpl.DefinedTemplates())
	return tpl
}

func getFolder(file string) string {
	return filepath.Dir(file)
}

func getTemplateFileLocations(folder string) (html, css string) {
	html = fmt.Sprintf("%s/template.html", folder)
	css = fmt.Sprintf("%s/style.css", folder)
	return
}
