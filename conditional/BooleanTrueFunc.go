package conditional

// Deprecated: Use IfTrueFuncBool instead.
func BooleanTrueFunc(
	isTrue bool,
	trueValueFunc func() bool,
) bool {
	return IfTrueFunc[bool](isTrue, trueValueFunc)
}
