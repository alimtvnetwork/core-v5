package conditional

// Deprecated: Use IfSlicePtr[int] instead.
func IntegersPtr(
	isTrue bool,
	trueValue, falseValue *[]int,
) *[]int {
	return IfSlicePtr[int](isTrue, trueValue, falseValue)
}
