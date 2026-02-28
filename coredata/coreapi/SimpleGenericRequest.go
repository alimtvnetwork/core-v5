package coreapi

import "gitlab.com/auk-go/core/coredata/coredynamic"

// SimpleGenericRequest is a type alias for TypedRequest[*coredynamic.SimpleRequest].
//
// It wraps a *coredynamic.SimpleRequest payload with a *RequestAttribute.
// As a type alias, it inherits all TypedRequest methods: Clone, ToGenericRequestIn,
// ToSimpleGenericRequest, ToTypedSimpleGenericRequest.
//
// Deprecated: Use TypedRequest[T] or TypedSimpleGenericRequest[T] with a concrete type
// for compile-time safety. SimpleGenericRequest remains for backward compatibility.
type SimpleGenericRequest = TypedRequest[*coredynamic.SimpleRequest]
