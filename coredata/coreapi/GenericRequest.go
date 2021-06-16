package coreapi

import "gitlab.com/evatix-go/core/coredata/coredynamic"

type GenericRequestIn struct {
	Attribute *RequestAttribute `json:"Attribute,omitempty"`
	Request   interface{}       `json:"Request,omitempty"`
}

func InvalidGenericRequestIn(attr *RequestAttribute) *GenericRequestIn {
	return &GenericRequestIn{
		Attribute: attr,
		Request:   nil,
	}
}

func (receiver *GenericRequestIn) SimpleGenericRequest(
	isValid bool,
	invalidMessage string,
) *SimpleGenericRequest {
	return &SimpleGenericRequest{
		Attribute: receiver.Attribute,
		Request:   coredynamic.NewSimpleRequest(receiver.Request, isValid, invalidMessage),
	}
}

func (receiver *GenericRequestIn) Clone() *GenericRequestIn {
	if receiver == nil {
		return nil
	}

	return &GenericRequestIn{
		Attribute: receiver.Attribute.Clone(),
		Request:   receiver.Clone(),
	}
}
