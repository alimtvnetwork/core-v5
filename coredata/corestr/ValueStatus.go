package corestr

import "gitlab.com/evatix-go/core/constants"

type ValueStatus struct {
	ValueValid *ValueValid
	Index      int
}

func InvalidValueStatusNoMessage() *ValueStatus {
	return InvalidValueStatus(constants.EmptyString)
}

func InvalidValueStatus(message string) *ValueStatus {
	return &ValueStatus{
		ValueValid: InvalidValueValid(message),
		Index:      constants.InvalidNotFoundCase,
	}
}

func (v *ValueStatus) Clone() *ValueStatus {
	return &ValueStatus{
		ValueValid: v.ValueValid.Clone(),
		Index:      v.Index,
	}
}
