package msgtype

import (
	"errors"
	"fmt"
)

type ExpectingRecord struct {
	ExpectingTitle string
	WasExpecting   interface{}
}

// Message
// Expecting
//
// returns
//      "%s - expecting (type:[%T]) : [\"%v\"], but received or actual (type:[%T]) : [\"%v\"]"
func (receiver *ExpectingRecord) Message(actual interface{}) string {
	return fmt.Sprintf(
		expectingMessageFormat,
		receiver.ExpectingTitle,
		receiver.WasExpecting, receiver.WasExpecting,
		actual, actual)
}

// Error
// Expecting
//
// returns
//      "%s - expecting (type:[%T]) : [\"%v\"], but received or actual (type:[%T]) : [\"%v\"]"
func (receiver *ExpectingRecord) Error(actual interface{}) error {
	return errors.New(receiver.Message(actual))
}
