package conditional

// Deprecated: Use IfSlicePtrFunc[int] instead.
func IntegersPtrFunc(
	isTrue bool,
	trueValueFunc, falseValueFunc func() *[]int,
) *[]int {
	return IfSlicePtrFunc[int](isTrue, trueValueFunc, falseValueFunc)
}
