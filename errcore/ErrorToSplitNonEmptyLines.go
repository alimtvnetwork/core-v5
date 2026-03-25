package errcore

import "github.com/alimtvnetwork/core/internal/strutilinternal"

func ErrorToSplitNonEmptyLines(err error) []string {
	lines := ErrorToSplitLines(err)

	return strutilinternal.NonEmptySlice(lines)
}
