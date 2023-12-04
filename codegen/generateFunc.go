package codegen

import (
	"reflect"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
)

type GenerateFunc struct {
	Func                    interface{}
	ArrangeInputs           interface{}
	Behaviours              corestr.SimpleSlice
	Repo                    string
	GeneratePath            string
	OverridingTestPkgName   string
	IsGenerateSeparateCases bool
	IsOverwrite             bool
}

func (it GenerateFunc) Generate() error {
	toWrap := it.toFunWrap()

	pkgName := it.testPkgName(toWrap)
	arrangeRt := reflect.TypeOf(it.ArrangeInputs)
	arrangePkg := arrangeRt.PkgPath()
	newPackages := corestr.New.SimpleSlice.SpreadStrings(
		toWrap.PkgPath(),
		arrangePkg,
	).WrapDoubleQuote()

	newPackagesLines := newPackages.JoinLine()

	actLines := it.generateActLines()

	_ := map[string]string{
		"$packageName":   pkgName,
		"$newPackages":   newPackagesLines,
		"$ArrangeType":   arrangeRt.String(),
		"$linesPossible": "100",
	}

	return nil
}

func (it GenerateFunc) testPkgName(toWrap *args.FuncWrap) string {
	return toWrap.PkgName() + "tests"
}

func (it GenerateFunc) toFunWrap() *args.FuncWrap {
	return args.
		NewFuncWrap.
		Default(it.Func)
}

func (it GenerateFunc) generateActLines() interface{} {

}
