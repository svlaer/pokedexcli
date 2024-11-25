package main

import (
	"fmt"
	"strings"
)

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  map[string]int
	Types  []string
}

func (p Pokemon) String() string {
	return fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\nStats: %s\nTypes: %s", p.Name, p.Height, p.Weight, formatStats(p.Stats), formatTypes(p.Types))
}

func formatStats(stats map[string]int) string {
	// Possible TODO: Sort keys before printing, to ensure consistent order
	var b strings.Builder

	for k, v := range stats {
		fmt.Fprintf(&b, "\n\t- %s: %d", k, v)
	}

	return b.String()
}

func formatTypes(types []string) string {
	var b strings.Builder

	for _, v := range types {
		fmt.Fprintf(&b, "\n\t- %s", v)
	}

	return b.String()
}
