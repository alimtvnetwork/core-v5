package coreoncetests

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/coreonce"
	"gitlab.com/auk-go/core/errcore"
)

func newErrorOnce(initError string) *coreonce.ErrorOnce {
	if initError == "" {
		return coreonce.NewErrorOncePtr(func() error { return nil })
	}

	if initError == "empty-marker" {
		return coreonce.NewErrorOncePtr(func() error { return errors.New("") })
	}

	return coreonce.NewErrorOncePtr(func() error { return errors.New(initError) })
}

func Test_ErrorOnce_Core(t *testing.T) {
	for caseIndex, tc := range errorOnceCoreTestCases {
		// Arrange
		once := newErrorOnce(tc.InitError)

		// Act
		actLines := []string{
			fmt.Sprintf("%v", once.HasError()),
			fmt.Sprintf("%v", once.IsValid()),
			fmt.Sprintf("%v", once.IsSuccess()),
			fmt.Sprintf("%v", once.IsEmpty()),
			fmt.Sprintf("%v", once.IsInvalid()),
			fmt.Sprintf("%v", once.IsFailed()),
			fmt.Sprintf("%v", once.HasAnyItem()),
			fmt.Sprintf("%v", once.IsDefined()),
			once.Message(),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Print diff on failure
		if errcore.LineDiffHasMismatch(actLines, expectedLines) {
			fmt.Printf(
				"\n=== ErrorOnce Core Diff (Case %d: %s) ===\n",
				caseIndex,
				tc.Case.Title,
			)
			fmt.Printf("  InitError: %q\n", tc.InitError)

			errcore.PrintLineDiff(caseIndex, tc.Case.Title, actLines, expectedLines)
			fmt.Println("=== End ===")
		}

		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_ErrorOnce_Caching(t *testing.T) {
	for caseIndex, tc := range errorOnceCachingTestCases {
		// Arrange
		callCount := 0
		initErr := tc.InitError
		once := coreonce.NewErrorOncePtr(func() error {
			callCount++

			return errors.New(initErr)
		})

		// Act
		r1 := once.Value()
		r2 := once.Value()
		r3 := once.Value()

		actLines := []string{
			r1.Error(),
			r2.Error(),
			r3.Error(),
			fmt.Sprintf("%d", callCount),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Print diff on failure
		if errcore.LineDiffHasMismatch(actLines, expectedLines) {
			fmt.Printf(
				"\n=== ErrorOnce Caching Diff (Case %d: %s) ===\n",
				caseIndex,
				tc.Case.Title,
			)
			fmt.Printf("  CallCount: %d\n", callCount)

			errcore.PrintLineDiff(caseIndex, tc.Case.Title, actLines, expectedLines)
			fmt.Println("=== End ===")
		}

		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_ErrorOnce_NullOrEmpty(t *testing.T) {
	for caseIndex, tc := range errorOnceNullOrEmptyTestCases {
		// Arrange
		once := newErrorOnce(tc.InitError)

		// Act
		actLines := []string{
			fmt.Sprintf("%v", once.IsNullOrEmpty()),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Print diff on failure
		if errcore.LineDiffHasMismatch(actLines, expectedLines) {
			fmt.Printf(
				"\n=== ErrorOnce NullOrEmpty Diff (Case %d: %s) ===\n",
				caseIndex,
				tc.Case.Title,
			)
			fmt.Printf("  InitError: %q\n", tc.InitError)

			errcore.PrintLineDiff(caseIndex, tc.Case.Title, actLines, expectedLines)
			fmt.Println("=== End ===")
		}

		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_ErrorOnce_MessageEqual(t *testing.T) {
	for caseIndex, tc := range errorOnceMessageEqualTestCases {
		// Arrange
		once := newErrorOnce(tc.InitError)

		// Act
		actLines := []string{
			fmt.Sprintf("%v", once.IsMessageEqual(tc.MatchMsg)),
			fmt.Sprintf("%v", once.IsMessageEqual("other")),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Print diff on failure
		if errcore.LineDiffHasMismatch(actLines, expectedLines) {
			fmt.Printf(
				"\n=== ErrorOnce MessageEqual Diff (Case %d: %s) ===\n",
				caseIndex,
				tc.Case.Title,
			)
			fmt.Printf("  InitError: %q, MatchMsg: %q\n", tc.InitError, tc.MatchMsg)

			errcore.PrintLineDiff(caseIndex, tc.Case.Title, actLines, expectedLines)
			fmt.Println("=== End ===")
		}

		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_ErrorOnce_ConcatNew(t *testing.T) {
	for caseIndex, tc := range errorOnceConcatTestCases {
		// Arrange
		once := newErrorOnce(tc.InitError)

		// Act
		result := once.ConcatNewString(tc.ExtraMsg)

		var actLines []string

		isNilError := tc.InitError == ""
		if isNilError {
			actLines = []string{result}
		} else {
			actLines = []string{
				fmt.Sprintf("%v", strings.Contains(result, tc.InitError)),
				fmt.Sprintf("%v", strings.Contains(result, tc.ExtraMsg)),
			}
		}

		expectedLines := tc.Case.ExpectedInput.([]string)

		// Print diff on failure
		if errcore.LineDiffHasMismatch(actLines, expectedLines) {
			fmt.Printf(
				"\n=== ErrorOnce ConcatNew Diff (Case %d: %s) ===\n",
				caseIndex,
				tc.Case.Title,
			)
			fmt.Printf("  InitError: %q, ExtraMsg: %q, Result: %q\n",
				tc.InitError, tc.ExtraMsg, result)

			errcore.PrintLineDiff(caseIndex, tc.Case.Title, actLines, expectedLines)
			fmt.Println("=== End ===")
		}

		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_ErrorOnce_Json(t *testing.T) {
	for caseIndex, tc := range errorOnceJsonTestCases {
		// Arrange
		once := newErrorOnce(tc.InitError)

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
				"\n=== ErrorOnce JSON Diff (Case %d: %s) ===\n",
				caseIndex,
				tc.Case.Title,
			)
			fmt.Printf("  InitError: %q, Error: %v\n", tc.InitError, err)

			errcore.PrintLineDiff(caseIndex, tc.Case.Title, actLines, expectedLines)
			fmt.Println("=== End ===")
		}

		// Assert
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
