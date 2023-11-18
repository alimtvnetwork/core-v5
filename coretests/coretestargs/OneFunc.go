package coretestargs

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type OneFunc struct {
	First    interface{} `json:",omitempty"`
	WorkFunc interface{} `json:",omitempty"`
	Expect   interface{} `json:",omitempty"`
	toSlice  *[]interface{}
	toString corestr.SimpleStringOnce
}

func (it OneFunc) ArgTwo() OneFunc {
	return OneFunc{
		First:  it.First,
		Expect: it.Expect,
	}
}

func (it *OneFunc) HasFirst() bool {
	return it != nil && reflectinternal.IsNotNull(it.First)
}

func (it *OneFunc) HasFunc() bool {
	return it != nil && reflectinternal.IsNotNull(it.WorkFunc)
}

func (it *OneFunc) HasExpect() bool {
	return it != nil && reflectinternal.IsNotNull(it.Expect)
}

func (it *OneFunc) GetFuncName() string {
	return reflectinternal.GetFuncName(it.WorkFunc)
}

func (it OneFunc) Slice() []interface{} {
	if it.toSlice != nil {
		return *it.toSlice
	}

	var args []interface{}

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

func (it OneFunc) GetByIndex(index int) interface{} {
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
		"%s { %s }",
		"OneFunc",
		strings.Join(args, constants.CommaSpace))

	return it.toString.GetSetOnce(toFinalString)
}

func (it OneFunc) LeftRight() LeftRight {
	return LeftRight{
		Left:   it.First,
		Expect: it.Expect,
	}
}
