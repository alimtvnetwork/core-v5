package args

import (
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

var (
	rvToInterfacesFunc = reflectinternal.Converter.ReflectValuesToInterfaces
	argsToRvFunc       = reflectinternal.Converter.ArgsToReflectValues
	NewFuncWrap        = newFuncWrapCreator{}
)
