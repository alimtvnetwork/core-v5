package coretests

type ArgFive struct {
	First  interface{} `json:",omitempty"`
	Second interface{} `json:",omitempty"`
	Third  interface{} `json:",omitempty"`
	Fourth interface{} `json:",omitempty"`
	Fifth  interface{} `json:",omitempty"`
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
