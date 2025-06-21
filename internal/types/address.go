package types

import (
	"encoding/json"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/Tsisar/starknet-indexer/internal/utils"
)

type Address struct {
	Value string
}

func NewAddress(f *felt.Felt) Address {
	address := utils.NormalizeStarkNetAddress(f.String())
	return Address{Value: address}
}

func (t Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Value)
}

func (t Address) String() string {
	return t.Value
}
