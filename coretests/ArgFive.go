package coretests

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type ArgFive struct {
	First  interface{} `json:",omitempty"`
	Second interface{} `json:",omitempty"`
	Third  interface{} `json:",omitempty"`
	Fourth interface{} `json:",omitempty"`
	Fifth  interface{} `json:",omitempty"`
	Expect interface{} `json:",omitempty"`
}

func (it ArgFive) ArgTwo() ArgTwo {
	return ArgTwo{
		First:  it.First,
		Second: it.Second,
	}
}

func (it ArgFive) ArgThree() ArgThree {
	return ArgThree{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

func (it ArgFive) ArgFour() ArgFour {
	return ArgFour{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

func (it *ArgFive) HasFirst() bool {
	return it != nil && reflectinternal.IsNotNull(it.First)
}

func (it *ArgFive) HasSecond() bool {
	return it != nil && reflectinternal.IsNotNull(it.Second)
}

func (it *ArgFive) HasThird() bool {
	return it != nil && reflectinternal.IsNotNull(it.Third)
}

func (it *ArgFive) HasFourth() bool {
	return it != nil && reflectinternal.IsNotNull(it.Fourth)
}

func (it *ArgFive) HasFifth() bool {
	return it != nil && reflectinternal.IsNotNull(it.Fifth)
}

func (it *ArgFive) HasExpect() bool {
	return it != nil && reflectinternal.IsNotNull(it.Expect)
}

func (it ArgFive) String() string {
	var args []string

	if it.HasFirst() {
		args = append(args, toString(it.First))
	}

	if it.HasSecond() {
		args = append(args, toString(it.Second))
	}

	if it.HasThird() {
		args = append(args, toString(it.Third))
	}

	if it.HasFourth() {
		args = append(args, toString(it.Fourth))
	}

	if it.HasFifth() {
		args = append(args, toString(it.Fifth))
	}

	if it.HasExpect() {
		args = append(args, toString(it.Expect))
	}

	return fmt.Sprintf(
		"%s { %s }",
		"ArgFive",
		strings.Join(args, constants.CommaSpace))
}
