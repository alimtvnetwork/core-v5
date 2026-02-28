package conditional

// Deprecated: Use IfSlicePtr[any] instead.
func InterfacesPtr(
	isTrue bool,
	trueValue, falseValue *[]any,
) *[]any {
	return IfSlicePtr[any](isTrue, trueValue, falseValue)
}
