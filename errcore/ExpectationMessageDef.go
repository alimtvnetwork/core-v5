package errcore

import (
	"fmt"
	"log/slog"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type ExpectationMessageDef struct {
	CaseIndex         int
	FuncName          string
	TestCaseName      string
	When              string
	Expected          any
	ActualProcessed   any
	ExpectedProcessed any
	IsNonWhiteSort    bool
	expectedString    *string
}

func (it ExpectationMessageDef) ExpectedSafeString() string {
	if it.expectedString != nil {
		return *it.expectedString
	}

	var expectedStr string

	if it.Expected != nil {
		expectedStr = fmt.Sprintf(
			constants.SprintValueFormat,
			it.Expected)
	}

	it.expectedString = &expectedStr

	return *it.expectedString
}

func (it ExpectationMessageDef) ExpectedStringTrim() string {
	return strings.TrimSpace(it.ExpectedString())
}

func (it ExpectationMessageDef) ExpectedString() string {
	if it.Expected == nil {
		panic("ExpectationMessageDef! Expected needs to be set before getting it!")
	}

	return it.ExpectedSafeString()
}

func (it ExpectationMessageDef) ToString(actual any) string {
	return GetWhenActualAndExpectProcessedMessage(
		actual,
		&it)
}

func (it ExpectationMessageDef) PrintIf(
	isPrint bool,
	actual any,
) {
	if !isPrint {
		return
	}

	it.Print(actual)
}

func (it ExpectationMessageDef) PrintIfFailed(
	isPrintOnFail,
	isFailed bool,
	actual any,
) {
	if isPrintOnFail && isFailed {
		it.Print(actual)
	}
}

func (it ExpectationMessageDef) Print(actual any) {
	msg := MsgHeaderPlusEnding(
		it.When,
		it.ToString(actual))

	slog.Warn("expectation", "message", msg)
}
