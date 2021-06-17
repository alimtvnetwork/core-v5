package chmodins

type RwxInstruction struct {
	RwxOwnerGroupOther
	IsSkipOnNonExist  bool `json:"IsSkipOnNonExist"`
	IsContinueOnError bool `json:"IsContinueOnError"`
	IsRecursive       bool `json:"IsRecursive"`
}

func (receiver *RwxInstruction) Clone() *RwxInstruction {
	if receiver == nil {
		return nil
	}

	return &RwxInstruction{
		RwxOwnerGroupOther: *receiver.RwxOwnerGroupOther.Clone(),
		IsSkipOnNonExist:   receiver.IsSkipOnNonExist,
		IsContinueOnError:  receiver.IsContinueOnError,
		IsRecursive:        receiver.IsRecursive,
	}
}
