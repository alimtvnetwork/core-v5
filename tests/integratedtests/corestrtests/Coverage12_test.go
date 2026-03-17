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
	expected.ShouldBeEqual(t, 0, "Collection basic", actual)
}

func Test_Cov12_Collection_Add(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.Add("hello")
	c.Add("world")
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Collection.Add", actual)
}

func Test_Cov12_Collection_AddIf(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AddIf(true, "yes")
	c.AddIf(false, "no")
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Collection.AddIf", actual)
}

func Test_Cov12_Collection_Adds(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.Adds("a", "b", "c")
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Collection.Adds", actual)
}

func Test_Cov12_Collection_List(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	list := c.List()
	actual := args.Map{"len": len(list)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Collection.List", actual)
}

func Test_Cov12_Collection_String(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	result := c.String()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection.String", actual)
}

func Test_Cov12_Collection_IsEmpty_Nil(t *testing.T) {
	var c *corestr.Collection
	actual := args.Map{"empty": c.IsEmpty(), "len": c.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "Collection nil", actual)
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
	expected.ShouldBeEqual(t, 0, "Hashmap basic", actual)
}

func Test_Cov12_Hashmap_Get(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
	val, found := hm.Get("k1")
	actual := args.Map{"val": val, "found": found}
	expected := args.Map{"val": "v1", "found": true}
	expected.ShouldBeEqual(t, 0, "Hashmap.Get", actual)
}

func Test_Cov12_Hashmap_Get_NotFound(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
	val, found := hm.Get("k2")
	actual := args.Map{"val": val, "found": found}
	expected := args.Map{"val": "", "found": false}
	expected.ShouldBeEqual(t, 0, "Hashmap.Get not found", actual)
}

func Test_Cov12_Hashmap_AddOrUpdate(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	isNew := hm.AddOrUpdate("k1", "v1")
	isNew2 := hm.AddOrUpdate("k1", "v2")
	val, _ := hm.Get("k1")
	actual := args.Map{"isNew": isNew, "isUpdate": !isNew2, "val": val}
	expected := args.Map{"isNew": true, "isUpdate": true, "val": "v2"}
	expected.ShouldBeEqual(t, 0, "Hashmap.AddOrUpdate", actual)
}

func Test_Cov12_Hashmap_Clear(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
	hm.Clear()
	actual := args.Map{"empty": hm.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap.Clear", actual)
}

func Test_Cov12_Hashmap_Nil(t *testing.T) {
	var hm *corestr.Hashmap
	actual := args.Map{"empty": hm.IsEmpty(), "len": hm.Length(), "hasAny": hm.HasAnyItem()}
	expected := args.Map{"empty": true, "len": 0, "hasAny": false}
	expected.ShouldBeEqual(t, 0, "Hashmap nil", actual)
}

func Test_Cov12_Hashmap_Clone(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
	cloned := hm.ClonePtr()
	actual := args.Map{"notNil": cloned != nil, "has": cloned.Has("k1")}
	expected := args.Map{"notNil": true, "has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap.Clone", actual)
}

func Test_Cov12_Hashmap_ClonePtr_Nil(t *testing.T) {
	var hm *corestr.Hashmap
	cloned := hm.ClonePtr()
	actual := args.Map{"nil": cloned == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Hashmap.ClonePtr nil", actual)
}

func Test_Cov12_Hashmap_IsEqualPtr(t *testing.T) {
	a := corestr.New.Hashmap.StringsKv("k1", "v1")
	b := corestr.New.Hashmap.StringsKv("k1", "v1")
	actual := args.Map{"equal": a.IsEqualPtr(b)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "Hashmap.IsEqualPtr", actual)
}

func Test_Cov12_Hashmap_IsEqualPtr_Different(t *testing.T) {
	a := corestr.New.Hashmap.StringsKv("k1", "v1")
	b := corestr.New.Hashmap.StringsKv("k1", "v2")
	actual := args.Map{"equal": a.IsEqualPtr(b)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "Hashmap.IsEqualPtr different", actual)
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
	expected.ShouldBeEqual(t, 0, "Hashset basic", actual)
}

func Test_Cov12_Hashset_Add(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.Add("x")
	hs.Add("y")
	hs.Add("x") // duplicate
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashset.Add", actual)
}

func Test_Cov12_Hashset_Nil(t *testing.T) {
	var hs *corestr.Hashset
	actual := args.Map{"empty": hs.IsEmpty(), "len": hs.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "Hashset nil", actual)
}

// ── SimpleSlice ──

func Test_Cov12_SimpleSlice_Basic(t *testing.T) {
	ss := corestr.New.SimpleSlice.Strings("a", "b", "c")
	actual := args.Map{
		"len":     ss.Length(),
		"isEmpty": ss.IsEmpty(),
		"hasAny":  ss.HasAnyItem(),
	}
	expected := args.Map{
		"len": 3, "isEmpty": false, "hasAny": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleSlice basic", actual)
}

func Test_Cov12_SimpleSlice_Nil(t *testing.T) {
	var ss *corestr.SimpleSlice
	actual := args.Map{"empty": ss.IsEmpty(), "len": ss.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "SimpleSlice nil", actual)
}

// ── LeftRight ──

func Test_Cov12_LeftRight(t *testing.T) {
	lr := corestr.LeftRight{Left: "left", Right: "right"}
	actual := args.Map{"left": lr.Left, "right": lr.Right}
	expected := args.Map{"left": "left", "right": "right"}
	expected.ShouldBeEqual(t, 0, "LeftRight", actual)
}

// ── LeftMiddleRight ──

func Test_Cov12_LeftMiddleRight(t *testing.T) {
	lmr := corestr.LeftMiddleRight{Left: "l", Middle: "m", Right: "r"}
	actual := args.Map{"left": lmr.Left, "middle": lmr.Middle, "right": lmr.Right}
	expected := args.Map{"left": "l", "middle": "m", "right": "r"}
	expected.ShouldBeEqual(t, 0, "LeftMiddleRight", actual)
}

// ── ValidValue ──

func Test_Cov12_ValidValue(t *testing.T) {
	vv := corestr.ValidValue{Value: "hello", IsValid: true}
	actual := args.Map{"val": vv.Value, "valid": vv.IsValid}
	expected := args.Map{"val": "hello", "valid": true}
	expected.ShouldBeEqual(t, 0, "ValidValue", actual)
}

func Test_Cov12_ValidValue_Invalid(t *testing.T) {
	vv := corestr.ValidValue{IsValid: false}
	actual := args.Map{"valid": vv.IsValid}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "ValidValue invalid", actual)
}

// ── ValueStatus ──

func Test_Cov12_ValueStatus(t *testing.T) {
	vs := corestr.ValueStatus{Value: "hello", IsFound: true}
	actual := args.Map{"val": vs.Value, "found": vs.IsFound}
	expected := args.Map{"val": "hello", "found": true}
	expected.ShouldBeEqual(t, 0, "ValueStatus", actual)
}

// ── KeyValuePair ──

func Test_Cov12_KeyValuePair(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	actual := args.Map{"key": kv.Key, "val": kv.Value}
	expected := args.Map{"key": "k", "val": "v"}
	expected.ShouldBeEqual(t, 0, "KeyValuePair", actual)
}

// ── KeyAnyValuePair ──

func Test_Cov12_KeyAnyValuePair(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: 42}
	actual := args.Map{"key": kv.Key, "val": kv.Value}
	expected := args.Map{"key": "k", "val": 42}
	expected.ShouldBeEqual(t, 0, "KeyAnyValuePair", actual)
}

// ── emptyCreator ──

func Test_Cov12_Empty_Hashmap(t *testing.T) {
	hm := corestr.Empty.Hashmap()
	actual := args.Map{"notNil": hm != nil, "empty": hm.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.Hashmap", actual)
}

func Test_Cov12_Empty_Hashset(t *testing.T) {
	hs := corestr.Empty.Hashset()
	actual := args.Map{"notNil": hs != nil, "empty": hs.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.Hashset", actual)
}

// ── TextWithLineNumber ──

func Test_Cov12_TextWithLineNumber(t *testing.T) {
	tln := corestr.TextWithLineNumber{Text: "hello", LineNumber: 1}
	actual := args.Map{"text": tln.Text, "num": tln.LineNumber}
	expected := args.Map{"text": "hello", "num": 1}
	expected.ShouldBeEqual(t, 0, "TextWithLineNumber", actual)
}

// ── HashsetsCollection ──

func Test_Cov12_HashsetsCollection_Basic(t *testing.T) {
	hsc := corestr.New.HashsetsCollection.Cap(5)
	actual := args.Map{"isEmpty": hsc.IsEmpty(), "len": hsc.Length()}
	expected := args.Map{"isEmpty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "HashsetsCollection basic", actual)
}

func Test_Cov12_HashsetsCollection_Add(t *testing.T) {
	hsc := corestr.New.HashsetsCollection.Cap(5)
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	hsc.Add(hs)
	actual := args.Map{"len": hsc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "HashsetsCollection.Add", actual)
}

func Test_Cov12_HashsetsCollection_Nil(t *testing.T) {
	var hsc *corestr.HashsetsCollection
	actual := args.Map{"empty": hsc.IsEmpty(), "len": hsc.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "HashsetsCollection nil", actual)
}

// ── CollectionsOfCollection ──

func Test_Cov12_CollectionsOfCollection_Basic(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Cap(5)
	actual := args.Map{"isEmpty": coc.IsEmpty(), "len": coc.Length()}
	expected := args.Map{"isEmpty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "CollectionsOfCollection basic", actual)
}

func Test_Cov12_CollectionsOfCollection_Add(t *testing.T) {
	coc := corestr.New.CollectionsOfCollection.Cap(5)
	c := corestr.New.Collection.Strings("a", "b")
	coc.Add(c)
	actual := args.Map{"len": coc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CollectionsOfCollection.Add", actual)
}

func Test_Cov12_CollectionsOfCollection_Nil(t *testing.T) {
	var coc *corestr.CollectionsOfCollection
	actual := args.Map{"empty": coc.IsEmpty(), "len": coc.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "CollectionsOfCollection nil", actual)
}

// ── KeyValueCollection ──

func Test_Cov12_KeyValueCollection_Basic(t *testing.T) {
	kvc := corestr.New.KeyValues.Cap(5)
	actual := args.Map{"isEmpty": kvc.IsEmpty(), "len": kvc.Length()}
	expected := args.Map{"isEmpty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "KeyValueCollection basic", actual)
}

func Test_Cov12_KeyValueCollection_Add(t *testing.T) {
	kvc := corestr.New.KeyValues.Cap(5)
	kvc.Add("key", "val")
	actual := args.Map{"len": kvc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KeyValueCollection.Add", actual)
}

func Test_Cov12_KeyValueCollection_Nil(t *testing.T) {
	var kvc *corestr.KeyValueCollection
	actual := args.Map{"empty": kvc.IsEmpty(), "len": kvc.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "KeyValueCollection nil", actual)
}

// ── SimpleStringOnce ──

func Test_Cov12_SimpleStringOnce(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Value("hello")
	actual := args.Map{
		"val":     sso.Value(),
		"isEmpty": sso.IsEmpty(),
		"hasVal":  sso.HasValue(),
	}
	expected := args.Map{
		"val": "hello", "isEmpty": false, "hasVal": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleStringOnce", actual)
}

func Test_Cov12_SimpleStringOnce_Nil(t *testing.T) {
	var sso *corestr.SimpleStringOnce
	actual := args.Map{"empty": sso.IsEmpty(), "val": sso.Value()}
	expected := args.Map{"empty": true, "val": ""}
	expected.ShouldBeEqual(t, 0, "SimpleStringOnce nil", actual)
}

// ── CharCollectionMap ──

func Test_Cov12_CharCollectionMap_Basic(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	actual := args.Map{"isEmpty": ccm.IsEmpty(), "len": ccm.Length()}
	expected := args.Map{"isEmpty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "CharCollectionMap basic", actual)
}

func Test_Cov12_CharCollectionMap_Nil(t *testing.T) {
	var ccm *corestr.CharCollectionMap
	actual := args.Map{"empty": ccm.IsEmpty(), "len": ccm.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "CharCollectionMap nil", actual)
}

// ── CharHashsetMap ──

func Test_Cov12_CharHashsetMap_Basic(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(5)
	actual := args.Map{"isEmpty": chm.IsEmpty(), "len": chm.Length()}
	expected := args.Map{"isEmpty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "CharHashsetMap basic", actual)
}

func Test_Cov12_CharHashsetMap_Nil(t *testing.T) {
	var chm *corestr.CharHashsetMap
	actual := args.Map{"empty": chm.IsEmpty(), "len": chm.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "CharHashsetMap nil", actual)
}

// ── CloneSlice ──

func Test_Cov12_CloneSlice(t *testing.T) {
	original := []string{"a", "b", "c"}
	cloned := corestr.CloneSlice(original)
	original[0] = "X"
	actual := args.Map{"clonedFirst": cloned[0], "len": len(cloned)}
	expected := args.Map{"clonedFirst": "a", "len": 3}
	expected.ShouldBeEqual(t, 0, "CloneSlice", actual)
}

func Test_Cov12_CloneSlice_Nil(t *testing.T) {
	cloned := corestr.CloneSlice(nil)
	actual := args.Map{"nil": cloned == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "CloneSlice nil", actual)
}

func Test_Cov12_CloneSliceIf_True(t *testing.T) {
	original := []string{"a"}
	cloned := corestr.CloneSliceIf(true, original)
	original[0] = "X"
	actual := args.Map{"cloned": cloned[0]}
	expected := args.Map{"cloned": "a"}
	expected.ShouldBeEqual(t, 0, "CloneSliceIf true", actual)
}

func Test_Cov12_CloneSliceIf_False(t *testing.T) {
	original := []string{"a"}
	cloned := corestr.CloneSliceIf(false, original)
	actual := args.Map{"len": len(cloned)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CloneSliceIf false", actual)
}

// ── LinkedList ──

func Test_Cov12_LinkedList_Basic(t *testing.T) {
	ll := corestr.New.LinkedList.Cap(5)
	actual := args.Map{"isEmpty": ll.IsEmpty(), "len": ll.Length()}
	expected := args.Map{"isEmpty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "LinkedList basic", actual)
}

func Test_Cov12_LinkedList_Nil(t *testing.T) {
	var ll *corestr.LinkedList
	actual := args.Map{"empty": ll.IsEmpty(), "len": ll.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "LinkedList nil", actual)
}

// ── LinkedCollections ──

func Test_Cov12_LinkedCollections_Basic(t *testing.T) {
	lc := corestr.New.LinkedCollections.Cap(5)
	actual := args.Map{"isEmpty": lc.IsEmpty(), "len": lc.Length()}
	expected := args.Map{"isEmpty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "LinkedCollections basic", actual)
}

func Test_Cov12_LinkedCollections_Nil(t *testing.T) {
	var lc *corestr.LinkedCollections
	actual := args.Map{"empty": lc.IsEmpty(), "len": lc.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "LinkedCollections nil", actual)
}

// ── ValidValues ──

func Test_Cov12_ValidValues(t *testing.T) {
	vvs := corestr.ValidValues{
		Items: []corestr.ValidValue{
			{Value: "a", IsValid: true},
			{Value: "b", IsValid: false},
		},
	}
	actual := args.Map{"len": len(vvs.Items)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ValidValues", actual)
}

// ── HashmapDiff ──

func Test_Cov12_HashmapDiff(t *testing.T) {
	diff := corestr.HashmapDiff{Key: "k", Left: "a", Right: "b"}
	actual := args.Map{"key": diff.Key, "left": diff.Left, "right": diff.Right}
	expected := args.Map{"key": "k", "left": "a", "right": "b"}
	expected.ShouldBeEqual(t, 0, "HashmapDiff", actual)
}
