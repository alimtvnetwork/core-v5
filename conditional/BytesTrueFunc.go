package conditional

// Deprecated: Use IfTrueFunc[[]byte] instead.
func BytesTrueFunc(
	isTrue bool,
	trueValueFunc func() []byte,
) []byte {
	return IfTrueFunc[[]byte](isTrue, trueValueFunc)
}
