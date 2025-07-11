// Code generated by generator-core; DO NOT EDIT.
package factory

import (
	"fmt"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/tsisar/starknet-indexer/internal/types"
)

func DecodeRoleGranted(keys, data []*felt.Felt) (interface{}, error) {
	if len(data) != 3 {
		return nil, fmt.Errorf("expected 3 data fields, got %d", len(data))
	}

	return RoleGranted{
		Data: RoleGrantedData{
			Role:    types.NewTextFelt(data[0]),
			Account: types.NewAddress(data[1]),
			Sender:  types.NewAddress(data[2]),
		},
	}, nil
}

func DecodeRoleRevoked(keys, data []*felt.Felt) (interface{}, error) {
	if len(data) != 3 {
		return nil, fmt.Errorf("expected 3 data fields, got %d", len(data))
	}

	return RoleRevoked{
		Data: RoleRevokedData{
			Role:    types.NewTextFelt(data[0]),
			Account: types.NewAddress(data[1]),
			Sender:  types.NewAddress(data[2]),
		},
	}, nil
}

func DecodeRoleAdminChanged(keys, data []*felt.Felt) (interface{}, error) {
	if len(data) != 3 {
		return nil, fmt.Errorf("expected 3 data fields, got %d", len(data))
	}

	return RoleAdminChanged{
		Data: RoleAdminChangedData{
			Role:              types.NewTextFelt(data[0]),
			PreviousAdminRole: types.NewTextFelt(data[1]),
			NewAdminRole:      types.NewTextFelt(data[2]),
		},
	}, nil
}

func DecodeFeeConfigUpdated(keys, data []*felt.Felt) (interface{}, error) {
	if len(keys) != 1 {
		return nil, fmt.Errorf("expected 1 keys, got %d", len(keys))
	}
	if len(data) != 2 {
		return nil, fmt.Errorf("expected 2 data fields, got %d", len(data))
	}

	return FeeConfigUpdated{
		Key: FeeConfigUpdatedKey{
			FeeRecipient: types.NewAddress(keys[0]),
		},
		Data: FeeConfigUpdatedData{
			FeeBps: types.U256{
				Low:  data[0].String(),
				High: data[1].String(),
			},
		},
	}, nil
}

func DecodeVaultDeployed(keys, data []*felt.Felt) (interface{}, error) {
	if len(keys) != 3 {
		return nil, fmt.Errorf("expected 3 keys, got %d", len(keys))
	}
	if len(data) != 6 {
		return nil, fmt.Errorf("expected 6 data fields, got %d", len(data))
	}

	return VaultDeployed{
		Key: VaultDeployedKey{
			Vault:      types.NewAddress(keys[0]),
			Asset:      types.NewAddress(keys[1]),
			Accountant: types.NewAddress(keys[2]),
		},
		Data: VaultDeployedData{
			ProfitMaxUnlockTime: data[0].Uint64(),
			Name:                types.NewTextFelt(data[1]),
			Symbol:              types.NewTextFelt(data[2]),
			Admin:               types.NewAddress(data[3]),
			AssetDecimals:       data[4].Uint64(),
			VaultDecimals:       data[5].Uint64(),
		},
	}, nil
}

func DecodeVaultClassHashUpdated(keys, data []*felt.Felt) (interface{}, error) {
	if len(keys) != 1 {
		return nil, fmt.Errorf("expected 1 keys, got %d", len(keys))
	}

	return VaultClassHashUpdated{
		Key: VaultClassHashUpdatedKey{
			NewClassHash: keys[0].String(),
		},
	}, nil
}

var Decoders = map[string]types.EventDecoderWithMeta{
	"0x9d4a59b844ac9d98627ddba326ab3707a7d7e105fd03c777569d0f61a91f1e":  {Func: DecodeRoleGranted, Name: "vaults::external::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleGranted"},
	"0x2842fd3b01bb0858fef6a2da51cdd9f995c7d36d7625fb68dd5d69fcc0a6d76": {Func: DecodeRoleRevoked, Name: "vaults::external::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleRevoked"},
	"0x2b23b0c08c7b22209aea4100552de1b7876a49f04ee5a4d94f83ad24bc4ec1c": {Func: DecodeRoleAdminChanged, Name: "vaults::external::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleAdminChanged"},
	"0x2dff38649b3190f61e2b96eb6a11d725e104dc2405365e5600b94b2013be9f8": {Func: DecodeFeeConfigUpdated, Name: "vaults::factory::Factory::FeeConfigUpdated"},
	"0x1049e6cdd616e547ac9d1ce9afb908baf99403775a5e96af0dcb28c73d02c8c": {Func: DecodeVaultDeployed, Name: "vaults::factory::Factory::VaultDeployed"},
	"0xb0b5839fe46a401509d701ea9b626f4a17e655d0c0249b11ab17f7c4d872b8":  {Func: DecodeVaultClassHashUpdated, Name: "vaults::factory::Factory::VaultClassHashUpdated"},
}
