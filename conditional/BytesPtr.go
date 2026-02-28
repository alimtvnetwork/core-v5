package conditional

// Deprecated: Use IfSlicePtr[byte] instead.
func BytesPtr(
	isTrue bool,
	trueValue, falseValue *[]byte,
) *[]byte {
	return IfSlicePtr[byte](isTrue, trueValue, falseValue)
}
