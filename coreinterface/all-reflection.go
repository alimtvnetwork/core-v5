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
	PkgName() string
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

type InArgsVerifier interface {
	VerifyInArgs(args []interface{}) (isOkay bool, err error)
}

type OutArgsVerifier interface {
	VerifyOutArgs(args []interface{}) (isOkay bool, err error)
}

type InArgsRvVerifier interface {
	InArgsVerifyRv(args []reflect.Type) (isOkay bool, err error)
}

type OutArgsRvVerifier interface {
	OutArgsVerifyRv(args []reflect.Type) (isOkay bool, err error)
}

type VoidCallNoReturner interface {
	VoidCallNoReturn(
		args ...interface{},
	) (processingErr error)
}

type MustBeValidater interface {
	mustBeValid()
}

type PrivateValidationErrorGetter interface {
	validationError() error
}

type MustInvoker interface {
	InvokeMust(
		args ...interface{},
	) []interface{}
}

type ReflectInvoker interface {
	Invoke(
		args ...interface{},
	) (results []interface{}, processingErr error)
}

type VoidCaller interface {
	VoidCall() ([]interface{}, error)
}

type ValidateMethodArgsGetter interface {
	ValidateMethodArgs(args []interface{}) error
}

type FirstResponseOfInvokeGetter interface {
	GetFirstResponseOfInvoke(
		args ...interface{},
	) (firstResponse interface{}, err error)
}

type InvokeResultOfIndexGetter interface {
	InvokeResultOfIndex(
		index int,
		args ...interface{},
	) (firstResponse interface{}, err error)
}

type InvokeErrorGetter interface {
	InvokeError(
		args ...interface{},
	) (funcErr, processingErr error)
}

type InvokeFirstAndErrorGetter interface {
	InvokeFirstAndError(
		args ...interface{},
	) (firstResponse interface{}, funcErr, processingErr error)
}

type FirstItemGetter interface {
	FirstItem() interface{}
}

type SecondItemGetter interface {
	SecondItem() interface{}
}

type ThirdItemGetter interface {
	ThirdItem() interface{}
}

type FourthItemGetter interface {
	FourthItem() interface{}
}

type FifthItemGetter interface {
	FifthItem() interface{}
}

type SixthItemGetter interface {
	SixthItem() interface{}
}

type ExpectGetter interface {
	Expected() interface{}
}

type ArrangeGetter interface {
	Arrange() interface{}
}

type ActualGetter interface {
	Actual() interface{}
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
	PrivateValidationErrorGetter

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
	Slice() []interface{}
}

type ArgsUptoGetter interface {
	Args(upTo int) []interface{}
}

type ValidArgsGetter interface {
	ValidArgs() []interface{}
}

type HasExpectChecker interface {
	HasExpect() bool
}

type ByIndexGetter interface {
	GetByIndex(index int) interface{}
}

type OneParameter interface {
	FirstItemGetter
	ExpectGetter

	HasFirst() bool
	HasExpectChecker

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
		results []interface{}, processingErr error,
	)
	InvokeArgs(upTo int) (
		results []interface{}, processingErr error,
	)
	ValidArgsGetter
	fmt.Stringer
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
