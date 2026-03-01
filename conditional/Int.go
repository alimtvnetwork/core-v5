package conditional

// Deprecated: Use IfInt instead.
func Int(
	isTrue bool,
	trueValue, falseValue int,
) int {
	return If[int](isTrue, trueValue, falseValue)
}
