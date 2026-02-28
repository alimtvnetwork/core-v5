package coredynamic

import (
	"encoding/json"
	"errors"
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
)

// TypedSimpleResult is the generic version of SimpleResult.
//
// Unlike SimpleResult (which embeds Dynamic and wraps any),
// TypedSimpleResult[T] provides compile-time type safety for the result payload.
//
// Usage:
//
//	result := coredynamic.NewTypedSimpleResultValid[User](user)
//	fmt.Println(result.Data().Name)  // "Alice" — strongly typed
//	fmt.Println(result.IsValid())    // true
type TypedSimpleResult[T any] struct {
	innerData T
	isValid   bool
	message   string
	err       error
}

// NewTypedSimpleResult creates a TypedSimpleResult with data, validity, and message.
func NewTypedSimpleResult[T any](
	result T,
	isValid bool,
	message string,
) *TypedSimpleResult[T] {
	return &TypedSimpleResult[T]{
		innerData: result,
		isValid:   isValid,
		message:   message,
	}
}

// NewTypedSimpleResultValid creates a valid TypedSimpleResult with empty message.
func NewTypedSimpleResultValid[T any](result T) *TypedSimpleResult[T] {
	return &TypedSimpleResult[T]{
		innerData: result,
		isValid:   true,
		message:   constants.EmptyString,
	}
}

// InvalidTypedSimpleResult creates an invalid TypedSimpleResult with a message.
func InvalidTypedSimpleResult[T any](message string) *TypedSimpleResult[T] {
	return &TypedSimpleResult[T]{
		isValid: false,
		message: message,
	}
}

// InvalidTypedSimpleResultNoMessage creates an invalid TypedSimpleResult with empty message.
func InvalidTypedSimpleResultNoMessage[T any]() *TypedSimpleResult[T] {
	return &TypedSimpleResult[T]{
		isValid: false,
		message: constants.EmptyString,
	}
}

// =============================================================================
// Core accessors
// =============================================================================

// Data returns the strongly-typed result data.
func (it *TypedSimpleResult[T]) Data() T {
	return it.innerData
}

// Result is an alias for Data, mirroring SimpleResult.Result.
func (it *TypedSimpleResult[T]) Result() T {
	return it.innerData
}

// IsValid returns whether the result is valid.
func (it *TypedSimpleResult[T]) IsValid() bool {
	return it != nil && it.isValid
}

// IsInvalid returns whether the result is invalid.
func (it *TypedSimpleResult[T]) IsInvalid() bool {
	return it == nil || !it.isValid
}

// Message returns the result message (typically an error/validation message).
func (it *TypedSimpleResult[T]) Message() string {
	if it == nil {
		return constants.EmptyString
	}

	return it.message
}

// String returns a string representation of the inner data.
func (it *TypedSimpleResult[T]) String() string {
	if it == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf("%v", it.innerData)
}

// InvalidError returns an error if the result has a message, otherwise nil.
func (it *TypedSimpleResult[T]) InvalidError() error {
	if it == nil {
		return nil
	}

	if it.err != nil {
		return it.err
	}

	if it.message == constants.EmptyString {
		return nil
	}

	it.err = errors.New(it.message)

	return it.err
}

// =============================================================================
// JSON operations
// =============================================================================

// JsonBytes serializes the inner data to JSON bytes.
func (it *TypedSimpleResult[T]) JsonBytes() ([]byte, error) {
	return json.Marshal(it.innerData)
}

// JsonResult returns a corejson.Result for the inner data.
func (it *TypedSimpleResult[T]) JsonResult() corejson.Result {
	return corejson.New(it.innerData)
}

// Json returns a corejson.Result.
func (it *TypedSimpleResult[T]) Json() corejson.Result {
	return corejson.New(it.innerData)
}

// JsonPtr returns a pointer to a corejson.Result.
func (it *TypedSimpleResult[T]) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it.innerData)
}

// MarshalJSON implements the json.Marshaler interface.
func (it *TypedSimpleResult[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.innerData)
}

// JsonModel returns the inner data for JSON serialization.
func (it *TypedSimpleResult[T]) JsonModel() T {
	return it.innerData
}

// JsonModelAny returns the inner data as any.
func (it *TypedSimpleResult[T]) JsonModelAny() any {
	return it.innerData
}

// =============================================================================
// GetAs* type assertion methods
// =============================================================================

// GetAsString attempts to retrieve the data as a string.
func (it *TypedSimpleResult[T]) GetAsString() (string, bool) {
	val, ok := any(it.innerData).(string)

	return val, ok
}

// GetAsInt attempts to retrieve the data as an int.
func (it *TypedSimpleResult[T]) GetAsInt() (int, bool) {
	val, ok := any(it.innerData).(int)

	return val, ok
}

// GetAsInt64 attempts to retrieve the data as an int64.
func (it *TypedSimpleResult[T]) GetAsInt64() (int64, bool) {
	val, ok := any(it.innerData).(int64)

	return val, ok
}

// GetAsFloat64 attempts to retrieve the data as a float64.
func (it *TypedSimpleResult[T]) GetAsFloat64() (float64, bool) {
	val, ok := any(it.innerData).(float64)

	return val, ok
}

// GetAsBool attempts to retrieve the data as a bool.
func (it *TypedSimpleResult[T]) GetAsBool() (bool, bool) {
	val, ok := any(it.innerData).(bool)

	return val, ok
}

// GetAsBytes attempts to retrieve the data as []byte.
func (it *TypedSimpleResult[T]) GetAsBytes() ([]byte, bool) {
	val, ok := any(it.innerData).([]byte)

	return val, ok
}

// GetAsStrings attempts to retrieve the data as []string.
func (it *TypedSimpleResult[T]) GetAsStrings() ([]string, bool) {
	val, ok := any(it.innerData).([]string)

	return val, ok
}

// =============================================================================
// Clone and conversion
// =============================================================================

// Clone returns a value copy of the TypedSimpleResult.
func (it *TypedSimpleResult[T]) Clone() TypedSimpleResult[T] {
	if it == nil {
		return TypedSimpleResult[T]{}
	}

	return TypedSimpleResult[T]{
		innerData: it.innerData,
		isValid:   it.isValid,
		message:   it.message,
	}
}

// ClonePtr returns a pointer to a copy of the TypedSimpleResult.
func (it *TypedSimpleResult[T]) ClonePtr() *TypedSimpleResult[T] {
	if it == nil {
		return nil
	}

	cloned := it.Clone()

	return &cloned
}

// ToSimpleResult converts to the non-generic SimpleResult for backward compatibility.
func (it *TypedSimpleResult[T]) ToSimpleResult() *SimpleResult {
	if it == nil {
		return InvalidSimpleResultNoMessage()
	}

	return NewSimpleResult(it.innerData, it.isValid, it.message)
}

// ToTypedDynamic converts to a TypedDynamic[T].
func (it *TypedSimpleResult[T]) ToTypedDynamic() TypedDynamic[T] {
	if it == nil {
		return InvalidTypedDynamic[T]()
	}

	return NewTypedDynamic[T](it.innerData, it.isValid)
}

// ToDynamic converts to the non-generic Dynamic for backward compatibility.
func (it *TypedSimpleResult[T]) ToDynamic() Dynamic {
	if it == nil {
		return InvalidDynamic()
	}

	return NewDynamic(it.innerData, it.isValid)
}
