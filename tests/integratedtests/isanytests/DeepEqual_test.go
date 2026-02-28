package isanytests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/isany"
)

// ==========================================
// Test: DeepEqual / NotDeepEqual
// ==========================================

func Test_DeepEqual_Verification(t *testing.T) {
	type testPair struct {
		left, right any
	}

	pairs := []testPair{
		{42, 42},
		{42, 99},
		{[]string{"a", "b"}, []string{"a", "b"}},
		{[]string{"a", "b"}, []string{"a", "c"}},
		{nil, nil},
	}

	for caseIndex, testCase := range deepEqualTestCases {
		// Arrange
		pair := pairs[caseIndex]

		// Act
		actLines := []string{
			fmt.Sprintf("%v", isany.DeepEqual(pair.left, pair.right)),
			fmt.Sprintf("%v", isany.NotDeepEqual(pair.left, pair.right)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Zero
// ==========================================

func Test_Zero_Verification(t *testing.T) {
	for caseIndex, testCase := range zeroTestCases {
		// Arrange
		values := []any{
			0,
			42,
			"",
			"hello",
			false,
		}

		// Act
		actLines := make([]string, 0, len(values))
		for _, v := range values {
			actLines = append(actLines, fmt.Sprintf("%v", isany.Zero(v)))
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ReflectNull
// ==========================================

func Test_ReflectNull_Verification(t *testing.T) {
	for caseIndex, testCase := range reflectNullTestCases {
		// Arrange
		var nilPtr *string
		nonNilPtr := new(string)
		var nilSlice []string

		// Act
		actLines := []string{
			fmt.Sprintf("%v", isany.ReflectNull(nilPtr)),
			fmt.Sprintf("%v", isany.ReflectNull(nonNilPtr)),
			fmt.Sprintf("%v", isany.ReflectNull(nilSlice)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: NotNull
// ==========================================

func Test_NotNull_Verification(t *testing.T) {
	for caseIndex, testCase := range notNullTestCases {
		// Arrange
		var nilPtr *string

		// Act
		actLines := []string{
			fmt.Sprintf("%v", isany.NotNull(nil)),
			fmt.Sprintf("%v", isany.NotNull(42)),
			fmt.Sprintf("%v", isany.NotNull(nilPtr) == isany.Null(nilPtr)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: StringEqual
// ==========================================

func Test_StringEqual_Verification(t *testing.T) {
	for caseIndex, testCase := range stringEqualTestCases {
		// Arrange — compare values by their string representation

		// Act
		actLines := []string{
			fmt.Sprintf("%v", isany.StringEqual("hello", "hello")),
			fmt.Sprintf("%v", isany.StringEqual("hello", "world")),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Pointer
// ==========================================

func Test_Pointer_Verification(t *testing.T) {
	for caseIndex, testCase := range pointerTestCases {
		// Arrange
		x := 42
		s := "hello"

		// Act
		actLines := []string{
			fmt.Sprintf("%v", isany.Pointer(&x)),
			fmt.Sprintf("%v", isany.Pointer(x)),
			fmt.Sprintf("%v", isany.Pointer(&s)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
