package gostringconverters

import (
	"strings"
)

var (
	basicToLeetMap = map[rune]rune{
		'a': '4',
		'e': '3',
		'g': '6',
		'i': '1',
		'o': '0',
		's': '5',
		't': '7',
	}

	basicFromLeetMap = map[rune]rune{
		'4': 'a',
		'3': 'e',
		'6': 'g',
		'1': 'i',
		'0': 'o',
		'5': 's',
		'7': 't',
	}
)

// ToLeetBasic - converts a string to Leet using basic Leet values
func ToLeetBasic(s string) string {
	return strings.Map(RuneMapToMapFunc(basicToLeetMap), s)
}

// FromLeetBasic - converts a string from Leet using basic Leet values
func FromLeetBasic(s string) string {
	return strings.Map(RuneMapToMapFunc(basicFromLeetMap), s)
}
