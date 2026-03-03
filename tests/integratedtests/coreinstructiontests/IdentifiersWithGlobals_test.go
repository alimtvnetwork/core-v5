package coreinstructiontests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coreinstruction"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================================================
// Test: Length
// ==========================================================================

func Test_IdentifiersWithGlobals_Length(t *testing.T) {
	// Case 0: empty
	{
		tc := idsLengthTestCases[0]
		ids := coreinstruction.EmptyIdentifiersWithGlobals()
		actLines := []string{fmt.Sprintf("%v", ids.Length())}
		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: 3 items
	{
		tc := idsLengthTestCases[1]
		ids := coreinstruction.NewIdentifiersWithGlobals(true, "a", "b", "c")
		actLines := []string{fmt.Sprintf("%v", ids.Length())}
		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 2: nil
	{
		tc := idsLengthTestCases[2]
		var nilIds *coreinstruction.IdentifiersWithGlobals
		actLines := []string{fmt.Sprintf("%v", nilIds.Length())}
		errcore.AssertDiffOnMismatch(t, 2, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: GetById
// ==========================================================================

func Test_IdentifiersWithGlobals_GetById(t *testing.T) {
	// Case 0: found
	{
		tc := idsGetByIdTestCases[0]
		ids := coreinstruction.NewIdentifiersWithGlobals(true, "alpha", "beta")
		found := ids.GetById("beta")
		actLines := []string{
			fmt.Sprintf("%v", found != nil),
			found.Id,
			fmt.Sprintf("%v", found.IsGlobal),
		}
		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: missing
	{
		tc := idsGetByIdTestCases[1]
		ids := coreinstruction.NewIdentifiersWithGlobals(false, "alpha")
		actLines := []string{fmt.Sprintf("%v", ids.GetById("missing") == nil)}
		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 2: empty id
	{
		tc := idsGetByIdTestCases[2]
		ids := coreinstruction.NewIdentifiersWithGlobals(true, "alpha")
		actLines := []string{fmt.Sprintf("%v", ids.GetById("") == nil)}
		errcore.AssertDiffOnMismatch(t, 2, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: Clone
// ==========================================================================

func Test_IdentifiersWithGlobals_Clone(t *testing.T) {
	// Case 0: independence
	{
		tc := idsCloneTestCases[0]
		orig := coreinstruction.NewIdentifiersWithGlobals(true, "x", "y")
		cloned := orig.Clone()
		cloned.Add(false, "z")
		actLines := []string{
			fmt.Sprintf("%v", orig.Length()),
			fmt.Sprintf("%v", cloned.Length()),
		}
		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: empty
	{
		tc := idsCloneTestCases[1]
		orig := coreinstruction.EmptyIdentifiersWithGlobals()
		cloned := orig.Clone()
		actLines := []string{
			fmt.Sprintf("%v", cloned != nil),
			fmt.Sprintf("%v", cloned.Length()),
		}
		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 2: preserves
	{
		tc := idsCloneTestCases[2]
		orig := coreinstruction.NewIdentifiersWithGlobals(false, "id-1", "id-2")
		cloned := orig.Clone()
		item := cloned.GetById("id-1")
		actLines := []string{
			fmt.Sprintf("%v", item != nil),
			item.Id,
			fmt.Sprintf("%v", item.IsGlobal),
		}
		errcore.AssertDiffOnMismatch(t, 2, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: Add
// ==========================================================================

func Test_IdentifiersWithGlobals_Add(t *testing.T) {
	// Case 0: single
	{
		tc := idsAddTestCases[0]
		ids := coreinstruction.EmptyIdentifiersWithGlobals()
		ids.Add(true, "new-id")
		found := ids.GetById("new-id")
		actLines := []string{
			fmt.Sprintf("%v", ids.Length()),
			fmt.Sprintf("%v", found != nil),
			fmt.Sprintf("%v", found.IsGlobal),
		}
		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: empty id ignored
	{
		tc := idsAddTestCases[1]
		ids := coreinstruction.EmptyIdentifiersWithGlobals()
		ids.Add(true, "")
		actLines := []string{fmt.Sprintf("%v", ids.Length())}
		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 2: multiple accumulate
	{
		tc := idsAddTestCases[2]
		ids := coreinstruction.NewIdentifiersWithGlobals(false, "existing")
		ids.Add(true, "second")
		ids.Add(false, "third")
		actLines := []string{
			fmt.Sprintf("%v", ids.Length()),
			fmt.Sprintf("%v", ids.GetById("second").IsGlobal),
			fmt.Sprintf("%v", ids.GetById("third").IsGlobal),
		}
		errcore.AssertDiffOnMismatch(t, 2, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: IsEmpty / HasAnyItem
// ==========================================================================

func Test_IdentifiersWithGlobals_IsEmpty(t *testing.T) {
	// Case 0: empty true
	{
		tc := idsIsEmptyTestCases[0]
		ids := coreinstruction.EmptyIdentifiersWithGlobals()
		actLines := []string{
			fmt.Sprintf("%v", ids.IsEmpty()),
			fmt.Sprintf("%v", ids.HasAnyItem()),
		}
		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: non-empty false
	{
		tc := idsIsEmptyTestCases[1]
		ids := coreinstruction.NewIdentifiersWithGlobals(true, "item")
		actLines := []string{
			fmt.Sprintf("%v", ids.IsEmpty()),
			fmt.Sprintf("%v", ids.HasAnyItem()),
		}
		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: IndexOf
// ==========================================================================

func Test_IdentifiersWithGlobals_IndexOf(t *testing.T) {
	// Case 0: found
	{
		tc := idsIndexOfTestCases[0]
		ids := coreinstruction.NewIdentifiersWithGlobals(true, "a", "b", "c")
		actLines := []string{
			fmt.Sprintf("%v", ids.IndexOf("a")),
			fmt.Sprintf("%v", ids.IndexOf("b")),
			fmt.Sprintf("%v", ids.IndexOf("c")),
		}
		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: missing
	{
		tc := idsIndexOfTestCases[1]
		ids := coreinstruction.NewIdentifiersWithGlobals(false, "x")
		actLines := []string{fmt.Sprintf("%v", ids.IndexOf("missing"))}
		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 2: empty string
	{
		tc := idsIndexOfTestCases[2]
		ids := coreinstruction.NewIdentifiersWithGlobals(true, "a")
		actLines := []string{fmt.Sprintf("%v", ids.IndexOf(""))}
		errcore.AssertDiffOnMismatch(t, 2, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 3: empty collection
	{
		tc := idsIndexOfTestCases[3]
		ids := coreinstruction.EmptyIdentifiersWithGlobals()
		actLines := []string{fmt.Sprintf("%v", ids.IndexOf("any"))}
		errcore.AssertDiffOnMismatch(t, 3, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: Adds
// ==========================================================================

func Test_IdentifiersWithGlobals_Adds(t *testing.T) {
	// Case 0: batch
	{
		tc := idsAddsTestCases[0]
		ids := coreinstruction.EmptyIdentifiersWithGlobals()
		ids.Adds(true, "one", "two", "three")
		actLines := []string{
			fmt.Sprintf("%v", ids.Length()),
			fmt.Sprintf("%v", ids.GetById("one") != nil),
			fmt.Sprintf("%v", ids.GetById("two") != nil),
			fmt.Sprintf("%v", ids.GetById("three") != nil),
		}
		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: empty
	{
		tc := idsAddsTestCases[1]
		ids := coreinstruction.EmptyIdentifiersWithGlobals()
		ids.Adds(true)
		actLines := []string{fmt.Sprintf("%v", ids.Length())}
		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: New edge
// ==========================================================================

func Test_IdentifiersWithGlobals_NewEdge(t *testing.T) {
	tc := idsNewEdgeTestCases[0]
	ids := coreinstruction.NewIdentifiersWithGlobals(true)
	actLines := []string{
		fmt.Sprintf("%v", ids != nil),
		fmt.Sprintf("%v", ids.Length()),
	}
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}
