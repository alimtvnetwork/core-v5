package coreapi

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/coredynamic"
)

type GenericResponse struct {
	Attribute *ResponseAttribute `json:"Attribute,omitempty"`
	Response  any                `json:"Response,omitempty"`
}

func InvalidGenericResponse(attr *ResponseAttribute) *GenericResponse {
	if attr == nil {
		return &GenericResponse{
			Attribute: InvalidResponseAttribute(constants.EmptyString),
			Response:  nil,
		}
	}

	return &GenericResponse{
		Attribute: attr,
		Response:  nil,
	}
}

func (it *GenericResponse) GenericResponseResult() *GenericResponseResult {
	return &GenericResponseResult{
		Attribute: it.Attribute,
		Response: coredynamic.NewSimpleResult(
			it,
			it.Attribute.IsValid,
			it.Attribute.Message),
	}
}

// Clone returns a deep copy. Response is copied by reference since any cannot be deep-cloned.
func (it *GenericResponse) Clone() *GenericResponse {
	if it == nil {
		return nil
	}

	return &GenericResponse{
		Attribute: it.Attribute.Clone(),
		Response:  it.Response,
	}
}
