package coredynamic

import "gitlab.com/evatix-go/core/constants"

type DynamicStatus struct {
	Dynamic
	Index   int
	Message string
}

func InvalidDynamicStatusNoMessage() *DynamicStatus {
	return InvalidDynamicStatus(constants.EmptyString)
}

func InvalidDynamicStatus(message string) *DynamicStatus {
	return &DynamicStatus{
		Dynamic: NewDynamic(nil, false),
		Index:   constants.InvalidNotFoundCase,
		Message: message,
	}
}

// Clone Warning: Cannot clone dynamic data or interface properly but set it again
//
// If it is a pointer one needs to copy it manually.
func (receiver *DynamicStatus) Clone() *DynamicStatus {
	if receiver == nil {
		return nil
	}

	return &DynamicStatus{
		Dynamic: *receiver.Dynamic.Clone(),
		Index:   constants.InvalidNotFoundCase,
		Message: receiver.Message,
	}
}
