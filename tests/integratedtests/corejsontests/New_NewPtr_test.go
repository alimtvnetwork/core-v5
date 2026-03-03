package corejsontests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/corejson"
)

// ==========================================================================
// Test: New - valid
// ==========================================================================

func Test_New_Valid(t *testing.T) {
	tc := newValidTestCase
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

	// Assert
	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: New - nil
// ==========================================================================

func Test_New_Nil(t *testing.T) {
	tc := newNilTestCase
	result := corejson.New(nil)

	actLines := []string{
		fmt.Sprintf("%v", result.HasError()),
		string(result.Bytes),
	}

	// Assert
	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: New - channel
// ==========================================================================

func Test_New_Channel(t *testing.T) {
	tc := newChannelTestCase
	result := corejson.New(make(chan int))

	actLines := []string{
		fmt.Sprintf("%v", result.HasError()),
		fmt.Sprintf("%v", strings.Contains(result.Error.Error(), "marshal")),
	}

	// Assert
	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: NewPtr - valid
// ==========================================================================

func Test_NewPtr_Valid(t *testing.T) {
	tc := newPtrValidTestCase
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

	// Assert
	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: NewPtr - nil
// ==========================================================================

func Test_NewPtr_Nil(t *testing.T) {
	tc := newPtrNilTestCase
	result := corejson.NewPtr(nil)

	actLines := []string{
		fmt.Sprintf("%v", result != nil),
		fmt.Sprintf("%v", result.HasError()),
		string(result.Bytes),
	}

	// Assert
	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: NewPtr - channel
// ==========================================================================

func Test_NewPtr_Channel(t *testing.T) {
	tc := newPtrChannelTestCase
	result := corejson.NewPtr(make(chan string))

	actLines := []string{
		fmt.Sprintf("%v", result != nil),
		fmt.Sprintf("%v", result.HasError()),
		fmt.Sprintf("%v", strings.Contains(result.Error.Error(), "marshal")),
	}

	// Assert
	tc.ShouldBeEqual(t, 0, actLines...)
}
