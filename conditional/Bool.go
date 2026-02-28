package conditional

// Deprecated: Use If[bool] instead.
func Bool(
	isTrue bool,
	trueValue, falseValue bool,
) bool {
	return If[bool](isTrue, trueValue, falseValue)
}
