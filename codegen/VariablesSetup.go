package codegen

import (
	"reflect"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

type variablesSetup struct {
	inArgsNames  corestr.SimpleSlice
	outArgsNames corestr.SimpleSlice
	setupLines   corestr.SimpleSlice
	inArgsTypes  []reflect.Type
	funcWrap     *args.FuncWrapAny
}

func (it variablesSetup) CompiledSetupLine() string {
	if it.setupLines.Length() == 0 {
		return ""
	}

	return it.setupLines.Join("\n\t\t")
}
