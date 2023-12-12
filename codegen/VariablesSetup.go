package codegen

import (
	"reflect"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
)

type variablesSetup struct {
	inArgsOrder             *corestr.SimpleSlice
	outArgsOrder            *corestr.SimpleSlice
	reassignedVariableNames *corestr.SimpleSlice
	setupLines              *corestr.SimpleSlice
	inArgs                  args.Map // left arg name, right type
	inArgsTypes             []reflect.Type
	funcWrap                *args.FuncWrap
}
