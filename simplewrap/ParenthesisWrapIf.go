package simplewrap

func ParenthesisWrapIf(
	isSquareWrap bool,
	source any,
) string {
	if !isSquareWrap {
		return toString(source)
	}

	return ParenthesisWrap(source)
}
