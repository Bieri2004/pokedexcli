package main

import (
	"strings"
)

func cleanInput(text string) []string {
	// Trim leading and trailing whitespace
	trimmed := strings.TrimSpace(text)

	// If empty after trimming, return empty slice
	if trimmed == "" {
		return []string{}
	}

	// Convert to lowercase
	lowered := strings.ToLower(trimmed)

	// Split on whitespace (Fields handles multiple spaces)
	words := strings.Fields(lowered)

	return words
}
