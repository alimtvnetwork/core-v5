package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================================================
// Test: EmptyHashmap
// ==========================================================================

func Test_Hashmap_Empty(t *testing.T) {
	tc := hashmapEmptyTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	actLines := []string{
		fmt.Sprintf("%v", hm.IsEmpty()),
		fmt.Sprintf("%v", hm.Length()),
		fmt.Sprintf("%v", hm.HasItems()),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: NewHashmap
// ==========================================================================

func Test_Hashmap_New(t *testing.T) {
	tc := hashmapNewTestCase
	hm := coregeneric.NewHashmap[string, int](10)

	actLines := []string{fmt.Sprintf("%v", hm.IsEmpty())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: HashmapFrom
// ==========================================================================

func Test_Hashmap_From(t *testing.T) {
	tc := hashmapFromTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})

	actLines := []string{
		fmt.Sprintf("%v", hm.Length()),
		fmt.Sprintf("%v", hm.Has("a")),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
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

	actLines := []string{
		fmt.Sprintf("%v", origVal),
		fmt.Sprintf("%v", clonedVal),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: Set
// ==========================================================================

func Test_Hashmap_SetNew(t *testing.T) {
	tc := hashmapSetNewTestCase
	hm := coregeneric.EmptyHashmap[string, int]()
	isNew := hm.Set("key", 42)

	actLines := []string{
		fmt.Sprintf("%v", isNew),
		fmt.Sprintf("%v", hm.Length()),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_SetExisting(t *testing.T) {
	tc := hashmapSetExistingTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"key": 1})
	isNew := hm.Set("key", 2)
	val, _ := hm.Get("key")

	actLines := []string{
		fmt.Sprintf("%v", isNew),
		fmt.Sprintf("%v", val),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: Get
// ==========================================================================

func Test_Hashmap_GetFound(t *testing.T) {
	tc := hashmapGetFoundTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"k": 42})
	val, found := hm.Get("k")

	actLines := []string{
		fmt.Sprintf("%v", found),
		fmt.Sprintf("%v", val),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_GetNotFound(t *testing.T) {
	tc := hashmapGetNotFoundTestCase
	hm := coregeneric.EmptyHashmap[string, int]()
	val, found := hm.Get("missing")

	actLines := []string{
		fmt.Sprintf("%v", found),
		fmt.Sprintf("%v", val),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: GetOrDefault
// ==========================================================================

func Test_Hashmap_GetOrDefaultMissing(t *testing.T) {
	tc := hashmapGetOrDefaultMissingTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	actLines := []string{fmt.Sprintf("%v", hm.GetOrDefault("x", 99))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_GetOrDefaultFound(t *testing.T) {
	tc := hashmapGetOrDefaultFoundTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"x": 5})

	actLines := []string{fmt.Sprintf("%v", hm.GetOrDefault("x", 99))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: Has / Contains / IsKeyMissing
// ==========================================================================

func Test_Hashmap_Has(t *testing.T) {
	tc := hashmapHasTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	actLines := []string{
		fmt.Sprintf("%v", hm.Has("a")),
		fmt.Sprintf("%v", hm.Contains("a")),
		fmt.Sprintf("%v", hm.IsKeyMissing("a")),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_IsKeyMissing(t *testing.T) {
	tc := hashmapIsKeyMissingTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	actLines := []string{fmt.Sprintf("%v", hm.IsKeyMissing("x"))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: Remove
// ==========================================================================

func Test_Hashmap_RemoveExisting(t *testing.T) {
	tc := hashmapRemoveExistingTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"k": 1})
	existed := hm.Remove("k")

	actLines := []string{
		fmt.Sprintf("%v", existed),
		fmt.Sprintf("%v", hm.IsEmpty()),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_RemoveMissing(t *testing.T) {
	tc := hashmapRemoveMissingTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	actLines := []string{fmt.Sprintf("%v", hm.Remove("x"))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: Keys
// ==========================================================================

func Test_Hashmap_Keys_NonEmpty(t *testing.T) {
	tc := hashmapKeysNonEmptyTestCase
	hm := coregeneric.HashmapFrom(map[int]string{1: "a", 2: "b"})

	actLines := []string{fmt.Sprintf("%v", len(hm.Keys()))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_Keys_Empty(t *testing.T) {
	tc := hashmapKeysEmptyTestCase
	hm := coregeneric.EmptyHashmap[int, string]()

	actLines := []string{fmt.Sprintf("%v", len(hm.Keys()))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: Values
// ==========================================================================

func Test_Hashmap_Values_NonEmpty(t *testing.T) {
	tc := hashmapValuesNonEmptyTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	vals := hm.Values()

	actLines := []string{
		fmt.Sprintf("%v", len(vals)),
		fmt.Sprintf("%v", vals[0]),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_Values_Empty(t *testing.T) {
	tc := hashmapValuesEmptyTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	actLines := []string{fmt.Sprintf("%v", len(hm.Values()))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: AddOrUpdateMap
// ==========================================================================

func Test_Hashmap_AddOrUpdateMap_Merges(t *testing.T) {
	tc := hashmapAddOrUpdateMapMergesTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm.AddOrUpdateMap(map[string]int{"b": 2, "a": 10})
	val, _ := hm.Get("a")

	actLines := []string{
		fmt.Sprintf("%v", hm.Length()),
		fmt.Sprintf("%v", val),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_AddOrUpdateMap_EmptyNoop(t *testing.T) {
	tc := hashmapAddOrUpdateMapEmptyNoopTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm.AddOrUpdateMap(map[string]int{})

	actLines := []string{fmt.Sprintf("%v", hm.Length())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: AddOrUpdateHashmap
// ==========================================================================

func Test_Hashmap_AddOrUpdateHashmap_Merges(t *testing.T) {
	tc := hashmapAddOrUpdateHashmapMergesTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm.AddOrUpdateHashmap(coregeneric.HashmapFrom(map[string]int{"b": 2}))

	actLines := []string{fmt.Sprintf("%v", hm.Length())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_AddOrUpdateHashmap_NilNoop(t *testing.T) {
	tc := hashmapAddOrUpdateHashmapNilNoopTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm.AddOrUpdateHashmap(nil)

	actLines := []string{fmt.Sprintf("%v", hm.Length())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: ConcatNew
// ==========================================================================

func Test_Hashmap_ConcatNew_Merged(t *testing.T) {
	tc := hashmapConcatNewMergedTestCase
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.HashmapFrom(map[string]int{"b": 2})
	result := hm1.ConcatNew(hm2)

	actLines := []string{
		fmt.Sprintf("%v", result.Length()),
		fmt.Sprintf("%v", hm1.Length()),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_ConcatNew_Nil(t *testing.T) {
	tc := hashmapConcatNewNilTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	result := hm.ConcatNew(nil)

	actLines := []string{fmt.Sprintf("%v", result.Length())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
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

	actLines := []string{fmt.Sprintf("%v", origVal)}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: IsEquals
// ==========================================================================

func Test_Hashmap_IsEquals_SameContent(t *testing.T) {
	tc := hashmapIsEqualsSameContentTestCase
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
	hm2 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})

	actLines := []string{fmt.Sprintf("%v", hm1.IsEquals(hm2))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_IsEquals_DifferentKeys(t *testing.T) {
	tc := hashmapIsEqualsDifferentKeysTestCase
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.HashmapFrom(map[string]int{"b": 1})

	actLines := []string{fmt.Sprintf("%v", hm1.IsEquals(hm2))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_IsEquals_DifferentLength(t *testing.T) {
	tc := hashmapIsEqualsDifferentLengthTestCase
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})

	actLines := []string{fmt.Sprintf("%v", hm1.IsEquals(hm2))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_IsEquals_BothNil(t *testing.T) {
	tc := hashmapIsEqualsBothNilTestCase
	var hm1, hm2 *coregeneric.Hashmap[string, int]

	actLines := []string{fmt.Sprintf("%v", hm1.IsEquals(hm2))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_IsEquals_OneNil(t *testing.T) {
	tc := hashmapIsEqualsOneNilTestCase
	hm := coregeneric.EmptyHashmap[string, int]()

	actLines := []string{fmt.Sprintf("%v", hm.IsEquals(nil))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_IsEquals_SamePointer(t *testing.T) {
	tc := hashmapIsEqualsSamePointerTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	actLines := []string{fmt.Sprintf("%v", hm.IsEquals(hm))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: ForEach
// ==========================================================================

func Test_Hashmap_ForEach(t *testing.T) {
	tc := hashmapForEachTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
	count := 0
	hm.ForEach(func(_ string, _ int) { count++ })

	actLines := []string{fmt.Sprintf("%v", count)}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_ForEachBreak(t *testing.T) {
	tc := hashmapForEachBreakTestCase
	hm := coregeneric.HashmapFrom(map[int]int{1: 1, 2: 2, 3: 3})
	count := 0
	hm.ForEachBreak(func(_ int, _ int) bool { count++; return count >= 2 })

	actLines := []string{fmt.Sprintf("%v", count)}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: String
// ==========================================================================

func Test_Hashmap_String(t *testing.T) {
	tc := hashmapStringTestCase
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	actLines := []string{fmt.Sprintf("%v", hm.String() != "")}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: Nil receiver
// ==========================================================================

func Test_Hashmap_NilReceiver_IsEmpty(t *testing.T) {
	tc := hashmapNilReceiverIsEmptyTestCase
	var hm *coregeneric.Hashmap[string, int]

	actLines := []string{fmt.Sprintf("%v", hm.IsEmpty())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_NilReceiver_Length(t *testing.T) {
	tc := hashmapNilReceiverLengthTestCase
	var hm *coregeneric.Hashmap[string, int]

	actLines := []string{fmt.Sprintf("%v", hm.Length())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_NilReceiver_HasItems(t *testing.T) {
	tc := hashmapNilReceiverHasItemsTestCase
	var hm *coregeneric.Hashmap[string, int]

	actLines := []string{fmt.Sprintf("%v", hm.HasItems())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}
