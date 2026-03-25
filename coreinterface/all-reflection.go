package coreinterface

import (
	"fmt"
	"reflect"
)

type FuncNameGetter interface {
	GetFuncName() string
}

type PkgPathGetter interface {
	PkgPath() string
}

type PkgNameGetter interface {
	PkgNameOnly() string
}

type HasValidFuncChecker interface {
	HasValidFunc() bool
}

type HasFuncChecker interface {
	HasFunc() bool
}

type ArgsCountGetter interface {
	ArgsCount() int
}

type ArgsLengthGetter interface {
	ArgsLength() int
}

type ReturnLengthGetter interface {
	ReturnLength() int
}

type IsPublicMethodGetter interface {
	IsPublicMethod() bool
}

type IsPrivateMethodGetter interface {
	IsPrivateMethod() bool
}

type TypeGetter interface {
	GetType() reflect.Type
}

type OutArgsTypesGetter interface {
	GetOutArgsTypes() []reflect.Type
}

type InArgsTypesGetter interface {
	GetInArgsTypes() []reflect.Type
}

type InArgsTypesNamesGetter interface {
	GetInArgsTypesNames() []string
}

type IsInTypeMatchesChecker interface {
	IsInTypeMatches(args ...any) (isOkay bool)
}

type IsOutTypeMatchesChecker interface {
	IsOutTypeMatches(outArgs ...any) (isOkay bool)
}

type InArgsVerifier interface {
	VerifyInArgs(args []any) (isOkay bool, err error)
}

type OutArgsVerifier interface {
	VerifyOutArgs(args []any) (isOkay bool, err error)
}

type InArgsRvVerifier interface {
	InArgsVerifyRv(args []reflect.Type) (isOkay bool, err error)
}

type OutArgsRvVerifier interface {
	OutArgsVerifyRv(args []reflect.Type) (isOkay bool, err error)
}

type VoidCallNoReturner interface {
	VoidCallNoReturn(
		args ...any,
	) (processingErr error)
}

type MustBeValidater interface {
	MustBeValid()
}

type MustInvoker interface {
	InvokeMust(
		args ...any,
	) []any
}

type ReflectInvoker interface {
	Invoke(
		args ...any,
	) (results []any, processingErr error)
}

type VoidCaller interface {
	VoidCall() ([]any, error)
}

type ValidateMethodArgsGetter interface {
	ValidateMethodArgs(args []any) error
}

type FirstResponseOfInvokeGetter interface {
	GetFirstResponseOfInvoke(
		args ...any,
	) (firstResponse any, err error)
}

type InvokeResultOfIndexGetter interface {
	InvokeResultOfIndex(
		index int,
		args ...any,
	) (firstResponse any, err error)
}

type InvokeErrorGetter interface {
	InvokeError(
		args ...any,
	) (funcErr, processingErr error)
}

type InvokeFirstAndErrorGetter interface {
	InvokeFirstAndError(
		args ...any,
	) (firstResponse any, funcErr, processingErr error)
}

type FirstItemGetter interface {
	FirstItem() any
}

type SecondItemGetter interface {
	SecondItem() any
}

type ThirdItemGetter interface {
	ThirdItem() any
}

type FourthItemGetter interface {
	FourthItem() any
}

type FifthItemGetter interface {
	FifthItem() any
}

type SixthItemGetter interface {
	SixthItem() any
}

type ExpectGetter interface {
	Expected() any
}

type ArrangeGetter interface {
	Arrange() any
}

type ActualGetter interface {
	Actual() any
}

type UptoSecondItemGetter interface {
	FirstItemGetter
	SecondItemGetter
}

type UptoThirdItemGetter interface {
	UptoSecondItemGetter
	ThirdItemGetter
}

type UptoFourthItemGetter interface {
	UptoThirdItemGetter
	FourthItemGetter
}

type UptoFifthItemGetter interface {
	UptoFourthItemGetter
	FifthItemGetter
}

type UptoSixthItemGetter interface {
	UptoFifthItemGetter
	SixthItemGetter
}

type FuncWrapContractsBinder interface {
	FuncNameGetter
	PkgPathGetter
	PkgNameGetter
	HasValidFuncChecker
	IsValidChecker
	IsInvalidChecker
	ArgsCountGetter
	ArgsLengthGetter
	ReturnLengthGetter
	IsPublicMethodGetter
	IsPrivateMethodGetter
	TypeGetter
	OutArgsTypesGetter
	InArgsTypesGetter
	InArgsTypesNamesGetter
	InArgsVerifier
	OutArgsVerifier
	InArgsRvVerifier
	OutArgsRvVerifier
	VoidCallNoReturner
	MustBeValidater
	ValidationErrorGetter

	GetInArgsTypesNames() []string
	GetOutArgsTypesNames() []string
	OutArgNames() []string
	InArgNames() []string
	OutArgsCount() int
	InArgsCount() int
	ArgsCount() int
	PkgNameOnly() string
	GetPascalCaseFuncName() string
	GetFuncName() string

	MustInvoker
	ReflectInvoker
	VoidCaller
	ValidateMethodArgsGetter

	FirstResponseOfInvokeGetter
	InvokeResultOfIndexGetter
	InvokeErrorGetter
	InvokeFirstAndErrorGetter
}

type SliceGetter interface {
	Slice() []any
}

type ArgsUptoGetter interface {
	Args(upTo int) []any
}

type ValidArgsGetter interface {
	ValidArgs() []any
}

type HasExpectChecker interface {
	HasExpect() bool
}

type ByIndexGetter interface {
	GetByIndex(index int) any
}

type OneParameter interface {
	FirstItemGetter
	ExpectGetter

	HasFirst() bool
	HasExpectChecker

	ArgsCountGetter

	ValidArgsGetter
	ArgsUptoGetter
	SliceGetter
	ByIndexGetter
}

type FuncParameter interface {
	HasFuncChecker
	FuncNameGetter
	ReflectInvoker
	MustInvoker
	InvokeWithValidArgs() (
		results []any, processingErr error,
	)
	ValidArgsGetter
	fmt.Stringer
}

type FuncByIndexParameter interface {
	FuncParameter
	InvokeArgs(upTo int) (
		results []any, processingErr error,
	)
}

type FuncByNameParameter interface {
	FuncParameter
	InvokeArgs(names ...string) (
		results []any, processingErr error,
	)
}

type TwoParameter interface {
	OneParameter
	UptoSecondItemGetter
}

type ThreeParameter interface {
	TwoParameter
	UptoThirdItemGetter
}

type FourthParameter interface {
	ThreeParameter
	UptoFourthItemGetter
}

type FifthParameter interface {
	FourthParameter
	UptoFifthItemGetter
}

type SixthParameter interface {
	FifthParameter
	UptoSixthItemGetter
}

type DirectFuncNameGetter interface {
	FuncName() string
}

type CompileStringWithError interface {
	Compile() (string, error)
}

type TemplateReplacer interface {
	ReplaceTemplate(format string, replacerMap map[string]string) string
}
