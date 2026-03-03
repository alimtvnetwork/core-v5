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

		// Print diff on failure
		if errcore.LineDiffHasMismatch(actLines, expectedLines) {
			fmt.Printf(
				"\n=== BoolOnce Core Diff (Case %d: %s) ===\n",
				caseIndex,
				tc.Case.Title,
			)
			fmt.Printf("  InitValue: %v\n", tc.InitValue)

			errcore.PrintLineDiff(caseIndex, tc.Case.Title, actLines, expectedLines)
			fmt.Println("=== End ===")
		}

		// Assert
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

		// Print diff on failure
		if errcore.LineDiffHasMismatch(actLines, expectedLines) {
			fmt.Printf(
				"\n=== BoolOnce Caching Diff (Case %d: %s) ===\n",
				caseIndex,
				tc.Case.Title,
			)
			fmt.Printf("  InitValue: %v, CallCount: %d\n", tc.InitValue, callCount)

			errcore.PrintLineDiff(caseIndex, tc.Case.Title, actLines, expectedLines)
			fmt.Println("=== End ===")
		}

		// Assert
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

		// Print diff on failure
		if errcore.LineDiffHasMismatch(actLines, expectedLines) {
			fmt.Printf(
				"\n=== BoolOnce JSON Diff (Case %d: %s) ===\n",
				caseIndex,
				tc.Case.Title,
			)
			fmt.Printf("  InitValue: %v, Error: %v\n", tc.InitValue, err)

			errcore.PrintLineDiff(caseIndex, tc.Case.Title, actLines, expectedLines)
			fmt.Println("=== End ===")
		}

		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
