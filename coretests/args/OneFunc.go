package args

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// OneFunc holds a single typed positional argument plus a WorkFunc for
// dynamic function invocation and an optional Expect field.
//
// Type parameter T1 represents the type of the First argument.
// The WorkFunc field is always any (for reflection-based invocation).
// Use OneFuncAny (= OneFunc[any]) for untyped usage.
//
// Example (typed):
//
//	tc := args.OneFunc[string]{
//	    First:    "hello",
//	    WorkFunc: strings.ToUpper,
//	    Expect:   "HELLO",
//	}
//	results, err := tc.InvokeWithValidArgs()
type OneFunc[T1 any] struct {
	First         T1                       `json:",omitempty"`
	WorkFunc      any                      `json:"-,omitempty"`
	Expect        any                      `json:",omitempty"`
	toSlice       []any                    `json:"-"`
	isSliceCached bool                     `json:"-"`
	toString      corestr.SimpleStringOnce `json:"-"`
}

// GetWorkFunc returns the wrapped function value.
func (it *OneFunc[T1]) GetWorkFunc() any {
	return it.WorkFunc
}

// FirstItem returns the First argument as any for interface compatibility.
func (it *OneFunc[T1]) FirstItem() any {
	return it.First
}

// Expected returns the expected value.
func (it *OneFunc[T1]) Expected() any {
	return it.Expect
}

// ArgTwo returns a copy of this OneFunc.
func (it *OneFunc[T1]) ArgTwo() OneFunc[T1] {
	return OneFunc[T1]{
		First:    it.First,
		WorkFunc: it.WorkFunc,
		Expect:   it.Expect,
	}
}

// HasFirst checks whether the First argument is defined.
func (it *OneFunc[T1]) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

// HasFunc checks whether the WorkFunc is defined.
func (it *OneFunc[T1]) HasFunc() bool {
	return it != nil && reflectinternal.Is.Defined(it.WorkFunc)
}

// HasExpect checks whether the Expect field is defined.
func (it *OneFunc[T1]) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

// GetFuncName returns the short name of the wrapped function.
func (it *OneFunc[T1]) GetFuncName() string {
	return reflectinternal.GetFunc.NameOnly(it.WorkFunc)
}

// FuncWrap wraps the WorkFunc in a FuncWrapAny for reflection-based invocation.
func (it *OneFunc[T1]) FuncWrap() *FuncWrapAny {
	return NewFuncWrap.Default(it.WorkFunc)
}

// Invoke dynamically calls the WorkFunc with the given arguments.
func (it *OneFunc[T1]) Invoke(args ...any) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(args...)
}

// InvokeMust invokes the WorkFunc, panicking on error.
func (it *OneFunc[T1]) InvokeMust(args ...any) []any {
	return invokeMustHelper(it.FuncWrap(), args...)
}

// InvokeWithValidArgs invokes the WorkFunc with all defined positional arguments.
func (it *OneFunc[T1]) InvokeWithValidArgs() (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(it.ValidArgs()...)
}

// InvokeArgs invokes the WorkFunc with positional arguments up to the given count.
func (it *OneFunc[T1]) InvokeArgs(upTo int) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(it.Args(upTo)...)
}

// ValidArgs returns all defined positional arguments as a slice.
func (it *OneFunc[T1]) ValidArgs() []any {
	var args []any

	args = appendIfDefined(args, it.First)

	return args
}

// ArgsCount returns the number of positional argument slots (always 1).
func (it *OneFunc[T1]) ArgsCount() int {
	return 1
}

// Args returns positional arguments up to the given count.
func (it *OneFunc[T1]) Args(upTo int) []any {
	var args []any

	if upTo >= 1 {
		args = append(args, it.First)
	}

	return args
}

// Slice returns all fields (First + FuncName + Expect) as a cached slice.
func (it *OneFunc[T1]) Slice() []any {
	if it.isSliceCached {
		return it.toSlice
	}

	var args []any

	args = appendIfDefined(args, it.First)

	if it.HasFunc() {
		args = append(args, it.GetFuncName())
	}

	args = appendIfDefined(args, it.Expect)

	it.toSlice = args
	it.isSliceCached = true

	return it.toSlice
}

// GetByIndex safely retrieves an item from the cached slice by index.
func (it *OneFunc[T1]) GetByIndex(index int) any {
	return getByIndex(it.Slice(), index)
}

// String returns a formatted string representation.
func (it OneFunc[T1]) String() string {
	return buildToString(
		"OneFunc",
		it.Slice(),
		&it.toString,
	)
}

// LeftRight converts to a LeftRight with First as Left, WorkFunc as Right.
func (it *OneFunc[T1]) LeftRight() LeftRight {
	return LeftRight{
		Left:   it.First,
		Right:  it.WorkFunc,
		Expect: it.Expect,
	}
}

// AsOneFuncParameter returns the OneFunc as a OneFuncParameter interface.
func (it OneFunc[T1]) AsOneFuncParameter() OneFuncParameter {
	return &it
}

// AsArgFuncContractsBinder returns the OneFunc as an ArgFuncContractsBinder interface.
func (it OneFunc[T1]) AsArgFuncContractsBinder() ArgFuncContractsBinder {
	return &it
}

// AsArgBaseContractsBinder returns the OneFunc as an ArgBaseContractsBinder interface.
func (it OneFunc[T1]) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
