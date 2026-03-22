package corestrtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// utils — WrapDouble, WrapSingle, WrapTilda, WrapDoubleIfMissing, WrapSingleIfMissing
// =============================================================================

func Test_Cov66_Utils_WrapDouble(t *testing.T) {
	actual := args.Map{"r": corestr.StringUtils.WrapDouble("hello")}
	expected := args.Map{"r": "\"hello\""}
	expected.ShouldBeEqual(t, 0, "WrapDouble", actual)
}

func Test_Cov66_Utils_WrapSingle(t *testing.T) {
	actual := args.Map{"r": corestr.StringUtils.WrapSingle("hello")}
	expected := args.Map{"r": "'hello'"}
	expected.ShouldBeEqual(t, 0, "WrapSingle", actual)
}

func Test_Cov66_Utils_WrapTilda(t *testing.T) {
	actual := args.Map{"r": corestr.StringUtils.WrapTilda("hello")}
	expected := args.Map{"r": "`hello`"}
	expected.ShouldBeEqual(t, 0, "WrapTilda", actual)
}

func Test_Cov66_Utils_WrapDoubleIfMissing_NoWrap(t *testing.T) {
	actual := args.Map{"r": corestr.StringUtils.WrapDoubleIfMissing("\"already\"")}
	expected := args.Map{"r": "\"already\""}
	expected.ShouldBeEqual(t, 0, "WrapDoubleIfMissing already wrapped", actual)
}

func Test_Cov66_Utils_WrapDoubleIfMissing_Wrap(t *testing.T) {
	actual := args.Map{"r": corestr.StringUtils.WrapDoubleIfMissing("need")}
	expected := args.Map{"r": "\"need\""}
	expected.ShouldBeEqual(t, 0, "WrapDoubleIfMissing wraps", actual)
}

func Test_Cov66_Utils_WrapDoubleIfMissing_Empty(t *testing.T) {
	actual := args.Map{"r": corestr.StringUtils.WrapDoubleIfMissing("")}
	expected := args.Map{"r": "\"\""}
	expected.ShouldBeEqual(t, 0, "WrapDoubleIfMissing empty", actual)
}

func Test_Cov66_Utils_WrapSingleIfMissing_NoWrap(t *testing.T) {
	actual := args.Map{"r": corestr.StringUtils.WrapSingleIfMissing("'already'")}
	expected := args.Map{"r": "'already'"}
	expected.ShouldBeEqual(t, 0, "WrapSingleIfMissing already wrapped", actual)
}

func Test_Cov66_Utils_WrapSingleIfMissing_Wrap(t *testing.T) {
	actual := args.Map{"r": corestr.StringUtils.WrapSingleIfMissing("need")}
	expected := args.Map{"r": "'need'"}
	expected.ShouldBeEqual(t, 0, "WrapSingleIfMissing wraps", actual)
}

func Test_Cov66_Utils_WrapSingleIfMissing_Empty(t *testing.T) {
	actual := args.Map{"r": corestr.StringUtils.WrapSingleIfMissing("")}
	expected := args.Map{"r": "''"}
	expected.ShouldBeEqual(t, 0, "WrapSingleIfMissing empty", actual)
}

// =============================================================================
// LeftMiddleRight — 15 uncovered methods
// =============================================================================

func Test_Cov66_LMR_LeftBytes(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")
	actual := args.Map{"len": len(lmr.LeftBytes())}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "LMR LeftBytes", actual)
}

func Test_Cov66_LMR_RightBytes(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")
	actual := args.Map{"len": len(lmr.RightBytes())}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "LMR RightBytes", actual)
}

func Test_Cov66_LMR_MiddleBytes(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")
	actual := args.Map{"len": len(lmr.MiddleBytes())}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "LMR MiddleBytes", actual)
}

func Test_Cov66_LMR_LeftTrim(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("  abc  ", "mid", "xyz")
	actual := args.Map{"r": lmr.LeftTrim()}
	expected := args.Map{"r": "abc"}
	expected.ShouldBeEqual(t, 0, "LMR LeftTrim", actual)
}

func Test_Cov66_LMR_RightTrim(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("abc", "mid", "  xyz  ")
	actual := args.Map{"r": lmr.RightTrim()}
	expected := args.Map{"r": "xyz"}
	expected.ShouldBeEqual(t, 0, "LMR RightTrim", actual)
}

func Test_Cov66_LMR_MiddleTrim(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("abc", "  mid  ", "xyz")
	actual := args.Map{"r": lmr.MiddleTrim()}
	expected := args.Map{"r": "mid"}
	expected.ShouldBeEqual(t, 0, "LMR MiddleTrim", actual)
}

func Test_Cov66_LMR_IsLeftEmpty(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("", "mid", "xyz")
	actual := args.Map{"r": lmr.IsLeftEmpty()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LMR IsLeftEmpty", actual)
}

func Test_Cov66_LMR_IsRightEmpty(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("abc", "mid", "")
	actual := args.Map{"r": lmr.IsRightEmpty()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LMR IsRightEmpty", actual)
}

func Test_Cov66_LMR_IsMiddleEmpty(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("abc", "", "xyz")
	actual := args.Map{"r": lmr.IsMiddleEmpty()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LMR IsMiddleEmpty", actual)
}

func Test_Cov66_LMR_IsMiddleWhitespace(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("abc", "   ", "xyz")
	actual := args.Map{"r": lmr.IsMiddleWhitespace()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LMR IsMiddleWhitespace", actual)
}

func Test_Cov66_LMR_IsLeftWhitespace(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("   ", "mid", "xyz")
	actual := args.Map{"r": lmr.IsLeftWhitespace()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LMR IsLeftWhitespace", actual)
}

func Test_Cov66_LMR_IsRightWhitespace(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("abc", "mid", "   ")
	actual := args.Map{"r": lmr.IsRightWhitespace()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LMR IsRightWhitespace", actual)
}

func Test_Cov66_LMR_HasValidNonEmptyLeft(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")
	actual := args.Map{"r": lmr.HasValidNonEmptyLeft()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LMR HasValidNonEmptyLeft", actual)
}

func Test_Cov66_LMR_HasValidNonEmptyRight(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")
	actual := args.Map{"r": lmr.HasValidNonEmptyRight()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LMR HasValidNonEmptyRight", actual)
}

func Test_Cov66_LMR_HasValidNonEmptyMiddle(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")
	actual := args.Map{"r": lmr.HasValidNonEmptyMiddle()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LMR HasValidNonEmptyMiddle", actual)
}

func Test_Cov66_LMR_HasValidNonWhitespaceLeft(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")
	actual := args.Map{"r": lmr.HasValidNonWhitespaceLeft()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LMR HasValidNonWhitespaceLeft", actual)
}

func Test_Cov66_LMR_HasValidNonWhitespaceRight(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")
	actual := args.Map{"r": lmr.HasValidNonWhitespaceRight()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LMR HasValidNonWhitespaceRight", actual)
}

func Test_Cov66_LMR_HasValidNonWhitespaceMiddle(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")
	actual := args.Map{"r": lmr.HasValidNonWhitespaceMiddle()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LMR HasValidNonWhitespaceMiddle", actual)
}

func Test_Cov66_LMR_HasSafeNonEmpty(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("abc", "mid", "xyz")
	actual := args.Map{"r": lmr.HasSafeNonEmpty()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LMR HasSafeNonEmpty", actual)
}

func Test_Cov66_LMR_HasSafeNonEmpty_False(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("", "mid", "xyz")
	actual := args.Map{"r": lmr.HasSafeNonEmpty()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "LMR HasSafeNonEmpty false", actual)
}

func Test_Cov66_LMR_IsAll(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	actual := args.Map{"r": lmr.IsAll("a", "b", "c")}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LMR IsAll", actual)
}

func Test_Cov66_LMR_Is(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	actual := args.Map{"r": lmr.Is("a", "c")}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LMR Is", actual)
}

func Test_Cov66_LMR_Clone(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	clone := lmr.Clone()
	actual := args.Map{"left": clone.Left, "mid": clone.Middle, "right": clone.Right}
	expected := args.Map{"left": "a", "mid": "b", "right": "c"}
	expected.ShouldBeEqual(t, 0, "LMR Clone", actual)
}

func Test_Cov66_LMR_ToLeftRight(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	lr := lmr.ToLeftRight()
	actual := args.Map{"left": lr.Left, "right": lr.Right}
	expected := args.Map{"left": "a", "right": "c"}
	expected.ShouldBeEqual(t, 0, "LMR ToLeftRight", actual)
}

func Test_Cov66_LMR_Clear(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	lmr.Clear()
	actual := args.Map{"left": lmr.Left, "mid": lmr.Middle, "right": lmr.Right}
	expected := args.Map{"left": "", "mid": "", "right": ""}
	expected.ShouldBeEqual(t, 0, "LMR Clear", actual)
}

func Test_Cov66_LMR_Dispose(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	lmr.Dispose()
	actual := args.Map{"left": lmr.Left}
	expected := args.Map{"left": ""}
	expected.ShouldBeEqual(t, 0, "LMR Dispose", actual)
}

func Test_Cov66_LMR_InvalidNoMessage(t *testing.T) {
	lmr := corestr.InvalidLeftMiddleRightNoMessage()
	actual := args.Map{"valid": lmr.IsValid}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "LMR InvalidNoMessage", actual)
}

func Test_Cov66_LMR_InvalidWithMessage(t *testing.T) {
	lmr := corestr.InvalidLeftMiddleRight("bad")
	actual := args.Map{"valid": lmr.IsValid, "msg": lmr.Message}
	expected := args.Map{"valid": false, "msg": "bad"}
	expected.ShouldBeEqual(t, 0, "LMR InvalidWithMessage", actual)
}

// =============================================================================
// NonChainedLinkedCollectionNodes — 6 uncovered
// =============================================================================

func Test_Cov66_NCLCN_IsEmpty_Empty(t *testing.T) {
	n := corestr.NewNonChainedLinkedCollectionNodes(0)
	actual := args.Map{"empty": n.IsEmpty(), "has": n.HasItems(), "len": n.Length()}
	expected := args.Map{"empty": true, "has": false, "len": 0}
	expected.ShouldBeEqual(t, 0, "NCLCN empty", actual)
}

func Test_Cov66_NCLCN_IsChainingApplied_False(t *testing.T) {
	n := corestr.NewNonChainedLinkedCollectionNodes(2)
	actual := args.Map{"chained": n.IsChainingApplied()}
	expected := args.Map{"chained": false}
	expected.ShouldBeEqual(t, 0, "NCLCN not chained", actual)
}

func Test_Cov66_NCLCN_FirstLast(t *testing.T) {
	lc := corestr.New.LinkedCollection.Create()
	lc.Add(corestr.New.Collection.Strings([]string{"a"}))
	lc.Add(corestr.New.Collection.Strings([]string{"b"}))
	nodes := lc.GetAllLinkedNodes()
	actual := args.Map{
		"firstNonNil": nodes.First() != nil,
		"lastNonNil":  nodes.Last() != nil,
		"len":         nodes.Length(),
	}
	expected := args.Map{
		"firstNonNil": true,
		"lastNonNil":  true,
		"len":         2,
	}
	expected.ShouldBeEqual(t, 0, "NCLCN First/Last", actual)
}

// =============================================================================
// NonChainedLinkedListNodes — 6 uncovered
// =============================================================================

func Test_Cov66_NCLLN_IsEmpty_Empty(t *testing.T) {
	n := corestr.NewNonChainedLinkedListNodes(0)
	actual := args.Map{"empty": n.IsEmpty(), "has": n.HasItems(), "len": n.Length()}
	expected := args.Map{"empty": true, "has": false, "len": 0}
	expected.ShouldBeEqual(t, 0, "NCLLN empty", actual)
}

func Test_Cov66_NCLLN_IsChainingApplied_False(t *testing.T) {
	n := corestr.NewNonChainedLinkedListNodes(2)
	actual := args.Map{"chained": n.IsChainingApplied()}
	expected := args.Map{"chained": false}
	expected.ShouldBeEqual(t, 0, "NCLLN not chained", actual)
}

func Test_Cov66_NCLLN_FirstLast(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	ll.Add("x")
	ll.Add("y")
	nodes := ll.GetAllLinkedNodes()
	actual := args.Map{
		"firstNonNil": nodes.First() != nil,
		"lastNonNil":  nodes.Last() != nil,
		"len":         nodes.Length(),
	}
	expected := args.Map{
		"firstNonNil": true,
		"lastNonNil":  true,
		"len":         2,
	}
	expected.ShouldBeEqual(t, 0, "NCLLN First/Last", actual)
}

// =============================================================================
// LinkedCollections — Tail
// =============================================================================

func Test_Cov66_LC_Tail(t *testing.T) {
	lc := corestr.New.LinkedCollection.Create()
	lc.Add(corestr.New.Collection.Strings([]string{"a"}))
	lc.Add(corestr.New.Collection.Strings([]string{"b"}))
	tail := lc.Tail()
	actual := args.Map{"nonNil": tail != nil}
	expected := args.Map{"nonNil": true}
	expected.ShouldBeEqual(t, 0, "LC Tail", actual)
}

func Test_Cov66_LC_Tail_Empty(t *testing.T) {
	lc := corestr.New.LinkedCollection.Create()
	tail := lc.Tail()
	actual := args.Map{"isNil": tail == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LC Tail empty", actual)
}

// =============================================================================
// CollectionsOfCollection — JSON methods (9 uncovered)
// =============================================================================

func Test_Cov66_COC_Json(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Create()
	coc.Add(corestr.New.Collection.Strings([]string{"a"}))
	r := coc.Json()
	actual := args.Map{"noErr": r.Error == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "COC Json", actual)
}

func Test_Cov66_COC_JsonModel(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Create()
	coc.Add(corestr.New.Collection.Strings([]string{"a"}))
	m := coc.JsonModel()
	actual := args.Map{"nonNil": m.Items != nil}
	expected := args.Map{"nonNil": true}
	expected.ShouldBeEqual(t, 0, "COC JsonModel", actual)
}

func Test_Cov66_COC_JsonModelAny(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Create()
	actual := args.Map{"nonNil": coc.JsonModelAny() != nil}
	expected := args.Map{"nonNil": true}
	expected.ShouldBeEqual(t, 0, "COC JsonModelAny", actual)
}

func Test_Cov66_COC_MarshalJSON(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Create()
	coc.Add(corestr.New.Collection.Strings([]string{"a"}))
	b, err := coc.MarshalJSON()
	actual := args.Map{"noErr": err == nil, "nonEmpty": len(b) > 0}
	expected := args.Map{"noErr": true, "nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "COC MarshalJSON", actual)
}

func Test_Cov66_COC_UnmarshalJSON(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Create()
	coc.Add(corestr.New.Collection.Strings([]string{"a"}))
	b, _ := coc.MarshalJSON()
	coc2 := corestr.New.CollectionsOfCollection.Create()
	err := coc2.UnmarshalJSON(b)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "COC UnmarshalJSON", actual)
}

func Test_Cov66_COC_ParseInjectUsingJson(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Create()
	coc.Add(corestr.New.Collection.Strings([]string{"a"}))
	jr := coc.JsonPtr()
	coc2 := corestr.New.CollectionsOfCollection.Create()
	r, err := coc2.ParseInjectUsingJson(jr)
	actual := args.Map{"noErr": err == nil, "nonNil": r != nil}
	expected := args.Map{"noErr": true, "nonNil": true}
	expected.ShouldBeEqual(t, 0, "COC ParseInjectUsingJson", actual)
}

func Test_Cov66_COC_ParseInjectUsingJson_Error(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Create()
	jr := &corejson.Result{Error: errors.New("fail")}
	_, err := coc.ParseInjectUsingJson(jr)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "COC ParseInjectUsingJson error", actual)
}

func Test_Cov66_COC_ParseInjectUsingJsonMust(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Create()
	coc.Add(corestr.New.Collection.Strings([]string{"a"}))
	jr := coc.JsonPtr()
	coc2 := corestr.New.CollectionsOfCollection.Create()
	r := coc2.ParseInjectUsingJsonMust(jr)
	actual := args.Map{"nonNil": r != nil}
	expected := args.Map{"nonNil": true}
	expected.ShouldBeEqual(t, 0, "COC ParseInjectUsingJsonMust", actual)
}

func Test_Cov66_COC_JsonParseSelfInject(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Create()
	coc.Add(corestr.New.Collection.Strings([]string{"a"}))
	jr := coc.JsonPtr()
	coc2 := corestr.New.CollectionsOfCollection.Create()
	err := coc2.JsonParseSelfInject(jr)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "COC JsonParseSelfInject", actual)
}

// =============================================================================
// HashsetsCollection — JSON methods (8 uncovered)
// =============================================================================

func Test_Cov66_HC_JsonModel(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Create()
	hc.Add(corestr.New.Hashset.Strings([]string{"x"}))
	m := hc.JsonModel()
	actual := args.Map{"nonNil": m != nil}
	expected := args.Map{"nonNil": true}
	expected.ShouldBeEqual(t, 0, "HC JsonModel", actual)
}

func Test_Cov66_HC_JsonModelAny(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Create()
	actual := args.Map{"nonNil": hc.JsonModelAny() != nil}
	expected := args.Map{"nonNil": true}
	expected.ShouldBeEqual(t, 0, "HC JsonModelAny", actual)
}

func Test_Cov66_HC_MarshalJSON(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Create()
	hc.Add(corestr.New.Hashset.Strings([]string{"x"}))
	b, err := hc.MarshalJSON()
	actual := args.Map{"noErr": err == nil, "nonEmpty": len(b) > 0}
	expected := args.Map{"noErr": true, "nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "HC MarshalJSON", actual)
}

func Test_Cov66_HC_UnmarshalJSON(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Create()
	hc.Add(corestr.New.Hashset.Strings([]string{"x"}))
	b, _ := hc.MarshalJSON()
	hc2 := corestr.New.HashsetsCollection.Create()
	err := hc2.UnmarshalJSON(b)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "HC UnmarshalJSON", actual)
}

func Test_Cov66_HC_ParseInjectUsingJson(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Create()
	hc.Add(corestr.New.Hashset.Strings([]string{"x"}))
	jr := hc.JsonPtr()
	hc2 := corestr.New.HashsetsCollection.Create()
	r, err := hc2.ParseInjectUsingJson(jr)
	actual := args.Map{"noErr": err == nil, "nonNil": r != nil}
	expected := args.Map{"noErr": true, "nonNil": true}
	expected.ShouldBeEqual(t, 0, "HC ParseInjectUsingJson", actual)
}

func Test_Cov66_HC_ParseInjectUsingJsonMust(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Create()
	hc.Add(corestr.New.Hashset.Strings([]string{"x"}))
	jr := hc.JsonPtr()
	hc2 := corestr.New.HashsetsCollection.Create()
	r := hc2.ParseInjectUsingJsonMust(jr)
	actual := args.Map{"nonNil": r != nil}
	expected := args.Map{"nonNil": true}
	expected.ShouldBeEqual(t, 0, "HC ParseInjectUsingJsonMust", actual)
}

func Test_Cov66_HC_JsonParseSelfInject(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Create()
	hc.Add(corestr.New.Hashset.Strings([]string{"x"}))
	jr := hc.JsonPtr()
	hc2 := corestr.New.HashsetsCollection.Create()
	err := hc2.JsonParseSelfInject(jr)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "HC JsonParseSelfInject", actual)
}

func Test_Cov66_HC_UnmarshalJSON_Error(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Create()
	err := hc.UnmarshalJSON([]byte("invalid"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "HC UnmarshalJSON error", actual)
}
