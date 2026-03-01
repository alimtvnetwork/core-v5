package conditional

// Deprecated: Use IfSliceInt instead.
func Integers(
	isTrue bool,
	trueValue, falseValue []int,
) []int {
	return IfSlice[int](isTrue, trueValue, falseValue)
}
