package codegen

import "gitlab.com/auk-go/core/coretests/args"

type GenerateFunc struct {
	Func                    interface{}
	ArrangeInputs           interface{}
	Repo                    string
	GeneratePath            string
	IsGenerateSeparateCases bool
	IsOverwrite             bool
}

func (it GenerateFunc) Generate() error {
	toWrap := args.
		NewFuncWrap.
		Default(it.Func)

}
