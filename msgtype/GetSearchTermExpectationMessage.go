package msgtype

import (
	"fmt"

	"gitlab.com/evatix-go/core/internal/msgformats"
)

func GetSearchTermExpectationMessage(
	counter int,
	expectationMessage string,
	contentProcessed interface{},
	searchTermProcessed interface{},
	additionalInfo interface{}, // can be nil
) string {
	if additionalInfo == nil {
		return fmt.Sprintf(
			msgformats.PrintHeaderForSearchWithActualAndExpectedProcessedWithoutAdditionalFormat,
			counter,
			expectationMessage,
			contentProcessed,
			searchTermProcessed,
		)
	}

	return fmt.Sprintf(
		msgformats.PrintHeaderForSearchWithActualAndExpectedProcessedFormat,
		counter,
		expectationMessage,
		contentProcessed,
		searchTermProcessed,
		additionalInfo,
	)
}
