package codegen

import "gitlab.com/auk-go/core/coredata/corestr"

type GeneratedCode struct {
	Code     string
	Packages corestr.SimpleSlice
}
