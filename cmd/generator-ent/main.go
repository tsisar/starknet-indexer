package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

const outDir = "generated/ent/schema"

var schemaTmpl = template.Must(template.New("schema").Parse(`
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	{{- if .HasEdges }}
	"entgo.io/ent/schema/edge"
	{{- end }}
)

type {{.Name}} struct {
	ent.Schema
}

func ({{.Name}}) Fields() []ent.Field {
	return []ent.Field{
		{{range .Fields}}{{.}},
		{{end}}
	}
}

{{ if .HasEdges }}
func ({{.Name}}) Edges() []ent.Edge {
	return []ent.Edge{
		{{range .Edges}}{{.}},
		{{end}}
	}
}
{{ end }}
`))

var enumTmpl = template.Must(template.New("enum").Parse(`
package schema

import "entgo.io/ent/schema/field"

var {{.Name}}Enum = field.Enum("{{.FieldName}}").
	Values({{range $i, $v := .Values}}{{if $i}}, {{end}}"{{$v}}"{{end}})
`))

func main() {
	var sources []*ast.Source

	mainSchema, err := os.ReadFile("graphql/schema.graphqls")
	if err != nil {
		panic(fmt.Errorf("read main schema failed: %w", err))
	}
	sources = append(sources, &ast.Source{
		Name:  "schema.graphqls",
		Input: string(mainSchema),
	})

	files, err := filepath.Glob("graphql/filters/*.graphqls")
	if err != nil {
		panic(fmt.Errorf("glob failed: %w", err))
	}

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			panic(fmt.Errorf("read filter schema failed: %w", err))
		}
		sources = append(sources, &ast.Source{
			Name:  filepath.Base(file),
			Input: string(content),
		})
	}

	schema, err := gqlparser.LoadSchema(sources...)
	if err != nil {
		panic(fmt.Errorf("parse schema failed: %w", err))
	}

	if err := os.MkdirAll(outDir, 0755); err != nil {
		panic(err)
	}

	for name, e := range schema.Types {
		if e.Kind != ast.Enum {
			continue
		}
		if strings.HasPrefix(name, "__") {
			continue
		}
		generateEnum(name, e)
	}

	for typeName, t := range schema.Types {
		if strings.HasPrefix(typeName, "__") || typeName == "Query" || t.Kind != ast.Object {
			continue
		}

		var fields, edges []string
		for _, f := range t.Fields {
			if isRelation(schema, f.Type) {
				edges = append(edges, generateEdge(f))
			} else {
				fields = append(fields, generateField(f))
			}
		}

		var buf bytes.Buffer
		if err := schemaTmpl.Execute(&buf, map[string]any{
			"Name":     typeName,
			"Fields":   fields,
			"Edges":    edges,
			"HasEdges": len(edges) > 0,
		}); err != nil {
			panic(err)
		}

		outFile := filepath.Join(outDir, toSnakeCase(typeName)+".go")
		if err := os.WriteFile(outFile, buf.Bytes(), 0644); err != nil {
			panic(err)
		}
		fmt.Println("Generated:", outFile)
	}
}

func generateField(f *ast.FieldDefinition) string {
	goType := mapType(f.Type)
	return fmt.Sprintf("field.%s(\"%s\")", goType, f.Name)
}

func mapType(t *ast.Type) string {
	inner := unwrapType(t)
	switch inner.Name() {
	case "String", "ID":
		return "String"
	case "Boolean":
		return "Bool"
	default:
		return "String"
	}
}

func isRelation(schema *ast.Schema, t *ast.Type) bool {
	inner := unwrapType(t)
	if scalarTypes(inner.Name()) {
		return false
	}
	if schema.Types[inner.Name()].Kind == ast.Enum {
		return false
	}
	return true
}

func scalarTypes(name string) bool {
	return name == "String" || name == "ID" || name == "Boolean"
}

func generateEdge(f *ast.FieldDefinition) string {
	inner := unwrapType(f.Type)
	if f.Type.Elem != nil {
		return fmt.Sprintf("edge.To(\"%s\", %s.Type)", f.Name, inner.Name())
	}
	return fmt.Sprintf("edge.To(\"%s\", %s.Type).Unique()", f.Name, inner.Name())
}

func unwrapType(t *ast.Type) *ast.Type {
	for t.Elem != nil {
		t = t.Elem
	}
	return t
}

func generateEnum(name string, e *ast.Definition) {
	values := []string{}
	for _, v := range e.EnumValues {
		values = append(values, v.Name)
	}

	var buf bytes.Buffer
	if err := enumTmpl.Execute(&buf, map[string]any{
		"Name":      name,
		"FieldName": toSnakeCase(name),
		"Values":    values,
	}); err != nil {
		panic(err)
	}

	outFile := filepath.Join(outDir, toSnakeCase(name)+"_enum.go")
	if err := os.WriteFile(outFile, buf.Bytes(), 0644); err != nil {
		panic(err)
	}
	fmt.Println("Generated enum:", outFile)
}

func toSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if unicode.IsUpper(r) && i > 0 {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}
