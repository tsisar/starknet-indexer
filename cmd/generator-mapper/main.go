package main

import (
	"bytes"
	"github.com/Tsisar/extended-log-go/log"
	"github.com/Tsisar/starknet-indexer/cmd/generator-mapper/abiutil"
	"github.com/Tsisar/starknet-indexer/internal/config"
	"github.com/Tsisar/starknet-indexer/internal/utils"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	contracts := config.App.Contracts
	subgraphDir := "generated/subgraph"
	mapperDir := "generated/registry"

	var contractPackages []string

	for name := range contracts {
		cleanName := strings.ToLower(strings.TrimPrefix(name, "CONTRACT_"))
		contractPackages = append(contractPackages, cleanName)

		abiPath := filepath.Join("abi", cleanName+".json")

		raw, err := os.ReadFile(abiPath)
		if err != nil {
			log.Warnf("Failed to read %s: %v", abiPath, err)
			continue
		}

		entries, err := abiutil.ParseAbi(raw)
		if err != nil {
			log.Warnf("Invalid ABI %s: %v", abiPath, err)
			continue
		}

		events := abiutil.ExtractEventTypes(entries)

		contractDir := filepath.Join(subgraphDir, cleanName)
		if err := os.MkdirAll(contractDir, 0755); err != nil {
			log.Fatalf("Failed to create directory for %s: %v", cleanName, err)
		}

		code := abiutil.GenerateGoStructs(events, cleanName)
		if err := os.WriteFile(filepath.Join(contractDir, "events.go"), []byte(code), 0644); err != nil {
			log.Fatalf("Failed to write event file for %s: %v", cleanName, err)
		}

		if err := os.MkdirAll(mapperDir, 0755); err != nil {
			log.Fatalf("Failed to create directory for %s: %v", cleanName, err)
		}

		registryCode := abiutil.GeneratePerContractRegistry(cleanName, events)
		if err := os.WriteFile(filepath.Join(mapperDir, cleanName+".go"), []byte(registryCode), 0644); err != nil {
			log.Fatalf("Failed to write registry file for %s: %v", cleanName, err)
		}

		log.Debugf("Generated for contract: %s", cleanName)
	}

	if err := generateMainMapper(contractPackages); err != nil {
		log.Fatalf("Failed to generate main mapper: %v", err)
	}
}

func generateMainMapper(packages []string) error {
	tmplFile := "cmd/generator-mapper/abiutil/templates/registry_map.tmpl"
	tmplRaw, err := os.ReadFile(tmplFile)
	if err != nil {
		return err
	}

	tmpl := template.Must(template.New("mapper").
		Funcs(template.FuncMap{
			"toCamelCase": utils.ToCamelCase,
		}).
		Parse(string(tmplRaw)))

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, struct {
		Packages []string
	}{Packages: packages}); err != nil {
		return err
	}

	outFile := "generated/registry/registry_map.go"
	return os.WriteFile(outFile, buf.Bytes(), 0644)
}
