package msgtype

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func CombineWithMsgType(
	genericMsg Variation,
	otherMsg string,
	reference interface{},
) string {
	if otherMsg == "" {
		return genericMsg.String() +
			getReferenceMessage(reference)
	}

	return genericMsg.String() +
		constants.Space +
		otherMsg +
		getReferenceMessage(reference)
}

func getReferenceMessage(
	reference interface{},
) string {
	if reference == nil {
		return ""
	}

	currentString, isString := reference.(string)
	if isString && currentString == "" {
		return ""
	}

	return fmt.Sprintf(
		ReferenceFormat,
		reference)
}
