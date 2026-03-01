package conditional

// Deprecated: Use IfTrueFuncString instead.
func StringTrueFunc(
	isTrue bool,
	trueValueFunc func() string,
) string {
	return IfTrueFunc[string](isTrue, trueValueFunc)
}
