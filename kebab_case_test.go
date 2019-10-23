package gostringconverters

import (
	"testing"
)

func TestKebabCase(t *testing.T) {
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
			givenS:   "ABC",
			expected: "a-b-c",
		},
		{
			givenS:   "a b c",
			expected: "a-b-c",
		},
		{
			givenS:   "abc",
			expected: "abc",
		},
		{
			givenS:   "a-b-c",
			expected: "a-b-c",
		},
		{
			givenS:   "a_b_c",
			expected: "a-b-c",
		},
		{
			givenS:   "  - a-b.c. help. -",
			expected: "a-b-c-help",
		},
		{
			givenS:   "  - a-b.c. helpই. -",
			expected: "a-b-c-helpই",
		},
		{
			givenS:   "  - a-b.c. help-ই. -",
			expected: "a-b-c-help-ই",
		},
		{
			givenS:   "thisIsASentence",
			expected: "this-is-a-sentence",
		},
		{
			givenS:   "ThisIsASentence",
			expected: "this-is-a-sentence",
		},
	}

	for i, testCase := range testCases {
		result := KebabCase(testCase.givenS)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.givenS, "expected", testCase.expected, "result", result)
		}
	}
}

func TestScreamingKebabCase(t *testing.T) {
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
			givenS:   "ABC",
			expected: "A-B-C",
		},
		{
			givenS:   "a b c",
			expected: "A-B-C",
		},
		{
			givenS:   "abc",
			expected: "ABC",
		},
		{
			givenS:   "a-b-c",
			expected: "A-B-C",
		},
		{
			givenS:   "a_b_c",
			expected: "A-B-C",
		},
		{
			givenS:   "  - a-b.c. help. -",
			expected: "A-B-C-HELP",
		},
		{
			givenS:   "  - a-b.c. helpই. -",
			expected: "A-B-C-HELPই",
		},
		{
			givenS:   "  - a-b.c. help-ই. -",
			expected: "A-B-C-HELP-ই",
		},
		{
			givenS:   "thisIsASentence",
			expected: "THIS-IS-A-SENTENCE",
		},
		{
			givenS:   "ThisIsASentence",
			expected: "THIS-IS-A-SENTENCE",
		},
	}

	for i, testCase := range testCases {
		result := ScreamingKebabCase(testCase.givenS)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.givenS, "expected", testCase.expected, "result", result)
		}
	}
}
