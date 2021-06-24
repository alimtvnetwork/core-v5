package coretests

import "fmt"

func GetActualAndExpectSortedMessage(
	counter int,
	when interface{},
	actual interface{},
	expected interface{},
	actualSorted interface{},
	expectedSorted interface{},
) string {
	return fmt.Sprintf(printValuesForActualAndSortedFormat,
		counter,
		when,
		actual,
		expected,
		actualSorted,
		expectedSorted,
	)
}
