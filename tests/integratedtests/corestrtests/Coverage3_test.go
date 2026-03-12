package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Collection — additional methods ──

func Test_Cov3_Collection_AddString(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AddString("hello")
	actual := args.Map{"length": c.Length(), "hasAny": c.HasAnyItem()}
	expected := args.Map{"length": 1, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "Collection AddString returns 1 -- single item", actual)
}

func Test_Cov3_Collection_AddStringPtr(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	val := "hello"
	c.AddStringPtr(&val)
	actual := args.Map{"length": c.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "Collection AddStringPtr returns 1 -- single ptr", actual)
}

func Test_Cov3_Collection_AllIndividualsLength(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AddStrings("ab", "cde")
	actual := args.Map{"len": c.AllIndividualsLength()}
	expected := args.Map{"len": 5}
	expected.ShouldBeEqual(t, 0, "Collection AllIndividualsLength returns 5 -- ab+cde", actual)
}

// ── Hashmap — additional methods ──

func Test_Cov3_Hashmap_AddIf(t *testing.T) {
	h := corestr.New.Hashmap.Cap(5)
	h.AddIf(true, "yes", "value")
	h.AddIf(false, "no", "value")
	actual := args.Map{"length": h.Length(), "hasYes": h.Has("yes"), "hasNo": h.Has("no")}
	expected := args.Map{"length": 1, "hasYes": true, "hasNo": false}
	expected.ShouldBeEqual(t, 0, "Hashmap AddIf returns 1 -- conditional add", actual)
}

func Test_Cov3_Hashmap_GetOrDefault(t *testing.T) {
	h := corestr.New.Hashmap.Cap(5)
	h.Add("key", "value")
	actual := args.Map{"exists": h.Get("key"), "notExists": h.GetOrDefault("missing", "default")}
	expected := args.Map{"exists": "value", "notExists": "default"}
	expected.ShouldBeEqual(t, 0, "Hashmap GetOrDefault returns expected -- hit and miss", actual)
}

// ── Hashset — additional methods ──

func Test_Cov3_Hashset_Remove(t *testing.T) {
	h := corestr.New.Hashset.Cap(5)
	h.Adds("a", "b", "c")
	h.Remove("b")
	actual := args.Map{"length": h.Length(), "hasB": h.Has("b")}
	expected := args.Map{"length": 2, "hasB": false}
	expected.ShouldBeEqual(t, 0, "Hashset Remove returns 2 -- removed b", actual)
}

// ── LeftRight — additional methods ──

func Test_Cov3_LeftRight_Empty(t *testing.T) {
	lr := corestr.LeftRight{}
	actual := args.Map{"hasLeft": lr.HasLeft(), "hasRight": lr.HasRight(), "hasBoth": lr.HasBoth()}
	expected := args.Map{"hasLeft": false, "hasRight": false, "hasBoth": false}
	expected.ShouldBeEqual(t, 0, "LeftRight empty returns false -- all checks", actual)
}

func Test_Cov3_LeftRight_PartialLeft(t *testing.T) {
	lr := corestr.LeftRight{Left: "l"}
	actual := args.Map{"hasLeft": lr.HasLeft(), "hasRight": lr.HasRight(), "hasBoth": lr.HasBoth()}
	expected := args.Map{"hasLeft": true, "hasRight": false, "hasBoth": false}
	expected.ShouldBeEqual(t, 0, "LeftRight partial returns mixed -- only left", actual)
}

// ── LeftMiddleRight — additional methods ──

func Test_Cov3_LeftMiddleRight_Empty(t *testing.T) {
	lmr := corestr.LeftMiddleRight{}
	actual := args.Map{"hasLeft": lmr.HasLeft(), "hasMiddle": lmr.HasMiddle(), "hasRight": lmr.HasRight()}
	expected := args.Map{"hasLeft": false, "hasMiddle": false, "hasRight": false}
	expected.ShouldBeEqual(t, 0, "LeftMiddleRight empty returns false -- all checks", actual)
}

// ── KeyValuePair ──

func Test_Cov3_KeyValuePair(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	actual := args.Map{"key": kv.Key, "value": kv.Value}
	expected := args.Map{"key": "k", "value": "v"}
	expected.ShouldBeEqual(t, 0, "KeyValuePair returns expected -- valid pair", actual)
}

// ── KeyAnyValuePair ──

func Test_Cov3_KeyAnyValuePair(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: 42}
	actual := args.Map{"key": kv.Key, "value": kv.Value}
	expected := args.Map{"key": "k", "value": 42}
	expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns expected -- valid pair", actual)
}

// ── SimpleSlice — additional methods ──

func Test_Cov3_SimpleSlice_Clear(t *testing.T) {
	s := corestr.New.SimpleSlice.Cap(5)
	s.Adds("a", "b")
	s.Clear()
	actual := args.Map{"isEmpty": s.IsEmpty(), "length": s.Length()}
	expected := args.Map{"isEmpty": true, "length": 0}
	expected.ShouldBeEqual(t, 0, "SimpleSlice Clear returns empty -- after clear", actual)
}

// ── NewValidValue ──

func Test_Cov3_NewValidValue(t *testing.T) {
	vv := corestr.NewValidValue("hello")
	actual := args.Map{"value": vv.Value, "isValid": vv.IsValid}
	expected := args.Map{"value": "hello", "isValid": true}
	expected.ShouldBeEqual(t, 0, "NewValidValue returns valid -- non-empty string", actual)
}

func Test_Cov3_NewValidValue_Empty(t *testing.T) {
	vv := corestr.NewValidValue("")
	actual := args.Map{"value": vv.Value, "isValid": vv.IsValid}
	expected := args.Map{"value": "", "isValid": false}
	expected.ShouldBeEqual(t, 0, "NewValidValue returns invalid -- empty string", actual)
}

// ── AllIndividualsLengthOfSimpleSlices ──

func Test_Cov3_AllIndividualsLengthOfSimpleSlices(t *testing.T) {
	s1 := corestr.New.SimpleSlice.Cap(5)
	s1.Adds("ab", "cde")
	s2 := corestr.New.SimpleSlice.Cap(5)
	s2.Add("f")
	actual := args.Map{"result": corestr.AllIndividualsLengthOfSimpleSlices([]*corestr.SimpleSlice{s1, s2})}
	expected := args.Map{"result": 6}
	expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices returns 6 -- 3 strings", actual)
}

func Test_Cov3_AllIndividualsLengthOfSimpleSlices_Nil(t *testing.T) {
	actual := args.Map{"result": corestr.AllIndividualsLengthOfSimpleSlices(nil)}
	expected := args.Map{"result": 0}
	expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices returns 0 -- nil input", actual)
}
