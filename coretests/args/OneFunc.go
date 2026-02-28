package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type OneFunc struct {
	First    any                      `json:",omitempty"`
	WorkFunc any                      `json:"-,omitempty"`
	Expect   any                      `json:",omitempty"`
	toSlice  *[]any                   `json:"-"`
	toString corestr.SimpleStringOnce `json:"-"`
}

func (it *OneFunc) GetWorkFunc() any {
	return it.WorkFunc
}

func (it *OneFunc) FirstItem() any {
	return it.First
}

func (it *OneFunc) Expected() any {
	return it.Expect
}

func (it *OneFunc) ArgTwo() OneFunc {
	return OneFunc{
		First:    it.First,
		WorkFunc: it.WorkFunc,
		Expect:   it.Expect,
	}
}

func (it *OneFunc) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

func (it *OneFunc) HasFunc() bool {
	return it != nil && reflectinternal.Is.Defined(it.WorkFunc)
}

func (it *OneFunc) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it *OneFunc) GetFuncName() string {
	return reflectinternal.GetFunc.NameOnly(it.WorkFunc)
}

func (it *OneFunc) FuncWrap() *FuncWrap {
	return NewFuncWrap.Default(it.WorkFunc)
}

func (it *OneFunc) Invoke(args ...any) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(args...)
}

func (it *OneFunc) InvokeMust(args ...any) (results []any) {
	results, err := it.FuncWrap().Invoke(args...)

	if err != nil {
		panic(err)
	}

	return results
}

func (it *OneFunc) InvokeWithValidArgs() (
	results []any, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.ValidArgs()

	return funcWrap.Invoke(validArgs...)
}

func (it *OneFunc) InvokeArgs(upTo int) (
	results []any, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.Args(upTo)

	return funcWrap.Invoke(validArgs...)
}

func (it *OneFunc) ValidArgs() []any {
	var args []any

	if it.HasFirst() {
		args = append(args, it.First)
	}

	return args
}

func (it *OneFunc) ArgsCount() int {
	return 1
}

func (it *OneFunc) Args(upTo int) []any {
	var args []any

	if upTo >= 1 {
		args = append(args, it.First)
	}

	return args
}

func (it *OneFunc) Slice() []any {
	if it.toSlice != nil {
		return *it.toSlice
	}

	var args []any

	if it.HasFirst() {
		args = append(args, it.First)
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

func (it *OneFunc) GetByIndex(index int) any {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it OneFunc) String() string {
	if it.toString.IsInitialized() {
		return it.toString.String()
	}

	var args []string

	for _, item := range it.Slice() {
		args = append(args, toString(item))
	}

	toFinalString := fmt.Sprintf(
		selfToStringFmt,
		"OneFunc",
		strings.Join(args, constants.CommaSpace),
	)

	return it.toString.GetSetOnce(toFinalString)
}

func (it *OneFunc) LeftRight() LeftRight {
	return LeftRight{
		Left:   it.First,
		Right:  it.WorkFunc,
		Expect: it.Expect,
	}
}

func (it OneFunc) AsOneFuncParameter() OneFuncParameter {
	return &it
}

func (it OneFunc) AsArgFuncContractsBinder() ArgFuncContractsBinder {
	return &it
}

func (it OneFunc) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
