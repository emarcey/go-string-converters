package gostringconverters

import (
	"testing"
)

func TestRuneMapConversion(t *testing.T) {
	var testCases = []struct {
		givenS   string
		givenM   map[rune]rune
		expected string
	}{
		{
			givenS:   "",
			givenM:   map[rune]rune{},
			expected: "",
		},
		{
			givenS:   "abc",
			givenM:   map[rune]rune{},
			expected: "abc",
		},
		{
			givenS: "",
			givenM: map[rune]rune{
				'a': '3',
				'b': '2',
				'c': '1',
			},
			expected: "",
		},
		{
			givenS: "abc",
			givenM: map[rune]rune{
				'a': '3',
				'b': '2',
				'c': '1',
			},
			expected: "321",
		},
	}

	for i, testCase := range testCases {
		result := RuneMapConversion(testCase.givenS, testCase.givenM)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.givenS, "and", testCase.givenM, "expected", testCase.expected, "result", result)
		}
	}
}
