package args

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/internal/convertinternal"
	"gitlab.com/auk-go/core/internal/reflectinternal"
	"gitlab.com/auk-go/core/internal/trydo"
)

func (it *FuncWrap) VoidCallNoReturn(args ...any) (processingErr error) {
	it.MustBeValid()
	_, err := it.Invoke(args)

	return err
}

func (it *FuncWrap) InvokeMust(args ...any) []any {
	results, err := it.Invoke(args...)

	if err != nil {
		panic(err)
	}

	return results
}

func (it *FuncWrap) Invoke(args ...any) (results []any, processingErr error) {
	return it.InvokeSkip(codestack.Skip1, args...)
}

func (it *FuncWrap) InvokeSkip(skipStack int, args ...any) (results []any, processingErr error) {
	firstErr := it.ValidationError()

	if firstErr != nil {
		return nil, firstErr
	}

	argsValidationErr := it.ValidateMethodArgs(args)

	if argsValidationErr != nil {
		return nil, argsValidationErr
	}

	rvs := argsToRvFunc(args)
	var resultsRawValues []reflect.Value
	exception := trydo.WrapPanic(func() { resultsRawValues = it.rv.Call(rvs) })

	if exception != nil {
		toMsg := convertinternal.AnyTo.SmartString(exception)
		finalError := fmt.Errorf(
			"%s - func invoke failed\nstack-trace:%s\nerr:%s",
			it.GetFuncName(),
			reflectinternal.CodeStack.StacksString(codestack.Skip1+skipStack),
			toMsg,
		)

		return rvToInterfacesFunc(resultsRawValues), finalError
	}

	return rvToInterfacesFunc(resultsRawValues), nil
}

func (it *FuncWrap) VoidCall() ([]any, error) {
	return it.Invoke()
}

func (it *FuncWrap) GetFirstResponseOfInvoke(args ...any) (firstResponse any, err error) {
	result, err := it.InvokeResultOfIndex(0, args...)

	if err != nil {
		return nil, err
	}

	return result, err
}

func (it *FuncWrap) InvokeResultOfIndex(index int, args ...any) (firstResponse any, err error) {
	results, err := it.Invoke(args...)

	if err != nil {
		return nil, err
	}

	return results[index], err
}

func (it *FuncWrap) InvokeError(args ...any) (funcErr, processingErr error) {
	result, err := it.GetFirstResponseOfInvoke(args...)

	if err != nil {
		return nil, err
	}

	return result.(error), err
}

// InvokeFirstAndError useful for method which looks like ReflectMethod() (something, error)
func (it *FuncWrap) InvokeFirstAndError(args ...any) (firstResponse any, funcErr, processingErr error) {
	results, processingErr := it.Invoke(args...)

	if processingErr != nil {
		return nil, nil, processingErr
	}

	if len(results) <= 1 {
		return results, nil, errors.New(it.GetFuncName() + " doesn't return at least 2 return args")
	}

	first := results[0]
	second := results[1].(error)

	return first, second, processingErr
}

func (it *FuncWrap) argsCountMismatchErrorMessage(
	expectedCount int,
	given int,
	args []any,
) string {
	expectedTypes := it.GetInArgsTypesNames()
	expectedToNames := strings.Join(expectedTypes, newLineSpaceIndent)
	actualTypes := reflectinternal.Converter.InterfacesToTypesNamesWithValues(args)
	actualTypesName := strings.Join(actualTypes, newLineSpaceIndent)

	return fmt.Sprintf(
		"%s [Func] =>\n  arguments count doesn't match for - count:\n    expected : %d\n    given    : %d\n  expected types listed :\n    - %s\n  actual given types list :\n    - %s",
		it.Name,
		expectedCount,
		given,
		expectedToNames,
		actualTypesName,
	)
}
