package utils

import (
	"fmt"
	"strings"
)

func NormalizeStarkNetAddress(addr string) string {
	addr = strings.TrimPrefix(addr, "0x")
	addr = strings.ToLower(addr)

	padded := fmt.Sprintf("%064s", addr)
	return "0x" + padded
}

func ToCamelCase(input string) string {
	parts := strings.Split(input, "_")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, "")
}
