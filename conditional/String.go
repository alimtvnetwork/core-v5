package conditional

func String(
	isTrue bool,
	trueValue, falseValue string,
) string {
	if isTrue {
		return trueValue
	}

	return falseValue
}
