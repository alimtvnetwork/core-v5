package codegentype

import (
	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
)

var (
	ranges = [...]string{
		Simple:           "Simple",
		MultipleArranges: "MultipleArranges",
	}

	basicEnumImpl = enumimpl.New.BasicByte.DefaultAllCases(
		Simple,
		ranges[:],
	)
)
