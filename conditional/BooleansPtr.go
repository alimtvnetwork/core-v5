package conditional

func BooleansPtr(
	isTrue bool,
	trueValue, falseValue *[]bool,
) *[]bool {
	if isTrue {
		return trueValue
	}

	return falseValue
}
