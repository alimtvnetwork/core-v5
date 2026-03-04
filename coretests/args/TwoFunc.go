package args

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// TwoFunc holds two typed positional arguments plus a WorkFunc for
// dynamic function invocation and an optional Expect field.
//
// Type parameters T1 and T2 represent the types of First and Second.
// Use TwoFuncAny (= TwoFunc[any, any]) for untyped usage.
//
// Example (typed):
//
//	tc := args.TwoFunc[string, int]{
//	    First:    "hello",
//	    Second:   42,
//	    WorkFunc: myFunc,
//	    Expect:   "expected",
//	}
type TwoFunc[T1, T2 any] struct {
	First         T1                       `json:",omitempty"`
	Second        T2                       `json:",omitempty"`
	WorkFunc      any                      `json:"-"`
	Expect        any                      `json:",omitempty"`
	toSlice       []any                    `json:"-"`
	isSliceCached bool                     `json:"-"`
	toString      corestr.SimpleStringOnce `json:"-"`
}

// GetWorkFunc returns the wrapped function value.
func (it *TwoFunc[T1, T2]) GetWorkFunc() any {
	return it.WorkFunc
}

// ArgsCount returns the number of positional argument slots (always 2).
func (it *TwoFunc[T1, T2]) ArgsCount() int {
	return 2
}

// FirstItem returns the First argument as any.
func (it *TwoFunc[T1, T2]) FirstItem() any {
	return it.First
}

// SecondItem returns the Second argument as any.
func (it *TwoFunc[T1, T2]) SecondItem() any {
	return it.Second
}

// Expected returns the expected value.
func (it *TwoFunc[T1, T2]) Expected() any {
	return it.Expect
}

// ArgTwo returns a copy with First and Second only.
func (it *TwoFunc[T1, T2]) ArgTwo() TwoFunc[T1, T2] {
	return TwoFunc[T1, T2]{
		First:  it.First,
		Second: it.Second,
	}
}

// HasFirst checks whether the First argument is defined.
func (it *TwoFunc[T1, T2]) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

// HasSecond checks whether the Second argument is defined.
func (it *TwoFunc[T1, T2]) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

// HasFunc checks whether the WorkFunc is defined.
func (it *TwoFunc[T1, T2]) HasFunc() bool {
	return it != nil && reflectinternal.Is.Defined(it.WorkFunc)
}

// HasExpect checks whether the Expect field is defined.
func (it *TwoFunc[T1, T2]) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

// GetFuncName returns the short name of the wrapped function.
func (it *TwoFunc[T1, T2]) GetFuncName() string {
	return reflectinternal.GetFunc.NameOnly(it.WorkFunc)
}

// FuncWrap wraps the WorkFunc in a FuncWrapAny for reflection-based invocation.
func (it *TwoFunc[T1, T2]) FuncWrap() *FuncWrapAny {
	return NewFuncWrap.Default(it.WorkFunc)
}

// Invoke dynamically calls the WorkFunc with the given arguments.
func (it *TwoFunc[T1, T2]) Invoke(args ...any) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(args...)
}

// InvokeMust invokes the WorkFunc, panicking on error.
func (it *TwoFunc[T1, T2]) InvokeMust(args ...any) []any {
	return invokeMustHelper(it.FuncWrap(), args...)
}

// InvokeWithValidArgs invokes the WorkFunc with all defined positional arguments.
func (it *TwoFunc[T1, T2]) InvokeWithValidArgs() (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(it.ValidArgs()...)
}

// InvokeArgs invokes the WorkFunc with positional arguments up to the given count.
func (it *TwoFunc[T1, T2]) InvokeArgs(upTo int) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(it.Args(upTo)...)
}

// ValidArgs returns all defined positional arguments as a slice.
func (it *TwoFunc[T1, T2]) ValidArgs() []any {
	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)

	return args
}

// Args returns positional arguments up to the given count.
func (it *TwoFunc[T1, T2]) Args(upTo int) []any {
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
func (it *TwoFunc[T1, T2]) Slice() []any {
	if it.isSliceCached {
		return it.toSlice
	}

	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)

	if it.HasFunc() {
		args = append(args, it.GetFuncName())
	}

	args = appendIfDefined(args, it.Expect)

	it.toSlice = args
	it.isSliceCached = true

	return it.toSlice
}

// GetByIndex safely retrieves an item from the cached slice by index.
func (it *TwoFunc[T1, T2]) GetByIndex(index int) any {
	return getByIndex(it.Slice(), index)
}

// String returns a formatted string representation.
func (it TwoFunc[T1, T2]) String() string {
	return buildToString(
		"TwoFunc",
		it.Slice(),
		&it.toString,
	)
}

// LeftRight converts to a LeftRight.
func (it *TwoFunc[T1, T2]) LeftRight() LeftRight {
	return LeftRight{
		Left:   it.First,
		Right:  it.Second,
		Expect: it.Expect,
	}
}

// AsTwoFuncParameter returns the TwoFunc as a TwoFuncParameter interface.
func (it TwoFunc[T1, T2]) AsTwoFuncParameter() TwoFuncParameter {
	return &it
}

// AsArgFuncContractsBinder returns the TwoFunc as an ArgFuncContractsBinder interface.
func (it TwoFunc[T1, T2]) AsArgFuncContractsBinder() ArgFuncContractsBinder {
	return &it
}

// AsArgBaseContractsBinder returns the TwoFunc as an ArgBaseContractsBinder interface.
func (it TwoFunc[T1, T2]) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
