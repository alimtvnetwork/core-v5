package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — IsEmpty / HasItems / Length
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_Hashmap_IsEmpty_New(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	actual := args.Map{"empty": hm.IsEmpty(), "items": hm.HasItems(), "len": hm.Length()}
	expected := args.Map{"empty": true, "items": false, "len": 0}
	expected.ShouldBeEqual(t, 0, "Hashmap empty", actual)
}

func Test_I29_Hashmap_HasAnyItem(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	actual := args.Map{"hasAny": hm.HasAnyItem(), "len": hm.Length()}
	expected := args.Map{"hasAny": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap HasAnyItem", actual)
}

func Test_I29_Hashmap_Length_Nil(t *testing.T) {
	var hm *corestr.Hashmap
	actual := args.Map{"len": hm.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Hashmap nil length", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — AddOrUpdate variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_Hashmap_AddOrUpdate(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	isNew1 := hm.AddOrUpdate("k", "v1")
	isNew2 := hm.AddOrUpdate("k", "v2")
	v, _ := hm.Get("k")
	actual := args.Map{"isNew1": isNew1, "isNew2": isNew2, "val": v}
	expected := args.Map{"isNew1": true, "isNew2": false, "val": "v2"}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdate", actual)
}

func Test_I29_Hashmap_Set(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	isNew := hm.Set("k", "v")
	actual := args.Map{"isNew": isNew}
	expected := args.Map{"isNew": true}
	expected.ShouldBeEqual(t, 0, "Hashmap Set", actual)
}

func Test_I29_Hashmap_SetTrim(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.SetTrim("  k  ", "  v  ")
	v, found := hm.Get("k")
	actual := args.Map{"found": found, "val": v}
	expected := args.Map{"found": true, "val": "v"}
	expected.ShouldBeEqual(t, 0, "Hashmap SetTrim", actual)
}

func Test_I29_Hashmap_SetBySplitter(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	isNew := hm.SetBySplitter("=", "key=value")
	v, _ := hm.Get("key")
	actual := args.Map{"isNew": isNew, "val": v}
	expected := args.Map{"isNew": true, "val": "value"}
	expected.ShouldBeEqual(t, 0, "Hashmap SetBySplitter", actual)
}

func Test_I29_Hashmap_SetBySplitter_NoSplit(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.SetBySplitter("=", "keyonly")
	v, found := hm.Get("keyonly")
	actual := args.Map{"found": found, "val": v}
	expected := args.Map{"found": true, "val": ""}
	expected.ShouldBeEqual(t, 0, "Hashmap SetBySplitter no split", actual)
}

func Test_I29_Hashmap_AddOrUpdateKeyStrValInt(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyStrValInt("k", 42)
	v, _ := hm.Get("k")
	actual := args.Map{"val": v}
	expected := args.Map{"val": "42"}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateKeyStrValInt", actual)
}

func Test_I29_Hashmap_AddOrUpdateKeyStrValFloat(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyStrValFloat("k", 3.14)
	v, found := hm.Get("k")
	actual := args.Map{"found": found, "notEmpty": v != ""}
	expected := args.Map{"found": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateKeyStrValFloat", actual)
}

func Test_I29_Hashmap_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyStrValFloat64("k", 2.718)
	v, found := hm.Get("k")
	actual := args.Map{"found": found, "notEmpty": v != ""}
	expected := args.Map{"found": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateKeyStrValFloat64", actual)
}

func Test_I29_Hashmap_AddOrUpdateKeyStrValAny(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyStrValAny("k", 99)
	_, found := hm.Get("k")
	actual := args.Map{"found": found}
	expected := args.Map{"found": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateKeyStrValAny", actual)
}

func Test_I29_Hashmap_AddOrUpdateKeyVal(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	kv := corestr.KeyValuePair{Key: "a", Value: "1"}
	isNew := hm.AddOrUpdateKeyVal(kv)
	actual := args.Map{"isNew": isNew}
	expected := args.Map{"isNew": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateKeyVal", actual)
}

func Test_I29_Hashmap_AddOrUpdateKeyValueAny(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	kav := corestr.KeyAnyValuePair{Key: "x", Value: "hello"}
	hm.AddOrUpdateKeyValueAny(kav)
	v, _ := hm.Get("x")
	actual := args.Map{"val": v}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateKeyValueAny", actual)
}

func Test_I29_Hashmap_AddOrUpdateHashmap_Nil(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	result := hm.AddOrUpdateHashmap(nil)
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateHashmap nil", actual)
}

func Test_I29_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	other := corestr.New.Hashmap.Cap(5)
	other.AddOrUpdate("b", "2")
	hm.AddOrUpdateHashmap(other)
	actual := args.Map{"len": hm.Length(), "hasB": hm.Has("b")}
	expected := args.Map{"len": 2, "hasB": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateHashmap", actual)
}

func Test_I29_Hashmap_AddOrUpdateMap_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdateMap(map[string]string{})
	actual := args.Map{"len": hm.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateMap empty", actual)
}

func Test_I29_Hashmap_AddOrUpdateMap(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdateMap(map[string]string{"x": "y"})
	actual := args.Map{"has": hm.Has("x")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateMap", actual)
}

func Test_I29_Hashmap_AddsOrUpdates_Nil(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	result := hm.AddsOrUpdates(nil...)
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddsOrUpdates nil", actual)
}

func Test_I29_Hashmap_AddsOrUpdates(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddsOrUpdates(
		corestr.KeyValuePair{Key: "a", Value: "1"},
		corestr.KeyValuePair{Key: "b", Value: "2"},
	)
	actual := args.Map{"len": hm.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashmap AddsOrUpdates", actual)
}

func Test_I29_Hashmap_AddOrUpdateKeyAnyValues_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyAnyValues()
	actual := args.Map{"empty": hm.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateKeyAnyValues empty", actual)
}

func Test_I29_Hashmap_AddOrUpdateKeyValues_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyValues()
	actual := args.Map{"empty": hm.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateKeyValues empty", actual)
}

func Test_I29_Hashmap_AddOrUpdateKeyValues(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyValues(
		corestr.KeyValuePair{Key: "a", Value: "1"},
	)
	actual := args.Map{"has": hm.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateKeyValues", actual)
}

func Test_I29_Hashmap_AddOrUpdateLock(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdateLock("k", "v")
	v, _ := hm.Get("k")
	actual := args.Map{"val": v}
	expected := args.Map{"val": "v"}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateLock", actual)
}

func Test_I29_Hashmap_AddOrUpdateWithWgLock(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	hm.AddOrUpdateWithWgLock("k", "v", wg)
	wg.Wait()
	actual := args.Map{"has": hm.Has("k")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateWithWgLock", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Has / Contains / IsKeyMissing
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_Hashmap_Has_Contains(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	actual := args.Map{"has": hm.Has("a"), "contains": hm.Contains("a"), "missing": hm.IsKeyMissing("b")}
	expected := args.Map{"has": true, "contains": true, "missing": true}
	expected.ShouldBeEqual(t, 0, "Hashmap Has/Contains", actual)
}

func Test_I29_Hashmap_ContainsLock(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	actual := args.Map{"cl": hm.ContainsLock("a"), "hl": hm.HasLock("a"), "hwl": hm.HasWithLock("a"), "mkl": hm.IsKeyMissingLock("z")}
	expected := args.Map{"cl": true, "hl": true, "hwl": true, "mkl": true}
	expected.ShouldBeEqual(t, 0, "Hashmap lock variants", actual)
}

func Test_I29_Hashmap_HasAllStrings(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	actual := args.Map{"all": hm.HasAllStrings("a", "b"), "missing": hm.HasAllStrings("a", "c")}
	expected := args.Map{"all": true, "missing": false}
	expected.ShouldBeEqual(t, 0, "Hashmap HasAllStrings", actual)
}

func Test_I29_Hashmap_HasAll(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	actual := args.Map{"all": hm.HasAll("a"), "miss": hm.HasAll("a", "z")}
	expected := args.Map{"all": true, "miss": false}
	expected.ShouldBeEqual(t, 0, "Hashmap HasAll", actual)
}

func Test_I29_Hashmap_HasAny(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	actual := args.Map{"any": hm.HasAny("z", "a"), "none": hm.HasAny("x", "y")}
	expected := args.Map{"any": true, "none": false}
	expected.ShouldBeEqual(t, 0, "Hashmap HasAny", actual)
}

func Test_I29_Hashmap_HasAllCollectionItems(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	coll := corestr.New.Collection.Strings([]string{"a"})
	actual := args.Map{"all": hm.HasAllCollectionItems(coll), "nil": hm.HasAllCollectionItems(nil)}
	expected := args.Map{"all": true, "nil": false}
	expected.ShouldBeEqual(t, 0, "Hashmap HasAllCollectionItems", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Get / GetValue
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_Hashmap_Get(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("k", "v")
	val, found := hm.Get("k")
	val2, found2 := hm.GetValue("k")
	_, notFound := hm.Get("missing")
	actual := args.Map{"val": val, "found": found, "val2": val2, "found2": found2, "notFound": notFound}
	expected := args.Map{"val": "v", "found": true, "val2": "v", "found2": true, "notFound": false}
	expected.ShouldBeEqual(t, 0, "Hashmap Get/GetValue", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Items / SafeItems / Keys / Values
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_Hashmap_Items(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("k", "v")
	actual := args.Map{"len": len(hm.Items())}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap Items", actual)
}

func Test_I29_Hashmap_SafeItems_Nil(t *testing.T) {
	var hm *corestr.Hashmap
	actual := args.Map{"nil": hm.SafeItems() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Hashmap SafeItems nil", actual)
}

func Test_I29_Hashmap_Keys(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	keys := hm.Keys()
	allKeys := hm.AllKeys()
	actual := args.Map{"keysLen": len(keys), "allKeysLen": len(allKeys)}
	expected := args.Map{"keysLen": 2, "allKeysLen": 2}
	expected.ShouldBeEqual(t, 0, "Hashmap Keys", actual)
}

func Test_I29_Hashmap_KeysCollection(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	kc := hm.KeysCollection()
	actual := args.Map{"len": kc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap KeysCollection", actual)
}

func Test_I29_Hashmap_ValuesList(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vl := hm.ValuesList()
	actual := args.Map{"len": len(vl), "val": vl[0]}
	expected := args.Map{"len": 1, "val": "1"}
	expected.ShouldBeEqual(t, 0, "Hashmap ValuesList", actual)
}

func Test_I29_Hashmap_ValuesCollection(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vc := hm.ValuesCollection()
	actual := args.Map{"len": vc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap ValuesCollection", actual)
}

func Test_I29_Hashmap_ValuesHashset(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vh := hm.ValuesHashset()
	actual := args.Map{"has": vh.Has("1")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap ValuesHashset", actual)
}

func Test_I29_Hashmap_Collection(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	coll := hm.Collection()
	actual := args.Map{"len": coll.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap Collection", actual)
}

func Test_I29_Hashmap_KeysValuesList(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	keys, vals := hm.KeysValuesList()
	actual := args.Map{"kLen": len(keys), "vLen": len(vals)}
	expected := args.Map{"kLen": 1, "vLen": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap KeysValuesList", actual)
}

func Test_I29_Hashmap_KeysValuesCollection(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	keys, vals := hm.KeysValuesCollection()
	actual := args.Map{"kLen": keys.Length(), "vLen": vals.Length()}
	expected := args.Map{"kLen": 1, "vLen": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap KeysValuesCollection", actual)
}

func Test_I29_Hashmap_KeysValuePairs(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	pairs := hm.KeysValuePairs()
	actual := args.Map{"len": len(pairs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap KeysValuePairs", actual)
}

func Test_I29_Hashmap_KeysValuePairsCollection(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	kvColl := hm.KeysValuePairsCollection()
	actual := args.Map{"len": kvColl.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap KeysValuePairsCollection", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Lock variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_Hashmap_IsEmptyLock(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	actual := args.Map{"empty": hm.IsEmptyLock()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap IsEmptyLock", actual)
}

func Test_I29_Hashmap_LengthLock(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	actual := args.Map{"len": hm.LengthLock()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap LengthLock", actual)
}

func Test_I29_Hashmap_KeysLock(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	kl := hm.KeysLock()
	actual := args.Map{"len": len(kl)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap KeysLock", actual)
}

func Test_I29_Hashmap_ValuesListCopyLock(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vl := hm.ValuesListCopyLock()
	actual := args.Map{"len": len(vl)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap ValuesListCopyLock", actual)
}

func Test_I29_Hashmap_ValuesCollectionLock(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vc := hm.ValuesCollectionLock()
	actual := args.Map{"len": vc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap ValuesCollectionLock", actual)
}

func Test_I29_Hashmap_ValuesHashsetLock(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vh := hm.ValuesHashsetLock()
	actual := args.Map{"has": vh.Has("1")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap ValuesHashsetLock", actual)
}

func Test_I29_Hashmap_ItemsCopyLock(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	cp := hm.ItemsCopyLock()
	actual := args.Map{"len": len(*cp)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap ItemsCopyLock", actual)
}

func Test_I29_Hashmap_KeysValuesListLock(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	keys, vals := hm.KeysValuesListLock()
	actual := args.Map{"kLen": len(keys), "vLen": len(vals)}
	expected := args.Map{"kLen": 1, "vLen": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap KeysValuesListLock", actual)
}

func Test_I29_Hashmap_StringLock(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	s := hm.StringLock()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap StringLock", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Diff / ConcatNew
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_Hashmap_ConcatNew_NoArgs(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	c := hm.ConcatNew(true)
	actual := args.Map{"len": c.Length(), "has": c.Has("a")}
	expected := args.Map{"len": 1, "has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap ConcatNew no args", actual)
}

func Test_I29_Hashmap_ConcatNew_WithArgs(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	other := corestr.New.Hashmap.Cap(5)
	other.AddOrUpdate("b", "2")
	c := hm.ConcatNew(true, other)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashmap ConcatNew with args", actual)
}

func Test_I29_Hashmap_ConcatNew_NilInArgs(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	c := hm.ConcatNew(true, nil)
	actual := args.Map{"has": c.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap ConcatNew nil in args", actual)
}

func Test_I29_Hashmap_ConcatNewUsingMaps_NoArgs(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	c := hm.ConcatNewUsingMaps(true)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap ConcatNewUsingMaps no args", actual)
}

func Test_I29_Hashmap_ConcatNewUsingMaps_WithArgs(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	c := hm.ConcatNewUsingMaps(true, map[string]string{"b": "2"}, nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashmap ConcatNewUsingMaps with args", actual)
}

func Test_I29_Hashmap_Diff(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	other := corestr.New.Hashmap.Cap(5)
	other.AddOrUpdate("a", "1")
	other.AddOrUpdate("b", "99")
	diff := hm.Diff(other)
	actual := args.Map{"hasDiff": diff.HasAnyItem()}
	expected := args.Map{"hasDiff": true}
	expected.ShouldBeEqual(t, 0, "Hashmap Diff", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Remove / Clear / Dispose
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_Hashmap_Remove(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.Remove("a")
	actual := args.Map{"has": hm.Has("a")}
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "Hashmap Remove", actual)
}

func Test_I29_Hashmap_RemoveWithLock(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.RemoveWithLock("a")
	actual := args.Map{"has": hm.Has("a")}
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "Hashmap RemoveWithLock", actual)
}

func Test_I29_Hashmap_Clear(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.Clear()
	actual := args.Map{"empty": hm.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap Clear", actual)
}

func Test_I29_Hashmap_Clear_Nil(t *testing.T) {
	var hm *corestr.Hashmap
	result := hm.Clear()
	actual := args.Map{"nil": result == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Hashmap Clear nil", actual)
}

func Test_I29_Hashmap_Dispose(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.Dispose()
	actual := args.Map{"empty": hm.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap Dispose", actual)
}

func Test_I29_Hashmap_Dispose_Nil(t *testing.T) {
	var hm *corestr.Hashmap
	hm.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Hashmap Dispose nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — IsEqual / Clone
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_Hashmap_IsEqualPtr_Same(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	actual := args.Map{"same": hm.IsEqualPtr(hm)}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "Hashmap IsEqualPtr same ptr", actual)
}

func Test_I29_Hashmap_IsEqualPtr_BothNil(t *testing.T) {
	var hm *corestr.Hashmap
	actual := args.Map{"eq": hm.IsEqualPtr(nil)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "Hashmap IsEqualPtr both nil", actual)
}

func Test_I29_Hashmap_IsEqualPtr_OneNil(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	actual := args.Map{"eq": hm.IsEqualPtr(nil)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "Hashmap IsEqualPtr one nil", actual)
}

func Test_I29_Hashmap_IsEqualPtr_BothEmpty(t *testing.T) {
	hm1 := corestr.New.Hashmap.Cap(5)
	hm2 := corestr.New.Hashmap.Cap(5)
	actual := args.Map{"eq": hm1.IsEqualPtr(hm2)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "Hashmap IsEqualPtr both empty", actual)
}

func Test_I29_Hashmap_IsEqualPtr_DiffLength(t *testing.T) {
	hm1 := corestr.New.Hashmap.Cap(5)
	hm1.AddOrUpdate("a", "1")
	hm2 := corestr.New.Hashmap.Cap(5)
	actual := args.Map{"eq": hm1.IsEqualPtr(hm2)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "Hashmap IsEqualPtr diff length", actual)
}

func Test_I29_Hashmap_IsEqualPtr_DiffValues(t *testing.T) {
	hm1 := corestr.New.Hashmap.Cap(5)
	hm1.AddOrUpdate("a", "1")
	hm2 := corestr.New.Hashmap.Cap(5)
	hm2.AddOrUpdate("a", "2")
	actual := args.Map{"eq": hm1.IsEqualPtr(hm2)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "Hashmap IsEqualPtr diff values", actual)
}

func Test_I29_Hashmap_IsEqual(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	cloned := hm.Clone()
	actual := args.Map{"eq": hm.IsEqual(cloned)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "Hashmap IsEqual", actual)
}

func Test_I29_Hashmap_IsEqualPtrLock(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	actual := args.Map{"eq": hm.IsEqualPtrLock(hm)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "Hashmap IsEqualPtrLock", actual)
}

func Test_I29_Hashmap_Clone(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	c := hm.Clone()
	actual := args.Map{"len": c.Length(), "has": c.Has("a")}
	expected := args.Map{"len": 1, "has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap Clone", actual)
}

func Test_I29_Hashmap_ClonePtr(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	cp := hm.ClonePtr()
	actual := args.Map{"notNil": cp != nil, "has": cp.Has("a")}
	expected := args.Map{"notNil": true, "has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap ClonePtr", actual)
}

func Test_I29_Hashmap_ClonePtr_Nil(t *testing.T) {
	var hm *corestr.Hashmap
	cp := hm.ClonePtr()
	actual := args.Map{"nil": cp == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Hashmap ClonePtr nil", actual)
}

func Test_I29_Hashmap_Clone_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	c := hm.Clone()
	actual := args.Map{"empty": c.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap Clone empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — String / Join / KeysToLower
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_Hashmap_String_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	actual := args.Map{"notEmpty": hm.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap String empty", actual)
}

func Test_I29_Hashmap_String_WithItems(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	actual := args.Map{"notEmpty": hm.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap String with items", actual)
}

func Test_I29_Hashmap_Join(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	actual := args.Map{"val": hm.Join(",")}
	expected := args.Map{"val": "1"}
	expected.ShouldBeEqual(t, 0, "Hashmap Join", actual)
}

func Test_I29_Hashmap_JoinKeys(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	actual := args.Map{"val": hm.JoinKeys(",")}
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "Hashmap JoinKeys", actual)
}

func Test_I29_Hashmap_KeysToLower(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("ABC", "1")
	lower := hm.KeysToLower()
	actual := args.Map{"has": lower.Has("abc")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap KeysToLower", actual)
}

func Test_I29_Hashmap_ValuesToLower(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("ABC", "1")
	lower := hm.ValuesToLower()
	actual := args.Map{"has": lower.Has("abc")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap ValuesToLower deprecated", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — JSON / Serialize / Deserialize
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_Hashmap_JsonModel(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	jm := hm.JsonModel()
	actual := args.Map{"len": len(jm)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap JsonModel", actual)
}

func Test_I29_Hashmap_JsonModelAny(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	actual := args.Map{"notNil": hm.JsonModelAny() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Hashmap JsonModelAny", actual)
}

func Test_I29_Hashmap_MarshalJSON(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	b, err := hm.MarshalJSON()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Hashmap MarshalJSON", actual)
}

func Test_I29_Hashmap_UnmarshalJSON(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	err := hm.UnmarshalJSON([]byte(`{"x":"y"}`))
	v, _ := hm.Get("x")
	actual := args.Map{"noErr": err == nil, "val": v}
	expected := args.Map{"noErr": true, "val": "y"}
	expected.ShouldBeEqual(t, 0, "Hashmap UnmarshalJSON", actual)
}

func Test_I29_Hashmap_UnmarshalJSON_Err(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	err := hm.UnmarshalJSON([]byte(`{invalid`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Hashmap UnmarshalJSON err", actual)
}

func Test_I29_Hashmap_Json(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	j := hm.Json()
	actual := args.Map{"hasBytes": j.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Hashmap Json", actual)
}

func Test_I29_Hashmap_JsonPtr(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	jp := hm.JsonPtr()
	actual := args.Map{"notNil": jp != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Hashmap JsonPtr", actual)
}

func Test_I29_Hashmap_Serialize(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	b, err := hm.Serialize()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Hashmap Serialize", actual)
}

func Test_I29_Hashmap_Deserialize(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	target := map[string]string{}
	err := hm.Deserialize(&target)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Hashmap Deserialize", actual)
}

func Test_I29_Hashmap_ParseInjectUsingJson(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	jr := hm.JsonPtr()
	hm2 := corestr.New.Hashmap.Cap(5)
	result, err := hm2.ParseInjectUsingJson(jr)
	actual := args.Map{"noErr": err == nil, "has": result.Has("a")}
	expected := args.Map{"noErr": true, "has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap ParseInjectUsingJson", actual)
}

func Test_I29_Hashmap_ParseInjectUsingJson_Err(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	badJson := corejson.NewPtr(42) // not a map
	_, err := hm.ParseInjectUsingJson(badJson)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Hashmap ParseInjectUsingJson err", actual)
}

func Test_I29_Hashmap_JsonParseSelfInject(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	jr := hm.JsonPtr()
	hm2 := corestr.New.Hashmap.Cap(5)
	err := hm2.JsonParseSelfInject(jr)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Hashmap JsonParseSelfInject", actual)
}

func Test_I29_Hashmap_AsJsoner(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	actual := args.Map{"notNil": hm.AsJsoner() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AsJsoner", actual)
}

func Test_I29_Hashmap_AsJsonContractsBinder(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	actual := args.Map{"notNil": hm.AsJsonContractsBinder() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AsJsonContractsBinder", actual)
}

func Test_I29_Hashmap_AsJsonParseSelfInjector(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	actual := args.Map{"notNil": hm.AsJsonParseSelfInjector() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AsJsonParseSelfInjector", actual)
}

func Test_I29_Hashmap_AsJsonMarshaller(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	actual := args.Map{"notNil": hm.AsJsonMarshaller() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AsJsonMarshaller", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — ToError / KeyValStringLines / ToStringsUsingCompiler
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_Hashmap_ToError(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	err := hm.ToError(", ")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Hashmap ToError", actual)
}

func Test_I29_Hashmap_ToDefaultError(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	err := hm.ToDefaultError()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Hashmap ToDefaultError", actual)
}

func Test_I29_Hashmap_KeyValStringLines(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	lines := hm.KeyValStringLines()
	actual := args.Map{"len": len(lines)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap KeyValStringLines", actual)
}

func Test_I29_Hashmap_ToStringsUsingCompiler_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	result := hm.ToStringsUsingCompiler(func(k, v string) string { return k + v })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Hashmap ToStringsUsingCompiler empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — GetValuesExcept / GetAllExcept / Filter
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_Hashmap_GetValuesExceptKeysInHashset(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	exclude := corestr.New.Hashset.Strings([]string{"a"})
	result := hm.GetValuesExceptKeysInHashset(exclude)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap GetValuesExceptKeysInHashset", actual)
}

func Test_I29_Hashmap_GetValuesExceptKeysInHashset_Nil(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	result := hm.GetValuesExceptKeysInHashset(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap GetValuesExceptKeysInHashset nil", actual)
}

func Test_I29_Hashmap_GetValuesKeysExcept(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	result := hm.GetValuesKeysExcept([]string{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap GetValuesKeysExcept", actual)
}

func Test_I29_Hashmap_GetValuesKeysExcept_Nil(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	result := hm.GetValuesKeysExcept(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap GetValuesKeysExcept nil", actual)
}

func Test_I29_Hashmap_GetAllExceptCollection(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	coll := corestr.New.Collection.Strings([]string{"a"})
	result := hm.GetAllExceptCollection(coll)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap GetAllExceptCollection", actual)
}

func Test_I29_Hashmap_GetAllExceptCollection_Nil(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	result := hm.GetAllExceptCollection(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap GetAllExceptCollection nil", actual)
}

func Test_I29_Hashmap_GetKeysFilteredItems(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("abc", "1")
	hm.AddOrUpdate("xyz", "2")
	result := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
		return s, len(s) == 3, false
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashmap GetKeysFilteredItems", actual)
}

func Test_I29_Hashmap_GetKeysFilteredItems_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	result := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
		return s, true, false
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Hashmap GetKeysFilteredItems empty", actual)
}

func Test_I29_Hashmap_GetKeysFilteredCollection(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	coll := hm.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
		return s, true, false
	})
	actual := args.Map{"len": coll.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap GetKeysFilteredCollection", actual)
}

func Test_I29_Hashmap_GetKeysFilteredCollection_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	coll := hm.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
		return s, true, false
	})
	actual := args.Map{"empty": coll.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap GetKeysFilteredCollection empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — AddOrUpdateCollection / AddOrUpdateStringsPtrWgLock
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_Hashmap_AddOrUpdateCollection_Nil(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	result := hm.AddOrUpdateCollection(nil, nil)
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateCollection nil", actual)
}

func Test_I29_Hashmap_AddOrUpdateCollection_Mismatch(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	keys := corestr.New.Collection.Strings([]string{"a", "b"})
	vals := corestr.New.Collection.Strings([]string{"1"})
	result := hm.AddOrUpdateCollection(keys, vals)
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateCollection mismatch", actual)
}

func Test_I29_Hashmap_AddOrUpdateCollection(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	keys := corestr.New.Collection.Strings([]string{"a"})
	vals := corestr.New.Collection.Strings([]string{"1"})
	hm.AddOrUpdateCollection(keys, vals)
	v, _ := hm.Get("a")
	actual := args.Map{"val": v}
	expected := args.Map{"val": "1"}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateCollection", actual)
}

func Test_I29_Hashmap_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(10)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	hm.AddOrUpdateStringsPtrWgLock(wg, []string{"a", "b"}, []string{"1", "2"})
	wg.Wait()
	actual := args.Map{"len": hm.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateStringsPtrWgLock", actual)
}

func Test_I29_Hashmap_AddOrUpdateStringsPtrWgLock_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	result := hm.AddOrUpdateStringsPtrWgLock(&sync.WaitGroup{}, []string{}, []string{})
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdateStringsPtrWgLock empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Filter variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_Hashmap_AddsOrUpdatesAnyUsingFilter_Nil(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	result := hm.AddsOrUpdatesAnyUsingFilter(nil, nil...)
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddsOrUpdatesAnyUsingFilter nil", actual)
}

func Test_I29_Hashmap_AddsOrUpdatesAnyUsingFilter_Break(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	filter := func(p corestr.KeyAnyValuePair) (string, bool, bool) {
		return p.ValueString(), true, true
	}
	hm.AddsOrUpdatesAnyUsingFilter(filter,
		corestr.KeyAnyValuePair{Key: "a", Value: "1"},
		corestr.KeyAnyValuePair{Key: "b", Value: "2"},
	)
	actual := args.Map{"len": hm.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap AddsOrUpdatesAnyUsingFilter break", actual)
}

func Test_I29_Hashmap_AddsOrUpdatesAnyUsingFilterLock_Nil(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	result := hm.AddsOrUpdatesAnyUsingFilterLock(nil, nil...)
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddsOrUpdatesAnyUsingFilterLock nil", actual)
}

func Test_I29_Hashmap_AddsOrUpdatesAnyUsingFilterLock_Break(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	filter := func(p corestr.KeyAnyValuePair) (string, bool, bool) {
		return p.ValueString(), true, true
	}
	hm.AddsOrUpdatesAnyUsingFilterLock(filter,
		corestr.KeyAnyValuePair{Key: "a", Value: "1"},
		corestr.KeyAnyValuePair{Key: "b", Value: "2"},
	)
	actual := args.Map{"len": hm.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap AddsOrUpdatesAnyUsingFilterLock break", actual)
}

func Test_I29_Hashmap_AddsOrUpdatesUsingFilter_Nil(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	result := hm.AddsOrUpdatesUsingFilter(nil, nil...)
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap AddsOrUpdatesUsingFilter nil", actual)
}

func Test_I29_Hashmap_AddsOrUpdatesUsingFilter_Break(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	filter := func(p corestr.KeyValuePair) (string, bool, bool) {
		return p.Value, true, true
	}
	hm.AddsOrUpdatesUsingFilter(filter,
		corestr.KeyValuePair{Key: "a", Value: "1"},
		corestr.KeyValuePair{Key: "b", Value: "2"},
	)
	actual := args.Map{"len": hm.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap AddsOrUpdatesUsingFilter break", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — GetKeysFilteredItems/Collection with break
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_Hashmap_GetKeysFilteredItems_Break(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	result := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
		return s, true, true // break on first
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap GetKeysFilteredItems break", actual)
}

func Test_I29_Hashmap_GetKeysFilteredCollection_Break(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	coll := hm.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
		return s, true, true // break on first
	})
	actual := args.Map{"len": coll.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap GetKeysFilteredCollection break", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// HashsetDataModel
// ══════════════════════════════════════════════════════════════════════════════

func Test_I29_HashsetDataModel_NewUsing(t *testing.T) {
	dm := &corestr.HashsetDataModel{Items: map[string]bool{"a": true}}
	hs := corestr.NewHashsetUsingDataModel(dm)
	actual := args.Map{"has": hs.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HashsetDataModel NewUsing", actual)
}

func Test_I29_HashsetDataModel_NewFromCollection(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	dm := corestr.NewHashsetsDataModelUsing(hs)
	actual := args.Map{"notNil": dm != nil, "len": len(dm.Items)}
	expected := args.Map{"notNil": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "HashsetDataModel NewFromCollection", actual)
}
