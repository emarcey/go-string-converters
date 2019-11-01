package gostringconverters

import (
	"reflect"
	"testing"
)

func TestConvertToSeparatedCase(t *testing.T) {
	var testCases = []struct {
		givenS    string
		givenOpts SeparatedCaseOptions
		expected  string
	}{
		{
			givenS:    "",
			givenOpts: SeparatedCaseOptions{},
			expected:  "",
		},
		{
			givenS: "",
			givenOpts: SeparatedCaseOptions{
				Separator:   '-',
				IsScreaming: true,
			},
			expected: "",
		},
		{
			givenS:    "abc",
			givenOpts: SeparatedCaseOptions{},
			expected:  "abc",
		},
		{
			givenS: "abc",
			givenOpts: SeparatedCaseOptions{
				Separator:   '-',
				IsScreaming: true,
			},
			expected: "ABC",
		},
		{
			givenS: "a-b-c",
			givenOpts: SeparatedCaseOptions{
				Separator: '-',
			},
			expected: "a-b-c",
		},
		{
			givenS: "a-b-c",
			givenOpts: SeparatedCaseOptions{
				Separator:   '-',
				IsScreaming: true,
			},
			expected: "A-B-C",
		},
		{
			givenS: "a-b-c",
			givenOpts: SeparatedCaseOptions{
				Separator:   '_',
				IsScreaming: true,
			},
			expected: "A_B_C",
		},
		{
			givenS: "  - a-b.c. help. -",
			givenOpts: SeparatedCaseOptions{
				Separator:   '_',
				IsScreaming: true,
			},
			expected: "A_B_C_HELP",
		},
		{
			givenS: "  - a-b.c. helpই. -",
			givenOpts: SeparatedCaseOptions{
				Separator:   '_',
				IsScreaming: true,
			},
			expected: "A_B_C_HELPই",
		},
		{
			givenS: "  - a-b.c. help-ই. -",
			givenOpts: SeparatedCaseOptions{
				Separator:   '_',
				IsScreaming: true,
			},
			expected: "A_B_C_HELP_ই",
		},
		{
			givenS: "thisIsASentence",
			givenOpts: SeparatedCaseOptions{
				Separator:   '_',
				IsScreaming: false,
			},
			expected: "this_is_a_sentence",
		},
		{
			givenS: "ThisIsASentence",
			givenOpts: SeparatedCaseOptions{
				Separator:   '_',
				IsScreaming: false,
			},
			expected: "this_is_a_sentence",
		},
	}

	for i, testCase := range testCases {
		result := ConvertToSeparatedCase(testCase.givenS, testCase.givenOpts)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.givenS, "and", testCase.givenOpts, "expected", testCase.expected, "result", result)
		}
	}
}

func TestIsSeparator(t *testing.T) {
	var testCases = []struct {
		given    rune
		expected bool
	}{
		{
			given:    'a',
			expected: false,
		},
		{
			given:    'A',
			expected: false,
		},
		{
			given:    '1',
			expected: false,
		},
		{
			given:    'ই',
			expected: false,
		},
		{
			given:    '-',
			expected: true,
		},
		{
			given:    '_',
			expected: true,
		},
		{
			given:    ' ',
			expected: true,
		},
		{
			given:    '?',
			expected: true,
		},
	}

	for i, testCase := range testCases {
		result := IsSeparator(testCase.given)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "expected", testCase.expected, "result", result)
		}
	}
}

func TestRuneMapToStringConversion(t *testing.T) {
	var testCases = []struct {
		givenS             string
		givenM             map[rune]string
		givenCaseSensitive bool
		expected           string
	}{
		{
			givenS:             "",
			givenM:             map[rune]string{},
			givenCaseSensitive: false,
			expected:           "",
		},
		{
			givenS:             "abc",
			givenM:             map[rune]string{},
			givenCaseSensitive: false,
			expected:           "abc",
		},
		{
			givenS: "abc",
			givenM: map[rune]string{
				'a': "this",
				'b': "is",
				'c': "cool",
			},
			givenCaseSensitive: false,
			expected:           "thisiscool",
		},
		{
			givenS: "abc",
			givenM: map[rune]string{
				'A': "this",
				'B': "is",
				'C': "cool",
			},
			givenCaseSensitive: false,
			expected:           "thisiscool",
		},
		{
			givenS: "abc",
			givenM: map[rune]string{
				'A': "this",
				'B': "is",
				'C': "cool",
			},
			givenCaseSensitive: true,
			expected:           "abc",
		},
	}

	for i, testCase := range testCases {
		result := RuneMapToStringConversion(testCase.givenS, testCase.givenM, testCase.givenCaseSensitive)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.givenS, "and", testCase.givenM, "and", testCase.givenCaseSensitive, "expected", testCase.expected, "result", result)
		}
	}
}
