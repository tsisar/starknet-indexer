package utils

import (
	"fmt"
	"math/big"
	"strconv"
)

var (
	WAD = big.NewInt(1e18)
	RAY = mustBigInt("1000000000000000000000000000")
	RAD = mustBigInt("1000000000000000000000000000000000000000000000")

	wadFloat = new(big.Float).SetInt(WAD)
	rayFloat = new(big.Float).SetInt(RAY)
	radFloat = new(big.Float).SetInt(RAD)
	Zero     = new(big.Float).SetInt64(0)
)

func mustBigInt(s string) *big.Int {
	n, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("invalid constant: " + s)
	}
	return n
}

// DivByWADToDecimal: BigInt / WAD → BigFloat
func DivByWADToDecimal(value *big.Int) *big.Float {
	v := new(big.Float).SetInt(value)
	return new(big.Float).Quo(v, wadFloat)
}

// DivByRAYToDecimal: BigInt / (WAD * 1e9) → BigFloat
func DivByRAYToDecimal(value *big.Int) *big.Float {
	v := new(big.Float).SetInt(value)
	denominator := new(big.Float).Mul(wadFloat, big.NewFloat(1e9))
	return new(big.Float).Quo(v, denominator)
}

// DivByRADToDecimal: BigInt / (WAD * WAD * 1e9) → BigFloat
func DivByRADToDecimal(value *big.Int) *big.Float {
	v := new(big.Float).SetInt(value)
	denominator := new(big.Float).Mul(wadFloat, wadFloat)
	denominator.Mul(denominator, big.NewFloat(1e9))
	return new(big.Float).Quo(v, denominator)
}

// DivByRAD: BigInt / (WAD * WAD * 1e9) → BigInt
func DivByRAD(value *big.Int) *big.Int {
	tmp := new(big.Int).Div(value, WAD)
	tmp.Div(tmp, WAD)
	tmp.Div(tmp, big.NewInt(1e9))
	return tmp
}

// DivByRAY: BigInt / (WAD * 1e9) → BigInt
func DivByRAY(value *big.Int) *big.Int {
	tmp := new(big.Int).Div(value, WAD)
	tmp.Div(tmp, big.NewInt(1e9))
	return tmp
}

func IsZero(v *big.Float) bool {
	return v.Cmp(Zero) == 0
}

func NormalizeHash(s string) string {
	if len(s) >= 2 && s[:2] == "0x" {
		return s[2:]
	}
	return s
}

// ToDecimalString converts big.Float to string (for ent)
func ToDecimalString(f *big.Float) string {
	return f.Text('f', 18)
}

func Uint64ToString(val uint64) string {
	return strconv.FormatUint(val, 10)
}

func EncodeStringToHex(s string) string {
	bytes := []byte(s)
	result := "0x"
	for _, b := range bytes {
		result += fmt.Sprintf("%02X", b)
	}
	return result
}
