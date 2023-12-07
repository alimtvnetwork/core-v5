package args

import "gitlab.com/auk-go/core/coreinterface"

type name interface {
}

type FuncWrapGetter interface {
	FuncWrap() *FuncWrap
}

type OneFuncParameter interface {
	coreinterface.OneParameter
	coreinterface.FuncParameter
	FuncWrapGetter
}

type TowFuncParameter interface {
	coreinterface.TwoParameter
	coreinterface.FuncParameter
	FuncWrapGetter
}

type ThreeFuncParameter interface {
	coreinterface.ThreeParameter
	coreinterface.FuncParameter
	FuncWrapGetter
}

type FourthFuncParameter interface {
	coreinterface.FourthParameter
	coreinterface.FuncParameter
	FuncWrapGetter
}

type FifthFuncParameter interface {
	coreinterface.FifthParameter
	coreinterface.FuncParameter
	FuncWrapGetter
}

type SixthFuncParameter interface {
	coreinterface.SixthParameter
	coreinterface.FuncParameter
	FuncWrapGetter
}
