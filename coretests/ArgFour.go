package coretests

type ArgFour struct {
	First  interface{} `json:",omitempty"`
	Second interface{} `json:",omitempty"`
	Third  interface{} `json:",omitempty"`
	Fourth interface{} `json:",omitempty"`
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
