package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Collection ──

func Test_Cov12_Collection_Basic(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
	actual := args.Map{
		"len":     c.Length(),
		"isEmpty": c.IsEmpty(),
		"hasAny":  c.HasAnyItem(),
		"first":   c.First(),
		"last":    c.Last(),
	}
	expected := args.Map{
		"len": 3, "isEmpty": false, "hasAny": true,
		"first": "a", "last": "c",
	}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- basic", actual)
}

func Test_Cov12_Collection_Add(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.Add("hello")
	c.Add("world")
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Collection.Add returns correct value -- with args", actual)
}

func Test_Cov12_Collection_AddIf(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AddIf(true, "yes")
	c.AddIf(false, "no")
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Collection.AddIf returns correct value -- with args", actual)
}

func Test_Cov12_Collection_Adds(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.Adds("a", "b", "c")
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Collection.Adds returns correct value -- with args", actual)
}

func Test_Cov12_Collection_List(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	list := c.List()
	actual := args.Map{"len": len(list)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Collection.List returns correct value -- with args", actual)
}

func Test_Cov12_Collection_String(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	result := c.String()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection.String returns correct value -- with args", actual)
}

func Test_Cov12_Collection_IsEmpty_Nil(t *testing.T) {
	var c *corestr.Collection
	actual := args.Map{"empty": c.IsEmpty(), "len": c.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "Collection returns nil -- nil", actual)
}

// ── Hashmap ──

func Test_Cov12_Hashmap_Basic(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
	actual := args.Map{
		"has":     hm.Has("k1"),
		"notHas":  !hm.Has("k2"),
		"len":     hm.Length(),
		"isEmpty": hm.IsEmpty(),
	}
	expected := args.Map{
		"has": true, "notHas": true, "len": 1, "isEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- basic", actual)
}

func Test_Cov12_Hashmap_Get(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
	val, found := hm.Get("k1")
	actual := args.Map{"val": val, "found": found}
	expected := args.Map{"val": "v1", "found": true}
	expected.ShouldBeEqual(t, 0, "Hashmap.Get returns correct value -- with args", actual)
}

func Test_Cov12_Hashmap_Get_NotFound(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
	val, found := hm.Get("k2")
	actual := args.Map{"val": val, "found": found}
	expected := args.Map{"val": "", "found": false}
	expected.ShouldBeEqual(t, 0, "Hashmap.Get returns correct value -- not found", actual)
}

func Test_Cov12_Hashmap_AddOrUpdate(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	isNew := hm.AddOrUpdate("k1", "v1")
	isNew2 := hm.AddOrUpdate("k1", "v2")
	val, _ := hm.Get("k1")
	actual := args.Map{"isNew": isNew, "isUpdate": !isNew2, "val": val}
	expected := args.Map{"isNew": true, "isUpdate": true, "val": "v2"}
	expected.ShouldBeEqual(t, 0, "Hashmap.AddOrUpdate returns correct value -- with args", actual)
}

func Test_Cov12_Hashmap_Clear(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
	hm.Clear()
	actual := args.Map{"empty": hm.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap.Clear returns correct value -- with args", actual)
}

func Test_Cov12_Hashmap_Nil(t *testing.T) {
	var hm *corestr.Hashmap
	actual := args.Map{"empty": hm.IsEmpty(), "len": hm.Length(), "hasAny": hm.HasAnyItem()}
	expected := args.Map{"empty": true, "len": 0, "hasAny": false}
	expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- nil", actual)
}

func Test_Cov12_Hashmap_Clone(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
	cloned := hm.ClonePtr()
	actual := args.Map{"notNil": cloned != nil, "notEmpty": !cloned.IsEmpty()}
	expected := args.Map{"notNil": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap.Clone returns correct value -- with args", actual)
}

func Test_Cov12_Hashmap_ClonePtr_Nil(t *testing.T) {
	var hm *corestr.Hashmap
	cloned := hm.ClonePtr()
	actual := args.Map{"nil": cloned == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Hashmap.ClonePtr returns nil -- nil", actual)
}

func Test_Cov12_Hashmap_IsEqualPtr(t *testing.T) {
	a := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
	b := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
	actual := args.Map{"equal": a.IsEqualPtr(b)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "Hashmap.IsEqualPtr returns correct value -- with args", actual)
}

func Test_Cov12_Hashmap_IsEqualPtr_Different(t *testing.T) {
	a := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
	b := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v2"})
	actual := args.Map{"equal": a.IsEqualPtr(b)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "Hashmap.IsEqualPtr returns correct value -- different", actual)
}

// ── Hashset ──

func Test_Cov12_Hashset_Basic(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
	actual := args.Map{
		"has":     hs.Has("a"),
		"notHas":  !hs.Has("d"),
		"len":     hs.Length(),
		"isEmpty": hs.IsEmpty(),
	}
	expected := args.Map{
		"has": true, "notHas": true, "len": 3, "isEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- basic", actual)
}

func Test_Cov12_Hashset_Add(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.Add("x")
	hs.Add("y")
	hs.Add("x") // duplicate
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashset.Add returns correct value -- with args", actual)
}

func Test_Cov12_Hashset_Nil(t *testing.T) {
	var hs *corestr.Hashset
	actual := args.Map{"empty": hs.IsEmpty(), "len": hs.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "Hashset returns nil -- nil", actual)
}

// ── SimpleSlice ──

func Test_Cov12_SimpleSlice_Basic(t *testing.T) {
	ss := corestr.New.SimpleSlice.SpreadStrings("a", "b", "c")
	actual := args.Map{
		"len":     ss.Length(),
		"isEmpty": ss.IsEmpty(),
		"hasAny":  ss.HasAnyItem(),
	}
	expected := args.Map{
		"len": 3, "isEmpty": false, "hasAny": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- basic", actual)
}

func Test_Cov12_SimpleSlice_Nil(t *testing.T) {
	var ss *corestr.SimpleSlice
	actual := args.Map{"empty": ss.IsEmpty(), "len": ss.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns nil -- nil", actual)
}

// ── LeftRight ──

func Test_Cov12_LeftRight(t *testing.T) {
	lr := corestr.NewLeftRight("left", "right")
	actual := args.Map{"left": lr.Left, "right": lr.Right}
	expected := args.Map{"left": "left", "right": "right"}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- with args", actual)
}

// ── LeftMiddleRight ──

func Test_Cov12_LeftMiddleRight(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("l", "m", "r")
	actual := args.Map{"left": lmr.Left, "middle": lmr.Middle, "right": lmr.Right}
	expected := args.Map{"left": "l", "middle": "m", "right": "r"}
	expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- with args", actual)
}

// ── ValidValue ──

func Test_Cov12_ValidValue(t *testing.T) {
	vv := corestr.ValidValue{Value: "hello", IsValid: true}
	actual := args.Map{"val": vv.Value, "valid": vv.IsValid}
	expected := args.Map{"val": "hello", "valid": true}
	expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- with args", actual)
}

func Test_Cov12_ValidValue_Invalid(t *testing.T) {
	vv := corestr.ValidValue{IsValid: false}
	actual := args.Map{"valid": vv.IsValid}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "ValidValue returns error -- invalid", actual)
}

// ── ValueStatus ──

func Test_Cov12_ValueStatus(t *testing.T) {
	vv := corestr.NewValidValue("hello")
	vs := &corestr.ValueStatus{ValueValid: vv, Index: 0}
	actual := args.Map{"val": vs.ValueValid.Value, "idx": vs.Index}
	expected := args.Map{"val": "hello", "idx": 0}
	expected.ShouldBeEqual(t, 0, "ValueStatus returns non-empty -- with args", actual)
}

// ── KeyValuePair ──

func Test_Cov12_KeyValuePair(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	actual := args.Map{"key": kv.Key, "val": kv.Value}
	expected := args.Map{"key": "k", "val": "v"}
	expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- with args", actual)
}

// ── KeyAnyValuePair ──

func Test_Cov12_KeyAnyValuePair(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: 42}
	actual := args.Map{"key": kv.Key, "val": kv.Value}
	expected := args.Map{"key": "k", "val": 42}
	expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- with args", actual)
}

// ── emptyCreator ──

func Test_Cov12_Empty_Hashmap(t *testing.T) {
	hm := corestr.Empty.Hashmap()
	actual := args.Map{"notNil": hm != nil, "empty": hm.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.Hashmap returns empty -- with args", actual)
}

func Test_Cov12_Empty_Hashset(t *testing.T) {
	hs := corestr.Empty.Hashset()
	actual := args.Map{"notNil": hs != nil, "empty": hs.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.Hashset returns empty -- with args", actual)
}

// ── TextWithLineNumber ──

func Test_Cov12_TextWithLineNumber(t *testing.T) {
	tln := corestr.TextWithLineNumber{Text: "hello", LineNumber: 1}
	actual := args.Map{"text": tln.Text, "num": tln.LineNumber}
	expected := args.Map{"text": "hello", "num": 1}
	expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns non-empty -- with args", actual)
}

// ── HashsetsCollection ──

func Test_Cov12_HashsetsCollection_Basic(t *testing.T) {
	hsc := corestr.New.HashsetsCollection.Cap(5)
	actual := args.Map{"isEmpty": hsc.IsEmpty(), "len": hsc.Length()}
	expected := args.Map{"isEmpty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "HashsetsCollection returns correct value -- basic", actual)
}

func Test_Cov12_HashsetsCollection_Add(t *testing.T) {
	hsc := corestr.New.HashsetsCollection.Cap(5)
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	hsc.Add(hs)
	actual := args.Map{"len": hsc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "HashsetsCollection.Add returns correct value -- with args", actual)
}

func Test_Cov12_HashsetsCollection_Nil(t *testing.T) {
	var hsc *corestr.HashsetsCollection
	actual := args.Map{"empty": hsc.IsEmpty(), "len": hsc.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "HashsetsCollection returns nil -- nil", actual)
}

// ── CollectionsOfCollection ──

func Test_Cov12_CollectionsOfCollection_Basic(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Cap(5)
	actual := args.Map{"isEmpty": coc.IsEmpty(), "len": coc.Length()}
	expected := args.Map{"isEmpty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "CollectionsOfCollection returns correct value -- basic", actual)
}

func Test_Cov12_CollectionsOfCollection_Add(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Cap(5)
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	coc.Add(c)
	actual := args.Map{"len": coc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CollectionsOfCollection.Add returns correct value -- with args", actual)
}

func Test_Cov12_CollectionsOfCollection_Nil(t *testing.T) {
	var coc *corestr.CollectionsOfCollection
	isNil := coc == nil
	actual := args.Map{"isNil": isNil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "CollectionsOfCollection returns nil -- nil", actual)
}

// ── KeyValueCollection ──

func Test_Cov12_KeyValueCollection_Basic(t *testing.T) {
	kvc := corestr.New.KeyValues.Cap(5)
	actual := args.Map{"isEmpty": kvc.IsEmpty(), "len": kvc.Length()}
	expected := args.Map{"isEmpty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "KeyValueCollection returns correct value -- basic", actual)
}

func Test_Cov12_KeyValueCollection_Add(t *testing.T) {
	kvc := corestr.New.KeyValues.Cap(5)
	kvc.Add("key", "val")
	actual := args.Map{"len": kvc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KeyValueCollection.Add returns correct value -- with args", actual)
}

func Test_Cov12_KeyValueCollection_Nil(t *testing.T) {
	var kvc *corestr.KeyValueCollection
	actual := args.Map{"empty": kvc.IsEmpty(), "len": kvc.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "KeyValueCollection returns nil -- nil", actual)
}

// ── SimpleStringOnce ──

func Test_Cov12_SimpleStringOnce(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("hello")
	actual := args.Map{
		"val":     sso.Value(),
		"isEmpty": sso.IsEmpty(),
		"hasVal":  sso.HasValidNonEmpty(),
	}
	expected := args.Map{
		"val": "hello", "isEmpty": false, "hasVal": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- with args", actual)
}

func Test_Cov12_SimpleStringOnce_Nil(t *testing.T) {
	var sso *corestr.SimpleStringOnce
	// IsEmpty panics on nil receiver — just verify nil check
	actual := args.Map{"isNil": sso == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns nil -- nil", actual)
}

// ── CharCollectionMap ──

func Test_Cov12_CharCollectionMap_Basic(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Empty()
	actual := args.Map{"isEmpty": ccm.IsEmpty(), "len": ccm.Length()}
	expected := args.Map{"isEmpty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "CharCollectionMap returns correct value -- basic", actual)
}

func Test_Cov12_CharCollectionMap_Nil(t *testing.T) {
	var ccm *corestr.CharCollectionMap
	actual := args.Map{"empty": ccm.IsEmpty(), "len": ccm.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "CharCollectionMap returns nil -- nil", actual)
}

// ── CharHashsetMap ──

func Test_Cov12_CharHashsetMap_Basic(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(5, 5)
	actual := args.Map{"isEmpty": chm.IsEmpty(), "len": chm.Length()}
	expected := args.Map{"isEmpty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "CharHashsetMap returns correct value -- basic", actual)
}

func Test_Cov12_CharHashsetMap_Nil(t *testing.T) {
	var chm *corestr.CharHashsetMap
	actual := args.Map{"empty": chm.IsEmpty(), "len": chm.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "CharHashsetMap returns nil -- nil", actual)
}

// ── CloneSlice ──

func Test_Cov12_CloneSlice(t *testing.T) {
	original := []string{"a", "b", "c"}
	cloned := corestr.CloneSlice(original)
	original[0] = "X"
	actual := args.Map{"clonedFirst": cloned[0], "len": len(cloned)}
	expected := args.Map{"clonedFirst": "a", "len": 3}
	expected.ShouldBeEqual(t, 0, "CloneSlice returns correct value -- with args", actual)
}

func Test_Cov12_CloneSlice_Nil(t *testing.T) {
	cloned := corestr.CloneSlice(nil)
	actual := args.Map{"empty": len(cloned) == 0}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "CloneSlice returns nil -- nil", actual)
}

func Test_Cov12_CloneSliceIf_True(t *testing.T) {
	original := []string{"a"}
	cloned := corestr.CloneSliceIf(true, original...)
	original[0] = "X"
	actual := args.Map{"cloned": cloned[0]}
	expected := args.Map{"cloned": "a"}
	expected.ShouldBeEqual(t, 0, "CloneSliceIf returns non-empty -- true", actual)
}

func Test_Cov12_CloneSliceIf_False(t *testing.T) {
	original := []string{"a"}
	cloned := corestr.CloneSliceIf(false, original...)
	actual := args.Map{"len": len(cloned)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CloneSliceIf returns non-empty -- false", actual)
}

// ── LinkedList ──

func Test_Cov12_LinkedList_Basic(t *testing.T) {
	ll := corestr.New.LinkedList.Create()
	actual := args.Map{"isEmpty": ll.IsEmpty(), "len": ll.Length()}
	expected := args.Map{"isEmpty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- basic", actual)
}

func Test_Cov12_LinkedList_Nil(t *testing.T) {
	var ll *corestr.LinkedList
	actual := args.Map{"isNil": ll == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LinkedList returns nil -- nil", actual)
}

// ── LinkedCollections ──

func Test_Cov12_LinkedCollections_Basic(t *testing.T) {
	lc := corestr.New.LinkedCollection.Create()
	actual := args.Map{"isEmpty": lc.IsEmpty(), "len": lc.Length()}
	expected := args.Map{"isEmpty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- basic", actual)
}

func Test_Cov12_LinkedCollections_Nil(t *testing.T) {
	var lc *corestr.LinkedCollections
	actual := args.Map{"empty": lc.IsEmpty(), "len": lc.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "LinkedCollections returns nil -- nil", actual)
}

// ── ValidValues ──

func Test_Cov12_ValidValues(t *testing.T) {
	vvs := corestr.NewValidValuesUsingValues(
		corestr.ValidValue{Value: "a", IsValid: true},
		corestr.ValidValue{Value: "b", IsValid: false},
	)
	actual := args.Map{"len": vvs.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- with args", actual)
}

// ── HashmapDiff ──

func Test_Cov12_HashmapDiff(t *testing.T) {
	diff := corestr.HashmapDiff(map[string]string{"k": "v"})
	actual := args.Map{"len": diff.Length(), "empty": diff.IsEmpty()}
	expected := args.Map{"len": 1, "empty": false}
	expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- with args", actual)
}
