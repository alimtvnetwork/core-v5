package codegen

import (
	"reflect"

	"gitlab.com/auk-go/core/codegen/codegentype"
	"gitlab.com/auk-go/core/codegen/fmtcodegentype"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

type BaseGenerator interface {
	Function() interface{}
	GetStruct() interface{}
	GenType() codegentype.Variant
	JoinFormatType() fmtcodegentype.Variant
	Cases() []coretestcases.CaseV1
	GetBehavioursGetter
	OverrideFuncCall() string
	GetOverrides() OverridingNames
	IsFunctionIncluded() bool
	Generate() error
	GenerateCodeOutput() *AllCodeOutput
	FmtJoin() string
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
	ArrangePackages() *corestr.Hashset
	TestPkgName() string
	FuncWrap() *args.FuncWrap

	coreinterface.DirectFuncNameGetter

	ArgsOutter
	ArgsInner

	VariableNameGetter
	DirectFuncInvokeName() string

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
