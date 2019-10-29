package gostringconverters

import (
	"strings"
	"unicode"
)

// RuneMapConversion - given a string and a map, converts every rune in string to corresponding map value
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
