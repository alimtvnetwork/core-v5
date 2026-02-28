package conditional

// Deprecated: Use IfTrueFunc[[]string] instead.
func StringsTrueFunc(
	isTrue bool,
	trueValueFunc func() []string,
) []string {
	return IfTrueFunc[[]string](isTrue, trueValueFunc)
}
