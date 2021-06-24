package coretests

import "fmt"

func PrintFailedExpected(
	isFailed bool,
	when,
	actual,
	expected interface{},
	counter int,
) {
	if isFailed {
		message := GetAssertMessageQuick(when, actual, expected, counter)

		fmt.Println(message)
	}
}
