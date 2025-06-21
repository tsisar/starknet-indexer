package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
)

type U128 struct {
	Low  string `json:"low"`
	High string `json:"high"`
}

func (u U128) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.ToBigInt())
}

func (u U128) String() string {
	return u.ToBigInt().String()
}

func (u U128) ToBigInt() *big.Int {
	low := new(big.Int)
	high := new(big.Int)

	if _, ok := low.SetString(strings.TrimPrefix(u.Low, "0x"), 16); !ok {
		panic(fmt.Sprintf("invalid low hex: %s", u.Low))
	}
	if _, ok := high.SetString(strings.TrimPrefix(u.High, "0x"), 16); !ok {
		panic(fmt.Sprintf("invalid high hex: %s", u.High))
	}

	high.Lsh(high, 64)
	return high.Or(high, low)
}
