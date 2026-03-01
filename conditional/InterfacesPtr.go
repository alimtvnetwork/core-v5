package conditional

// Deprecated: Use IfSlice[any] instead.
func InterfacesPtr(
	isTrue bool,
	trueValue, falseValue []any,
) []any {
	return IfSlice[any](isTrue, trueValue, falseValue)
}
