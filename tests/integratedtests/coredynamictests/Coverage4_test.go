package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── DynamicCollection basic ops ──

func Test_Cov4_DynamicCollection_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{
		"isEmpty": dc.IsEmpty(),
		"count":   dc.Count(),
		"hasAny":  dc.HasAnyItem(),
	}
	expected := args.Map{
		"isEmpty": true,
		"count":   0,
		"hasAny":  false,
	}
	expected.ShouldBeEqual(t, 0, "EmptyDynamicCollection returns empty -- new", actual)
}

func Test_Cov4_DynamicCollection_AddAny(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello")
	dc.AddAny(42)
	actual := args.Map{
		"count":     dc.Count(),
		"hasAny":    dc.HasAnyItem(),
		"isEmpty":   dc.IsEmpty(),
		"lastIndex": dc.LastIndex(),
		"hasIdx0":   dc.HasIndex(0),
		"hasIdx99":  dc.HasIndex(99),
	}
	expected := args.Map{
		"count":     2,
		"hasAny":    true,
		"isEmpty":   false,
		"lastIndex": 1,
		"hasIdx0":   true,
		"hasIdx99":  false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.AddAny works -- two items", actual)
}

func Test_Cov4_DynamicCollection_FirstLast(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("first")
	dc.AddAny("last")
	actual := args.Map{
		"first":           dc.First().Value(),
		"last":            dc.Last().Value(),
		"firstOrDefault":  dc.FirstOrDefault().Value(),
		"lastOrDefault":   dc.LastOrDefault().Value(),
	}
	expected := args.Map{
		"first":           "first",
		"last":            "last",
		"firstOrDefault":  "first",
		"lastOrDefault":   "last",
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection First/Last correct -- two items", actual)
}

func Test_Cov4_DynamicCollection_SkipTakeLimit(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a")
	dc.AddAny("b")
	dc.AddAny("c")
	actual := args.Map{
		"skipLen":  len(dc.Skip(1)),
		"takeLen":  len(dc.Take(2)),
		"limitLen": len(dc.Limit(1)),
	}
	expected := args.Map{
		"skipLen":  2,
		"takeLen":  2,
		"limitLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Skip/Take/Limit correct -- 3 items", actual)
}

func Test_Cov4_DynamicCollection_ListStrings(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello")
	dc.AddAny("world")
	strs := dc.ListStrings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.ListStrings returns 2 -- two items", actual)
}

func Test_Cov4_DynamicCollection_RemoveAt(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a")
	dc.AddAny("b")
	dc.AddAny("c")
	dc.RemoveAt(1)
	actual := args.Map{"count": dc.Count()}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.RemoveAt removes middle -- 3 to 2", actual)
}

func Test_Cov4_DynamicCollection_Strings(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello")
	strs := dc.Strings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.Strings returns 1 -- single item", actual)
}

func Test_Cov4_DynamicCollection_String(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello")
	s := dc.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.String returns non-empty -- single item", actual)
}

func Test_Cov4_DynamicCollection_Loop(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a")
	dc.AddAny("b")
	count := 0
	dc.Loop(func(index int, item coredynamic.Dynamic) {
		count++
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.Loop iterates all -- two items", actual)
}

func Test_Cov4_DynamicCollection_AddAnyNonNull(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyNonNull(nil)
	dc.AddAnyNonNull("hello")
	actual := args.Map{"count": dc.Count()}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.AddAnyNonNull skips nil -- 1 valid", actual)
}

func Test_Cov4_DynamicCollection_At(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello")
	d := dc.At(0)
	actual := args.Map{"val": d.Value()}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.At returns correct -- index 0", actual)
}

func Test_Cov4_DynamicCollection_Items(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a")
	items := dc.Items()
	actual := args.Map{"len": len(items)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.Items returns slice -- 1 item", actual)
}

func Test_Cov4_DynamicCollection_AnyItems(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello")
	items := dc.AnyItems()
	actual := args.Map{"len": len(items)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.AnyItems returns 1 -- single item", actual)
}

func Test_Cov4_DynamicCollection_SkipTakeCollections(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a")
	dc.AddAny("b")
	dc.AddAny("c")
	skipCol := dc.SkipCollection(1)
	takeCol := dc.TakeCollection(2)
	limitCol := dc.LimitCollection(1)
	safeLimitCol := dc.SafeLimitCollection(100)
	actual := args.Map{
		"skipLen":      skipCol.Count(),
		"takeLen":      takeCol.Count(),
		"limitLen":     limitCol.Count(),
		"safeLimitLen": safeLimitCol.Count(),
	}
	expected := args.Map{
		"skipLen":      2,
		"takeLen":      2,
		"limitLen":     1,
		"safeLimitLen": 3,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Skip/Take/Limit collections -- 3 items", actual)
}

// ── CollectionTypes ──

func Test_Cov4_NewStringCollection(t *testing.T) {
	sc := coredynamic.NewStringCollection([]string{"a", "b"})
	actual := args.Map{"count": sc.Count()}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "NewStringCollection returns correct count -- 2 items", actual)
}

func Test_Cov4_EmptyStringCollection(t *testing.T) {
	sc := coredynamic.EmptyStringCollection()
	actual := args.Map{"isEmpty": sc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "EmptyStringCollection returns empty -- new", actual)
}

func Test_Cov4_NewIntCollection(t *testing.T) {
	ic := coredynamic.NewIntCollection([]int{1, 2, 3})
	actual := args.Map{"count": ic.Count()}
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "NewIntCollection returns correct count -- 3 items", actual)
}

func Test_Cov4_EmptyIntCollection(t *testing.T) {
	ic := coredynamic.EmptyIntCollection()
	actual := args.Map{"isEmpty": ic.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "EmptyIntCollection returns empty -- new", actual)
}

func Test_Cov4_NewInt64Collection(t *testing.T) {
	c := coredynamic.NewInt64Collection([]int64{1, 2})
	actual := args.Map{"count": c.Count()}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "NewInt64Collection returns correct count -- 2 items", actual)
}

func Test_Cov4_NewByteCollection(t *testing.T) {
	c := coredynamic.NewByteCollection([]byte{1, 2, 3})
	actual := args.Map{"count": c.Count()}
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "NewByteCollection returns correct count -- 3 bytes", actual)
}

func Test_Cov4_NewBoolCollection(t *testing.T) {
	c := coredynamic.NewBoolCollection([]bool{true, false})
	actual := args.Map{"count": c.Count()}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "NewBoolCollection returns correct count -- 2 items", actual)
}

func Test_Cov4_NewFloat64Collection(t *testing.T) {
	c := coredynamic.NewFloat64Collection([]float64{1.1, 2.2})
	actual := args.Map{"count": c.Count()}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "NewFloat64Collection returns correct count -- 2 items", actual)
}

func Test_Cov4_NewAnyMapCollection(t *testing.T) {
	c := coredynamic.NewAnyMapCollection(map[string]any{"k": "v"})
	actual := args.Map{"count": c.Count()}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "NewAnyMapCollection returns correct count -- 1 item", actual)
}

func Test_Cov4_NewStringMapCollection(t *testing.T) {
	c := coredynamic.NewStringMapCollection(map[string]string{"k": "v"})
	actual := args.Map{"count": c.Count()}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "NewStringMapCollection returns correct count -- 1 item", actual)
}

// ── KeyVal ──

func Test_Cov4_KeyVal_Basic(t *testing.T) {
	kv := coredynamic.NewKeyVal("myKey", "myVal")
	actual := args.Map{
		"key":         kv.Key,
		"val":         kv.Val,
		"isKeyNull":   kv.IsKeyNullOrEmptyString(),
	}
	expected := args.Map{
		"key":         "myKey",
		"val":         "myVal",
		"isKeyNull":   false,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal basic getters -- key and val set", actual)
}

func Test_Cov4_KeyVal_IsKeyNull_Empty(t *testing.T) {
	kv := coredynamic.NewKeyVal("", "val")
	actual := args.Map{"isNull": kv.IsKeyNullOrEmptyString()}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "KeyVal.IsKeyNullOrEmptyString returns true -- empty key", actual)
}

func Test_Cov4_KeyVal_IsKeyNull_Nil(t *testing.T) {
	kv := coredynamic.NewKeyVal(nil, "val")
	actual := args.Map{"isNull": kv.IsKeyNullOrEmptyString()}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "KeyVal.IsKeyNullOrEmptyString returns true -- nil key", actual)
}

// ── KeyValCollection ──

func Test_Cov4_KeyValCollection_Basic(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.NewKeyVal("k1", "v1"))
	kvc.Add(coredynamic.NewKeyVal("k2", "v2"))
	actual := args.Map{
		"length":  kvc.Length(),
		"hasAny":  kvc.HasAnyItem(),
		"isEmpty": kvc.IsEmpty(),
	}
	expected := args.Map{
		"length":  2,
		"hasAny":  true,
		"isEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection basic ops -- 2 items", actual)
}

func Test_Cov4_KeyValCollection_AllKeys(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.NewKeyVal("b", "2"))
	kvc.Add(coredynamic.NewKeyVal("a", "1"))
	keys := kvc.AllKeys()
	actual := args.Map{"len": len(keys)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyValCollection.AllKeys returns 2 -- two items", actual)
}

func Test_Cov4_KeyValCollection_AllKeysSorted(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.NewKeyVal("b", "2"))
	kvc.Add(coredynamic.NewKeyVal("a", "1"))
	keys := kvc.AllKeysSorted()
	actual := args.Map{"first": keys[0]}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "KeyValCollection.AllKeysSorted first is a -- sorted", actual)
}

func Test_Cov4_KeyValCollection_AllValues(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.NewKeyVal("k", "v"))
	vals := kvc.AllValues()
	actual := args.Map{"len": len(vals)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KeyValCollection.AllValues returns 1 -- single item", actual)
}

func Test_Cov4_KeyValCollection_Items(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.NewKeyVal("k", "v"))
	items := kvc.Items()
	actual := args.Map{"len": len(items)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KeyValCollection.Items returns 1 -- single item", actual)
}

func Test_Cov4_KeyValCollection_String(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.NewKeyVal("k", "v"))
	s := kvc.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection.String returns non-empty -- single item", actual)
}

func Test_Cov4_KeyValCollection_Clone(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.NewKeyVal("k", "v"))
	cloned := kvc.Clone()
	actual := args.Map{"sameLen": cloned.Length() == kvc.Length()}
	expected := args.Map{"sameLen": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection.Clone returns same len -- single item", actual)
}

func Test_Cov4_KeyValCollection_ClonePtr(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.NewKeyVal("k", "v"))
	cloned := kvc.ClonePtr()
	actual := args.Map{
		"notNil":     cloned != nil,
		"notSamePtr": cloned != kvc,
	}
	expected := args.Map{
		"notNil":     true,
		"notSamePtr": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection.ClonePtr returns different ptr -- single item", actual)
}

func Test_Cov4_KeyValCollection_ClonePtr_Nil(t *testing.T) {
	var kvc *coredynamic.KeyValCollection
	cloned := kvc.ClonePtr()
	actual := args.Map{"isNil": cloned == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection.ClonePtr returns nil -- nil receiver", actual)
}

// ── MapAnyItems ──

func Test_Cov4_MapAnyItems_Basic(t *testing.T) {
	m := coredynamic.NewMapAnyItems(map[string]any{"k1": "v1", "k2": 42})
	actual := args.Map{
		"length":   m.Length(),
		"hasItems": m.HasAnyItem(),
		"isEmpty":  m.IsEmpty(),
		"hasK1":    m.Has("k1"),
		"hasX":     m.Has("x"),
	}
	expected := args.Map{
		"length":   2,
		"hasItems": true,
		"isEmpty":  false,
		"hasK1":    true,
		"hasX":     false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems basic ops -- 2 items", actual)
}

func Test_Cov4_MapAnyItems_Get(t *testing.T) {
	m := coredynamic.NewMapAnyItems(map[string]any{"k1": "v1"})
	val := m.Get("k1")
	actual := args.Map{"val": val}
	expected := args.Map{"val": "v1"}
	expected.ShouldBeEqual(t, 0, "MapAnyItems.Get returns correct -- k1", actual)
}

func Test_Cov4_MapAnyItems_Set(t *testing.T) {
	m := coredynamic.NewMapAnyItems(map[string]any{})
	m.Set("k1", "v1")
	actual := args.Map{"has": m.Has("k1"), "val": m.Get("k1")}
	expected := args.Map{"has": true, "val": "v1"}
	expected.ShouldBeEqual(t, 0, "MapAnyItems.Set adds item -- k1", actual)
}

func Test_Cov4_MapAnyItems_AllKeysSorted(t *testing.T) {
	m := coredynamic.NewMapAnyItems(map[string]any{"c": 3, "a": 1, "b": 2})
	keys := m.AllKeysSorted()
	actual := args.Map{"first": keys[0], "last": keys[2]}
	expected := args.Map{"first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "MapAnyItems.AllKeysSorted returns sorted -- 3 keys", actual)
}

func Test_Cov4_MapAnyItems_AllValues(t *testing.T) {
	m := coredynamic.NewMapAnyItems(map[string]any{"k": "v"})
	vals := m.AllValues()
	actual := args.Map{"len": len(vals)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems.AllValues returns 1 -- single item", actual)
}

func Test_Cov4_MapAnyItems_String(t *testing.T) {
	m := coredynamic.NewMapAnyItems(map[string]any{"k": "v"})
	s := m.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems.String returns non-empty -- single item", actual)
}

func Test_Cov4_MapAnyItems_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItems
	actual := args.Map{
		"isEmpty": m.IsEmpty(),
		"length":  m.Length(),
	}
	expected := args.Map{
		"isEmpty": true,
		"length":  0,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems nil receiver safe -- isEmpty and length", actual)
}

// ── LeftRight ──

func Test_Cov4_LeftRight_Basic(t *testing.T) {
	lr := coredynamic.NewLeftRight("left", "right")
	actual := args.Map{
		"left":    lr.Left,
		"right":   lr.Right,
		"isEqual": lr.IsEqual(),
	}
	expected := args.Map{
		"left":    "left",
		"right":   "right",
		"isEqual": false,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight basic -- different values", actual)
}

func Test_Cov4_LeftRight_Equal(t *testing.T) {
	lr := coredynamic.NewLeftRight("same", "same")
	actual := args.Map{"isEqual": lr.IsEqual()}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "LeftRight.IsEqual returns true -- same values", actual)
}

// ── MapAnyItemDiff ──

func Test_Cov4_MapAnyItemDiff_Basic(t *testing.T) {
	diff := coredynamic.NewMapAnyItemDiff(map[string]any{"k": "v"})
	actual := args.Map{
		"length":  diff.Length(),
		"hasAny":  diff.HasAnyItem(),
		"isEmpty": diff.IsEmpty(),
	}
	expected := args.Map{
		"length":  1,
		"hasAny":  true,
		"isEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff basic ops -- 1 item", actual)
}

func Test_Cov4_MapAnyItemDiff_AllKeysSorted(t *testing.T) {
	diff := coredynamic.NewMapAnyItemDiff(map[string]any{"b": 2, "a": 1})
	keys := diff.AllKeysSorted()
	actual := args.Map{"first": keys[0]}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff.AllKeysSorted returns sorted -- 2 keys", actual)
}

func Test_Cov4_MapAnyItemDiff_Clear(t *testing.T) {
	diff := coredynamic.NewMapAnyItemDiff(map[string]any{"k": "v"})
	diff.Clear()
	actual := args.Map{"isEmpty": diff.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff.Clear empties -- after clear", actual)
}

// ── TypeStatus ──

func Test_Cov4_TypeStatus_Basic(t *testing.T) {
	ts := coredynamic.NewTypeStatus("hello")
	actual := args.Map{
		"isValid":    ts.IsValid,
		"isPtr":      ts.IsPointer,
		"typeName":   ts.TypeName != "",
	}
	expected := args.Map{
		"isValid":    true,
		"isPtr":      false,
		"typeName":   true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus basic -- string value", actual)
}

func Test_Cov4_TypeStatus_Nil(t *testing.T) {
	ts := coredynamic.NewTypeStatus(nil)
	actual := args.Map{"isValid": ts.IsValid}
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns invalid -- nil input", actual)
}

// ── Dynamic advanced ──

func Test_Cov4_Dynamic_ReflectValue(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	rv := d.ReflectValue()
	actual := args.Map{"notNil": rv != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic.ReflectValue returns non-nil -- valid string", actual)
}

func Test_Cov4_Dynamic_ReflectType(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	rt := d.ReflectType()
	actual := args.Map{"notNil": rt != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic.ReflectType returns non-nil -- valid string", actual)
}
