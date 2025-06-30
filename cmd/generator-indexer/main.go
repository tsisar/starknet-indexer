package main

import (
	"fmt"
	"github.com/tsisar/extended-log-go/log"
	"github.com/tsisar/starknet-indexer/cmd/generator-indexer/abiutil"
	"github.com/tsisar/starknet-indexer/internal/config"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	contracts := config.App.IndexerConfig.Contracts
	rpcURL := config.App.RPCEndpoint

	var contractNames []string

	for name, contract := range contracts {
		cleanName := strings.ToLower(name)
		contractNames = append(contractNames, cleanName)

		abiPath := filepath.Join("abi", cleanName+".json")

		fmt.Println("---------------------------------")
		log.Debugf("Fetching ABI for %s (%s)", name, contract.Address)
		err := abiutil.FetchAndSaveABI(cleanName, contract.Address, rpcURL, "abi")
		if err != nil {
			log.Warnf("Failed to fetch ABI for %s: %v", name, err)
			continue
		}

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

		contractDir := filepath.Join("generated/indexer", cleanName)
		if err := os.MkdirAll(contractDir, 0755); err != nil {
			log.Fatalf("Failed to create directory for %s: %v", cleanName, err)
		}

		code := abiutil.GenerateGoStructs(events, cleanName)
		if err := os.WriteFile(filepath.Join(contractDir, "events.go"), []byte(code), 0644); err != nil {
			log.Fatalf("Failed to write event file for %s: %v", cleanName, err)
		}
		log.Debugf("Generated Go structs for %s with %d events", cleanName, len(events))

		decoderCode := abiutil.GenerateDecoderFunctions(events, cleanName)
		if err := os.WriteFile(filepath.Join(contractDir, "decoders.go"), []byte(decoderCode), 0644); err != nil {
			log.Fatalf("Failed to write decoder file for %s: %v", cleanName, err)
		}
		log.Debugf("Generated decoders for %s", cleanName)
	}

	registry := abiutil.GenerateRegistryFile(contractNames)
	if err := os.WriteFile("generated/indexer/registry.go", []byte(registry), 0644); err != nil {
		log.Fatalf("Failed to write registry file: %v", err)
	}
	log.Infof("Generated ABI registry with %d contracts", len(contractNames))
}
