package coretests

import "fmt"

func PrintFailedExpectedSortedMessage(
	isFailed bool,
	when,
	actual, actualSorted,
	expected, expectedSorted interface{},
	counter int,
) {
	if isFailed {
		message :=
			GetActualAndExpectSortedMessage(
				counter,
				when,
				actual,
				expected,
				actualSorted,
				expectedSorted)
		fmt.Println(message)
	}
}
