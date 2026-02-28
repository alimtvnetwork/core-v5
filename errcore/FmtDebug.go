package errcore

import "fmt"

func FmtDebug(
	format string,
	items ...any,
) {
	fmt.Printf(format, items...)
}

func ValidPrint(
	isValid bool,
	items ...any,
) {
	if isValid {
		fmt.Print(items...)
	}
}

func FailedPrint(
	isFailed bool,
	items ...any,
) {
	if isFailed {
		fmt.Print(items...)
	}
}
