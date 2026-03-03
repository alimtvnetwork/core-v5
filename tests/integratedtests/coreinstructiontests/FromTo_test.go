package coreinstructiontests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coreinstruction"
)

// ==========================================
// Test: ClonePtr
// ==========================================

func Test_FromTo_ClonePtr(t *testing.T) {
	// Case 0: positive
	{
		tc := fromToClonePtrCopiesTestCase
		orig := &coreinstruction.FromTo{From: "source", To: "destination"}
		cloned := orig.ClonePtr()

		actLines := []string{
			fmt.Sprintf("%v", cloned != nil),
			cloned.From,
			cloned.To,
		}

		tc.ShouldBeEqual(t, 0, actLines...)
	}

	// Case 1: nil receiver
	{
		tc := fromToClonePtrNilTestCase
		var nilFT *coreinstruction.FromTo

		actLines := []string{
			fmt.Sprintf("%v", nilFT.ClonePtr() == nil),
		}

		tc.ShouldBeEqual(t, 1, actLines...)
	}
}

// ==========================================
// Test: Clone
// ==========================================

func Test_FromTo_Clone(t *testing.T) {
	tc := fromToCloneCopiesTestCase
	orig := coreinstruction.FromTo{From: "a", To: "b"}
	c := orig.Clone()

	actLines := []string{c.From, c.To}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================
// Test: IsNull
// ==========================================

func Test_FromTo_IsNull(t *testing.T) {
	// Case 0: nil returns true
	{
		tc := fromToIsNullNilTestCase
		var nilFT *coreinstruction.FromTo

		actLines := []string{fmt.Sprintf("%v", nilFT.IsNull())}

		tc.ShouldBeEqual(t, 0, actLines...)
	}

	// Case 1: non-nil returns false
	{
		tc := fromToIsNullNonNilTestCase
		ft := &coreinstruction.FromTo{From: "x", To: "y"}

		actLines := []string{fmt.Sprintf("%v", ft.IsNull())}

		tc.ShouldBeEqual(t, 1, actLines...)
	}
}

// ==========================================
// Test: IsFromEmpty
// ==========================================

func Test_FromTo_IsFromEmpty(t *testing.T) {
	// Case 0: empty From returns true
	{
		tc := fromToIsFromEmptyEmptyTestCase
		ft := &coreinstruction.FromTo{From: "", To: "dest"}

		actLines := []string{fmt.Sprintf("%v", ft.IsFromEmpty())}

		tc.ShouldBeEqual(t, 0, actLines...)
	}

	// Case 1: nil receiver returns true
	{
		tc := fromToIsFromEmptyNilTestCase
		var nilFT *coreinstruction.FromTo

		actLines := []string{fmt.Sprintf("%v", nilFT.IsFromEmpty())}

		tc.ShouldBeEqual(t, 1, actLines...)
	}
}

// ==========================================
// Test: IsToEmpty
// ==========================================

func Test_FromTo_IsToEmpty(t *testing.T) {
	// Case 0: empty To returns true
	{
		tc := fromToIsToEmptyEmptyTestCase
		ft := &coreinstruction.FromTo{From: "src", To: ""}

		actLines := []string{fmt.Sprintf("%v", ft.IsToEmpty())}

		tc.ShouldBeEqual(t, 0, actLines...)
	}

	// Case 1: non-empty returns false
	{
		tc := fromToIsToEmptyNonEmptyTestCase
		ft := &coreinstruction.FromTo{From: "src", To: "dest"}

		actLines := []string{fmt.Sprintf("%v", ft.IsToEmpty())}

		tc.ShouldBeEqual(t, 1, actLines...)
	}
}

// ==========================================
// Test: String
// ==========================================

func Test_FromTo_String(t *testing.T) {
	tc := fromToStringContainsTestCase
	ft := coreinstruction.FromTo{From: "alpha", To: "beta"}
	s := ft.String()

	actLines := []string{
		fmt.Sprintf("%v", len(s) > 0 && strings.Contains(s, "alpha")),
		fmt.Sprintf("%v", strings.Contains(s, "beta")),
	}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================
// Test: FromName / ToName
// ==========================================

func Test_FromTo_Names(t *testing.T) {
	tc := fromToNamesTestCase
	ft := coreinstruction.FromTo{From: "src", To: "dst"}

	actLines := []string{ft.FromName(), ft.ToName()}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================
// Test: SetFromName
// ==========================================

func Test_FromTo_SetFromName(t *testing.T) {
	// Case 0: updates From
	{
		tc := fromToSetFromNameUpdatesTestCase
		ft := &coreinstruction.FromTo{From: "old", To: "t"}
		ft.SetFromName("new")

		actLines := []string{ft.From}

		tc.ShouldBeEqual(t, 0, actLines...)
	}

	// Case 1: nil receiver no panic
	{
		tc := fromToSetFromNameNilTestCase
		var nilFT *coreinstruction.FromTo
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			nilFT.SetFromName("x")
		}()

		actLines := []string{fmt.Sprintf("%v", !didPanic)}

		tc.ShouldBeEqual(t, 1, actLines...)
	}
}

// ==========================================
// Test: SetToName
// ==========================================

func Test_FromTo_SetToName(t *testing.T) {
	tc := fromToSetToNameUpdatesTestCase
	ft := &coreinstruction.FromTo{From: "f", To: "old"}
	ft.SetToName("new")

	actLines := []string{ft.To}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================
// Test: SourceDestination
// ==========================================

func Test_FromTo_SourceDestination(t *testing.T) {
	// Case 0: maps From->Source To->Destination
	{
		tc := fromToSourceDestMapsTestCase
		ft := &coreinstruction.FromTo{From: "src", To: "dst"}
		sd := ft.SourceDestination()

		actLines := []string{
			fmt.Sprintf("%v", sd != nil),
			sd.Source,
			sd.Destination,
		}

		tc.ShouldBeEqual(t, 0, actLines...)
	}

	// Case 1: nil returns nil
	{
		tc := fromToSourceDestNilTestCase
		var nilFT *coreinstruction.FromTo

		actLines := []string{fmt.Sprintf("%v", nilFT.SourceDestination() == nil)}

		tc.ShouldBeEqual(t, 1, actLines...)
	}
}

// ==========================================
// Test: Rename
// ==========================================

func Test_FromTo_Rename(t *testing.T) {
	// Case 0: maps From->Existing To->New
	{
		tc := fromToRenameMapsTestCase
		ft := &coreinstruction.FromTo{From: "old", To: "new"}
		rn := ft.Rename()

		actLines := []string{
			fmt.Sprintf("%v", rn != nil),
			rn.Existing,
			rn.New,
		}

		tc.ShouldBeEqual(t, 0, actLines...)
	}

	// Case 1: nil returns nil
	{
		tc := fromToRenameNilTestCase
		var nilFT *coreinstruction.FromTo

		actLines := []string{fmt.Sprintf("%v", nilFT.Rename() == nil)}

		tc.ShouldBeEqual(t, 1, actLines...)
	}
}
