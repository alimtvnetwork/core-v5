package coreoncetests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coreonce"
	"gitlab.com/auk-go/core/errcore"
)

func Test_BoolOnce_Core(t *testing.T) {
	for caseIndex, tc := range boolOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewBoolOncePtr(func() bool { return initVal })

		// Act
		val := once.Value()
		str := once.String()

		actLines := []string{
			fmt.Sprintf("%v", val),
			str,
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.PrintDiffOnMismatch(caseIndex, tc.Case.Title, actLines, expectedLines,
			fmt.Sprintf("  InitValue: %v", tc.InitValue),
		)
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_BoolOnce_Caching(t *testing.T) {
	for caseIndex, tc := range boolOnceCachingTestCases {
		// Arrange
		callCount := 0
		initVal := tc.InitValue
		once := coreonce.NewBoolOncePtr(func() bool {
			callCount++

			return initVal
		})

		// Act
		r1 := once.Value()
		r2 := once.Value()
		r3 := once.Value()

		actLines := []string{
			fmt.Sprintf("%v", r1),
			fmt.Sprintf("%v", r2),
			fmt.Sprintf("%v", r3),
			fmt.Sprintf("%d", callCount),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.PrintDiffOnMismatch(caseIndex, tc.Case.Title, actLines, expectedLines,
			fmt.Sprintf("  InitValue: %v, CallCount: %d", tc.InitValue, callCount),
		)
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_BoolOnce_Json(t *testing.T) {
	for caseIndex, tc := range boolOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewBoolOncePtr(func() bool { return initVal })

		// Act
		data, err := once.MarshalJSON()
		noError := err == nil

		actLines := []string{
			fmt.Sprintf("%v", noError),
			string(data),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.PrintDiffOnMismatch(caseIndex, tc.Case.Title, actLines, expectedLines,
			fmt.Sprintf("  InitValue: %v, Error: %v", tc.InitValue, err),
		)
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
