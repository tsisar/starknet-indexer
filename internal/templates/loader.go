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

// LoadTemplateSetWithFuncs loads all .tmpl files in the templateDir as a single template set with FuncMap.
func (l *Loader) LoadTemplateSetWithFuncs(mainName string, mainFile string, funcs template.FuncMap) *template.Template {
	files, err := filepath.Glob(filepath.Join(l.templateDir, "*.tmpl"))
	if err != nil {
		panic(fmt.Errorf("failed to glob templates: %w", err))
	}
	var mainPath string
	for _, f := range files {
		if filepath.Base(f) == mainFile {
			mainPath = f
			break
		}
	}
	if mainPath == "" {
		panic(fmt.Errorf("main template %s not found in %s", mainFile, l.templateDir))
	}
	tmpl := template.New(mainName).Funcs(funcs)
	tmpl, err = tmpl.ParseFiles(files...)
	if err != nil {
		panic(fmt.Errorf("failed to parse template set: %w", err))
	}
	return tmpl.Lookup(mainName)
}
