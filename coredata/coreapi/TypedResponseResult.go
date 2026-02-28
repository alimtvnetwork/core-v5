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

// IsValid returns true if the attribute is present and valid.
func (it *TypedResponseResult[T]) IsValid() bool {
	return it != nil &&
		it.Attribute != nil &&
		it.Attribute.IsValid
}

// IsInvalid returns true if the result is invalid.
func (it *TypedResponseResult[T]) IsInvalid() bool {
	return !it.IsValid()
}

// Message returns the attribute message, or empty string if nil.
func (it *TypedResponseResult[T]) Message() string {
	if it == nil || it.Attribute == nil {
		return constants.EmptyString
	}

	return it.Attribute.Message
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

// ToGenericResponseResult converts to the legacy GenericResponseResult.
func (it *TypedResponseResult[T]) ToGenericResponseResult() *GenericResponseResult {
	if it == nil {
		return nil
	}

	return InvalidGenericResponseResult(it.Attribute)
}

// ToGenericResponse converts to GenericResponse (TypedResponse[any]) for backward compatibility.
//
// Deprecated: GenericResponse is now a type alias for TypedResponse[any].
func (it *TypedResponseResult[T]) ToGenericResponse() *GenericResponse {
	if it == nil {
		return nil
	}

	return &GenericResponse{
		Attribute: it.Attribute,
		Response:  it.Response,
	}
}

// ToTypedResponse converts to TypedResponse[T].
func (it *TypedResponseResult[T]) ToTypedResponse() *TypedResponse[T] {
	if it == nil {
		return nil
	}

	return &TypedResponse[T]{
		Attribute: it.Attribute,
		Response:  it.Response,
	}
}
