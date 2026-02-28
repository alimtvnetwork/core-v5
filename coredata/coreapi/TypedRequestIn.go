package coreapi

import "gitlab.com/auk-go/core/coredata/coredynamic"

// TypedRequestIn is the generic version of GenericRequestIn.
//
// T represents the strongly-typed request payload, replacing the dynamic interface{} field.
//
// Usage:
//
//	req := coreapi.NewTypedRequestIn[MyPayload](attr, payload)
//	req.Request.SomeField // fully typed, no assertion needed
type TypedRequestIn[T any] struct {
	Attribute *RequestAttribute `json:"Attribute,omitempty"`
	Request   T                 `json:"Request,omitempty"`
}

// NewTypedRequestIn creates a valid TypedRequestIn with the given attribute and request.
func NewTypedRequestIn[T any](
	attribute *RequestAttribute,
	request T,
) *TypedRequestIn[T] {
	return &TypedRequestIn[T]{
		Attribute: attribute,
		Request:   request,
	}
}

// InvalidTypedRequestIn creates an invalid TypedRequestIn with a zero-value request.
func InvalidTypedRequestIn[T any](
	attribute *RequestAttribute,
) *TypedRequestIn[T] {
	if attribute == nil {
		attribute = InvalidRequestAttribute()
	}

	return &TypedRequestIn[T]{
		Attribute: attribute,
	}
}

// Clone returns a deep copy of the TypedRequestIn.
func (it *TypedRequestIn[T]) Clone() *TypedRequestIn[T] {
	if it == nil {
		return nil
	}

	return &TypedRequestIn[T]{
		Attribute: it.Attribute.Clone(),
		Request:   it.Request,
	}
}

// ToGenericRequestIn converts to the non-generic GenericRequestIn for backward compatibility.
func (it *TypedRequestIn[T]) ToGenericRequestIn() *GenericRequestIn {
	if it == nil {
		return nil
	}

	return &GenericRequestIn{
		Attribute: it.Attribute,
		Request:   it.Request,
	}
}

// TypedSimpleGenericRequest converts to a TypedSimpleGenericRequest[T]
// by wrapping the request in a TypedSimpleRequest.
//
// This mirrors GenericRequestIn.SimpleGenericRequest().
func (it *TypedRequestIn[T]) TypedSimpleGenericRequest(
	isValid bool,
	invalidMessage string,
) *TypedSimpleGenericRequest[T] {
	if it == nil {
		return nil
	}

	return &TypedSimpleGenericRequest[T]{
		Attribute: it.Attribute,
		Request: coredynamic.NewTypedSimpleRequest[T](
			it.Request,
			isValid,
			invalidMessage),
	}
}

// SimpleGenericRequest converts to the legacy SimpleGenericRequest.
//
// This mirrors GenericRequestIn.SimpleGenericRequest() for backward compatibility.
func (it *TypedRequestIn[T]) SimpleGenericRequest(
	isValid bool,
	invalidMessage string,
) *SimpleGenericRequest {
	if it == nil {
		return nil
	}

	return &SimpleGenericRequest{
		Attribute: it.Attribute,
		Request: coredynamic.NewSimpleRequest(
			it.Request,
			isValid,
			invalidMessage),
	}
}
