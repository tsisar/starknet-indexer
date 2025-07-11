// Code generated by generator-core; DO NOT EDIT.
package price_feed_wbtc

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

func DecodeLogSetOracle(keys, data []*felt.Felt) (interface{}, error) {
	if len(keys) != 1 {
		return nil, fmt.Errorf("expected 1 keys, got %d", len(keys))
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("expected 1 data fields, got %d", len(data))
	}

	return LogSetOracle{
		Key: LogSetOracleKey{
			Caller: types.NewAddress(keys[0]),
		},
		Data: LogSetOracleData{
			Oracle: types.NewAddress(data[0]),
		},
	}, nil
}

func DecodeLogSetTimeDelay(keys, data []*felt.Felt) (interface{}, error) {
	if len(keys) != 1 {
		return nil, fmt.Errorf("expected 1 keys, got %d", len(keys))
	}
	if len(data) != 2 {
		return nil, fmt.Errorf("expected 2 data fields, got %d", len(data))
	}

	return LogSetTimeDelay{
		Key: LogSetTimeDelayKey{
			Caller: types.NewAddress(keys[0]),
		},
		Data: LogSetTimeDelayData{
			TimeDelay: types.U256{
				Low:  data[0].String(),
				High: data[1].String(),
			},
		},
	}, nil
}

func DecodeLogPeekPriceFailed(keys, data []*felt.Felt) (interface{}, error) {
	if len(keys) != 1 {
		return nil, fmt.Errorf("expected 1 keys, got %d", len(keys))
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("expected 1 data fields, got %d", len(data))
	}

	return LogPeekPriceFailed{
		Key: LogPeekPriceFailedKey{
			Caller: types.NewAddress(keys[0]),
		},
		Data: LogPeekPriceFailedData{
			Reason: data[0].String(),
		},
	}, nil
}

func DecodeLogSetPriceLife(keys, data []*felt.Felt) (interface{}, error) {
	if len(keys) != 1 {
		return nil, fmt.Errorf("expected 1 keys, got %d", len(keys))
	}
	if len(data) != 2 {
		return nil, fmt.Errorf("expected 2 data fields, got %d", len(data))
	}

	return LogSetPriceLife{
		Key: LogSetPriceLifeKey{
			Caller: types.NewAddress(keys[0]),
		},
		Data: LogSetPriceLifeData{
			PriceLife: types.U256{
				Low:  data[0].String(),
				High: data[1].String(),
			},
		},
	}, nil
}

func DecodeLogSetPoolId(keys, data []*felt.Felt) (interface{}, error) {
	if len(keys) != 1 {
		return nil, fmt.Errorf("expected 1 keys, got %d", len(keys))
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("expected 1 data fields, got %d", len(data))
	}

	return LogSetPoolId{
		Key: LogSetPoolIdKey{
			Caller: types.NewAddress(keys[0]),
		},
		Data: LogSetPoolIdData{
			PoolId: types.NewTextFelt(data[0]),
		},
	}, nil
}

func DecodeLogSetPrice(keys, data []*felt.Felt) (interface{}, error) {
	if len(keys) != 1 {
		return nil, fmt.Errorf("expected 1 keys, got %d", len(keys))
	}
	if len(data) != 2 {
		return nil, fmt.Errorf("expected 2 data fields, got %d", len(data))
	}

	return LogSetPrice{
		Key: LogSetPriceKey{
			Caller: types.NewAddress(keys[0]),
		},
		Data: LogSetPriceData{
			Price: types.U256{
				Low:  data[0].String(),
				High: data[1].String(),
			},
		},
	}, nil
}

func DecodeLogSetAccessControlConfig(keys, data []*felt.Felt) (interface{}, error) {
	if len(keys) != 1 {
		return nil, fmt.Errorf("expected 1 keys, got %d", len(keys))
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("expected 1 data fields, got %d", len(data))
	}

	return LogSetAccessControlConfig{
		Key: LogSetAccessControlConfigKey{
			Caller: types.NewAddress(keys[0]),
		},
		Data: LogSetAccessControlConfigData{
			AccessControlConfig: types.NewAddress(data[0]),
		},
	}, nil
}

var Decoders = map[string]types.EventDecoderWithMeta{
	"0x2eb5248cf3d8cd81a5ba6d3cc6e1997df7b174eb894aac081867c1a2bc43c90": {Func: DecodePaused, Name: "stablecoin::components::pausable::PausableComponent::Paused"},
	"0xece5baf71f670bcb771481fd7bd9efd6d6b8053246fe67b5a13db8bf5f50f1":  {Func: DecodeUnpaused, Name: "stablecoin::components::pausable::PausableComponent::Unpaused"},
	"0x26106c38f17534bb06d5ea9485d0f06313ec1b1fa61d3e434e4aeda3ac850e6": {Func: DecodeLogSetOracle, Name: "stablecoin::core::price::price_feed::PriceFeed::LogSetOracle"},
	"0x5f3b9e56ce7ea42cf3224e2925aa027fbbbf31406dc8c6e1616f884ac11f1":   {Func: DecodeLogSetTimeDelay, Name: "stablecoin::core::price::price_feed::PriceFeed::LogSetTimeDelay"},
	"0x717f102da503614744f8c80783267c2a4f4761cf628a29a8a27b102a0ca9cf":  {Func: DecodeLogPeekPriceFailed, Name: "stablecoin::core::price::price_feed::PriceFeed::LogPeekPriceFailed"},
	"0xfad4a5ec66d2cf032f70b7d8658298eb66a8c011d6809d95fd57b3a0a67d1c":  {Func: DecodeLogSetPriceLife, Name: "stablecoin::core::price::price_feed::PriceFeed::LogSetPriceLife"},
	"0xc3f6fbdd9cdf2ab936cc809294d5df120a1cc2e2b5980b979e235885e4e5b2":  {Func: DecodeLogSetPoolId, Name: "stablecoin::core::price::price_feed::PriceFeed::LogSetPoolId"},
	"0x12afa7db819f5a23f7890708cc38ac2806814f52c923102f9fcb6c0ff6f5f87": {Func: DecodeLogSetPrice, Name: "stablecoin::core::price::price_feed::PriceFeed::LogSetPrice"},
	"0x299667613c997fd6fc15c51a5c3b55e66701061d593bfa04867d1dd048a0f40": {Func: DecodeLogSetAccessControlConfig, Name: "stablecoin::core::price::price_feed::PriceFeed::LogSetAccessControlConfig"},
}
