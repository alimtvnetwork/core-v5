package errcore

import "fmt"

func FmtDebugIf(
	isDebug bool,
	format string,
	items ...any,
) {
	if !isDebug {
		return
	}

	fmt.Printf(format, items...)
}
