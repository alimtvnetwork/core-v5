package codestack

import (
	"runtime"
)

func New(skipIndex int) Trace {
	pc, file, line, isOkay := runtime.Caller(skipIndex + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	fullMethodSignature, packageName, methodName := MethodNamePackageName(fullFuncName)

	return Trace{
		SkipIndex:         skipIndex,
		PackageName:       packageName,
		MethodName:        methodName,
		PackageMethodName: fullMethodSignature,
		FileName:          file,
		Line:              line,
		IsOkay:            isOkay,
	}
}
