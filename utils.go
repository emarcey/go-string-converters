package gostringconverters

import (
	"strings"
)

func RuneMapConversion(s string, m map[rune]rune) string {
	var output strings.Builder
	for _, r := range s {
		l, ok := m[r]
		if !ok {
			output.WriteRune(r)
			continue
		}
		output.WriteRune(l)
	}
	return output.String()
}
