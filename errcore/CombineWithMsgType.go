package errcore

import (
	"gitlab.com/auk-go/core/constants"
)

func CombineWithMsgTypeStackTrace(
	genericMsg RawErrorType,
	otherMsg string,
	reference interface{},
) string {
	msg := CombineWithMsgTypeNoStack(
		genericMsg,
		otherMsg,
		reference,
	)

	return StackEnhance.MsgSkip(1, msg)
}

func CombineWithMsgTypeNoStack(
	genericMsg RawErrorType,
	otherMsg string,
	reference interface{},
) string {
	if otherMsg == "" {
		return genericMsg.String() +
			getReferenceMessage(reference)
	}

	msg := genericMsg.String() +
		constants.Space +
		otherMsg +
		getReferenceMessage(reference)

	return msg
}
