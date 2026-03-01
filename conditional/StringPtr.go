package conditional

// Deprecated: Use IfPtrString instead.
func StringPtr(
	isTrue bool,
	trueValue, falseValue *string,
) *string {
	return IfPtr[string](isTrue, trueValue, falseValue)
}
