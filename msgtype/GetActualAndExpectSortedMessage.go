package msgtype

import (
	"fmt"

	"gitlab.com/evatix-go/core/internal/msgformats"
)

func GetWhenActualAndExpectProcessedMessage(
	counter int,
	when interface{},
	actual interface{},
	expected interface{},
	actualSorted interface{},
	expectedSorted interface{},
) string {
	return fmt.Sprintf(
		msgformats.PrintWhenActualAndExpectedProcessedFormat,
		counter,
		when,
		actual,
		expected,
		actualSorted,
		expectedSorted,
	)
}
