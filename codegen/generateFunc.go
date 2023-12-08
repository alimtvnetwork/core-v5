package codegen

import (
	"errors"
	"reflect"

	"gitlab.com/auk-go/core/codegen/codegentype"
	"gitlab.com/auk-go/core/codegen/fmtcodegentype"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

type GenerateFunc struct {
	Func                    interface{}
	GenerateType            codegentype.Variant
	FmtType                 fmtcodegentype.Variant
	TestCases               []coretestcases.CaseV1
	Behaviours              corestr.SimpleSlice
	Repo                    string
	GeneratePath            string
	OverridingTestPkgName   string
	IsGenerateSeparateCases bool
	IsIncludeFunction       bool
	IsOverwrite             bool
}

func (it GenerateFunc) Generate() error {
	toWrap := it.toFunWrap()

	pkgName := it.testPkgName(toWrap)
	newPackagesLines := it.allPackages(toWrap)
	firstArrangeTypeName := it.firstArrangeTypeName()

	actLines := it.generateActLines()

	_ := map[string]string{
		"$packageName":   pkgName,
		"$fmtJoin":       it.generateFmtJoin(),
		"$newPackages":   newPackagesLines,
		"$ArrangeType":   firstArrangeTypeName,
		"$linesPossible": "100",
	}

	return nil
}

func (it GenerateFunc) firstArrangeTypeName() string {
	rt := it.firstArrangeType()

	if rt == nil {
		return constants.NilAngelBracket
	}

	return (*rt).String()
}

func (it GenerateFunc) allPackages(toWrap *args.FuncWrap) string {
	arrangePkgPaths := it.arrangePackages()

	newPackages := corestr.
		New.
		SimpleSlice.
		SpreadStrings(
			arrangePkgPaths...,
		).
		Add(toWrap.PkgPath()).
		WrapDoubleQuote()

	newPackagesLines := newPackages.JoinLine()

	return newPackagesLines
}

func (it GenerateFunc) firstArrangeType() *reflect.Type {
	if len(it.TestCases) == 0 {
		return nil
	}

	rt := reflect.TypeOf(
		it.TestCases[0].ArrangeInput,
	)

	return &rt
}

func (it GenerateFunc) arrangeReflectTypes() []reflect.Type {
	var results []reflect.Type

	for _, testCase := range it.TestCases {
		results = append(
			results,
			reflect.TypeOf(testCase.ArrangeInput),
		)
	}

	return results
}

func (it GenerateFunc) arrangePackages() []string {
	allReflectTypes := it.arrangeReflectTypes()

	var pks []string
	for _, reflectType := range allReflectTypes {
		pks = append(pks, reflectType.PkgPath())
	}

	return pks
}

func (it GenerateFunc) testPkgName(toWrap *args.FuncWrap) string {
	return toWrap.PkgName() + "tests"
}

func (it GenerateFunc) toFunWrap() *args.FuncWrap {
	return args.
		NewFuncWrap.
		Default(it.Func)
}

func (it GenerateFunc) generateActLines() []string {

}

func (it GenerateFunc) generateFmtJoin() string {
	return it.FmtType.Fmt()
}

func (it GenerateFunc) generateFmtOutputs(
	joiner string,
	funcName string,
	expected string,
	outArs, inArgs *corestr.SimpleSlice,
) (*corestr.SimpleSlice, error) {
	slice := corestr.New.SimpleSlice.Cap(20)
	slice.Add("caseIndex")

	switch it.FmtType {
	case fmtcodegentype.Default: // "%d : %s -> %s",
		outArgsString := outArs.Join(joiner)
		inArgsString := inArgs.Join(joiner)
		slice.Add(inArgsString)
		slice.Add(outArgsString)

		return slice, nil
	case fmtcodegentype.WithFunction: // "%d : %s(%s) -> %s | %s",
		outArgsString := outArs.Join(joiner)
		inArgsString := inArgs.Join(joiner)
		slice.Add(funcName)
		slice.Add(inArgsString)
		slice.Add(outArgsString)
		slice.Add(expected)

		return slice, nil
	}

	return slice, it.FmtType.OnlySupportedMsgErr()
}

func (it GenerateFunc) getFuncName() (string, error) {
	funcWrap := it.toFunWrap()

	if funcWrap.IsInvalid() {
		return "", errors.New("func wrap is invalid - func name")
	}

	return funcWrap.GetFuncName(), nil
}

// outArgs
//
//	Aka returns Args
//
// - if one then return "result" only
// - Or else, result1, result2 ...
func (it GenerateFunc) outArgs() (*corestr.SimpleSlice, error) {
	funcWrap := it.toFunWrap()

	if funcWrap.IsInvalid() {
		return it.emptySlice(), errors.New("func wrap is invalid - return args")
	}

	length := funcWrap.ReturnLength()
	slice := corestr.New.SimpleSlice.Cap(length)

	if length == 1 {
		return slice.Add("result"), nil
	}

	for i := 0; i < length; i++ {
		slice.AppendFmt("result%d", i+1)
	}

	return slice, nil
}

// inArgs
//
// - if one then return "result" only
// - Or else, result1, result2 ...
func (it GenerateFunc) inArgs() (*corestr.SimpleSlice, error) {
	funcWrap := it.toFunWrap()

	if funcWrap.IsInvalid() {
		return it.emptySlice(), errors.New("func wrap is invalid - return args")
	}

	length := funcWrap.ArgsCount()
	slice := corestr.New.SimpleSlice.Cap(length)

	if length == 1 {
		return slice.Add("result"), nil
	}

	for i := 0; i < length; i++ {
		slice.AppendFmt("result%d", i+1)
	}

	return slice, nil
}

func (it GenerateFunc) emptySlice() *corestr.SimpleSlice {
	return corestr.Empty.SimpleSlice()
}
