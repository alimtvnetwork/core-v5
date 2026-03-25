package simplewrap

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

// MsgWrapMsg
//
// Format : constants.StringWrapValueFormat
func MsgWrapMsg(msg, wrappedMsg string) string {
	if msg == "" && wrappedMsg == "" {
		return ""
	}

	if msg == "" && wrappedMsg != "" {
		return wrappedMsg
	}

	if msg != "" && wrappedMsg == "" {
		return msg
	}

	return fmt.Sprintf(
		constants.StringWrapValueFormat,
		msg,
		wrappedMsg)
}
