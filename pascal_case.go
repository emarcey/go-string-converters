package gostringconverters

import (
	"strings"
	"unicode"
)

// PascalCase - Punctuation Aggressive PascalCase converter
// Will treat any space or punctuation as a separator
func PascalCase(s string) string {
	// this is a string -> ThisIsAString
	var output strings.Builder
	nextUpper := true
	for _, r := range s {
		if IsSeparator(r) {
			nextUpper = true
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
	}
	return output.String()
}

// UpperCamelCase - alias since PascalCase and UpperCamelCase are synonyms
func UpperCamelCase(s string) string {
	return PascalCase(s)
}

// ConvertSnakeCaseToPascalCase - wrapper for ConvertToPascalCase for specific separator (_)
func ConvertSnakeCaseToPascalCase(s string) string {
	return ConvertToPascalCase(s, '_')
}

// ConvertKebabCaseToPascalCase - wrapper for ConvertToPascalCase for specific separator (-)
func ConvertKebabCaseToPascalCase(s string) string {
	return ConvertToPascalCase(s, '-')
}

// ConvertSentenceCaseToPascalCase - wrapper for ConvertToPascalCase for specific separator (space)
func ConvertSentenceCaseToPascalCase(s string) string {
	return ConvertToPascalCase(s, ' ')
}

// ConvertToPascalCase - Simple PascalCase converter that splits on a single separator
func ConvertToPascalCase(s string, sep rune) string {
	// this is a string -> ThisIsAString
	var output strings.Builder
	nextUpper := true
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
