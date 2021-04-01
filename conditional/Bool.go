package conditional

func Bool(
	isTrue bool,
	trueValue, falseValue bool,
) bool {
	if isTrue {
		return trueValue
	}

	return falseValue
}
