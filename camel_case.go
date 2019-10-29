package gostringconverters

import (
	"strings"
	"unicode"
)

// CamelCase - Punctuation Aggressive CamelCase converter.
// Will treat any space or punctuation as a separator
func CamelCase(s string) string {
	// this is a string -> thisIsAString
	var output strings.Builder
	nextUpper := false
	isFirst := true
	for _, r := range s {
		if IsSeparator(r) {
			if !isFirst {
				nextUpper = true
			}
			continue
		}
		if !unicode.IsLetter(r) {
			output.WriteRune(r)
			continue
		}
		if nextUpper {
			output.WriteRune(unicode.ToUpper(r))
			nextUpper = false
			continue
		}
		output.WriteRune(unicode.ToLower(r))
		isFirst = false
	}
	return output.String()
}

// LowerCamelCase - alias since CamelCase and LowerCamelCase are synonyms
func LowerCamelCase(s string) string {
	return CamelCase(s)
}

// ConvertSnakeCaseToCamelCase - wrapper for CovertToCamelCase for specific separator (_)
func ConvertSnakeCaseToCamelCase(s string) string {
	return ConvertToCamelCase(s, '_')
}

// ConvertKebabCaseToCamelCase - wrapper for CovertToCamelCase for specific separator (-)
func ConvertKebabCaseToCamelCase(s string) string {
	return ConvertToCamelCase(s, '-')
}

// ConvertSentenceCaseToCamelCase - wrapper for CovertToCamelCase for specific separator (space)
func ConvertSentenceCaseToCamelCase(s string) string {
	return ConvertToCamelCase(s, ' ')
}

// ConvertToCamelCase - Simple CamelCase converter that splits on a single separator
func ConvertToCamelCase(s string, sep rune) string {
	// this is a string -> ThisIsAString
	var output strings.Builder
	nextUpper := false
	for _, r := range s {
		if r == sep {
			nextUpper = true
			continue
		}
		if nextUpper {
			output.WriteRune(unicode.ToUpper(r))
			nextUpper = false
			continue
		}
		output.WriteRune(unicode.ToLower(r))
	}
	return output.String()
}
