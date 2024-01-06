package main

import (
	"fmt"
	"go/ast"

	"gitlab.com/auk-go/core/codegen"
	"gitlab.com/auk-go/core/codegen/aukast"
	"gitlab.com/auk-go/core/codegen/codegentype"
	"gitlab.com/auk-go/core/codegen/fmtcodegentype"
	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
)

type unitTestGenerator struct{}

type AlimStruct struct {
	First     string
	LeftRight args.LeftRight
	Draft     coretests.DraftType
}

func (it unitTestGenerator) Generate() {
	curFunc := it.SampleFunc

	/**
	func (it unitTestGenerator) SampleFunc(
		x int,
		arg1, arg2 string,
		alim *AlimStruct,
		alim2 []AlimStruct,
	) (r1 string, r2 int, r3 **AlimStruct) {
	*/
	generateFunc := codegen.GenerateFunc{
		Func:         curFunc,
		GenerateType: codegentype.MultipleArranges,
		FmtType:      fmtcodegentype.WithExpect,
		TestCases: []coretestcases.CaseV1{
			{
				Title: "Some",
				ArrangeInput: []args.Five{
					// {
					// 	First: "x",
					// },
					{
						First:  1,
						Second: "alim 1",
						Third:  "alim 2",
						Fourth: &AlimStruct{
							First: "alim 1",
							LeftRight: args.LeftRight{
								Left:   "l",
								Right:  "r",
								Expect: "e",
							},
							Draft: coretests.DraftType{
								SampleString1: "alim 1",
							},
						},
						Fifth: []AlimStruct{
							{
								First: "alim 2",
								LeftRight: args.LeftRight{
									Left:   "a2-l",
									Right:  "a2-r",
									Expect: "a2-e",
								},
								Draft: coretests.DraftType{
									SampleString1: "alim 2",
								},
							},
						},
						Expect: "some expect",
					},
				},
				CustomFormat:    "",
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

func (it unitTestGenerator) SampleFunc(
	x int,
	arg1, arg2 string,
	alim *AlimStruct,
	alim2 []AlimStruct,
) (r1 string, r2 int, r3 **AlimStruct) {
	toAlim := &AlimStruct{
		First:     "someName - " + alim.First + alim2[0].First,
		LeftRight: alim.LeftRight,
		Draft:     alim2[0].Draft,
	}

	return arg1 + " " + arg2 + "-> Processed", x + 1, &toAlim
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
