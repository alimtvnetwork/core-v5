package conditional

func StringFunc(
	isTrue bool,
	trueValueFunc, falseValueFunc func() string,
) string {
	if isTrue {
		return trueValueFunc()
	}

	return falseValueFunc()
}
