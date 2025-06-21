package types

import (
	"encoding/json"
	"github.com/NethermindEth/juno/core/felt"
)

type TextFelt struct {
	Value string
}

func NewTextFelt(f *felt.Felt) TextFelt {
	raw := f.Bytes()
	bytes := make([]byte, 32)
	copy(bytes, raw[:])

	start := 0
	end := len(bytes)

	for start < end && bytes[start] == 0 {
		start++
	}
	for end > start && bytes[end-1] == 0 {
		end--
	}

	trimmed := bytes[start:end]
	str := string(trimmed)

	return TextFelt{Value: str}
}

func (t TextFelt) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Value)
}

func (t TextFelt) String() string {
	return t.Value
}
