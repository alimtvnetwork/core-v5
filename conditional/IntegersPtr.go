package conditional

// Deprecated: Use IfSlice[int] instead.
func IntegersPtr(
	isTrue bool,
	trueValue, falseValue []int,
) []int {
	return IfSlice[int](isTrue, trueValue, falseValue)
}
