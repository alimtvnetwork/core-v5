package conditional

func IntegersPtrFunc(
	isTrue bool,
	trueValueFunc, falseValueFunc func() *[]int,
) *[]int {
	if isTrue {
		return trueValueFunc()
	}

	return falseValueFunc()
}
