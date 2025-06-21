package abiutil

import (
	"encoding/json"
	"fmt"
)

type CairoAbiEntry struct {
	Type     string            `json:"type"`
	Name     string            `json:"name"`
	Kind     string            `json:"kind,omitempty"`
	Variants []CairoAbiVariant `json:"variants,omitempty"`
	Members  []CairoAbiMember  `json:"members,omitempty"`
}

type CairoAbiVariant struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Kind string `json:"kind"`
}

type CairoAbiMember struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Kind string `json:"kind"`
}

func ParseAbi(raw []byte) ([]CairoAbiEntry, error) {
	var container any
	if err := json.Unmarshal(raw, &container); err != nil {
		return nil, fmt.Errorf("failed to unmarshal outer ABI: %w", err)
	}

	switch v := container.(type) {
	case string:
		// raw stringified JSON → decode again
		var entries []CairoAbiEntry
		if err := json.Unmarshal([]byte(v), &entries); err != nil {
			return nil, fmt.Errorf("failed to unmarshal inner ABI: %w", err)
		}
		return entries, nil

	case []any:
		// already correct format → convert directly
		var entries []CairoAbiEntry
		raw2, _ := json.Marshal(v)
		if err := json.Unmarshal(raw2, &entries); err != nil {
			return nil, fmt.Errorf("failed to unmarshal ABI array: %w", err)
		}
		return entries, nil

	default:
		return nil, fmt.Errorf("unsupported ABI format: %T", v)
	}
}

func ExtractEventTypes(entries []CairoAbiEntry) []CairoAbiEntry {
	var events []CairoAbiEntry

	for _, e := range entries {
		if e.Type != "event" {
			continue
		}

		var argTypes []string
		for _, m := range e.Members {
			argTypes = append(argTypes, m.Type)
		}

		events = append(events, e)
	}

	return events
}
