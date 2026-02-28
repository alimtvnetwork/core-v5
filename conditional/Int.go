package conditional

// Deprecated: Use If[int] instead.
func Int(
	isTrue bool,
	trueValue, falseValue int,
) int {
	return If[int](isTrue, trueValue, falseValue)
}
