package conditional

// Deprecated: Use If[string] instead.
func String(
	isTrue bool,
	trueValue, falseValue string,
) string {
	return If[string](isTrue, trueValue, falseValue)
}
