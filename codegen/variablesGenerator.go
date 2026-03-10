package codegen

import (
	"reflect"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coreindexes"
	"github.com/alimtvnetwork/core/coretests/args"
)

type variablesGenerator struct {
	baseGenerator BaseGenerator
}

func (it variablesGenerator) FuncWrap() *args.FuncWrapAny {
	return it.baseGenerator.FuncWrap()
}

func (it variablesGenerator) Generate() variablesSetup {
	funcWrap := it.FuncWrap()
	inArgsNames := funcWrap.InArgNames()
	inArgsTypes := funcWrap.GetInArgsTypes()

	return variablesSetup{
		inArgsNames:  inArgsNames,
		outArgsNames: funcWrap.OutArgNames(),
		setupLines: it.SetupLines(
			vars.inputPrefix,
			inArgsNames,
			inArgsTypes,
		),
		inArgsTypes: inArgsTypes,
		funcWrap:    funcWrap,
	}
}

func (it variablesGenerator) SetupLines(
	parentVariableName string,
	inArgNames []string,
	inArgsTypes []reflect.Type,
) corestr.SimpleSlice {
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
			parentVariableName,
			rightName,
			inArgsTypes[i].String(),
		)
	}

	return *toSlice
}
