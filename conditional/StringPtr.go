package conditional

// Deprecated: Use IfPtr[string] instead.
func StringPtr(
	isTrue bool,
	trueValue, falseValue *string,
) *string {
	return IfPtr[string](isTrue, trueValue, falseValue)
}
