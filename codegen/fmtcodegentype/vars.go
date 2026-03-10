package fmtcodegentype

import (
	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
)

var (
	ranges = [...]string{
		Default:        "Default",
		WithExpect:     "WithExpect",
		WithFuncExpect: "WithFuncExpect",
		WithFuncError:  "WithFuncError",
	}

	rangesFmt = map[Variant]string{
		Default:        "%d : %s -> %s",
		WithExpect:     "%d : %s -> %s | %s",
		WithFuncExpect: "%d : %s(%s) -> %s | %s",
		WithFuncError:  "%d : %s - %s",
	}

	basicEnumImpl = enumimpl.New.BasicByte.DefaultAllCases(
		Default,
		ranges[:],
	)
)
