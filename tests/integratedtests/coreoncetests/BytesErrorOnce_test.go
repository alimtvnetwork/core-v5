package coreoncetests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/coreonce"
	"gitlab.com/auk-go/core/errcore"
)

// =============================================================================
// BytesErrorOnce — Core
// =============================================================================

func Test_BytesErrorOnce_Core(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceCoreTestCases {
		// Arrange
		initBytes := tc.InitBytes
		initErr := tc.InitErr
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			return initBytes, initErr
		})

		// Act
		val, err := once.Value()
		actLines := []string{
			string(val),
			fmt.Sprintf("%v", err == nil),
			fmt.Sprintf("%v", once.Length()),
			fmt.Sprintf("%v", once.HasAnyItem()),
			fmt.Sprintf("%v", once.IsEmpty()),
			fmt.Sprintf("%v", once.IsEmptyBytes()),
			fmt.Sprintf("%v", once.IsBytesEmpty()),
			fmt.Sprintf("%v", once.IsNull()),
			fmt.Sprintf("%v", once.IsDefined()),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines)
	}
}

// =============================================================================
// BytesErrorOnce — Caching
// =============================================================================

func Test_BytesErrorOnce_Caching(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceCachingTestCases {
		// Arrange
		initBytes := tc.InitBytes
		initErr := tc.InitErr

		if initErr != nil {
			once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
				return initBytes, initErr
			})

			// Act
			val, err := once.Value()
			actLines := []string{
				string(val),
				fmt.Sprintf("%v", val == nil),
				err.Error(),
			}
			expectedLines := tc.Case.ExpectedInput.([]string)

			// Assert
			errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines)

			continue
		}

		callCount := 0
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			callCount++

			return initBytes, nil
		})

		// Act
		r1, e1 := once.Value()
		r2, e2 := once.Value()

		actLines := []string{
			string(r1),
			string(r2),
			fmt.Sprintf("%v", e1 == nil),
			fmt.Sprintf("%v", e2 == nil),
			fmt.Sprintf("%d", callCount),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines,
			fmt.Sprintf("  CallCount: %d", callCount),
		)
	}
}

// =============================================================================
// BytesErrorOnce — Access (Execute, ValueOnly, ValueWithError)
// =============================================================================

func Test_BytesErrorOnce_Access(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceAccessTestCases {
		// Arrange
		initBytes := tc.InitBytes
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			return initBytes, nil
		})

		// Act
		var actLines []string

		switch caseIndex {
		case 0: // Execute == Value
			v1, _ := once.Execute()
			v2, _ := once.Value()
			actLines = []string{
				fmt.Sprintf("%v", string(v1) == string(v2)),
			}
		case 1: // ValueOnly
			actLines = []string{
				string(once.ValueOnly()),
			}
		case 2: // ValueWithError
			v, e := once.ValueWithError()
			actLines = []string{
				string(v),
				fmt.Sprintf("%v", e == nil),
			}
		}

		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines)
	}
}

// =============================================================================
// BytesErrorOnce — Error State
// =============================================================================

func Test_BytesErrorOnce_ErrorState(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceErrorStateTestCases {
		// Arrange
		initBytes := tc.InitBytes
		initErr := tc.InitErr
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			return initBytes, initErr
		})

		// Act
		actLines := []string{
			fmt.Sprintf("%v", once.HasError()),
			fmt.Sprintf("%v", once.IsEmptyError()),
			fmt.Sprintf("%v", once.IsValid()),
			fmt.Sprintf("%v", once.IsSuccess()),
			fmt.Sprintf("%v", once.IsInvalid()),
			fmt.Sprintf("%v", once.IsFailed()),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines)
	}
}

// =============================================================================
// BytesErrorOnce — HasIssues
// =============================================================================

func Test_BytesErrorOnce_HasIssues(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceHasIssuesTestCases {
		// Arrange
		var once *coreonce.BytesErrorOnce

		if tc.IsNilReceiver {
			once = nil
		} else {
			initBytes := tc.InitBytes
			initErr := tc.InitErr
			once = coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
				return initBytes, initErr
			})
		}

		// Act
		actLines := []string{
			fmt.Sprintf("%v", once.HasIssuesOrEmpty()),
			fmt.Sprintf("%v", once.HasSafeItems()),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines)
	}
}

// =============================================================================
// BytesErrorOnce — String
// =============================================================================

func Test_BytesErrorOnce_String(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceStringTestCases {
		// Arrange
		initBytes := tc.InitBytes
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			return initBytes, nil
		})

		// Act
		actLines := []string{
			once.String(),
			fmt.Sprintf("%v", once.IsStringEmpty()),
			fmt.Sprintf("%v", once.IsStringEmptyOrWhitespace()),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines)
	}
}

// =============================================================================
// BytesErrorOnce — Deserialize
// =============================================================================

func Test_BytesErrorOnce_Deserialize(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceDeserializeTestCases {
		// Arrange
		initJson := tc.InitJson
		initErr := tc.InitErr
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			return []byte(initJson), initErr
		})

		// Act
		var actLines []string

		if tc.IsMust {
			var result map[string]string
			panicked := callPanics(func() { once.DeserializeMust(&result) })
			actLines = []string{fmt.Sprintf("%v", panicked)}

			if !panicked {
				actLines = append(actLines, result["key"])
			}
		} else if initErr != nil {
			var result map[string]string
			err := once.Deserialize(&result)
			actLines = []string{
				fmt.Sprintf("%v", err != nil),
				fmt.Sprintf("%v", strings.Contains(err.Error(), "existing error cannot deserialize")),
				fmt.Sprintf("%v", strings.Contains(err.Error(), initErr.Error())),
			}
		} else if initJson == "not-json" {
			var result map[string]string
			err := once.Deserialize(&result)
			actLines = []string{
				fmt.Sprintf("%v", err != nil),
				fmt.Sprintf("%v", strings.Contains(err.Error(), "deserialize failed")),
			}
		} else {
			var result map[string]string
			err := once.Deserialize(&result)
			actLines = []string{
				fmt.Sprintf("%v", err == nil),
				result["name"],
			}
		}

		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines)
	}
}

// =============================================================================
// BytesErrorOnce — Serialization
// =============================================================================

func Test_BytesErrorOnce_Serialization(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceSerializationTestCases {
		// Arrange
		initBytes := tc.InitBytes
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			return initBytes, nil
		})

		// Act
		var actLines []string

		if caseIndex == 0 {
			data, err := once.MarshalJSON()
			actLines = []string{
				fmt.Sprintf("%v", err == nil),
				string(data),
			}
		} else {
			data, err := once.Serialize()
			actLines = []string{
				fmt.Sprintf("%v", err == nil),
				string(data),
			}
		}

		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines)
	}
}

// =============================================================================
// BytesErrorOnce — SerializeMust
// =============================================================================

func Test_BytesErrorOnce_SerializeMust(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceSerializeMustTestCases {
		// Arrange
		initBytes := tc.InitBytes
		initErr := tc.InitErr
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			return initBytes, initErr
		})

		// Act
		var result []byte
		panicked := callPanics(func() { result = once.SerializeMust() })
		actLines := []string{fmt.Sprintf("%v", panicked)}

		if !panicked {
			actLines = append(actLines, string(result))
		}

		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines)
	}
}

// =============================================================================
// BytesErrorOnce — Lifecycle (panic guards + IsInitialized)
// =============================================================================

func Test_BytesErrorOnce_Lifecycle(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceLifecycleTestCases {
		// Arrange
		initBytes := tc.InitBytes
		initErr := tc.InitErr
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			return initBytes, initErr
		})

		// Act
		actLines := []string{
			fmt.Sprintf("%v", callPanics(func() { once.HandleError() })),
			fmt.Sprintf("%v", callPanics(func() { once.MustBeEmptyError() })),
			fmt.Sprintf("%v", callPanics(func() { once.MustHaveSafeItems() })),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines)
	}
}

// =============================================================================
// BytesErrorOnce — IsInitialized
// =============================================================================

func Test_BytesErrorOnce_IsInitialized(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceInitializedTestCases {
		// Arrange
		initBytes := tc.InitBytes
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			return initBytes, nil
		})

		// Act
		beforeInit := once.IsInitialized()
		_, _ = once.Value()
		afterInit := once.IsInitialized()

		actLines := []string{
			fmt.Sprintf("%v", beforeInit),
			fmt.Sprintf("%v", afterInit),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines)
	}
}

// =============================================================================
// BytesErrorOnce — Constructor
// =============================================================================

func Test_BytesErrorOnce_Constructor(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceConstructorTestCases {
		// Arrange
		initBytes := tc.InitBytes
		once := coreonce.NewBytesErrorOnce(func() ([]byte, error) {
			return initBytes, nil
		})

		// Act
		v, e := once.Value()
		actLines := []string{
			string(v),
			fmt.Sprintf("%v", e == nil),
		}
		expectedLines := tc.Case.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, expectedLines)
	}
}
