package corerangestests

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

func Test_IntRanges_ValidCases(t *testing.T) {
	for _, testCase := range validIntRangeTestCases {
		// Arrange
		arrangeInputs := testCase.Arrange()
		first := arrangeInputs[0]
		rest := arrangeInputs[1:]

		// Act
		actualRanges := first.CreateRanges(rest...)

		// Assert
		convey.Convey(testCase.Title, t, func() {
			convey.So(actualRanges, should.Equal, testCase.ExpectedInput)
		})
	}
}
