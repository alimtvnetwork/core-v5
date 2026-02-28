package conditional

func InterfaceFunc(
	isTrue bool,
	trueValueFunc, falseValueFunc func() any,
) any {
	if isTrue {
		return trueValueFunc()
	}

	return falseValueFunc()
}
