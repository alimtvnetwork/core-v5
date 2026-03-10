package errcore

import (
	"fmt"

	"github.com/alimtvnetwork/core/internal/msgformats"
)

func MsgHeader(
	items ...any,
) string {
	return fmt.Sprintf(
		msgformats.MsgHeaderFormat,
		items...)
}
