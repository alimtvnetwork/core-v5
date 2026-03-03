package coregenerictests

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coredata/coregeneric"
)

// =============================================================================
// Hashmap — Constructors
// =============================================================================

func Test_Hashmap_Empty(t *testing.T) {
	convey.Convey("EmptyHashmap creates empty map", t, func() {
		hm := coregeneric.EmptyHashmap[string, int]()
		convey.So(hm.IsEmpty(), should.BeTrue)
		convey.So(hm.Length(), should.Equal, 0)
		convey.So(hm.HasItems(), should.BeFalse)
	})
}

func Test_Hashmap_New(t *testing.T) {
	convey.Convey("NewHashmap creates with capacity", t, func() {
		hm := coregeneric.NewHashmap[string, int](10)
		convey.So(hm.IsEmpty(), should.BeTrue)
	})
}

func Test_Hashmap_From(t *testing.T) {
	convey.Convey("HashmapFrom wraps existing map", t, func() {
		hm := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
		convey.So(hm.Length(), should.Equal, 2)
		convey.So(hm.Has("a"), should.BeTrue)
	})
}

func Test_Hashmap_Clone(t *testing.T) {
	convey.Convey("HashmapClone creates independent copy", t, func() {
		original := coregeneric.HashmapFrom(map[string]int{"k": 1})
		cloned := coregeneric.HashmapClone(original.Map())
		cloned.Set("k", 99)

		origVal, _ := original.Get("k")
		clonedVal, _ := cloned.Get("k")
		convey.So(origVal, should.Equal, 1)
		convey.So(clonedVal, should.Equal, 99)
	})
}

// =============================================================================
// Hashmap — Set / Get
// =============================================================================

func Test_Hashmap_Set_NewKey(t *testing.T) {
	convey.Convey("Hashmap.Set returns true for new key", t, func() {
		hm := coregeneric.EmptyHashmap[string, int]()
		isNew := hm.Set("key", 42)
		convey.So(isNew, should.BeTrue)
		convey.So(hm.Length(), should.Equal, 1)
	})
}

func Test_Hashmap_Set_ExistingKey(t *testing.T) {
	convey.Convey("Hashmap.Set returns false for existing key", t, func() {
		hm := coregeneric.HashmapFrom(map[string]int{"key": 1})
		isNew := hm.Set("key", 2)
		convey.So(isNew, should.BeFalse)
		val, _ := hm.Get("key")
		convey.So(val, should.Equal, 2)
	})
}

func Test_Hashmap_Get_Found(t *testing.T) {
	convey.Convey("Hashmap.Get returns value and true", t, func() {
		hm := coregeneric.HashmapFrom(map[string]int{"k": 42})
		val, found := hm.Get("k")
		convey.So(found, should.BeTrue)
		convey.So(val, should.Equal, 42)
	})
}

func Test_Hashmap_Get_NotFound(t *testing.T) {
	convey.Convey("Hashmap.Get returns zero and false for missing", t, func() {
		hm := coregeneric.EmptyHashmap[string, int]()
		val, found := hm.Get("missing")
		convey.So(found, should.BeFalse)
		convey.So(val, should.Equal, 0)
	})
}

func Test_Hashmap_GetOrDefault(t *testing.T) {
	convey.Convey("Hashmap.GetOrDefault returns default for missing key", t, func() {
		hm := coregeneric.EmptyHashmap[string, int]()
		convey.So(hm.GetOrDefault("x", 99), should.Equal, 99)
	})
}

func Test_Hashmap_GetOrDefault_Found(t *testing.T) {
	convey.Convey("Hashmap.GetOrDefault returns value for existing key", t, func() {
		hm := coregeneric.HashmapFrom(map[string]int{"x": 5})
		convey.So(hm.GetOrDefault("x", 99), should.Equal, 5)
	})
}

// =============================================================================
// Hashmap — Has / Contains / IsKeyMissing
// =============================================================================

func Test_Hashmap_Has(t *testing.T) {
	convey.Convey("Hashmap.Has returns true for existing key", t, func() {
		hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
		convey.So(hm.Has("a"), should.BeTrue)
		convey.So(hm.Contains("a"), should.BeTrue)
		convey.So(hm.IsKeyMissing("a"), should.BeFalse)
	})
}

func Test_Hashmap_IsKeyMissing(t *testing.T) {
	convey.Convey("Hashmap.IsKeyMissing returns true for missing key", t, func() {
		hm := coregeneric.EmptyHashmap[string, int]()
		convey.So(hm.IsKeyMissing("x"), should.BeTrue)
	})
}

// =============================================================================
// Hashmap — Remove
// =============================================================================

func Test_Hashmap_Remove_Existing(t *testing.T) {
	convey.Convey("Hashmap.Remove returns true for existing key", t, func() {
		hm := coregeneric.HashmapFrom(map[string]int{"k": 1})
		existed := hm.Remove("k")
		convey.So(existed, should.BeTrue)
		convey.So(hm.IsEmpty(), should.BeTrue)
	})
}

func Test_Hashmap_Remove_Missing(t *testing.T) {
	convey.Convey("Hashmap.Remove returns false for missing key", t, func() {
		hm := coregeneric.EmptyHashmap[string, int]()
		convey.So(hm.Remove("x"), should.BeFalse)
	})
}

// =============================================================================
// Hashmap — Keys / Values / Map
// =============================================================================

func Test_Hashmap_Keys(t *testing.T) {
	convey.Convey("Hashmap.Keys returns all keys", t, func() {
		hm := coregeneric.HashmapFrom(map[int]string{1: "a", 2: "b"})
		keys := hm.Keys()
		convey.So(len(keys), should.Equal, 2)
	})
}

func Test_Hashmap_Keys_Empty(t *testing.T) {
	convey.Convey("Hashmap.Keys returns empty for empty map", t, func() {
		hm := coregeneric.EmptyHashmap[int, string]()
		convey.So(hm.Keys(), should.Resemble, []int{})
	})
}

func Test_Hashmap_Values(t *testing.T) {
	convey.Convey("Hashmap.Values returns all values", t, func() {
		hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
		vals := hm.Values()
		convey.So(len(vals), should.Equal, 1)
		convey.So(vals[0], should.Equal, 1)
	})
}

func Test_Hashmap_Values_Empty(t *testing.T) {
	convey.Convey("Hashmap.Values returns empty for empty map", t, func() {
		hm := coregeneric.EmptyHashmap[string, int]()
		convey.So(hm.Values(), should.Resemble, []int{})
	})
}

func Test_Hashmap_Map(t *testing.T) {
	convey.Convey("Hashmap.Map returns underlying map", t, func() {
		m := map[string]int{"x": 1}
		hm := coregeneric.HashmapFrom(m)
		convey.So(hm.Map(), should.Equal, m)
	})
}

// =============================================================================
// Hashmap — AddOrUpdateMap / AddOrUpdateHashmap
// =============================================================================

func Test_Hashmap_AddOrUpdateMap(t *testing.T) {
	convey.Convey("Hashmap.AddOrUpdateMap merges entries", t, func() {
		hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
		hm.AddOrUpdateMap(map[string]int{"b": 2, "a": 10})
		convey.So(hm.Length(), should.Equal, 2)
		val, _ := hm.Get("a")
		convey.So(val, should.Equal, 10)
	})
}

func Test_Hashmap_AddOrUpdateMap_Empty(t *testing.T) {
	convey.Convey("Hashmap.AddOrUpdateMap with empty map is no-op", t, func() {
		hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
		hm.AddOrUpdateMap(map[string]int{})
		convey.So(hm.Length(), should.Equal, 1)
	})
}

func Test_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	convey.Convey("Hashmap.AddOrUpdateHashmap merges another hashmap", t, func() {
		hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
		other := coregeneric.HashmapFrom(map[string]int{"b": 2})
		hm.AddOrUpdateHashmap(other)
		convey.So(hm.Length(), should.Equal, 2)
	})
}

func Test_Hashmap_AddOrUpdateHashmap_Nil(t *testing.T) {
	convey.Convey("Hashmap.AddOrUpdateHashmap with nil is no-op", t, func() {
		hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
		hm.AddOrUpdateHashmap(nil)
		convey.So(hm.Length(), should.Equal, 1)
	})
}

// =============================================================================
// Hashmap — ConcatNew
// =============================================================================

func Test_Hashmap_ConcatNew(t *testing.T) {
	convey.Convey("Hashmap.ConcatNew creates merged copy", t, func() {
		hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
		hm2 := coregeneric.HashmapFrom(map[string]int{"b": 2})
		result := hm1.ConcatNew(hm2)
		convey.So(result.Length(), should.Equal, 2)
		convey.So(hm1.Length(), should.Equal, 1) // original unchanged
	})
}

func Test_Hashmap_ConcatNew_WithNil(t *testing.T) {
	convey.Convey("Hashmap.ConcatNew handles nil others", t, func() {
		hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
		result := hm.ConcatNew(nil)
		convey.So(result.Length(), should.Equal, 1)
	})
}

// =============================================================================
// Hashmap — Clone
// =============================================================================

func Test_Hashmap_CloneMethod(t *testing.T) {
	convey.Convey("Hashmap.Clone creates independent copy", t, func() {
		hm := coregeneric.HashmapFrom(map[string]int{"k": 1})
		cloned := hm.Clone()
		cloned.Set("k", 99)

		origVal, _ := hm.Get("k")
		convey.So(origVal, should.Equal, 1)
	})
}

// =============================================================================
// Hashmap — IsEquals
// =============================================================================

func Test_Hashmap_IsEquals_Same(t *testing.T) {
	convey.Convey("Hashmap.IsEquals true for same content", t, func() {
		hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
		hm2 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
		convey.So(hm1.IsEquals(hm2), should.BeTrue)
	})
}

func Test_Hashmap_IsEquals_Different(t *testing.T) {
	convey.Convey("Hashmap.IsEquals false for different keys", t, func() {
		hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
		hm2 := coregeneric.HashmapFrom(map[string]int{"b": 1})
		convey.So(hm1.IsEquals(hm2), should.BeFalse)
	})
}

func Test_Hashmap_IsEquals_DifferentLength(t *testing.T) {
	convey.Convey("Hashmap.IsEquals false for different lengths", t, func() {
		hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
		hm2 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
		convey.So(hm1.IsEquals(hm2), should.BeFalse)
	})
}

func Test_Hashmap_IsEquals_BothNil(t *testing.T) {
	convey.Convey("Hashmap.IsEquals true for both nil", t, func() {
		var hm1, hm2 *coregeneric.Hashmap[string, int]
		convey.So(hm1.IsEquals(hm2), should.BeTrue)
	})
}

func Test_Hashmap_IsEquals_OneNil(t *testing.T) {
	convey.Convey("Hashmap.IsEquals false when one is nil", t, func() {
		hm := coregeneric.EmptyHashmap[string, int]()
		convey.So(hm.IsEquals(nil), should.BeFalse)
	})
}

func Test_Hashmap_IsEquals_SamePointer(t *testing.T) {
	convey.Convey("Hashmap.IsEquals true for same pointer", t, func() {
		hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
		convey.So(hm.IsEquals(hm), should.BeTrue)
	})
}

// =============================================================================
// Hashmap — ForEach / ForEachBreak
// =============================================================================

func Test_Hashmap_ForEach(t *testing.T) {
	convey.Convey("Hashmap.ForEach visits all entries", t, func() {
		hm := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
		count := 0
		hm.ForEach(func(key string, val int) { count++ })
		convey.So(count, should.Equal, 2)
	})
}

func Test_Hashmap_ForEachBreak(t *testing.T) {
	convey.Convey("Hashmap.ForEachBreak can stop early", t, func() {
		hm := coregeneric.HashmapFrom(map[int]int{1: 1, 2: 2, 3: 3})
		count := 0
		hm.ForEachBreak(func(key int, val int) bool {
			count++
			return count >= 2
		})
		convey.So(count, should.Equal, 2)
	})
}

// =============================================================================
// Hashmap — String
// =============================================================================

func Test_Hashmap_String(t *testing.T) {
	convey.Convey("Hashmap.String returns string representation", t, func() {
		hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
		convey.So(hm.String(), should.NotBeEmpty)
	})
}

// =============================================================================
// Hashmap — Nil receiver guards
// =============================================================================

func Test_Hashmap_IsEmpty_NilReceiver(t *testing.T) {
	convey.Convey("Hashmap.IsEmpty true on nil receiver", t, func() {
		var hm *coregeneric.Hashmap[string, int]
		convey.So(hm.IsEmpty(), should.BeTrue)
	})
}

func Test_Hashmap_Length_NilReceiver(t *testing.T) {
	convey.Convey("Hashmap.Length returns 0 on nil receiver", t, func() {
		var hm *coregeneric.Hashmap[string, int]
		convey.So(hm.Length(), should.Equal, 0)
	})
}

func Test_Hashmap_HasItems_NilReceiver(t *testing.T) {
	convey.Convey("Hashmap.HasItems false on nil receiver", t, func() {
		var hm *coregeneric.Hashmap[string, int]
		convey.So(hm.HasItems(), should.BeFalse)
	})
}
