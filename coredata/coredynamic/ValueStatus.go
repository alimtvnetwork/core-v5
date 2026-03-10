package coredynamic

import "github.com/alimtvnetwork/core/constants"

type ValueStatus struct {
	IsValid bool
	Message string
	Index   int
	Value   any
}

func InvalidValueStatusNoMessage() *ValueStatus {
	return InvalidValueStatus(constants.EmptyString)
}

func InvalidValueStatus(message string) *ValueStatus {
	return &ValueStatus{
		IsValid: false,
		Message: message,
		Index:   constants.InvalidNotFoundCase,
		Value:   nil,
	}
}
