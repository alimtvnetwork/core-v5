package conditional

// Deprecated: Use IfByte instead.
func Byte(
	isTrue bool,
	trueValue, falseValue byte,
) byte {
	return If[byte](isTrue, trueValue, falseValue)
}
