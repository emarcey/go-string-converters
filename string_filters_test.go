package gostringconverters

import (
	"testing"
)

func RuneIsA() RuneFilterFunc {
	return func(r rune) bool {
		return r == 'a'
	}
}

func RuneIsNotA() RuneFilterFunc {
	return func(r rune) bool {
		return r != 'a'
	}
}

func RuneFilterAlwaysTrue() RuneFilterFunc {
	return func(r rune) bool {
		return true
	}
}

func RuneFilterAlwaysFalse() RuneFilterFunc {
	return func(r rune) bool {
		return false
	}
}

func TestRuneIsFiltered(t *testing.T) {
	var testCases = []struct {
		givenRune    rune
		givenFilters []RuneFilterFunc
		expected     bool
	}{
		{
			givenRune:    'a',
			givenFilters: []RuneFilterFunc{},
			expected:     false,
		},
		{
			givenRune: 'a',
			givenFilters: []RuneFilterFunc{
				RuneIsNotA(),
			},
			expected: false,
		},
		{
			givenRune: 'a',
			givenFilters: []RuneFilterFunc{
				RuneIsA(),
			},
			expected: true,
		},
		{
			givenRune: 'a',
			givenFilters: []RuneFilterFunc{
				RuneIsNotA(),
				RuneFilterAlwaysTrue(),
				RuneIsA(),
			},
			expected: true,
		},
		{
			givenRune: 'a',
			givenFilters: []RuneFilterFunc{
				RuneIsA(),
				RuneIsNotA(),
				RuneFilterAlwaysTrue(),
			},
			expected: true,
		},
	}

	for i, testCase := range testCases {
		result := runeIsFiltered(testCase.givenRune, testCase.givenFilters)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.givenRune, "and", testCase.givenFilters, "expected", testCase.expected, "result", result)
		}
	}
}

func TestFilterStringChars(t *testing.T) {
	var testCases = []struct {
		givenString  string
		givenFilters []RuneFilterFunc
		expected     string
	}{
		{
			givenString:  "",
			givenFilters: []RuneFilterFunc{},
			expected:     "",
		},
		{
			givenString:  "abc",
			givenFilters: []RuneFilterFunc{},
			expected:     "abc",
		},
		{
			givenString: "abc",
			givenFilters: []RuneFilterFunc{
				RuneFilterAlwaysFalse(),
			},
			expected: "abc",
		},
		{
			givenString: "abc",
			givenFilters: []RuneFilterFunc{
				RuneFilterAlwaysTrue(),
			},
			expected: "",
		},
		{
			givenString: "abc",
			givenFilters: []RuneFilterFunc{
				RuneIsA(),
			},
			expected: "bc",
		},
		{
			givenString: "abc",
			givenFilters: []RuneFilterFunc{
				RuneIsNotA(),
			},
			expected: "a",
		},
		{
			givenString: "abc",
			givenFilters: []RuneFilterFunc{
				RuneIsNotA(),
				RuneIsA(),
			},
			expected: "",
		},
	}

	for i, testCase := range testCases {
		result := FilterStringChars(testCase.givenString, testCase.givenFilters)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.givenString, "and", testCase.givenFilters, "expected", testCase.expected, "result", result)
		}
	}
}

func TestRemoveControlCharsFromString(t *testing.T) {
	var testCases = []struct {
		given    string
		expected string
	}{
		{
			given:    "",
			expected: "",
		},
		{
			given:    "adfadgafdag32412@!@$",
			expected: "adfadgafdag32412@!@$",
		},
		{
			given:    "adfadgafda" + string(rune(1)) + "g32412@!@$",
			expected: "adfadgafdag32412@!@$",
		},
		{
			given:    string(rune(1)) + "adfadgafda" + string(rune(1)) + "g32412@!@$" + string(rune(1)),
			expected: "adfadgafdag32412@!@$",
		},
	}

	for i, testCase := range testCases {
		result := RemoveControlCharsFromString(testCase.given)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "expected", testCase.expected, "result", result)
		}
	}
}
