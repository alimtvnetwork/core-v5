package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

func GetSearchTermExpectationSimpleMessage(
	counter int,
	expectationErrorMessage string,
	processingIndex int,
	contentProcessed any,
	searchTermProcessed any,
) string {
	return fmt.Sprintf(
		msgformats.PrintHeaderForSearchActualAndExpectedProcessedSimpleFormat,
		counter,
		expectationErrorMessage,
		processingIndex,
		contentProcessed,
		searchTermProcessed,
	)
}
