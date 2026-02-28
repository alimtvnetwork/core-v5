package conditional

// Deprecated: Use If[any] instead.
func Interface(
	isTrue bool,
	trueValue, falseValue any,
) any {
	return If[any](isTrue, trueValue, falseValue)
}
