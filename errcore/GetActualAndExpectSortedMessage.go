package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

func GetWhenActualAndExpectProcessedMessage(
	actual any,
	expectationMessageDef *ExpectationMessageDef,
) string {
	return fmt.Sprintf(
		msgformats.PrintWhenActualAndExpectedProcessedFormat,
		expectationMessageDef.CaseIndex,
		expectationMessageDef.When,
		expectationMessageDef.FuncName,
		actual,
		expectationMessageDef.Expected,
		expectationMessageDef.ActualProcessed,
		expectationMessageDef.ExpectedProcessed,
		expectationMessageDef.TestCaseName,
	)
}
