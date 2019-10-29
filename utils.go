package gostringconverters

import (
	"strings"
	"unicode"
)

// ConvertToSeparatedCase - given a string and SeparatedCaseOptions, modifies the string case according to those options
// Handles other separated cases (like snake_case) and uppercase-demarcated cases (like PascalCase)
func ConvertToSeparatedCase(s string, opts SeparatedCaseOptions) string {
	var output strings.Builder

	capsFunc := unicode.ToLower
	if opts.IsScreaming {
		capsFunc = unicode.ToUpper
	}
	prevIsSeparator := true
	hasWrittenChar := false
	for _, r := range s {
		if IsSeparator(r) {
			prevIsSeparator = true
			continue
		}
		if (unicode.IsUpper(r) && !prevIsSeparator) || (prevIsSeparator && hasWrittenChar) {
			output.WriteRune(opts.Separator)
		}
		output.WriteRune(capsFunc(r))
		hasWrittenChar = true
		prevIsSeparator = false

	}
	return output.String()
}

// IsSeparator - checks if rune is a text separator type (punctuation or space)
func IsSeparator(r rune) bool {
	return unicode.IsPunct(r) || unicode.IsSpace(r)
}

// RuneMapToMapFunc - converts a map[rune]rune into a func that can be called by strings.Map
func RuneMapToMapFunc(runeMap map[rune]rune) func(rune) rune {
	return func(r rune) rune {
		_, ok := runeMap[r]
		if !ok {
			return r
		}
		return runeMap[r]
	}
}
