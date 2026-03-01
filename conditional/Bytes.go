package conditional

// Deprecated: Use IfSliceByte instead.
func Bytes(
	isTrue bool,
	trueValue, falseValue []byte,
) []byte {
	return IfSlice[byte](isTrue, trueValue, falseValue)
}
