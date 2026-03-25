package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Test: EmptyHashmap
// ==========================================================================

func Test_Hashmap_Empty(t *testing.T) {
	tc := hashmapEmptyTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	actual := args.Map{
		"isEmpty":  hm.IsEmpty(),
		"length":   hm.Length(),
		"hasItems": hm.HasItems(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: NewHashmap
// ==========================================================================

func Test_Hashmap_New(t *testing.T) {
	tc := hashmapNewTestCase
	hm := coregeneric.NewHashmap[string, int](10)

	actual := args.Map{"isEmpty": hm.IsEmpty()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: HashmapFrom
// ==========================================================================

func Test_Hashmap_From(t *testing.T) {
	tc := hashmapFromTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})

	actual := args.Map{
		"length": hm.Length(),
		"hasKey": hm.Has("a"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: HashmapClone (function)
// ==========================================================================

func Test_Hashmap_CloneFunc(t *testing.T) {
	tc := hashmapCloneFuncTestCase
	orig := coregeneric.HashmapFrom(map[string]int{"k": 1})
	cloned := coregeneric.HashmapClone(orig.Map())
	cloned.Set("k", 99)

	origVal, _ := orig.Get("k")
	clonedVal, _ := cloned.Get("k")

	actual := args.Map{
		"origValue":   origVal,
		"clonedValue": clonedVal,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Set
// ==========================================================================

func Test_Hashmap_SetNew(t *testing.T) {
	tc := hashmapSetNewTestCase
	hm := coregeneric.EmptyHashmap[string, int]()
	isNew := hm.Set("key", 42)

	actual := args.Map{
		"isNew":  isNew,
		"length": hm.Length(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_SetExisting(t *testing.T) {
	tc := hashmapSetExistingTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"key": 1})
	isNew := hm.Set("key", 2)
	val, _ := hm.Get("key")

	actual := args.Map{
		"isNew":        isNew,
		"updatedValue": val,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Get
// ==========================================================================

func Test_Hashmap_GetFound(t *testing.T) {
	tc := hashmapGetFoundTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"k": 42})
	val, found := hm.Get("k")

	actual := args.Map{
		"found": found,
		"value": val,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_GetNotFound(t *testing.T) {
	tc := hashmapGetNotFoundTestCase
	hm := coregeneric.EmptyHashmap[string, int]()
	val, found := hm.Get("missing")

	actual := args.Map{
		"found": found,
		"value": val,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: GetOrDefault
// ==========================================================================

func Test_Hashmap_GetOrDefaultMissing(t *testing.T) {
	tc := hashmapGetOrDefaultMissingTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	actual := args.Map{"value": hm.GetOrDefault("x", 99)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_GetOrDefaultFound(t *testing.T) {
	tc := hashmapGetOrDefaultFoundTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"x": 5})

	actual := args.Map{"value": hm.GetOrDefault("x", 99)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Has / Contains / IsKeyMissing
// ==========================================================================

func Test_Hashmap_Has(t *testing.T) {
	tc := hashmapHasTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	actual := args.Map{
		"has":          hm.Has("a"),
		"contains":     hm.Contains("a"),
		"isKeyMissing": hm.IsKeyMissing("a"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsKeyMissing(t *testing.T) {
	tc := hashmapIsKeyMissingTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	actual := args.Map{"isKeyMissing": hm.IsKeyMissing("x")}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Remove
// ==========================================================================

func Test_Hashmap_RemoveExisting(t *testing.T) {
	tc := hashmapRemoveExistingTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"k": 1})
	existed := hm.Remove("k")

	actual := args.Map{
		"removed": existed,
		"isGone":  hm.IsEmpty(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_RemoveMissing(t *testing.T) {
	tc := hashmapRemoveMissingTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	actual := args.Map{"removed": hm.Remove("x")}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Keys
// ==========================================================================

func Test_Hashmap_Keys_NonEmpty(t *testing.T) {
	tc := hashmapKeysNonEmptyTestCase
	hm := coregeneric.HashmapFrom(map[int]string{1: "a", 2: "b"})

	actual := args.Map{"keyCount": len(hm.Keys())}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_Keys_Empty(t *testing.T) {
	tc := hashmapKeysEmptyTestCase
	hm := coregeneric.EmptyHashmap[int, string]()

	actual := args.Map{"keyCount": len(hm.Keys())}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Values
// ==========================================================================

func Test_Hashmap_Values_NonEmpty(t *testing.T) {
	tc := hashmapValuesNonEmptyTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	vals := hm.Values()

	actual := args.Map{
		"valueCount":       len(vals),
		"containsExpected": vals[0],
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_Values_Empty(t *testing.T) {
	tc := hashmapValuesEmptyTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	actual := args.Map{"valueCount": len(hm.Values())}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: AddOrUpdateMap
// ==========================================================================

func Test_Hashmap_AddOrUpdateMap_Merges(t *testing.T) {
	tc := hashmapAddOrUpdateMapMergesTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm.AddOrUpdateMap(map[string]int{"b": 2, "a": 10})
	val, _ := hm.Get("a")

	actual := args.Map{
		"length":      hm.Length(),
		"mergedValue": val,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_AddOrUpdateMap_EmptyNoop(t *testing.T) {
	tc := hashmapAddOrUpdateMapEmptyNoopTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm.AddOrUpdateMap(map[string]int{})

	actual := args.Map{"length": hm.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: AddOrUpdateHashmap
// ==========================================================================

func Test_Hashmap_AddOrUpdateHashmap_Merges(t *testing.T) {
	tc := hashmapAddOrUpdateHashmapMergesTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm.AddOrUpdateHashmap(coregeneric.HashmapFrom(map[string]int{"b": 2}))

	actual := args.Map{"length": hm.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_AddOrUpdateHashmap_NilNoop(t *testing.T) {
	tc := hashmapAddOrUpdateHashmapNilNoopTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm.AddOrUpdateHashmap(nil)

	actual := args.Map{"length": hm.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: ConcatNew
// ==========================================================================

func Test_Hashmap_ConcatNew_Merged(t *testing.T) {
	tc := hashmapConcatNewMergedTestCase
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.HashmapFrom(map[string]int{"b": 2})
	result := hm1.ConcatNew(hm2)

	actual := args.Map{
		"mergedLength":   result.Length(),
		"originalLength": hm1.Length(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_ConcatNew_Nil(t *testing.T) {
	tc := hashmapConcatNewNilTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	result := hm.ConcatNew(nil)

	actual := args.Map{"length": result.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Clone method
// ==========================================================================

func Test_Hashmap_CloneMethod(t *testing.T) {
	tc := hashmapCloneMethodTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"k": 1})
	cloned := hm.Clone()
	cloned.Set("k", 99)
	origVal, _ := hm.Get("k")

	actual := args.Map{"origValue": origVal}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: IsEquals
// ==========================================================================

func Test_Hashmap_IsEquals_SameContent(t *testing.T) {
	tc := hashmapIsEqualsSameContentTestCase
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
	hm2 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})

	actual := args.Map{"isEquals": hm1.IsEquals(hm2)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsEquals_DifferentKeys(t *testing.T) {
	tc := hashmapIsEqualsDifferentKeysTestCase
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.HashmapFrom(map[string]int{"b": 1})

	actual := args.Map{"isEquals": hm1.IsEquals(hm2)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsEquals_DifferentLength(t *testing.T) {
	tc := hashmapIsEqualsDifferentLengthTestCase
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})

	actual := args.Map{"isEquals": hm1.IsEquals(hm2)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// Test_Hashmap_IsEquals_BothNil is defined in CollectionBranch_test.go (line 384).
// Removed duplicate declaration here.

func Test_Hashmap_IsEquals_OneNil(t *testing.T) {
	tc := hashmapIsEqualsOneNilTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	actual := args.Map{"isEquals": hm.IsEquals(nil)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsEquals_SamePointer(t *testing.T) {
	tc := hashmapIsEqualsSamePointerTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	actual := args.Map{"isEquals": hm.IsEquals(hm)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: ForEach
// ==========================================================================

func Test_Hashmap_ForEach(t *testing.T) {
	tc := hashmapForEachTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
	count := 0
	hm.ForEach(func(_ string, _ int) { count++ })

	actual := args.Map{"visitCount": count}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_ForEachBreak(t *testing.T) {
	tc := hashmapForEachBreakTestCase
	hm := coregeneric.HashmapFrom(map[int]int{1: 1, 2: 2, 3: 3})
	count := 0
	hm.ForEachBreak(func(_ int, _ int) bool { count++; return count >= 2 })

	actual := args.Map{"visitCount": count}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: String
// ==========================================================================

func Test_Hashmap_String(t *testing.T) {
	tc := hashmapStringTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	actual := args.Map{"isNonEmpty": hm.String() != ""}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// Note: Nil receiver tests migrated to NilReceiver_test.go using CaseNilSafe pattern.
