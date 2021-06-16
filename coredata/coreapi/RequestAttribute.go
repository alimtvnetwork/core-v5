package coreapi

import (
	"gitlab.com/evatix-go/core/reqtype"
)

type RequestAttribute struct {
	Url           string `json:"Url,omitempty"`
	Host          string `json:"Host,omitempty"`
	ResourceName  string `json:"ResourceName,omitempty"`
	RequestType   reqtype.Request
	IsValid       bool
	SearchRequest *SearchRequest `json:"SearchRequest,omitempty"`
	PageRequest   *PageRequest   `json:"PageRequest,omitempty"`
}

func (receiver *RequestAttribute) Clone() *RequestAttribute {
	if receiver == nil {
		return nil
	}

	return &RequestAttribute{
		Url:           receiver.Url,
		Host:          receiver.Host,
		ResourceName:  receiver.ResourceName,
		RequestType:   receiver.RequestType,
		IsValid:       receiver.IsValid,
		SearchRequest: receiver.SearchRequest.Clone(),
		PageRequest:   receiver.PageRequest.Clone(),
	}
}
