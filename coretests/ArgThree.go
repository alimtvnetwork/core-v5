package coretests

type ArgThree struct {
	First  interface{} `json:",omitempty"`
	Second interface{} `json:",omitempty"`
	Third  interface{} `json:",omitempty"`
	Expect interface{} `json:",omitempty"`
}

func (it ArgThree) ArgTwo() ArgTwo {
	return ArgTwo{
		First:  it.First,
		Second: it.Second,
	}
}

func (it ArgThree) LeftRightExpect() LeftRight {
	return LeftRight{
		Left:   it.First,
		Right:  it.Second,
		Expect: it.Third,
	}
}
