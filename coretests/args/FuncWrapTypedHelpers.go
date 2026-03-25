package args

import "reflect"

// Signature checker methods — detect common function return patterns.
// These check the wrapped function's return types without invoking it.

// IsBoolFunc returns true if the function returns exactly one value of type bool.
//
// Matches: func(...) bool
func (it *FuncWrap[T]) IsBoolFunc() bool {
	if it.IsInvalid() {
		return false
	}

	return it.OutArgsCount() == 1 &&
		it.rvType.Out(0).Kind() == reflect.Bool
}

// IsErrorFunc returns true if the function returns exactly one value of type error.
//
// Matches: func(...) error
func (it *FuncWrap[T]) IsErrorFunc() bool {
	if it.IsInvalid() {
		return false
	}

	return it.OutArgsCount() == 1 &&
		it.rvType.Out(0).Implements(errorType)
}

// IsStringFunc returns true if the function returns exactly one value of type string.
//
// Matches: func(...) string
func (it *FuncWrap[T]) IsStringFunc() bool {
	if it.IsInvalid() {
		return false
	}

	return it.OutArgsCount() == 1 &&
		it.rvType.Out(0).Kind() == reflect.String
}

// IsAnyFunc returns true if the function returns exactly one value of any type.
//
// Matches: func(...) T (single return)
func (it *FuncWrap[T]) IsAnyFunc() bool {
	if it.IsInvalid() {
		return false
	}

	return it.OutArgsCount() == 1
}

// IsValueErrorFunc returns true if the function returns exactly two values
// where the second is an error.
//
// Matches: func(...) (T, error)
func (it *FuncWrap[T]) IsValueErrorFunc() bool {
	if it.IsInvalid() {
		return false
	}

	return it.OutArgsCount() == 2 &&
		it.rvType.Out(1).Implements(errorType)
}

// IsAnyErrorFunc returns true if the function returns exactly two values
// where the second implements error. Alias for IsValueErrorFunc.
//
// Matches: func(...) (any, error)
func (it *FuncWrap[T]) IsAnyErrorFunc() bool {
	return it.IsValueErrorFunc()
}

// IsVoidFunc returns true if the function returns no values.
//
// Matches: func(...)
func (it *FuncWrap[T]) IsVoidFunc() bool {
	if it.IsInvalid() {
		return false
	}

	return it.OutArgsCount() == 0
}

// Typed invoke helpers — invoke and cast the result to a specific Go type.

// InvokeAsBool invokes the function and returns the first result as a bool.
// Returns false if invocation fails or the result is not a bool.
func (it *FuncWrap[T]) InvokeAsBool(args ...any) (bool, error) {
	results, err := it.Invoke(args...)
	if err != nil {
		return false, err
	}

	if len(results) == 0 {
		return false, nil
	}

	val, ok := results[0].(bool)
	if !ok {
		return false, nil
	}

	return val, nil
}

// InvokeAsError invokes the function and returns the first result as an error.
// Returns nil if invocation succeeds and the function returned nil error.
// The second return is the processing error from invocation itself.
func (it *FuncWrap[T]) InvokeAsError(args ...any) (funcErr, processingErr error) {
	results, err := it.Invoke(args...)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	if results[0] == nil {
		return nil, nil
	}

	funcErr, ok := results[0].(error)
	if !ok {
		return nil, nil
	}

	return funcErr, nil
}

// InvokeAsString invokes the function and returns the first result as a string.
// Returns empty string if invocation fails or the result is not a string.
func (it *FuncWrap[T]) InvokeAsString(args ...any) (string, error) {
	results, err := it.Invoke(args...)
	if err != nil {
		return "", err
	}

	if len(results) == 0 {
		return "", nil
	}

	val, ok := results[0].(string)
	if !ok {
		return "", nil
	}

	return val, nil
}

// InvokeAsAny invokes the function and returns the first result as any.
// Returns nil if invocation fails or there are no results.
func (it *FuncWrap[T]) InvokeAsAny(args ...any) (any, error) {
	results, err := it.Invoke(args...)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	return results[0], nil
}

// InvokeAsAnyError invokes the function and returns the first result as any
// and the second result as error. For functions returning (T, error).
// Returns (nil, nil, processingErr) if invocation itself fails.
func (it *FuncWrap[T]) InvokeAsAnyError(args ...any) (result any, funcErr error, processingErr error) {
	results, err := it.Invoke(args...)
	if err != nil {
		return nil, nil, err
	}

	if len(results) == 0 {
		return nil, nil, nil
	}

	first := results[0]

	if len(results) < 2 || results[1] == nil {
		return first, nil, nil
	}

	funcErr, _ = results[1].(error)

	return first, funcErr, nil
}

// errorType is cached for signature checking against the error interface.
var errorType = reflect.TypeOf((*error)(nil)).Elem()
