package conditional

// Deprecated: Use IfSlice[int] with func wrappers instead.
func IntegersPtrFunc(
	isTrue bool,
	trueValueFunc, falseValueFunc func() []int,
) []int {
	if isTrue {
		return trueValueFunc()
	}
	return falseValueFunc()
}
