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
		"first":          dc.First().Value(),
		"last":           dc.Last().Value(),
		"firstOrDefault": dc.FirstOrDefault().Value(),
		"lastOrDefault":  dc.LastOrDefault().Value(),
	}
	expected := args.Map{
		"first":          "first",
		"last":           "last",
		"firstOrDefault": "first",
		"lastOrDefault":  "last",
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

func Test_Cov4_DynamicCollection_String(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello")
	s := dc.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.String returns non-empty -- single item", actual)
}

func Test_Cov4_DynamicCollection_AddAnyNonNull(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyNonNull(nil)
	dc.AddAnyNonNull("hello")
	actual := args.Map{"count": dc.Count()}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.AddAnyNonNull skips nil -- 1 valid", actual)
}

// ── CollectionTypes (constructors take capacity int, not slices) ──

func Test_Cov4_NewStringCollection(t *testing.T) {
	sc := coredynamic.NewStringCollection(5)
	actual := args.Map{"isEmpty": sc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewStringCollection returns empty -- capacity 5", actual)
}

func Test_Cov4_EmptyStringCollection(t *testing.T) {
	sc := coredynamic.EmptyStringCollection()
	actual := args.Map{"isEmpty": sc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "EmptyStringCollection returns empty -- new", actual)
}

func Test_Cov4_NewIntCollection(t *testing.T) {
	ic := coredynamic.NewIntCollection(5)
	actual := args.Map{"isEmpty": ic.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewIntCollection returns empty -- capacity 5", actual)
}

func Test_Cov4_EmptyIntCollection(t *testing.T) {
	ic := coredynamic.EmptyIntCollection()
	actual := args.Map{"isEmpty": ic.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "EmptyIntCollection returns empty -- new", actual)
}

func Test_Cov4_NewInt64Collection(t *testing.T) {
	c := coredynamic.NewInt64Collection(5)
	actual := args.Map{"isEmpty": c.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewInt64Collection returns empty -- capacity 5", actual)
}

func Test_Cov4_NewByteCollection(t *testing.T) {
	c := coredynamic.NewByteCollection(5)
	actual := args.Map{"isEmpty": c.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewByteCollection returns empty -- capacity 5", actual)
}

func Test_Cov4_NewBoolCollection(t *testing.T) {
	c := coredynamic.NewBoolCollection(5)
	actual := args.Map{"isEmpty": c.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewBoolCollection returns empty -- capacity 5", actual)
}

func Test_Cov4_NewFloat64Collection(t *testing.T) {
	c := coredynamic.NewFloat64Collection(5)
	actual := args.Map{"isEmpty": c.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewFloat64Collection returns empty -- capacity 5", actual)
}

func Test_Cov4_NewAnyMapCollection(t *testing.T) {
	c := coredynamic.NewAnyMapCollection(5)
	actual := args.Map{"isEmpty": c.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewAnyMapCollection returns empty -- capacity 5", actual)
}

func Test_Cov4_NewStringMapCollection(t *testing.T) {
	c := coredynamic.NewStringMapCollection(5)
	actual := args.Map{"isEmpty": c.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewStringMapCollection returns empty -- capacity 5", actual)
}

// ── KeyVal ──

func Test_Cov4_KeyVal_Basic(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "myKey", Val: "myVal"}
	actual := args.Map{
		"key":       kv.Key,
		"val":       kv.Val,
		"isKeyNull": kv.IsKeyNull(),
	}
	expected := args.Map{
		"key":       "myKey",
		"val":       "myVal",
		"isKeyNull": false,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal basic getters -- key and val set", actual)
}

// ── KeyValCollection ──

func Test_Cov4_KeyValCollection_Basic(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.KeyVal{Key: "k1", Val: "v1"})
	kvc.Add(coredynamic.KeyVal{Key: "k2", Val: "v2"})
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
	kvc.Add(coredynamic.KeyVal{Key: "b", Val: "2"})
	kvc.Add(coredynamic.KeyVal{Key: "a", Val: "1"})
	keys := kvc.AllKeys()
	actual := args.Map{"len": len(keys)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyValCollection.AllKeys returns 2 -- two items", actual)
}

func Test_Cov4_KeyValCollection_AllKeysSorted(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.KeyVal{Key: "b", Val: "2"})
	kvc.Add(coredynamic.KeyVal{Key: "a", Val: "1"})
	keys := kvc.AllKeysSorted()
	actual := args.Map{"first": keys[0]}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "KeyValCollection.AllKeysSorted first is a -- sorted", actual)
}

func Test_Cov4_KeyValCollection_AllValues(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.KeyVal{Key: "k", Val: "v"})
	vals := kvc.AllValues()
	actual := args.Map{"len": len(vals)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KeyValCollection.AllValues returns 1 -- single item", actual)
}

func Test_Cov4_KeyValCollection_String(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.KeyVal{Key: "k", Val: "v"})
	s := kvc.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection.String returns non-empty -- single item", actual)
}

// ── MapAnyItems ──

func Test_Cov4_MapAnyItems_Basic(t *testing.T) {
	m := coredynamic.NewMapAnyItems(5)
	m.Set("k1", "v1")
	m.Set("k2", 42)
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
	m := coredynamic.NewMapAnyItems(5)
	m.Set("k1", "v1")
	val := m.Get("k1")
	actual := args.Map{"val": val}
	expected := args.Map{"val": "v1"}
	expected.ShouldBeEqual(t, 0, "MapAnyItems.Get returns correct -- k1", actual)
}

func Test_Cov4_MapAnyItems_AllKeysSorted(t *testing.T) {
	m := coredynamic.NewMapAnyItems(5)
	m.Set("c", 3)
	m.Set("a", 1)
	m.Set("b", 2)
	keys := m.AllKeysSorted()
	actual := args.Map{"first": keys[0], "last": keys[2]}
	expected := args.Map{"first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "MapAnyItems.AllKeysSorted returns sorted -- 3 keys", actual)
}

func Test_Cov4_MapAnyItems_String(t *testing.T) {
	m := coredynamic.NewMapAnyItems(5)
	m.Set("k", "v")
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
	lr := &coredynamic.LeftRight{Left: "left", Right: "right"}
	actual := args.Map{
		"isEmpty":  lr.IsEmpty(),
		"hasAny":   lr.HasAnyItem(),
		"hasLeft":  lr.HasLeft(),
		"hasRight": lr.HasRight(),
	}
	expected := args.Map{
		"isEmpty":  false,
		"hasAny":   true,
		"hasLeft":  true,
		"hasRight": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight basic -- both set", actual)
}

func Test_Cov4_LeftRight_Empty(t *testing.T) {
	lr := &coredynamic.LeftRight{}
	actual := args.Map{"isEmpty": lr.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "LeftRight.IsEmpty returns true -- no values", actual)
}
