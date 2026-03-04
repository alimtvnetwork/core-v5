package args

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// LeftRight is a two-item holder with Left/Right semantics,
// providing a semantic alternative to Two for cases where
// the directionality of arguments matters.
type LeftRight struct {
	Left          any                      `json:",omitempty"`
	Right         any                      `json:",omitempty"`
	Expect        any                      `json:",omitempty"`
	toSlice       []any                    `json:"-"`
	isSliceCached bool                     `json:"-"`
	toString      corestr.SimpleStringOnce `json:"-"`
}

// ArgsCount returns the number of positional argument slots (always 2).
func (it *LeftRight) ArgsCount() int {
	return 2
}

// FirstItem returns the Left field as any.
func (it *LeftRight) FirstItem() any {
	return it.Left
}

// SecondItem returns the Right field as any.
func (it *LeftRight) SecondItem() any {
	return it.Right
}

// Expected returns the expected value.
func (it *LeftRight) Expected() any {
	return it.Expect
}

// ArgTwo returns a TwoFuncAny with Left and Right fields.
func (it *LeftRight) ArgTwo() TwoFuncAny {
	return TwoFuncAny{
		First:  it.Left,
		Second: it.Right,
	}
}

// HasFirst checks whether the Left field is defined.
func (it *LeftRight) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.Left)
}

// HasSecond checks whether the Right field is defined.
func (it *LeftRight) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Right)
}

// HasLeft checks whether the Left field is defined (alias for HasFirst).
func (it *LeftRight) HasLeft() bool {
	return it != nil && reflectinternal.Is.Defined(it.Left)
}

// HasRight checks whether the Right field is defined (alias for HasSecond).
func (it *LeftRight) HasRight() bool {
	return it != nil && reflectinternal.Is.Defined(it.Right)
}

// HasExpect checks whether the Expect field is defined.
func (it *LeftRight) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

// ValidArgs returns all defined positional arguments as a slice.
func (it *LeftRight) ValidArgs() []any {
	var args []any

	args = appendIfDefined(args, it.Left)
	args = appendIfDefined(args, it.Right)

	return args
}

// Args returns positional arguments up to the given count.
func (it *LeftRight) Args(upTo int) []any {
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
func (it *LeftRight) Slice() []any {
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
func (it *LeftRight) GetByIndex(index int) any {
	return getByIndex(it.Slice(), index)
}

// String returns a formatted string representation.
func (it *LeftRight) String() string {
	return buildToString(
		"LeftRight",
		it.Slice(),
		&it.toString,
	)
}

// Clone returns an independent copy of this LeftRight.
func (it *LeftRight) Clone() LeftRight {
	return LeftRight{
		Left:   it.Left,
		Right:  it.Right,
		Expect: it.Expect,
	}
}

// AsTwoParameter returns the LeftRight as a TwoParameter interface.
func (it LeftRight) AsTwoParameter() TwoParameter {
	return &it
}

// AsArgBaseContractsBinder returns the LeftRight as an ArgBaseContractsBinder interface.
func (it LeftRight) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
