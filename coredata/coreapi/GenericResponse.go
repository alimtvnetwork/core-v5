package coreapi

// GenericResponse is a type alias for TypedResponse[any].
//
// Deprecated: Use TypedResponse[T] with a concrete type for compile-time safety.
// GenericResponse remains for backward compatibility.
type GenericResponse = TypedResponse[any]

// InvalidGenericResponse creates an invalid GenericResponse (TypedResponse[any])
// with a nil response.
//
// Deprecated: Use InvalidTypedResponse[T] with a concrete type instead.
func InvalidGenericResponse(attr *ResponseAttribute) *GenericResponse {
	return InvalidTypedResponse[any](attr)
}
