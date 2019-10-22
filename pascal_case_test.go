package gostringconverters

import (
	"testing"
)

func TestPascalCase(t *testing.T) {
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
			expected: "Something",
		},
		{
			given:    "some-tHing",
			expected: "SomeThing",
		},
		{
			given:    "some_thing",
			expected: "SomeThing",
		},
		{
			given:    "some thing else",
			expected: "SomeThingElse",
		},
		{
			given:    "What is going on here? I don't ",
			expected: "WhatIsGoingOnHereIDonT",
		},
		{
			given:    "   - What is going on here? I don't ",
			expected: "WhatIsGoingOnHereIDonT",
		},
		{
			given:    "th1s",
			expected: "Th1s",
		},
		{
			given:    "thÃt-is-ãwesΘme-ص-ই",
			expected: "ThãtIsÃwesθmeصই",
		},
	}

	for i, testCase := range testCases {
		result := PascalCase(testCase.given)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "expected", testCase.expected, "result", result)
		}
	}
}

func TestUpperCamelCase(t *testing.T) {
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
			expected: "Something",
		},
		{
			given:    "some-tHing",
			expected: "SomeThing",
		},
		{
			given:    "some_thing",
			expected: "SomeThing",
		},
		{
			given:    "some thing else",
			expected: "SomeThingElse",
		},
		{
			given:    "What is going on here? I don't ",
			expected: "WhatIsGoingOnHereIDonT",
		},
		{
			given:    "   - What is going on here? I don't ",
			expected: "WhatIsGoingOnHereIDonT",
		},
		{
			given:    "th1s",
			expected: "Th1s",
		},
		{
			given:    "thÃt-is-ãwesΘme-ص-ই",
			expected: "ThãtIsÃwesθmeصই",
		},
	}

	for i, testCase := range testCases {
		result := UpperCamelCase(testCase.given)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "expected", testCase.expected, "result", result)
		}
	}
}

func TestConvertToPascalCase(t *testing.T) {
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
			expected: "Something",
		},
		{
			given:    "some-tHing",
			givenSep: '-',
			expected: "SomeThing",
		},
		{
			given:    "some-thing",
			givenSep: '_',
			expected: "Some-thing",
		},
	}

	for i, testCase := range testCases {
		result := ConvertToPascalCase(testCase.given, testCase.givenSep)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "and", testCase.givenSep, "expected", testCase.expected, "result", result)
		}
	}
}

func TestConvertSnakeCaseToPascalCase(t *testing.T) {
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
			expected: "Something",
		},
		{
			given:    "some-tHing",
			expected: "Some-thing",
		},
		{
			given:    "some tHing",
			expected: "Some thing",
		},
		{
			given:    "some_tHing",
			expected: "SomeThing",
		},
	}

	for i, testCase := range testCases {
		result := ConvertSnakeCaseToPascalCase(testCase.given)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "expected", testCase.expected, "result", result)
		}
	}
}

func TestConvertKebabCaseToPascalCase(t *testing.T) {
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
			expected: "Something",
		},
		{
			given:    "some-tHing",
			expected: "SomeThing",
		},
		{
			given:    "some tHing",
			expected: "Some thing",
		},
		{
			given:    "some_tHing",
			expected: "Some_thing",
		},
	}

	for i, testCase := range testCases {
		result := ConvertKebabCaseToPascalCase(testCase.given)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "expected", testCase.expected, "result", result)
		}
	}
}

func TestConvertSentenceCaseToPascalCase(t *testing.T) {
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
			expected: "Something",
		},
		{
			given:    "some-tHing",
			expected: "Some-thing",
		},
		{
			given:    "some tHing",
			expected: "SomeThing",
		},
		{
			given:    "some_tHing",
			expected: "Some_thing",
		},
	}

	for i, testCase := range testCases {
		result := ConvertSentenceCaseToPascalCase(testCase.given)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "expected", testCase.expected, "result", result)
		}
	}
}
