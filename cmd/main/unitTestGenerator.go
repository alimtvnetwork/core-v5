package main

import (
	"fmt"
	"go/ast"

	"gitlab.com/auk-go/core/codegen"
	"gitlab.com/auk-go/core/codegen/aukast"
	"gitlab.com/auk-go/core/codegen/codegentype"
	"gitlab.com/auk-go/core/codegen/coreproperty"
	"gitlab.com/auk-go/core/codegen/fmtcodegentype"
	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
)

type unitTestGenerator struct{}

func (it unitTestGenerator) Generate() {
	curFunc := coreproperty.Writer.Write

	/**
	func (it unitTestGenerator) samplefunc(
		x int,
		arg1, arg2 string,
		alim *AlimStruct,
		alim2 []AlimStruct,
	) (r1 string, r2 int, r3 **AlimStruct) {

	{
						First:  1,
						Second: "alim 1",
						Third:  "alim 2",
						Fourth: &samplefunc.AlimStruct{
							First: "alim 1",
							LeftRight: args.LeftRight{
								Left:   "l",
								Right:  "r",
								Expect: "e",
							},
							Draft: &coretests.DraftType{
								SampleString1: "alim 1",
							},
						},
						Fifth: &[]*samplefunc.AlimStruct{
							{
								First: "alim 2",
								LeftRight: args.LeftRight{
									Left:   "a2-l",
									Right:  "a2-r",
									Expect: "a2-e",
								},
								Draft: &coretests.DraftType{
									SampleString1: "alim 2",
								},
							},
						},
						Expect: "some expect",
					},
				},
	*/
	generateFunc := codegen.GenerateFunc{
		Func:         curFunc,
		Struct:       nil,
		GenerateType: codegentype.MultipleArranges,
		FmtType:      fmtcodegentype.WithExpect,
		TestCases: []coretestcases.CaseV1{
			{
				Title: "When nil / null is given, nil returned as is hardcoded.",
				ArrangeInput: []args.One{
					{
						First:  nil,
						Expect: nil,
					},
				},
				CustomFormat:    "",
				HasError:        false,
				HasPanic:        false,
				IsValidateError: false,
			},
			{
				Title: "Some string given outputs double quoted string.",
				ArrangeInput: []args.One{
					{
						First:  "some string",
						Expect: "some string",
					},
				},
				CustomFormat:    "",
				HasError:        false,
				HasPanic:        false,
				IsValidateError: false,
			},
			{
				Title: "Slice of string ,int, bytes, boolean outputs in similar fashion",
				ArrangeInput: []args.One{
					{
						First: []string{
							"some val 1",
							"some val 2",
							"some val 3",
						},
					},
					{
						First: []int{
							-1,
							5,
							10,
							255,
							1500,
						},
					},
					{
						First: []byte{
							0,
							5,
							10,
							255,
							100,
						},
					},
					{
						First: []bool{
							true,
							false,
							true,
						},
					},
				},
				CustomFormat:    "",
				HasError:        false,
				HasPanic:        false,
				IsValidateError: false,
			},
			{
				Title: "Struct with slice, etc outputs as it was given.",
				ArrangeInput: []args.One{
					{
						First: &coretests.DraftType{
							SampleString1: "sample 1",
							SampleString2: "sample 2",
							SampleInteger: -59,
							Lines: []string{
								"hello 1",
								"hello 2",
							},
							RawBytes: []byte("really !"),
						},
					},
				},
				CustomFormat:    "",
				HasError:        false,
				HasPanic:        false,
				IsValidateError: false,
			},
			{
				Title: "Pointer struct with slice, etc outputs as it was given.",
				ArrangeInput: []args.One{
					{
						First: &coretests.DraftType{
							SampleString1: "sample 1 Ptr",
							SampleString2: "sample 2 Ptr",
							SampleInteger: -59,
							Lines: []string{
								"hello 1 Ptr",
								"hello 2 Ptr",
							},
							RawBytes: []byte("really ! Ptr"),
						},
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
		OverridingNames: codegen.OverridingNames{
			TestPkgName: "",
			FuncCall:    "coreproperty.Writer.Write",
		},
		UnitTestRootPath: codestack.Dir.RepoDirJoin("tests/integratedtests/codegentests"),
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
