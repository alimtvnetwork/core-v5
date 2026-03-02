package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
)

// ==========================================================================
// Test: Pair — IsEqual extended edge cases
// ==========================================================================

func Test_Pair_IsEqual_Extended(t *testing.T) {
	// Case 0: same values, different validity
	{
		tc := pairIsEqualExtendedTestCases[0]
		a := coregeneric.NewPair("x", "y")
		b := coregeneric.NewPairWithMessage("x", "y", false, "")
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEqual(b)))
	}

	// Case 1: different right values
	{
		tc := pairIsEqualExtendedTestCases[1]
		a := coregeneric.NewPair("x", "y")
		b := coregeneric.NewPair("x", "z")
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%v", a.IsEqual(b)))
	}

	// Case 2: both invalid with same zero values
	{
		tc := pairIsEqualExtendedTestCases[2]
		a := coregeneric.InvalidPairNoMessage[string, string]()
		b := coregeneric.InvalidPairNoMessage[string, string]()
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%v", a.IsEqual(b)))
	}

	// Case 3: Pair[int,int] same values
	{
		tc := pairIsEqualExtendedTestCases[3]
		a := coregeneric.NewPair(10, 20)
		b := coregeneric.NewPair(10, 20)
		tc.ShouldBeEqual(t, 3, fmt.Sprintf("%v", a.IsEqual(b)))
	}

	// Case 4: Pair[int,int] different values
	{
		tc := pairIsEqualExtendedTestCases[4]
		a := coregeneric.NewPair(10, 20)
		b := coregeneric.NewPair(10, 30)
		tc.ShouldBeEqual(t, 4, fmt.Sprintf("%v", a.IsEqual(b)))
	}

	// Case 5: Pair[string,int] mixed types
	{
		tc := pairIsEqualExtendedTestCases[5]
		a := coregeneric.NewPair("key", 42)
		b := coregeneric.NewPair("key", 42)
		tc.ShouldBeEqual(t, 5, fmt.Sprintf("%v", a.IsEqual(b)))
	}
}

// ==========================================================================
// Test: Pair — HasMessage edge cases
// ==========================================================================

func Test_Pair_HasMessage(t *testing.T) {
	// Case 0: valid pair, no message
	{
		tc := pairHasMessageTestCases[0]
		p := coregeneric.NewPair("a", "b")
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", p.HasMessage()))
	}

	// Case 1: invalid pair with message
	{
		tc := pairHasMessageTestCases[1]
		p := coregeneric.InvalidPair[string, string]("error")
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%v", p.HasMessage()))
	}

	// Case 2: whitespace-only message
	{
		tc := pairHasMessageTestCases[2]
		p := coregeneric.NewPairWithMessage("a", "b", true, "   ")
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%v", p.HasMessage()))
	}

	// Case 3: nil pair
	{
		tc := pairHasMessageTestCases[3]
		var p *coregeneric.Pair[string, string]
		tc.ShouldBeEqual(t, 3, fmt.Sprintf("%v", p.HasMessage()))
	}
}

// ==========================================================================
// Test: Pair — IsInvalid edge cases
// ==========================================================================

func Test_Pair_IsInvalid(t *testing.T) {
	// Case 0: valid pair
	{
		tc := pairIsInvalidTestCases[0]
		p := coregeneric.NewPair("a", "b")
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", p.IsInvalid()))
	}

	// Case 1: invalid pair
	{
		tc := pairIsInvalidTestCases[1]
		p := coregeneric.InvalidPairNoMessage[string, string]()
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%v", p.IsInvalid()))
	}

	// Case 2: nil pair
	{
		tc := pairIsInvalidTestCases[2]
		var p *coregeneric.Pair[string, string]
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%v", p.IsInvalid()))
	}
}

// ==========================================================================
// Test: Pair — String output
// ==========================================================================

func Test_Pair_String(t *testing.T) {
	// Case 0: valid Pair[string,string]
	{
		tc := pairStringTestCases[0]
		p := coregeneric.NewPair("hello", "world")
		tc.ShouldBeEqual(t, 0, p.String())
	}

	// Case 1: invalid with zero values
	{
		tc := pairStringTestCases[1]
		p := coregeneric.InvalidPairNoMessage[string, string]()
		tc.ShouldBeEqual(t, 1, p.String())
	}

	// Case 2: nil pair
	{
		tc := pairStringTestCases[2]
		var p *coregeneric.Pair[string, string]
		tc.ShouldBeEqual(t, 2, p.String())
	}

	// Case 3: Pair[string,int]
	{
		tc := pairStringTestCases[3]
		p := coregeneric.NewPair("key", 42)
		tc.ShouldBeEqual(t, 3, p.String())
	}
}

// ==========================================================================
// Test: Triple — IsEqual extended edge cases
// ==========================================================================

func Test_Triple_IsEqual_Extended(t *testing.T) {
	// Case 0: same values same validity
	{
		tc := tripleIsEqualExtendedTestCases[0]
		a := coregeneric.NewTriple("a", "b", "c")
		b := coregeneric.NewTriple("a", "b", "c")
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEqual(b)))
	}

	// Case 1: same values different validity
	{
		tc := tripleIsEqualExtendedTestCases[1]
		a := coregeneric.NewTriple("a", "b", "c")
		b := coregeneric.NewTripleWithMessage("a", "b", "c", false, "")
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%v", a.IsEqual(b)))
	}

	// Case 2: different middle
	{
		tc := tripleIsEqualExtendedTestCases[2]
		a := coregeneric.NewTriple("a", "b", "c")
		b := coregeneric.NewTriple("a", "X", "c")
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%v", a.IsEqual(b)))
	}

	// Case 3: both nil
	{
		tc := tripleIsEqualExtendedTestCases[3]
		var a *coregeneric.Triple[string, string, string]
		var b *coregeneric.Triple[string, string, string]
		tc.ShouldBeEqual(t, 3, fmt.Sprintf("%v", a.IsEqual(b)))
	}

	// Case 4: nil vs non-nil
	{
		tc := tripleIsEqualExtendedTestCases[4]
		var a *coregeneric.Triple[string, string, string]
		b := coregeneric.NewTriple("a", "b", "c")
		tc.ShouldBeEqual(t, 4, fmt.Sprintf("%v", a.IsEqual(b)))
	}
}

// ==========================================================================
// Test: Triple — HasMessage edge cases
// ==========================================================================

func Test_Triple_HasMessage(t *testing.T) {
	// Case 0: valid, no message
	{
		tc := tripleHasMessageTestCases[0]
		tr := coregeneric.NewTriple("a", "b", "c")
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tr.HasMessage()))
	}

	// Case 1: invalid with message
	{
		tc := tripleHasMessageTestCases[1]
		tr := coregeneric.InvalidTriple[string, string, string]("err")
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%v", tr.HasMessage()))
	}

	// Case 2: nil triple
	{
		tc := tripleHasMessageTestCases[2]
		var tr *coregeneric.Triple[string, string, string]
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%v", tr.HasMessage()))
	}
}

// ==========================================================================
// Test: Triple — IsInvalid edge cases
// ==========================================================================

func Test_Triple_IsInvalid(t *testing.T) {
	// Case 0: valid
	{
		tc := tripleIsInvalidTestCases[0]
		tr := coregeneric.NewTriple("a", "b", "c")
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tr.IsInvalid()))
	}

	// Case 1: invalid
	{
		tc := tripleIsInvalidTestCases[1]
		tr := coregeneric.InvalidTripleNoMessage[string, string, string]()
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%v", tr.IsInvalid()))
	}

	// Case 2: nil
	{
		tc := tripleIsInvalidTestCases[2]
		var tr *coregeneric.Triple[string, string, string]
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%v", tr.IsInvalid()))
	}
}

// ==========================================================================
// Test: Triple — String output
// ==========================================================================

func Test_Triple_String(t *testing.T) {
	// Case 0: valid
	{
		tc := tripleStringTestCases[0]
		tr := coregeneric.NewTriple("a", "b", "c")
		tc.ShouldBeEqual(t, 0, tr.String())
	}

	// Case 1: invalid zero values
	{
		tc := tripleStringTestCases[1]
		tr := coregeneric.InvalidTripleNoMessage[string, string, string]()
		tc.ShouldBeEqual(t, 1, tr.String())
	}

	// Case 2: nil
	{
		tc := tripleStringTestCases[2]
		var tr *coregeneric.Triple[string, string, string]
		tc.ShouldBeEqual(t, 2, tr.String())
	}
}

// ==========================================================================
// Test: Pair — NewPairWithMessage
// ==========================================================================

func Test_Pair_NewPairWithMessage(t *testing.T) {
	// Case 0: valid with message
	{
		tc := pairWithMessageTestCases[0]
		p := coregeneric.NewPairWithMessage("hello", "world", true, "ok")
		tc.ShouldBeEqual(t, 0, p.Left, p.Right, fmt.Sprintf("%v", p.IsValid), p.Message)
	}

	// Case 1: invalid with error message
	{
		tc := pairWithMessageTestCases[1]
		p := coregeneric.NewPairWithMessage("", "", false, "failed")
		tc.ShouldBeEqual(t, 1, p.Left, p.Right, fmt.Sprintf("%v", p.IsValid), p.Message)
	}
}

// ==========================================================================
// Test: Triple — NewTripleWithMessage
// ==========================================================================

func Test_Triple_NewTripleWithMessage(t *testing.T) {
	// Case 0: valid with message
	{
		tc := tripleWithMessageTestCases[0]
		tr := coregeneric.NewTripleWithMessage("a", "b", "c", true, "success")
		tc.ShouldBeEqual(t, 0, tr.Left, tr.Middle, tr.Right, fmt.Sprintf("%v", tr.IsValid), tr.Message)
	}

	// Case 1: invalid with error
	{
		tc := tripleWithMessageTestCases[1]
		tr := coregeneric.NewTripleWithMessage("", "", "", false, "error occurred")
		tc.ShouldBeEqual(t, 1, tr.Left, tr.Middle, tr.Right, fmt.Sprintf("%v", tr.IsValid), tr.Message)
	}
}

// ==========================================================================
// Test: Pair — Dispose
// ==========================================================================

func Test_Pair_Dispose(t *testing.T) {
	tc := pairDisposeTestCases[0]
	p := coregeneric.NewPairWithMessage("a", "b", true, "msg")
	p.Dispose()
	tc.ShouldBeEqual(t, 0, p.Left, p.Right, fmt.Sprintf("%v", p.IsValid), p.Message)
}

// ==========================================================================
// Test: Triple — Dispose
// ==========================================================================

func Test_Triple_Dispose(t *testing.T) {
	tc := tripleDisposeTestCases[0]
	tr := coregeneric.NewTripleWithMessage("a", "b", "c", true, "msg")
	tr.Dispose()
	tc.ShouldBeEqual(t, 0, tr.Left, tr.Middle, tr.Right, fmt.Sprintf("%v", tr.IsValid), tr.Message)
}

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
