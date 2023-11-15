package coretests

type LeftRightExpect struct {
	Left   interface{} `json:",omitempty"`
	Right  interface{} `json:",omitempty"`
	Expect interface{} `json:",omitempty"`
}

func (it LeftRightExpect) ArgTwo() ArgTwo {
	return ArgTwo{
		First:  it.Left,
		Second: it.Right,
	}
}

func (it LeftRightExpect) ArgThree() ArgThree {
	return ArgThree{
		First:  it.Left,
		Second: it.Right,
		Third:  it.Expect,
	}
}
