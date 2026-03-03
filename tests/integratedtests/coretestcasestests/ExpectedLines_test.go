package coretestcasestests

import (
	"testing"

	"gitlab.com/auk-go/core/errcore"
)

func Test_CaseV1_ExpectedLines_Int(t *testing.T) {
	tc := expectedLinesIntTestCase

	// Arrange — ExpectedInput is int(42)

	// Act
	actLines := tc.ExpectedLines()

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, expectedLinesIntExpected)
}

func Test_CaseV1_ExpectedLines_BoolTrue(t *testing.T) {
	tc := expectedLinesBoolTrueTestCase

	// Arrange — ExpectedInput is true

	// Act
	actLines := tc.ExpectedLines()

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, expectedLinesBoolTrueExpected)
}

func Test_CaseV1_ExpectedLines_BoolFalse(t *testing.T) {
	tc := expectedLinesBoolFalseTestCase

	// Arrange — ExpectedInput is false

	// Act
	actLines := tc.ExpectedLines()

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, expectedLinesBoolFalseExpected)
}

func Test_CaseV1_ExpectedLines_IntSlice(t *testing.T) {
	tc := expectedLinesIntSliceTestCase

	// Arrange — ExpectedInput is []int{10, 20, 30}

	// Act
	actLines := tc.ExpectedLines()

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, expectedLinesIntSliceExpected)
}

func Test_CaseV1_ExpectedLines_BoolSlice(t *testing.T) {
	tc := expectedLinesBoolSliceTestCase

	// Arrange — ExpectedInput is []bool{true, false, true}

	// Act
	actLines := tc.ExpectedLines()

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, expectedLinesBoolSliceExpected)
}

func Test_CaseV1_ExpectedLines_String(t *testing.T) {
	tc := expectedLinesStringTestCase

	// Arrange — ExpectedInput is "hello"

	// Act
	actLines := tc.ExpectedLines()

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, expectedLinesStringExpected)
}

func Test_CaseV1_ExpectedLines_StringSlice(t *testing.T) {
	tc := expectedLinesStringSliceTestCase

	// Arrange — ExpectedInput is []string{"a", "b", "c"}

	// Act
	actLines := tc.ExpectedLines()

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, expectedLinesStringSliceExpected)
}

func Test_CaseV1_ExpectedLines_MapStringInt(t *testing.T) {
	tc := expectedLinesMapStringIntTestCase

	// Arrange — ExpectedInput is map[string]int{"age": 30, "count": 5}

	// Act
	actLines := tc.ExpectedLines()

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, expectedLinesMapStringIntExpected)
}
