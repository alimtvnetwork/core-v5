package args

import (
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// Four holds four typed positional arguments plus an optional Expect field.
//
// Type parameters T1, T2, T3, T4 represent the types of First through Fourth.
// Use FourAny (= Four[any, any, any, any]) for untyped usage.
//
// Example (typed):
//
//	arg := args.Four[string, int, bool, float64]{
//	    First:  "hello",
//	    Second: 42,
//	    Third:  true,
//	    Fourth: 3.14,
//	}
type Four[T1, T2, T3, T4 any] struct {
	First         T1                       `json:",omitempty"`
	Second        T2                       `json:",omitempty"`
	Third         T3                       `json:",omitempty"`
	Fourth        T4                       `json:",omitempty"`
	Expect        any                      `json:",omitempty"`
	toSlice       []any                    `json:"-"`
	isSliceCached bool                     `json:"-"`
	toString      corestr.SimpleStringOnce `json:"-"`
}

// ArgsCount returns the number of positional argument slots (always 4).
func (it *Four[T1, T2, T3, T4]) ArgsCount() int {
	return 4
}

// FirstItem returns the First argument as any.
func (it *Four[T1, T2, T3, T4]) FirstItem() any {
	return it.First
}

// SecondItem returns the Second argument as any.
func (it *Four[T1, T2, T3, T4]) SecondItem() any {
	return it.Second
}

// ThirdItem returns the Third argument as any.
func (it *Four[T1, T2, T3, T4]) ThirdItem() any {
	return it.Third
}

// FourthItem returns the Fourth argument as any.
func (it *Four[T1, T2, T3, T4]) FourthItem() any {
	return it.Fourth
}

// Expected returns the expected value.
func (it *Four[T1, T2, T3, T4]) Expected() any {
	return it.Expect
}

// ArgTwo returns a Two with the first two arguments.
func (it *Four[T1, T2, T3, T4]) ArgTwo() Two[T1, T2] {
	return Two[T1, T2]{
		First:  it.First,
		Second: it.Second,
	}
}

// ArgThree returns a Three with the first three arguments.
func (it *Four[T1, T2, T3, T4]) ArgThree() Three[T1, T2, T3] {
	return Three[T1, T2, T3]{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

// HasFirst checks whether the First argument is defined.
func (it *Four[T1, T2, T3, T4]) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

// HasSecond checks whether the Second argument is defined.
func (it *Four[T1, T2, T3, T4]) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

// HasThird checks whether the Third argument is defined.
func (it *Four[T1, T2, T3, T4]) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

// HasFourth checks whether the Fourth argument is defined.
func (it *Four[T1, T2, T3, T4]) HasFourth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fourth)
}

// HasExpect checks whether the Expect field is defined.
func (it *Four[T1, T2, T3, T4]) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

// ValidArgs returns all defined positional arguments as a slice.
func (it *Four[T1, T2, T3, T4]) ValidArgs() []any {
	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)
	args = appendIfDefined(args, it.Third)
	args = appendIfDefined(args, it.Fourth)

	return args
}

// Args returns positional arguments up to the given count.
func (it *Four[T1, T2, T3, T4]) Args(upTo int) []any {
	var args []any

	if upTo >= 1 {
		args = append(args, it.First)
	}

	if upTo >= 2 {
		args = append(args, it.Second)
	}

	if upTo >= 3 {
		args = append(args, it.Third)
	}

	if upTo >= 4 {
		args = append(args, it.Fourth)
	}

	return args
}

// Slice returns all fields as a cached slice.
func (it *Four[T1, T2, T3, T4]) Slice() []any {
	if it.isSliceCached {
		return it.toSlice
	}

	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)
	args = appendIfDefined(args, it.Third)
	args = appendIfDefined(args, it.Fourth)
	args = appendIfDefined(args, it.Expect)

	it.toSlice = args
	it.isSliceCached = true

	return it.toSlice
}

// GetByIndex safely retrieves an item from the cached slice by index.
func (it *Four[T1, T2, T3, T4]) GetByIndex(index int) any {
	return getByIndex(it.Slice(), index)
}

// String returns a formatted string representation.
func (it Four[T1, T2, T3, T4]) String() string {
	return buildToString(
		"Four",
		it.Slice(),
		&it.toString,
	)
}

// AsFourParameter returns the Four as a FourParameter interface.
func (it Four[T1, T2, T3, T4]) AsFourParameter() FourParameter {
	return &it
}

// AsArgBaseContractsBinder returns the Four as an ArgBaseContractsBinder interface.
func (it Four[T1, T2, T3, T4]) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
