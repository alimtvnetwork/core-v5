package coretests

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type ArgFour struct {
	First  interface{} `json:",omitempty"`
	Second interface{} `json:",omitempty"`
	Third  interface{} `json:",omitempty"`
	Fourth interface{} `json:",omitempty"`
	Expect interface{} `json:",omitempty"`
}

func (it ArgFour) ArgTwo() ArgTwo {
	return ArgTwo{
		First:  it.First,
		Second: it.Second,
	}
}

func (it ArgFour) ArgThree() ArgThree {
	return ArgThree{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

func (it *ArgFour) HasFirst() bool {
	return it != nil && reflectinternal.IsNotNull(it.First)
}

func (it *ArgFour) HasSecond() bool {
	return it != nil && reflectinternal.IsNotNull(it.Second)
}

func (it *ArgFour) HasThird() bool {
	return it != nil && reflectinternal.IsNotNull(it.Third)
}

func (it *ArgFour) HasFourth() bool {
	return it != nil && reflectinternal.IsNotNull(it.Fourth)
}

func (it *ArgFour) HasExpect() bool {
	return it != nil && reflectinternal.IsNotNull(it.Expect)
}

func (it ArgFour) String() string {
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

	if it.HasExpect() {
		args = append(args, toString(it.Expect))
	}

	return fmt.Sprintf(
		"%s { %s }",
		"ArgFour",
		strings.Join(args, constants.CommaSpace))
}
