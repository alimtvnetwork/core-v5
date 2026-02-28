package conditional

// Deprecated: Use IfSlice[any] instead.
func Interfaces(
	isTrue bool,
	trueValue, falseValue []any,
) []any {
	return IfSlice[any](isTrue, trueValue, falseValue)
}
