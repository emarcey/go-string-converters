package gostringconverters

import (
	"testing"
)

func TestToLeetBasic(t *testing.T) {
	leetKeys := ""
	leetValues := ""

	for k, v := range basicToLeetMap {
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
		{
			given:    "leet hacker",
			expected: "l337 h4ck3r",
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

	for k, v := range basicFromLeetMap {
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
		{
			given:    "l337 h4ck3r",
			expected: "leet hacker",
		},
	}
	for i, testCase := range testCases {
		result := FromLeetBasic(testCase.given)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "expected", testCase.expected, "result", result)
		}
	}
}
