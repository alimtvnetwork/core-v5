package args

type LeftRight struct {
	Left   interface{} `json:",omitempty"`
	Right  interface{} `json:",omitempty"`
	Expect interface{} `json:",omitempty"`
}

func (it LeftRight) ArgTwo() TwoFunc {
	return TwoFunc{
		First:  it.Left,
		Second: it.Right,
	}
}

func (it LeftRight) ArgThree() ThreeFunc {
	return ThreeFunc{
		First:  it.Left,
		Second: it.Right,
		Third:  it.Expect,
	}
}
