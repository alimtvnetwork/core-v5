package coreapi

import "gitlab.com/auk-go/core/constants"

// InvalidGenericResponseResult creates an invalid GenericResponseResult
// (TypedResponseResult[*coredynamic.SimpleResult]) with a nil response.
//
// Deprecated: Use InvalidTypedResponseResult[T] with a concrete type instead.
func InvalidGenericResponseResult(attribute *ResponseAttribute) *GenericResponseResult {
	if attribute == nil {
		return &GenericResponseResult{
			Attribute: InvalidResponseAttribute(constants.EmptyString),
			Response:  nil,
		}
	}

	return &GenericResponseResult{
		Attribute: attribute,
		Response:  nil,
	}
}
