package simplewrap

func SquareWrapIf(
	isSquareWrap bool,
	source any,
) string {
	if !isSquareWrap {
		return toString(source)
	}
	
	return SquareWrap(source)
}
