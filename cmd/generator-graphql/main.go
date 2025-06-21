package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

//go:embed templates/where.tmpl
var whereTmpl string

//go:embed templates/order.tmpl
var orderTmpl string

//go:embed templates/query.tmpl
var queryTmpl string

type Field struct {
	Name string
	Type string
}

type GraphQLType struct {
	Name   string
	Fields []Field
}

type GraphQLEnum struct {
	Name   string
	Values []string
}

type FilterContext struct {
	Name   string
	Fields []Field
}

type QueryContext struct {
	Types []GraphQLType
}

var skipTypes = map[string]bool{
	"Query":    true,
	"Mutation": true,
	"Block":    true,
	"Meta":     true,
}

func parseSchema(path string) ([]GraphQLType, []GraphQLEnum, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var types []GraphQLType
	var enums []GraphQLEnum
	var currentType *GraphQLType
	var currentEnum *GraphQLEnum

	reType := regexp.MustCompile(`^type\s+(\w+)\s*{`)
	reField := regexp.MustCompile(`^\s*(\w+):\s+([\w!$begin:math:display$$end:math:display$]+)`)
	reEnum := regexp.MustCompile(`^enum\s+(\w+)\s*{`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		if m := reType.FindStringSubmatch(line); m != nil {
			currentType = &GraphQLType{Name: m[1]}
			continue
		}
		if m := reEnum.FindStringSubmatch(line); m != nil {
			currentEnum = &GraphQLEnum{Name: m[1]}
			continue
		}
		if currentType != nil {
			if line == "}" {
				types = append(types, *currentType)
				currentType = nil
				continue
			}
			if m := reField.FindStringSubmatch(line); m != nil {
				fieldType := normalizeType(m[2])
				currentType.Fields = append(currentType.Fields, Field{Name: m[1], Type: fieldType})
			}
		}
		if currentEnum != nil {
			if line == "}" {
				enums = append(enums, *currentEnum)
				currentEnum = nil
				continue
			}
			currentEnum.Values = append(currentEnum.Values, line)
		}
	}
	return types, enums, scanner.Err()
}

func normalizeType(raw string) string {
	t := strings.TrimSuffix(raw, "!")
	t = strings.TrimPrefix(t, "[")
	t = strings.TrimSuffix(t, "]")
	t = strings.TrimSuffix(t, "!")
	return t
}

func baseType(t string) string {
	// scalar mapping only for builtins, all others — leave as is
	switch t {
	case "Int", "ID", "Boolean", "Float", "String":
		return t
	default:
		return t
	}
}

func isScalar(t string, enums map[string]bool) bool {
	return t == "String" || t == "Int" || t == "ID" || t == "Boolean" || t == "Float" || enums[t]
}

func main() {
	types, enums, err := parseSchema("graphql/schema.graphqls")
	if err != nil {
		panic(err)
	}

	// build enum map
	enumMap := make(map[string]bool)
	for _, e := range enums {
		enumMap[e.Name] = true
	}

	whereTpl := template.Must(template.New("where").
		Funcs(template.FuncMap{
			"baseType": baseType,
			"isScalar": func(t string) bool { return isScalar(t, enumMap) },
		}).
		Parse(whereTmpl))

	orderTpl := template.Must(template.New("order").Parse(orderTmpl))

	queryTpl := template.Must(template.New("query").
		Funcs(template.FuncMap{"lower": toLowerCamelCase}).
		Parse(queryTmpl))

	for _, t := range types {
		if skipTypes[t.Name] {
			continue
		}
		ctx := FilterContext{Name: t.Name, Fields: t.Fields}
		var buf bytes.Buffer

		if err := whereTpl.Execute(&buf, ctx); err != nil {
			panic(err)
		}
		buf.WriteString("\n")
		if err := orderTpl.Execute(&buf, ctx); err != nil {
			panic(err)
		}

		out := filepath.Join("graphql/filters", strings.ToLower(t.Name)+".graphqls")
		if err := os.MkdirAll(filepath.Dir(out), 0755); err != nil {
			panic(err)
		}
		if err := os.WriteFile(out, buf.Bytes(), 0644); err != nil {
			panic(err)
		}
	}

	var filteredTypes []GraphQLType
	for _, t := range types {
		if skipTypes[t.Name] {
			continue
		}
		filteredTypes = append(filteredTypes, t)
	}

	qctx := QueryContext{Types: filteredTypes}
	var qbuf bytes.Buffer
	if err := queryTpl.Execute(&qbuf, qctx); err != nil {
		panic(err)
	}
	if err := os.WriteFile("graphql/filters/query.graphqls", qbuf.Bytes(), 0644); err != nil {
		panic(err)
	}

	fmt.Println("Filters generated to ./graphql/filters/")
}

func toLowerCamelCase(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
