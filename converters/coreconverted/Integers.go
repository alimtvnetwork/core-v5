package coreconverted

type Integers struct {
	Values        *[]int
	CombinedError error
}

func (receiver *Integers) HasError() bool {
	return receiver.CombinedError != nil
}

func (receiver *Integers) Length() int {
	if receiver.Values == nil {
		return 0
	}

	return len(*receiver.Values)
}

func (receiver *Integers) IsEmpty() bool {
	return receiver.Length() == 0
}

func (receiver *Integers) HandleWithPanic() {
	if receiver.CombinedError == nil {
		return
	}

	panic(receiver.CombinedError)
}
