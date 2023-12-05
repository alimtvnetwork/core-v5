package coreinterface

import "reflect"

type FuncNameGetter interface {
	FuncName() string
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

type FuncWrapContracts interface {
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

type FuncWrapContracts interface {
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
