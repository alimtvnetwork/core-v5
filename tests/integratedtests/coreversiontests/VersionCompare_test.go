package coreversiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreversion"
)

func Test_Version_Compare_Verification(t *testing.T) {
	for caseIndex, testCase := range versionCompareTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftStr, _ := input.GetAsString("left")
		rightStr, _ := input.GetAsString("right")

		// Act
		leftV := coreversion.New.Create(leftStr)
		rightV := coreversion.New.Create(rightStr)
		result := leftV.Compare(&rightV)

		// Assert
		actual := args.Map{
			"result": result.Name(),
		}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
