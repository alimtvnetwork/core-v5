package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type ThreeFunc struct {
	First    interface{} `json:",omitempty"`
	Second   interface{} `json:",omitempty"`
	Third    interface{} `json:",omitempty"`
	WorkFunc interface{} `json:",omitempty"`
	Expect   interface{} `json:",omitempty"`
	toSlice  *[]interface{}
	toString corestr.SimpleStringOnce
}

func (it ThreeFunc) ArgTwo() TwoFunc {
	return TwoFunc{
		First:  it.First,
		Second: it.Second,
	}
}

func (it ThreeFunc) ArgThree() ThreeFunc {
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

func (it *ThreeFunc) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it *ThreeFunc) GetFuncName() string {
	return reflectinternal.GetFuncName(it.WorkFunc)
}

func (it ThreeFunc) Slice() []interface{} {
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

	if it.HasFunc() {
		args = append(args, it.GetFuncName())
	}

	if it.HasExpect() {
		args = append(args, it.Expect)
	}

	it.toSlice = &args

	return *it.toSlice
}

func (it ThreeFunc) GetByIndex(index int) interface{} {
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
		"%s { %s }",
		"ThreeFunc",
		strings.Join(args, constants.CommaSpace),
	)

	return it.toString.GetSetOnce(toFinalString)
}

func (it ThreeFunc) LeftRight() LeftRight {
	return LeftRight{
		Left:   it.First,
		Right:  it.Second,
		Expect: it.Expect,
	}
}
