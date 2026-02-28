package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type FourFunc struct {
	First    any                      `json:",omitempty"`
	Second   any                      `json:",omitempty"`
	Third    any                      `json:",omitempty"`
	Fourth   any                      `json:",omitempty"`
	WorkFunc any                      `json:",omitempty"`
	Expect   any                      `json:",omitempty"`
	toSlice  *[]any
	toString corestr.SimpleStringOnce
}

func (it *FourFunc) GetWorkFunc() any {
	return it.WorkFunc
}

func (it *FourFunc) ArgsCount() int {
	return 4
}

func (it *FourFunc) FirstItem() any {
	return it.First
}

func (it *FourFunc) SecondItem() any {
	return it.Second
}

func (it *FourFunc) ThirdItem() any {
	return it.Third
}

func (it *FourFunc) FourthItem() any {
	return it.Fourth
}

func (it *FourFunc) Expected() any {
	return it.Expect
}

func (it *FourFunc) ArgTwo() TwoFunc {
	return TwoFunc{
		First:  it.First,
		Second: it.Second,
	}
}

func (it *FourFunc) ArgThree() ThreeFunc {
	return ThreeFunc{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

func (it *FourFunc) ArgFour() FourFunc {
	return FourFunc{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

func (it *FourFunc) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

func (it *FourFunc) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

func (it *FourFunc) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

func (it *FourFunc) HasFourth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fourth)
}

func (it *FourFunc) HasFunc() bool {
	return it != nil && reflectinternal.Is.Defined(it.WorkFunc)
}

func (it *FourFunc) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it *FourFunc) GetFuncName() string {
	return reflectinternal.GetFunc.NameOnly(it.WorkFunc)
}

func (it *FourFunc) FuncWrap() *FuncWrap {
	return NewFuncWrap.Default(it.WorkFunc)
}

func (it *FourFunc) Invoke(args ...any) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(args...)
}

func (it *FourFunc) InvokeMust(args ...any) (results []any) {
	results, err := it.FuncWrap().Invoke(args...)

	if err != nil {
		panic(err)
	}

	return results
}

func (it *FourFunc) InvokeWithValidArgs() (
	results []any, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.ValidArgs()

	return funcWrap.Invoke(validArgs...)
}

func (it *FourFunc) InvokeArgs(upTo int) (
	results []any, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.Args(upTo)

	return funcWrap.Invoke(validArgs...)
}

func (it *FourFunc) ValidArgs() []any {
	var args []any

	if it.HasFirst() {
		args = append(args, it.First)
	}

	if it.HasSecond() {
		args = append(args, it.Second)
	}

	if it.HasThird() {
		args = append(args, it.Third)
	}

	if it.HasFourth() {
		args = append(args, it.Fourth)
	}

	return args
}

func (it *FourFunc) Args(upTo int) []any {
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

func (it *FourFunc) Slice() []any {
	if it.toSlice != nil {
		return *it.toSlice
	}

	var args []any

	if it.HasFirst() {
		args = append(args, it.First)
	}

	if it.HasSecond() {
		args = append(args, it.Second)
	}

	if it.HasThird() {
		args = append(args, it.Third)
	}

	if it.HasFourth() {
		args = append(args, it.Fourth)
	}

	if it.HasFunc() {
		args = append(args, it.GetFuncName())
	}

	if it.HasExpect() {
		args = append(args, it.Expect)
	}

	it.toSlice = &args

	return *it.toSlice
}

func (it *FourFunc) GetByIndex(index int) any {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it FourFunc) String() string {
	if it.toString.IsInitialized() {
		return it.toString.String()
	}

	var args []string

	for _, item := range it.Slice() {
		args = append(args, toString(item))
	}

	toFinalString := fmt.Sprintf(
		selfToStringFmt,
		"FourFunc",
		strings.Join(args, constants.CommaSpace),
	)

	return it.toString.GetSetOnce(toFinalString)
}

func (it FourFunc) AsFourFuncParameter() FourFuncParameter {
	return &it
}

func (it FourFunc) AsArgFuncContractsBinder() ArgFuncContractsBinder {
	return &it
}

func (it FourFunc) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
