package conditional

func InterfacesPtr(
	isTrue bool,
	trueValue, falseValue *[]any,
) *[]any {
	if isTrue {
		return trueValue
	}

	return falseValue
}
