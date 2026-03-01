package conditional

// Deprecated: Use IfBool instead.
func Bool(
	isTrue bool,
	trueValue, falseValue bool,
) bool {
	return If[bool](isTrue, trueValue, falseValue)
}
