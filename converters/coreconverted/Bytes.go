package coreconverted

type Bytes struct {
	Values        *[]byte
	CombinedError error
}

func (receiver *Bytes) HasError() bool {
	return receiver.CombinedError != nil
}

func (receiver *Bytes) Length() int {
	if receiver.Values == nil {
		return 0
	}

	return len(*receiver.Values)
}

func (receiver *Bytes) IsEmpty() bool {
	return receiver.Length() == 0
}

func (receiver *Bytes) HandleWithPanic() {
	if receiver.CombinedError == nil {
		return
	}

	panic(receiver.CombinedError)
}
