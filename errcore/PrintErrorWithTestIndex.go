package errcore

import "log/slog"

func PrintErrorWithTestIndex(
	caseIndex int,
	header string,
	err error,
) {
	if err != nil {
		slog.Error("test error",
			"caseIndex", caseIndex,
			"title", header,
			"summary", err,
		)
	}
}
