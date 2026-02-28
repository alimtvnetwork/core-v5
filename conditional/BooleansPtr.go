package conditional

// Deprecated: Use IfSlicePtr[bool] instead.
func BooleansPtr(
	isTrue bool,
	trueValue, falseValue *[]bool,
) *[]bool {
	return IfSlicePtr[bool](isTrue, trueValue, falseValue)
}
