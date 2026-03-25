package errcore

import "log/slog"

func PrintError(
	err error,
) {
	if err != nil {
		slog.Error("error occurred", "error", err)
	}
}
