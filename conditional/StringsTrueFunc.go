package conditional

// Deprecated: Use IfTrueFuncStrings instead.
func StringsTrueFunc(
	isTrue bool,
	trueValueFunc func() []string,
) []string {
	return IfTrueFunc[[]string](isTrue, trueValueFunc)
}
