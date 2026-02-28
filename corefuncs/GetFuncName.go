package corefuncs

import "gitlab.com/auk-go/core/internal/reflectinternal"

func GetFuncName(i any) string {
	return reflectinternal.GetFunc.NameOnly(i)
}
