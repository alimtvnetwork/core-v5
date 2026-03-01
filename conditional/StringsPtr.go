package conditional

// Deprecated: Use IfSliceString instead.
func StringsPtr(
	isTrue bool,
	trueValue, falseValue []string,
) []string {
	return IfSlice[string](isTrue, trueValue, falseValue)
}
