package errcore

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/namevalue"
)

func TestVarTwo_WithType(t *testing.T) {
	s := VarTwo(true, "a", 1, "b", 2)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestVarTwo_WithoutType(t *testing.T) {
	s := VarTwo(false, "a", 1, "b", 2)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestVarTwoNoType(t *testing.T) {
	s := VarTwoNoType("a", 1, "b", 2)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestVarThree_WithType(t *testing.T) {
	s := VarThree(true, "a", 1, "b", 2, "c", 3)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestVarThree_WithoutType(t *testing.T) {
	s := VarThree(false, "a", 1, "b", 2, "c", 3)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestVarThreeNoType(t *testing.T) {
	s := VarThreeNoType("a", 1, "b", 2, "c", 3)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestVarMap_Empty(t *testing.T) {
	if VarMap(nil) != "" {
		t.Fatal("expected empty")
	}
}

func TestVarMap_WithItems(t *testing.T) {
	s := VarMap(map[string]any{"a": 1})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestVarMapStrings_Empty(t *testing.T) {
	if len(VarMapStrings(nil)) != 0 {
		t.Fatal("expected empty")
	}
}

func TestVarMapStrings_WithItems(t *testing.T) {
	s := VarMapStrings(map[string]any{"a": 1})
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func TestVarNameValues_Empty(t *testing.T) {
	if VarNameValues() != "" {
		t.Fatal("expected empty")
	}
}

func TestVarNameValues_WithItems(t *testing.T) {
	s := VarNameValues(namevalue.StringAny{Name: "a", Value: 1})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestVarNameValuesJoiner_Empty(t *testing.T) {
	if VarNameValuesJoiner(",") != "" {
		t.Fatal("expected empty")
	}
}

func TestVarNameValuesJoiner_WithItems(t *testing.T) {
	s := VarNameValuesJoiner(",", namevalue.StringAny{Name: "a", Value: 1})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestVarNameValuesStrings_Empty(t *testing.T) {
	if len(VarNameValuesStrings()) != 0 {
		t.Fatal("expected empty")
	}
}

func TestVarNameValuesStrings_WithItems(t *testing.T) {
	s := VarNameValuesStrings(namevalue.StringAny{Name: "a", Value: 1})
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func TestMessageVarTwo(t *testing.T) {
	s := MessageVarTwo("msg", "a", 1, "b", 2)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestMessageVarThree(t *testing.T) {
	s := MessageVarThree("msg", "a", 1, "b", 2, "c", 3)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestMessageVarMap_Empty(t *testing.T) {
	s := MessageVarMap("msg", nil)
	if s != "msg" {
		t.Fatal("expected just msg")
	}
}

func TestMessageVarMap_WithItems(t *testing.T) {
	s := MessageVarMap("msg", map[string]any{"a": 1})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestMessageNameValues_Empty(t *testing.T) {
	s := MessageNameValues("msg")
	if s != "msg" {
		t.Fatal("expected just msg")
	}
}

func TestMessageNameValues_WithItems(t *testing.T) {
	s := MessageNameValues("msg", namevalue.StringAny{Name: "a", Value: 1})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestMessageWithRef(t *testing.T) {
	s := MessageWithRef("msg", "ref")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestMessageWithRefToError(t *testing.T) {
	err := MessageWithRefToError("msg", "ref")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRef_Nil(t *testing.T) {
	if Ref(nil) != "" {
		t.Fatal("expected empty")
	}
}

func TestRef_WithRef(t *testing.T) {
	s := Ref("ref")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRefToError_Nil(t *testing.T) {
	if RefToError(nil) != nil {
		t.Fatal("expected nil")
	}
}

func TestRefToError_WithRef(t *testing.T) {
	err := RefToError("ref")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestToError_Empty(t *testing.T) {
	if ToError("") != nil {
		t.Fatal("expected nil")
	}
}

func TestToError_WithMsg(t *testing.T) {
	err := ToError("msg")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestToString_Nil(t *testing.T) {
	if ToString(nil) != "" {
		t.Fatal("expected empty")
	}
}

func TestToString_WithErr(t *testing.T) {
	s := ToString(errors.New("e"))
	if s != "e" {
		t.Fatal("expected e")
	}
}

func TestToStringPtr_Nil(t *testing.T) {
	p := ToStringPtr(nil)
	if p == nil || *p != "" {
		t.Fatal("expected empty ptr")
	}
}

func TestToStringPtr_WithErr(t *testing.T) {
	p := ToStringPtr(errors.New("e"))
	if p == nil || *p != "e" {
		t.Fatal("expected e")
	}
}

func TestToValueString(t *testing.T) {
	s := ToValueString("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestToExitError_Nil(t *testing.T) {
	if ToExitError(nil) != nil {
		t.Fatal("expected nil")
	}
}

func TestToExitError_NonExitError(t *testing.T) {
	if ToExitError(errors.New("e")) != nil {
		t.Fatal("expected nil")
	}
}

func TestSourceDestination(t *testing.T) {
	s := SourceDestination(true, "src", "dst")
	if s == "" {
		t.Fatal("expected non-empty")
	}
	s2 := SourceDestination(false, "src", "dst")
	if s2 == "" {
		t.Fatal("expected non-empty")
	}
}

func TestSourceDestinationNoType(t *testing.T) {
	s := SourceDestinationNoType("src", "dst")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestSourceDestinationErr(t *testing.T) {
	err := SourceDestinationErr(true, "src", "dst")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestCombine(t *testing.T) {
	s := Combine("gen", "other", "ref")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestCombineWithMsgTypeNoStack(t *testing.T) {
	s := CombineWithMsgTypeNoStack(InvalidType, "", nil)
	if s == "" {
		t.Fatal("expected non-empty")
	}
	s2 := CombineWithMsgTypeNoStack(InvalidType, "msg", "ref")
	if s2 == "" {
		t.Fatal("expected non-empty")
	}
}

func TestCombineWithMsgTypeStackTrace(t *testing.T) {
	s := CombineWithMsgTypeStackTrace(InvalidType, "msg", nil)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestStackTracesCompiled(t *testing.T) {
	s := StackTracesCompiled([]string{"a", "b"})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestGherkinsString(t *testing.T) {
	s := GherkinsString(0, "f", "g", "w", "t")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestGherkinsStringWithExpectation(t *testing.T) {
	s := GherkinsStringWithExpectation(0, "f", "g", "w", "t", "a", "e")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRangeNotMeet_WithRange(t *testing.T) {
	s := RangeNotMeet("msg", 0, 10, []int{1, 2})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRangeNotMeet_WithoutRange(t *testing.T) {
	s := RangeNotMeet("msg", 0, 10, nil)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestPanicRangeNotMeet_WithRange(t *testing.T) {
	s := PanicRangeNotMeet("msg", 0, 10, []int{1})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestPanicRangeNotMeet_WithoutRange(t *testing.T) {
	s := PanicRangeNotMeet("msg", 0, 10, nil)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestEnumRangeNotMeet_WithRange(t *testing.T) {
	s := EnumRangeNotMeet(0, 10, "range")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestEnumRangeNotMeet_WithoutRange(t *testing.T) {
	s := EnumRangeNotMeet(0, 10, nil)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestMsgHeader(t *testing.T) {
	s := MsgHeader("test")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestMsgHeaderIf_True(t *testing.T) {
	s := MsgHeaderIf(true, "test")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestMsgHeaderIf_False(t *testing.T) {
	s := MsgHeaderIf(false, "test")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestMsgHeaderPlusEnding(t *testing.T) {
	s := MsgHeaderPlusEnding("header", "msg")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestSliceError_Empty(t *testing.T) {
	if SliceError(",", nil) != nil {
		t.Fatal("expected nil")
	}
}

func TestSliceError_WithItems(t *testing.T) {
	err := SliceError(",", []string{"a", "b"})
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestSliceErrorDefault(t *testing.T) {
	err := SliceErrorDefault([]string{"a"})
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestSliceToError_Empty(t *testing.T) {
	if SliceToError(nil) != nil {
		t.Fatal("expected nil")
	}
}

func TestSliceToError_WithItems(t *testing.T) {
	err := SliceToError([]string{"a"})
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestSliceToErrorPtr_Empty(t *testing.T) {
	if SliceToErrorPtr(nil) != nil {
		t.Fatal("expected nil")
	}
}

func TestSliceToErrorPtr_WithItems(t *testing.T) {
	err := SliceToErrorPtr([]string{"a"})
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestSliceErrorsToStrings_Nil(t *testing.T) {
	s := SliceErrorsToStrings(nil...)
	if len(s) != 0 {
		t.Fatal("expected empty")
	}
}

func TestSliceErrorsToStrings_WithItems(t *testing.T) {
	s := SliceErrorsToStrings(errors.New("a"), nil, errors.New("b"))
	if len(s) != 2 {
		t.Fatal("expected 2")
	}
}

func TestManyErrorToSingle(t *testing.T) {
	err := ManyErrorToSingle([]error{errors.New("a"), nil})
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestManyErrorToSingleDirect(t *testing.T) {
	err := ManyErrorToSingleDirect(errors.New("a"))
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestMergeErrors(t *testing.T) {
	err := MergeErrors(errors.New("a"), errors.New("b"))
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestMergeErrorsToString_Nil(t *testing.T) {
	if MergeErrorsToString(",", nil...) != "" {
		t.Fatal("expected empty")
	}
}

func TestMergeErrorsToString_WithItems(t *testing.T) {
	s := MergeErrorsToString(",", errors.New("a"))
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestMergeErrorsToStringDefault_Nil(t *testing.T) {
	if MergeErrorsToStringDefault(nil...) != "" {
		t.Fatal("expected empty")
	}
}

func TestMergeErrorsToStringDefault_WithItems(t *testing.T) {
	s := MergeErrorsToStringDefault(errors.New("a"))
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestStringLinesToQuoteLines_Empty(t *testing.T) {
	if len(StringLinesToQuoteLines(nil)) != 0 {
		t.Fatal("expected empty")
	}
}

func TestStringLinesToQuoteLines_WithItems(t *testing.T) {
	s := StringLinesToQuoteLines([]string{"a"})
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func TestStringLinesToQuoteLinesToSingle(t *testing.T) {
	s := StringLinesToQuoteLinesToSingle([]string{"a", "b"})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestLinesToDoubleQuoteLinesWithTabs_Empty(t *testing.T) {
	if len(LinesToDoubleQuoteLinesWithTabs(2, nil)) != 0 {
		t.Fatal("expected empty")
	}
}

func TestLinesToDoubleQuoteLinesWithTabs_WithItems(t *testing.T) {
	s := LinesToDoubleQuoteLinesWithTabs(4, []string{"a"})
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func TestFmtDebug(t *testing.T) {
	FmtDebug("test %s", "v")
}

func TestFmtDebugIf_False(t *testing.T) {
	FmtDebugIf(false, "skip")
}

func TestFmtDebugIf_True(t *testing.T) {
	FmtDebugIf(true, "test %s", "v")
}

func TestValidPrint(t *testing.T) {
	ValidPrint(false, "skip")
	ValidPrint(true, "show")
}

func TestFailedPrint(t *testing.T) {
	FailedPrint(false, "skip")
	FailedPrint(true, "show")
}

func TestGetReferenceMessage_Nil(t *testing.T) {
	if getReferenceMessage(nil) != "" {
		t.Fatal("expected empty")
	}
}

func TestGetReferenceMessage_EmptyString(t *testing.T) {
	if getReferenceMessage("") != "" {
		t.Fatal("expected empty")
	}
}

func TestGetReferenceMessage_WithRef(t *testing.T) {
	s := getReferenceMessage("ref")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestTypesNamesString(t *testing.T) {
	s := typesNamesString("a", 1)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestGetActualAndExpectProcessedMessage(t *testing.T) {
	s := GetActualAndExpectProcessedMessage(0, "a", "e", "ap", "ep")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestGetSearchLineNumberExpectationMessage(t *testing.T) {
	s := GetSearchLineNumberExpectationMessage(0, 1, 2, "content", "search", "info")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestGetSearchTermExpectationMessage_WithInfo(t *testing.T) {
	s := GetSearchTermExpectationMessage(0, "h", "e", 1, "a", "e", "info")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestGetSearchTermExpectationMessage_NilInfo(t *testing.T) {
	s := GetSearchTermExpectationMessage(0, "h", "e", 1, "a", "e", nil)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestGetSearchTermExpectationSimpleMessage(t *testing.T) {
	s := GetSearchTermExpectationSimpleMessage(0, "e", 1, "c", "s")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestExpected_But(t *testing.T) {
	err := Expected.But("t", "e", "a")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestExpected_ButFoundAsMsg(t *testing.T) {
	s := Expected.ButFoundAsMsg("t", "e", "a")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestExpected_ButFoundWithTypeAsMsg(t *testing.T) {
	s := Expected.ButFoundWithTypeAsMsg("t", "e", "a")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestExpected_ButUsingType(t *testing.T) {
	err := Expected.ButUsingType("t", "e", "a")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestExpected_ReflectButFound(t *testing.T) {
	err := Expected.ReflectButFound(reflect.Int, reflect.String)
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestExpected_PrimitiveButFound(t *testing.T) {
	err := Expected.PrimitiveButFound(reflect.Map)
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestExpected_ValueHasNoElements(t *testing.T) {
	err := Expected.ValueHasNoElements(reflect.Slice)
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestShouldBe_StrEqMsg(t *testing.T) {
	s := ShouldBe.StrEqMsg("a", "b")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestShouldBe_StrEqErr(t *testing.T) {
	err := ShouldBe.StrEqErr("a", "b")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestShouldBe_AnyEqMsg(t *testing.T) {
	s := ShouldBe.AnyEqMsg(1, 2)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestShouldBe_AnyEqErr(t *testing.T) {
	err := ShouldBe.AnyEqErr(1, 2)
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestShouldBe_JsonEqMsg(t *testing.T) {
	s := ShouldBe.JsonEqMsg("a", "b")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestShouldBe_JsonEqErr(t *testing.T) {
	err := ShouldBe.JsonEqErr("a", "b")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}
