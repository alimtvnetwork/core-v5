package conditional

// Deprecated: Use IfFunc[string] instead.
func StringFunc(
	isTrue bool,
	trueValueFunc, falseValueFunc func() string,
) string {
	return IfFunc[string](isTrue, trueValueFunc, falseValueFunc)
}
