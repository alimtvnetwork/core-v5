package conditional

// Deprecated: Use IfSlice[bool] instead.
func Booleans(
	isTrue bool,
	trueValue, falseValue []bool,
) []bool {
	return IfSlice[bool](isTrue, trueValue, falseValue)
}
