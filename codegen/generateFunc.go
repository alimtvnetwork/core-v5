package codegen

import (
	"errors"
	"fmt"
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
	firstArrangeTypeName := it.firstArrangeTypeName()
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

	funcTemplateReplacer := map[string]string{
		unitTestVars.FuncName:         funcName,
		unitTestVars.ArrangeType:      firstArrangeTypeName,
		unitTestVars.LinesPossible:    totalSliceLength,
		unitTestVars.ActArgsSetup:     actLines.JoinLine(),
		unitTestVars.InArgs:           inArgs.Join(ArgsJoiner),
		unitTestVars.OutArgs:          outArgs.Join(ArgsJoiner),
		unitTestVars.FmtJoin:          it.generateFmtJoin(),
		unitTestVars.FmtOutputs:       fmtOutputs.Join(fmtJoiner),
		unitTestVars.DirectFuncInvoke: it.directFuncInvoke(),
	}

	unitTests, unitErr := it.unitTests(
		inArgs,
		outArgs,
		funcTemplateReplacer,
	)

	if iserror.Defined(unitErr) {
		return NewCodeOutput.Invalid(unitErr)
	}

	finalUnitTest := stringslice.Joins(
		constants.NewLineUnix,
		packageHeader,
		"",
		unitTests.JoinLine(),
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

func (it GenerateFunc) unitTests(
	inArgs,
	outArgs *corestr.SimpleSlice,
	tempMap map[string]string,
) (*corestr.SimpleSlice, error) {
	totalBehaviours := len(it.Behaviours)
	testsSlice := corestr.
		New.
		SimpleSlice.
		Cap(totalBehaviours)

	if totalBehaviours == 0 {
		return testsSlice, errors.New("must set behaviours it cannot be empty")
	}

	funcName := it.funcName()

	for _, behaviour := range it.Behaviours {
		fmtOutputs, fmtErr := it.generateFmtOutputs(
			fmtJoiner,
			funcName,
			unitTestVars.inputExpectedVar,
			outArgs,
			inArgs,
		)

		tempMap[unitTestVars.FmtOutputs] = fmtOutputs.Join(fmtJoiner)
		tempMap[unitTestVars.Behaviour] = behaviour
		tempMap[unitTestVars.TestCaseName] = it.testCaseName(
			totalBehaviours,
			funcName,
			behaviour,
		)

		if iserror.Defined(fmtErr) {
			return testsSlice, fmtErr
		}

		unitTest := stringutil.
			ReplaceTemplate.
			DirectKeyUsingMapTrim(
				funcTemplate,
				tempMap,
			)

		testsSlice.Add(unitTest)
	}

	return testsSlice, nil
}

func (it GenerateFunc) testCaseName(
	totalBehaviours int,
	funcName,
	behaviour string,
) string {
	if totalBehaviours == 1 {
		return camelCaseFunc(
			fmt.Sprintf(
				"%sTestCases",
				funcName,
			),
		)
	}

	return camelCaseFunc(
		fmt.Sprintf(
			"%sTestCases%s",
			funcName,
			pascalCaseFunc(behaviour),
		),
	)
}

func (it GenerateFunc) packageHeader(toWrap *args.FuncWrap) (string, string) {
	testPkgName := it.testPkgName(toWrap)
	newPackagesLines := it.allPackages(toWrap)
	packagesTemplate := map[string]string{
		"$packageName": testPkgName,
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

func (it GenerateFunc) funcName() string {
	funcWrap := it.toFunWrap()

	if funcWrap.IsInvalid() {
		return ""
	}

	return funcWrap.GetFuncName()
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

	if length == 0 {
		return slice, nil
	}

	if length == 1 {
		return slice.Add(it.variableName("input", 0)), nil
	}

	for i := 0; i < length; i++ {
		slice.Add(it.variableName("input", i))
	}

	return slice, nil
}

func (it GenerateFunc) variableName(parentVar string, index int) string {
	return parentVar + "." + it.indexByName(index)
}

func (it GenerateFunc) indexByName(index int) string {
	return indexByNameMap[index]
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
