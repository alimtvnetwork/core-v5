package conditional

func IntegersPtr(
	isTrue bool,
	trueValue, falseValue *[]int,
) *[]int {
	if isTrue {
		return trueValue
	}

	return falseValue
}
