package coreapi

import "gitlab.com/auk-go/core/coredata/coredynamic"

// InvalidSimpleGenericRequest creates an invalid SimpleGenericRequest
// (TypedRequest[*coredynamic.SimpleRequest]) with a nil request.
//
// Deprecated: Use InvalidTypedRequest[T] with a concrete type instead.
func InvalidSimpleGenericRequest(attribute *RequestAttribute) *SimpleGenericRequest {
	if attribute == nil {
		return &SimpleGenericRequest{
			Attribute: InvalidRequestAttribute(),
			Request:   nil,
		}
	}

	return &SimpleGenericRequest{
		Attribute: attribute,
		Request:   nil,
	}
}

// NewSimpleGenericRequest creates a SimpleGenericRequest from an attribute and
// a *coredynamic.SimpleRequest.
//
// Deprecated: Use NewTypedRequest[T] or NewTypedSimpleGenericRequest[T] instead.
func NewSimpleGenericRequest(
	attribute *RequestAttribute,
	request *coredynamic.SimpleRequest,
) *SimpleGenericRequest {
	return &SimpleGenericRequest{
		Attribute: attribute,
		Request:   request,
	}
}
