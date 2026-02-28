package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// Holder is used to hold more dynamic parameters for
// the Act function in the unit or integration test.
// If parameters are not enough use the Hashmap.
type Holder struct {
	First    any              `json:",omitempty"`
	Second   any              `json:",omitempty"`
	Third    any              `json:",omitempty"`
	Fourth   any              `json:",omitempty"`
	Fifth    any              `json:",omitempty"`
	Sixth    any              `json:",omitempty"`
	WorkFunc any              `json:"-"`
	Expect   any              `json:",omitempty"`
	Hashmap  Map              `json:",omitempty"`
	toSlice  *[]any           `json:"-"`
	toString corestr.SimpleStringOnce `json:"-"`
}

func (it *Holder) GetWorkFunc() any { return it.WorkFunc }
func (it *Holder) ArgsCount() int   { return 7 }
func (it *Holder) FirstItem() any   { return it.First }
func (it *Holder) SecondItem() any  { return it.Second }
func (it *Holder) ThirdItem() any   { return it.Third }
func (it *Holder) FourthItem() any  { return it.Fourth }
func (it *Holder) FifthItem() any   { return it.Fifth }
func (it *Holder) SixthItem() any   { return it.Sixth }
func (it *Holder) Expected() any    { return it.Expect }

func (it *Holder) ArgTwo() TwoFunc { return TwoFunc{First: it.First, Second: it.Second} }
func (it *Holder) ArgThree() ThreeFunc { return ThreeFunc{First: it.First, Second: it.Second, Third: it.Third} }
func (it *Holder) ArgFour() FourFunc { return FourFunc{First: it.First, Second: it.Second, Third: it.Third, Fourth: it.Fourth} }
func (it *Holder) ArgFive() FiveFunc { return FiveFunc{First: it.First, Second: it.Second, Third: it.Third, Fourth: it.Fourth} }

func (it *Holder) HasFirst() bool  { return it != nil && reflectinternal.Is.Defined(it.First) }
func (it *Holder) HasSecond() bool { return it != nil && reflectinternal.Is.Defined(it.Second) }
func (it *Holder) HasThird() bool  { return it != nil && reflectinternal.Is.Defined(it.Third) }
func (it *Holder) HasFourth() bool { return it != nil && reflectinternal.Is.Defined(it.Fourth) }
func (it *Holder) HasFifth() bool  { return it != nil && reflectinternal.Is.Defined(it.Fifth) }
func (it *Holder) HasSixth() bool  { return it != nil && reflectinternal.Is.Defined(it.Sixth) }
func (it *Holder) HasFunc() bool   { return it != nil && reflectinternal.Is.Defined(it.WorkFunc) }
func (it *Holder) HasExpect() bool { return it != nil && reflectinternal.Is.Defined(it.Expect) }

func (it *Holder) GetFuncName() string { return reflectinternal.GetFunc.NameOnly(it.WorkFunc) }
func (it *Holder) FuncWrap() *FuncWrap { return NewFuncWrap.Default(it.WorkFunc) }

func (it *Holder) Invoke(args ...any) (results []any, processingErr error) {
	return it.FuncWrap().Invoke(args...)
}

func (it *Holder) InvokeMust(args ...any) (results []any) {
	results, err := it.FuncWrap().Invoke(args...)
	if err != nil { panic(err) }
	return results
}

func (it *Holder) InvokeWithValidArgs() (results []any, processingErr error) {
	return it.FuncWrap().Invoke(it.ValidArgs()...)
}

func (it *Holder) InvokeArgs(upTo int) (results []any, processingErr error) {
	return it.FuncWrap().Invoke(it.Args(upTo)...)
}

func (it *Holder) ValidArgs() []any {
	var args []any
	if it.HasFirst() { args = append(args, it.First) }
	if it.HasSecond() { args = append(args, it.Second) }
	if it.HasThird() { args = append(args, it.Third) }
	if it.HasFourth() { args = append(args, it.Fourth) }
	if it.HasFifth() { args = append(args, it.Fifth) }
	if it.HasSixth() { args = append(args, it.Sixth) }
	return args
}

func (it *Holder) Args(upTo int) []any {
	var args []any
	if upTo >= 1 { args = append(args, it.First) }
	if upTo >= 2 { args = append(args, it.Second) }
	if upTo >= 3 { args = append(args, it.Third) }
	if upTo >= 4 { args = append(args, it.Fourth) }
	if upTo >= 5 { args = append(args, it.Fifth) }
	if upTo >= 6 { args = append(args, it.Sixth) }
	return args
}

func (it *Holder) Slice() []any {
	if it.toSlice != nil { return *it.toSlice }
	var args []any
	if it.HasFirst() { args = append(args, it.First) }
	if it.HasSecond() { args = append(args, it.Second) }
	if it.HasThird() { args = append(args, it.Third) }
	if it.HasFourth() { args = append(args, it.Fourth) }
	if it.HasFifth() { args = append(args, it.Fifth) }
	if it.HasSixth() { args = append(args, it.Sixth) }
	if it.HasFunc() { args = append(args, it.GetFuncName()) }
	if it.HasExpect() { args = append(args, it.Expect) }
	it.toSlice = &args
	return *it.toSlice
}

func (it *Holder) GetByIndex(index int) any {
	slice := it.Slice()
	if len(slice)-1 < index { return nil }
	return slice[index]
}

func (it *Holder) String() string {
	if it.toString.IsInitialized() { return it.toString.String() }
	var args []string
	for _, item := range it.Slice() { args = append(args, toString(item)) }
	toFinalString := fmt.Sprintf(selfToStringFmt, "Holder", strings.Join(args, constants.CommaSpace))
	return it.toString.GetSetOnce(toFinalString)
}

func (it Holder) AsSixthParameter() coreinterface.SixthParameter { return &it }
func (it Holder) AsArgFuncContractsBinder() ArgFuncContractsBinder { return &it }
