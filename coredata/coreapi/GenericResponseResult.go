package coreapi

import (
	"gitlab.com/evatix-go/core/coredata/coredynamic"
)

type GenericResponseResult struct {
	Attribute *ResponseAttribute        `json:"Attribute,omitempty"`
	Response  *coredynamic.SimpleResult `json:"Response,omitempty"`
}

func (receiver *GenericResponseResult) Clone() *GenericResponseResult {
	if receiver == nil {
		return nil
	}

	return &GenericResponseResult{
		Attribute: receiver.Attribute.Clone(),
		Response:  receiver.Response.Clone(),
	}
}
