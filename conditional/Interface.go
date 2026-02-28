package conditional

func Interface(
	isTrue bool,
	trueValue, falseValue any,
) any {
	if isTrue {
		return trueValue
	}

	return falseValue
}
