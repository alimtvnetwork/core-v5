package conditional

func StringPtr(
	isTrue bool,
	trueValue, falseValue *string,
) *string {
	if isTrue {
		return trueValue
	}

	return falseValue
}
