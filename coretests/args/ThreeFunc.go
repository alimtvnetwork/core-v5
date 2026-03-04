package args

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// ThreeFunc holds three typed positional arguments plus a WorkFunc for
// dynamic function invocation and an optional Expect field.
//
// Type parameters T1, T2, T3 represent the types of First, Second, Third.
// Use ThreeFuncAny (= ThreeFunc[any, any, any]) for untyped usage.
//
// Example (typed):
//
//	tc := args.ThreeFunc[string, int, bool]{
//	    First:    "hello",
//	    Second:   42,
//	    Third:    true,
//	    WorkFunc: myFunc,
//	    Expect:   "expected",
//	}
type ThreeFunc[T1, T2, T3 any] struct {
	First         T1                       `json:",omitempty"`
	Second        T2                       `json:",omitempty"`
	Third         T3                       `json:",omitempty"`
	WorkFunc      any                      `json:"-"`
	Expect        any                      `json:",omitempty"`
	toSlice       []any                    `json:"-"`
	isSliceCached bool                     `json:"-"`
	toString      corestr.SimpleStringOnce `json:"-"`
}

// GetWorkFunc returns the wrapped function value.
func (it *ThreeFunc[T1, T2, T3]) GetWorkFunc() any {
	return it.WorkFunc
}

// ArgsCount returns the number of positional argument slots (always 3).
func (it *ThreeFunc[T1, T2, T3]) ArgsCount() int {
	return 3
}

// FirstItem returns the First argument as any.
func (it *ThreeFunc[T1, T2, T3]) FirstItem() any {
	return it.First
}

// SecondItem returns the Second argument as any.
func (it *ThreeFunc[T1, T2, T3]) SecondItem() any {
	return it.Second
}

// ThirdItem returns the Third argument as any.
func (it *ThreeFunc[T1, T2, T3]) ThirdItem() any {
	return it.Third
}

// Expected returns the expected value.
func (it *ThreeFunc[T1, T2, T3]) Expected() any {
	return it.Expect
}

// ArgTwo returns a TwoFunc with the first two arguments.
func (it *ThreeFunc[T1, T2, T3]) ArgTwo() TwoFunc[T1, T2] {
	return TwoFunc[T1, T2]{
		First:  it.First,
		Second: it.Second,
	}
}

// ArgThree returns a copy of this ThreeFunc with positional args only.
func (it *ThreeFunc[T1, T2, T3]) ArgThree() ThreeFunc[T1, T2, T3] {
	return ThreeFunc[T1, T2, T3]{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

// HasFirst checks whether the First argument is defined.
func (it *ThreeFunc[T1, T2, T3]) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

// HasSecond checks whether the Second argument is defined.
func (it *ThreeFunc[T1, T2, T3]) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

// HasThird checks whether the Third argument is defined.
func (it *ThreeFunc[T1, T2, T3]) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

// HasFunc checks whether the WorkFunc is defined.
func (it *ThreeFunc[T1, T2, T3]) HasFunc() bool {
	return it != nil && reflectinternal.Is.Defined(it.WorkFunc)
}

// HasExpect checks whether the Expect field is defined.
func (it *ThreeFunc[T1, T2, T3]) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

// GetFuncName returns the short name of the wrapped function.
func (it *ThreeFunc[T1, T2, T3]) GetFuncName() string {
	return reflectinternal.GetFunc.NameOnly(it.WorkFunc)
}

// FuncWrap wraps the WorkFunc in a FuncWrapAny for reflection-based invocation.
func (it *ThreeFunc[T1, T2, T3]) FuncWrap() *FuncWrapAny {
	return NewFuncWrap.Default(it.WorkFunc)
}

// Invoke dynamically calls the WorkFunc with the given arguments.
func (it *ThreeFunc[T1, T2, T3]) Invoke(args ...any) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(args...)
}

// InvokeMust invokes the WorkFunc, panicking on error.
func (it *ThreeFunc[T1, T2, T3]) InvokeMust(args ...any) []any {
	return invokeMustHelper(it.FuncWrap(), args...)
}

// InvokeWithValidArgs invokes the WorkFunc with all defined positional arguments.
func (it *ThreeFunc[T1, T2, T3]) InvokeWithValidArgs() (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(it.ValidArgs()...)
}

// InvokeArgs invokes the WorkFunc with positional arguments up to the given count.
func (it *ThreeFunc[T1, T2, T3]) InvokeArgs(upTo int) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(it.Args(upTo)...)
}

// ValidArgs returns all defined positional arguments as a slice.
func (it *ThreeFunc[T1, T2, T3]) ValidArgs() []any {
	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)
	args = appendIfDefined(args, it.Third)

	return args
}

// Args returns positional arguments up to the given count.
func (it *ThreeFunc[T1, T2, T3]) Args(upTo int) []any {
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
func (it *ThreeFunc[T1, T2, T3]) Slice() []any {
	if it.isSliceCached {
		return it.toSlice
	}

	var args []any

	args = appendIfDefined(args, it.First)
	args = appendIfDefined(args, it.Second)
	args = appendIfDefined(args, it.Third)

	if it.HasFunc() {
		args = append(args, it.GetFuncName())
	}

	args = appendIfDefined(args, it.Expect)

	it.toSlice = args
	it.isSliceCached = true

	return it.toSlice
}

// GetByIndex safely retrieves an item from the cached slice by index.
func (it *ThreeFunc[T1, T2, T3]) GetByIndex(index int) any {
	return getByIndex(it.Slice(), index)
}

// String returns a formatted string representation.
func (it ThreeFunc[T1, T2, T3]) String() string {
	return buildToString(
		"ThreeFunc",
		it.Slice(),
		&it.toString,
	)
}

// LeftRight converts to a LeftRight with First as Left, Second as Right.
func (it *ThreeFunc[T1, T2, T3]) LeftRight() LeftRight {
	return LeftRight{
		Left:   it.First,
		Right:  it.Second,
		Expect: it.Expect,
	}
}

// AsThreeFuncParameter returns the ThreeFunc as a ThreeFuncParameter interface.
func (it ThreeFunc[T1, T2, T3]) AsThreeFuncParameter() ThreeFuncParameter {
	return &it
}

// AsArgFuncContractsBinder returns the ThreeFunc as an ArgFuncContractsBinder interface.
func (it ThreeFunc[T1, T2, T3]) AsArgFuncContractsBinder() ArgFuncContractsBinder {
	return &it
}

// AsArgBaseContractsBinder returns the ThreeFunc as an ArgBaseContractsBinder interface.
func (it ThreeFunc[T1, T2, T3]) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
