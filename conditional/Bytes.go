package conditional

func Bytes(
	isTrue bool,
	trueValue, falseValue []byte,
) []byte {
	if isTrue {
		return trueValue
	}

	return falseValue
}
