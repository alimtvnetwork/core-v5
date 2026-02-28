package conditional

// FunctionsExecuteResults is the generic version of AnyFunctionsExecuteResults.
// It executes the appropriate set of functions based on the condition
// and collects results of type T.
//
// Each function returns (result T, isTake bool, isBreak bool):
//   - isTake: include this result in the output slice
//   - isBreak: stop executing further functions
//
// Usage:
//
//	results := conditional.FunctionsExecuteResults[int](true, trueFuncs, falseFuncs)
func FunctionsExecuteResults[T any](
	isTrue bool,
	trueValueFunctions, falseValueFunctions []func() (
		result T,
		isTake,
		isBreak bool,
	),
) []T {
	if isTrue {
		return executeFunctions[T](trueValueFunctions)
	}

	return executeFunctions[T](falseValueFunctions)
}

func executeFunctions[T any](
	functions []func() (
		result T,
		isTake,
		isBreak bool,
	),
) []T {
	if len(functions) == 0 {
		return nil
	}

	results := make([]T, 0, len(functions))
	for _, curFunc := range functions {
		if curFunc == nil {
			continue
		}

		result, isTake, isBreak := curFunc()

		if isTake {
			results = append(results, result)
		}

		if isBreak {
			return results
		}
	}

	return results
}
