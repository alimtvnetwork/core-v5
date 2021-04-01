package conditional

func BytesPtr(
	isTrue bool,
	trueValue, falseValue *[]byte,
) *[]byte {
	if isTrue {
		return trueValue
	}

	return falseValue
}
