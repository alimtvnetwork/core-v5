package coreapi

// GenericRequestIn is a type alias for TypedRequestIn[any].
//
// Deprecated: Use TypedRequestIn[T] with a concrete type for compile-time safety.
// GenericRequestIn remains for backward compatibility.
type GenericRequestIn = TypedRequestIn[any]

// InvalidGenericRequestIn creates an invalid GenericRequestIn (TypedRequestIn[any])
// with a zero-value request.
//
// Deprecated: Use InvalidTypedRequestIn[T] with a concrete type instead.
func InvalidGenericRequestIn(
	attr *RequestAttribute,
) *GenericRequestIn {
	return InvalidTypedRequestIn[any](attr)
}
