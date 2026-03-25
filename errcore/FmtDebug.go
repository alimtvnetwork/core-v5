package errcore

import (
	"fmt"
	"log/slog"
)

func FmtDebug(
	format string,
	items ...any,
) {
	slog.Debug(fmt.Sprintf(format, items...))
}

func ValidPrint(
	isValid bool,
	items ...any,
) {
	if isValid {
		slog.Info("valid", "values", fmt.Sprint(items...))
	}
}

func FailedPrint(
	isFailed bool,
	items ...any,
) {
	if isFailed {
		slog.Warn("failed", "values", fmt.Sprint(items...))
	}
}
