package abiutil

import (
	"bytes"
	_ "embed"
	"github.com/tsisar/starknet-indexer/internal/utils"
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
	Package string
	Events  []EventData
}

func GenerateGoStructs(events []CairoAbiEntry, packageName string) string {
	ctx := TemplateContext{Package: packageName}
	for _, ev := range events {
		if ev.Kind != "struct" {
			continue
		}
		ctx.Events = append(ctx.Events, EventData{
			Name:       ev.Name,
			StructName: simplifyTypeName(ev.Name),
			Kind:       ev.Kind,
			Selector:   "",
			KeyFields:  buildFields(ev.Members, "key"),
			DataFields: buildFields(ev.Members, "data"),
		})
	}

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
	case "core::starknet::contract_address::ContractAddress":
		return "types.Address"
	case "core::felt252":
		return "types.TextFelt"
	case "core::bool":
		return "bool"
	case "core::integer::u8", "core::integer::u16", "core::integer::u32", "core::integer::u64":
		return "uint64"
	case "core::integer::u256":
		return "types.U256"
	case "core::integer::u128":
		return "types.U128"
	case "alexandria_math::i257::i257":
		return "types.I257"
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
