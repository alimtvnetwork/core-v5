package conditional

func Interface(
	isTrue bool,
	trueValue, falseValue interface{},
) interface{} {
	if isTrue {
		return trueValue
	}

	return falseValue
}
