package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================================================
// Test: Pair — IsEqual extended edge cases
// ==========================================================================

func Test_Pair_IsEqual_SameValuesDiffValidity(t *testing.T) {
	tc := pairIsEqualSameValuesDiffValidityTestCase
	a := coregeneric.NewPair("x", "y")
	b := coregeneric.NewPairWithMessage("x", "y", false, "")

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", a.IsEqual(b)))
}

func Test_Pair_IsEqual_DiffRight(t *testing.T) {
	tc := pairIsEqualDiffRightTestCase
	a := coregeneric.NewPair("x", "y")
	b := coregeneric.NewPair("x", "z")

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", a.IsEqual(b)))
}

func Test_Pair_IsEqual_BothInvalidZero(t *testing.T) {
	tc := pairIsEqualBothInvalidZeroTestCase
	a := coregeneric.InvalidPairNoMessage[string, string]()
	b := coregeneric.InvalidPairNoMessage[string, string]()

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", a.IsEqual(b)))
}

func Test_Pair_IsEqual_IntSame(t *testing.T) {
	tc := pairIsEqualIntSameTestCase
	a := coregeneric.NewPair(10, 20)
	b := coregeneric.NewPair(10, 20)

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", a.IsEqual(b)))
}

func Test_Pair_IsEqual_IntDiff(t *testing.T) {
	tc := pairIsEqualIntDiffTestCase
	a := coregeneric.NewPair(10, 20)
	b := coregeneric.NewPair(10, 30)

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", a.IsEqual(b)))
}

func Test_Pair_IsEqual_MixedTypes(t *testing.T) {
	tc := pairIsEqualMixedTypesTestCase
	a := coregeneric.NewPair("key", 42)
	b := coregeneric.NewPair("key", 42)

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", a.IsEqual(b)))
}

// ==========================================================================
// Test: Pair — HasMessage edge cases
// ==========================================================================

func Test_Pair_HasMessage_ValidNoMsg(t *testing.T) {
	tc := pairHasMessageValidNoMsgTestCase
	p := coregeneric.NewPair("a", "b")

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", p.HasMessage()))
}

func Test_Pair_HasMessage_InvalidWithMsg(t *testing.T) {
	tc := pairHasMessageInvalidWithMsgTestCase
	p := coregeneric.InvalidPair[string, string]("error")

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", p.HasMessage()))
}

func Test_Pair_HasMessage_Whitespace(t *testing.T) {
	tc := pairHasMessageWhitespaceTestCase
	p := coregeneric.NewPairWithMessage("a", "b", true, "   ")

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", p.HasMessage()))
}

func Test_Pair_HasMessage_Nil(t *testing.T) {
	tc := pairHasMessageNilTestCase
	var p *coregeneric.Pair[string, string]

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", p.HasMessage()))
}

// ==========================================================================
// Test: Pair — IsInvalid edge cases
// ==========================================================================

func Test_Pair_IsInvalid_Valid(t *testing.T) {
	tc := pairIsInvalidValidTestCase
	p := coregeneric.NewPair("a", "b")

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", p.IsInvalid()))
}

func Test_Pair_IsInvalid_Invalid(t *testing.T) {
	tc := pairIsInvalidInvalidTestCase
	p := coregeneric.InvalidPairNoMessage[string, string]()

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", p.IsInvalid()))
}

func Test_Pair_IsInvalid_Nil(t *testing.T) {
	tc := pairIsInvalidNilTestCase
	var p *coregeneric.Pair[string, string]

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", p.IsInvalid()))
}

// ==========================================================================
// Test: Pair — String output
// ==========================================================================

func Test_Pair_String_Valid(t *testing.T) {
	tc := pairStringValidTestCase
	p := coregeneric.NewPair("hello", "world")

	tc.ShouldBeEqualFirst(t, p.String())
}

func Test_Pair_String_InvalidZero(t *testing.T) {
	tc := pairStringInvalidZeroTestCase
	p := coregeneric.InvalidPairNoMessage[string, string]()

	tc.ShouldBeEqualFirst(t, p.String())
}

func Test_Pair_String_Nil(t *testing.T) {
	tc := pairStringNilTestCase
	var p *coregeneric.Pair[string, string]

	tc.ShouldBeEqualFirst(t, p.String())
}

func Test_Pair_String_MixedType(t *testing.T) {
	tc := pairStringMixedTypeTestCase
	p := coregeneric.NewPair("key", 42)

	tc.ShouldBeEqualFirst(t, p.String())
}

// ==========================================================================
// Test: Triple — IsEqual extended edge cases
// ==========================================================================

func Test_Triple_IsEqual_Same(t *testing.T) {
	tc := tripleIsEqualSameTestCase
	a := coregeneric.NewTriple("a", "b", "c")
	b := coregeneric.NewTriple("a", "b", "c")

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", a.IsEqual(b)))
}

func Test_Triple_IsEqual_DiffValidity(t *testing.T) {
	tc := tripleIsEqualDiffValidityTestCase
	a := coregeneric.NewTriple("a", "b", "c")
	b := coregeneric.NewTripleWithMessage("a", "b", "c", false, "")

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", a.IsEqual(b)))
}

func Test_Triple_IsEqual_DiffMiddle(t *testing.T) {
	tc := tripleIsEqualDiffMiddleTestCase
	a := coregeneric.NewTriple("a", "b", "c")
	b := coregeneric.NewTriple("a", "X", "c")

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", a.IsEqual(b)))
}

func Test_Triple_IsEqual_BothNil(t *testing.T) {
	tc := tripleIsEqualBothNilTestCase
	var a *coregeneric.Triple[string, string, string]
	var b *coregeneric.Triple[string, string, string]

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", a.IsEqual(b)))
}

func Test_Triple_IsEqual_NilVsNonNil(t *testing.T) {
	tc := tripleIsEqualNilVsNonNilTestCase
	var a *coregeneric.Triple[string, string, string]
	b := coregeneric.NewTriple("a", "b", "c")

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", a.IsEqual(b)))
}

// ==========================================================================
// Test: Triple — HasMessage edge cases
// ==========================================================================

func Test_Triple_HasMessage_ValidNoMsg(t *testing.T) {
	tc := tripleHasMessageValidNoMsgTestCase
	tr := coregeneric.NewTriple("a", "b", "c")

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", tr.HasMessage()))
}

func Test_Triple_HasMessage_InvalidWithMsg(t *testing.T) {
	tc := tripleHasMessageInvalidWithMsgTestCase
	tr := coregeneric.InvalidTriple[string, string, string]("err")

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", tr.HasMessage()))
}

func Test_Triple_HasMessage_Nil(t *testing.T) {
	tc := tripleHasMessageNilTestCase
	var tr *coregeneric.Triple[string, string, string]

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", tr.HasMessage()))
}

// ==========================================================================
// Test: Triple — IsInvalid edge cases
// ==========================================================================

func Test_Triple_IsInvalid_Valid(t *testing.T) {
	tc := tripleIsInvalidValidTestCase
	tr := coregeneric.NewTriple("a", "b", "c")

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", tr.IsInvalid()))
}

func Test_Triple_IsInvalid_Invalid(t *testing.T) {
	tc := tripleIsInvalidInvalidTestCase
	tr := coregeneric.InvalidTripleNoMessage[string, string, string]()

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", tr.IsInvalid()))
}

func Test_Triple_IsInvalid_Nil(t *testing.T) {
	tc := tripleIsInvalidNilTestCase
	var tr *coregeneric.Triple[string, string, string]

	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%v", tr.IsInvalid()))
}

// ==========================================================================
// Test: Triple — String output
// ==========================================================================

func Test_Triple_String_Valid(t *testing.T) {
	tc := tripleStringValidTestCase
	tr := coregeneric.NewTriple("a", "b", "c")

	tc.ShouldBeEqualFirst(t, tr.String())
}

func Test_Triple_String_InvalidZero(t *testing.T) {
	tc := tripleStringInvalidZeroTestCase
	tr := coregeneric.InvalidTripleNoMessage[string, string, string]()

	tc.ShouldBeEqualFirst(t, tr.String())
}

func Test_Triple_String_Nil(t *testing.T) {
	tc := tripleStringNilTestCase
	var tr *coregeneric.Triple[string, string, string]

	tc.ShouldBeEqualFirst(t, tr.String())
}

// ==========================================================================
// Test: Pair — NewPairWithMessage
// ==========================================================================

func Test_Pair_NewPairWithMessage_Valid(t *testing.T) {
	tc := pairWithMessageValidTestCase
	p := coregeneric.NewPairWithMessage("hello", "world", true, "ok")

	// Act
	actual := args.Map{
		"left":         p.Left,
		"right":        p.Right,
		"isValid":      p.IsValid,
		"errorMessage": p.Message,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

func Test_Pair_NewPairWithMessage_Invalid(t *testing.T) {
	tc := pairWithMessageInvalidTestCase
	p := coregeneric.NewPairWithMessage("", "", false, "failed")

	// Act
	actual := args.Map{
		"left":         p.Left,
		"right":        p.Right,
		"isValid":      p.IsValid,
		"errorMessage": p.Message,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

// ==========================================================================
// Test: Triple — NewTripleWithMessage
// ==========================================================================

func Test_Triple_NewTripleWithMessage_Valid(t *testing.T) {
	tc := tripleWithMessageValidTestCase
	tr := coregeneric.NewTripleWithMessage("a", "b", "c", true, "success")

	// Act
	actual := args.Map{
		"left":         tr.Left,
		"middle":       tr.Middle,
		"right":        tr.Right,
		"isValid":      tr.IsValid,
		"errorMessage": tr.Message,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

func Test_Triple_NewTripleWithMessage_Invalid(t *testing.T) {
	tc := tripleWithMessageInvalidTestCase
	tr := coregeneric.NewTripleWithMessage("", "", "", false, "error occurred")

	// Act
	actual := args.Map{
		"left":         tr.Left,
		"middle":       tr.Middle,
		"right":        tr.Right,
		"isValid":      tr.IsValid,
		"errorMessage": tr.Message,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

// ==========================================================================
// Test: Pair — Dispose
// ==========================================================================

func Test_Pair_Dispose(t *testing.T) {
	tc := pairDisposeTestCase
	p := coregeneric.NewPairWithMessage("a", "b", true, "msg")
	p.Dispose()

	// Act
	actual := args.Map{
		"left":         p.Left,
		"right":        p.Right,
		"isValid":      p.IsValid,
		"errorMessage": p.Message,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

// ==========================================================================
// Test: Triple — Dispose
// ==========================================================================

func Test_Triple_Dispose(t *testing.T) {
	tc := tripleDisposeTestCase
	tr := coregeneric.NewTripleWithMessage("a", "b", "c", true, "msg")
	tr.Dispose()

	// Act
	actual := args.Map{
		"left":         tr.Left,
		"middle":       tr.Middle,
		"right":        tr.Right,
		"isValid":      tr.IsValid,
		"errorMessage": tr.Message,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

// ==========================================================================
// Test: All typed Pair creator shortcuts
// ==========================================================================

// ==========================================================================
// Test: All typed Pair creator shortcuts
// ==========================================================================

func Test_New_Pair_Creator_AllShortcuts(t *testing.T) {
	// StringInt64
	{
		p := coregeneric.New.Pair.StringInt64("k", int64(99))
		if p.Left != "k" || p.Right != int64(99) || !p.IsValid {
			t.Errorf("New.Pair.StringInt64 failed: got %v", p)
		}
	}

	// StringFloat64
	{
		p := coregeneric.New.Pair.StringFloat64("pi", 3.14)
		if p.Left != "pi" || p.Right != 3.14 || !p.IsValid {
			t.Errorf("New.Pair.StringFloat64 failed: got %v", p)
		}
	}

	// StringBool
	{
		p := coregeneric.New.Pair.StringBool("flag", true)
		if p.Left != "flag" || p.Right != true || !p.IsValid {
			t.Errorf("New.Pair.StringBool failed: got %v", p)
		}
	}

	// StringAny
	{
		p := coregeneric.New.Pair.StringAny("key", []int{1, 2})
		if p.Left != "key" || !p.IsValid {
			t.Errorf("New.Pair.StringAny failed: got %v", p)
		}
	}

	// IntInt
	{
		p := coregeneric.New.Pair.IntInt(1, 2)
		if p.Left != 1 || p.Right != 2 || !p.IsValid {
			t.Errorf("New.Pair.IntInt failed: got %v", p)
		}
	}

	// IntString
	{
		p := coregeneric.New.Pair.IntString(42, "answer")
		if p.Left != 42 || p.Right != "answer" || !p.IsValid {
			t.Errorf("New.Pair.IntString failed: got %v", p)
		}
	}

	// InvalidAny
	{
		p := coregeneric.New.Pair.InvalidAny("bad")
		if p.IsValid || p.Message != "bad" {
			t.Errorf("New.Pair.InvalidAny failed: got %v", p)
		}
	}
}

// ==========================================================================
// Test: All typed Triple creator shortcuts
// ==========================================================================

func Test_New_Triple_Creator_AllShortcuts(t *testing.T) {
	// StringIntString
	{
		tr := coregeneric.New.Triple.StringIntString("left", 42, "right")
		if tr.Left != "left" || tr.Middle != 42 || tr.Right != "right" || !tr.IsValid {
			t.Errorf("New.Triple.StringIntString failed: got %v", tr)
		}
	}

	// StringAnyAny
	{
		tr := coregeneric.New.Triple.StringAnyAny("key", 3.14, true)
		if tr.Left != "key" || !tr.IsValid {
			t.Errorf("New.Triple.StringAnyAny failed: got %v", tr)
		}
	}

	// InvalidAny
	{
		tr := coregeneric.New.Triple.InvalidAny("err")
		if tr.IsValid || tr.Message != "err" {
			t.Errorf("New.Triple.InvalidAny failed: got %v", tr)
		}
	}
}

// ==========================================================================
// Test: Pair — nil receiver Clear (no panic)
// ==========================================================================

func Test_Pair_Nil_Clear_NoPanic(t *testing.T) {
	var p *coregeneric.Pair[string, string]
	p.Clear() // should not panic
}

// ==========================================================================
// Test: Triple — nil receiver Clear (no panic)
// ==========================================================================

func Test_Triple_Nil_Clear_NoPanic(t *testing.T) {
	var tr *coregeneric.Triple[string, string, string]
	tr.Clear() // should not panic
}
