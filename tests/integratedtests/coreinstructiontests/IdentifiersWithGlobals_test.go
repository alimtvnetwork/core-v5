package coreinstructiontests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coreinstruction"
)

// ==========================================================================
// Test: Length
// ==========================================================================

func Test_IdentifiersWithGlobals_Length_Empty(t *testing.T) {
	tc := idsLengthEmptyTestCase
	ids := coreinstruction.EmptyIdentifiersWithGlobals()
	actLines := []string{fmt.Sprintf("%v", ids.Length())}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_IdentifiersWithGlobals_Length_ThreeItems(t *testing.T) {
	tc := idsLengthThreeItemsTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "a", "b", "c")
	actLines := []string{fmt.Sprintf("%v", ids.Length())}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_IdentifiersWithGlobals_Length_Nil(t *testing.T) {
	tc := idsLengthNilTestCase
	var nilIds *coreinstruction.IdentifiersWithGlobals
	actLines := []string{fmt.Sprintf("%v", nilIds.Length())}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: GetById
// ==========================================================================

func Test_IdentifiersWithGlobals_GetById_Found(t *testing.T) {
	tc := idsGetByIdFoundTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "alpha", "beta")
	found := ids.GetById("beta")
	actLines := []string{
		fmt.Sprintf("%v", found != nil),
		found.Id,
		fmt.Sprintf("%v", found.IsGlobal),
	}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_IdentifiersWithGlobals_GetById_Missing(t *testing.T) {
	tc := idsGetByIdMissingTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(false, "alpha")
	actLines := []string{fmt.Sprintf("%v", ids.GetById("missing") == nil)}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_IdentifiersWithGlobals_GetById_EmptyId(t *testing.T) {
	tc := idsGetByIdEmptyTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "alpha")
	actLines := []string{fmt.Sprintf("%v", ids.GetById("") == nil)}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: Clone
// ==========================================================================

func Test_IdentifiersWithGlobals_Clone_Independence(t *testing.T) {
	tc := idsCloneIndependenceTestCase
	orig := coreinstruction.NewIdentifiersWithGlobals(true, "x", "y")
	cloned := orig.Clone()
	cloned.Add(false, "z")
	actLines := []string{
		fmt.Sprintf("%v", orig.Length()),
		fmt.Sprintf("%v", cloned.Length()),
	}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_IdentifiersWithGlobals_Clone_Empty(t *testing.T) {
	tc := idsCloneEmptyTestCase
	orig := coreinstruction.EmptyIdentifiersWithGlobals()
	cloned := orig.Clone()
	actLines := []string{
		fmt.Sprintf("%v", cloned != nil),
		fmt.Sprintf("%v", cloned.Length()),
	}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_IdentifiersWithGlobals_Clone_Preserves(t *testing.T) {
	tc := idsClonePreservesTestCase
	orig := coreinstruction.NewIdentifiersWithGlobals(false, "id-1", "id-2")
	cloned := orig.Clone()
	item := cloned.GetById("id-1")
	actLines := []string{
		fmt.Sprintf("%v", item != nil),
		item.Id,
		fmt.Sprintf("%v", item.IsGlobal),
	}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: Add
// ==========================================================================

func Test_IdentifiersWithGlobals_Add_Single(t *testing.T) {
	tc := idsAddSingleTestCase
	ids := coreinstruction.EmptyIdentifiersWithGlobals()
	ids.Add(true, "new-id")
	found := ids.GetById("new-id")
	actLines := []string{
		fmt.Sprintf("%v", ids.Length()),
		fmt.Sprintf("%v", found != nil),
		fmt.Sprintf("%v", found.IsGlobal),
	}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_IdentifiersWithGlobals_Add_EmptyIdIgnored(t *testing.T) {
	tc := idsAddEmptyIdTestCase
	ids := coreinstruction.EmptyIdentifiersWithGlobals()
	ids.Add(true, "")
	actLines := []string{fmt.Sprintf("%v", ids.Length())}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_IdentifiersWithGlobals_Add_MultipleAccumulate(t *testing.T) {
	tc := idsAddMultipleTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(false, "existing")
	ids.Add(true, "second")
	ids.Add(false, "third")
	actLines := []string{
		fmt.Sprintf("%v", ids.Length()),
		fmt.Sprintf("%v", ids.GetById("second").IsGlobal),
		fmt.Sprintf("%v", ids.GetById("third").IsGlobal),
	}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: IsEmpty / HasAnyItem
// ==========================================================================

func Test_IdentifiersWithGlobals_IsEmpty_True(t *testing.T) {
	tc := idsIsEmptyTrueTestCase
	ids := coreinstruction.EmptyIdentifiersWithGlobals()
	actLines := []string{
		fmt.Sprintf("%v", ids.IsEmpty()),
		fmt.Sprintf("%v", ids.HasAnyItem()),
	}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_IdentifiersWithGlobals_IsEmpty_False(t *testing.T) {
	tc := idsIsEmptyFalseTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "item")
	actLines := []string{
		fmt.Sprintf("%v", ids.IsEmpty()),
		fmt.Sprintf("%v", ids.HasAnyItem()),
	}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: IndexOf
// ==========================================================================

func Test_IdentifiersWithGlobals_IndexOf_Found(t *testing.T) {
	tc := idsIndexOfFoundTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "a", "b", "c")
	actLines := []string{
		fmt.Sprintf("%v", ids.IndexOf("a")),
		fmt.Sprintf("%v", ids.IndexOf("b")),
		fmt.Sprintf("%v", ids.IndexOf("c")),
	}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_IdentifiersWithGlobals_IndexOf_Missing(t *testing.T) {
	tc := idsIndexOfMissingTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(false, "x")
	actLines := []string{fmt.Sprintf("%v", ids.IndexOf("missing"))}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_IdentifiersWithGlobals_IndexOf_EmptyString(t *testing.T) {
	tc := idsIndexOfEmptyStringTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "a")
	actLines := []string{fmt.Sprintf("%v", ids.IndexOf(""))}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_IdentifiersWithGlobals_IndexOf_EmptyCollection(t *testing.T) {
	tc := idsIndexOfEmptyCollectionTestCase
	ids := coreinstruction.EmptyIdentifiersWithGlobals()
	actLines := []string{fmt.Sprintf("%v", ids.IndexOf("any"))}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: Adds
// ==========================================================================

func Test_IdentifiersWithGlobals_Adds_Batch(t *testing.T) {
	tc := idsAddsBatchTestCase
	ids := coreinstruction.EmptyIdentifiersWithGlobals()
	ids.Adds(true, "one", "two", "three")
	actLines := []string{
		fmt.Sprintf("%v", ids.Length()),
		fmt.Sprintf("%v", ids.GetById("one") != nil),
		fmt.Sprintf("%v", ids.GetById("two") != nil),
		fmt.Sprintf("%v", ids.GetById("three") != nil),
	}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_IdentifiersWithGlobals_Adds_Empty(t *testing.T) {
	tc := idsAddsEmptyTestCase
	ids := coreinstruction.EmptyIdentifiersWithGlobals()
	ids.Adds(true)
	actLines := []string{fmt.Sprintf("%v", ids.Length())}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: New edge
// ==========================================================================

func Test_IdentifiersWithGlobals_NewEdge(t *testing.T) {
	tc := idsNewEdgeEmptyTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(true)
	actLines := []string{
		fmt.Sprintf("%v", ids != nil),
		fmt.Sprintf("%v", ids.Length()),
	}

	tc.ShouldBeEqual(t, 0, actLines...)
}
