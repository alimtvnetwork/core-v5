package conditional

// Functions is the generic version of AnyFunctions.
// It selects the appropriate function slice based on the condition.
//
// Usage:
//
//	funcs := conditional.Functions[string](true, trueFuncs, falseFuncs)
func Functions[T any](
	isTrue bool,
	trueValueFunctions, falseValueFunctions []func() (
		result T,
		isTake,
		isBreak bool,
	),
) []func() (
	result T,
	isTake,
	isBreak bool,
) {
	if isTrue {
		return trueValueFunctions
	}

	return falseValueFunctions
}
