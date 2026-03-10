package args

import (
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// Six holds six typed positional arguments plus an optional Expect field.
//
// Type parameters T1–T6 represent the types of First through Sixth.
// Use SixAny (= Six[any, any, any, any, any, any]) for untyped usage.
//
// Example (typed):
//
//	arg := args.Six[string, int, bool, float64, []byte, rune]{
//	    First:  "hello",
//	    Second: 42,
//	    Third:  true,
//	    Fourth: 3.14,
//	    Fifth:  []byte("data"),
//	    Sixth:  'x',
//	}
type Six[T1, T2, T3, T4, T5, T6 any] struct {
	First         T1                       `json:",omitempty"`
	Second        T2                       `json:",omitempty"`
	Third         T3                       `json:",omitempty"`
	Fourth        T4                       `json:",omitempty"`
	Fifth         T5                       `json:",omitempty"`
	Sixth         T6                       `json:",omitempty"`
	Expect        any                      `json:",omitempty"`
	toSlice       []any                    `json:"-"`
	isSliceCached bool                     `json:"-"`
	toString      corestr.SimpleStringOnce `json:"-"`
}

// ArgsCount returns the number of positional argument slots (always 6).
func (it *Six[T1, T2, T3, T4, T5, T6]) ArgsCount() int {
	return 6
}

// FirstItem returns the First argument as any.
func (it *Six[T1, T2, T3, T4, T5, T6]) FirstItem() any {
	return it.First
}

// SecondItem returns the Second argument as any.
func (it *Six[T1, T2, T3, T4, T5, T6]) SecondItem() any {
	return it.Second
}

// ThirdItem returns the Third argument as any.
func (it *Six[T1, T2, T3, T4, T5, T6]) ThirdItem() any {
	return it.Third
}

// FourthItem returns the Fourth argument as any.
func (it *Six[T1, T2, T3, T4, T5, T6]) FourthItem() any {
	return it.Fourth
}

// FifthItem returns the Fifth argument as any.
func (it *Six[T1, T2, T3, T4, T5, T6]) FifthItem() any {
	return it.Fifth
}

// SixthItem returns the Sixth argument as any.
func (it *Six[T1, T2, T3, T4, T5, T6]) SixthItem() any {
	return it.Sixth
}

// Expected returns the expected value.
func (it *Six[T1, T2, T3, T4, T5, T6]) Expected() any {
	return it.Expect
}

// ArgTwo returns a Two with the first two arguments.
func (it *Six[T1, T2, T3, T4, T5, T6]) ArgTwo() Two[T1, T2] {
	return Two[T1, T2]{
		First:  it.First,
		Second: it.Second,
	}
}

// ArgThree returns a Three with the first three arguments.
func (it *Six[T1, T2, T3, T4, T5, T6]) ArgThree() Three[T1, T2, T3] {
	return Three[T1, T2, T3]{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

// ArgFour returns a Four with the first four arguments.
func (it *Six[T1, T2, T3, T4, T5, T6]) ArgFour() Four[T1, T2, T3, T4] {
	return Four[T1, T2, T3, T4]{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

// ArgFive returns a Five with the first five arguments.
func (it *Six[T1, T2, T3, T4, T5, T6]) ArgFive() Five[T1, T2, T3, T4, T5] {
	return Five[T1, T2, T3, T4, T5]{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
		Fifth:  it.Fifth,
	}
}

// HasFirst checks whether the First argument is defined.
func (it *Six[T1, T2, T3, T4, T5, T6]) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

// HasSecond checks whether the Second argument is defined.
func (it *Six[T1, T2, T3, T4, T5, T6]) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

// HasThird checks whether the Third argument is defined.
func (it *Six[T1, T2, T3, T4, T5, T6]) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

// HasFourth checks whether the Fourth argument is defined.
func (it *Six[T1, T2, T3, T4, T5, T6]) HasFourth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fourth)
}

// HasFifth checks whether the Fifth argument is defined.
func (it *Six[T1, T2, T3, T4, T5, T6]) HasFifth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fifth)
}

// HasSixth checks whether the Sixth argument is defined.
func (it *Six[T1, T2, T3, T4, T5, T6]) HasSixth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Sixth)
}

// HasExpect checks whether the Expect field is defined.
func (it *Six[T1, T2, T3, T4, T5, T6]) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

// ValidArgs returns all defined positional arguments as a slice.
func (it *Six[T1, T2, T3, T4, T5, T6]) ValidArgs() []any {
	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)
	args = appendIfDefined(args, it.Third)
	args = appendIfDefined(args, it.Fourth)
	args = appendIfDefined(args, it.Fifth)
	args = appendIfDefined(args, it.Sixth)

	return args
}

// Args returns positional arguments up to the given count.
func (it *Six[T1, T2, T3, T4, T5, T6]) Args(upTo int) []any {
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

	if upTo >= 6 {
		args = append(args, it.Sixth)
	}

	return args
}

// Slice returns all fields as a cached slice.
func (it *Six[T1, T2, T3, T4, T5, T6]) Slice() []any {
	if it.isSliceCached {
		return it.toSlice
	}

	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)
	args = appendIfDefined(args, it.Third)
	args = appendIfDefined(args, it.Fourth)
	args = appendIfDefined(args, it.Fifth)
	args = appendIfDefined(args, it.Sixth)
	args = appendIfDefined(args, it.Expect)

	it.toSlice = args
	it.isSliceCached = true

	return it.toSlice
}

// GetByIndex safely retrieves an item from the cached slice by index.
func (it *Six[T1, T2, T3, T4, T5, T6]) GetByIndex(index int) any {
	return getByIndex(it.Slice(), index)
}

// String returns a formatted string representation.
func (it Six[T1, T2, T3, T4, T5, T6]) String() string {
	return buildToString(
		"Six",
		it.Slice(),
		&it.toString,
	)
}

// AsSixthParameter returns the Six as a SixthParameter interface.
func (it Six[T1, T2, T3, T4, T5, T6]) AsSixthParameter() SixthParameter {
	return &it
}

// AsArgBaseContractsBinder returns the Six as an ArgBaseContractsBinder interface.
func (it Six[T1, T2, T3, T4, T5, T6]) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
