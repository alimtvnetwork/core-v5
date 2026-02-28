package conditional

// Deprecated: Use IfSlice[string] instead.
func Strings(
	isTrue bool,
	trueValue, falseValue []string,
) []string {
	return IfSlice[string](isTrue, trueValue, falseValue)
}
