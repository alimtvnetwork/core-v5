package conditional

// Deprecated: Use IfFunc[bool] instead.
func BoolFunc(
	isTrue bool,
	trueValueFunc, falseValueFunc func() bool,
) bool {
	return IfFunc[bool](isTrue, trueValueFunc, falseValueFunc)
}
