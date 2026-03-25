package corefuncs

import "github.com/alimtvnetwork/core/internal/reflectinternal"

func GetFuncName(i any) string {
	return reflectinternal.GetFunc.NameOnly(i)
}
