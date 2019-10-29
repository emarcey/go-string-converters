package gostringconverters

// KebabCase - Wraps ConvertToSeparatedCase to make lowercase, hyphen delimited text
func KebabCase(s string) string {
	// thisIsAString -> this-is-a-string
	opts := SeparatedCaseOptions{
		Separator:   '-',
		IsScreaming: false,
	}
	return ConvertToSeparatedCase(s, opts)
}

// ScreamingKebabCase - Wraps ConvertToSeparatedCase to make uppercase, hyphen delimited text
func ScreamingKebabCase(s string) string {
	// thisIsAString -> THIS-IS-A-STRING
	opts := SeparatedCaseOptions{
		Separator:   '-',
		IsScreaming: true,
	}
	return ConvertToSeparatedCase(s, opts)
}
