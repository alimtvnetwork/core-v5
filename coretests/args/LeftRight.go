package args

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// LeftRight is a generic two-item holder with Left/Right semantics,
// providing a semantic alternative to Two for cases where
// the directionality of arguments matters.
//
// Type parameters TLeft and TRight represent the types of the Left and Right fields.
// Use LeftRightAny (= LeftRight[any, any]) for untyped usage.
//
// Example (typed):
//
//	lr := args.LeftRight[string, int]{
//	    Left:   "hello",
//	    Right:  42,
//	    Expect: true,
//	}
//
// Example (untyped):
//
//	lr := args.LeftRightAny{
//	    Left:   someValue,
//	    Right:  anotherValue,
//	    Expect: "expected",
//	}
type LeftRight[TLeft, TRight any] struct {
	Left          TLeft                    `json:",omitempty"`
	Right         TRight                   `json:",omitempty"`
	Expect        any                      `json:",omitempty"`
	toSlice       []any                    `json:"-"`
	isSliceCached bool                     `json:"-"`
	toString      corestr.SimpleStringOnce `json:"-"`
}

// ArgsCount returns the number of positional argument slots (always 2).
func (it *LeftRight[TLeft, TRight]) ArgsCount() int {
	return 2
}

// FirstItem returns the Left field as any.
func (it *LeftRight[TLeft, TRight]) FirstItem() any {
	return it.Left
}

// SecondItem returns the Right field as any.
func (it *LeftRight[TLeft, TRight]) SecondItem() any {
	return it.Right
}

// Expected returns the expected value.
func (it *LeftRight[TLeft, TRight]) Expected() any {
	return it.Expect
}

// ArgTwo returns a TwoFuncAny with Left and Right fields.
func (it *LeftRight[TLeft, TRight]) ArgTwo() TwoFuncAny {
	return TwoFuncAny{
		First:  it.Left,
		Second: it.Right,
	}
}

// HasFirst checks whether the Left field is defined.
func (it *LeftRight[TLeft, TRight]) HasFirst() bool {
	return it != nil &&
		reflectinternal.Is.Defined(it.Left)
}

// HasSecond checks whether the Right field is defined.
func (it *LeftRight[TLeft, TRight]) HasSecond() bool {
	return it != nil &&
		reflectinternal.Is.Defined(it.Right)
}

// HasLeft checks whether the Left field is defined (alias for HasFirst).
func (it *LeftRight[TLeft, TRight]) HasLeft() bool {
	return it != nil &&
		reflectinternal.Is.Defined(it.Left)
}

// HasRight checks whether the Right field is defined (alias for HasSecond).
func (it *LeftRight[TLeft, TRight]) HasRight() bool {
	return it != nil &&
		reflectinternal.Is.Defined(it.Right)
}

// HasExpect checks whether the Expect field is defined.
func (it *LeftRight[TLeft, TRight]) HasExpect() bool {
	return it != nil &&
		reflectinternal.Is.Defined(it.Expect)
}

// ValidArgs returns all defined positional arguments as a slice.
func (it *LeftRight[TLeft, TRight]) ValidArgs() []any {
	var args []any

	args = appendIfDefined(args, it.Left)
	args = appendIfDefined(args, it.Right)

	return args
}

// Args returns positional arguments up to the given count.
func (it *LeftRight[TLeft, TRight]) Args(upTo int) []any {
	var args []any

	if upTo >= 1 {
		args = append(args, it.Left)
	}

	if upTo >= 2 {
		args = append(args, it.Right)
	}

	return args
}

// Slice returns all fields as a cached slice.
func (it *LeftRight[TLeft, TRight]) Slice() []any {
	if it.isSliceCached {
		return it.toSlice
	}

	var args []any

	args = appendIfDefined(args, it.Left)
	args = appendIfDefined(args, it.Right)
	args = appendIfDefined(args, it.Expect)

	it.toSlice = args
	it.isSliceCached = true

	return it.toSlice
}

// GetByIndex safely retrieves an item from the cached slice by index.
func (it *LeftRight[TLeft, TRight]) GetByIndex(index int) any {
	return getByIndex(it.Slice(), index)
}

// String returns a formatted string representation.
func (it *LeftRight[TLeft, TRight]) String() string {
	return buildToString(
		"LeftRight",
		it.Slice(),
		&it.toString,
	)
}

// Clone returns an independent copy of this LeftRight.
func (it *LeftRight[TLeft, TRight]) Clone() LeftRight[TLeft, TRight] {
	return LeftRight[TLeft, TRight]{
		Left:   it.Left,
		Right:  it.Right,
		Expect: it.Expect,
	}
}

// AsTwoParameter returns the LeftRight as a TwoParameter interface.
func (it LeftRight[TLeft, TRight]) AsTwoParameter() TwoParameter {
	return &it
}

// AsArgBaseContractsBinder returns the LeftRight as an ArgBaseContractsBinder interface.
func (it LeftRight[TLeft, TRight]) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
