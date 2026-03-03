package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
)

// ==========================================================================
// Test: Hashset — Add / AddBool edge cases
// ==========================================================================

func Test_Hashset_AddDuplicate(t *testing.T) {
	tc := hashsetAddDuplicateTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.Add(1).Add(2).Add(3).Add(1).Add(2)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", hs.Length()))
}

func Test_Hashset_AddBool(t *testing.T) {
	tc := hashsetAddBoolTestCase
	hs := coregeneric.EmptyHashset[string]()
	first := hs.AddBool("a")
	second := hs.AddBool("a")
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", first),
		fmt.Sprintf("%v", second),
		fmt.Sprintf("%d", hs.Length()),
	)
}

func Test_Hashset_AddsVariadic(t *testing.T) {
	tc := hashsetAddsVariadicTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.Adds(10, 20, 30, 10)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", hs.Length()))
}

func Test_Hashset_AddSlice(t *testing.T) {
	tc := hashsetAddSliceTestCase
	hs := coregeneric.EmptyHashset[string]()
	hs.AddSlice([]string{"x", "y", "z"})
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", hs.Length()),
		fmt.Sprintf("%v", hs.Has("x")),
		fmt.Sprintf("%v", hs.Has("y")),
		fmt.Sprintf("%v", hs.Has("z")),
	)
}

// ==========================================================================
// Test: Hashset — AddIf / AddIfMany
// ==========================================================================

func Test_Hashset_AddIfTrue(t *testing.T) {
	tc := hashsetAddIfTrueTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.AddIf(true, 42)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", hs.Length()),
		fmt.Sprintf("%v", hs.Has(42)),
	)
}

func Test_Hashset_AddIfFalse(t *testing.T) {
	tc := hashsetAddIfFalseTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.AddIf(false, 42)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", hs.Length()))
}

func Test_Hashset_AddIfManyTrue(t *testing.T) {
	tc := hashsetAddIfManyTrueTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.AddIfMany(true, 1, 2, 3)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", hs.Length()))
}

func Test_Hashset_AddIfManyFalse(t *testing.T) {
	tc := hashsetAddIfManyFalseTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.AddIfMany(false, 1, 2, 3)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", hs.Length()))
}

// ==========================================================================
// Test: Hashset — AddHashsetItems / AddItemsMap
// ==========================================================================

func Test_Hashset_MergeOtherSet(t *testing.T) {
	tc := hashsetMergeOtherSetTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2})
	other := coregeneric.HashsetFrom([]int{3, 4})
	hs.AddHashsetItems(other)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", hs.Length()),
		fmt.Sprintf("%v", hs.Has(3)),
		fmt.Sprintf("%v", hs.Has(4)),
	)
}

func Test_Hashset_MergeNilOther(t *testing.T) {
	tc := hashsetMergeNilOtherTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2})
	hs.AddHashsetItems(nil)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", hs.Length()))
}

func Test_Hashset_MergeEmptyOther(t *testing.T) {
	tc := hashsetMergeEmptyOtherTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2})
	hs.AddHashsetItems(coregeneric.EmptyHashset[int]())
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", hs.Length()))
}

func Test_Hashset_AddItemsMap(t *testing.T) {
	tc := hashsetAddItemsMapTestCase
	hs := coregeneric.EmptyHashset[string]()
	hs.AddItemsMap(map[string]bool{
		"yes":  true,
		"also": true,
		"nope": false,
	})
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", hs.Length()),
		fmt.Sprintf("%v", hs.Has("yes")),
		fmt.Sprintf("%v", hs.Has("nope")),
	)
}

// ==========================================================================
// Test: Hashset — Remove edge cases
// ==========================================================================

func Test_Hashset_RemoveExisting(t *testing.T) {
	tc := hashsetRemoveExistingTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	existed := hs.Remove(2)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", existed),
		fmt.Sprintf("%d", hs.Length()),
		fmt.Sprintf("%v", hs.Has(2)),
	)
}

func Test_Hashset_RemoveNonExisting(t *testing.T) {
	tc := hashsetRemoveNonExistingTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	existed := hs.Remove(99)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", existed),
		fmt.Sprintf("%d", hs.Length()),
	)
}

// ==========================================================================
// Test: Hashset — Has / Contains
// ==========================================================================

func Test_Hashset_Has(t *testing.T) {
	tc := hashsetHasTestCase
	hs := coregeneric.HashsetFrom([]string{"a", "b", "c"})
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", hs.Has("a")),
		fmt.Sprintf("%v", hs.Has("z")),
	)
}

func Test_Hashset_ContainsAlias(t *testing.T) {
	tc := hashsetContainsAliasTestCase
	hs := coregeneric.HashsetFrom([]string{"a", "b", "c"})
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", hs.Contains("b")),
		fmt.Sprintf("%v", hs.Contains("z")),
	)
}

// ==========================================================================
// Test: Hashset — HasAll / HasAny
// ==========================================================================

func Test_Hashset_HasAllTrue(t *testing.T) {
	tc := hashsetHasAllTrueTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", hs.HasAll(1, 3, 5)))
}

func Test_Hashset_HasAllFalse(t *testing.T) {
	tc := hashsetHasAllFalseTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", hs.HasAll(1, 99)))
}

func Test_Hashset_HasAnyTrue(t *testing.T) {
	tc := hashsetHasAnyTrueTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", hs.HasAny(99, 3)))
}

func Test_Hashset_HasAnyFalse(t *testing.T) {
	tc := hashsetHasAnyFalseTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", hs.HasAny(99, 100)))
}

func Test_Hashset_HasAllEmptyArgs(t *testing.T) {
	tc := hashsetHasAllEmptyArgsTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", hs.HasAll()))
}

func Test_Hashset_HasAnyEmptyArgs(t *testing.T) {
	tc := hashsetHasAnyEmptyArgsTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", hs.HasAny()))
}

// ==========================================================================
// Test: Hashset — IsEquals
// ==========================================================================

func Test_Hashset_IsEquals_SameItems(t *testing.T) {
	tc := hashsetIsEqualsSameItemsTestCase
	a := coregeneric.HashsetFrom([]int{1, 2, 3})
	b := coregeneric.HashsetFrom([]int{3, 2, 1})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEquals(b)))
}

func Test_Hashset_IsEquals_DifferentItems(t *testing.T) {
	tc := hashsetIsEqualsDifferentItemsTestCase
	a := coregeneric.HashsetFrom([]int{1, 2, 3})
	b := coregeneric.HashsetFrom([]int{1, 2, 4})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEquals(b)))
}

func Test_Hashset_IsEquals_DifferentLength(t *testing.T) {
	tc := hashsetIsEqualsDifferentLengthTestCase
	a := coregeneric.HashsetFrom([]int{1, 2})
	b := coregeneric.HashsetFrom([]int{1, 2, 3})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEquals(b)))
}

func Test_Hashset_IsEquals_BothNil(t *testing.T) {
	tc := hashsetIsEqualsBothNilTestCase
	var a *coregeneric.Hashset[int]
	var b *coregeneric.Hashset[int]
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEquals(b)))
}

func Test_Hashset_IsEquals_NilVsNonNil(t *testing.T) {
	tc := hashsetIsEqualsNilVsNonNilTestCase
	var a *coregeneric.Hashset[int]
	b := coregeneric.EmptyHashset[int]()
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEquals(b)))
}

func Test_Hashset_IsEquals_SamePointer(t *testing.T) {
	tc := hashsetIsEqualsSamePointerTestCase
	a := coregeneric.HashsetFrom([]int{1, 2})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEquals(a)))
}

func Test_Hashset_IsEquals_BothEmpty(t *testing.T) {
	tc := hashsetIsEqualsBothEmptyTestCase
	a := coregeneric.EmptyHashset[int]()
	b := coregeneric.EmptyHashset[int]()
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEquals(b)))
}

// ==========================================================================
// Test: Hashset — Resize
// ==========================================================================

func Test_Hashset_ResizeLarger(t *testing.T) {
	tc := hashsetResizeLargerTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	hs.Resize(100)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", hs.Length()),
		fmt.Sprintf("%v", hs.Has(1)),
		fmt.Sprintf("%v", hs.Has(2)),
		fmt.Sprintf("%v", hs.Has(3)),
	)
}

func Test_Hashset_ResizeSmaller(t *testing.T) {
	tc := hashsetResizeSmallerTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	hs.Resize(1)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", hs.Length()))
}

// ==========================================================================
// Test: Hashset — List / ListPtr / Map / Collection / String
// ==========================================================================

func Test_Hashset_OutputList(t *testing.T) {
	tc := hashsetOutputListTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", len(hs.List())))
}

func Test_Hashset_OutputListEmpty(t *testing.T) {
	tc := hashsetOutputListEmptyTestCase
	hs := coregeneric.EmptyHashset[int]()
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", len(hs.List())))
}

func Test_Hashset_OutputListPtr(t *testing.T) {
	tc := hashsetOutputListPtrTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", hs.ListPtr() != nil))
}

func Test_Hashset_OutputMap(t *testing.T) {
	tc := hashsetOutputMapTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", len(hs.Map())))
}

func Test_Hashset_OutputCollection(t *testing.T) {
	tc := hashsetOutputCollectionTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	col := hs.Collection()
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.Length()))
}

// ==========================================================================
// Test: Hashset — Lock variants
// ==========================================================================

func Test_Hashset_LockAddContains(t *testing.T) {
	tc := hashsetLockAddContainsTestCase
	hs := coregeneric.EmptyHashset[string]()
	hs.AddLock("a")
	hs.AddLock("b")
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", hs.Length()),
		fmt.Sprintf("%v", hs.ContainsLock("a")),
		fmt.Sprintf("%v", hs.ContainsLock("b")),
		fmt.Sprintf("%v", hs.ContainsLock("z")),
	)
}

func Test_Hashset_LockAddSlice(t *testing.T) {
	tc := hashsetLockAddSliceTestCase
	hs := coregeneric.EmptyHashset[int]()
	hs.AddSliceLock([]int{10, 20, 30})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", hs.Length()))
}

func Test_Hashset_LockRemove(t *testing.T) {
	tc := hashsetLockRemoveTestCase
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	existed := hs.RemoveLock(2)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", existed),
		fmt.Sprintf("%d", hs.Length()),
		fmt.Sprintf("%v", hs.Has(2)),
	)
}

func Test_Hashset_LockIsEmptyLength(t *testing.T) {
	tc := hashsetLockIsEmptyLengthTestCase
	hs := coregeneric.EmptyHashset[int]()
	emptyBefore := fmt.Sprintf("%v", hs.IsEmptyLock())
	lenBefore := fmt.Sprintf("%d", hs.LengthLock())
	hs.Adds(1, 2)
	emptyAfter := fmt.Sprintf("%v", hs.IsEmptyLock())
	lenAfter := fmt.Sprintf("%d", hs.LengthLock())
	tc.ShouldBeEqual(t, 0, emptyBefore, lenBefore, emptyAfter, lenAfter)
}

// ==========================================================================
// Test: Hashset — Constructors
// ==========================================================================

func Test_Hashset_ConstructorEmpty(t *testing.T) {
	tc := hashsetConstructorEmptyTestCase
	hs := coregeneric.EmptyHashset[int]()
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", hs.Length()),
		fmt.Sprintf("%v", hs.IsEmpty()),
	)
}

func Test_Hashset_ConstructorNewCap(t *testing.T) {
	tc := hashsetConstructorNewCapTestCase
	hs := coregeneric.NewHashset[string](10)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", hs.Length()),
		fmt.Sprintf("%v", hs.IsEmpty()),
	)
}

func Test_Hashset_ConstructorFrom(t *testing.T) {
	tc := hashsetConstructorFromTestCase
	hs := coregeneric.HashsetFrom([]string{"a", "b", "c"})
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", hs.Length()),
		fmt.Sprintf("%v", hs.Has("a")),
		fmt.Sprintf("%v", hs.Has("b")),
		fmt.Sprintf("%v", hs.Has("c")),
	)
}

func Test_Hashset_ConstructorFromMap(t *testing.T) {
	tc := hashsetConstructorFromMapTestCase
	hs := coregeneric.HashsetFromMap(map[int]bool{10: true, 20: true})
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", hs.Length()),
		fmt.Sprintf("%v", hs.Has(10)),
		fmt.Sprintf("%v", hs.Has(20)),
	)
}

func Test_Hashset_ConstructorHasItems(t *testing.T) {
	tc := hashsetConstructorHasItemsTestCase
	pop := coregeneric.HashsetFrom([]int{1})
	empty := coregeneric.EmptyHashset[int]()
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", pop.HasItems()),
		fmt.Sprintf("%v", empty.HasItems()),
	)
}
