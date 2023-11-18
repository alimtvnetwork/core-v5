package coretests

type ArgSix struct {
	First  interface{} `json:",omitempty"`
	Second interface{} `json:",omitempty"`
	Third  interface{} `json:",omitempty"`
	Fourth interface{} `json:",omitempty"`
	Fifth  interface{} `json:",omitempty"`
	Six    interface{} `json:",omitempty"`
	Expect interface{} `json:",omitempty"`
}

func (it ArgSix) ArgTwo() ArgTwo {
	return ArgTwo{
		First:  it.First,
		Second: it.Second,
	}
}

func (it ArgSix) ArgThree() ArgThree {
	return ArgThree{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
	}
}

func (it ArgSix) ArgFour() ArgFour {
	return ArgFour{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}

func (it ArgSix) ArgFive() ArgFive {
	return ArgFive{
		First:  it.First,
		Second: it.Second,
		Third:  it.Third,
		Fourth: it.Fourth,
	}
}
