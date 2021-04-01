package conditional

func Booleans(
	isTrue bool,
	trueValue, falseValue []bool,
) []bool {
	if isTrue {
		return trueValue
	}

	return falseValue
}
