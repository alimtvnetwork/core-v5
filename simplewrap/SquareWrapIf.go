package simplewrap

func SquareWrapIf(
	isSquareWrap bool,
	source any,
) string {
	isSkipWrap := !isSquareWrap

	if isSkipWrap {
		return toString(source)
	}

	return SquareWrap(source)
}
