package coreinstruction

import "github.com/alimtvnetwork/core/coredata/corestr"

type NameRequests struct {
	Name     string               `json:"Name,omitempty"`
	Requests *corestr.SimpleSlice `json:"Requests,omitempty"`
}
