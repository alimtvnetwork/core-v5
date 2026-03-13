package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/errcore"
)

// TestSliceToError verifies SliceToError.
func TestSliceToError(t *testing.T) {
	// Empty returns nil
	if errcore.SliceToError(nil) != nil {
		t.Error("nil should return nil")
	}
	if errcore.SliceToError([]string{}) != nil {
		t.Error("empty should return nil")
	}

	// Non-empty returns error
	err := errcore.SliceToError([]string{"err1", "err2"})
	if err == nil {
		t.Error("should return error")
	}
}

// TestSliceToErrorPtr verifies SliceToErrorPtr.
func TestSliceToErrorPtr(t *testing.T) {
	if errcore.SliceToErrorPtr(nil) != nil {
		t.Error("nil should return nil")
	}
	err := errcore.SliceToErrorPtr([]string{"e1"})
	if err == nil {
		t.Error("should return error")
	}
}

// TestToError verifies ToError.
func TestToError(t *testing.T) {
	if errcore.ToError("") != nil {
		t.Error("empty should return nil")
	}
	err := errcore.ToError("fail")
	if err == nil || err.Error() != "fail" {
		t.Error("should return error with message")
	}
}

// TestToString verifies ToString.
func TestToString(t *testing.T) {
	if errcore.ToString(nil) != "" {
		t.Error("nil should return empty")
	}
	if errcore.ToString(errors.New("test")) != "test" {
		t.Error("should return error string")
	}
}

// TestToStringPtr verifies ToStringPtr.
func TestToStringPtr(t *testing.T) {
	r := errcore.ToStringPtr(nil)
	if r == nil {
		t.Error("should return pointer")
	}
	if *r != "" {
		t.Error("nil error should give empty string")
	}
}

// TestMergeErrors verifies MergeErrors.
func TestMergeErrors(t *testing.T) {
	if errcore.MergeErrors() != nil {
		t.Error("no errors should return nil")
	}
	if errcore.MergeErrors(nil, nil) != nil {
		t.Error("all nil should return nil")
	}
	err := errcore.MergeErrors(errors.New("a"), errors.New("b"))
	if err == nil {
		t.Error("should return combined error")
	}
}

// TestCombine verifies Combine.
func TestCombine(t *testing.T) {
	result := errcore.Combine("generic", "other", "ref")
	if result == "" {
		t.Error("should return non-empty")
	}
}

// TestSliceError verifies SliceError.
func TestSliceError(t *testing.T) {
	if errcore.SliceError(",", nil) != nil {
		t.Error("nil should return nil")
	}
	if errcore.SliceError(",", []string{}) != nil {
		t.Error("empty should return nil")
	}
	err := errcore.SliceError(",", []string{"a", "b"})
	if err == nil {
		t.Error("should return error")
	}
}

// TestSliceErrorDefault verifies SliceErrorDefault.
func TestSliceErrorDefault(t *testing.T) {
	if errcore.SliceErrorDefault(nil) != nil {
		t.Error("nil should return nil")
	}
}

// TestManyErrorToSingle verifies ManyErrorToSingle.
func TestManyErrorToSingle(t *testing.T) {
	r := errcore.ManyErrorToSingle(nil)
	if r != nil {
		t.Error("nil should return nil")
	}
	r = errcore.ManyErrorToSingle([]error{errors.New("x")})
	if r == nil {
		t.Error("single error should return it")
	}
}

// TestManyErrorToSingleDirect verifies ManyErrorToSingleDirect.
func TestManyErrorToSingleDirect(t *testing.T) {
	r := errcore.ManyErrorToSingleDirect()
	if r != nil {
		t.Error("empty should return nil")
	}
	r = errcore.ManyErrorToSingleDirect(errors.New("a"))
	if r == nil {
		t.Error("should return error")
	}
}

// TestConcatMessageWithErr verifies ConcatMessageWithErr.
func TestConcatMessageWithErr(t *testing.T) {
	r := errcore.ConcatMessageWithErr("prefix", nil)
	if r != nil {
		t.Error("nil error should return nil")
	}
	r = errcore.ConcatMessageWithErr("prefix", errors.New("err"))
	if r == nil {
		t.Error("should return concatenated error")
	}
}

// TestExpecting verifies Expecting error message.
func TestExpecting(t *testing.T) {
	r := errcore.Expecting("header", "expected", "actual")
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestExpectingSimple verifies ExpectingSimple.
func TestExpectingSimple(t *testing.T) {
	r := errcore.ExpectingSimple("header", "expected", "actual")
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestExpectingSimpleNoType verifies ExpectingSimpleNoType.
func TestExpectingSimpleNoType(t *testing.T) {
	r := errcore.ExpectingSimpleNoType("header", "expected", "actual")
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestExpectingError verifies ExpectingError.
func TestExpectingError(t *testing.T) {
	err := errcore.ExpectingErrorSimpleNoType("header", "expected", "actual")
	if err == nil {
		t.Error("should return error")
	}
}

// TestExpectingErrorSimpleNoType verifies ExpectingErrorSimpleNoType.
func TestExpectingErrorSimpleNoType(t *testing.T) {
	err := errcore.ExpectingErrorSimpleNoType("header", "expected", "actual")
	if err == nil {
		t.Error("should return error")
	}
}

// TestExpectingNotEqualSimpleNoType verifies ExpectingNotEqualSimpleNoType.
func TestExpectingNotEqualSimpleNoType(t *testing.T) {
	r := errcore.ExpectingNotEqualSimpleNoType("header", "a", "b")
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestVarTwo verifies VarTwo.
func TestVarTwo(t *testing.T) {
	r := errcore.VarTwo(false, "a", 1, "b", 2)
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestVarThree verifies VarThree.
func TestVarThree(t *testing.T) {
	r := errcore.VarThree(false, "a", 1, "b", 2, "c", 3)
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestMessageVarTwo verifies MessageVarTwo.
func TestMessageVarTwo(t *testing.T) {
	r := errcore.MessageVarTwo("msg", "a", 1, "b", 2)
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestMessageVarThree verifies MessageVarThree.
func TestMessageVarThree(t *testing.T) {
	r := errcore.MessageVarThree("msg", "a", 1, "b", 2, "c", 3)
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestMessageVarMap verifies MessageVarMap.
func TestMessageVarMap(t *testing.T) {
	r := errcore.MessageVarMap("msg", map[string]any{"k": "v"})
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestVarMap verifies VarMap.
func TestVarMap(t *testing.T) {
	r := errcore.VarMap(map[string]any{"k": "v"})
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestShouldBe verifies ShouldBe.
func TestShouldBe(t *testing.T) {
	r := errcore.ShouldBe.StrEqMsg("actual", "expected")
	if r == "" {
		t.Error("should return non-empty")
	}
	err := errcore.ShouldBe.StrEqErr("actual", "expected")
	if err == nil {
		t.Error("should return error")
	}
}

// TestRawErrCollection verifies RawErrCollection.
func TestRawErrCollection(t *testing.T) {
	c := errcore.RawErrCollection{}
	if c.HasError() {
		t.Error("empty should not have error")
	}
	c.Add(errors.New("err1"))
	if !c.HasError() {
		t.Error("should have error")
	}
	if c.CompiledError() == nil {
		t.Error("should return compiled error")
	}
}

// TestSliceErrorsToStrings verifies SliceErrorsToStrings.
func TestSliceErrorsToStrings(t *testing.T) {
	r := errcore.SliceErrorsToStrings(nil)
	if len(r) != 0 {
		t.Error("nil should return empty")
	}
	r = errcore.SliceErrorsToStrings(errors.New("a"), errors.New("b"))
	if len(r) != 2 {
		t.Errorf("expected 2, got %d", len(r))
	}
}

// TestErrorToSplitLines verifies ErrorToSplitLines.
func TestErrorToSplitLines(t *testing.T) {
	r := errcore.ErrorToSplitLines(nil)
	if r != nil {
		t.Error("nil should return nil")
	}
	r = errcore.ErrorToSplitLines(errors.New("a\nb"))
	if len(r) != 2 {
		t.Errorf("expected 2 lines, got %d", len(r))
	}
}

// TestErrorToSplitNonEmptyLines verifies ErrorToSplitNonEmptyLines.
func TestErrorToSplitNonEmptyLines(t *testing.T) {
	r := errcore.ErrorToSplitNonEmptyLines(nil)
	if r != nil {
		t.Error("nil should return nil")
	}
}

// TestRef verifies Ref.
func TestRef(t *testing.T) {
	r := errcore.Ref("ref")
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestMessageWithRef verifies MessageWithRef.
func TestMessageWithRef(t *testing.T) {
	r := errcore.MessageWithRef("msg", "ref")
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestVarTwoNoType verifies VarTwoNoType.
func TestVarTwoNoType(t *testing.T) {
	r := errcore.VarTwoNoType("a", 1, "b", 2)
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestVarThreeNoType verifies VarThreeNoType.
func TestVarThreeNoType(t *testing.T) {
	r := errcore.VarThreeNoType("a", 1, "b", 2, "c", 3)
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestGetSearchTermExpectationMessage verifies search term message.
func TestGetSearchTermExpectationMessage(t *testing.T) {
	r := errcore.GetSearchTermExpectationMessage(1, "header", "expectation", 0, "actual", "expected", nil)
	if r == "" {
		t.Error("should return non-empty")
	}
}

// TestGetSearchTermExpectationSimpleMessage verifies simple search term message.
func TestGetSearchTermExpectationSimpleMessage(t *testing.T) {
	r := errcore.GetSearchTermExpectationSimpleMessage("search", "content", true)
	if r == "" {
		t.Error("should return non-empty")
	}
}
