package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
)

// ==========================================================================
// Test: Hashset — Add / AddBool edge cases
// ==========================================================================

func Test_Hashset_Add_Edge(t *testing.T) {
	// Case 0: duplicate
	{
		tc := hashsetAddEdgeTestCases[0]
		hs := coregeneric.EmptyHashset[int]()
		hs.Add(1).Add(2).Add(3).Add(1).Add(2)
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", hs.Length()))
	}

	// Case 1: AddBool
	{
		tc := hashsetAddEdgeTestCases[1]
		hs := coregeneric.EmptyHashset[string]()
		first := hs.AddBool("a")
		second := hs.AddBool("a")
		tc.ShouldBeEqual(t, 1,
			fmt.Sprintf("%v", first),
			fmt.Sprintf("%v", second),
			fmt.Sprintf("%d", hs.Length()),
		)
	}

	// Case 2: Adds variadic
	{
		tc := hashsetAddEdgeTestCases[2]
		hs := coregeneric.EmptyHashset[int]()
		hs.Adds(10, 20, 30, 10)
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%d", hs.Length()))
	}

	// Case 3: AddSlice
	{
		tc := hashsetAddEdgeTestCases[3]
		hs := coregeneric.EmptyHashset[string]()
		hs.AddSlice([]string{"x", "y", "z"})
		tc.ShouldBeEqual(t, 3,
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has("x")),
			fmt.Sprintf("%v", hs.Has("y")),
			fmt.Sprintf("%v", hs.Has("z")),
		)
	}
}

// ==========================================================================
// Test: Hashset — AddIf / AddIfMany
// ==========================================================================

func Test_Hashset_AddIf(t *testing.T) {
	// Case 0: true
	{
		tc := hashsetAddIfTestCases[0]
		hs := coregeneric.EmptyHashset[int]()
		hs.AddIf(true, 42)
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has(42)),
		)
	}

	// Case 1: false
	{
		tc := hashsetAddIfTestCases[1]
		hs := coregeneric.EmptyHashset[int]()
		hs.AddIf(false, 42)
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", hs.Length()))
	}

	// Case 2: AddIfMany true
	{
		tc := hashsetAddIfTestCases[2]
		hs := coregeneric.EmptyHashset[int]()
		hs.AddIfMany(true, 1, 2, 3)
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%d", hs.Length()))
	}

	// Case 3: AddIfMany false
	{
		tc := hashsetAddIfTestCases[3]
		hs := coregeneric.EmptyHashset[int]()
		hs.AddIfMany(false, 1, 2, 3)
		tc.ShouldBeEqual(t, 3, fmt.Sprintf("%d", hs.Length()))
	}
}

// ==========================================================================
// Test: Hashset — AddHashsetItems / AddItemsMap
// ==========================================================================

func Test_Hashset_Merge(t *testing.T) {
	// Case 0: merge other set
	{
		tc := hashsetMergeTestCases[0]
		hs := coregeneric.HashsetFrom([]int{1, 2})
		other := coregeneric.HashsetFrom([]int{3, 4})
		hs.AddHashsetItems(other)
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has(3)),
			fmt.Sprintf("%v", hs.Has(4)),
		)
	}

	// Case 1: nil other
	{
		tc := hashsetMergeTestCases[1]
		hs := coregeneric.HashsetFrom([]int{1, 2})
		hs.AddHashsetItems(nil)
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", hs.Length()))
	}

	// Case 2: empty other
	{
		tc := hashsetMergeTestCases[2]
		hs := coregeneric.HashsetFrom([]int{1, 2})
		hs.AddHashsetItems(coregeneric.EmptyHashset[int]())
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%d", hs.Length()))
	}

	// Case 3: AddItemsMap only true entries
	{
		tc := hashsetMergeTestCases[3]
		hs := coregeneric.EmptyHashset[string]()
		hs.AddItemsMap(map[string]bool{
			"yes":    true,
			"also":   true,
			"nope":   false,
		})
		tc.ShouldBeEqual(t, 3,
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has("yes")),
			fmt.Sprintf("%v", hs.Has("nope")),
		)
	}
}

// ==========================================================================
// Test: Hashset — Remove edge cases
// ==========================================================================

func Test_Hashset_Remove_Edge(t *testing.T) {
	// Case 0: remove existing
	{
		tc := hashsetRemoveEdgeTestCases[0]
		hs := coregeneric.HashsetFrom([]int{1, 2, 3})
		existed := hs.Remove(2)
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%v", existed),
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has(2)),
		)
	}

	// Case 1: remove non-existing
	{
		tc := hashsetRemoveEdgeTestCases[1]
		hs := coregeneric.HashsetFrom([]int{1, 2, 3})
		existed := hs.Remove(99)
		tc.ShouldBeEqual(t, 1,
			fmt.Sprintf("%v", existed),
			fmt.Sprintf("%d", hs.Length()),
		)
	}
}

// ==========================================================================
// Test: Hashset — Has / Contains
// ==========================================================================

func Test_Hashset_Has_Edge(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a", "b", "c"})

	// Case 0: Has
	{
		tc := hashsetHasEdgeTestCases[0]
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%v", hs.Has("a")),
			fmt.Sprintf("%v", hs.Has("z")),
		)
	}

	// Case 1: Contains alias
	{
		tc := hashsetHasEdgeTestCases[1]
		tc.ShouldBeEqual(t, 1,
			fmt.Sprintf("%v", hs.Contains("b")),
			fmt.Sprintf("%v", hs.Contains("z")),
		)
	}
}

// ==========================================================================
// Test: Hashset — HasAll / HasAny
// ==========================================================================

func Test_Hashset_HasAll_HasAny(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})

	// Case 0: HasAll true
	{
		tc := hashsetHasAllAnyTestCases[0]
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", hs.HasAll(1, 3, 5)))
	}

	// Case 1: HasAll false
	{
		tc := hashsetHasAllAnyTestCases[1]
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%v", hs.HasAll(1, 99)))
	}

	// Case 2: HasAny true
	{
		tc := hashsetHasAllAnyTestCases[2]
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%v", hs.HasAny(99, 3)))
	}

	// Case 3: HasAny false
	{
		tc := hashsetHasAllAnyTestCases[3]
		tc.ShouldBeEqual(t, 3, fmt.Sprintf("%v", hs.HasAny(99, 100)))
	}

	// Case 4: HasAll empty args
	{
		tc := hashsetHasAllAnyTestCases[4]
		tc.ShouldBeEqual(t, 4, fmt.Sprintf("%v", hs.HasAll()))
	}

	// Case 5: HasAny empty args
	{
		tc := hashsetHasAllAnyTestCases[5]
		tc.ShouldBeEqual(t, 5, fmt.Sprintf("%v", hs.HasAny()))
	}
}

// ==========================================================================
// Test: Hashset — IsEquals
// ==========================================================================

func Test_Hashset_IsEquals(t *testing.T) {
	// Case 0: same items
	{
		tc := hashsetIsEqualsTestCases[0]
		a := coregeneric.HashsetFrom([]int{1, 2, 3})
		b := coregeneric.HashsetFrom([]int{3, 2, 1})
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEquals(b)))
	}

	// Case 1: different items
	{
		tc := hashsetIsEqualsTestCases[1]
		a := coregeneric.HashsetFrom([]int{1, 2, 3})
		b := coregeneric.HashsetFrom([]int{1, 2, 4})
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%v", a.IsEquals(b)))
	}

	// Case 2: different length
	{
		tc := hashsetIsEqualsTestCases[2]
		a := coregeneric.HashsetFrom([]int{1, 2})
		b := coregeneric.HashsetFrom([]int{1, 2, 3})
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%v", a.IsEquals(b)))
	}

	// Case 3: both nil
	{
		tc := hashsetIsEqualsTestCases[3]
		var a *coregeneric.Hashset[int]
		var b *coregeneric.Hashset[int]
		tc.ShouldBeEqual(t, 3, fmt.Sprintf("%v", a.IsEquals(b)))
	}

	// Case 4: nil vs non-nil
	{
		tc := hashsetIsEqualsTestCases[4]
		var a *coregeneric.Hashset[int]
		b := coregeneric.EmptyHashset[int]()
		tc.ShouldBeEqual(t, 4, fmt.Sprintf("%v", a.IsEquals(b)))
	}

	// Case 5: same pointer
	{
		tc := hashsetIsEqualsTestCases[5]
		a := coregeneric.HashsetFrom([]int{1, 2})
		tc.ShouldBeEqual(t, 5, fmt.Sprintf("%v", a.IsEquals(a)))
	}

	// Case 6: both empty
	{
		tc := hashsetIsEqualsTestCases[6]
		a := coregeneric.EmptyHashset[int]()
		b := coregeneric.EmptyHashset[int]()
		tc.ShouldBeEqual(t, 6, fmt.Sprintf("%v", a.IsEquals(b)))
	}
}

// ==========================================================================
// Test: Hashset — Resize
// ==========================================================================

func Test_Hashset_Resize(t *testing.T) {
	// Case 0: larger capacity
	{
		tc := hashsetResizeTestCases[0]
		hs := coregeneric.HashsetFrom([]int{1, 2, 3})
		hs.Resize(100)
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has(1)),
			fmt.Sprintf("%v", hs.Has(2)),
			fmt.Sprintf("%v", hs.Has(3)),
		)
	}

	// Case 1: smaller than length — no-op
	{
		tc := hashsetResizeTestCases[1]
		hs := coregeneric.HashsetFrom([]int{1, 2, 3})
		hs.Resize(1)
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", hs.Length()))
	}
}

// ==========================================================================
// Test: Hashset — List / ListPtr / Map / Collection / String
// ==========================================================================

func Test_Hashset_Output(t *testing.T) {
	// Case 0: List
	{
		tc := hashsetOutputTestCases[0]
		hs := coregeneric.HashsetFrom([]int{1, 2, 3})
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", len(hs.List())))
	}

	// Case 1: List on empty
	{
		tc := hashsetOutputTestCases[1]
		hs := coregeneric.EmptyHashset[int]()
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", len(hs.List())))
	}

	// Case 2: ListPtr non-nil
	{
		tc := hashsetOutputTestCases[2]
		hs := coregeneric.HashsetFrom([]int{1, 2, 3})
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%v", hs.ListPtr() != nil))
	}

	// Case 3: Map
	{
		tc := hashsetOutputTestCases[3]
		hs := coregeneric.HashsetFrom([]int{1, 2, 3})
		tc.ShouldBeEqual(t, 3, fmt.Sprintf("%d", len(hs.Map())))
	}

	// Case 4: Collection
	{
		tc := hashsetOutputTestCases[4]
		hs := coregeneric.HashsetFrom([]int{1, 2, 3})
		col := hs.Collection()
		tc.ShouldBeEqual(t, 4, fmt.Sprintf("%d", col.Length()))
	}
}

// ==========================================================================
// Test: Hashset — Lock variants
// ==========================================================================

func Test_Hashset_Lock(t *testing.T) {
	// Case 0: AddLock + ContainsLock
	{
		tc := hashsetLockTestCases[0]
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

	// Case 1: AddSliceLock
	{
		tc := hashsetLockTestCases[1]
		hs := coregeneric.EmptyHashset[int]()
		hs.AddSliceLock([]int{10, 20, 30})
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", hs.Length()))
	}

	// Case 2: RemoveLock
	{
		tc := hashsetLockTestCases[2]
		hs := coregeneric.HashsetFrom([]int{1, 2, 3})
		existed := hs.RemoveLock(2)
		tc.ShouldBeEqual(t, 2,
			fmt.Sprintf("%v", existed),
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has(2)),
		)
	}

	// Case 3: IsEmptyLock + LengthLock
	{
		tc := hashsetLockTestCases[3]
		hs := coregeneric.EmptyHashset[int]()
		emptyBefore := fmt.Sprintf("%v", hs.IsEmptyLock())
		lenBefore := fmt.Sprintf("%d", hs.LengthLock())
		hs.Adds(1, 2)
		emptyAfter := fmt.Sprintf("%v", hs.IsEmptyLock())
		lenAfter := fmt.Sprintf("%d", hs.LengthLock())
		tc.ShouldBeEqual(t, 3, emptyBefore, lenBefore, emptyAfter, lenAfter)
	}
}

// ==========================================================================
// Test: Hashset — Constructors
// ==========================================================================

func Test_Hashset_Constructors(t *testing.T) {
	// Case 0: EmptyHashset
	{
		tc := hashsetConstructorTestCases[0]
		hs := coregeneric.EmptyHashset[int]()
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.IsEmpty()),
		)
	}

	// Case 1: NewHashset with capacity
	{
		tc := hashsetConstructorTestCases[1]
		hs := coregeneric.NewHashset[string](10)
		tc.ShouldBeEqual(t, 1,
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.IsEmpty()),
		)
	}

	// Case 2: HashsetFrom
	{
		tc := hashsetConstructorTestCases[2]
		hs := coregeneric.HashsetFrom([]string{"a", "b", "c"})
		tc.ShouldBeEqual(t, 2,
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has("a")),
			fmt.Sprintf("%v", hs.Has("b")),
			fmt.Sprintf("%v", hs.Has("c")),
		)
	}

	// Case 3: HashsetFromMap
	{
		tc := hashsetConstructorTestCases[3]
		hs := coregeneric.HashsetFromMap(map[int]bool{10: true, 20: true})
		tc.ShouldBeEqual(t, 3,
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has(10)),
			fmt.Sprintf("%v", hs.Has(20)),
		)
	}

	// Case 4: HasItems
	{
		tc := hashsetConstructorTestCases[4]
		pop := coregeneric.HashsetFrom([]int{1})
		empty := coregeneric.EmptyHashset[int]()
		tc.ShouldBeEqual(t, 4,
			fmt.Sprintf("%v", pop.HasItems()),
			fmt.Sprintf("%v", empty.HasItems()),
		)
	}
}
