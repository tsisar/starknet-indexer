package abiutil

import (
	"bytes"
	_ "embed"
	"github.com/Tsisar/starknet-indexer/internal/utils"
	"log"
	"strings"
	"text/template"
)

//go:embed templates/events.tmpl
var structsTmpl string

type Field struct {
	Name     string
	GoType   string
	IndexOff int
}

type EventData struct {
	Name       string
	StructName string
	Kind       string
	Selector   string
	KeyFields  []Field
	DataFields []Field
}

type TemplateContext struct {
	Package        string
	Events         []EventData
	NeedsBigImport bool
}

func GenerateGoStructs(events []CairoAbiEntry, packageName string) string {
	ctx := TemplateContext{Package: packageName}
	needsBig := false

	for _, ev := range events {
		if ev.Kind != "struct" {
			continue
		}

		keyFields := buildFields(ev.Members, "key")
		dataFields := buildFields(ev.Members, "data")

		if containsBigInt(keyFields) || containsBigInt(dataFields) {
			needsBig = true
		}

		ctx.Events = append(ctx.Events, EventData{
			Name:       ev.Name,
			StructName: simplifyTypeName(ev.Name),
			Kind:       ev.Kind,
			Selector:   "",
			KeyFields:  keyFields,
			DataFields: dataFields,
		})
	}

	ctx.NeedsBigImport = needsBig

	tmpl := template.Must(template.
		New("events.tmpl").
		Funcs(template.FuncMap{
			"snake": camelToSnake,
		}).
		Parse(structsTmpl),
	)

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, ctx); err != nil {
		log.Fatalf("structs template execution failed: %v", err)
	}
	return buf.String()
}

func containsBigInt(fields []Field) bool {
	for _, f := range fields {
		if f.GoType == "big.Int" {
			return true
		}
	}
	return false
}

func camelToSnake(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_', r+('a'-'A'))
		} else {
			result = append(result, []rune(strings.ToLower(string(r)))...)
		}
	}
	return string(result)
}

func simplifyTypeName(full string) string {
	parts := strings.Split(full, "::")
	return parts[len(parts)-1]
}

func mapCairoTypeToGo(t string) string {
	switch t {
	case "core::felt252":
		return "string" // felt252 is typically represented as string in Go
	case "core::bool":
		return "bool"
	case "core::integer::u8", "core::integer::u16", "core::integer::u32", "core::integer::u64":
		return "uint64"
	case "core::integer::u256":
		return "big.Int"
	case "core::integer::u128":
		return "big.Int"
	case "alexandria_math::i257::i257":
		return "big.Int"
	default:
		return "string"
	}
}

func buildFields(members []CairoAbiMember, kind string) []Field {
	var fields []Field
	index := 0
	for _, m := range members {
		isKey := m.Kind == "key"
		if (kind == "key" && !isKey) || (kind == "data" && isKey) {
			continue
		}
		goType := mapCairoTypeToGo(m.Type)
		fields = append(fields, Field{
			Name:     utils.ToCamelCase(m.Name),
			GoType:   goType,
			IndexOff: index,
		})
		switch goType {
		case "types.U256", "types.U128":
			index += 2
		case "types.I257":
			index += 3
		default:
			index++
		}
	}
	return fields
}
