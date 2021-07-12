package coretests

import (
	"gitlab.com/evatix-go/core/msgtype"
)

func GetActualAndExpectSortedMessage(
	counter int,
	when interface{},
	actual interface{},
	expected interface{},
	actualProcessed interface{},
	expectedProcessed interface{},
) string {
	return msgtype.GetWhenActualAndExpectProcessedMessage(
		counter,
		when,
		actual,
		expected,
		actualProcessed,
		expectedProcessed)
}
