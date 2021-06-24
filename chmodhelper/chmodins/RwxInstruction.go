package chmodins

type RwxInstruction struct {
	RwxOwnerGroupOther
	Condition
}

func (receiver *RwxInstruction) Clone() *RwxInstruction {
	if receiver == nil {
		return nil
	}

	return &RwxInstruction{
		RwxOwnerGroupOther: *receiver.RwxOwnerGroupOther.Clone(),
		Condition: Condition{
			IsSkipOnInvalid:   receiver.IsSkipOnInvalid,
			IsContinueOnError: receiver.IsContinueOnError,
			IsRecursive:       receiver.IsRecursive,
		},
	}
}
