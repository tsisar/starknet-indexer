// Code generated by generator-core; DO NOT EDIT.
package aggregator_wbtc

import (
	"fmt"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/tsisar/starknet-indexer/internal/types"
)

func DecodeLogNodeAdded(keys, data []*felt.Felt) (interface{}, error) {
	if len(keys) != 1 {
		return nil, fmt.Errorf("expected 1 keys, got %d", len(keys))
	}

	return LogNodeAdded{
		Key: LogNodeAddedKey{
			Node: types.NewAddress(keys[0]),
		},
	}, nil
}

func DecodeLogNodeRemoved(keys, data []*felt.Felt) (interface{}, error) {
	if len(keys) != 1 {
		return nil, fmt.Errorf("expected 1 keys, got %d", len(keys))
	}

	return LogNodeRemoved{
		Key: LogNodeRemovedKey{
			Node: types.NewAddress(keys[0]),
		},
	}, nil
}

func DecodeLogAnswerPublished(keys, data []*felt.Felt) (interface{}, error) {
	if len(keys) != 3 {
		return nil, fmt.Errorf("expected 3 keys, got %d", len(keys))
	}

	return LogAnswerPublished{
		Key: LogAnswerPublishedKey{
			Value: types.U256{
				Low:  keys[0].String(),
				High: keys[1].String(),
			},
			Timestamp: keys[2].Uint64(),
		},
	}, nil
}

var Decoders = map[string]types.EventDecoderWithMeta{
	"0x2da5e931b1dddc705b0d73b5f9e133b726f4a94a80d4788cc232a42c2111d56": {Func: DecodeLogNodeAdded, Name: "oracle_network::core::aggregator::Aggregator::LogNodeAdded"},
	"0x2ca196df22a12166c0a879ddf856cf0753d6468edc0f9b83a07d1f706cf3a05": {Func: DecodeLogNodeRemoved, Name: "oracle_network::core::aggregator::Aggregator::LogNodeRemoved"},
	"0x30fbd47556a86d3eabef15936171fb51b08efa2380646fa3966add50025f73b": {Func: DecodeLogAnswerPublished, Name: "oracle_network::core::aggregator::Aggregator::LogAnswerPublished"},
}
