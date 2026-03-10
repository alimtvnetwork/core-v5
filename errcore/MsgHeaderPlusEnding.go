package errcore

import (
	"fmt"

	"github.com/alimtvnetwork/core/internal/msgformats"
)

func MsgHeaderPlusEnding(
	header, message any,
) string {
	return fmt.Sprintf(
		msgformats.MsgHeaderPlusEndingFormat,
		header,
		message)
}
