package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

func Test_Time_Compare_Verification(t *testing.T) {
	for caseIndex, testCase := range timeCompareTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, ok := input["left"].(time.Time)
		if !ok {
			errcore.HandleErrMessage("left must be time.Time")
		}
		right, ok := input["right"].(time.Time)
		if !ok {
			errcore.HandleErrMessage("right must be time.Time")
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
