package coreapi

import "gitlab.com/auk-go/core/coredata/coredynamic"

// GenericResponseResult is a type alias for TypedResponseResult[*coredynamic.SimpleResult].
//
// It wraps a *coredynamic.SimpleResult response with a *ResponseAttribute.
// As a type alias, it inherits all TypedResponseResult methods: Clone, ClonePtr,
// IsValid, IsInvalid, Message, ToGenericResponseResult, ToGenericResponse, ToTypedResponse.
//
// Deprecated: Use TypedResponseResult[T] with a concrete type for compile-time safety.
// GenericResponseResult remains for backward compatibility.
type GenericResponseResult = TypedResponseResult[*coredynamic.SimpleResult]
