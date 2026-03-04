package args

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// Five holds five typed positional arguments plus an optional Expect field.
//
// Type parameters T1–T5 represent the types of First through Fifth.
// Use FiveAny (= Five[any, any, any, any, any]) for untyped usage.
//
// Example (typed):
//
//	arg := args.Five[string, int, bool, float64, []byte]{
//	    First:  "hello",
//	    Second: 42,
//	    Third:  true,
//	    Fourth: 3.14,
//	    Fifth:  []byte("data"),
//	}
type Five[T1, T2, T3, T4, T5 any] struct {
	First         T1                       `json:",omitempty"`
	Second        T2                       `json:",omitempty"`
	Third         T3                       `json:",omitempty"`
	Fourth        T4                       `json:",omitempty"`
	Fifth         T5                       `json:",omitempty"`
	Expect        any                      `json:",omitempty"`
	toSlice       []any                    `json:"-"`
	isSliceCached bool                     `json:"-"`
	toString      corestr.SimpleStringOnce `json:"-"`
}

// ArgsCount returns the number of positional argument slots (always 5).
func (it *Five[T1, T2, T3, T4, T5]) ArgsCount() int {
	return 5
}

// FirstItem returns the First argument as any.
func (it *Five[T1, T2, T3, T4, T5]) FirstItem() any {
	return it.First
}

// SecondItem returns the Second argument as any.
func (it *Five[T1, T2, T3, T4, T5]) SecondItem() any {
	return it.Second
}

// ThirdItem returns the Third argument as any.
func (it *Five[T1, T2, T3, T4, T5]) ThirdItem() any {
	return it.Third
}

// FourthItem returns the Fourth argument as any.
func (it *Five[T1, T2, T3, T4, T5]) FourthItem() any {
	return it.Fourth
}

// FifthItem returns the Fifth argument as any.
func (it *Five[T1, T2, T3, T4, T5]) FifthItem() any {
	return it.Fifth
}

// Expected returns the expected value.
func (it *Five[T1, T2, T3, T4, T5]) Expected() any {
	return it.Expect
}

// ArgTwo returns a Two with the first two arguments.
func (it *Five[T1, T2, T3, T4, T5]) ArgTwo() Two[T1, T2] {
	return Two[T1, T2]{
		First:  it.First,
		Second: it.Second,
	}
}

// ArgThree returns a Three with the first three arguments.
func (it *Five[T1, T2, T3, T4, T5]) ArgThree() Three[T1, T2, T3] {
	return Three[T1, T2, T3]{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

// ArgFour returns a Four with the first four arguments.
func (it *Five[T1, T2, T3, T4, T5]) ArgFour() Four[T1, T2, T3, T4] {
	return Four[T1, T2, T3, T4]{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

// HasFirst checks whether the First argument is defined.
func (it *Five[T1, T2, T3, T4, T5]) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

// HasSecond checks whether the Second argument is defined.
func (it *Five[T1, T2, T3, T4, T5]) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

// HasThird checks whether the Third argument is defined.
func (it *Five[T1, T2, T3, T4, T5]) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

// HasFourth checks whether the Fourth argument is defined.
func (it *Five[T1, T2, T3, T4, T5]) HasFourth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fourth)
}

// HasFifth checks whether the Fifth argument is defined.
func (it *Five[T1, T2, T3, T4, T5]) HasFifth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fifth)
}

// HasExpect checks whether the Expect field is defined.
func (it *Five[T1, T2, T3, T4, T5]) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

// ValidArgs returns all defined positional arguments as a slice.
func (it *Five[T1, T2, T3, T4, T5]) ValidArgs() []any {
	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)
	args = appendIfDefined(args, it.Third)
	args = appendIfDefined(args, it.Fourth)
	args = appendIfDefined(args, it.Fifth)

	return args
}

// Args returns positional arguments up to the given count.
func (it *Five[T1, T2, T3, T4, T5]) Args(upTo int) []any {
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

	if upTo >= 5 {
		args = append(args, it.Fifth)
	}

	return args
}

// Slice returns all fields as a cached slice.
func (it *Five[T1, T2, T3, T4, T5]) Slice() []any {
	if it.isSliceCached {
		return it.toSlice
	}

	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)
	args = appendIfDefined(args, it.Third)
	args = appendIfDefined(args, it.Fourth)
	args = appendIfDefined(args, it.Fifth)
	args = appendIfDefined(args, it.Expect)

	it.toSlice = args
	it.isSliceCached = true

	return it.toSlice
}

// GetByIndex safely retrieves an item from the cached slice by index.
func (it *Five[T1, T2, T3, T4, T5]) GetByIndex(index int) any {
	return getByIndex(it.Slice(), index)
}

// String returns a formatted string representation.
func (it Five[T1, T2, T3, T4, T5]) String() string {
	return buildToString(
		"Five",
		it.Slice(),
		&it.toString,
	)
}

// AsFifthParameter returns the Five as a FifthParameter interface.
func (it Five[T1, T2, T3, T4, T5]) AsFifthParameter() FifthParameter {
	return &it
}

// AsArgBaseContractsBinder returns the Five as an ArgBaseContractsBinder interface.
func (it Five[T1, T2, T3, T4, T5]) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
