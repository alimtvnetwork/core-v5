package errcore

import (
	"fmt"

	"github.com/alimtvnetwork/core/internal/msgformats"
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
