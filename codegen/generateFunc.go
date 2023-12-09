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
	"gitlab.com/auk-go/core/coreindexes"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coreutils/stringutil"
	"gitlab.com/auk-go/core/internal/pathinternal"
	"gitlab.com/auk-go/core/internal/reflectinternal"
	"gitlab.com/auk-go/core/isany"
	"gitlab.com/auk-go/core/iserror"
)

type GenerateFunc struct {
	Func                    interface{}
	Struct                  interface{}
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

func (it GenerateFunc) Function() interface{} {
	return it.Func
}

func (it GenerateFunc) CurStruct() interface{} {
	return it.Struct
}

func (it GenerateFunc) GenType() codegentype.Variant {
	return it.GenerateType
}

func (it GenerateFunc) JoinFormatType() fmtcodegentype.Variant {
	return it.FmtType
}

func (it GenerateFunc) Cases() []coretestcases.CaseV1 {
	return it.TestCases
}

func (it GenerateFunc) CurBehaviours() corestr.SimpleSlice {
	return it.Behaviours
}

func (it GenerateFunc) CurFuncOverrideCall() interface{} {
	return it.FuncOverrideCall
}

func (it GenerateFunc) IsFunctionInclude() bool {
	return it.IsIncludeFunction
}

func (it GenerateFunc) Generate() error {
	codeOutput := it.GenerateCodeOutput()

	return codeOutput.Write().CompiledError()
}

func (it GenerateFunc) GenerateCodeOutput() *CodeOutput {
	toWrap := it.FuncWrap()

	if toWrap.IsInvalid() {
		return NewCodeOutput.Invalid(toWrap.InvalidError())
	}

	testPkgName, packageHeader := it.PackageHeader()

	inArgs, inArgsErr := it.InArgs()

	if iserror.Defined(inArgsErr) {
		return NewCodeOutput.Invalid(inArgsErr)
	}

	outArgs, outArgsErr := it.OutArgs()

	if iserror.Defined(outArgsErr) {
		return NewCodeOutput.Invalid(outArgsErr)
	}

	funcName := toWrap.GetFuncName()
	firstArrangeTypeName := it.FirstArrangeTypeName()
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
		unitTestVars.InArgs:           inArgs.Join(ArgsJoiner),
		unitTestVars.OutArgs:          outArgs.Join(ArgsJoiner),
		unitTestVars.FmtJoin:          it.generateFmtJoin(),
		unitTestVars.FmtOutputs:       fmtOutputs.Join(fmtJoiner),
		unitTestVars.DirectFuncInvoke: it.DirectFuncInvokeName(),
	}

	unitTests, unitErr := it.UnitTests(
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

	testCaseCompiled := it.TestCasesCompiled()

	return &CodeOutput{
		UnitTest:   finalUnitTest,
		TestCase:   testCaseCompiled,
		StructName: it.StructName(),
		FuncName:   funcName,
		FileWriter: it.fileWriter(testPkgName),
	}
}

func (it GenerateFunc) UnitTests(
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

	funcName := it.FuncName()

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
		tempMap[unitTestVars.TestCaseName] = it.TestCaseName(
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

func (it GenerateFunc) TestCaseName(
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

func (it GenerateFunc) PackageHeader() (testPkgName string, packageHeader string) {
	testPkgName = it.TestPkgName()
	newPackagesLines := it.AllPackages()
	packagesTemplate := map[string]string{
		"$packageName": testPkgName,
		"$newPackages": newPackagesLines,
	}

	packageHeader = stringutil.
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

func (it GenerateFunc) FirstArrangeTypeName() string {
	rt := it.FirstArrangeType()

	if rt == nil {
		return constants.NilAngelBracket
	}

	return (*rt).String()
}

func (it GenerateFunc) AllPackages() string {
	arrangePkgPaths := it.ArrangePackages()

	newPackages := corestr.
		New.
		SimpleSlice.
		Hashset(arrangePkgPaths).
		Add(it.FuncWrap().PkgPath()).
		WrapDoubleQuote()

	newPackagesLines := newPackages.JoinLine()

	return newPackagesLines
}

func (it GenerateFunc) FirstArrangeType() *reflect.Type {
	if len(it.TestCases) == 0 {
		return nil
	}

	rt := reflect.TypeOf(
		it.TestCases[0].ArrangeInput,
	)

	return &rt
}

func (it GenerateFunc) ArrangeReflectTypes() []reflect.Type {
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

func (it GenerateFunc) ArrangePackages() *corestr.Hashset {
	allReflectTypes := it.ArrangeReflectTypes()

	pks := corestr.New.Hashset.Cap(len(allReflectTypes))

	for _, reflectType := range allReflectTypes {
		pks.Add(reflectType.PkgPath())
	}

	return pks
}

func (it GenerateFunc) TestPkgName() string {
	return it.FuncWrap().PkgNameOnly() + "tests"
}

func (it GenerateFunc) FuncWrap() *args.FuncWrap {
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

func (it GenerateFunc) FuncName() string {
	funcWrap := it.FuncWrap()

	if funcWrap.IsInvalid() {
		return ""
	}

	return funcWrap.GetFuncName()
}

// OutArgs
//
//	Aka returns Args
//
// - if one then return "result" only
// - Or else, result1, result2 ...
func (it GenerateFunc) OutArgs() (*corestr.SimpleSlice, error) {
	funcWrap := it.FuncWrap()

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

// InArgs
//
// - if one then return "result" only
// - Or else, result1, result2 ...
func (it GenerateFunc) InArgs() (*corestr.SimpleSlice, error) {
	funcWrap := it.FuncWrap()

	if funcWrap.IsInvalid() {
		return it.emptySlice(), errors.New("func wrap is invalid - return args")
	}

	length := funcWrap.ArgsCount()
	slice := corestr.New.SimpleSlice.Cap(length)

	if length == 0 {
		return slice, nil
	}

	if length == 1 {
		return slice.Add(it.VariableName("input", 0)), nil
	}

	for i := 0; i < length; i++ {
		slice.Add(it.VariableName("input", i))
	}

	return slice, nil
}

// VariableName
//
// variable.First or variable.Second ... based on index.
func (it GenerateFunc) VariableName(parentVar string, index int) string {
	return parentVar + "." + it.indexByName(index)
}

func (it GenerateFunc) indexByName(index int) string {
	return coreindexes.NameByIndex(index)
}

func (it GenerateFunc) emptySlice() *corestr.SimpleSlice {
	return corestr.Empty.SimpleSlice()
}

func (it GenerateFunc) DirectFuncInvokeName() string {
	if len(it.FuncOverrideCall) > 0 {
		return it.FuncOverrideCall
	}

	return it.FuncWrap().FuncDirectInvokeName()
}

func (it GenerateFunc) StructName() string {
	if isany.Null(it.Struct) {
		return ""
	}

	return reflectinternal.TypeName(it.Struct)
}

func (it GenerateFunc) TestCasesCompiled() string {
	return ""
}
