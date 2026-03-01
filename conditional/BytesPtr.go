package conditional

// Deprecated: Use IfSlice[byte] instead.
func BytesPtr(
	isTrue bool,
	trueValue, falseValue []byte,
) []byte {
	return IfSlice[byte](isTrue, trueValue, falseValue)
}
