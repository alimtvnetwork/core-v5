package codegen

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/codegen/codegentype"
	"github.com/alimtvnetwork/core/codegen/fmtcodegentype"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coreindexes"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/internal/convertinternal"
	"github.com/alimtvnetwork/core/internal/pathinternal"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
	"github.com/alimtvnetwork/core/isany"
	"github.com/alimtvnetwork/core/iserror"
	"github.com/alimtvnetwork/core/simplewrap"
)

type GenerateFunc struct {
	Func                 any
	Struct               any
	GenerateType         codegentype.Variant
	FmtType              fmtcodegentype.Variant
	TestCases            []coretestcases.CaseV1
	Behaviours           corestr.SimpleSlice
	OverridingNames      OverridingNames
	UnitTestRootPath     string
	Options              Options
	packageHeader        corestr.SimpleStringOnce
	funcWrap             *args.FuncWrapAny
	setupVariable        *variablesSetup
	directFuncInvokeName string
}

func (it GenerateFunc) GetOverrides() OverridingNames {
	return it.OverridingNames
}

func (it GenerateFunc) Function() any {
	return it.Func
}

func (it GenerateFunc) GetStruct() any {
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

func (it GenerateFunc) GetBehaviours() corestr.SimpleSlice {
	return it.Behaviours
}

func (it GenerateFunc) OverrideFuncCall() string {
	return it.OverridingNames.FuncCall
}

func (it GenerateFunc) IsFunctionIncluded() bool {
	return it.Options.IsIncludeFunction
}

func (it GenerateFunc) Generate() error {
	codeOutput := it.GenerateCodeOutput()

	return codeOutput.Write().CompiledError()
}

func (it GenerateFunc) GenerateCodeOutput() *FinalCode {
	toWrap := it.FuncWrap()

	if toWrap.IsInvalid() {
		return New.FinalCode.Invalid(toWrap.InvalidError())
	}

	inArgs, inArgsErr := it.InArgs()

	if iserror.Defined(inArgsErr) {
		return New.FinalCode.Invalid(inArgsErr)
	}

	outArgs, outArgsErr := it.OutArgs()

	if iserror.Defined(outArgsErr) {
		return New.FinalCode.Invalid(outArgsErr)
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
		return New.FinalCode.Invalid(fmtErr)
	}

	funcTemplateReplacer := map[string]string{
		vars.FuncName:         funcName,
		vars.ArrangeType:      firstArrangeTypeName,
		vars.LinesPossible:    totalSliceLength,
		vars.InArgs:           Utils.ParametersJoin(inArgs),
		vars.OutArgs:          Utils.ParametersJoin(outArgs),
		vars.FmtJoin:          it.FmtJoin(),
		vars.FmtOutputs:       fmtOutputs.Join(fmtJoiner),
		vars.DirectFuncInvoke: it.DirectFuncInvokeName(),
	}

	unitTests, unitErr := it.UnitTests(
		inArgs,
		outArgs,
		funcTemplateReplacer,
	)

	if iserror.Defined(unitErr) {
		return New.FinalCode.Invalid(unitErr)
	}

	unitTestCode := it.NewGoCode(unitTests.JoinLine())
	testCaseCompiled, testCaseErr := it.TestCasesCompiledCode()

	return &FinalCode{
		UnitTest:   unitTestCode,
		TestCase:   testCaseCompiled,
		StructName: it.StructName(),
		FuncName:   funcName,
		Error:      testCaseErr,
		FileWriter: it.internalFileWriter(it.TestPkgName()),
		Options:    it.Options,
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
			vars.inputExpected,
			outArgs,
			inArgs,
		)

		tempMap[vars.FmtOutputs] = fmtOutputs.Join(fmtJoiner)
		tempMap[vars.Behaviour] = behaviour
		tempMap[vars.TestCaseName] = it.TestCaseName(
			totalBehaviours,
			funcName,
			behaviour,
		)
		tempMap[vars.VariablesSetup] = it.CompiledVariablesSetup()
		if iserror.Defined(fmtErr) {
			return testsSlice, fmtErr
		}

		unitTest := it.ReplaceTemplate(
			it.FuncTemplatedCode(),
			tempMap,
		)

		testsSlice.Add(unitTest)
	}

	return testsSlice, nil
}

func (it GenerateFunc) FuncTemplatedCode() string {
	return functionTemplatesMap[it.GenerateType]
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

//
// func (it GenerateFunc) PackageHeader() (testPkgName string, packageHeader string) {
// 	testPkgName = it.testPkgName()
//
// 	if it.packageHeader.IsDefined() {
// 		return testPkgName, it.packageHeader.String()
// 	}
//
// 	newPackagesLines := it.AllPackages()
// 	packagesTemplate := map[string]string{
// 		"$packageName": testPkgName,
// 		"$newPackages": newPackagesLines,
// 	}
//
// 	packageHeader = it.ReplaceTemplate(
// 		testPkgHeaderTemplate,
// 		packagesTemplate,
// 	)
//
// 	return testPkgName, it.packageHeader.GetSetOnce(packageHeader)
// }

func (it GenerateFunc) internalFileWriter(unitTestPackageName string) *chmodhelper.SimpleFileReaderWriter {
	finalUnitTestPath := it.unitTestRootPath(unitTestPackageName)

	return chmodhelper.
		New.
		SimpleFileReaderWriter.
		Options(
			it.Options.IsOverwrite,
			true,
			true,
			finalUnitTestPath,
		)
}

func (it GenerateFunc) FileWriter() *chmodhelper.SimpleFileReaderWriter {
	return it.internalFileWriter(it.TestPkgName())
}

func (it GenerateFunc) unitTestRootPath(unitTestPackageName string) string {
	return pathinternal.Join(
		it.UnitTestRootPath,
		unitTestPackageName,
		"x.go", // we are writing to the parent dir
	)
}

func (it GenerateFunc) TestRootPath() string {
	return it.unitTestRootPath(it.TestPkgName())
}

func (it GenerateFunc) FirstArrangeTypeName() string {
	if len(it.TestCases) == 0 {
		return ""
	}

	return convertinternal.AnyTo.TypeName(
		it.TestCases[0].ArrangeInput,
	)
}

func (it GenerateFunc) FirstTestCase() *coretestcases.CaseV1 {
	if len(it.TestCases) == 0 {
		return nil
	}

	return &it.TestCases[0]
}

func (it GenerateFunc) AllPackages() *corestr.Hashset {
	arrangePkgPaths := it.ArrangeImports()
	funcPkgPath := it.FuncWrap().PkgPath()

	newPackages := Utils.AllPackages(
		funcPkgPath,
		arrangePkgPaths.List()...,
	)

	return newPackages
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

func (it GenerateFunc) ArrangeImports() *corestr.Hashset {
	allReflectTypes := it.ArrangeReflectTypes()

	pks := corestr.New.Hashset.Cap(len(allReflectTypes))

	for _, reflectType := range allReflectTypes {
		pks.Add(simplewrap.WithDoubleQuote(reflectType.PkgPath()))
	}

	return pks
}

func (it GenerateFunc) TestPkgName() string {
	if it.OverridingNames.HasTestPkgName() {
		return it.OverridingNames.TestPkgName
	}

	return it.FuncWrap().PkgNameOnly() + "tests"
}

func (it GenerateFunc) FuncWrap() *args.FuncWrapAny {
	if it.funcWrap != nil {
		return it.funcWrap
	}

	it.funcWrap = args.
		NewFuncWrap.
		Default(it.Func)

	return it.funcWrap
}

// FmtJoin
//
//	https://prnt.sc/Susd-5ZPdDvp
//
//	returns the AppendFmt format string inside the unit test
func (it GenerateFunc) FmtJoin() string {
	return it.FmtType.Fmt()
}

// HasInnerLoop checks if the GenerateFunc has an inner loop.
//
// It returns a boolean indicating whether the GenerateType has multiple arranges.
func (it GenerateFunc) HasInnerLoop() bool {
	return it.GenerateType.IsMultipleArranges()
}

// generateFmtOutputs generates the formatted outputs based on the given joiner, function name, expected output,
// and input and output arguments. It returns a *corestr.SimpleSlice containing the formatted outputs and an error, if any.
//
// Parameters:
//   - joiner: The string used to join the arguments.
//   - funcName: The name of the function.
//   - expected: The expected output.
//   - outArs: A *corestr.SimpleSlice containing the output arguments.
//   - inArgs: A *corestr.SimpleSlice containing the input arguments.
//
// Returns:
//   - https://prnt.sc/rFtWpwFxFnVm, https://prnt.sc/9a3wEshInxJv
//   - Generates each comma single line join for the AppendFmt(format, allItems....)
//   - The function will generate allItems as it comes from args
//   - error: An error, if any.
func (it GenerateFunc) generateFmtOutputs(
	joiner string,
	funcName string,
	expected string,
	outArs, inArgs *corestr.SimpleSlice,
) (*corestr.SimpleSlice, error) {
	slice := corestr.New.SimpleSlice.Cap(20)

	switch it.FmtType {
	case fmtcodegentype.Default: // "%d : %s -> %s",
		// outArgsString := outArs.Join(joiner)
		// inArgsString := inArgs.Join(joiner)

		slice.Add(vars.allInArgsCompiled)
		slice.Add(vars.allOutArgsCompiled)

		return slice, nil
	case fmtcodegentype.WithExpect: // %d : %s -> %s | %s,
		// outArgsString := outArs.Join(joiner)
		// inArgsString := inArgs.Join(joiner)

		slice.Add(vars.allInArgsCompiled)
		slice.Add(vars.allOutArgsCompiled)
		slice.Add(expected)

		return slice, nil
	case fmtcodegentype.WithFuncExpect: // "%d : %s(%s) -> %s | %s",
		// outArgsString := outArs.Join(joiner)
		// inArgsString := inArgs.Join(joiner)

		slice.Add(funcName)
		slice.Add(vars.allInArgsCompiled)
		slice.Add(vars.allOutArgsCompiled)
		slice.Add(expected)

		return slice, nil
	}

	return slice, it.FmtType.OnlySupportedMsgErr(
		"only supported",
		fmtcodegentype.Default.Name(),
		fmtcodegentype.WithExpect.Name(),
		fmtcodegentype.WithFuncExpect.Name(),
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

	return &it.VariablesSetup().inArgsNames, nil
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

func (it *GenerateFunc) DirectFuncInvokeName() string {
	if len(it.directFuncInvokeName) > 0 {
		return it.directFuncInvokeName
	}

	if len(it.OverridingNames.FuncCall) > 0 {
		it.directFuncInvokeName = Utils.ChainEachLine(
			it.OverridingNames.FuncCall,
		)
	} else {
		it.directFuncInvokeName = Utils.ChainEachLine(
			it.FuncWrap().FuncDirectInvokeName(),
		)
	}

	return it.directFuncInvokeName
}

func (it GenerateFunc) StructName() string {
	if isany.Null(it.Struct) {
		return ""
	}

	return reflectinternal.TypeName(it.Struct)
}

func (it GenerateFunc) ReplaceTemplate(
	format string,
	replacerMap map[string]string,
) string {
	if len(format) == 0 {
		return ""
	}

	return templateReplacerFunc(
		format,
		replacerMap,
	)
}

func (it GenerateFunc) TestCasesCompiledCode() (*GoCode, error) {
	caseGenerator := testCaseGenerator{
		baseGenerator: it.AsBaseGenerator(),
	}

	return caseGenerator.Compile()
}

func (it *GenerateFunc) VariablesSetup() *variablesSetup {
	if it.setupVariable != nil {
		return it.setupVariable
	}

	generator := variablesGenerator{
		baseGenerator: it.AsBaseGenerator(),
	}

	vs := generator.Generate()
	it.setupVariable = &vs

	return it.setupVariable
}

func (it GenerateFunc) CompiledVariablesSetup() string {
	return it.VariablesSetup().CompiledSetupLine()
}

func (it GenerateFunc) TestRootDir() string {
	return pathinternal.ParentDir(it.TestRootPath())
}

func (it GenerateFunc) SuccessMessage() string {
	return fmt.Sprintf(
		"Created tests successfully at %s",
		it.TestRootDir(),
	)
}

func (it GenerateFunc) FailedMessage() string {
	return fmt.Sprintf(
		"Tried to create tests at %s, but failed",
		it.TestRootDir(),
	)
}

func (it GenerateFunc) NewGoCode(codes ...string) *GoCode {
	return New.GoCode.Create(it.AsBaseGenerator(), codes...)
}

func (it GenerateFunc) AsBaseGenerator() BaseGenerator {
	return &it
}
