package conditional

// Deprecated: Use IfSliceBool instead.
func Booleans(
	isTrue bool,
	trueValue, falseValue []bool,
) []bool {
	return IfSlice[bool](isTrue, trueValue, falseValue)
}
