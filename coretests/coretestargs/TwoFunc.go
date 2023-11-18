package coretestargs

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type TwoFunc struct {
	First    interface{} `json:",omitempty"`
	Second   interface{} `json:",omitempty"`
	WorkFunc interface{} `json:",omitempty"`
	Expect   interface{} `json:",omitempty"`
	toSlice  *[]interface{}
	toString corestr.SimpleStringOnce
}

func (it TwoFunc) ArgTwo() TwoFunc {
	return TwoFunc{
		First:  it.First,
		Second: it.Second,
	}
}

func (it *TwoFunc) HasFirst() bool {
	return it != nil && reflectinternal.IsNotNull(it.First)
}

func (it *TwoFunc) HasSecond() bool {
	return it != nil && reflectinternal.IsNotNull(it.Second)
}

func (it *TwoFunc) HasFunc() bool {
	return it != nil && reflectinternal.IsNotNull(it.WorkFunc)
}

func (it *TwoFunc) HasExpect() bool {
	return it != nil && reflectinternal.IsNotNull(it.Expect)
}

func (it *TwoFunc) GetFuncName() string {
	return reflectinternal.GetFuncName(it.WorkFunc)
}

func (it TwoFunc) Slice() []interface{} {
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

	if it.HasFunc() {
		args = append(args, it.GetFuncName())
	}

	if it.HasExpect() {
		args = append(args, it.Expect)
	}

	it.toSlice = &args

	return *it.toSlice
}

func (it TwoFunc) GetByIndex(index int) interface{} {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it TwoFunc) String() string {
	if it.toString.IsInitialized() {
		return it.toString.String()
	}

	var args []string

	for _, item := range it.Slice() {
		args = append(args, toString(item))
	}

	toFinalString := fmt.Sprintf(
		"%s { %s }",
		"TwoFunc",
		strings.Join(args, constants.CommaSpace))

	return it.toString.GetSetOnce(toFinalString)
}

func (it TwoFunc) LeftRightExpect() LeftRight {
	return LeftRight{
		Left:   it.First,
		Right:  it.Second,
		Expect: it.Expect,
	}
}
