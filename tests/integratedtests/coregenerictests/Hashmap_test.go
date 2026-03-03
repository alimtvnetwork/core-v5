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
	tc := hashmapEmptyTestCases[0]
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
	tc := hashmapNewTestCases[0]
	hm := coregeneric.NewHashmap[string, int](10)

	actLines := []string{fmt.Sprintf("%v", hm.IsEmpty())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: HashmapFrom
// ==========================================================================

func Test_Hashmap_From(t *testing.T) {
	tc := hashmapFromTestCases[0]
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
	tc := hashmapCloneFuncTestCases[0]
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
	tc := hashmapSetNewTestCases[0]
	hm := coregeneric.EmptyHashmap[string, int]()
	isNew := hm.Set("key", 42)

	actLines := []string{
		fmt.Sprintf("%v", isNew),
		fmt.Sprintf("%v", hm.Length()),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_SetExisting(t *testing.T) {
	tc := hashmapSetExistingTestCases[0]
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
	tc := hashmapGetFoundTestCases[0]
	hm := coregeneric.HashmapFrom(map[string]int{"k": 42})
	val, found := hm.Get("k")

	actLines := []string{
		fmt.Sprintf("%v", found),
		fmt.Sprintf("%v", val),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_GetNotFound(t *testing.T) {
	tc := hashmapGetNotFoundTestCases[0]
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
	tc := hashmapGetOrDefaultMissingTestCases[0]
	hm := coregeneric.EmptyHashmap[string, int]()

	actLines := []string{fmt.Sprintf("%v", hm.GetOrDefault("x", 99))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_GetOrDefaultFound(t *testing.T) {
	tc := hashmapGetOrDefaultFoundTestCases[0]
	hm := coregeneric.HashmapFrom(map[string]int{"x": 5})

	actLines := []string{fmt.Sprintf("%v", hm.GetOrDefault("x", 99))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: Has / Contains / IsKeyMissing
// ==========================================================================

func Test_Hashmap_Has(t *testing.T) {
	tc := hashmapHasTestCases[0]
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	actLines := []string{
		fmt.Sprintf("%v", hm.Has("a")),
		fmt.Sprintf("%v", hm.Contains("a")),
		fmt.Sprintf("%v", hm.IsKeyMissing("a")),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_IsKeyMissing(t *testing.T) {
	tc := hashmapIsKeyMissingTestCases[0]
	hm := coregeneric.EmptyHashmap[string, int]()

	actLines := []string{fmt.Sprintf("%v", hm.IsKeyMissing("x"))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: Remove
// ==========================================================================

func Test_Hashmap_RemoveExisting(t *testing.T) {
	tc := hashmapRemoveExistingTestCases[0]
	hm := coregeneric.HashmapFrom(map[string]int{"k": 1})
	existed := hm.Remove("k")

	actLines := []string{
		fmt.Sprintf("%v", existed),
		fmt.Sprintf("%v", hm.IsEmpty()),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_RemoveMissing(t *testing.T) {
	tc := hashmapRemoveMissingTestCases[0]
	hm := coregeneric.EmptyHashmap[string, int]()

	actLines := []string{fmt.Sprintf("%v", hm.Remove("x"))}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: Keys
// ==========================================================================

func Test_Hashmap_Keys(t *testing.T) {
	for caseIndex, tc := range hashmapKeysTestCases {
		var actLines []string

		if caseIndex == 0 {
			hm := coregeneric.HashmapFrom(map[int]string{1: "a", 2: "b"})
			actLines = []string{fmt.Sprintf("%v", len(hm.Keys()))}
		} else {
			hm := coregeneric.EmptyHashmap[int, string]()
			actLines = []string{fmt.Sprintf("%v", len(hm.Keys()))}
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: Values
// ==========================================================================

func Test_Hashmap_Values(t *testing.T) {
	for caseIndex, tc := range hashmapValuesTestCases {
		var actLines []string

		if caseIndex == 0 {
			hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
			vals := hm.Values()
			actLines = []string{fmt.Sprintf("%v", len(vals)), fmt.Sprintf("%v", vals[0])}
		} else {
			hm := coregeneric.EmptyHashmap[string, int]()
			actLines = []string{fmt.Sprintf("%v", len(hm.Values()))}
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: AddOrUpdateMap
// ==========================================================================

func Test_Hashmap_AddOrUpdateMap(t *testing.T) {
	for caseIndex, tc := range hashmapAddOrUpdateMapTestCases {
		hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
		var actLines []string

		if caseIndex == 0 {
			hm.AddOrUpdateMap(map[string]int{"b": 2, "a": 10})
			val, _ := hm.Get("a")
			actLines = []string{fmt.Sprintf("%v", hm.Length()), fmt.Sprintf("%v", val)}
		} else {
			hm.AddOrUpdateMap(map[string]int{})
			actLines = []string{fmt.Sprintf("%v", hm.Length())}
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: AddOrUpdateHashmap
// ==========================================================================

func Test_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	for caseIndex, tc := range hashmapAddOrUpdateHashmapTestCases {
		hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
		var actLines []string

		if caseIndex == 0 {
			hm.AddOrUpdateHashmap(coregeneric.HashmapFrom(map[string]int{"b": 2}))
			actLines = []string{fmt.Sprintf("%v", hm.Length())}
		} else {
			hm.AddOrUpdateHashmap(nil)
			actLines = []string{fmt.Sprintf("%v", hm.Length())}
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: ConcatNew
// ==========================================================================

func Test_Hashmap_ConcatNew(t *testing.T) {
	for caseIndex, tc := range hashmapConcatNewTestCases {
		var actLines []string

		if caseIndex == 0 {
			hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
			hm2 := coregeneric.HashmapFrom(map[string]int{"b": 2})
			result := hm1.ConcatNew(hm2)
			actLines = []string{fmt.Sprintf("%v", result.Length()), fmt.Sprintf("%v", hm1.Length())}
		} else {
			hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
			result := hm.ConcatNew(nil)
			actLines = []string{fmt.Sprintf("%v", result.Length())}
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: Clone method
// ==========================================================================

func Test_Hashmap_CloneMethod(t *testing.T) {
	tc := hashmapCloneMethodTestCases[0]
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

func Test_Hashmap_IsEquals(t *testing.T) {
	for caseIndex, tc := range hashmapIsEqualsTestCases {
		var actLines []string

		switch caseIndex {
		case 0:
			hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
			hm2 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
			actLines = []string{fmt.Sprintf("%v", hm1.IsEquals(hm2))}
		case 1:
			hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
			hm2 := coregeneric.HashmapFrom(map[string]int{"b": 1})
			actLines = []string{fmt.Sprintf("%v", hm1.IsEquals(hm2))}
		case 2:
			hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
			hm2 := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
			actLines = []string{fmt.Sprintf("%v", hm1.IsEquals(hm2))}
		case 3:
			var hm1, hm2 *coregeneric.Hashmap[string, int]
			actLines = []string{fmt.Sprintf("%v", hm1.IsEquals(hm2))}
		case 4:
			hm := coregeneric.EmptyHashmap[string, int]()
			actLines = []string{fmt.Sprintf("%v", hm.IsEquals(nil))}
		case 5:
			hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
			actLines = []string{fmt.Sprintf("%v", hm.IsEquals(hm))}
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: ForEach
// ==========================================================================

func Test_Hashmap_ForEach(t *testing.T) {
	tc := hashmapForEachTestCases[0]
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
	count := 0
	hm.ForEach(func(_ string, _ int) { count++ })

	actLines := []string{fmt.Sprintf("%v", count)}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Hashmap_ForEachBreak(t *testing.T) {
	tc := hashmapForEachBreakTestCases[0]
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
	tc := hashmapStringTestCases[0]
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	actLines := []string{fmt.Sprintf("%v", hm.String() != "")}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: Nil receiver
// ==========================================================================

func Test_Hashmap_NilReceiver(t *testing.T) {
	for caseIndex, tc := range hashmapNilReceiverTestCases {
		var hm *coregeneric.Hashmap[string, int]
		var actLines []string

		switch caseIndex {
		case 0:
			actLines = []string{fmt.Sprintf("%v", hm.IsEmpty())}
		case 1:
			actLines = []string{fmt.Sprintf("%v", hm.Length())}
		case 2:
			actLines = []string{fmt.Sprintf("%v", hm.HasItems())}
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, tc.ExpectedInput)
	}
}
