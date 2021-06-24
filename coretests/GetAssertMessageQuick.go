package coretests

import "fmt"

func GetAssertMessageQuick(
	when,
	actual,
	expected interface{},
	counter int,
) string {
	return fmt.Sprintf(quickActualExpectedMessageFormat,
		counter,
		when,
		actual,
		expected,
	)
}
