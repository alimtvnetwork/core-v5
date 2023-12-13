package codegen

import "gitlab.com/auk-go/core/coretests/args"

type generateVariables struct {
	baseGenerator BaseGenerator
}

func (it generateVariables) FuncWrap() *args.FuncWrap {
	return it.baseGenerator.FuncWrap()
}

func (it generateVariables) Generate() variablesSetup {
	// funcWrap := it.FuncWrap()

	return variablesSetup{}
}
