package isanytests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
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
		actual := args.Map{
			"isDeepEqual":    fmt.Sprintf("%v", isany.DeepEqual(pair.left, pair.right)),
			"isNotDeepEqual": fmt.Sprintf("%v", isany.NotDeepEqual(pair.left, pair.right)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Zero
// ==========================================

func Test_Zero_Verification(t *testing.T) {
	for caseIndex, testCase := range zeroTestCases {
		// Act
		actual := args.Map{
			"intZero":   fmt.Sprintf("%v", isany.Zero(0)),
			"int42":     fmt.Sprintf("%v", isany.Zero(42)),
			"emptyStr":  fmt.Sprintf("%v", isany.Zero("")),
			"helloStr":  fmt.Sprintf("%v", isany.Zero("hello")),
			"boolFalse": fmt.Sprintf("%v", isany.Zero(false)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"nilPtr":    fmt.Sprintf("%v", isany.ReflectNull(nilPtr)),
			"nonNilPtr": fmt.Sprintf("%v", isany.ReflectNull(nonNilPtr)),
			"nilSlice":  fmt.Sprintf("%v", isany.ReflectNull(nilSlice)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"notNullNil":    fmt.Sprintf("%v", isany.NotNull(nil)),
			"notNull42":     fmt.Sprintf("%v", isany.NotNull(42)),
			"inverseEquals": fmt.Sprintf("%v", isany.NotNull(nilPtr) == isany.Null(nilPtr)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: StringEqual
// ==========================================

func Test_StringEqual_Verification(t *testing.T) {
	for caseIndex, testCase := range stringEqualTestCases {
		// Act
		actual := args.Map{
			"sameStrings": fmt.Sprintf("%v", isany.StringEqual("hello", "hello")),
			"diffStrings": fmt.Sprintf("%v", isany.StringEqual("hello", "world")),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"intPtr":    fmt.Sprintf("%v", isany.Pointer(&x)),
			"intVal":    fmt.Sprintf("%v", isany.Pointer(x)),
			"stringPtr": fmt.Sprintf("%v", isany.Pointer(&s)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
