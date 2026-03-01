package conditional

// Deprecated: Use IfSliceBool instead.
func BooleansPtr(
	isTrue bool,
	trueValue, falseValue []bool,
) []bool {
	return IfSlice[bool](isTrue, trueValue, falseValue)
}
