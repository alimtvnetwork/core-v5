package codegen

import (
	"reflect"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/codegen/codegentype"
	"github.com/alimtvnetwork/core/codegen/fmtcodegentype"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coreinterface"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

type FmtOptions struct {
	HasExpect bool
	BaseArgs  args.AsArgBaseContractsBinder
	FuncArgs  args.ArgFuncContractsBinder
}

type BaseGenerator interface {
	Function() any
	GetStruct() any
	GenType() codegentype.Variant
	JoinFormatType() fmtcodegentype.Variant
	Cases() []coretestcases.CaseV1
	GetBehavioursGetter
	OverrideFuncCall() string
	GetOverrides() OverridingNames
	IsFunctionIncluded() bool
	Generate() error
	GenerateCodeOutput() *FinalCode
	FmtJoin() string
	// HasInnerLoop checks if the GenerateFunc has an inner loop.
	//
	// It returns a boolean indicating whether the GenerateType has multiple arranges.
	HasInnerLoop() bool
	UnitTests(
		inArgs,
		outArgs *corestr.SimpleSlice,
		tempMap map[string]string,
	) (*corestr.SimpleSlice, error)
	TestCaseName(
		totalBehaviours int,
		funcName,
		behaviour string,
	) string
	FirstArrangeTypeName() string
	AllPackages() *corestr.Hashset
	FirstArrangeType() *reflect.Type
	ArrangeReflectTypes() []reflect.Type
	FirstTestCaseGetter
	ArrangeImports() *corestr.Hashset
	TestPkgName() string
	FuncWrap() *args.FuncWrapAny

	coreinterface.DirectFuncNameGetter

	ArgsOutter
	ArgsInner

	VariableNameGetter
	DirectFuncInvokeName() string

	FileWriter() *chmodhelper.SimpleFileReaderWriter

	StructNameGetter
	NewGoCode(codes ...string) *GoCode

	TestCasesCompiler
}

type TestCasesCompiler interface {
	TestCasesCompiledCode() (*GoCode, error)
}

type VariableNameGetter interface {
	VariableName(parentVar string, index int) string
}

type StructNameGetter interface {
	StructName() string
}

type FirstTestCaseGetter interface {
	FirstTestCase() *coretestcases.CaseV1
}

type ArgsInner interface {
	InArgs() (*corestr.SimpleSlice, error)
}

type ArgsOutter interface {
	OutArgs() (*corestr.SimpleSlice, error)
}

type GetBehavioursGetter interface {
	GetBehaviours() corestr.SimpleSlice
}
