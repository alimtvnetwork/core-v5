package main

import "gitlab.com/auk-go/core/codegen"

type unitTestGenerator struct{}

func (it unitTestGenerator) Generate() error {
	generateFunc := codegen.GenerateFunc{
		Func:                    nil,
		GenerateType:            0,
		FmtType:                 0,
		TestCases:               nil,
		Behaviours:              nil,
		Repo:                    "",
		GeneratePath:            "",
		OverridingTestPkgName:   "",
		IsGenerateSeparateCases: false,
		IsIncludeFunction:       false,
		IsOverwrite:             false,
	}
}
