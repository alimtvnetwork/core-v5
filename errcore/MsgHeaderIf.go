package errcore

import (
	"fmt"
)

func MsgHeaderIf(
	isHeader bool,
	items ...any,
) string {
	if isHeader {
		return MsgHeader(items...)
	}

	return fmt.Sprint(items...)
}
