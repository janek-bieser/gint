package gint

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin/render"
)

const (
	defaultTemplateDir string = "templates"
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
	layoutPath := fmt.Sprintf("%s/%s.%s", defaultTemplateDir, r.LayoutFile, defaultTemplateExt)
	path := fmt.Sprintf("%s/%s.%s", defaultTemplateDir, name, defaultTemplateExt)

	return template.ParseFiles(layoutPath, path)
}
