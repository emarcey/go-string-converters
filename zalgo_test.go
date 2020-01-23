package gostringconverters

import (
	"math/rand"
	"testing"
)

func TestZalgoEncoder(t *testing.T) {
	var testCases = []struct {
		givenString   string
		givenNumChars int
		expected      string
	}{
		{
			givenString:   "",
			givenNumChars: 1,
			expected:      "",
		},
		{
			givenString:   "abcdefg",
			givenNumChars: 0,
			expected:      "abcdefg",
		},
		{
			givenString:   "abcdefg",
			givenNumChars: 1,
			expected:      "aͦb͉c̹d̶ėf̯g̻",
		},
		{
			givenString:   "abcdefg",
			givenNumChars: 5,
			expected:      "a̵ͯ̐͗ͦb͓̺͂ͨͩc̷̨̜̔̿d̻̽̊͠͞e̻̠͗̇̀f̮̩͛̽͢g̵̞͓͋ͅ",
		},
		{
			givenString:   "abcdefg",
			givenNumChars: 10,
			expected:      "a̸̡̰̓́͑ͧ̽͘͡b̵̨̨̲̰̝̽̍̎ͧc͎̟̞̽ͭͣͦ̋̆͞d̴̢͖̙̙̮̑͊ͥ͊e̢̨̗͙̱ͥ̎̊͘͘f̴̲̬̳͕̻ͥ̆͋ͪg̷̫̗̰̺͙͍̒̔̇",
		},
	}

	rand.Seed(1)
	for i, testCase := range testCases {
		result := ZalgoEncoder(testCase.givenString, testCase.givenNumChars)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.givenString, "and", testCase.givenNumChars, "expected", testCase.expected, "result", result)
		}
	}
}

func TestZalgoDecoder(t *testing.T) {
	var testCases = []struct {
		given    string
		expected string
	}{
		{
			given:    "",
			expected: "",
		},
		{
			given:    "abcdefg",
			expected: "abcdefg",
		},
		{
			given:    "aͦb͉c̹d̶ėf̯g̻",
			expected: "abcdefg",
		},
		{
			given:    "a̵ͯ̐͗ͦb͓̺͂ͨͩc̷̨̜̔̿d̻̽̊͠͞e̻̠͗̇̀f̮̩͛̽͢g̵̞͓͋ͅ",
			expected: "abcdefg",
		},
		{
			given:    "a̸̡̰̓́͑ͧ̽͘͡b̵̨̨̲̰̝̽̍̎ͧc͎̟̞̽ͭͣͦ̋̆͞d̴̢͖̙̙̮̑͊ͥ͊e̢̨̗͙̱ͥ̎̊͘͘f̴̲̬̳͕̻ͥ̆͋ͪg̷̫̗̰̺͙͍̒̔̇",
			expected: "abcdefg",
		},
	}

	rand.Seed(1)
	for i, testCase := range testCases {
		result := ZalgoDecoder(testCase.given)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given, "expected", testCase.expected, "result", result)
		}
	}
}
