package conditional

func Interfaces(
	isTrue bool,
	trueValue, falseValue []any,
) []any {
	if isTrue {
		return trueValue
	}

	return falseValue
}
