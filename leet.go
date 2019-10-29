package gostringconverters

var (
	basicToLeetMap = map[rune]rune{
		'a': '4',
		'e': '3',
		'g': '6',
		'i': '1',
		'o': '0',
		's': '5',
		't': '7',
	}

	basicFromLeetMap = map[rune]rune{
		'a': '4',
		'e': '3',
		'g': '6',
		'i': '1',
		'o': '0',
		's': '5',
		't': '7',
	}
)

// ToLeetBasic - converts a string to Leet using basic Leet values
func ToLeetBasic(s string) string {
	return RuneMapConversion(s, basicToLeetMap)
}

// FromLeetBasic - converts a string from Leet using basic Leet values
func FromLeetBasic(s string) string {
	return RuneMapConversion(s, basicFromLeetMap)
}
