package gostringconverters

import (
	"testing"
)

func TestToLeetBasic(t *testing.T) {
	leetKeys := ""
	leetValues := ""

	for k, v := range BASIC_TO_LEET_MAP {
		leetKeys += string(k)
		leetValues += string(v)
	}
	var testCases = []struct {
		given    string
		expected string
	}{
		{
			given:    "",
			expected: "",
		},
		{
			given:    leetKeys,
			expected: leetValues,
		},
	}
	for i, testCase := range testCases {
		result := ToLeetBasic(testCase.given)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "expected", testCase.expected, "result", result)
		}
	}
}

func TestFromLeetBasic(t *testing.T) {
	leetKeys := ""
	leetValues := ""

	for k, v := range BASIC_FROM_LEET_MAP {
		leetKeys += string(k)
		leetValues += string(v)
	}
	var testCases = []struct {
		given    string
		expected string
	}{
		{
			given:    "",
			expected: "",
		},
		{
			given:    leetKeys,
			expected: leetValues,
		},
	}
	for i, testCase := range testCases {
		result := FromLeetBasic(testCase.given)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "expected", testCase.expected, "result", result)
		}
	}
}
