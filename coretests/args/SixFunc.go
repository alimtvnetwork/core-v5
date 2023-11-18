package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type SixFunc struct {
	First    interface{} `json:",omitempty"`
	Second   interface{} `json:",omitempty"`
	Third    interface{} `json:",omitempty"`
	Fourth   interface{} `json:",omitempty"`
	Fifth    interface{} `json:",omitempty"`
	Sixth    interface{} `json:",omitempty"`
	WorkFunc interface{} `json:",omitempty"`
	Expect   interface{} `json:",omitempty"`
	toSlice  *[]interface{}
	toString corestr.SimpleStringOnce
}

func (it SixFunc) ArgTwo() TwoFunc {
	return TwoFunc{
		First:  it.First,
		Second: it.Second,
	}
}

func (it SixFunc) ArgThree() ThreeFunc {
	return ThreeFunc{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

func (it SixFunc) ArgFour() FourFunc {
	return FourFunc{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

func (it SixFunc) ArgFive() FiveFunc {
	return FiveFunc{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

func (it *SixFunc) HasFirst() bool {
	return it != nil && reflectinternal.IsNotNull(it.First)
}

func (it *SixFunc) HasSecond() bool {
	return it != nil && reflectinternal.IsNotNull(it.Second)
}

func (it *SixFunc) HasThird() bool {
	return it != nil && reflectinternal.IsNotNull(it.Third)
}

func (it *SixFunc) HasFourth() bool {
	return it != nil && reflectinternal.IsNotNull(it.Fourth)
}

func (it *SixFunc) HasFifth() bool {
	return it != nil && reflectinternal.IsNotNull(it.Fifth)
}

func (it *SixFunc) HasSixth() bool {
	return it != nil && reflectinternal.IsNotNull(it.Sixth)
}

func (it *SixFunc) HasFunc() bool {
	return it != nil && reflectinternal.IsNotNull(it.WorkFunc)
}

func (it *SixFunc) HasExpect() bool {
	return it != nil && reflectinternal.IsNotNull(it.Expect)
}

func (it *SixFunc) GetFuncName() string {
	return reflectinternal.GetFuncName(it.WorkFunc)
}

func (it SixFunc) Slice() []interface{} {
	if it.toSlice != nil {
		return *it.toSlice
	}

	var args []interface{}

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

func (it SixFunc) GetByIndex(index int) interface{} {
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
		"%s { %s }",
		"SixFunc",
		strings.Join(args, constants.CommaSpace))

	return it.toString.GetSetOnce(toFinalString)
}
