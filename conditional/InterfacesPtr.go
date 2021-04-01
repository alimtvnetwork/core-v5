package conditional

func InterfacesPtr(
	isTrue bool,
	trueValue, falseValue *[]interface{},
) *[]interface{} {
	if isTrue {
		return trueValue
	}

	return falseValue
}
