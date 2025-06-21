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

	if handler, ok := r[event.EventName]; ok {
		return handler(ctx, client, event)
	}
	log.Warnf("No mapping implemented for event: %s", event.EventName)
	return nil
}

func logEvent(event model.Event) {
	entry := map[string]interface{}{
		"contract_name":    event.Contract.Name,
		"contract_address": event.Contract.Address,
		"event_name":       event.EventName,
		"event_index":      event.EventIndex,
		"event":            event.RawData,
		"block_number":     event.BlockNumber,
		"block_timestamp":  event.Timestamp,
		"block_hash":       event.Block.Hash,
		"tx_hash":          event.Transaction.Hash,
	}

	jsonBytes, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		log.Errorf("Failed to marshal decoded event %s: %v", event.EventName, err)
		return
	}

	log.Infof("[mapper] Decoded event:\n%s", string(jsonBytes))
}
