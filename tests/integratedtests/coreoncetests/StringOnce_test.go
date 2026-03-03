package coreoncetests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/coreonce"
	"gitlab.com/auk-go/core/errcore"
)

func Test_StringOnce_Core(t *testing.T) {
	for caseIndex, tc := range stringOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringOncePtr(func() string { return initVal })

		// Act
		actLines := []string{
			once.Value(),
			once.String(),
			fmt.Sprintf("%v", once.IsEmpty()),
			fmt.Sprintf("%v", once.IsEmptyOrWhitespace()),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines,
			fmt.Sprintf("  InitValue: %q", tc.InitValue),
		)
	}
}

func Test_StringOnce_Caching(t *testing.T) {
	for caseIndex, tc := range stringOnceCachingTestCases {
		// Arrange
		callCount := 0
		initVal := tc.InitValue
		once := coreonce.NewStringOncePtr(func() string {
			callCount++

			return initVal
		})

		// Act
		r1 := once.Value()
		r2 := once.Value()
		r3 := once.Value()

		actLines := []string{
			r1,
			r2,
			r3,
			fmt.Sprintf("%d", callCount),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines,
			fmt.Sprintf("  CallCount: %d", callCount),
		)
	}
}

func Test_StringOnce_Match(t *testing.T) {
	for caseIndex, tc := range stringOnceMatchTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringOncePtr(func() string { return initVal })

		// Act
		var actLines []string

		hasPrefix := tc.MatchArg == "prefix"
		hasSuffix := tc.MatchArg == "suffix"

		if hasPrefix {
			actLines = []string{
				fmt.Sprintf("%v", once.HasPrefix(tc.MatchArg)),
				fmt.Sprintf("%v", once.HasPrefix("data")),
			}
		} else if hasSuffix {
			actLines = []string{
				fmt.Sprintf("%v", once.HasSuffix(tc.MatchArg)),
				fmt.Sprintf("%v", once.HasSuffix("data")),
			}
		} else if tc.MatchArg == tc.InitValue {
			actLines = []string{
				fmt.Sprintf("%v", once.IsEqual(tc.MatchArg)),
				fmt.Sprintf("%v", once.IsEqual("xyz")),
			}
		} else {
			actLines = []string{
				fmt.Sprintf("%v", once.IsContains(tc.MatchArg)),
				fmt.Sprintf("%v", once.IsContains("xyz")),
			}
		}

		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines,
			fmt.Sprintf("  InitValue: %q, MatchArg: %q", tc.InitValue, tc.MatchArg),
		)
	}
}

func Test_StringOnce_Split(t *testing.T) {
	for caseIndex, tc := range stringOnceSplitTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringOncePtr(func() string { return initVal })

		// Act
		var actLines []string

		isSplitBy := tc.InitValue == "a,b,c"
		isTrimSplit := tc.InitValue == " key = value "

		if isSplitBy {
			parts := once.SplitBy(tc.Splitter)
			actLines = []string{
				fmt.Sprintf("%d", len(parts)),
				parts[0],
				parts[2],
			}
		} else if isTrimSplit {
			left, right := once.SplitLeftRightTrim(tc.Splitter)
			actLines = []string{left, right}
		} else {
			left, right := once.SplitLeftRight(tc.Splitter)
			actLines = []string{left, right}
		}

		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines,
			fmt.Sprintf("  InitValue: %q, Splitter: %q", tc.InitValue, tc.Splitter),
		)
	}
}

func Test_StringOnce_Json(t *testing.T) {
	for caseIndex, tc := range stringOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringOncePtr(func() string { return initVal })

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
			fmt.Sprintf("  InitValue: %q, Error: %v", tc.InitValue, err),
		)
	}
}
