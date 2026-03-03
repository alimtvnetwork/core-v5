package coreoncetests

import (
	"encoding/json"
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coreonce"
	"gitlab.com/auk-go/core/errcore"
)

func Test_BytesOnce_Core(t *testing.T) {
	for caseIndex, tc := range bytesOnceCoreTestCases {
		// Arrange
		var once *coreonce.BytesOnce

		if tc.UseNilInit {
			once = &coreonce.BytesOnce{}
		} else {
			initBytes := tc.InitBytes
			once = coreonce.NewBytesOncePtr(func() []byte { return initBytes })
		}

		// Act
		val := once.Value()
		actLines := []string{
			string(val),
			once.String(),
			fmt.Sprintf("%v", once.IsEmpty()),
			fmt.Sprintf("%v", once.Length()),
			fmt.Sprintf("%v", val == nil),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines)
	}
}

func Test_BytesOnce_Caching(t *testing.T) {
	for caseIndex, tc := range bytesOnceCachingTestCases {
		// Arrange
		callCount := 0
		initBytes := tc.InitBytes
		once := coreonce.NewBytesOncePtr(func() []byte {
			callCount++

			return initBytes
		})

		// Act
		r1 := once.Value()
		r2 := once.Value()
		r3 := once.Value()

		actLines := []string{
			string(r1),
			string(r2),
			string(r3),
			fmt.Sprintf("%d", callCount),
			fmt.Sprintf("%v", string(once.Execute()) == string(once.Value())),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines,
			fmt.Sprintf("  CallCount: %d", callCount),
		)
	}
}

func Test_BytesOnce_JSON(t *testing.T) {
	for caseIndex, tc := range bytesOnceJsonTestCases {
		// Arrange
		initBytes := tc.InitBytes
		once := coreonce.NewBytesOncePtr(func() []byte { return initBytes })

		// Act
		var actLines []string

		if tc.ReplaceBytes != nil {
			input, _ := json.Marshal(tc.ReplaceBytes)
			err := once.UnmarshalJSON(input)

			actLines = []string{
				fmt.Sprintf("%v", err == nil),
				string(once.Value()),
			}
		} else {
			data, err := once.MarshalJSON()

			actLines = []string{
				fmt.Sprintf("%v", err == nil),
				fmt.Sprintf("%v", len(data) > 0),
			}
		}

		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines)
	}
}

func Test_BytesOnce_Constructor(t *testing.T) {
	for caseIndex, tc := range bytesOnceConstructorTestCases {
		// Arrange
		initBytes := tc.InitBytes
		once := coreonce.NewBytesOnce(func() []byte { return initBytes })

		// Act
		actLines := []string{
			string(once.Value()),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines)
	}
}
