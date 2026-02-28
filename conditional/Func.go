package conditional

func Func(
	isTrue bool,
	trueValueFunc, falseValueFunc func() any,
) func() any {
	if isTrue {
		return trueValueFunc
	}

	return falseValueFunc
}
