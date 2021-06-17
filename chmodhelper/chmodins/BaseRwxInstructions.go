package chmodins

type BaseRwxInstructions struct {
	RwxInstructions *[]*RwxInstruction `json:"RwxInstructions,omitempty"`
}

func (receiver *BaseRwxInstructions) Length() int {
	if receiver == nil || receiver.RwxInstructions == nil {
		return 0
	}

	return len(* receiver.RwxInstructions)
}

func (receiver *BaseRwxInstructions) IsEmpty() bool {
	return receiver.Length() == 0
}

func (receiver *BaseRwxInstructions) HasAnyItem() bool {
	return receiver.Length() > 0
}

func (receiver *BaseRwxInstructions) Clone() *BaseRwxInstructions {
	if receiver == nil {
		return nil
	}

	instructions := make(
		[]*RwxInstruction,
		receiver.Length())

	for i, instruction := range *receiver.RwxInstructions {
		instructions[i] = instruction.Clone()
	}

	return &BaseRwxInstructions{
		RwxInstructions: &instructions,
	}
}
