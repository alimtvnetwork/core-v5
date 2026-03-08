package results

import (
	"fmt"
	"reflect"
)

// Result is the base typed result for a single-return function invocation.
//
// T is the type of the primary return value.
//
// Fields:
//   - Value      — the primary return value
//   - Error      — error returned by the function (or wrapped panic)
//   - Panicked   — true if the invocation recovered from a panic
//   - PanicValue — the raw value recovered from the panic (nil if no panic)
type Result[T any] struct {
	Value      T
	Error      error
	Panicked   bool
	PanicValue any
}

// IsSafe returns true if no panic occurred and no error was returned.
func (it Result[T]) IsSafe() bool {
	return !it.Panicked && it.Error == nil
}

// HasError returns true if the Error field is non-nil.
func (it Result[T]) HasError() bool {
	return it.Error != nil
}

// HasPanicked returns true if the invocation recovered from a panic.
func (it Result[T]) HasPanicked() bool {
	return it.Panicked
}

// IsResult checks whether Value matches the given expected value
// using fmt.Sprintf for comparison.
func (it Result[T]) IsResult(expected any) bool {
	return fmt.Sprintf("%v", it.Value) == fmt.Sprintf("%v", expected)
}

// IsResultTypeOf checks whether Value is assignable to the given type.
func (it Result[T]) IsResultTypeOf(expected any) bool {
	if expected == nil {
		return reflect.ValueOf(it.Value).IsZero()
	}

	expectedType := reflect.TypeOf(expected)
	actualType := reflect.TypeOf(it.Value)

	if actualType == nil {
		return false
	}

	return actualType.AssignableTo(expectedType)
}

// IsError checks whether the Error field matches the given error string.
// Returns false if Error is nil.
func (it Result[T]) IsError(msg string) bool {
	if it.Error == nil {
		return false
	}

	return it.Error.Error() == msg
}

// ValueString returns the Value formatted via %v.
func (it Result[T]) ValueString() string {
	return fmt.Sprintf("%v", it.Value)
}

// String returns a human-readable summary of the result.
func (it Result[T]) String() string {
	if it.Panicked {
		return fmt.Sprintf(
			"Result{panicked: %v, panicValue: %v}",
			it.Panicked,
			it.PanicValue,
		)
	}

	if it.Error != nil {
		return fmt.Sprintf(
			"Result{value: %v, error: %s}",
			it.Value,
			it.Error.Error(),
		)
	}

	return fmt.Sprintf(
		"Result{value: %v}",
		it.Value,
	)
}
