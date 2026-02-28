package coreapi

import "gitlab.com/auk-go/core/constants"

// TypedResponse is the generic version of GenericResponse.
//
// T represents the strongly-typed response payload, replacing the dynamic interface{} field.
//
// Usage:
//
//	resp := coreapi.NewTypedResponse[MyResult](attr, result)
//	resp.Response.SomeField // fully typed
type TypedResponse[T any] struct {
	Attribute *ResponseAttribute `json:"Attribute,omitempty"`
	Response  T                  `json:"Response,omitempty"`
}

// NewTypedResponse creates a valid TypedResponse with the given attribute and response.
func NewTypedResponse[T any](
	attribute *ResponseAttribute,
	response T,
) *TypedResponse[T] {
	return &TypedResponse[T]{
		Attribute: attribute,
		Response:  response,
	}
}

// InvalidTypedResponse creates an invalid TypedResponse with a zero-value response.
func InvalidTypedResponse[T any](
	attribute *ResponseAttribute,
) *TypedResponse[T] {
	if attribute == nil {
		attribute = InvalidResponseAttribute(constants.EmptyString)
	}

	return &TypedResponse[T]{
		Attribute: attribute,
	}
}

// Clone returns a deep copy of the TypedResponse.
func (it *TypedResponse[T]) Clone() *TypedResponse[T] {
	if it == nil {
		return nil
	}

	return &TypedResponse[T]{
		Attribute: it.Attribute.Clone(),
		Response:  it.Response,
	}
}

// ToGenericResponse converts to the non-generic GenericResponse for backward compatibility.
func (it *TypedResponse[T]) ToGenericResponse() *GenericResponse {
	if it == nil {
		return nil
	}

	return &GenericResponse{
		Attribute: it.Attribute,
		Response:  it.Response,
	}
}
