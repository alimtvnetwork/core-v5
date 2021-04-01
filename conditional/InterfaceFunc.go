package conditional

func InterfaceFunc(
	isTrue bool,
	trueValueFunc, falseValueFunc func() interface{},
) interface{} {
	if isTrue {
		return trueValueFunc()
	}

	return falseValueFunc()
}
