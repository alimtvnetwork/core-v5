package conditional

func StringsIndexVal(
	isTrue bool,
	slice []string,
	trueValue, falseValue int,
) string {
	if isTrue {
		return slice[trueValue]
	}

	return slice[falseValue]
}
