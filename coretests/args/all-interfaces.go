package args

import "gitlab.com/auk-go/core/coreinterface"

type FuncWrapGetter interface {
	FuncWrap() *FuncWrap
}

type FuncNumber interface {
	coreinterface.FuncByIndexParameter
	FuncWrapGetter
}

type FuncNamer interface {
	coreinterface.FuncByNameParameter
	FuncWrapGetter
}

type OneFuncParameter interface {
	coreinterface.OneParameter
	FuncNumber
}

type TowFuncParameter interface {
	OneFuncParameter
	coreinterface.TwoParameter
	FuncNumber
}

type ThreeFuncParameter interface {
	TowFuncParameter
	coreinterface.ThreeParameter
	FuncNumber
}

type FourthFuncParameter interface {
	ThreeFuncParameter
	coreinterface.FourthParameter
	FuncNumber
}

type FifthFuncParameter interface {
	FourthFuncParameter
	coreinterface.FifthParameter
	FuncNumber
}

type SixthFuncParameter interface {
	FifthFuncParameter
	coreinterface.SixthParameter
	FuncNumber
}

type ArgsMapper interface {
	coreinterface.FirstItemGetter
	coreinterface.ExpectGetter
	HasFirst() bool
	coreinterface.HasExpectChecker
	coreinterface.ValidArgsGetter
	coreinterface.SliceGetter
	coreinterface.ByIndexGetter
	coreinterface.UptoSixthItemGetter

	FuncNamer
}

type FuncWrapper interface {
	coreinterface.FuncWrapContractsBinder
	InvalidError() error
	IsEqual(
		another *FuncWrap,
	) bool
	IsNotEqual(
		another *FuncWrap,
	) bool
}
