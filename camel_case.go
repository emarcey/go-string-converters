package gostringconverters

import (
	"strings"
	"unicode"
)

// Punctuation Aggressive CamelCase converter
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

// alias since CamelCase and LowerCamelCase are synonyms
func LowerCamelCase(s string) string {
	return CamelCase(s)
}

func ConvertSnakeCaseToCamelCase(s string) string {
	return ConvertToCamelCase(s, '_')
}

func ConvertKebabCaseToCamelCase(s string) string {
	return ConvertToCamelCase(s, '-')
}

func ConvertSentenceCaseToCamelCase(s string) string {
	return ConvertToCamelCase(s, ' ')
}

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
