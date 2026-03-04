package args

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// FiveFunc holds five typed positional arguments plus a WorkFunc for
// dynamic function invocation and an optional Expect field.
//
// Type parameters T1–T5 represent the types of First through Fifth.
// Use FiveFuncAny (= FiveFunc[any, any, any, any, any]) for untyped usage.
type FiveFunc[T1, T2, T3, T4, T5 any] struct {
	First         T1                       `json:",omitempty"`
	Second        T2                       `json:",omitempty"`
	Third         T3                       `json:",omitempty"`
	Fourth        T4                       `json:",omitempty"`
	Fifth         T5                       `json:",omitempty"`
	WorkFunc      any                      `json:"-"`
	Expect        any                      `json:",omitempty"`
	toSlice       []any                    `json:"-"`
	isSliceCached bool                     `json:"-"`
	toString      corestr.SimpleStringOnce `json:"-"`
}

// GetWorkFunc returns the wrapped function value.
func (it *FiveFunc[T1, T2, T3, T4, T5]) GetWorkFunc() any {
	return it.WorkFunc
}

// ArgsCount returns the number of positional argument slots (always 5).
func (it *FiveFunc[T1, T2, T3, T4, T5]) ArgsCount() int {
	return 5
}

// FirstItem returns the First argument as any.
func (it *FiveFunc[T1, T2, T3, T4, T5]) FirstItem() any {
	return it.First
}

// SecondItem returns the Second argument as any.
func (it *FiveFunc[T1, T2, T3, T4, T5]) SecondItem() any {
	return it.Second
}

// ThirdItem returns the Third argument as any.
func (it *FiveFunc[T1, T2, T3, T4, T5]) ThirdItem() any {
	return it.Third
}

// FourthItem returns the Fourth argument as any.
func (it *FiveFunc[T1, T2, T3, T4, T5]) FourthItem() any {
	return it.Fourth
}

// FifthItem returns the Fifth argument as any.
func (it *FiveFunc[T1, T2, T3, T4, T5]) FifthItem() any {
	return it.Fifth
}

// Expected returns the expected value.
func (it *FiveFunc[T1, T2, T3, T4, T5]) Expected() any {
	return it.Expect
}

// ArgTwo returns a TwoFunc with the first two arguments.
func (it *FiveFunc[T1, T2, T3, T4, T5]) ArgTwo() TwoFunc[T1, T2] {
	return TwoFunc[T1, T2]{
		First:  it.First,
		Second: it.Second,
	}
}

// ArgThree returns a ThreeFunc with the first three arguments.
func (it *FiveFunc[T1, T2, T3, T4, T5]) ArgThree() ThreeFunc[T1, T2, T3] {
	return ThreeFunc[T1, T2, T3]{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

// ArgFour returns a FourFunc with the first four arguments.
func (it *FiveFunc[T1, T2, T3, T4, T5]) ArgFour() FourFunc[T1, T2, T3, T4] {
	return FourFunc[T1, T2, T3, T4]{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

// HasFirst checks whether the First argument is defined.
func (it *FiveFunc[T1, T2, T3, T4, T5]) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

// HasSecond checks whether the Second argument is defined.
func (it *FiveFunc[T1, T2, T3, T4, T5]) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

// HasThird checks whether the Third argument is defined.
func (it *FiveFunc[T1, T2, T3, T4, T5]) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

// HasFourth checks whether the Fourth argument is defined.
func (it *FiveFunc[T1, T2, T3, T4, T5]) HasFourth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fourth)
}

// HasFifth checks whether the Fifth argument is defined.
func (it *FiveFunc[T1, T2, T3, T4, T5]) HasFifth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fifth)
}

// HasFunc checks whether the WorkFunc is defined.
func (it *FiveFunc[T1, T2, T3, T4, T5]) HasFunc() bool {
	return it != nil && reflectinternal.Is.Defined(it.WorkFunc)
}

// HasExpect checks whether the Expect field is defined.
func (it *FiveFunc[T1, T2, T3, T4, T5]) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

// GetFuncName returns the short name of the wrapped function.
func (it *FiveFunc[T1, T2, T3, T4, T5]) GetFuncName() string {
	return reflectinternal.GetFunc.NameOnly(it.WorkFunc)
}

// FuncWrap wraps the WorkFunc in a FuncWrapAny for reflection-based invocation.
func (it *FiveFunc[T1, T2, T3, T4, T5]) FuncWrap() *FuncWrapAny {
	return NewFuncWrap.Default(it.WorkFunc)
}

// Invoke dynamically calls the WorkFunc with the given arguments.
func (it *FiveFunc[T1, T2, T3, T4, T5]) Invoke(args ...any) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(args...)
}

// InvokeMust invokes the WorkFunc, panicking on error.
func (it *FiveFunc[T1, T2, T3, T4, T5]) InvokeMust(args ...any) []any {
	return invokeMustHelper(it.FuncWrap(), args...)
}

// InvokeWithValidArgs invokes the WorkFunc with all defined positional arguments.
func (it *FiveFunc[T1, T2, T3, T4, T5]) InvokeWithValidArgs() (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(it.ValidArgs()...)
}

// InvokeArgs invokes the WorkFunc with positional arguments up to the given count.
func (it *FiveFunc[T1, T2, T3, T4, T5]) InvokeArgs(upTo int) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(it.Args(upTo)...)
}

// ValidArgs returns all defined positional arguments as a slice.
func (it *FiveFunc[T1, T2, T3, T4, T5]) ValidArgs() []any {
	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)
	args = appendIfDefined(args, it.Third)
	args = appendIfDefined(args, it.Fourth)
	args = appendIfDefined(args, it.Fifth)

	return args
}

// Args returns positional arguments up to the given count.
func (it *FiveFunc[T1, T2, T3, T4, T5]) Args(upTo int) []any {
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
func (it *FiveFunc[T1, T2, T3, T4, T5]) Slice() []any {
	if it.isSliceCached {
		return it.toSlice
	}

	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)
	args = appendIfDefined(args, it.Third)
	args = appendIfDefined(args, it.Fourth)
	args = appendIfDefined(args, it.Fifth)

	if it.HasFunc() {
		args = append(args, it.GetFuncName())
	}

	args = appendIfDefined(args, it.Expect)

	it.toSlice = args
	it.isSliceCached = true

	return it.toSlice
}

// GetByIndex safely retrieves an item from the cached slice by index.
func (it *FiveFunc[T1, T2, T3, T4, T5]) GetByIndex(index int) any {
	return getByIndex(it.Slice(), index)
}

// String returns a formatted string representation.
func (it FiveFunc[T1, T2, T3, T4, T5]) String() string {
	return buildToString(
		"FiveFunc",
		it.Slice(),
		&it.toString,
	)
}

// AsFifthFuncParameter returns the FiveFunc as a FifthFuncParameter interface.
func (it FiveFunc[T1, T2, T3, T4, T5]) AsFifthFuncParameter() FifthFuncParameter {
	return &it
}

// AsArgFuncContractsBinder returns the FiveFunc as an ArgFuncContractsBinder interface.
func (it FiveFunc[T1, T2, T3, T4, T5]) AsArgFuncContractsBinder() ArgFuncContractsBinder {
	return &it
}

// AsArgBaseContractsBinder returns the FiveFunc as an ArgBaseContractsBinder interface.
func (it FiveFunc[T1, T2, T3, T4, T5]) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
