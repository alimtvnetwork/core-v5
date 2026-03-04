package args

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// One holds a single typed positional argument plus an optional Expect field.
//
// Type parameter T1 represents the type of the First argument.
// Use OneAny (= One[any]) for untyped usage.
//
// Example (typed):
//
//	arg := args.One[string]{
//	    First:  "hello",
//	    Expect: 42,
//	}
//
// Example (untyped):
//
//	arg := args.OneAny{First: "hello", Expect: 42}
type One[T1 any] struct {
	First         T1                       `json:",omitempty"`
	Expect        any                      `json:",omitempty"`
	toSlice       []any                    `json:"-"`
	isSliceCached bool                     `json:"-"`
	toString      corestr.SimpleStringOnce `json:"-"`
}

// FirstItem returns the First argument as any for interface compatibility.
func (it *One[T1]) FirstItem() any {
	return it.First
}

// Expected returns the expected value.
func (it *One[T1]) Expected() any {
	return it.Expect
}

// ArgTwo returns a copy as an equivalent One (identity downcast).
func (it *One[T1]) ArgTwo() One[T1] {
	return One[T1]{
		First:  it.First,
		Expect: it.Expect,
	}
}

// HasFirst checks whether the First argument is defined (non-nil, non-zero).
func (it *One[T1]) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

// HasExpect checks whether the Expect field is defined.
func (it *One[T1]) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

// ValidArgs returns all defined positional arguments as a slice.
func (it *One[T1]) ValidArgs() []any {
	var args []any

	args = appendIfDefined(args, it.First)

	return args
}

// Args returns positional arguments up to the given count.
func (it *One[T1]) Args(upTo int) []any {
	var args []any

	if upTo >= 1 {
		args = append(args, it.First)
	}

	return args
}

// ArgsCount returns the number of positional argument slots (always 1).
func (it *One[T1]) ArgsCount() int {
	return 1
}

// Slice returns all fields (First + Expect) as a cached slice.
func (it *One[T1]) Slice() []any {
	if it.isSliceCached {
		return it.toSlice
	}

	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Expect)

	it.toSlice = args
	it.isSliceCached = true

	return it.toSlice
}

// GetByIndex safely retrieves an item from the cached slice by index.
func (it *One[T1]) GetByIndex(index int) any {
	return getByIndex(it.Slice(), index)
}

// String returns a formatted string representation of the One instance.
func (it One[T1]) String() string {
	return buildToString(
		"One",
		it.Slice(),
		&it.toString,
	)
}

// LeftRight converts the One to a LeftRight with First as Left.
func (it *One[T1]) LeftRight() LeftRight {
	return LeftRight{
		Left:   it.First,
		Expect: it.Expect,
	}
}

// AsOneParameter returns the One as a OneParameter interface.
func (it One[T1]) AsOneParameter() OneParameter {
	return &it
}

// AsArgBaseContractsBinder returns the One as an ArgBaseContractsBinder interface.
func (it One[T1]) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
