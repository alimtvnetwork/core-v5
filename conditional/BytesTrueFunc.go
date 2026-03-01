package conditional

// Deprecated: Use IfTrueFuncBytes instead.
func BytesTrueFunc(
	isTrue bool,
	trueValueFunc func() []byte,
) []byte {
	return IfTrueFunc[[]byte](isTrue, trueValueFunc)
}
