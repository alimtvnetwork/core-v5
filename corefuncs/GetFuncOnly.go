package corefuncs

import "gitlab.com/auk-go/core/internal/reflectinternal"

func GetFuncOnly(i interface{}) string {
	return reflectinternal.GetFuncOnly(i)
}
