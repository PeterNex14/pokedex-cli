package main

import (
	"strings"
)

func cleanInput(text string) []string {
	texts := strings.Fields(strings.ToLower(text))

	return texts
}