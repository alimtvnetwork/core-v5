package conditional

func Byte(
	isTrue bool,
	trueValue, falseValue byte,
) byte {
	if isTrue {
		return trueValue
	}

	return falseValue
}
