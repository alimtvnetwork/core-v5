package conditional

func Func(
	isTrue bool,
	trueValueFunc, falseValueFunc func() interface{},
) func() interface{} {
	if isTrue {
		return trueValueFunc
	}

	return falseValueFunc
}
