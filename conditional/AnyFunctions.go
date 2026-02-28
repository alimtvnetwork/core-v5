package conditional

func AnyFunctions(
	isTrue bool,
	trueValueFunctions, falseValueFunctions []func() (
		result any,
		isTake,
		isBreak bool,
	),
) []func() (
	result any,
	isTake,
	isBreak bool,
) {
	if isTrue {
		return trueValueFunctions
	}

	return falseValueFunctions
}
