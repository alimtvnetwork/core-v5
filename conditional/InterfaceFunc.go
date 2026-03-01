package conditional

// Deprecated: Use IfFuncAny instead.
func InterfaceFunc(
	isTrue bool,
	trueValueFunc, falseValueFunc func() any,
) any {
	return IfFunc[any](isTrue, trueValueFunc, falseValueFunc)
}
