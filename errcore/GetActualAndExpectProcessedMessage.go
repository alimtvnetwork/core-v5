package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

func GetActualAndExpectProcessedMessage(
	counter int,
	actual any,
	expected any,
	actualProcessed any,
	expectedProcessed any,
) string {
	return fmt.Sprintf(
		msgformats.PrintActualAndExpectedProcessedFormat,
		counter,
		actual,
		expected,
		actualProcessed,
		expectedProcessed,
	)
}
