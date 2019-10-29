package gostringconverters

// SeparatedCaseOptions - options passed to ConvertToSeparatedCase
// Allows run separator & IsScreaming (uppercase) bool
type SeparatedCaseOptions struct {
	Separator   rune
	IsScreaming bool
}
