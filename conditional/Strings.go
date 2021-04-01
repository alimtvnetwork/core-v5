package conditional

func Strings(
	isTrue bool,
	trueValue, falseValue []string,
) []string {
	if isTrue {
		return trueValue
	}

	return falseValue
}
