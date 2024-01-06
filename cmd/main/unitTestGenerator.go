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

	generateFunc := codegen.GenerateFunc{
		Func:         curFunc,
		GenerateType: codegentype.MultipleArranges,
		FmtType:      fmtcodegentype.WithExpect,
		TestCases: []coretestcases.CaseV1{
			{
				Title: "Some",
				ArrangeInput: []args.One{
					{
						First: "x",
					},
					// {
					// 	First: "in 1 - first - ",
					// 	Second: args.LeftRight{
					// 		Left:   "l",
					// 		Right:  "r",
					// 		Expect: "e",
					// 	},
					// 	Third:  nil,
					// 	Fourth: nil,
					// 	Fifth:  nil,
					// 	Expect: "some expect",
					// },
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
	alim2 AlimStruct,
) (r1 string, r2 int, r3 **AlimStruct) {
	toAlim := &AlimStruct{
		First:     "someName - " + alim.First + alim2.First,
		LeftRight: alim.LeftRight,
		Draft:     alim2.Draft,
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
