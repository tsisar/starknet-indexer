package types

import "github.com/NethermindEth/juno/core/felt"

type EventDecoderWithMeta struct {
	Func func([]*felt.Felt, []*felt.Felt) (interface{}, error)
	Name string
}

func (e *EventDecoderWithMeta) FullName() string {
	if len(e.Name) > 0 {
		return e.Name
	}
	return "unknown"
}
