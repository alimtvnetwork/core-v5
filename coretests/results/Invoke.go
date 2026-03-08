package results

import (
	"fmt"
	"reflect"
)

// InvokeWithPanicRecovery calls the given function reference with the
// provided receiver and args, recovering from any panic.
//
// This is the core invocation engine used by CaseNilSafe.
//
// funcRef must be a method expression like (*MyStruct).Method.
// receiver is the value to bind as the first argument (may be nil).
// args are additional arguments to pass after the receiver.
func InvokeWithPanicRecovery(
	funcRef any,
	receiver any,
	args ...any,
) ResultAny {
	var result ResultAny

	func() {
		defer func() {
			if r := recover(); r != nil {
				result.Panicked = true
				result.PanicValue = r
			}
		}()

		rv := reflect.ValueOf(funcRef)

		if rv.Kind() != reflect.Func {
			result.Panicked = true
			result.PanicValue = fmt.Sprintf(
				"funcRef is not a function: %T",
				funcRef,
			)

			return
		}

		callArgs := buildCallArgs(rv, receiver, args)
		returnValues := rv.Call(callArgs)
		result = extractResult(returnValues)
	}()

	return result
}

// buildCallArgs constructs the reflect.Value slice for the function call.
func buildCallArgs(
	rv reflect.Value,
	receiver any,
	args []any,
) []reflect.Value {
	funcType := rv.Type()
	callArgs := make([]reflect.Value, 0, 1+len(args))

	if receiver == nil && funcType.NumIn() > 0 {
		firstParam := funcType.In(0)
		callArgs = append(callArgs, reflect.Zero(firstParam))
	} else if receiver != nil {
		callArgs = append(callArgs, reflect.ValueOf(receiver))
	}

	for _, arg := range args {
		callArgs = append(callArgs, reflect.ValueOf(arg))
	}

	return callArgs
}

// extractResult converts reflect return values into a ResultAny.
func extractResult(returnValues []reflect.Value) ResultAny {
	var result ResultAny

	if len(returnValues) == 0 {
		return result
	}

	first := returnValues[0].Interface()
	result.Value = first

	if len(returnValues) > 1 {
		last := returnValues[len(returnValues)-1]

		if last.Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) &&
			!last.IsNil() {
			result.Error = last.Interface().(error)
		}
	}

	return result
}
