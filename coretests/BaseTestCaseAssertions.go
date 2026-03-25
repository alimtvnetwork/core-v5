package coretests

import (
	"errors"
	"fmt"
	"log/slog"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// =============================================================================
// Assertion Methods
// =============================================================================

// ShouldBe
//
// Disabled testcases will not be executed.
func (it *BaseTestCase) ShouldBe(
	caseIndex int,
	t *testing.T,
	assert convey.Assertion,
	actual any,
) {
	if it.IsEnable.IsFalse() {
		it.noPrintAssert(caseIndex, t, assert, actual)

		return
	}

	it.ShouldBeExplicit(
		true,
		caseIndex,
		t,
		it.Title,
		actual,
		assert,
		it.Expected(),
	)
}

func (it *BaseTestCase) noPrintAssert(
	caseIndex int,
	t *testing.T,
	assert convey.Assertion,
	actual any,
) {
	toTile := it.FormTitle(caseIndex)

	it.SetActual(actual)

	convey.Convey(
		toTile, t, func() {
			convey.SoMsg(
				toTile,
				actual,
				assert,
				it.ExpectedInput,
			)
		},
	)
}

func (it *BaseTestCase) ShouldBeExplicit(
	isValidateType bool,
	caseIndex int,
	t *testing.T,
	title string,
	actual any,
	assert convey.Assertion,
	expected any,
) {
	if it.IsEnable.IsFalse() {
		it.noPrintAssert(caseIndex, t, assert, actual)

		return
	}

	it.SetActual(actual)
	headerTitle := it.CustomTitle(caseIndex, title)
	actualLines := GetAssert.ToStrings(actual)
	expectedLines := GetAssert.ToStrings(expected)
	compare := assert(actualLines, expectedLines)
	isFailed := compare != ""

	convey.Convey(
		headerTitle, t, func() {
			if isFailed {
				toString := it.LinesString(caseIndex)

				slog.Warn("test case mismatch", "caseIndex", caseIndex, "details", toString)
			}

			convey.SoMsg(
				headerTitle,
				actualLines,
				assert,
				expectedLines,
			)
		},
	)

	isSkipTypeValidation := !isValidateType

	if isSkipTypeValidation {
		return
	}

	it.TypeShouldMatch(t, caseIndex, title)
}

func (it *BaseTestCase) TypeShouldMatch(
	t *testing.T,
	caseIndex int,
	title string,
) {
	err := it.TypeValidationError()

	if err == nil {
		return
	}

	errHeader := fmt.Sprintf(
		"%d : %s - type verification failed",
		caseIndex,
		title,
	)

	var finalError error

	if err != nil {
		finalError = errors.New(errHeader + err.Error())
	}

	convey.Convey(
		errHeader, t, func() {
			convey.So(
				finalError,
				convey.ShouldBeNil,
			)
		},
	)
}
