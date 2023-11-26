package codestack

import "runtime"

type newCreator struct{}

func NewDefault() Trace {
	return New(defaultInternalSkip)
}

func NewFirst() Trace {
	return New(Skip2)
}

func NewPtr(skipIndex int) *Trace {
	pc, file, line, isOkay := runtime.Caller(skipIndex + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	fullMethodSignature, packageName, methodName := MethodNamePackageName(fullFuncName)

	return &Trace{
		SkipIndex:         skipIndex,
		PackageName:       packageName,
		MethodName:        methodName,
		PackageMethodName: fullMethodSignature,
		FilePath:          file,
		Line:              line,
		IsOkay:            isOkay,
	}
}
