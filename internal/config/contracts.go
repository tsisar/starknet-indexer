package config

import (
	"encoding/json"
	"github.com/tsisar/extended-log-go/log"
	"github.com/tsisar/starknet-indexer/internal/utils"
	"os"
	"path/filepath"
)

type ContractEntry struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type indexerConfig struct {
	Network    string                   `json:"network"`
	Indexer    string                   `json:"indexer"`
	StartBlock uint64                   `json:"start_block"`
	Contracts  map[string]ContractEntry `json:"contracts"`
}

func loadIndexerConfigFromEnv() indexerConfig {
	indexer := getString("INDEXER_NAME", "")
	if indexer == "" {
		log.Fatalf("INDEXER_NAME not set")
	}

	configPath := filepath.Join("config", "starknet", indexer+".json")
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("failed to read config file %s: %v", configPath, err)
	}

	var parsed indexerConfig
	if err := json.Unmarshal(data, &parsed); err != nil {
		log.Fatalf("failed to parse config file %s: %v", configPath, err)
	}

	for key, entry := range parsed.Contracts {
		entry.Address = utils.NormalizeStarkNetAddress(entry.Address)
		parsed.Contracts[key] = entry
	}

	return parsed
}
