package fmtcodegentype

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	ranges = [...]string{
		Default:         "Default",
		WithExpectation: "WithExpectation",
		WithFuncError:   "WithFuncError",
	}

	rangesFmt = map[Variant]string{
		Default:         "%d : %s -> %s",
		WithExpectation: "%d : %s(%s) -> %s | %s",
		WithFuncError:   "%d : %s - %s",
	}

	basicEnumImpl = enumimpl.New.BasicByte.DefaultAllCases(
		Default,
		ranges[:],
	)
)
