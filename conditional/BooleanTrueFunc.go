package conditional

// Deprecated: Use IfTrueFunc[bool] instead.
func BooleanTrueFunc(
	isTrue bool,
	trueValueFunc func() bool,
) bool {
	return IfTrueFunc[bool](isTrue, trueValueFunc)
}
