package fmtcodegentype

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

var (
	ranges = map[byte]string{
		Default.ValueByte():       "Default",
		WithFunction.ValueByte():  "WithFunction",
		WithFuncError.ValueByte(): "WithFuncError",
	}

	rangesFmt = map[Variant]string{
		Default:       "%d : %s -> %s",
		WithFunction:  "%d : %s(%s) -> %s | %s",
		WithFuncError: "%d : %s - %s",
	}

	basicEnumImpl = enumimpl.New.BasicByte.CreateUsingMap(
		reflectinternal.TypeName(Default),
		ranges,
	)
)
