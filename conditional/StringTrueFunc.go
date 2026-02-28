package conditional

// Deprecated: Use IfTrueFunc[string] instead.
func StringTrueFunc(
	isTrue bool,
	trueValueFunc func() string,
) string {
	return IfTrueFunc[string](isTrue, trueValueFunc)
}
