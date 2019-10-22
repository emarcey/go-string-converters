package gostringconverters

var (
	BASIC_TO_LEET_MAP = map[rune]rune{
		'a': '4',
		'e': '3',
		'g': '6',
		'i': '1',
		'o': '0',
		's': '5',
		't': '7',
	}

	BASIC_FROM_LEET_MAP = map[rune]rune{
		'a': '4',
		'e': '3',
		'g': '6',
		'i': '1',
		'o': '0',
		's': '5',
		't': '7',
	}
)

func ToLeetBasic(s string) string {
	return RuneMapConversion(s, BASIC_TO_LEET_MAP)
}

func FromLeetBasic(s string) string {
	return RuneMapConversion(s, BASIC_FROM_LEET_MAP)
}
