package abiutil

import (
	"bytes"
	_ "embed"
	"log"
	"text/template"
)

//go:embed templates/registry.tmpl
var registryTmpl string

type registryContext struct {
	Package   string
	Contracts []string
}

func GenerateRegistryFile(contractNames []string) string {
	ctx := registryContext{
		Package:   "generated",
		Contracts: contractNames,
	}

	tmpl := template.Must(template.
		New("registry.tmpl").
		Parse(registryTmpl),
	)

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, ctx); err != nil {
		log.Fatalf("registry template execution failed: %v", err)
	}
	return buf.String()
}
