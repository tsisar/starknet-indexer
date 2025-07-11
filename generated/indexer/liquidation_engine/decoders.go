// Code generated by generator-core; DO NOT EDIT.
package liquidation_engine

import (
	"fmt"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/tsisar/starknet-indexer/internal/types"
)

func DecodePaused(keys, data []*felt.Felt) (interface{}, error) {
	if len(data) != 1 {
		return nil, fmt.Errorf("expected 1 data fields, got %d", len(data))
	}

	return Paused{
		Data: PausedData{
			Account: types.NewAddress(data[0]),
		},
	}, nil
}

func DecodeUnpaused(keys, data []*felt.Felt) (interface{}, error) {
	if len(data) != 1 {
		return nil, fmt.Errorf("expected 1 data fields, got %d", len(data))
	}

	return Unpaused{
		Data: UnpausedData{
			Account: types.NewAddress(data[0]),
		},
	}, nil
}

func DecodeLiquidationFail(keys, data []*felt.Felt) (interface{}, error) {
	if len(keys) != 1 {
		return nil, fmt.Errorf("expected 1 keys, got %d", len(keys))
	}
	if len(data) != 3 {
		return nil, fmt.Errorf("expected 3 data fields, got %d", len(data))
	}

	return LiquidationFail{
		Key: LiquidationFailKey{
			Liquidator: types.NewAddress(keys[0]),
		},
		Data: LiquidationFailData{
			CollateralPoolIds: data[0].String(),
			PositionAddresses: data[1].String(),
			Reason:            types.NewTextFelt(data[2]),
		},
	}, nil
}

func DecodeLogSetBookKeeper(keys, data []*felt.Felt) (interface{}, error) {
	if len(data) != 1 {
		return nil, fmt.Errorf("expected 1 data fields, got %d", len(data))
	}

	return LogSetBookKeeper{
		Data: LogSetBookKeeperData{
			NewAddress: types.NewAddress(data[0]),
		},
	}, nil
}

func DecodeLogAddToWhitelist(keys, data []*felt.Felt) (interface{}, error) {
	if len(keys) != 1 {
		return nil, fmt.Errorf("expected 1 keys, got %d", len(keys))
	}

	return LogAddToWhitelist{
		Key: LogAddToWhitelistKey{
			User: types.NewAddress(keys[0]),
		},
	}, nil
}

func DecodeLogRemoveFromWhitelist(keys, data []*felt.Felt) (interface{}, error) {
	if len(keys) != 1 {
		return nil, fmt.Errorf("expected 1 keys, got %d", len(keys))
	}

	return LogRemoveFromWhitelist{
		Key: LogRemoveFromWhitelistKey{
			User: types.NewAddress(keys[0]),
		},
	}, nil
}

func DecodeLogCage(keys, data []*felt.Felt) (interface{}, error) {

	return LogCage{}, nil
}

var Decoders = map[string]types.EventDecoderWithMeta{
	"0x2eb5248cf3d8cd81a5ba6d3cc6e1997df7b174eb894aac081867c1a2bc43c90": {Func: DecodePaused, Name: "stablecoin::components::pausable::PausableComponent::Paused"},
	"0xece5baf71f670bcb771481fd7bd9efd6d6b8053246fe67b5a13db8bf5f50f1":  {Func: DecodeUnpaused, Name: "stablecoin::components::pausable::PausableComponent::Unpaused"},
	"0x176d4e5ebb4a804704f29ceaa88e54c90680e32ea27dd3ffffe93db405fe3f5": {Func: DecodeLiquidationFail, Name: "stablecoin::core::liquidation_engine::LiquidationEngine::LiquidationFail"},
	"0x1ff221b30a3b491affd2f37d096cf0e8d5565722b8ff03d4751f1c2f067083f": {Func: DecodeLogSetBookKeeper, Name: "stablecoin::core::liquidation_engine::LiquidationEngine::LogSetBookKeeper"},
	"0x1338d8637e5f1d9f0538699896d04ece348accfaac81d0924f421b1cfbfdfec": {Func: DecodeLogAddToWhitelist, Name: "stablecoin::core::liquidation_engine::LiquidationEngine::LogAddToWhitelist"},
	"0x32176b0461878b055af0a933e1c41f4a0d653a05fd879ca66fc1f39baa2f79":  {Func: DecodeLogRemoveFromWhitelist, Name: "stablecoin::core::liquidation_engine::LiquidationEngine::LogRemoveFromWhitelist"},
	"0x75ef26d8af3732b337ded7c33086749b36a2c1ec660f2de592fe60717edfd":   {Func: DecodeLogCage, Name: "stablecoin::core::liquidation_engine::LiquidationEngine::LogCage"},
}
