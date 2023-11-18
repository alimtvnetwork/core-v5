package coretests

import "gitlab.com/auk-go/core/coretests/coretestargs"

type LeftRight struct {
	Left   interface{} `json:",omitempty"`
	Right  interface{} `json:",omitempty"`
	Expect interface{} `json:",omitempty"`
}

func (it LeftRight) ArgTwo() coretestargs.ArgTwo {
	return coretestargs.ArgTwo{
		First:  it.Left,
		Second: it.Right,
	}
}

func (it LeftRight) ArgThree() coretestargs.ArgThree {
	return coretestargs.ArgThree{
		First:  it.Left,
		Second: it.Right,
		Third:  it.Expect,
	}
}
