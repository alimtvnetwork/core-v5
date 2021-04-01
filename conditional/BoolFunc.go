package conditional

func BoolFunc(
	isTrue bool,
	trueValueFunc, falseValueFunc func() bool,
) bool {
	if isTrue {
		return trueValueFunc()
	}

	return falseValueFunc()
}
