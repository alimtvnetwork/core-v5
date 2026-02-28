package coreapi

// TypedRequest is the generic version of SimpleGenericRequest.
//
// T represents the strongly-typed request payload, providing compile-time safety
// instead of the dynamic SimpleRequest wrapper.
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

// ToSimpleGenericRequest converts to the non-generic SimpleGenericRequest for backward compatibility.
//
// Note: This wraps the typed request in a coredynamic.SimpleRequest via Dynamic.
func (it *TypedRequest[T]) ToGenericRequestIn() *GenericRequestIn {
	if it == nil {
		return nil
	}

	return &GenericRequestIn{
		Attribute: it.Attribute,
		Request:   it.Request,
	}
}
