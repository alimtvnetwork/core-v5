package conditional

// Deprecated: Use IfSlicePtr[string] instead.
func StringsPtr(
	isTrue bool,
	trueValue, falseValue *[]string,
) *[]string {
	return IfSlicePtr[string](isTrue, trueValue, falseValue)
}
