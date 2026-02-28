package coreapi

import "gitlab.com/auk-go/core/constants"

// TypedResponseResult is the generic version of GenericResponseResult.
//
// T represents the strongly-typed response data, replacing the dynamic SimpleResult wrapper.
//
// Usage:
//
//	result := coreapi.NewTypedResponseResult[MyOutput](attr, output)
//	result.Response.Field // strongly typed
type TypedResponseResult[T any] struct {
	Attribute *ResponseAttribute `json:"Attribute,omitempty"`
	Response  T                  `json:"Response,omitempty"`
}

// NewTypedResponseResult creates a valid TypedResponseResult.
func NewTypedResponseResult[T any](
	attribute *ResponseAttribute,
	response T,
) *TypedResponseResult[T] {
	return &TypedResponseResult[T]{
		Attribute: attribute,
		Response:  response,
	}
}

// InvalidTypedResponseResult creates an invalid TypedResponseResult with a zero-value response.
func InvalidTypedResponseResult[T any](
	attribute *ResponseAttribute,
) *TypedResponseResult[T] {
	if attribute == nil {
		attribute = InvalidResponseAttribute(constants.EmptyString)
	}

	return &TypedResponseResult[T]{
		Attribute: attribute,
	}
}

// Clone returns a deep copy of the TypedResponseResult.
func (it TypedResponseResult[T]) Clone() TypedResponseResult[T] {
	return TypedResponseResult[T]{
		Attribute: it.Attribute.Clone(),
		Response:  it.Response,
	}
}

// ClonePtr returns a pointer to a deep copy.
func (it *TypedResponseResult[T]) ClonePtr() *TypedResponseResult[T] {
	if it == nil {
		return nil
	}

	cloned := it.Clone()

	return &cloned
}
