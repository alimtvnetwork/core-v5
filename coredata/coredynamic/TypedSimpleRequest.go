package coredynamic

import (
	"encoding/json"
	"errors"

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

// Data returns the strongly-typed request data.
func (it *TypedSimpleRequest[T]) Data() T {
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

// JsonBytes serializes the inner data to JSON bytes.
func (it *TypedSimpleRequest[T]) JsonBytes() ([]byte, error) {
	return json.Marshal(it.innerData)
}

// JsonResult returns a corejson.Result for the inner data.
func (it *TypedSimpleRequest[T]) JsonResult() corejson.Result {
	return corejson.New(it.innerData)
}

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

// GetAsFloat64 attempts to retrieve the data as a float64.
func (it *TypedSimpleRequest[T]) GetAsFloat64() (float64, bool) {
	val, ok := any(it.innerData).(float64)

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
