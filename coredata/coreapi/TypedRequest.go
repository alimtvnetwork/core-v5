package coreapi

import "gitlab.com/auk-go/core/coredata/coredynamic"

// TypedRequest is the generic request type.
//
// T represents the strongly-typed request payload, providing compile-time safety.
//
// Usage:
//
//	req := coreapi.NewTypedRequest[MyInput](attr, input)
//	req.Request.Field // strongly typed access
type TypedRequest[T any] struct {
	Attribute *RequestAttribute `json:"Attribute,omitempty"`
	Request   T                 `json:"Request,omitempty"`
}

// NewTypedRequest creates a valid TypedRequest with the given attribute and request.
func NewTypedRequest[T any](
	attribute *RequestAttribute,
	request T,
) *TypedRequest[T] {
	return &TypedRequest[T]{
		Attribute: attribute,
		Request:   request,
	}
}

// InvalidTypedRequest creates an invalid TypedRequest with a zero-value request.
func InvalidTypedRequest[T any](
	attribute *RequestAttribute,
) *TypedRequest[T] {
	if attribute == nil {
		attribute = InvalidRequestAttribute()
	}

	return &TypedRequest[T]{
		Attribute: attribute,
	}
}

// Clone returns a deep copy of the TypedRequest.
func (it *TypedRequest[T]) Clone() *TypedRequest[T] {
	if it == nil {
		return nil
	}

	return &TypedRequest[T]{
		Attribute: it.Attribute.Clone(),
		Request:   it.Request,
	}
}

// ToTypedSimpleGenericRequest converts to TypedSimpleGenericRequest[T]
// by wrapping the typed request in a TypedSimpleRequest[T].
func (it *TypedRequest[T]) ToTypedSimpleGenericRequest(
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
