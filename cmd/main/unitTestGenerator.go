package main

import (
	"fmt"
	"go/ast"

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
		GenerateType: codegentype.MultipleArranges,
		FmtType:      fmtcodegentype.Default,
		TestCases: []coretestcases.CaseV1{
			{
				Title: "Some",
				ArrangeInput: []args.One{
					{
						First:  "someName",
						Expect: "some expect",
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

func (it unitTestGenerator) DummyFunc(someVar1, someVar2 *[]string, x *unitTestGenerator) (
	r1, r2 string, ix *unitTestGenerator,
) {
	return "", "", nil
}

func DummyFuncX(someVar1, someVar2 *[]string, x *unitTestGenerator) (
	r1, r2 string, ix *unitTestGenerator,
) {
	return "", "", nil
}

func (it unitTestGenerator) AstChecker() {
	astReader, err := aukast.New.AstReader.FilePath(it.curFile())

	errcore.HandleErr(err)

	// nodesMap, err := astReader.NodesMap()
	// errcore.HandleErr(err)
	//
	// structTypes, err := astReader.StructTypes()
	// errcore.HandleErr(err)
	//
	// firstNode, _ := astReader.SubstringByNode(structTypes[0])

	fmt.Println()
	fmt.Println()
	functions, err := astReader.Functions()
	errcore.HandleErr(err)

	fmt.Println(functions)

	c := astReader.Filter(
		func(elem *aukast.AstElem) (isTake bool) {
			switch elem.Node.(type) {
			case *ast.CompositeLit, *ast.KeyValueExpr:
				return true
			}

			return false
		},
	)

	fmt.Println(c)
	// fmt.Println(structTypes)
	// fmt.Println(firstNode)
}
