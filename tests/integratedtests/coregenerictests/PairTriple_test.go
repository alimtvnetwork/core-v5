package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================
// Test: Pair — NewPair valid
// ==========================================

func Test_Pair_NewPair_Valid(t *testing.T) {
	for caseIndex, testCase := range pairNewValidTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, leftErr := input.GetAsString("left")
		errcore.HandleErrMessage("left", leftErr)
		right, rightErr := input.GetAsString("right")
		errcore.HandleErrMessage("right", rightErr)

		// Act
		pair := coregeneric.NewPair(left, right)
		actLines := []string{
			pair.Left,
			pair.Right,
			fmt.Sprintf("%v", pair.IsValid),
			pair.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Pair — InvalidPair
// ==========================================

func Test_Pair_InvalidPair(t *testing.T) {
	for caseIndex, testCase := range pairInvalidTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		message, _ := input.GetAsString("message")

		// Act
		var pair *coregeneric.Pair[string, string]
		if message == "" {
			pair = coregeneric.InvalidPairNoMessage[string, string]()
		} else {
			pair = coregeneric.InvalidPair[string, string](message)
		}

		actLines := []string{
			pair.Left,
			pair.Right,
			fmt.Sprintf("%v", pair.IsValid),
			pair.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Pair — Clone independence
// ==========================================

func Test_Pair_Clone_Independence(t *testing.T) {
	for caseIndex, testCase := range pairCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		right, _ := input.GetAsString("right")

		// Act
		original := coregeneric.NewPair(left, right)
		cloned := original.Clone()
		cloned.Left = "mutated-left"

		actLines := []string{
			original.Left,
			original.Right,
			fmt.Sprintf("%v", original.IsValid),
			cloned.Left,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Pair — nil Clone
// ==========================================

func Test_Pair_Nil_Clone(t *testing.T) {
	for caseIndex, testCase := range pairNilCloneTestCases {
		// Arrange — nil pair

		// Act
		var pair *coregeneric.Pair[string, string]
		cloned := pair.Clone()

		actLines := []string{
			fmt.Sprintf("%v", cloned == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Pair — IsEqual
// ==========================================

func Test_Pair_IsEqual(t *testing.T) {
	// Case 0: Equal pairs
	{
		testCase := pairIsEqualTestCases[0]
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		right, _ := input.GetAsString("right")

		a := coregeneric.NewPair(left, right)
		b := coregeneric.NewPair(left, right)
		actLines := []string{fmt.Sprintf("%v", a.IsEqual(b))}
		testCase.ShouldBeEqual(t, 0, actLines...)
	}

	// Case 1: Unequal - different left
	{
		testCase := pairIsEqualTestCases[1]
		a := coregeneric.NewPair("a", "b")
		b := coregeneric.NewPair("x", "b")
		actLines := []string{fmt.Sprintf("%v", a.IsEqual(b))}
		testCase.ShouldBeEqual(t, 1, actLines...)
	}

	// Case 2: Nil vs non-nil
	{
		testCase := pairIsEqualTestCases[2]
		a := coregeneric.NewPair("a", "b")
		var b *coregeneric.Pair[string, string]
		actLines := []string{fmt.Sprintf("%v", a.IsEqual(b))}
		testCase.ShouldBeEqual(t, 2, actLines...)
	}

	// Case 3: Both nil
	{
		testCase := pairIsEqualTestCases[3]
		var a *coregeneric.Pair[string, string]
		var b *coregeneric.Pair[string, string]
		actLines := []string{fmt.Sprintf("%v", a.IsEqual(b))}
		testCase.ShouldBeEqual(t, 3, actLines...)
	}
}

// ==========================================
// Test: Pair — Values()
// ==========================================

func Test_Pair_Values(t *testing.T) {
	for caseIndex, testCase := range pairValuesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		right, _ := input.GetAsString("right")

		// Act
		pair := coregeneric.NewPair(left, right)
		l, r := pair.Values()
		actLines := []string{l, r}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Pair — Clear
// ==========================================

func Test_Pair_Clear(t *testing.T) {
	for caseIndex, testCase := range pairClearTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		right, _ := input.GetAsString("right")

		// Act
		pair := coregeneric.NewPair(left, right)
		pair.Clear()
		actLines := []string{
			pair.Left,
			pair.Right,
			fmt.Sprintf("%v", pair.IsValid),
			pair.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Triple — NewTriple valid
// ==========================================

func Test_Triple_NewTriple_Valid(t *testing.T) {
	for caseIndex, testCase := range tripleNewValidTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		middle, _ := input.GetAsString("middle")
		right, _ := input.GetAsString("right")

		// Act
		triple := coregeneric.NewTriple(left, middle, right)
		actLines := []string{
			triple.Left,
			triple.Middle,
			triple.Right,
			fmt.Sprintf("%v", triple.IsValid),
			triple.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Triple — InvalidTriple
// ==========================================

func Test_Triple_InvalidTriple(t *testing.T) {
	for caseIndex, testCase := range tripleInvalidTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		message, _ := input.GetAsString("message")

		// Act
		var triple *coregeneric.Triple[string, string, string]
		if message == "" {
			triple = coregeneric.InvalidTripleNoMessage[string, string, string]()
		} else {
			triple = coregeneric.InvalidTriple[string, string, string](message)
		}

		actLines := []string{
			triple.Left,
			triple.Middle,
			triple.Right,
			fmt.Sprintf("%v", triple.IsValid),
			triple.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Triple — Clone independence
// ==========================================

func Test_Triple_Clone_Independence(t *testing.T) {
	for caseIndex, testCase := range tripleCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		middle, _ := input.GetAsString("middle")
		right, _ := input.GetAsString("right")

		// Act
		original := coregeneric.NewTriple(left, middle, right)
		cloned := original.Clone()
		cloned.Left = "mutated"

		actLines := []string{
			original.Left,
			original.Middle,
			original.Right,
			fmt.Sprintf("%v", original.IsValid),
			cloned.Left,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Triple — nil Clone
// ==========================================

func Test_Triple_Nil_Clone(t *testing.T) {
	for caseIndex, testCase := range tripleNilCloneTestCases {
		// Act
		var triple *coregeneric.Triple[string, string, string]
		cloned := triple.Clone()

		actLines := []string{
			fmt.Sprintf("%v", cloned == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Triple — Values()
// ==========================================

func Test_Triple_Values(t *testing.T) {
	for caseIndex, testCase := range tripleValuesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		middle, _ := input.GetAsString("middle")
		right, _ := input.GetAsString("right")

		// Act
		triple := coregeneric.NewTriple(left, middle, right)
		a, b, c := triple.Values()
		actLines := []string{a, b, c}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Triple — Clear
// ==========================================

func Test_Triple_Clear(t *testing.T) {
	for caseIndex, testCase := range tripleClearTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		middle, _ := input.GetAsString("middle")
		right, _ := input.GetAsString("right")

		// Act
		triple := coregeneric.NewTriple(left, middle, right)
		triple.Clear()
		actLines := []string{
			triple.Left,
			triple.Middle,
			triple.Right,
			fmt.Sprintf("%v", triple.IsValid),
			triple.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: New.Pair Creator shortcuts
// ==========================================

func Test_New_Pair_Creator(t *testing.T) {
	// StringString
	p := coregeneric.New.Pair.StringString("k", "v")
	if p.Left != "k" || p.Right != "v" || !p.IsValid {
		t.Errorf("New.Pair.StringString failed: got %v", p)
	}

	// StringInt
	pi := coregeneric.New.Pair.StringInt("age", 30)
	if pi.Left != "age" || pi.Right != 30 || !pi.IsValid {
		t.Errorf("New.Pair.StringInt failed: got %v", pi)
	}

	// Any
	pa := coregeneric.New.Pair.Any("x", 42)
	if pa.Left != "x" || pa.Right != 42 || !pa.IsValid {
		t.Errorf("New.Pair.Any failed: got %v", pa)
	}

	// InvalidStringString
	inv := coregeneric.New.Pair.InvalidStringString("err")
	if inv.IsValid || inv.Message != "err" {
		t.Errorf("New.Pair.InvalidStringString failed: got %v", inv)
	}
}

// ==========================================
// Test: New.Triple Creator shortcuts
// ==========================================

func Test_New_Triple_Creator(t *testing.T) {
	// StringStringString
	tr := coregeneric.New.Triple.StringStringString("a", "b", "c")
	if tr.Left != "a" || tr.Middle != "b" || tr.Right != "c" || !tr.IsValid {
		t.Errorf("New.Triple.StringStringString failed: got %v", tr)
	}

	// Any
	ta := coregeneric.New.Triple.Any("x", 1, true)
	if ta.Left != "x" || ta.Middle != 1 || ta.Right != true || !ta.IsValid {
		t.Errorf("New.Triple.Any failed: got %v", ta)
	}

	// InvalidStringStringString
	inv := coregeneric.New.Triple.InvalidStringStringString("bad")
	if inv.IsValid || inv.Message != "bad" {
		t.Errorf("New.Triple.InvalidStringStringString failed: got %v", inv)
	}
}
