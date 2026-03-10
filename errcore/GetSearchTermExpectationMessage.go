package errcore

import (
	"fmt"

	"github.com/alimtvnetwork/core/internal/msgformats"
)

func GetSearchTermExpectationMessage(
	counter int,
	header string,
	expectationMessage string,
	lineProcessingIndex int,
	actual any,
	expected any,
	additionalInfo any, // can be nil
) string {
	if additionalInfo == nil {
		return fmt.Sprintf(
			msgformats.PrintHeaderForSearchWithActualAndExpectedProcessedWithoutAdditionalFormat,
			counter,
			header,
			expectationMessage,
			lineProcessingIndex,
			actual,
			expected,
		)
	}

	return fmt.Sprintf(
		msgformats.PrintHeaderForSearchWithActualAndExpectedProcessedFormat,
		counter,
		header,
		expectationMessage,
		lineProcessingIndex,
		actual,
		expected,
		additionalInfo,
	)
}
