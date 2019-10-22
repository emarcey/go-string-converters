package gostringconverters

import (
	"testing"
)

func TestCamelCase(t *testing.T) {
	var testCases = []struct {
		given    string
		expected string
	}{
		{
			given:    "",
			expected: "",
		},
		{
			given:    "something",
			expected: "something",
		},
		{
			given:    "some-tHing",
			expected: "someThing",
		},
		{
			given:    "some_thing",
			expected: "someThing",
		},
		{
			given:    "some thing else",
			expected: "someThingElse",
		},
		{
			given:    "What is going on here? I don't ",
			expected: "whatIsGoingOnHereIDonT",
		},
		{
			given:    "   - What is going on here? I don't ",
			expected: "whatIsGoingOnHereIDonT",
		},
		{
			given:    "th1s",
			expected: "th1s",
		},
		{
			given:    "thÃt-is-ãwesΘme-ص-ই",
			expected: "thãtIsÃwesθmeصই",
		},
	}

	for i, testCase := range testCases {
		result := CamelCase(testCase.given)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "expected", testCase.expected, "result", result)
		}
	}
}

func TestConvertToCamelCase(t *testing.T) {
	var testCases = []struct {
		given    string
		givenSep rune
		expected string
	}{
		{
			given:    "",
			givenSep: '-',
			expected: "",
		},
		{
			given:    "something",
			givenSep: '-',
			expected: "something",
		},
		{
			given:    "some-tHing",
			givenSep: '-',
			expected: "someThing",
		},
		{
			given:    "some-thing",
			givenSep: '_',
			expected: "some-thing",
		},
	}

	for i, testCase := range testCases {
		result := ConvertToCamelCase(testCase.given, testCase.givenSep)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "and", testCase.givenSep, "expected", testCase.expected, "result", result)
		}
	}
}

func TestConvertSnakeCaseToCamelCase(t *testing.T) {
	var testCases = []struct {
		given    string
		expected string
	}{
		{
			given:    "",
			expected: "",
		},
		{
			given:    "something",
			expected: "something",
		},
		{
			given:    "some-tHing",
			expected: "some-thing",
		},
		{
			given:    "some tHing",
			expected: "some thing",
		},
		{
			given:    "some_tHing",
			expected: "someThing",
		},
	}

	for i, testCase := range testCases {
		result := ConvertSnakeCaseToCamelCase(testCase.given)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "expected", testCase.expected, "result", result)
		}
	}
}

func TestConvertKebabCaseToCamelCase(t *testing.T) {
	var testCases = []struct {
		given    string
		expected string
	}{
		{
			given:    "",
			expected: "",
		},
		{
			given:    "something",
			expected: "something",
		},
		{
			given:    "some-tHing",
			expected: "someThing",
		},
		{
			given:    "some tHing",
			expected: "some thing",
		},
		{
			given:    "some_tHing",
			expected: "some_thing",
		},
	}

	for i, testCase := range testCases {
		result := ConvertKebabCaseToCamelCase(testCase.given)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "expected", testCase.expected, "result", result)
		}
	}
}

func TestConvertSentenceCaseToCamelCase(t *testing.T) {
	var testCases = []struct {
		given    string
		expected string
	}{
		{
			given:    "",
			expected: "",
		},
		{
			given:    "something",
			expected: "something",
		},
		{
			given:    "some-tHing",
			expected: "some-thing",
		},
		{
			given:    "some tHing",
			expected: "someThing",
		},
		{
			given:    "some_tHing",
			expected: "some_thing",
		},
	}

	for i, testCase := range testCases {
		result := ConvertSentenceCaseToCamelCase(testCase.given)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "expected", testCase.expected, "result", result)
		}
	}
}
