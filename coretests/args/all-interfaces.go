package args

import (
	"fmt"

	"gitlab.com/auk-go/core/coreinterface"
)

type FuncWrapGetter interface {
	FuncWrap() *FuncWrap
}

type FuncNumber interface {
	GetWorkFunc() interface{}
	coreinterface.FuncByIndexParameter
	FuncWrapGetter
}

type FuncNamer interface {
	GetWorkFunc() interface{}
	coreinterface.FuncByNameParameter
	FuncWrapGetter
}

type OneParameter interface {
	ArgBaseContractsBinder
	coreinterface.OneParameter
}

type OneFuncParameter interface {
	ArgFuncContractsBinder
	OneParameter
	FuncNumber
}

type TwoParameter interface {
	ArgBaseContractsBinder
	coreinterface.TwoParameter
}

type TwoFuncParameter interface {
	OneFuncParameter
	TwoParameter
	FuncNumber
}

type ThreeParameter interface {
	TwoParameter
	coreinterface.ThreeParameter
}

type ThreeFuncParameter interface {
	TwoFuncParameter
	ThreeParameter
	FuncNumber
}

type FourthParameter interface {
	ThreeParameter
	coreinterface.FourthParameter
}

type FourthFuncParameter interface {
	ThreeFuncParameter
	FourthParameter
	FuncNumber
}

type FifthParameter interface {
	FourthParameter
	coreinterface.FifthParameter
}

type FifthFuncParameter interface {
	FourthFuncParameter
	FifthParameter
	FuncNumber
}

type SixthParameter interface {
	FifthParameter
	coreinterface.SixthParameter
}

type SixthFuncParameter interface {
	FifthFuncParameter
	SixthParameter
	FuncNumber
}

type ArgsMapper interface {
	ArgBaseContractsBinder

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

type ArgBaseContractsBinder interface {
	coreinterface.FirstItemGetter
	coreinterface.ExpectGetter
	HasFirst() bool

	coreinterface.HasExpectChecker
	coreinterface.ValidArgsGetter
	coreinterface.SliceGetter
	coreinterface.ByIndexGetter

	coreinterface.ArgsCountGetter

	fmt.Stringer
}

type ArgFuncContractsBinder interface {
	ArgBaseContractsBinder
	FuncNumber
}

type ArgFuncNameContractsBinder interface {
	ArgBaseContractsBinder
	FuncNamer
}
