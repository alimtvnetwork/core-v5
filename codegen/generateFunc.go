package codegen

import (
	"errors"
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

func (it GenerateFunc) generateActLines() []string {

}

func (it GenerateFunc) generateFmtJoin() string {
	return "   %d : "
}

func (it GenerateFunc) getFuncName() (string, error) {
	funcWrap := it.toFunWrap()

	if funcWrap.IsInvalid() {
		return "", errors.New("func wrap is invalid - func name")
	}

	return funcWrap.GetFuncName(), nil
}

func (it GenerateFunc) getReturnArgs() (*corestr.SimpleSlice, error) {
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

func (it GenerateFunc) emptySlice() *corestr.SimpleSlice {
	return corestr.Empty.SimpleSlice()
}
