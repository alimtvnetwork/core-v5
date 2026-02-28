package coredynamic

import (
	"encoding/json"
	"errors"
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
)

// TypedSimpleRequest is a generic, strongly-typed version of SimpleRequest.
//
// T represents the request payload type, eliminating the need for
// type assertions when accessing request data.
//
// Usage:
//
//	type UserInput struct { Name string; Age int }
//	req := coredynamic.NewTypedSimpleRequest[UserInput](
//	    UserInput{Name: "Alice", Age: 30}, true, "",
//	)
//	fmt.Println(req.Data().Name)  // "Alice" — fully typed
//	fmt.Println(req.IsValid())    // true
type TypedSimpleRequest[T any] struct {
	innerData T
	isValid   bool
	message   string
	err       error
}

// NewTypedSimpleRequest creates a TypedSimpleRequest with the given data, validity, and message.
func NewTypedSimpleRequest[T any](
	request T,
	isValid bool,
	message string,
) *TypedSimpleRequest[T] {
	return &TypedSimpleRequest[T]{
		innerData: request,
		isValid:   isValid,
		message:   message,
	}
}

// NewTypedSimpleRequestValid creates a valid TypedSimpleRequest with empty message.
func NewTypedSimpleRequestValid[T any](request T) *TypedSimpleRequest[T] {
	return &TypedSimpleRequest[T]{
		innerData: request,
		isValid:   true,
		message:   constants.EmptyString,
	}
}

// InvalidTypedSimpleRequest creates an invalid TypedSimpleRequest with the given message.
func InvalidTypedSimpleRequest[T any](message string) *TypedSimpleRequest[T] {
	return &TypedSimpleRequest[T]{
		isValid: false,
		message: message,
	}
}

// InvalidTypedSimpleRequestNoMessage creates an invalid TypedSimpleRequest with empty message.
func InvalidTypedSimpleRequestNoMessage[T any]() *TypedSimpleRequest[T] {
	return &TypedSimpleRequest[T]{
		isValid: false,
		message: constants.EmptyString,
	}
}

// =============================================================================
// Core accessors
// =============================================================================

// Data returns the strongly-typed request data.
func (it *TypedSimpleRequest[T]) Data() T {
	return it.innerData
}

// Request is an alias for Data, mirroring SimpleRequest.Request().
func (it *TypedSimpleRequest[T]) Request() T {
	return it.innerData
}

// Value is an alias for Data, mirroring SimpleRequest.Value().
func (it *TypedSimpleRequest[T]) Value() T {
	return it.innerData
}

// IsValid returns whether the request is valid.
func (it *TypedSimpleRequest[T]) IsValid() bool {
	return it != nil && it.isValid
}

// IsInvalid returns whether the request is invalid.
func (it *TypedSimpleRequest[T]) IsInvalid() bool {
	return it == nil || !it.isValid
}

// Message returns the request's message (typically an error/validation message).
func (it *TypedSimpleRequest[T]) Message() string {
	if it == nil {
		return constants.EmptyString
	}

	return it.message
}

// String returns a string representation of the inner data.
func (it *TypedSimpleRequest[T]) String() string {
	if it == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf("%v", it.innerData)
}

// InvalidError returns an error if the request has a message, otherwise nil.
func (it *TypedSimpleRequest[T]) InvalidError() error {
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
func (it *TypedSimpleRequest[T]) JsonBytes() ([]byte, error) {
	return json.Marshal(it.innerData)
}

// JsonResult returns a corejson.Result for the inner data.
func (it *TypedSimpleRequest[T]) JsonResult() corejson.Result {
	return corejson.New(it.innerData)
}

// Json returns a corejson.Result (alias matching SimpleRequest pattern).
func (it *TypedSimpleRequest[T]) Json() corejson.Result {
	return corejson.New(it.innerData)
}

// JsonPtr returns a pointer to a corejson.Result.
func (it *TypedSimpleRequest[T]) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it.innerData)
}

// MarshalJSON implements the json.Marshaler interface.
func (it *TypedSimpleRequest[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.innerData)
}

// JsonModel returns the inner data for JSON serialization.
func (it *TypedSimpleRequest[T]) JsonModel() T {
	return it.innerData
}

// JsonModelAny returns the inner data as any.
func (it *TypedSimpleRequest[T]) JsonModelAny() any {
	return it.innerData
}

// =============================================================================
// GetAs* type assertion methods
// =============================================================================

// GetAsString attempts to retrieve the data as a string.
// Returns the string value and whether the conversion was successful.
func (it *TypedSimpleRequest[T]) GetAsString() (string, bool) {
	val, ok := any(it.innerData).(string)

	return val, ok
}

// GetAsInt attempts to retrieve the data as an int.
func (it *TypedSimpleRequest[T]) GetAsInt() (int, bool) {
	val, ok := any(it.innerData).(int)

	return val, ok
}

// GetAsInt64 attempts to retrieve the data as an int64.
func (it *TypedSimpleRequest[T]) GetAsInt64() (int64, bool) {
	val, ok := any(it.innerData).(int64)

	return val, ok
}

// GetAsFloat64 attempts to retrieve the data as a float64.
func (it *TypedSimpleRequest[T]) GetAsFloat64() (float64, bool) {
	val, ok := any(it.innerData).(float64)

	return val, ok
}

// GetAsFloat32 attempts to retrieve the data as a float32.
func (it *TypedSimpleRequest[T]) GetAsFloat32() (float32, bool) {
	val, ok := any(it.innerData).(float32)

	return val, ok
}

// GetAsBool attempts to retrieve the data as a bool.
func (it *TypedSimpleRequest[T]) GetAsBool() (bool, bool) {
	val, ok := any(it.innerData).(bool)

	return val, ok
}

// GetAsBytes attempts to retrieve the data as []byte.
func (it *TypedSimpleRequest[T]) GetAsBytes() ([]byte, bool) {
	val, ok := any(it.innerData).([]byte)

	return val, ok
}

// GetAsStrings attempts to retrieve the data as []string.
func (it *TypedSimpleRequest[T]) GetAsStrings() ([]string, bool) {
	val, ok := any(it.innerData).([]string)

	return val, ok
}

// =============================================================================
// Clone and conversion
// =============================================================================

// Clone returns a copy of the TypedSimpleRequest.
func (it *TypedSimpleRequest[T]) Clone() *TypedSimpleRequest[T] {
	if it == nil {
		return nil
	}

	return &TypedSimpleRequest[T]{
		innerData: it.innerData,
		isValid:   it.isValid,
		message:   it.message,
	}
}

// ToSimpleRequest converts to the non-generic SimpleRequest for backward compatibility.
func (it *TypedSimpleRequest[T]) ToSimpleRequest() *SimpleRequest {
	if it == nil {
		return InvalidSimpleRequestNoMessage()
	}

	return NewSimpleRequest(it.innerData, it.isValid, it.message)
}

// ToTypedDynamic converts to a TypedDynamic[T].
func (it *TypedSimpleRequest[T]) ToTypedDynamic() TypedDynamic[T] {
	if it == nil {
		return InvalidTypedDynamic[T]()
	}

	return NewTypedDynamic[T](it.innerData, it.isValid)
}

// ToDynamic converts to the non-generic Dynamic for backward compatibility.
func (it *TypedSimpleRequest[T]) ToDynamic() Dynamic {
	if it == nil {
		return InvalidDynamic()
	}

	return NewDynamic(it.innerData, it.isValid)
}
