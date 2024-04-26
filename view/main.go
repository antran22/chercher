package view

import (
	"bytes"
	"embed"
	"errors"
	"github.com/Masterminds/sprig/v3"
	"html/template"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
)

// This example illustrates that the FuncMap *must* be set before the
// templates themselves are loaded.

type templateMap map[string]*template.Template
type Renderer struct {
	templates *templateMap
}

//go:embed all:templates
var tmplFS embed.FS

func MakeRenderer() (*Renderer, error) {
	tmplMap := make(templateMap)
	sprigFuncMap := sprig.FuncMap()

	err := fs.WalkDir(tmplFS, "templates", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(path, ".html") {
			if err != nil {
				return err
			}
			parts := strings.Split(path, string(os.PathSeparator))
			if len(parts) < 2 || parts[1] == "layout" {
				return nil
			}
			name := strings.Join(parts[1:], string(os.PathSeparator))
			log.Println("Register template", path, "with name", name)

			fileTmpl, err := template.New(name).Funcs(sprigFuncMap).ParseFS(tmplFS, "templates/layout/*.html", path)

			if err != nil {
				return err
			}

			tmplMap[name] = fileTmpl
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &Renderer{
		templates: &tmplMap,
	}, nil
}

func (r *Renderer) render(w io.Writer, name string, data any) error {
	tmpl, ok := (*r.templates)[name]
	if !ok {
		return errors.New("template not found: " + name)
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func (r *Renderer) renderString(name string, data any) (string, error) {
	buff := bytes.Buffer{}
	err := r.render(&buff, name, data)

	if err != nil {
		return "", err
	}
	return buff.String(), nil

}
