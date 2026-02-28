package conditional

import (
	"strconv"

	"gitlab.com/auk-go/core/errcore"
)

// TypedErrorFunctionsExecuteResults is the generic version of ErrorFunctionsExecuteResults.
// It executes the appropriate set of functions based on the condition,
// collects results of type T, and aggregates any errors.
//
// Each function returns (result T, err error).
// Results are collected only when err is nil.
//
// Usage:
//
//	results, err := conditional.TypedErrorFunctionsExecuteResults[int](true, trueFuncs, falseFuncs)
func TypedErrorFunctionsExecuteResults[T any](
	isTrue bool,
	trueValueFunctions []func() (T, error),
	falseValueFunctions []func() (T, error),
) ([]T, error) {
	if isTrue {
		return executeTypedErrorFunctions[T](trueValueFunctions)
	}

	return executeTypedErrorFunctions[T](falseValueFunctions)
}

func executeTypedErrorFunctions[T any](
	functions []func() (T, error),
) ([]T, error) {
	if len(functions) == 0 {
		return nil, nil
	}

	results := make([]T, 0, len(functions))
	var sliceErr []string

	for index, currentFunction := range functions {
		if currentFunction == nil {
			continue
		}

		result, err := currentFunction()

		if err != nil {
			sliceErr = append(sliceErr, err.Error()+"- index of - "+strconv.Itoa(index))
			continue
		}

		results = append(results, result)
	}

	return results, errcore.SliceToError(sliceErr)
}
