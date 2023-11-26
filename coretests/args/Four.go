package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type Four struct {
	First  interface{} `json:",omitempty"`
	Second interface{} `json:",omitempty"`
	Third  interface{} `json:",omitempty"`
	Fourth interface{} `json:",omitempty"`
	Expect interface{} `json:",omitempty"`
}

func (it Four) ArgTwo() Two {
	return Two{
		First:  it.First,
		Second: it.Second,
	}
}

func (it Four) ArgThree() Three {
	return Three{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

func (it *Four) HasFirst() bool {
	return it != nil && reflectinternal.Is.Defined(it.First)
}

func (it *Four) HasSecond() bool {
	return it != nil && reflectinternal.Is.Defined(it.Second)
}

func (it *Four) HasThird() bool {
	return it != nil && reflectinternal.Is.Defined(it.Third)
}

func (it *Four) HasFourth() bool {
	return it != nil && reflectinternal.Is.Defined(it.Fourth)
}

func (it *Four) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it Four) String() string {
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
		"Four",
		strings.Join(args, constants.CommaSpace),
	)
}
