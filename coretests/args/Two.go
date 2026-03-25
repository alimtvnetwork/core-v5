package args

import (
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// Two holds two typed positional arguments plus an optional Expect field.
//
// Type parameters T1 and T2 represent the types of the First and Second arguments.
// Use TwoAny (= Two[any, any]) for untyped usage.
//
// Example (typed):
//
//	arg := args.Two[string, int]{
//	    First:  "hello",
//	    Second: 42,
//	    Expect: true,
//	}
type Two[T1, T2 any] struct {
	First         T1                       `json:",omitempty"`
	Second        T2                       `json:",omitempty"`
	Expect        any                      `json:",omitempty"`
	toSlice       []any                    `json:"-"`
	isSliceCached bool                     `json:"-"`
	toString      corestr.SimpleStringOnce `json:"-"`
}

// FirstItem returns the First argument as any for interface compatibility.
func (it *Two[T1, T2]) FirstItem() any {
	return it.First
}

// SecondItem returns the Second argument as any for interface compatibility.
func (it *Two[T1, T2]) SecondItem() any {
	return it.Second
}

// Expected returns the expected value.
func (it *Two[T1, T2]) Expected() any {
	return it.Expect
}

// ArgTwo returns a TwoFunc with First and Second fields.
func (it *Two[T1, T2]) ArgTwo() TwoFuncAny {
	return TwoFuncAny{
		First:  it.First,
		Second: it.Second,
	}
}

// HasFirst checks whether the First argument is defined.
func (it *Two[T1, T2]) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

// HasSecond checks whether the Second argument is defined.
func (it *Two[T1, T2]) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

// HasExpect checks whether the Expect field is defined.
func (it *Two[T1, T2]) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

// ValidArgs returns all defined positional arguments as a slice.
func (it *Two[T1, T2]) ValidArgs() []any {
	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)

	return args
}

// ArgsCount returns the number of positional argument slots (always 2).
func (it *Two[T1, T2]) ArgsCount() int {
	return 2
}

// Args returns positional arguments up to the given count.
func (it *Two[T1, T2]) Args(upTo int) []any {
	var args []any

	if upTo >= 1 {
		args = append(args, it.First)
	}

	if upTo >= 2 {
		args = append(args, it.Second)
	}

	return args
}

// Slice returns all fields as a cached slice.
func (it *Two[T1, T2]) Slice() []any {
	if it.isSliceCached {
		return it.toSlice
	}

	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)
	args = appendIfDefined(args, it.Expect)

	it.toSlice = args
	it.isSliceCached = true

	return it.toSlice
}

// GetByIndex safely retrieves an item from the cached slice by index.
func (it *Two[T1, T2]) GetByIndex(index int) any {
	return getByIndex(it.Slice(), index)
}

// String returns a formatted string representation of the Two instance.
func (it *Two[T1, T2]) String() string {
	return buildToString(
		"Two",
		it.Slice(),
		&it.toString,
	)
}

// LeftRight converts the Two to a LeftRightAny.
func (it *Two[T1, T2]) LeftRight() LeftRightAny {
	return LeftRightAny{
		Left:   it.First,
		Right:  it.Second,
		Expect: it.Expect,
	}
}

// AsTwoParameter returns the Two as a TwoParameter interface.
func (it Two[T1, T2]) AsTwoParameter() TwoParameter {
	return &it
}

// AsArgBaseContractsBinder returns the Two as an ArgBaseContractsBinder interface.
func (it Two[T1, T2]) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
