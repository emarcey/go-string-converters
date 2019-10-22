package gostringconverters

import (
	"strings"
	"unicode"
)

// Punctuation Aggressive PascalCase converter
// Will treat any space or punctuation as a separator
func PascalCase(s string) string {
	// this is a string -> ThisIsAString
	var output strings.Builder
	nextUpper := true
	for _, r := range s {
		if unicode.IsPunct(r) || unicode.IsSpace(r) {
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

// alias since PascalCase and UpperCamelCase are synonyms
func UpperCamelCase(s string) string {
	return PascalCase(s)
}

func ConvertSnakeCaseToPascalCase(s string) string {
	return ConvertToPascalCase(s, '_')
}

func ConvertKebabCaseToPascalCase(s string) string {
	return ConvertToPascalCase(s, '-')
}

func ConvertSentenceCaseToPascalCase(s string) string {
	return ConvertToPascalCase(s, ' ')
}

// tightly restricted pascal case converted that allows an input rune for a delimiter
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
