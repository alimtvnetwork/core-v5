package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type Four struct {
	First    any `json:",omitempty"`
	Second   any `json:",omitempty"`
	Third    any `json:",omitempty"`
	Fourth   any `json:",omitempty"`
	Expect   any `json:",omitempty"`
	toSlice  *[]any
	toString corestr.SimpleStringOnce
}

func (it *Four) ArgsCount() int { return 4 }
func (it *Four) FirstItem() any { return it.First }
func (it *Four) SecondItem() any { return it.Second }
func (it *Four) ThirdItem() any { return it.Third }
func (it *Four) FourthItem() any { return it.Fourth }
func (it *Four) Expected() any { return it.Expect }

func (it *Four) ArgTwo() Two {
	return Two{First: it.First, Second: it.Second}
}

func (it *Four) ArgThree() Three {
	return Three{First: it.First, Second: it.Second, Third: it.Third}
}

func (it *Four) HasFirst() bool  { return it != nil && reflectinternal.Is.Defined(it.First) }
func (it *Four) HasSecond() bool { return it != nil && reflectinternal.Is.Defined(it.Second) }
func (it *Four) HasThird() bool  { return it != nil && reflectinternal.Is.Defined(it.Third) }
func (it *Four) HasFourth() bool { return it != nil && reflectinternal.Is.Defined(it.Fourth) }
func (it *Four) HasExpect() bool { return it != nil && reflectinternal.Is.Defined(it.Expect) }

func (it *Four) ValidArgs() []any {
	var args []any
	if it.HasFirst() { args = append(args, it.First) }
	if it.HasSecond() { args = append(args, it.Second) }
	if it.HasThird() { args = append(args, it.Third) }
	if it.HasFourth() { args = append(args, it.Fourth) }
	return args
}

func (it *Four) Args(upTo int) []any {
	var args []any
	if upTo >= 1 { args = append(args, it.First) }
	if upTo >= 2 { args = append(args, it.Second) }
	if upTo >= 3 { args = append(args, it.Third) }
	if upTo >= 4 { args = append(args, it.Fourth) }
	return args
}

func (it *Four) Slice() []any {
	if it.toSlice != nil { return *it.toSlice }
	var args []any
	if it.HasFirst() { args = append(args, it.First) }
	if it.HasSecond() { args = append(args, it.Second) }
	if it.HasThird() { args = append(args, it.Third) }
	if it.HasFourth() { args = append(args, it.Fourth) }
	if it.HasExpect() { args = append(args, it.Expect) }
	it.toSlice = &args
	return *it.toSlice
}

func (it *Four) GetByIndex(index int) any {
	slice := it.Slice()
	if len(slice)-1 < index { return nil }
	return slice[index]
}

func (it *Four) String() string {
	var args []string
	if it.HasFirst() { args = append(args, toString(it.First)) }
	if it.HasSecond() { args = append(args, toString(it.Second)) }
	if it.HasThird() { args = append(args, toString(it.Third)) }
	if it.HasFourth() { args = append(args, toString(it.Fourth)) }
	if it.HasExpect() { args = append(args, toString(it.Expect)) }
	return fmt.Sprintf(selfToStringFmt, "Four", strings.Join(args, constants.CommaSpace))
}

func (it Four) AsFourParameter() FourParameter               { return &it }
func (it Four) AsArgBaseContractsBinder() ArgBaseContractsBinder { return &it }
