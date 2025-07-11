// Code generated by generator-core; DO NOT EDIT.
package access_control_manager

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

var Decoders = map[string]types.EventDecoderWithMeta{
	"0x9d4a59b844ac9d98627ddba326ab3707a7d7e105fd03c777569d0f61a91f1e":  {Func: DecodeRoleGranted, Name: "oracle_network::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleGranted"},
	"0x2842fd3b01bb0858fef6a2da51cdd9f995c7d36d7625fb68dd5d69fcc0a6d76": {Func: DecodeRoleRevoked, Name: "oracle_network::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleRevoked"},
	"0x2b23b0c08c7b22209aea4100552de1b7876a49f04ee5a4d94f83ad24bc4ec1c": {Func: DecodeRoleAdminChanged, Name: "oracle_network::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleAdminChanged"},
}
