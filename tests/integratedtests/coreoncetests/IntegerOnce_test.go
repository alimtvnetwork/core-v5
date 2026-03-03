package coreoncetests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coreonce"
	"gitlab.com/auk-go/core/errcore"
)

func Test_IntegerOnce_Core(t *testing.T) {
	for caseIndex, tc := range integerOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegerOncePtr(func() int { return initVal })

		// Act
		actLines := []string{
			fmt.Sprintf("%d", once.Value()),
			once.String(),
			fmt.Sprintf("%v", once.IsZero()),
			fmt.Sprintf("%v", once.IsEmpty()),
			fmt.Sprintf("%v", once.IsAboveZero()),
			fmt.Sprintf("%v", once.IsPositive()),
			fmt.Sprintf("%v", once.IsLessThanZero()),
			fmt.Sprintf("%v", once.IsNegative()),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines,
			fmt.Sprintf("  InitValue: %d", tc.InitValue),
		)
	}
}

func Test_IntegerOnce_Caching(t *testing.T) {
	for caseIndex, tc := range integerOnceCachingTestCases {
		// Arrange
		callCount := 0
		initVal := tc.InitValue
		once := coreonce.NewIntegerOncePtr(func() int {
			callCount++

			return initVal
		})

		// Act
		r1 := once.Value()
		r2 := once.Value()

		actLines := []string{
			fmt.Sprintf("%d", r1),
			fmt.Sprintf("%d", r2),
			fmt.Sprintf("%d", callCount),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines,
			fmt.Sprintf("  CallCount: %d", callCount),
		)
	}
}

func Test_IntegerOnce_Compare(t *testing.T) {
	for caseIndex, tc := range integerOnceCompareTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegerOncePtr(func() int { return initVal })
		cmpVal := tc.CompareValue

		// Act
		actLines := []string{
			fmt.Sprintf("%v", once.IsAbove(cmpVal)),
			fmt.Sprintf("%v", once.IsAbove(initVal)),
			fmt.Sprintf("%v", once.IsAboveEqual(initVal)),
		}

		isLessThanCase := tc.InitValue < tc.CompareValue
		if isLessThanCase {
			actLines = []string{
				fmt.Sprintf("%v", once.IsLessThan(cmpVal)),
				fmt.Sprintf("%v", once.IsLessThan(initVal)),
				fmt.Sprintf("%v", once.IsLessThanEqual(initVal)),
			}
		}

		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines,
			fmt.Sprintf("  InitValue: %d, CompareValue: %d", tc.InitValue, tc.CompareValue),
		)
	}
}

func Test_IntegerOnce_Json(t *testing.T) {
	for caseIndex, tc := range integerOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegerOncePtr(func() int { return initVal })

		// Act
		data, err := once.MarshalJSON()
		noError := err == nil

		actLines := []string{
			fmt.Sprintf("%v", noError),
			string(data),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines,
			fmt.Sprintf("  InitValue: %d, Error: %v", tc.InitValue, err),
		)
	}
}
