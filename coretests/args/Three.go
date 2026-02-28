package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type Three struct {
	First    any `json:",omitempty"`
	Second   any `json:",omitempty"`
	Third    any `json:",omitempty"`
	Expect   any `json:",omitempty"`
	toSlice  *[]any
	toString corestr.SimpleStringOnce
}

func (it *Three) ArgsCount() int { return 3 }
func (it *Three) FirstItem() any { return it.First }
func (it *Three) SecondItem() any { return it.Second }
func (it *Three) ThirdItem() any { return it.Third }
func (it *Three) Expected() any { return it.Expect }

func (it *Three) ArgTwo() TwoFunc {
	return TwoFunc{First: it.First, Second: it.Second}
}

func (it *Three) ArgThree() ThreeFunc {
	return ThreeFunc{First: it.First, Second: it.Second, Third: it.Third}
}

func (it *Three) HasFirst() bool  { return it != nil && reflectinternal.Is.Defined(it.First) }
func (it *Three) HasSecond() bool { return it != nil && reflectinternal.Is.Defined(it.Second) }
func (it *Three) HasThird() bool  { return it != nil && reflectinternal.Is.Defined(it.Third) }
func (it *Three) HasExpect() bool { return it != nil && reflectinternal.Is.Defined(it.Expect) }

func (it *Three) ValidArgs() []any {
	var args []any
	if it.HasFirst() { args = append(args, it.First) }
	if it.HasSecond() { args = append(args, it.Second) }
	if it.HasThird() { args = append(args, it.Third) }
	return args
}

func (it *Three) Args(upTo int) []any {
	var args []any
	if upTo >= 1 { args = append(args, it.First) }
	if upTo >= 2 { args = append(args, it.Second) }
	if upTo >= 3 { args = append(args, it.Third) }
	return args
}

func (it *Three) Slice() []any {
	if it.toSlice != nil { return *it.toSlice }
	var args []any
	if it.HasFirst() { args = append(args, it.First) }
	if it.HasSecond() { args = append(args, it.Second) }
	if it.HasThird() { args = append(args, it.Third) }
	if it.HasExpect() { args = append(args, it.Expect) }
	it.toSlice = &args
	return *it.toSlice
}

func (it *Three) GetByIndex(index int) any {
	slice := it.Slice()
	if len(slice)-1 < index { return nil }
	return slice[index]
}

func (it Three) String() string {
	if it.toString.IsInitialized() { return it.toString.String() }
	var args []string
	for _, item := range it.Slice() { args = append(args, toString(item)) }
	toFinalString := fmt.Sprintf(selfToStringFmt, "ThreeFunc", strings.Join(args, constants.CommaSpace))
	return it.toString.GetSetOnce(toFinalString)
}

func (it *Three) LeftRight() LeftRight {
	return LeftRight{Left: it.First, Right: it.Second, Expect: it.Expect}
}

func (it Three) AsThreeParameter() ThreeParameter          { return &it }
func (it Three) AsArgBaseContractsBinder() ArgBaseContractsBinder { return &it }
