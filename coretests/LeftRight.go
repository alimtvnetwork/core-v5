package coretests

type LeftRight struct {
	Left   interface{} `json:",omitempty"`
	Right  interface{} `json:",omitempty"`
	Expect interface{} `json:",omitempty"`
}

func (it LeftRight) ArgTwo() ArgTwo {
	return ArgTwo{
		First:  it.Left,
		Second: it.Right,
	}
}

func (it LeftRight) ArgThree() ArgThree {
	return ArgThree{
		First:  it.Left,
		Second: it.Right,
		Third:  it.Expect,
	}
}
