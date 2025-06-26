package mapper

import (
	"context"
	"encoding/json"
	"github.com/tsisar/extended-log-go/log"
	"github.com/tsisar/starknet-indexer/generated/ent"
	"github.com/tsisar/starknet-indexer/generated/registry"
	"github.com/tsisar/starknet-indexer/internal/storage/model"
)

func MapEvents(ctx context.Context, client *ent.Client, event model.Event) error {
	logEvent(event)

	contract := event.Contract.Name // contract name as package name
	r, ok := registry.RegistryMap[contract]
	if !ok {
		log.Warnf("No r found for contract: %s", contract)
		return nil
	}

	if handler, ok := r[event.Name]; ok {
		return handler(ctx, client, event)
	}
	log.Warnf("No mapping implemented for event: %s", event.Name)
	return nil
}

func logEvent(event model.Event) {
	entry := map[string]interface{}{
		"contract_name":    event.Contract.Name,
		"contract_address": event.Contract.Address,
		"name":             event.Name,
		"index":            event.Index,
		"event":            event.JsonEv,
		"block_number":     event.BlockNumber,
		"block_timestamp":  event.Timestamp,
		"block_hash":       event.Block.Hash,
		"tx_hash":          event.Transaction.Hash,
	}

	jsonBytes, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		log.Errorf("Failed to marshal decoded event %s: %v", event.Name, err)
		return
	}

	log.Infof("[mapper] Decoded event:\n%s", string(jsonBytes))
}
