package coreinstructiontests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Test: Length
// ==========================================================================

func Test_IdentifiersWithGlobals_Length_Empty(t *testing.T) {
	tc := idsLengthEmptyTestCase
	ids := coreinstruction.EmptyIdentifiersWithGlobals()
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", ids.Length()))
}

func Test_IdentifiersWithGlobals_Length_ThreeItems(t *testing.T) {
	tc := idsLengthThreeItemsTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "a", "b", "c")
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", ids.Length()))
}

func Test_IdentifiersWithGlobals_Length_Nil(t *testing.T) {
	tc := idsLengthNilTestCase
	var nilIds *coreinstruction.IdentifiersWithGlobals
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", nilIds.Length()))
}

// ==========================================================================
// Test: GetById
// ==========================================================================

func Test_IdentifiersWithGlobals_GetById_Found(t *testing.T) {
	tc := idsGetByIdFoundTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "alpha", "beta")
	found := ids.GetById("beta")

	actual := args.Map{
		"found":    found != nil,
		"id":       found.Id,
		"isGlobal": found.IsGlobal,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IdentifiersWithGlobals_GetById_Missing(t *testing.T) {
	tc := idsGetByIdMissingTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(false, "alpha")
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", ids.GetById("missing") == nil))
}

func Test_IdentifiersWithGlobals_GetById_EmptyId(t *testing.T) {
	tc := idsGetByIdEmptyTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "alpha")
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", ids.GetById("") == nil))
}

// ==========================================================================
// Test: Clone
// ==========================================================================

func Test_IdentifiersWithGlobals_Clone_Independence(t *testing.T) {
	tc := idsCloneIndependenceTestCase
	orig := coreinstruction.NewIdentifiersWithGlobals(true, "x", "y")
	cloned := orig.Clone()
	cloned.Add(false, "z")

	actual := args.Map{
		"originalLength": orig.Length(),
		"cloneLength":    cloned.Length(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IdentifiersWithGlobals_Clone_Empty(t *testing.T) {
	tc := idsCloneEmptyTestCase
	orig := coreinstruction.EmptyIdentifiersWithGlobals()
	cloned := orig.Clone()

	actual := args.Map{
		"isNotNil": cloned != nil,
		"length":   cloned.Length(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IdentifiersWithGlobals_Clone_Preserves(t *testing.T) {
	tc := idsClonePreservesTestCase
	orig := coreinstruction.NewIdentifiersWithGlobals(false, "id-1", "id-2")
	cloned := orig.Clone()
	item := cloned.GetById("id-1")

	actual := args.Map{
		"isNotNil": item != nil,
		"firstId":  item.Id,
		"isGlobal": item.IsGlobal,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Add
// ==========================================================================

func Test_IdentifiersWithGlobals_Add_Single(t *testing.T) {
	tc := idsAddSingleTestCase
	ids := coreinstruction.EmptyIdentifiersWithGlobals()
	ids.Add(true, "new-id")
	found := ids.GetById("new-id")

	actual := args.Map{
		"length":   ids.Length(),
		"found":    found != nil,
		"isGlobal": found.IsGlobal,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IdentifiersWithGlobals_Add_EmptyIdIgnored(t *testing.T) {
	tc := idsAddEmptyIdTestCase
	ids := coreinstruction.EmptyIdentifiersWithGlobals()
	ids.Add(true, "")
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", ids.Length()))
}

func Test_IdentifiersWithGlobals_Add_MultipleAccumulate(t *testing.T) {
	tc := idsAddMultipleTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(false, "existing")
	ids.Add(true, "second")
	ids.Add(false, "third")

	actual := args.Map{
		"length":         ids.Length(),
		"secondIsGlobal": ids.GetById("second").IsGlobal,
		"thirdIsGlobal":  ids.GetById("third").IsGlobal,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: IsEmpty / HasAnyItem
// ==========================================================================

func Test_IdentifiersWithGlobals_IsEmpty_True(t *testing.T) {
	tc := idsIsEmptyTrueTestCase
	ids := coreinstruction.EmptyIdentifiersWithGlobals()

	actual := args.Map{
		"isEmpty":    ids.IsEmpty(),
		"hasAnyItem": ids.HasAnyItem(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IdentifiersWithGlobals_IsEmpty_False(t *testing.T) {
	tc := idsIsEmptyFalseTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "item")

	actual := args.Map{
		"isEmpty":    ids.IsEmpty(),
		"hasAnyItem": ids.HasAnyItem(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: IndexOf
// ==========================================================================

func Test_IdentifiersWithGlobals_IndexOf_Found(t *testing.T) {
	tc := idsIndexOfFoundTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "a", "b", "c")

	actual := args.Map{
		"indexOfA": ids.IndexOf("a"),
		"indexOfB": ids.IndexOf("b"),
		"indexOfC": ids.IndexOf("c"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IdentifiersWithGlobals_IndexOf_Missing(t *testing.T) {
	tc := idsIndexOfMissingTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(false, "x")
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", ids.IndexOf("missing")))
}

func Test_IdentifiersWithGlobals_IndexOf_EmptyString(t *testing.T) {
	tc := idsIndexOfEmptyStringTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "a")
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", ids.IndexOf("")))
}

func Test_IdentifiersWithGlobals_IndexOf_EmptyCollection(t *testing.T) {
	tc := idsIndexOfEmptyCollectionTestCase
	ids := coreinstruction.EmptyIdentifiersWithGlobals()
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", ids.IndexOf("any")))
}

// ==========================================================================
// Test: Adds
// ==========================================================================

func Test_IdentifiersWithGlobals_Adds_Batch(t *testing.T) {
	tc := idsAddsBatchTestCase
	ids := coreinstruction.EmptyIdentifiersWithGlobals()
	ids.Adds(true, "one", "two", "three")

	actual := args.Map{
		"length":     ids.Length(),
		"foundOne":   ids.GetById("one") != nil,
		"foundTwo":   ids.GetById("two") != nil,
		"foundThree": ids.GetById("three") != nil,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IdentifiersWithGlobals_Adds_Empty(t *testing.T) {
	tc := idsAddsEmptyTestCase
	ids := coreinstruction.EmptyIdentifiersWithGlobals()
	ids.Adds(true)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", ids.Length()))
}

// ==========================================================================
// Test: New edge
// ==========================================================================

func Test_IdentifiersWithGlobals_NewEdge(t *testing.T) {
	tc := idsNewEdgeEmptyTestCase
	ids := coreinstruction.NewIdentifiersWithGlobals(true)

	actual := args.Map{
		"isNotNil": ids != nil,
		"length":   ids.Length(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}
