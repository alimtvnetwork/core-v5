package codestack

import "runtime"

type nameGetter struct{}

func MethodName() (methodName string) {
	_, _, methodName = MethodNamePackageNameUsingStackSkip(defaultInternalSkip)

	return methodName
}

func MethodNameOf(fullName string) (packageName string) {
	_, _, methodName := MethodNamePackageName(
		fullName,
	)

	return methodName
}

func MethodNamePackageName(fullFuncName string) (fullMethodName, packageName, methodName string) {
	if fullFuncName == "" {
		return "", "", ""
	}

	return reflectinternal.MethodNamePackageName(fullFuncName)
}

func MethodNamePackageNameUsingStackSkip(stackSkipIndex int) (fullMethodName, packageName, methodName string) {
	pc, _, _, _ := runtime.Caller(stackSkipIndex + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	return MethodNamePackageName(fullFuncName)
}
func MethodNameUsingStackSkip(stackSkipIndex int) (methodName string) {
	_, _, methodName = MethodNamePackageNameUsingStackSkip(
		stackSkipIndex + defaultInternalSkip,
	)

	return methodName
}

func JoinPackageNameWithRelative(
	fullNameExtractPackageName, relativeNamesJoin string,
) (packageName string) {
	_, packageName, _ = MethodNamePackageName(
		fullNameExtractPackageName,
	)

	return packageName + "." + relativeNamesJoin
}
