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
func RuneMapToMapFunc(runeMap map[rune]rune, caseSensitive bool) func(rune) rune {
	return func(r rune) rune {
		tempR := r
		if !caseSensitive {
			tempR = unicode.ToLower(r)
		}
		_, ok := runeMap[tempR]
		if !(ok || caseSensitive) {
			tempR = unicode.ToUpper(r)
			_, ok = runeMap[tempR]
		}
		if !ok {
			return r
		}
		return runeMap[tempR]
	}
}

// RuneMapConversion - given a string and a map, converts every rune in string to corresponding map value
func RuneMapToStringConversion(s string, m map[rune]string, caseSensitive bool) string {
	var output strings.Builder
	for _, r := range s {
		tempR := r
		if !caseSensitive {
			tempR = unicode.ToLower(r)
		}
		l, ok := m[tempR]
		if !(ok || caseSensitive) {
			tempR = unicode.ToUpper(r)
			l, ok = m[tempR]
		}
		if !ok {
			output.WriteRune(r)
			continue
		}
		output.WriteString(l)
	}
	return output.String()
}

func DecomposeMapStringToRune(m map[string]rune, caseSensitive bool) map[rune]DecomposedRune {
	decomposedRuneMap := make(map[rune]DecomposedRune)
	for key, value := range m {
		decomposedRuneMap = addRunesToDecomposedRuneMap(decomposedRuneMap, key, value, caseSensitive)
	}
	return decomposedRuneMap
}

func addRunesToDecomposedRuneMap(decomposedRuneMap map[rune]DecomposedRune, s string, value rune, caseSensitive bool) map[rune]DecomposedRune {
	if len(s) == 0 {
		return decomposedRuneMap
	}

	r, ok := decomposedRuneMapHasRune(decomposedRuneMap, rune(s[0]), caseSensitive)
	decomposedRune := DecomposedRune{}
	if ok {
		decomposedRune = decomposedRuneMap[r]
	}
	if decomposedRune.NextRunes == nil {
		decomposedRune.NextRunes = make(map[rune]DecomposedRune)
	}

	if len(s) == 1 {
		decomposedRune.Value = value
		decomposedRuneMap[r] = decomposedRune
		return decomposedRuneMap
	}
	decomposedRune.NextRunes = addRunesToDecomposedRuneMap(decomposedRune.NextRunes, s[1:], value, caseSensitive)
	decomposedRuneMap[r] = decomposedRune
	return decomposedRuneMap
}

func decomposedRuneMapHasRune(m map[rune]DecomposedRune, r rune, caseSensitive bool) (rune, bool) {
	tempR := r
	if !caseSensitive {
		tempR = unicode.ToLower(r)
	}
	_, ok := m[tempR]
	if ok {
		return tempR, true
	}
	if !(ok || caseSensitive) {
		tempR = unicode.ToUpper(r)
		_, ok = m[tempR]
	}
	return tempR, ok
}
