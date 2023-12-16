package main

import (
	"fmt"

	"gitlab.com/auk-go/core/codegen"
	"gitlab.com/auk-go/core/codegen/aukast"
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
	curFunc := reflectinternal.GetFunc.PascalFuncName

	generateFunc := codegen.GenerateFunc{
		Func:         curFunc,
		GenerateType: codegentype.Simple,
		FmtType:      fmtcodegentype.Default,
		TestCases: []coretestcases.CaseV1{
			{
				Title: "Some",
				ArrangeInput: []args.One{
					{
						First:  "someName",
						Expect: nil,
					},
					{
						First:  "someName 2",
						Expect: nil,
					},
				},
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
		Behaviours: []string{
			"Verification",
		},
		UnitTestRootPath: codestack.Dir.CurDirJoin("unit-test"),
		Options: codegen.Options{
			IsGenerateInSameFile:  true,
			IsWriteTestCasesFirst: true,
			IsIncludeFunction:     false,
			IsOverwrite:           true,
		},
	}

	err := generateFunc.Generate()

	errcore.HandleErr(err)

	fmt.Println(generateFunc.SuccessMessage())
}

func (it unitTestGenerator) curFile() string {
	return codestack.File.CurrentFilePath()
}

func (it unitTestGenerator) AstChecker() {
	astReader := aukast.New.FilePath(it.curFile())

	node, err := astReader.Initialize()

	errcore.HandleErr(err)

	nodesMap, err := astReader.NodesMap()
	errcore.HandleErr(err)

	structTypes, err := astReader.StructTypes()
	errcore.HandleErr(err)

	firstNode, _ := astReader.SubstringByNode(structTypes[0])

	for key := range nodesMap {
		fmt.Printf("%s,\n", key)
	}

	fmt.Println()
	fmt.Println()

	fmt.Println(node.Decls)
	fmt.Println(nodesMap)
	fmt.Println(structTypes)
	fmt.Println(firstNode)
}
