package stringslicetests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/stringslice"
)

// ==========================================
// Clone
// ==========================================

func Test_StringSlice_Clone_NonEmpty(t *testing.T) {
	src := []string{"a", "b", "c"}
	result := stringslice.Clone(src)
	if len(result) != 3 {
		t.Errorf("Clone: expected 3, got %d", len(result))
	}
	// independence
	result[0] = "z"
	if src[0] == "z" {
		t.Error("Clone should produce independent copy")
	}
}

func Test_StringSlice_Clone_Empty(t *testing.T) {
	result := stringslice.Clone([]string{})
	if len(result) != 0 {
		t.Errorf("Clone empty: expected 0, got %d", len(result))
	}
}

func Test_StringSlice_Clone_Nil(t *testing.T) {
	result := stringslice.Clone(nil)
	if result == nil {
		t.Error("Clone nil should return non-nil empty slice")
	}
	if len(result) != 0 {
		t.Errorf("Clone nil: expected 0, got %d", len(result))
	}
}

// ==========================================
// CloneUsingCap
// ==========================================

func Test_StringSlice_CloneUsingCap_AddsCapacity(t *testing.T) {
	src := []string{"a"}
	result := stringslice.CloneUsingCap(10, src)
	if len(result) != 1 {
		t.Errorf("CloneUsingCap len: expected 1, got %d", len(result))
	}
	if cap(result) < 11 {
		t.Errorf("CloneUsingCap cap: expected >= 11, got %d", cap(result))
	}
}

func Test_StringSlice_CloneUsingCap_Empty(t *testing.T) {
	result := stringslice.CloneUsingCap(5, []string{})
	if len(result) != 0 {
		t.Errorf("expected 0, got %d", len(result))
	}
}

// ==========================================
// FirstOrDefault / LastOrDefault
// ==========================================

func Test_StringSlice_FirstOrDefault_NonEmpty(t *testing.T) {
	result := stringslice.FirstOrDefault([]string{"x", "y"})
	if result != "x" {
		t.Errorf("expected 'x', got '%s'", result)
	}
}

func Test_StringSlice_FirstOrDefault_Empty(t *testing.T) {
	result := stringslice.FirstOrDefault([]string{})
	if result != "" {
		t.Errorf("expected empty string, got '%s'", result)
	}
}

func Test_StringSlice_LastOrDefault_NonEmpty(t *testing.T) {
	result := stringslice.LastOrDefault([]string{"a", "b", "c"})
	if result != "c" {
		t.Errorf("expected 'c', got '%s'", result)
	}
}

func Test_StringSlice_LastOrDefault_Empty(t *testing.T) {
	result := stringslice.LastOrDefault([]string{})
	if result != "" {
		t.Errorf("expected empty string, got '%s'", result)
	}
}

func Test_StringSlice_LastOrDefault_Single(t *testing.T) {
	result := stringslice.LastOrDefault([]string{"only"})
	if result != "only" {
		t.Errorf("expected 'only', got '%s'", result)
	}
}

// ==========================================
// SafeIndexAt
// ==========================================

func Test_StringSlice_SafeIndexAt_Valid(t *testing.T) {
	result := stringslice.SafeIndexAt([]string{"a", "b", "c"}, 1)
	if result != "b" {
		t.Errorf("expected 'b', got '%s'", result)
	}
}

func Test_StringSlice_SafeIndexAt_OutOfBounds(t *testing.T) {
	result := stringslice.SafeIndexAt([]string{"a"}, 5)
	if result != "" {
		t.Errorf("expected empty string, got '%s'", result)
	}
}

func Test_StringSlice_SafeIndexAt_NegativeIndex(t *testing.T) {
	result := stringslice.SafeIndexAt([]string{"a"}, -1)
	if result != "" {
		t.Errorf("expected empty string, got '%s'", result)
	}
}

func Test_StringSlice_SafeIndexAt_EmptySlice(t *testing.T) {
	result := stringslice.SafeIndexAt([]string{}, 0)
	if result != "" {
		t.Errorf("expected empty string, got '%s'", result)
	}
}

// ==========================================
// InPlaceReverse
// ==========================================

func Test_StringSlice_InPlaceReverse_Multiple(t *testing.T) {
	s := []string{"a", "b", "c", "d"}
	result := stringslice.InPlaceReverse(&s)
	r := *result
	if r[0] != "d" || r[1] != "c" || r[2] != "b" || r[3] != "a" {
		t.Errorf("expected [d c b a], got %v", r)
	}
}

func Test_StringSlice_InPlaceReverse_Two(t *testing.T) {
	s := []string{"x", "y"}
	result := stringslice.InPlaceReverse(&s)
	r := *result
	if r[0] != "y" || r[1] != "x" {
		t.Errorf("expected [y x], got %v", r)
	}
}

func Test_StringSlice_InPlaceReverse_Single(t *testing.T) {
	s := []string{"only"}
	result := stringslice.InPlaceReverse(&s)
	if (*result)[0] != "only" {
		t.Error("single element should remain unchanged")
	}
}

func Test_StringSlice_InPlaceReverse_Empty(t *testing.T) {
	s := []string{}
	result := stringslice.InPlaceReverse(&s)
	if len(*result) != 0 {
		t.Error("empty should remain empty")
	}
}

func Test_StringSlice_InPlaceReverse_Nil(t *testing.T) {
	result := stringslice.InPlaceReverse(nil)
	if result == nil || len(*result) != 0 {
		t.Error("nil should return empty slice ptr")
	}
}

// ==========================================
// MergeNew
// ==========================================

func Test_StringSlice_MergeNew_BothNonEmpty(t *testing.T) {
	result := stringslice.MergeNew([]string{"a", "b"}, "c", "d")
	if len(result) != 4 {
		t.Errorf("expected 4, got %d", len(result))
	}
	if result[0] != "a" || result[3] != "d" {
		t.Errorf("unexpected order: %v", result)
	}
}

func Test_StringSlice_MergeNew_EmptyFirst(t *testing.T) {
	result := stringslice.MergeNew([]string{}, "x")
	if len(result) != 1 {
		t.Errorf("expected 1, got %d", len(result))
	}
}

func Test_StringSlice_MergeNew_NoAdditional(t *testing.T) {
	result := stringslice.MergeNew([]string{"a", "b"})
	if len(result) != 2 {
		t.Errorf("expected 2, got %d", len(result))
	}
}

func Test_StringSlice_MergeNew_BothEmpty(t *testing.T) {
	result := stringslice.MergeNew([]string{})
	if len(result) != 0 {
		t.Errorf("expected 0, got %d", len(result))
	}
}

// ==========================================
// NonEmptySlice
// ==========================================

func Test_StringSlice_NonEmpty_FiltersEmpty(t *testing.T) {
	result := stringslice.NonEmptySlice([]string{"a", "", "b", "", "c"})
	if len(result) != 3 {
		t.Errorf("expected 3, got %d", len(result))
	}
}

func Test_StringSlice_NonEmpty_AllEmpty(t *testing.T) {
	result := stringslice.NonEmptySlice([]string{"", "", ""})
	if len(result) != 0 {
		t.Errorf("expected 0, got %d", len(result))
	}
}

func Test_StringSlice_NonEmpty_NoneEmpty(t *testing.T) {
	result := stringslice.NonEmptySlice([]string{"a", "b"})
	if len(result) != 2 {
		t.Errorf("expected 2, got %d", len(result))
	}
}

func Test_StringSlice_NonEmpty_EmptySlice(t *testing.T) {
	result := stringslice.NonEmptySlice([]string{})
	if len(result) != 0 {
		t.Errorf("expected 0, got %d", len(result))
	}
}

// ==========================================
// NonWhitespace
// ==========================================

func Test_StringSlice_NonWhitespace_FiltersWhitespace(t *testing.T) {
	result := stringslice.NonWhitespace([]string{"a", "  ", "b", "\t", "c"})
	if len(result) != 3 {
		t.Errorf("expected 3, got %d", len(result))
	}
}

func Test_StringSlice_NonWhitespace_Nil(t *testing.T) {
	result := stringslice.NonWhitespace(nil)
	if result == nil || len(result) != 0 {
		t.Error("nil should return empty slice")
	}
}

func Test_StringSlice_NonWhitespace_Empty(t *testing.T) {
	result := stringslice.NonWhitespace([]string{})
	if len(result) != 0 {
		t.Errorf("expected 0, got %d", len(result))
	}
}

// ==========================================
// IsEmpty / HasAnyItem
// ==========================================

func Test_StringSlice_IsEmpty_True(t *testing.T) {
	if !stringslice.IsEmpty([]string{}) {
		t.Error("empty slice should be empty")
	}
}

func Test_StringSlice_IsEmpty_False(t *testing.T) {
	if stringslice.IsEmpty([]string{"a"}) {
		t.Error("non-empty slice should not be empty")
	}
}

func Test_StringSlice_HasAnyItem_True(t *testing.T) {
	if !stringslice.HasAnyItem([]string{"a"}) {
		t.Error("should have items")
	}
}

func Test_StringSlice_HasAnyItem_False(t *testing.T) {
	if stringslice.HasAnyItem([]string{}) {
		t.Error("empty should not have items")
	}
}

// ==========================================
// SortIf
// ==========================================

func Test_StringSlice_SortIf_True(t *testing.T) {
	s := []string{"c", "a", "b"}
	result := stringslice.SortIf(true, s)
	if result[0] != "a" || result[1] != "b" || result[2] != "c" {
		t.Errorf("expected sorted, got %v", result)
	}
}

func Test_StringSlice_SortIf_False(t *testing.T) {
	s := []string{"c", "a", "b"}
	result := stringslice.SortIf(false, s)
	if result[0] != "c" {
		t.Errorf("expected unsorted, got %v", result)
	}
}

// ==========================================
// SafeRangeItems
// ==========================================

func Test_StringSlice_SafeRangeItems_ValidRange(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e"}
	result := stringslice.SafeRangeItems(s, 1, 3)
	if len(result) != 2 {
		t.Errorf("expected 2, got %d", len(result))
	}
	if result[0] != "b" || result[1] != "c" {
		t.Errorf("expected [b c], got %v", result)
	}
}

func Test_StringSlice_SafeRangeItems_Nil(t *testing.T) {
	result := stringslice.SafeRangeItems(nil, 0, 1)
	if len(result) != 0 {
		t.Errorf("nil: expected 0, got %d", len(result))
	}
}

func Test_StringSlice_SafeRangeItems_Empty(t *testing.T) {
	result := stringslice.SafeRangeItems([]string{}, 0, 1)
	if len(result) != 0 {
		t.Errorf("empty: expected 0, got %d", len(result))
	}
}

func Test_StringSlice_SafeRangeItems_StartBeyondLength(t *testing.T) {
	result := stringslice.SafeRangeItems([]string{"a"}, 5, 10)
	if len(result) != 0 {
		t.Errorf("start beyond: expected 0, got %d", len(result))
	}
}

// ==========================================
// ExpandByFunc
// ==========================================

func Test_StringSlice_ExpandByFunc_Basic(t *testing.T) {
	result := stringslice.ExpandByFunc(
		[]string{"a,b", "c,d"},
		func(line string) []string {
			return []string{line + "-1", line + "-2"}
		},
	)
	if len(result) != 4 {
		t.Errorf("expected 4, got %d", len(result))
	}
}

func Test_StringSlice_ExpandByFunc_Empty(t *testing.T) {
	result := stringslice.ExpandByFunc(
		[]string{},
		func(line string) []string { return []string{line} },
	)
	if len(result) != 0 {
		t.Errorf("expected 0, got %d", len(result))
	}
}

func Test_StringSlice_ExpandByFunc_SkipsNilReturn(t *testing.T) {
	result := stringslice.ExpandByFunc(
		[]string{"a", "skip", "b"},
		func(line string) []string {
			if line == "skip" {
				return nil
			}
			return []string{line}
		},
	)
	if len(result) != 2 {
		t.Errorf("expected 2 (skip nil return), got %d", len(result))
	}
}
