package types

import (
	"encoding/json"
	"math/big"
)

type I257 struct {
	Abs        U256 `json:"abs"`
	IsNegative bool `json:"is_negative"`
}

// MarshalJSON implements the json.Marshaler interface.
func (i I257) MarshalJSON() ([]byte, error) {
	s := i.ToBigInt()
	return json.Marshal(s)
}

// MarshalText implements the encoding.TextMarshaler interface.
// Useful for encoding in plain text formats (e.g., CSV, YAML).
func (i I257) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// String returns the signed string representation of the number.
func (i I257) String() string {
	absInt := i.Abs.ToBigInt()
	if i.IsNegative {
		return "-" + absInt.String()
	}
	return absInt.String()
}

// ToBigInt returns the signed big.Int representation of the number.
func (i I257) ToBigInt() *big.Int {
	absInt := i.Abs.ToBigInt()
	if i.IsNegative {
		return absInt.Neg(absInt)
	}
	return absInt
}
