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
//
// Edge cases handled:
//   - nil funcRef → Panicked with descriptive PanicValue
//   - void methods → ReturnCount=0, Value=nil
//   - multi-return → AllResults populated, last error extracted
//   - value receivers with nil pointer → natural Go panic recovered
//   - non-error second return → AllResults populated, Error stays nil
//   - interface returns with nil → properly extracted as nil
func InvokeWithPanicRecovery(
	funcRef any,
	receiver any,
	args ...any,
) ResultAny {
	var result ResultAny

	if funcRef == nil {
		result.Panicked = true
		result.PanicValue = "funcRef is nil"

		return result
	}

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
//
// For nil receivers, it creates a typed zero value matching the first
// parameter type. This produces a typed nil pointer for pointer receivers,
// which is exactly what nil-safety tests need.
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
		if arg == nil {
			// Preserve nil as untyped — caller must use typed nil if needed
			callArgs = append(callArgs, reflect.Zero(reflect.TypeOf((*any)(nil)).Elem()))
		} else {
			callArgs = append(callArgs, reflect.ValueOf(arg))
		}
	}

	return callArgs
}

// extractResult converts reflect return values into a ResultAny.
//
// Handles:
//   - 0 returns (void) → ReturnCount=0, Value=nil
//   - 1 return → Value = first, ReturnCount=1
//   - 2+ returns → Value = first, Error = last if implements error
//   - Interface nil → properly detected via IsNil()
func extractResult(returnValues []reflect.Value) ResultAny {
	var result ResultAny

	result.ReturnCount = len(returnValues)
	result.AllResults = make([]any, len(returnValues))

	for i, rv := range returnValues {
		result.AllResults[i] = safeInterface(rv)
	}

	if len(returnValues) == 0 {
		return result
	}

	result.Value = safeInterface(returnValues[0])

	if len(returnValues) > 1 {
		last := returnValues[len(returnValues)-1]
		result.Error = extractErrorFromValue(last)
	}

	return result
}

// safeInterface extracts the interface value from a reflect.Value,
// handling nil interfaces and invalid values gracefully.
func safeInterface(rv reflect.Value) any {
	if !rv.IsValid() {
		return nil
	}

	// Check for nil pointers, interfaces, maps, slices, channels, funcs
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Map,
		reflect.Slice, reflect.Chan, reflect.Func:

		if rv.IsNil() {
			return nil
		}
	}

	return rv.Interface()
}

// errorType is cached for signature checking.
var errorType = reflect.TypeOf((*error)(nil)).Elem()

// extractErrorFromValue checks if a reflect.Value implements error
// and returns it. Returns nil if the value is not an error or is nil.
func extractErrorFromValue(rv reflect.Value) error {
	if !rv.IsValid() {
		return nil
	}

	if !rv.Type().Implements(errorType) {
		return nil
	}

	if rv.Kind() == reflect.Interface && rv.IsNil() {
		return nil
	}

	if rv.Kind() == reflect.Ptr && rv.IsNil() {
		return nil
	}

	funcErr, ok := rv.Interface().(error)

	if !ok {
		return nil
	}

	return funcErr
}
