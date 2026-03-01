package conditional

// Deprecated: Use IfSliceAny instead.
func Interfaces(
	isTrue bool,
	trueValue, falseValue []any,
) []any {
	return IfSlice[any](isTrue, trueValue, falseValue)
}
