package coreapi

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/coredynamic"
)

type GenericResponse struct {
	Attribute *ResponseAttribute `json:"Attribute,omitempty"`
	Response  interface{}        `json:"Response,omitempty"`
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

func (receiver *GenericResponse) GenericResponseResult() *GenericResponseResult {
	return &GenericResponseResult{
		Attribute: receiver.Attribute,
		Response: coredynamic.NewSimpleResult(
			receiver,
			receiver.Attribute.IsValid,
			receiver.Attribute.Message),
	}
}

// Clone Cannot copy interface, just putting response in response field.
func (receiver *GenericResponse) Clone() *GenericResponse {
	if receiver == nil {
		return nil
	}

	return &GenericResponse{
		Attribute: receiver.Attribute.Clone(),
		Response:  receiver.Response,
	}
}
