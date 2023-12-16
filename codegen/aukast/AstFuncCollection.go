package aukast

import "gitlab.com/auk-go/core/coredata/corestr"

type AstFuncCollection struct {
	Names  *corestr.SimpleSlice
	Map    map[string]AstFunction
	Parent *AstElem
}
