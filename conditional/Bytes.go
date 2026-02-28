package conditional

// Deprecated: Use IfSlice[byte] instead.
func Bytes(
	isTrue bool,
	trueValue, falseValue []byte,
) []byte {
	return IfSlice[byte](isTrue, trueValue, falseValue)
}
