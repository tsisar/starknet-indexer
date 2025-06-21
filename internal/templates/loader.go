package templates

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Loader struct {
	templateDir string
}

func NewLoader(templateDir string) *Loader {
	return &Loader{
		templateDir: templateDir,
	}
}

func (l *Loader) LoadTemplate(name, filename string) *template.Template {
	path := filepath.Join(l.templateDir, filename)
	content, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("failed to read template %s from %s: %w", name, path, err))
	}
	return template.Must(template.New(name).Parse(string(content)))
}

func (l *Loader) LoadTemplateWithFuncs(name, filename string, funcs template.FuncMap) *template.Template {
	path := filepath.Join(l.templateDir, filename)
	content, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("failed to read template %s from %s: %w", name, path, err))
	}
	tmpl := template.Must(template.New(name).Funcs(funcs).Parse(string(content)))
	return tmpl
}

func (l *Loader) LoadAllTemplates() map[string]*template.Template {
	templates := make(map[string]*template.Template)

	files, err := filepath.Glob(filepath.Join(l.templateDir, "*.tmpl"))
	if err != nil {
		panic(fmt.Errorf("failed to glob templates: %w", err))
	}

	for _, file := range files {
		name := filepath.Base(file)
		name = strings.TrimSuffix(name, ".tmpl")
		templates[name] = l.LoadTemplate(name, name+".tmpl")
	}

	return templates
}
