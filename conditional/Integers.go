package conditional

func Integers(
	isTrue bool,
	trueValue, falseValue []int,
) []int {
	if isTrue {
		return trueValue
	}

	return falseValue
}
