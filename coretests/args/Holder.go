package args

import (
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coreinterface"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// Holder is a flexible 6-slot argument holder with a typed WorkFunc
// and a fallback Hashmap for dynamic parameters.
//
// Type parameter T represents the type of the WorkFunc field.
// Use HolderAny (= Holder[any]) for untyped usage.
//
// If positional parameters are not enough, use the Hashmap field
// for additional key-value arguments.
//
// Example (typed):
//
//	h := args.Holder[func(string) error]{
//	    First:    "input",
//	    WorkFunc: myProcessor,
//	}
//
// Example (untyped):
//
//	h := args.HolderAny{
//	    First:    "input",
//	    WorkFunc: myProcessor,
//	    Hashmap:  args.Map{"extra": "value"},
//	}
type Holder[T any] struct {
	First         any                      `json:",omitempty"`
	Second        any                      `json:",omitempty"`
	Third         any                      `json:",omitempty"`
	Fourth        any                      `json:",omitempty"`
	Fifth         any                      `json:",omitempty"`
	Sixth         any                      `json:",omitempty"`
	WorkFunc      T                        `json:"-"`
	Expect        any                      `json:",omitempty"`
	Hashmap       Map                      `json:",omitempty"`
	toSlice       []any                    `json:"-"`
	isSliceCached bool                     `json:"-"`
	toString      corestr.SimpleStringOnce `json:"-"`
}

// GetWorkFunc returns the wrapped function value as any.
func (it *Holder[T]) GetWorkFunc() any {
	return it.WorkFunc
}

// ArgsCount returns the number of positional argument slots (always 7).
func (it *Holder[T]) ArgsCount() int {
	return 7
}

// FirstItem returns the First argument.
func (it *Holder[T]) FirstItem() any {
	return it.First
}

// SecondItem returns the Second argument.
func (it *Holder[T]) SecondItem() any {
	return it.Second
}

// ThirdItem returns the Third argument.
func (it *Holder[T]) ThirdItem() any {
	return it.Third
}

// FourthItem returns the Fourth argument.
func (it *Holder[T]) FourthItem() any {
	return it.Fourth
}

// FifthItem returns the Fifth argument.
func (it *Holder[T]) FifthItem() any {
	return it.Fifth
}

// SixthItem returns the Sixth argument.
func (it *Holder[T]) SixthItem() any {
	return it.Sixth
}

// Expected returns the expected value.
func (it *Holder[T]) Expected() any {
	return it.Expect
}

// ArgTwo returns a TwoFuncAny with the first two arguments.
func (it *Holder[T]) ArgTwo() TwoFuncAny {
	return TwoFuncAny{
		First:  it.First,
		Second: it.Second,
	}
}

// ArgThree returns a ThreeFuncAny with the first three arguments.
func (it *Holder[T]) ArgThree() ThreeFuncAny {
	return ThreeFuncAny{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

// ArgFour returns a FourFuncAny with the first four arguments.
func (it *Holder[T]) ArgFour() FourFuncAny {
	return FourFuncAny{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

// ArgFive returns a FiveFuncAny with the first five arguments.
func (it *Holder[T]) ArgFive() FiveFuncAny {
	return FiveFuncAny{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
		Fifth:  it.Fifth,
	}
}

// HasFirst checks whether the First argument is defined.
func (it *Holder[T]) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

// HasSecond checks whether the Second argument is defined.
func (it *Holder[T]) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

// HasThird checks whether the Third argument is defined.
func (it *Holder[T]) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

// HasFourth checks whether the Fourth argument is defined.
func (it *Holder[T]) HasFourth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fourth)
}

// HasFifth checks whether the Fifth argument is defined.
func (it *Holder[T]) HasFifth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fifth)
}

// HasSixth checks whether the Sixth argument is defined.
func (it *Holder[T]) HasSixth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Sixth)
}

// HasFunc checks whether the WorkFunc is defined.
func (it *Holder[T]) HasFunc() bool {
	return it != nil && reflectinternal.Is.Defined(it.WorkFunc)
}

// HasExpect checks whether the Expect field is defined.
func (it *Holder[T]) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

// GetFuncName returns the short name of the wrapped function.
func (it *Holder[T]) GetFuncName() string {
	return reflectinternal.GetFunc.NameOnly(it.WorkFunc)
}

// FuncWrap wraps the WorkFunc in a FuncWrapAny for reflection-based invocation.
func (it *Holder[T]) FuncWrap() *FuncWrapAny {
	return NewFuncWrap.Default(it.WorkFunc)
}

// Invoke dynamically calls the WorkFunc with the given arguments.
func (it *Holder[T]) Invoke(args ...any) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(args...)
}

// InvokeMust invokes the WorkFunc, panicking on error.
func (it *Holder[T]) InvokeMust(args ...any) []any {
	return invokeMustHelper(it.FuncWrap(), args...)
}

// InvokeWithValidArgs invokes the WorkFunc with all defined positional arguments.
func (it *Holder[T]) InvokeWithValidArgs() (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(it.ValidArgs()...)
}

// InvokeArgs invokes the WorkFunc with positional arguments up to the given count.
func (it *Holder[T]) InvokeArgs(upTo int) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(it.Args(upTo)...)
}

// ValidArgs returns all defined positional arguments as a slice.
func (it *Holder[T]) ValidArgs() []any {
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
func (it *Holder[T]) Args(upTo int) []any {
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
func (it *Holder[T]) Slice() []any {
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
func (it *Holder[T]) GetByIndex(index int) any {
	return getByIndex(it.Slice(), index)
}

// String returns a formatted string representation.
func (it *Holder[T]) String() string {
	return buildToString(
		"Holder",
		it.Slice(),
		&it.toString,
	)
}

// AsSixthParameter returns the Holder as a SixthParameter interface.
func (it Holder[T]) AsSixthParameter() coreinterface.SixthParameter {
	return &it
}

// AsArgFuncContractsBinder returns the Holder as an ArgFuncContractsBinder interface.
func (it Holder[T]) AsArgFuncContractsBinder() ArgFuncContractsBinder {
	return &it
}
