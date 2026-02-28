package coreapi

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/coredynamic"
)

// TypedResponse is the generic API response type.
//
// T represents the strongly-typed response payload.
// When T is `any`, this is equivalent to the legacy GenericResponse.
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

// TypedResponseResult converts to a TypedResponseResult[T].
func (it *TypedResponse[T]) TypedResponseResult() *TypedResponseResult[T] {
	if it == nil {
		return nil
	}

	return &TypedResponseResult[T]{
		Attribute: it.Attribute,
		Response:  it.Response,
	}
}

// GenericResponseResult converts to the legacy GenericResponseResult
// by wrapping the response in a coredynamic.SimpleResult.
func (it *TypedResponse[T]) GenericResponseResult() *GenericResponseResult {
	if it == nil {
		return nil
	}

	return &GenericResponseResult{
		Attribute: it.Attribute,
		Response: coredynamic.NewSimpleResult(
			it.Response,
			it.Attribute.IsValid,
			it.Attribute.Message),
	}
}
