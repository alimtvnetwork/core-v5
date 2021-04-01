package conditional

func Interfaces(
	isTrue bool,
	trueValue, falseValue []interface{},
) []interface{} {
	if isTrue {
		return trueValue
	}

	return falseValue
}
