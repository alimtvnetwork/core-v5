package corefuncs

import (
	"runtime"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func GetFunc(i any) *runtime.Func {
	return reflectinternal.GetFunc(i)
}
