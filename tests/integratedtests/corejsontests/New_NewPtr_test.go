package corejsontests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================================================
// Test: New - valid
// ==========================================================================

func Test_New_Valid(t *testing.T) {
	tc := newValidTestCases[0]
	result := corejson.New(struct {
		Name string
		Age  int
	}{Name: "Alice", Age: 30})

	actLines := []string{
		fmt.Sprintf("%v", result.HasError()),
		fmt.Sprintf("%v", result.IsEmpty()),
		fmt.Sprintf("%v", len(result.Bytes) > 0),
		fmt.Sprintf("%v", result.TypeName != ""),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: New - nil
// ==========================================================================

func Test_New_Nil(t *testing.T) {
	tc := newNilTestCases[0]
	result := corejson.New(nil)

	actLines := []string{
		fmt.Sprintf("%v", result.HasError()),
		string(result.Bytes),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: New - channel
// ==========================================================================

func Test_New_Channel(t *testing.T) {
	tc := newChannelTestCases[0]
	result := corejson.New(make(chan int))

	actLines := []string{
		fmt.Sprintf("%v", result.HasError()),
		fmt.Sprintf("%v", strings.Contains(result.Error.Error(), "marshal")),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: NewPtr - valid
// ==========================================================================

func Test_NewPtr_Valid(t *testing.T) {
	tc := newPtrValidTestCases[0]
	result := corejson.NewPtr(struct {
		Name string
		Age  int
	}{Name: "Bob", Age: 25})

	actLines := []string{
		fmt.Sprintf("%v", result != nil),
		fmt.Sprintf("%v", result.HasError()),
		fmt.Sprintf("%v", result.IsEmpty()),
		fmt.Sprintf("%v", len(result.Bytes) > 0),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: NewPtr - nil
// ==========================================================================

func Test_NewPtr_Nil(t *testing.T) {
	tc := newPtrNilTestCases[0]
	result := corejson.NewPtr(nil)

	actLines := []string{
		fmt.Sprintf("%v", result != nil),
		fmt.Sprintf("%v", result.HasError()),
		string(result.Bytes),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: NewPtr - channel
// ==========================================================================

func Test_NewPtr_Channel(t *testing.T) {
	tc := newPtrChannelTestCases[0]
	result := corejson.NewPtr(make(chan string))

	actLines := []string{
		fmt.Sprintf("%v", result != nil),
		fmt.Sprintf("%v", result.HasError()),
		fmt.Sprintf("%v", strings.Contains(result.Error.Error(), "marshal")),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}
