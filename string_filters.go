package gostringconverters

import (
	"strings"
	"unicode"
)

// FilterStringChars - given a string, removes all characters meeting filter conditions
func FilterStringChars(s string, filters []RuneFilterFunc) string {
	var result strings.Builder
	for _, r := range s {
		if !runeIsFiltered(r, filters) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func runeIsFiltered(r rune, filters []RuneFilterFunc) bool {
	for _, filter := range filters {
		if filter(r) {
			return true
		}
	}
	return false
}

// RuneFilterFunc - function that accepts a rune and returns a true if the rune meets a given condition
type RuneFilterFunc func(rune) bool

// ControlCharFilterFunc - checks if rune is a control character
func ControlCharFilterFunc() RuneFilterFunc {
	return func(r rune) bool {
		return unicode.IsControl(r)
	}
}

// RemoveControlCharsFromString - wraps FilterStringByRune to remove all control characters from a string
func RemoveControlCharsFromString(s string) string {
	return FilterStringChars(s, []RuneFilterFunc{
		ControlCharFilterFunc(),
	})
}
