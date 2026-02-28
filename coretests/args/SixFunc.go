package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type SixFunc struct {
	First    any                      `json:",omitempty"`
	Second   any                      `json:",omitempty"`
	Third    any                      `json:",omitempty"`
	Fourth   any                      `json:",omitempty"`
	Fifth    any                      `json:",omitempty"`
	Sixth    any                      `json:",omitempty"`
	WorkFunc any                      `json:"-"`
	Expect   any                      `json:",omitempty"`
	toSlice  *[]any                   `json:"-"`
	toString corestr.SimpleStringOnce `json:"-"`
}

func (it *SixFunc) GetWorkFunc() any {
	return it.WorkFunc
}

func (it *SixFunc) ArgsCount() int {
	return 6
}

func (it *SixFunc) FirstItem() any {
	return it.First
}

func (it *SixFunc) SecondItem() any {
	return it.Second
}

func (it *SixFunc) ThirdItem() any {
	return it.Third
}

func (it *SixFunc) FourthItem() any {
	return it.Fourth
}

func (it *SixFunc) FifthItem() any {
	return it.Fifth
}

func (it *SixFunc) SixthItem() any {
	return it.Sixth
}

func (it *SixFunc) Expected() any {
	return it.Expect
}

func (it *SixFunc) ArgTwo() TwoFunc {
	return TwoFunc{
		First:  it.First,
		Second: it.Second,
	}
}

func (it *SixFunc) ArgThree() ThreeFunc {
	return ThreeFunc{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

func (it *SixFunc) ArgFour() FourFunc {
	return FourFunc{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

func (it *SixFunc) ArgFive() FiveFunc {
	return FiveFunc{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

func (it *SixFunc) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

func (it *SixFunc) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

func (it *SixFunc) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

func (it *SixFunc) HasFourth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fourth)
}

func (it *SixFunc) HasFifth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fifth)
}

func (it *SixFunc) HasSixth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Sixth)
}

func (it *SixFunc) HasFunc() bool {
	return it != nil && reflectinternal.Is.Defined(it.WorkFunc)
}

func (it *SixFunc) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it *SixFunc) GetFuncName() string {
	return reflectinternal.GetFunc.NameOnly(it.WorkFunc)
}

func (it *SixFunc) FuncWrap() *FuncWrap {
	return NewFuncWrap.Default(it.WorkFunc)
}

func (it *SixFunc) Invoke(args ...any) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(args...)
}

func (it *SixFunc) InvokeMust(args ...any) (results []any) {
	results, err := it.FuncWrap().Invoke(args...)

	if err != nil {
		panic(err)
	}

	return results
}

func (it *SixFunc) InvokeWithValidArgs() (
	results []any, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.ValidArgs()

	return funcWrap.Invoke(validArgs...)
}

func (it *SixFunc) InvokeArgs(upTo int) (
	results []any, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.Args(upTo)

	return funcWrap.Invoke(validArgs...)
}

func (it *SixFunc) ValidArgs() []any {
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

	if it.HasFifth() {
		args = append(args, it.Fifth)
	}

	if it.HasSixth() {
		args = append(args, it.Sixth)
	}

	return args
}

func (it *SixFunc) Args(upTo int) []any {
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

func (it *SixFunc) Slice() []any {
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

	if it.HasFifth() {
		args = append(args, it.Fifth)
	}

	if it.HasSixth() {
		args = append(args, it.Sixth)
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

func (it *SixFunc) GetByIndex(index int) any {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it SixFunc) String() string {
	if it.toString.IsInitialized() {
		return it.toString.String()
	}

	var args []string

	for _, item := range it.Slice() {
		args = append(args, toString(item))
	}

	toFinalString := fmt.Sprintf(
		selfToStringFmt,
		"SixFunc",
		strings.Join(args, constants.CommaSpace),
	)

	return it.toString.GetSetOnce(toFinalString)
}

func (it SixFunc) AsSixthFuncParameter() SixthFuncParameter {
	return &it
}

func (it SixFunc) AsArgFuncContractsBinder() ArgFuncContractsBinder {
	return &it
}

func (it SixFunc) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
