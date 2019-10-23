package gostringconverters

func KebabCase(s string) string {
	// thisIsAString -> this-is-a-string
	opts := SeparatedCaseOptions{
		Separator:   '-',
		IsScreaming: false,
	}
	return ConvertToSeparatedCase(s, opts)
}

func ScreamingKebabCase(s string) string {
	// thisIsAString -> THIS-IS-A-STRING
	opts := SeparatedCaseOptions{
		Separator:   '-',
		IsScreaming: true,
	}
	return ConvertToSeparatedCase(s, opts)
}
