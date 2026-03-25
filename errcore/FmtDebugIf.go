package errcore

import (
	"fmt"
	"log/slog"
)

func FmtDebugIf(
	isDebug bool,
	format string,
	items ...any,
) {
	if !isDebug {
		return
	}

	slog.Debug(fmt.Sprintf(format, items...))
}
