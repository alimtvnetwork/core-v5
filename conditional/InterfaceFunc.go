package conditional

// Deprecated: Use IfFunc[any] instead.
func InterfaceFunc(
	isTrue bool,
	trueValueFunc, falseValueFunc func() any,
) any {
	return IfFunc[any](isTrue, trueValueFunc, falseValueFunc)
}
