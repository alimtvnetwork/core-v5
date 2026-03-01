package conditional

// Deprecated: Use IfSliceInt instead.
func IntegersPtr(
	isTrue bool,
	trueValue, falseValue []int,
) []int {
	return IfSlice[int](isTrue, trueValue, falseValue)
}
