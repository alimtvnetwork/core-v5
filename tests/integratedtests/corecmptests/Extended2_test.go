package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_IsStringsEqualPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range isStringsEqualPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftNil := input.GetAsBoolDefault("leftNil", false)
		rightNil := input.GetAsBoolDefault("rightNil", false)

		var left, right []string
		if !leftNil {
			rawLeft, _ := input.Get("left")
			left = rawLeft.([]string)
		}
		if !rightNil {
			rawRight, _ := input.Get("right")
			right = rawRight.([]string)
		}

		// Act
		result := corecmp.IsStringsEqualPtr(left, right)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TimePtr_Verification(t *testing.T) {
	for caseIndex, testCase := range timePtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftNil := input.GetAsBoolDefault("leftNil", false)
		rightNil := input.GetAsBoolDefault("rightNil", false)

		var left, right *time.Time
		now := time.Now()

		if !leftNil {
			left = &now
		}
		if !rightNil {
			sameTime := input.GetAsBoolDefault("sameTime", false)
			if sameTime {
				right = &now
			} else {
				later := now.Add(time.Hour)
				right = &later
			}
		}

		// Act
		result := corecmp.TimePtr(left, right)

		actual := args.Map{
			"isEqual": result == corecomparator.Equal,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
