package corefuncs

import (
	"runtime"

	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

func GetFunc(i any) *runtime.Func {
	return reflectinternal.GetFunc.RunTime(i)
}
