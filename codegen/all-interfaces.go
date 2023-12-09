package codegen

import (
	"reflect"

	"gitlab.com/auk-go/core/codegen/codegentype"
	"gitlab.com/auk-go/core/codegen/fmtcodegentype"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

type BaseGenerator interface {
	Function() interface{}
	CurStruct() interface{}
	GenType() codegentype.Variant
	JoinFormatType() fmtcodegentype.Variant
	Cases() []coretestcases.CaseV1
	CurBehaviours() corestr.SimpleSlice
	CurFuncOverrideCall() interface{}
	IsFunctionInclude() bool
	Generate() error
	GenerateCodeOutput() *CodeOutput
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
	PackageHeader() (testPkgName string, packageHeader string)
	FirstArrangeTypeName() string
	AllPackages() string
	FirstArrangeType() *reflect.Type
	ArrangeReflectTypes() []reflect.Type

	ArrangePackages() *corestr.Hashset
	TestPkgName() string
	FuncWrap() *args.FuncWrap
	FuncName() string
}
