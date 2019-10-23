package gostringconverters

func SnakeCase(s string) string {
	// thisIsAString -> this_is_a_string
	opts := SeparatedCaseOptions{
		Separator:   '_',
		IsScreaming: false,
	}
	return ConvertToSeparatedCase(s, opts)
}

func ScreamingSnakeCase(s string) string {
	// thisIsAString -> THIS_IS_A_STRING
	opts := SeparatedCaseOptions{
		Separator:   '_',
		IsScreaming: true,
	}
	return ConvertToSeparatedCase(s, opts)
}
