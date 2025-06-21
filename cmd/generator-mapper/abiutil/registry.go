package abiutil

import (
	"bytes"
	_ "embed"
	"github.com/tsisar/starknet-indexer/internal/utils"
	"text/template"
)

//go:embed templates/registry.tmpl
var registryTmplStr string

var registryTmpl = template.Must(template.New("registry").
	Funcs(template.FuncMap{
		"toCamelCase": utils.ToCamelCase,
	}).
	Parse(registryTmplStr))

type RegistryEvent struct {
	EventName   string
	StructName  string
	PackageName string
}

func GeneratePerContractRegistry(packageName string, events []CairoAbiEntry) string {
	ctx := struct {
		PackageName string
		Events      []RegistryEvent
	}{
		PackageName: packageName,
	}

	for _, e := range events {
		if e.Kind != "struct" {
			continue
		}
		name := simplifyTypeName(e.Name)
		ctx.Events = append(ctx.Events, RegistryEvent{
			EventName:   e.Name,
			StructName:  name,
			PackageName: packageName,
		})
	}

	var buf bytes.Buffer
	if err := registryTmpl.Execute(&buf, ctx); err != nil {
		panic(err)
	}
	return buf.String()
}
