package codegen

import (
	"errors"
	"reflect"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/codegen/codegentype"
	"gitlab.com/auk-go/core/codegen/fmtcodegentype"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coreutils/stringutil"
	"gitlab.com/auk-go/core/internal/pathinternal"
	"gitlab.com/auk-go/core/internal/reflectinternal"
	"gitlab.com/auk-go/core/iserror"
)

type GenerateFunc struct {
	Func                    interface{}
	FuncOverrideCall        string
	GenerateType            codegentype.Variant
	FmtType                 fmtcodegentype.Variant
	TestCases               []coretestcases.CaseV1
	Behaviours              corestr.SimpleSlice
	UnitTestRootPath        string
	OverridingTestPkgName   string
	IsGenerateSeparateCases bool
	IsIncludeFunction       bool
	IsOverwrite             bool
	funcWrap                *args.FuncWrap
}

func (it GenerateFunc) Generate() error {
	codeOutput := it.GenerateCodeOutput()

	return codeOutput.Write().CompiledError()
}

func (it GenerateFunc) GenerateCodeOutput() *CodeOutput {
	toWrap := it.toFunWrap()

	if toWrap.IsInvalid() {
		return NewCodeOutput.Invalid(toWrap.InvalidError())
	}

	testPkgName, packageHeader := it.packageHeader(toWrap)

	actLines := it.generateActLines()
	inArgs, inArgsErr := it.inArgs()

	if iserror.Defined(inArgsErr) {
		return NewCodeOutput.Invalid(inArgsErr)
	}

	outArgs, outArgsErr := it.outArgs()

	if iserror.Defined(outArgsErr) {
		return NewCodeOutput.Invalid(outArgsErr)
	}

	funcName := toWrap.GetFuncName()
	fmtOutputs, fmtErr := it.generateFmtOutputs(
		fmtJoiner,
		funcName,
		"",
		outArgs,
		inArgs,
	)

	if iserror.Defined(fmtErr) {
		return NewCodeOutput.Invalid(fmtErr)
	}

	firstArrangeTypeName := it.firstArrangeTypeName()

	funcTemplateReplacer := map[string]string{
		"$FuncName":         funcName,
		"$ArrangeType":      firstArrangeTypeName,
		"$linesPossible":    "100",
		"$actArgsSetup":     actLines.JoinLine(),
		"$inArgs":           inArgs.Join(ArgsJoiner),
		"$outArgs":          outArgs.Join(ArgsJoiner),
		"$fmtJoin":          it.generateFmtJoin(),
		"$fmtOutputs":       fmtOutputs.Join(fmtJoiner),
		"$directFuncInvoke": it.directFuncInvoke(),
	}

	unitTest := stringutil.
		ReplaceTemplate.
		DirectKeyUsingMapTrim(
			funcTemplate,
			funcTemplateReplacer,
		)

	finalUnitTest := stringslice.JoinWith(
		constants.NewLineUnix,
		packageHeader,
		"",
		unitTest,
		"",
	)

	return &CodeOutput{
		UnitTest:   finalUnitTest,
		TestCase:   "",
		StructName: "",
		FuncName:   funcName,
		FileWriter: it.fileWriter(testPkgName),
	}
}

func (it GenerateFunc) packageHeader(toWrap *args.FuncWrap) (string, string) {
	testPkgName := it.testPkgName(toWrap)
	newPackagesLines := it.allPackages(toWrap)
	packagesTemplate := map[string]string{
		"$packageName": testPkgName,
		"$fmtJoin":     it.generateFmtJoin(),
		"$newPackages": newPackagesLines,
	}

	packageHeader := stringutil.
		ReplaceTemplate.
		DirectKeyUsingMapTrim(
			testPkgHeaderTemplate,
			packagesTemplate,
		)

	return testPkgName, packageHeader
}

func (it GenerateFunc) fileWriter(unitTestPackageName string) *chmodhelper.SimpleFileReaderWriter {
	finalUnitTestPath := it.unitTestRootPath(unitTestPackageName)

	return chmodhelper.
		New.
		SimpleFileReaderWriter.
		Options(
			true,
			true,
			true,
			finalUnitTestPath,
		)
}

func (it GenerateFunc) unitTestRootPath(unitTestPackageName string) string {
	return pathinternal.Join(
		it.UnitTestRootPath,
		unitTestPackageName,
		"x.go",
	)
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
		Hashset(arrangePkgPaths).
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

	reducerFunc := reflectinternal.Looper.ReducePointerDefault

	for _, testCase := range it.TestCases {
		r := reducerFunc(testCase)

		if r.IsInvalid() {
			continue
		}

		results = append(
			results,
			r.FinalReflectVal.Type(),
		)
	}

	return results
}

func (it GenerateFunc) arrangePackages() *corestr.Hashset {
	allReflectTypes := it.arrangeReflectTypes()

	pks := corestr.New.Hashset.Cap(len(allReflectTypes))

	for _, reflectType := range allReflectTypes {
		pks.Add(reflectType.PkgPath())
	}

	return pks
}

func (it GenerateFunc) testPkgName(toWrap *args.FuncWrap) string {
	return toWrap.PkgNameOnly() + "tests"
}

func (it GenerateFunc) toFunWrap() *args.FuncWrap {
	if it.funcWrap != nil {
		return it.funcWrap
	}

	it.funcWrap = args.
		NewFuncWrap.
		Default(it.Func)

	return it.funcWrap
}

func (it GenerateFunc) generateActLines() *corestr.SimpleSlice {
	return nil
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

	return slice, it.FmtType.OnlySupportedMsgErr(
		"only supported",
		fmtcodegentype.Default.Name(),
		fmtcodegentype.WithFunction.Name(),
	)
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

func (it GenerateFunc) directFuncInvoke() string {
	if len(it.FuncOverrideCall) > 0 {
		return it.FuncOverrideCall
	}

	return it.toFunWrap().FuncDirectInvokeName()
}
