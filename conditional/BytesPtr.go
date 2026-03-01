package conditional

// Deprecated: Use IfSliceByte instead.
func BytesPtr(
	isTrue bool,
	trueValue, falseValue []byte,
) []byte {
	return IfSlice[byte](isTrue, trueValue, falseValue)
}
