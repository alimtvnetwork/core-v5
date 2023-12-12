package codegen

import (
	"reflect"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
)

type variablesSetup struct {
	reassignedVariableNames *corestr.SimpleSlice
	setupLines              *corestr.SimpleSlice
	inArgs                  *corestr.KeyValueCollection // left arg name, right type
	inArgsTypes             []reflect.Type
	funcWrap                *args.FuncWrap
}
