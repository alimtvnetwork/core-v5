package coreapi

import "gitlab.com/auk-go/core/coredata/coredynamic"

type GenericRequestIn struct {
	Attribute *RequestAttribute `json:"Attribute,omitempty"`
	Request   any               `json:"Request,omitempty"`
}

func InvalidGenericRequestIn(
	attr *RequestAttribute,
) *GenericRequestIn {
	return &GenericRequestIn{
		Attribute: attr,
	}
}

func (it *GenericRequestIn) SimpleGenericRequest(
	isValid bool,
	invalidMessage string,
) *SimpleGenericRequest {
	return &SimpleGenericRequest{
		Attribute: it.Attribute,
		Request: coredynamic.NewSimpleRequest(
			it.Request,
			isValid,
			invalidMessage),
	}
}

// Clone returns a deep copy of GenericRequestIn.
//
// Note: The Request field is copied by reference since any cannot be deep-cloned generically.
func (it *GenericRequestIn) Clone() *GenericRequestIn {
	if it == nil {
		return nil
	}

	return &GenericRequestIn{
		Attribute: it.Attribute.Clone(),
		Request:   it.Request,
	}
}
