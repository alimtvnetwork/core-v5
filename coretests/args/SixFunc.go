package args

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// SixFunc holds six typed positional arguments plus a WorkFunc for
// dynamic function invocation and an optional Expect field.
//
// Type parameters T1–T6 represent the types of First through Sixth.
// Use SixFuncAny (= SixFunc[any, any, any, any, any, any]) for untyped usage.
type SixFunc[T1, T2, T3, T4, T5, T6 any] struct {
	First         T1                       `json:",omitempty"`
	Second        T2                       `json:",omitempty"`
	Third         T3                       `json:",omitempty"`
	Fourth        T4                       `json:",omitempty"`
	Fifth         T5                       `json:",omitempty"`
	Sixth         T6                       `json:",omitempty"`
	WorkFunc      any                      `json:"-"`
	Expect        any                      `json:",omitempty"`
	toSlice       []any                    `json:"-"`
	isSliceCached bool                     `json:"-"`
	toString      corestr.SimpleStringOnce `json:"-"`
}

// GetWorkFunc returns the wrapped function value.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) GetWorkFunc() any {
	return it.WorkFunc
}

// ArgsCount returns the number of positional argument slots (always 6).
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) ArgsCount() int {
	return 6
}

// FirstItem returns the First argument as any.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) FirstItem() any {
	return it.First
}

// SecondItem returns the Second argument as any.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) SecondItem() any {
	return it.Second
}

// ThirdItem returns the Third argument as any.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) ThirdItem() any {
	return it.Third
}

// FourthItem returns the Fourth argument as any.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) FourthItem() any {
	return it.Fourth
}

// FifthItem returns the Fifth argument as any.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) FifthItem() any {
	return it.Fifth
}

// SixthItem returns the Sixth argument as any.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) SixthItem() any {
	return it.Sixth
}

// Expected returns the expected value.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) Expected() any {
	return it.Expect
}

// ArgTwo returns a TwoFunc with the first two arguments.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) ArgTwo() TwoFunc[T1, T2] {
	return TwoFunc[T1, T2]{
		First:  it.First,
		Second: it.Second,
	}
}

// ArgThree returns a ThreeFunc with the first three arguments.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) ArgThree() ThreeFunc[T1, T2, T3] {
	return ThreeFunc[T1, T2, T3]{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

// ArgFour returns a FourFunc with the first four arguments.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) ArgFour() FourFunc[T1, T2, T3, T4] {
	return FourFunc[T1, T2, T3, T4]{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

// ArgFive returns a FiveFunc with the first five arguments.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) ArgFive() FiveFunc[T1, T2, T3, T4, T5] {
	return FiveFunc[T1, T2, T3, T4, T5]{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
		Fifth:  it.Fifth,
	}
}

// HasFirst checks whether the First argument is defined.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

// HasSecond checks whether the Second argument is defined.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

// HasThird checks whether the Third argument is defined.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

// HasFourth checks whether the Fourth argument is defined.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) HasFourth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fourth)
}

// HasFifth checks whether the Fifth argument is defined.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) HasFifth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fifth)
}

// HasSixth checks whether the Sixth argument is defined.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) HasSixth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Sixth)
}

// HasFunc checks whether the WorkFunc is defined.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) HasFunc() bool {
	return it != nil && reflectinternal.Is.Defined(it.WorkFunc)
}

// HasExpect checks whether the Expect field is defined.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

// GetFuncName returns the short name of the wrapped function.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) GetFuncName() string {
	return reflectinternal.GetFunc.NameOnly(it.WorkFunc)
}

// FuncWrap wraps the WorkFunc in a FuncWrapAny for reflection-based invocation.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) FuncWrap() *FuncWrapAny {
	return NewFuncWrap.Default(it.WorkFunc)
}

// Invoke dynamically calls the WorkFunc with the given arguments.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) Invoke(args ...any) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(args...)
}

// InvokeMust invokes the WorkFunc, panicking on error.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) InvokeMust(args ...any) []any {
	return invokeMustHelper(it.FuncWrap(), args...)
}

// InvokeWithValidArgs invokes the WorkFunc with all defined positional arguments.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) InvokeWithValidArgs() (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(it.ValidArgs()...)
}

// InvokeArgs invokes the WorkFunc with positional arguments up to the given count.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) InvokeArgs(upTo int) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(it.Args(upTo)...)
}

// ValidArgs returns all defined positional arguments as a slice.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) ValidArgs() []any {
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
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) Args(upTo int) []any {
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
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) Slice() []any {
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

	if it.HasFunc() {
		args = append(args, it.GetFuncName())
	}

	args = appendIfDefined(args, it.Expect)

	it.toSlice = args
	it.isSliceCached = true

	return it.toSlice
}

// GetByIndex safely retrieves an item from the cached slice by index.
func (it *SixFunc[T1, T2, T3, T4, T5, T6]) GetByIndex(index int) any {
	return getByIndex(it.Slice(), index)
}

// String returns a formatted string representation.
func (it SixFunc[T1, T2, T3, T4, T5, T6]) String() string {
	return buildToString(
		"SixFunc",
		it.Slice(),
		&it.toString,
	)
}

// AsSixthFuncParameter returns the SixFunc as a SixthFuncParameter interface.
func (it SixFunc[T1, T2, T3, T4, T5, T6]) AsSixthFuncParameter() SixthFuncParameter {
	return &it
}

// AsArgFuncContractsBinder returns the SixFunc as an ArgFuncContractsBinder interface.
func (it SixFunc[T1, T2, T3, T4, T5, T6]) AsArgFuncContractsBinder() ArgFuncContractsBinder {
	return &it
}

// AsArgBaseContractsBinder returns the SixFunc as an ArgBaseContractsBinder interface.
func (it SixFunc[T1, T2, T3, T4, T5, T6]) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
