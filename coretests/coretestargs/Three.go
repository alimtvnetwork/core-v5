package coretestargs

type Three struct {
	First  interface{} `json:",omitempty"`
	Second interface{} `json:",omitempty"`
	Third  interface{} `json:",omitempty"`
	Expect interface{} `json:",omitempty"`
}

func (it Three) ArgTwo() Two {
	return Two{
		First:  it.First,
		Second: it.Second,
	}
}

func (it Three) LeftRight() LeftRight {
	return LeftRight{
		Left:   it.First,
		Right:  it.Second,
		Expect: it.Expect,
	}
}
