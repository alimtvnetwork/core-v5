package conditional

func AnyFunctionsExecuteResults(
	isTrue bool,
	trueValueFunctions, falseValueFunctions []func() (
		result any,
		isTake,
		isBreak bool,
	),
) []any {
	if isTrue {
		return executeAnyFunctions(trueValueFunctions)
	}

	return executeAnyFunctions(falseValueFunctions)
}
