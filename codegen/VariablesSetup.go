package codegen

import (
	"reflect"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
)

type variablesSetup struct {
	inArgsNames  corestr.SimpleSlice
	outArgsNames corestr.SimpleSlice
	setupLines   corestr.SimpleSlice
	inArgsTypes  []reflect.Type
	funcWrap     *args.FuncWrap
}
