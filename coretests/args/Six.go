package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type ArgSix struct {
	First    interface{} `json:",omitempty"`
	Second   interface{} `json:",omitempty"`
	Third    interface{} `json:",omitempty"`
	Fourth   interface{} `json:",omitempty"`
	Fifth    interface{} `json:",omitempty"`
	Sixth    interface{} `json:",omitempty"`
	Expect   interface{} `json:",omitempty"`
	toSlice  *[]interface{}
	toString corestr.SimpleStringOnce
}

func (it ArgSix) ArgTwo() Two {
	return Two{
		First:  it.First,
		Second: it.Second,
	}
}

func (it ArgSix) ArgThree() Three {
	return Three{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

func (it ArgSix) ArgFour() Four {
	return Four{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

func (it ArgSix) ArgFive() Five {
	return Five{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

func (it *ArgSix) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

func (it *ArgSix) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

func (it *ArgSix) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

func (it *ArgSix) HasFourth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fourth)
}

func (it *ArgSix) HasFifth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fifth)
}

func (it *ArgSix) HasSixth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Sixth)
}

func (it *ArgSix) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it ArgSix) Slice() []interface{} {
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

	if it.HasExpect() {
		args = append(args, it.Expect)
	}

	it.toSlice = &args

	return *it.toSlice
}

func (it ArgSix) GetByIndex(index int) interface{} {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it ArgSix) String() string {
	if it.toString.IsInitialized() {
		return it.toString.String()
	}

	var args []string

	for _, item := range it.Slice() {
		args = append(args, toString(item))
	}

	toFinalString := fmt.Sprintf(
		selfToStringFmt,
		"ArgSix",
		strings.Join(args, constants.CommaSpace),
	)

	return it.toString.GetSetOnce(toFinalString)
}
