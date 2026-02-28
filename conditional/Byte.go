package conditional

// Deprecated: Use If[byte] instead.
func Byte(
	isTrue bool,
	trueValue, falseValue byte,
) byte {
	return If[byte](isTrue, trueValue, falseValue)
}
