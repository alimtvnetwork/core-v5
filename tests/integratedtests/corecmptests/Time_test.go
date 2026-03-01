package corecmptests

import (
	"testing"
	"time"

	"gitlab.com/auk-go/core/corecmp"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

func Test_Time_Compare_Verification(t *testing.T) {
	for caseIndex, testCase := range timeCompareTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, ok := input["left"].(time.Time)
		if !ok {
			errcore.HandleErrMessage(nil, "left must be time.Time")
		}
		right, ok := input["right"].(time.Time)
		if !ok {
			errcore.HandleErrMessage(nil, "right must be time.Time")
		}

		// Act
		result := corecmp.Time(left, right)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			result.String(),
		)
	}
}
