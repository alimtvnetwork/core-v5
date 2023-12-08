package codegentype

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

var (
	ranges = map[byte]string{
		Invalid.ValueByte():          "Invalid",
		SimpleType.ValueByte():       "SimpleType",
		MultipleArranges.ValueByte(): "MultipleArranges",
	}

	basicEnumImpl = enumimpl.New.BasicByte.CreateUsingMap(
		reflectinternal.TypeName(Invalid),
		ranges,
	)
)
