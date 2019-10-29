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

	advancedToLeetMap = map[rune]string{
		'a': "4",
		'b': "|3",
		'c': "(",
		'd': "|)",
		'e': "3",
		'f': "|=",
		'g': "9",
		'h': "|-|",
		'i': "!",
		'j': "_|",
		'k': "|<",
		'l': "|_",
		'm': "/\\/\\",
		'n': "|\\|",
		'o': "0",
		'p': "|D",
		'q': "q",
		'r': "|2",
		's': "5",
		't': "7",
		'u': "(_)",
		'v': "\\/",
		'w': "\\/\\/",
		'x': "><",
		'y': "`/",
		'z': "2",
	}

	advancedFromLeetMap = map[string]rune{
		"4":      'a',
		"|3":     'b',
		"(":      'c',
		"|)":     'd',
		"3":      'e',
		"|=":     'f',
		"9":      'g',
		"|-|":    'h',
		"!":      'i',
		"_|":     'j',
		"|<":     'k',
		"|_":     'l',
		"/\\/\\": 'm',
		"|\\|":   'n',
		"0":      'o',
		"|D":     'p',
		"q":      'q',
		"|2":     'r',
		"5":      's',
		"7":      't',
		"(_)":    'u',
		"\\/":    'v',
		"\\/\\/": 'w',
		"><":     'x',
		"`/":     'y',
		"2":      'z',
	}
)

// ToLeetBasic - converts a string to Leet using basic Leet values
func ToLeetBasic(s string) string {
	return strings.Map(RuneMapToMapFunc(basicToLeetMap, false), s)
}

// FromLeetBasic - converts a string from Leet using basic Leet values
func FromLeetBasic(s string) string {
	return strings.Map(RuneMapToMapFunc(basicFromLeetMap, false), s)
}

// ToLeetAdvanced - converts a string to Leet using advanced Leet values
func ToLeetAdvanced(s string) string {
	return RuneMapToStringConversion(s, advancedToLeetMap, false)
}
