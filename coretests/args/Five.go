package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type Five struct {
	First    interface{} `json:",omitempty"`
	Second   interface{} `json:",omitempty"`
	Third    interface{} `json:",omitempty"`
	Fourth   interface{} `json:",omitempty"`
	Fifth    interface{} `json:",omitempty"`
	WorkFunc interface{} `json:",omitempty"`
	Expect   interface{} `json:",omitempty"`
	toSlice  *[]interface{}
	toString corestr.SimpleStringOnce
}

func (it Five) ArgTwo() Two {
	return Two{
		First:  it.First,
		Second: it.Second,
	}
}

func (it Five) ArgThree() Three {
	return Three{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

func (it Five) ArgFour() Four {
	return Four{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

func (it *Five) HasFirst() bool {
	return it != nil && reflectinternal.IsNotNull(it.First)
}

func (it *Five) HasSecond() bool {
	return it != nil && reflectinternal.IsNotNull(it.Second)
}

func (it *Five) HasThird() bool {
	return it != nil && reflectinternal.IsNotNull(it.Third)
}

func (it *Five) HasFourth() bool {
	return it != nil && reflectinternal.IsNotNull(it.Fourth)
}

func (it *Five) HasFifth() bool {
	return it != nil && reflectinternal.IsNotNull(it.Fifth)
}

func (it *Five) HasExpect() bool {
	return it != nil && reflectinternal.IsNotNull(it.Expect)
}

func (it *Five) GetFuncName() string {
	return reflectinternal.GetFuncName(it.WorkFunc)
}

func (it Five) Slice() []interface{} {
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

	if it.HasExpect() {
		args = append(args, it.Expect)
	}

	it.toSlice = &args

	return *it.toSlice
}

func (it Five) GetByIndex(index int) interface{} {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it Five) String() string {
	if it.toString.IsInitialized() {
		return it.toString.String()
	}

	var args []string

	for _, item := range it.Slice() {
		args = append(args, toString(item))
	}

	toFinalString := fmt.Sprintf(
		"%s { %s }",
		"Five",
		strings.Join(args, constants.CommaSpace))

	return it.toString.GetSetOnce(toFinalString)
}
