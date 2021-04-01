package conditional

func Int(
	isTrue bool,
	trueValue, falseValue int,
) int {
	if isTrue {
		return trueValue
	}

	return falseValue
}
