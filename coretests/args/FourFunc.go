package args

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// FourFunc holds four typed positional arguments plus a WorkFunc for
// dynamic function invocation and an optional Expect field.
//
// Type parameters T1–T4 represent the types of First through Fourth.
// Use FourFuncAny (= FourFunc[any, any, any, any]) for untyped usage.
type FourFunc[T1, T2, T3, T4 any] struct {
	First         T1                       `json:",omitempty"`
	Second        T2                       `json:",omitempty"`
	Third         T3                       `json:",omitempty"`
	Fourth        T4                       `json:",omitempty"`
	WorkFunc      any                      `json:"-"`
	Expect        any                      `json:",omitempty"`
	toSlice       []any                    `json:"-"`
	isSliceCached bool                     `json:"-"`
	toString      corestr.SimpleStringOnce `json:"-"`
}

// GetWorkFunc returns the wrapped function value.
func (it *FourFunc[T1, T2, T3, T4]) GetWorkFunc() any {
	return it.WorkFunc
}

// ArgsCount returns the number of positional argument slots (always 4).
func (it *FourFunc[T1, T2, T3, T4]) ArgsCount() int {
	return 4
}

// FirstItem returns the First argument as any.
func (it *FourFunc[T1, T2, T3, T4]) FirstItem() any {
	return it.First
}

// SecondItem returns the Second argument as any.
func (it *FourFunc[T1, T2, T3, T4]) SecondItem() any {
	return it.Second
}

// ThirdItem returns the Third argument as any.
func (it *FourFunc[T1, T2, T3, T4]) ThirdItem() any {
	return it.Third
}

// FourthItem returns the Fourth argument as any.
func (it *FourFunc[T1, T2, T3, T4]) FourthItem() any {
	return it.Fourth
}

// Expected returns the expected value.
func (it *FourFunc[T1, T2, T3, T4]) Expected() any {
	return it.Expect
}

// ArgTwo returns a TwoFunc with the first two arguments.
func (it *FourFunc[T1, T2, T3, T4]) ArgTwo() TwoFunc[T1, T2] {
	return TwoFunc[T1, T2]{
		First:  it.First,
		Second: it.Second,
	}
}

// ArgThree returns a ThreeFunc with the first three arguments.
func (it *FourFunc[T1, T2, T3, T4]) ArgThree() ThreeFunc[T1, T2, T3] {
	return ThreeFunc[T1, T2, T3]{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

// ArgFour returns a copy with all four positional args.
func (it *FourFunc[T1, T2, T3, T4]) ArgFour() FourFunc[T1, T2, T3, T4] {
	return FourFunc[T1, T2, T3, T4]{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

// HasFirst checks whether the First argument is defined.
func (it *FourFunc[T1, T2, T3, T4]) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

// HasSecond checks whether the Second argument is defined.
func (it *FourFunc[T1, T2, T3, T4]) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

// HasThird checks whether the Third argument is defined.
func (it *FourFunc[T1, T2, T3, T4]) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

// HasFourth checks whether the Fourth argument is defined.
func (it *FourFunc[T1, T2, T3, T4]) HasFourth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fourth)
}

// HasFunc checks whether the WorkFunc is defined.
func (it *FourFunc[T1, T2, T3, T4]) HasFunc() bool {
	return it != nil && reflectinternal.Is.Defined(it.WorkFunc)
}

// HasExpect checks whether the Expect field is defined.
func (it *FourFunc[T1, T2, T3, T4]) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

// GetFuncName returns the short name of the wrapped function.
func (it *FourFunc[T1, T2, T3, T4]) GetFuncName() string {
	return reflectinternal.GetFunc.NameOnly(it.WorkFunc)
}

// FuncWrap wraps the WorkFunc in a FuncWrapAny for reflection-based invocation.
func (it *FourFunc[T1, T2, T3, T4]) FuncWrap() *FuncWrapAny {
	return NewFuncWrap.Default(it.WorkFunc)
}

// Invoke dynamically calls the WorkFunc with the given arguments.
func (it *FourFunc[T1, T2, T3, T4]) Invoke(args ...any) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(args...)
}

// InvokeMust invokes the WorkFunc, panicking on error.
func (it *FourFunc[T1, T2, T3, T4]) InvokeMust(args ...any) []any {
	return invokeMustHelper(it.FuncWrap(), args...)
}

// InvokeWithValidArgs invokes the WorkFunc with all defined positional arguments.
func (it *FourFunc[T1, T2, T3, T4]) InvokeWithValidArgs() (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(it.ValidArgs()...)
}

// InvokeArgs invokes the WorkFunc with positional arguments up to the given count.
func (it *FourFunc[T1, T2, T3, T4]) InvokeArgs(upTo int) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(it.Args(upTo)...)
}

// ValidArgs returns all defined positional arguments as a slice.
func (it *FourFunc[T1, T2, T3, T4]) ValidArgs() []any {
	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)
	args = appendIfDefined(args, it.Third)
	args = appendIfDefined(args, it.Fourth)

	return args
}

// Args returns positional arguments up to the given count.
func (it *FourFunc[T1, T2, T3, T4]) Args(upTo int) []any {
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
func (it *FourFunc[T1, T2, T3, T4]) Slice() []any {
	if it.isSliceCached {
		return it.toSlice
	}

	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)
	args = appendIfDefined(args, it.Third)
	args = appendIfDefined(args, it.Fourth)

	if it.HasFunc() {
		args = append(args, it.GetFuncName())
	}

	args = appendIfDefined(args, it.Expect)

	it.toSlice = args
	it.isSliceCached = true

	return it.toSlice
}

// GetByIndex safely retrieves an item from the cached slice by index.
func (it *FourFunc[T1, T2, T3, T4]) GetByIndex(index int) any {
	return getByIndex(it.Slice(), index)
}

// String returns a formatted string representation.
func (it FourFunc[T1, T2, T3, T4]) String() string {
	return buildToString(
		"FourFunc",
		it.Slice(),
		&it.toString,
	)
}

// AsFourFuncParameter returns the FourFunc as a FourFuncParameter interface.
func (it FourFunc[T1, T2, T3, T4]) AsFourFuncParameter() FourFuncParameter {
	return &it
}

// AsArgFuncContractsBinder returns the FourFunc as an ArgFuncContractsBinder interface.
func (it FourFunc[T1, T2, T3, T4]) AsArgFuncContractsBinder() ArgFuncContractsBinder {
	return &it
}

// AsArgBaseContractsBinder returns the FourFunc as an ArgBaseContractsBinder interface.
func (it FourFunc[T1, T2, T3, T4]) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
