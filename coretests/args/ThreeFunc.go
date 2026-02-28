package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type ThreeFunc struct {
	First    any                      `json:",omitempty"`
	Second   any                      `json:",omitempty"`
	Third    any                      `json:",omitempty"`
	WorkFunc any                      `json:"-"`
	Expect   any                      `json:",omitempty"`
	toSlice  *[]any                   `json:"-"`
	toString corestr.SimpleStringOnce `json:"-"`
}

func (it *ThreeFunc) GetWorkFunc() any {
	return it.WorkFunc
}

func (it *ThreeFunc) ArgsCount() int {
	return 3
}

func (it *ThreeFunc) FirstItem() any {
	return it.First
}

func (it *ThreeFunc) SecondItem() any {
	return it.Second
}

func (it *ThreeFunc) ThirdItem() any {
	return it.Third
}

func (it *ThreeFunc) Expected() any {
	return it.Expect
}

func (it *ThreeFunc) ArgTwo() TwoFunc {
	return TwoFunc{
		First:  it.First,
		Second: it.Second,
	}
}

func (it *ThreeFunc) ArgThree() ThreeFunc {
	return ThreeFunc{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

func (it *ThreeFunc) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

func (it *ThreeFunc) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

func (it *ThreeFunc) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

func (it *ThreeFunc) HasFunc() bool {
	return it != nil && reflectinternal.Is.Defined(it.WorkFunc)
}

func (it *ThreeFunc) GetFuncName() string {
	return reflectinternal.GetFunc.NameOnly(it.WorkFunc)
}

func (it *ThreeFunc) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it *ThreeFunc) FuncWrap() *FuncWrap {
	return NewFuncWrap.Default(it.WorkFunc)
}

func (it *ThreeFunc) Invoke(args ...any) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(args...)
}

func (it *ThreeFunc) InvokeMust(args ...any) (results []any) {
	results, err := it.FuncWrap().Invoke(args...)

	if err != nil {
		panic(err)
	}

	return results
}

func (it *ThreeFunc) InvokeWithValidArgs() (
	results []any, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.ValidArgs()

	return funcWrap.Invoke(validArgs...)
}

func (it *ThreeFunc) InvokeArgs(upTo int) (
	results []any, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.Args(upTo)

	return funcWrap.Invoke(validArgs...)
}

func (it *ThreeFunc) ValidArgs() []any {
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

	return args
}

func (it *ThreeFunc) Args(upTo int) []any {
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

func (it *ThreeFunc) Slice() []any {
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

	if it.HasExpect() {
		args = append(args, it.Expect)
	}

	it.toSlice = &args

	return *it.toSlice
}

func (it *ThreeFunc) GetByIndex(index int) any {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it ThreeFunc) String() string {
	if it.toString.IsInitialized() {
		return it.toString.String()
	}

	var args []string

	for _, item := range it.Slice() {
		args = append(args, toString(item))
	}

	toFinalString := fmt.Sprintf(
		selfToStringFmt,
		"ThreeFunc",
		strings.Join(args, constants.CommaSpace),
	)

	return it.toString.GetSetOnce(toFinalString)
}

func (it *ThreeFunc) LeftRight() LeftRight {
	return LeftRight{
		Left:   it.First,
		Right:  it.Second,
		Expect: it.Expect,
	}
}

func (it ThreeFunc) AsThreeFuncParameter() ThreeFuncParameter {
	return &it
}

func (it ThreeFunc) AsArgFuncContractsBinder() ArgFuncContractsBinder {
	return &it
}

func (it ThreeFunc) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
