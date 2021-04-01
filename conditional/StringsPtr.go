package conditional

func StringsPtr(
	isTrue bool,
	trueValue, falseValue *[]string,
) *[]string {
	if isTrue {
		return trueValue
	}

	return falseValue
}
