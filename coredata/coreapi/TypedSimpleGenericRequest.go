package coreapi

import "gitlab.com/auk-go/core/coredata/coredynamic"

// TypedSimpleGenericRequest is the generic version of SimpleGenericRequest.
//
// It wraps a *coredynamic.TypedSimpleRequest[T] instead of the dynamic *coredynamic.SimpleRequest,
// providing compile-time type safety for the request payload.
//
// Usage:
//
//	type UserInput struct { Name string; Email string }
//
//	req := coreapi.NewTypedSimpleGenericRequest[UserInput](
//	    attr,
//	    coredynamic.NewTypedSimpleRequestValid[UserInput](
//	        UserInput{Name: "Alice", Email: "alice@example.com"},
//	    ),
//	)
//
//	fmt.Println(req.Request.Data().Name) // "Alice" — strongly typed
type TypedSimpleGenericRequest[T any] struct {
	Attribute *RequestAttribute                   `json:"Attribute,omitempty"`
	Request   *coredynamic.TypedSimpleRequest[T] `json:"Request,omitempty"`
}

// NewTypedSimpleGenericRequest creates a valid TypedSimpleGenericRequest.
func NewTypedSimpleGenericRequest[T any](
	attribute *RequestAttribute,
	request *coredynamic.TypedSimpleRequest[T],
) *TypedSimpleGenericRequest[T] {
	return &TypedSimpleGenericRequest[T]{
		Attribute: attribute,
		Request:   request,
	}
}

// InvalidTypedSimpleGenericRequest creates an invalid TypedSimpleGenericRequest
// with a nil request.
func InvalidTypedSimpleGenericRequest[T any](
	attribute *RequestAttribute,
) *TypedSimpleGenericRequest[T] {
	if attribute == nil {
		attribute = InvalidRequestAttribute()
	}

	return &TypedSimpleGenericRequest[T]{
		Attribute: attribute,
		Request:   nil,
	}
}

// IsValid returns true if both attribute and request are valid.
func (it *TypedSimpleGenericRequest[T]) IsValid() bool {
	if it == nil || it.Request == nil {
		return false
	}

	return it.Attribute != nil &&
		it.Attribute.IsValid &&
		it.Request.IsValid()
}

// IsInvalid returns true if the request is invalid.
func (it *TypedSimpleGenericRequest[T]) IsInvalid() bool {
	return !it.IsValid()
}

// Data returns the underlying typed data from the request.
// Panics if Request is nil.
func (it *TypedSimpleGenericRequest[T]) Data() T {
	return it.Request.Data()
}

// Message returns the request message (typically validation/error message).
func (it *TypedSimpleGenericRequest[T]) Message() string {
	if it == nil || it.Request == nil {
		return ""
	}

	return it.Request.Message()
}

// InvalidError returns the request's error if it has one.
func (it *TypedSimpleGenericRequest[T]) InvalidError() error {
	if it == nil || it.Request == nil {
		return nil
	}

	return it.Request.InvalidError()
}

// Clone returns a deep copy of the TypedSimpleGenericRequest.
func (it *TypedSimpleGenericRequest[T]) Clone() *TypedSimpleGenericRequest[T] {
	if it == nil {
		return nil
	}

	return &TypedSimpleGenericRequest[T]{
		Attribute: it.Attribute.Clone(),
		Request:   it.Request.Clone(),
	}
}

// ToSimpleGenericRequest converts to the non-generic SimpleGenericRequest
// for backward compatibility.
func (it *TypedSimpleGenericRequest[T]) ToSimpleGenericRequest() *SimpleGenericRequest {
	if it == nil {
		return nil
	}

	var simpleReq *coredynamic.SimpleRequest
	if it.Request != nil {
		simpleReq = it.Request.ToSimpleRequest()
	}

	return &SimpleGenericRequest{
		Attribute: it.Attribute,
		Request:   simpleReq,
	}
}

// ToGenericRequestIn converts to GenericRequestIn for backward compatibility.
func (it *TypedSimpleGenericRequest[T]) ToGenericRequestIn() *GenericRequestIn {
	if it == nil {
		return nil
	}

	var requestData any
	if it.Request != nil {
		requestData = it.Request.Data()
	}

	return &GenericRequestIn{
		Attribute: it.Attribute,
		Request:   requestData,
	}
}
