package codegen

import (
	"reflect"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coreindexes"
	"gitlab.com/auk-go/core/coretests/args"
)

type generateVariables struct {
	baseGenerator BaseGenerator
}

func (it generateVariables) FuncWrap() *args.FuncWrap {
	return it.baseGenerator.FuncWrap()
}

func (it generateVariables) Generate() variablesSetup {
	funcWrap := it.FuncWrap()
	inArgsNames := funcWrap.InArgNames()
	inArgsTypes := funcWrap.GetInArgsTypes()

	return variablesSetup{
		inArgsNames:  inArgsNames,
		outArgsNames: funcWrap.OutArgNames(),
		setupLines:   it.SetupLines(inArgsNames, inArgsTypes),
		inArgsTypes:  inArgsTypes,
		funcWrap:     funcWrap,
	}
}

func (it generateVariables) SetupLines(inArgNames []string, inArgsTypes []reflect.Type) corestr.SimpleSlice {
	if len(inArgNames) == 0 {
		return []string{}
	}

	toSlice := corestr.
		New.
		SimpleSlice.
		ByLen(inArgNames)

	for i, name := range inArgNames {
		rightName := coreindexes.NameByIndex(i)

		toSlice.AppendFmt(
			"%s := %s.%s.(%s)",
			name,
			"input",
			rightName,
			inArgsTypes[i].String(),
		)
	}

	return *toSlice
}
