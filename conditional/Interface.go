package conditional

// Deprecated: Use IfAny instead.
func Interface(
	isTrue bool,
	trueValue, falseValue any,
) any {
	return If[any](isTrue, trueValue, falseValue)
}
