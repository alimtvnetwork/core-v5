package main

import (
	"gitlab.com/auk-go/core/codegen"
	"gitlab.com/auk-go/core/codegen/codegentype"
	"gitlab.com/auk-go/core/codegen/fmtcodegentype"
	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type unitTestGenerator struct{}

func (it unitTestGenerator) Generate() {
	generateFunc := codegen.GenerateFunc{
		Func:         reflectinternal.GetFunc.PublicFuncName,
		GenerateType: codegentype.Simple,
		FmtType:      fmtcodegentype.Default,
		TestCases: []coretestcases.CaseV1{
			{
				Title:           "Some",
				ArrangeInput:    []args.One{},
				ActualInput:     nil,
				ExpectedInput:   nil,
				Additional:      nil,
				CustomFormat:    "",
				VerifyTypeOf:    nil,
				Parameters:      nil,
				IsEnable:        0,
				HasError:        false,
				HasPanic:        false,
				IsValidateError: false,
			},
		},
		Behaviours:              nil,
		Repo:                    codestack.Dir.RepoDir(),
		GeneratePath:            codestack.Dir.CurDirJoin("unit-test"),
		IsGenerateSeparateCases: false,
		IsIncludeFunction:       false,
		IsOverwrite:             true,
	}

	err := generateFunc.Generate()

	errcore.HandleErr(err)
}
