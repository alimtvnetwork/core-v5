package codestack

import "runtime"

type newCreator struct{}

func (it newCreator) Default() Trace {
	return New(defaultInternalSkip)
}

func (it newCreator) First() Trace {
	return New(Skip2)
}

func (it newCreator) Ptr(skipIndex int) *Trace {
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
