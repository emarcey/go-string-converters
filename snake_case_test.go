package gostringconverters

import (
	"testing"
)

func TestSnakeCase(t *testing.T) {
	var testCases = []struct {
		givenS   string
		expected string
	}{
		{
			givenS:   "",
			expected: "",
		},
		{
			givenS:   "",
			expected: "",
		},
		{
			givenS:   "abc",
			expected: "abc",
		},
		{
			givenS:   "ABC",
			expected: "a_b_c",
		},
		{
			givenS:   "a b c",
			expected: "a_b_c",
		},
		{
			givenS:   "a-b-c",
			expected: "a_b_c",
		},
		{
			givenS:   "a_b_c",
			expected: "a_b_c",
		},
		{
			givenS:   "  - a-b.c. help. -",
			expected: "a_b_c_help",
		},
		{
			givenS:   "  - a-b.c. helpই. -",
			expected: "a_b_c_helpই",
		},
		{
			givenS:   "  - a-b.c. help-ই. -",
			expected: "a_b_c_help_ই",
		},
		{
			givenS:   "thisIsASentence",
			expected: "this_is_a_sentence",
		},
		{
			givenS:   "ThisIsASentence",
			expected: "this_is_a_sentence",
		},
	}

	for i, testCase := range testCases {
		result := SnakeCase(testCase.givenS)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.givenS, "expected", testCase.expected, "result", result)
		}
	}
}

func TestScreamingSnakeCase(t *testing.T) {
	var testCases = []struct {
		givenS   string
		expected string
	}{
		{
			givenS:   "",
			expected: "",
		},
		{
			givenS:   "",
			expected: "",
		},
		{
			givenS:   "abc",
			expected: "ABC",
		},
		{
			givenS:   "ABC",
			expected: "A_B_C",
		},
		{
			givenS:   "a b c",
			expected: "A_B_C",
		},
		{
			givenS:   "a-b-c",
			expected: "A_B_C",
		},
		{
			givenS:   "a_b_c",
			expected: "A_B_C",
		},
		{
			givenS:   "  - a-b.c. help. -",
			expected: "A_B_C_HELP",
		},
		{
			givenS:   "  - a-b.c. helpই. -",
			expected: "A_B_C_HELPই",
		},
		{
			givenS:   "  - a-b.c. help-ই. -",
			expected: "A_B_C_HELP_ই",
		},
		{
			givenS:   "thisIsASentence",
			expected: "THIS_IS_A_SENTENCE",
		},
		{
			givenS:   "ThisIsASentence",
			expected: "THIS_IS_A_SENTENCE",
		},
	}

	for i, testCase := range testCases {
		result := ScreamingSnakeCase(testCase.givenS)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.givenS, "expected", testCase.expected, "result", result)
		}
	}
}
