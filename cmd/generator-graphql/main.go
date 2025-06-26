package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/tsisar/starknet-indexer/internal/templates"
)

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
	reField := regexp.MustCompile(`^\s*(\w+)(?:\([^)]*\))?:\s+(.+)$`)
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
				fmt.Printf("Field: %s, Raw type: %s, Normalized type: %s\n", m[1], m[2], fieldType)
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
	// Handle complex field types like: positions(where: PositionWhereInput, orderBy: PositionOrderBy, first: Int, skip: Int): [Position!]!

	// If it contains parameters (parentheses), extract the return type
	if strings.Contains(raw, "(") && strings.Contains(raw, "):") {
		// Find the return type part after "):"
		parts := strings.Split(raw, "):")
		if len(parts) >= 2 {
			raw = strings.TrimSpace(parts[1])
		}
	}

	// Remove array brackets and non-null markers
	t := strings.TrimSuffix(raw, "!")
	t = strings.TrimPrefix(t, "[")
	t = strings.TrimSuffix(t, "]")
	t = strings.TrimSuffix(t, "!")

	return strings.TrimSpace(t)
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
	loader := templates.NewLoader("cmd/generator-graphql/templates")

	types, enums, err := parseSchema("graphql/schema.graphqls")
	if err != nil {
		panic(err)
	}

	// build enum map
	enumMap := make(map[string]bool)
	for _, e := range enums {
		enumMap[e.Name] = true
	}

	funcMap := template.FuncMap{
		"baseType": baseType,
		"isScalar": func(t string) bool { return isScalar(t, enumMap) },
		"lower":    toLowerCamelCase,
	}

	whereTpl := loader.LoadTemplateWithFuncs("where", "where.tmpl", funcMap)
	orderTpl := loader.LoadTemplateWithFuncs("order", "order.tmpl", funcMap)
	queryTpl := loader.LoadTemplateWithFuncs("query", "query.tmpl", funcMap)

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
		return s
	}
	r := []rune(s)
	r[0] = []rune(strings.ToLower(string(r[0])))[0]
	return string(r)
}
