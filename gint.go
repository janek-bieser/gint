package gint

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin/render"
)

const (
	defaultTemplateDir string = "templates"
	defaultPartialsDir string = "templates/_partials"
	defaultTemplateExt string = "tmpl"
	defaultLayoutFile  string = "layout"
)

// HTMLRender is an implementation of the render.HTMLRender interface
// defined by the gin framework.
type HTMLRender struct {
	TemplateDir string
	TemplateExt string
	LayoutFile  string
}

// NewHTMLRender constructs a default HTMLRender object.
func NewHTMLRender() *HTMLRender {
	return &HTMLRender{
		TemplateDir: defaultTemplateDir,
		TemplateExt: defaultTemplateExt,
		LayoutFile:  defaultLayoutFile,
	}
}

// Instance is the implementation of the render.HTMLRender interface of
// the gin framework.
func (r *HTMLRender) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: template.Must(r.loadTemplate(name)),
		Data:     data,
	}
}

func (r *HTMLRender) loadTemplate(name string) (*template.Template, error) {
	layoutContent, err := loadTemplateFile("layout")
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("layout").Parse(layoutContent))

	contentString, err := loadTemplateFile(name)
	if err != nil {
		panic(err)
	}

	template.Must(tpl.New("content").Parse(contentString))
	addPartials(tpl)

	return tpl, nil
}

func addPartials(tpl *template.Template) {
	root, err := filepath.Glob(fmt.Sprintf("%s/*.%s", defaultPartialsDir, defaultTemplateExt))
	if err != nil {
		panic(err)
	}

	subfolders, err := filepath.Glob(fmt.Sprintf("%s/**/*.%s", defaultPartialsDir, defaultTemplateExt))
	if err != nil {
		panic(err)
	}

	paths := append(root, subfolders...)
	tplPrefix := defaultTemplateDir + "/"
	tplSuffix := "." + defaultTemplateExt

	for _, path := range paths {
		name := strings.TrimSuffix(strings.TrimPrefix(path, tplPrefix), tplSuffix)
		bytes, err := ioutil.ReadFile(path)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			template.Must(tpl.New(name).Parse(string(bytes)))
		}
	}
}

func loadTemplateFile(name string) (string, error) {
	path := fmt.Sprintf("%s/%s.%s", defaultTemplateDir, name, defaultTemplateExt)
	bytes, err := ioutil.ReadFile(path)
	return string(bytes), err
}
