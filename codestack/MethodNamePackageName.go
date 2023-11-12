package codestack

import (
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func MethodNamePackageName(fullFuncName string) (fullMethodName, packageName, methodName string) {
	if fullFuncName == "" {
		return "", "", ""
	}

	return reflectinternal.MethodNamePackageName(fullFuncName)
}
