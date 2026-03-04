package args

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// Three holds three typed positional arguments plus an optional Expect field.
//
// Type parameters T1, T2, T3 represent the types of the First, Second, Third arguments.
// Use ThreeAny (= Three[any, any, any]) for untyped usage.
//
// Example (typed):
//
//	arg := args.Three[string, int, bool]{
//	    First:  "hello",
//	    Second: 42,
//	    Third:  true,
//	}
type Three[T1, T2, T3 any] struct {
	First         T1                       `json:",omitempty"`
	Second        T2                       `json:",omitempty"`
	Third         T3                       `json:",omitempty"`
	Expect        any                      `json:",omitempty"`
	toSlice       []any                    `json:"-"`
	isSliceCached bool                     `json:"-"`
	toString      corestr.SimpleStringOnce `json:"-"`
}

// ArgsCount returns the number of positional argument slots (always 3).
func (it *Three[T1, T2, T3]) ArgsCount() int {
	return 3
}

// FirstItem returns the First argument as any for interface compatibility.
func (it *Three[T1, T2, T3]) FirstItem() any {
	return it.First
}

// SecondItem returns the Second argument as any for interface compatibility.
func (it *Three[T1, T2, T3]) SecondItem() any {
	return it.Second
}

// ThirdItem returns the Third argument as any for interface compatibility.
func (it *Three[T1, T2, T3]) ThirdItem() any {
	return it.Third
}

// Expected returns the expected value.
func (it *Three[T1, T2, T3]) Expected() any {
	return it.Expect
}

// ArgTwo returns a Two with the first two arguments.
func (it *Three[T1, T2, T3]) ArgTwo() Two[T1, T2] {
	return Two[T1, T2]{
		First:  it.First,
		Second: it.Second,
	}
}

// ArgThree returns a copy of this Three.
func (it *Three[T1, T2, T3]) ArgThree() Three[T1, T2, T3] {
	return Three[T1, T2, T3]{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

// HasFirst checks whether the First argument is defined.
func (it *Three[T1, T2, T3]) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

// HasSecond checks whether the Second argument is defined.
func (it *Three[T1, T2, T3]) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

// HasThird checks whether the Third argument is defined.
func (it *Three[T1, T2, T3]) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

// HasExpect checks whether the Expect field is defined.
func (it *Three[T1, T2, T3]) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

// ValidArgs returns all defined positional arguments as a slice.
func (it *Three[T1, T2, T3]) ValidArgs() []any {
	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)
	args = appendIfDefined(args, it.Third)

	return args
}

// Args returns positional arguments up to the given count.
func (it *Three[T1, T2, T3]) Args(upTo int) []any {
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

	return args
}

// Slice returns all fields as a cached slice.
func (it *Three[T1, T2, T3]) Slice() []any {
	if it.isSliceCached {
		return it.toSlice
	}

	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)
	args = appendIfDefined(args, it.Third)
	args = appendIfDefined(args, it.Expect)

	it.toSlice = args
	it.isSliceCached = true

	return it.toSlice
}

// GetByIndex safely retrieves an item from the cached slice by index.
func (it *Three[T1, T2, T3]) GetByIndex(index int) any {
	return getByIndex(it.Slice(), index)
}

// String returns a formatted string representation.
func (it Three[T1, T2, T3]) String() string {
	return buildToString(
		"Three",
		it.Slice(),
		&it.toString,
	)
}

// LeftRight converts to a LeftRight with First as Left, Second as Right.
func (it *Three[T1, T2, T3]) LeftRight() LeftRightAny {
	return LeftRightAny{
		Left:   it.First,
		Right:  it.Second,
		Expect: it.Expect,
	}
}

// AsThreeParameter returns the Three as a ThreeParameter interface.
func (it Three[T1, T2, T3]) AsThreeParameter() ThreeParameter {
	return &it
}

// AsArgBaseContractsBinder returns the Three as an ArgBaseContractsBinder interface.
func (it Three[T1, T2, T3]) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
