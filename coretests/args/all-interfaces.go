package args

import (
	"fmt"

	"gitlab.com/auk-go/core/coreinterface"
)

// FuncWrapGetter provides access to a FuncWrapAny instance.
type FuncWrapGetter interface {
	FuncWrap() *FuncWrapAny
}

// FuncNumber combines function access with index-based parameter retrieval.
type FuncNumber interface {
	GetWorkFunc() any
	coreinterface.FuncByIndexParameter
	FuncWrapGetter
}

// FuncNamer combines function access with name-based parameter retrieval.
type FuncNamer interface {
	GetWorkFunc() any
	coreinterface.FuncByNameParameter
	FuncWrapGetter
}

// OneParameter defines the contract for a single-argument holder.
type OneParameter interface {
	ArgBaseContractsBinder
	AsArgBaseContractsBinder
	coreinterface.OneParameter
}

// OneFuncParameter extends OneParameter with function invocation support.
type OneFuncParameter interface {
	ArgFuncContractsBinder
	AsArgFuncContractsBinder
	OneParameter
	FuncNumber
}

// TwoParameter defines the contract for a two-argument holder.
type TwoParameter interface {
	ArgBaseContractsBinder
	OneParameter
	coreinterface.TwoParameter
}

// TwoFuncParameter extends TwoParameter with function invocation support.
type TwoFuncParameter interface {
	OneFuncParameter
	TwoParameter
	FuncNumber
}

// ThreeParameter defines the contract for a three-argument holder.
type ThreeParameter interface {
	TwoParameter
	coreinterface.ThreeParameter
}

// ThreeFuncParameter extends ThreeParameter with function invocation support.
type ThreeFuncParameter interface {
	TwoFuncParameter
	ThreeParameter
	FuncNumber
}

// FourParameter defines the contract for a four-argument holder.
type FourParameter interface {
	ThreeParameter
	coreinterface.FourthParameter
}

// FourFuncParameter extends FourParameter with function invocation support.
type FourFuncParameter interface {
	ThreeFuncParameter
	FourParameter
	FuncNumber
}

// FifthParameter defines the contract for a five-argument holder.
type FifthParameter interface {
	FourParameter
	coreinterface.FifthParameter
}

// FifthFuncParameter extends FifthParameter with function invocation support.
type FifthFuncParameter interface {
	FourFuncParameter
	FifthParameter
	FuncNumber
}

// SixthParameter defines the contract for a six-argument holder.
type SixthParameter interface {
	FifthParameter
	coreinterface.SixthParameter
}

// SixthFuncParameter extends SixthParameter with function invocation support.
type SixthFuncParameter interface {
	FifthFuncParameter
	SixthParameter
	FuncNumber
}

// ArgsMapper provides map-based argument access with function invocation.
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

// FuncWrapper defines the full contract for a FuncWrapAny instance.
type FuncWrapper interface {
	coreinterface.FuncWrapContractsBinder
	InvalidError() error
	IsEqual(
		another *FuncWrapAny,
	) bool
	IsNotEqual(
		another *FuncWrapAny,
	) bool
}

// HasFirstChecker checks whether the first argument is defined.
type HasFirstChecker interface {
	HasFirst() bool
}

// ArgBaseContractsBinder is the core contract for all argument holders.
// It provides item access, validation, slicing, and string formatting.
type ArgBaseContractsBinder interface {
	coreinterface.FirstItemGetter
	coreinterface.ExpectGetter
	HasFirstChecker

	coreinterface.HasExpectChecker
	coreinterface.ValidArgsGetter
	coreinterface.SliceGetter
	coreinterface.ByIndexGetter

	coreinterface.ArgsCountGetter

	fmt.Stringer
}

// ArgFuncContractsBinder extends ArgBaseContractsBinder with function support.
type ArgFuncContractsBinder interface {
	ArgBaseContractsBinder
	FuncNumber
}

// AsArgBaseContractsBinder provides conversion to ArgBaseContractsBinder.
type AsArgBaseContractsBinder interface {
	AsArgBaseContractsBinder() ArgBaseContractsBinder
}

// AsArgFuncContractsBinder provides conversion to ArgFuncContractsBinder.
type AsArgFuncContractsBinder interface {
	AsArgFuncContractsBinder() ArgFuncContractsBinder
}

// ArgFuncNameContractsBinder extends ArgBaseContractsBinder with named function support.
type ArgFuncNameContractsBinder interface {
	ArgBaseContractsBinder
	FuncNamer
}

// AsArgFuncNameContractsBinder provides conversion to ArgFuncNameContractsBinder.
type AsArgFuncNameContractsBinder interface {
	AsArgFuncNameContractsBinder() ArgFuncNameContractsBinder
}

type AsOneFuncParameter interface {
	AsOneFuncParameter() OneFuncParameter
}

type AsTwoFuncParameter interface {
	AsTwoFuncParameter() TwoFuncParameter
}

type AsThreeFuncParameter interface {
	AsThreeFuncParameter() ThreeFuncParameter
}

type AsFourFuncParameter interface {
	AsFourFuncParameter() FourFuncParameter
}

type AsFifthFuncParameter interface {
	AsFifthFuncParameter() FifthFuncParameter
}

type AsSixthFuncParameter interface {
	AsSixthFuncParameter() SixthFuncParameter
}

type AsOneParameter interface {
	AsOneParameter() OneParameter
}

type AsTwoParameter interface {
	AsTwoParameter() TwoParameter
}

type AsThreeParameter interface {
	AsThreeParameter() ThreeParameter
}

type AsFourParameter interface {
	AsFourParameter() FourParameter
}

type AsFifthParameter interface {
	AsFifthParameter() FifthParameter
}

type AsSixthParameter interface {
	AsSixthParameter() SixthParameter
}
