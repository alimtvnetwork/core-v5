package args

import (
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

var (
	rvToInterfacesFunc = reflectinternal.Converter.ReflectValuesToInterfaces
	argsToRvFunc       = reflectinternal.Converter.ArgsToReflectValues
	pascalCaseFunc     = reflectinternal.
				GetFunc.
				PascalFuncName
	NewFuncWrap  = newFuncWrapCreator{}
	FuncDetector = funcDetector{}
	Empty        = emptyCreator{}
)
