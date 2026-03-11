package coreindexestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreindexes"
)

// TestHasIndex verifies index existence.
func TestHasIndex(t *testing.T) {
	indexes := []int{1, 3, 5}
	if !coreindexes.HasIndex(indexes, 3) {
		t.Error("3 should be found")
	}
	if coreindexes.HasIndex(indexes, 2) {
		t.Error("2 should not be found")
	}
}

// TestIsInvalidIndex verifies invalid index.
func TestIsInvalidIndex(t *testing.T) {
	if !coreindexes.IsInvalidIndex(-1) {
		t.Error("-1 should be invalid")
	}
	if coreindexes.IsInvalidIndex(0) {
		t.Error("0 should be valid")
	}
}

// TestIsWithinIndexRange verifies index range.
func TestIsWithinIndexRange(t *testing.T) {
	if !coreindexes.IsWithinIndexRange(2, 5) {
		t.Error("index 2 should be within length 5")
	}
	if coreindexes.IsWithinIndexRange(5, 5) {
		t.Error("index 5 should not be within length 5")
	}
}

// TestLastIndex verifies last index.
func TestLastIndex(t *testing.T) {
	if coreindexes.LastIndex(5) != 4 {
		t.Error("last index of 5 should be 4")
	}
	if coreindexes.LastIndex(0) != -1 {
		t.Error("last index of 0 should be -1")
	}
}

// TestNameByIndex verifies name lookup.
func TestNameByIndex(t *testing.T) {
	if coreindexes.NameByIndex(0) != "First" {
		t.Errorf("expected 'First', got '%s'", coreindexes.NameByIndex(0))
	}
	if coreindexes.NameByIndex(9) != "Tenth" {
		t.Errorf("expected 'Tenth', got '%s'", coreindexes.NameByIndex(9))
	}
	if coreindexes.NameByIndex(99) != "" {
		t.Error("out of range should return empty")
	}
}

// TestOf verifies index-of search.
func TestOf(t *testing.T) {
	indexes := []int{10, 20, 30}
	if coreindexes.Of(indexes, 20) != 1 {
		t.Error("20 should be at position 1")
	}
	if coreindexes.Of(indexes, 99) != -1 {
		t.Error("99 should return -1")
	}
}

// TestSafeEndingIndex verifies safe ending.
func TestSafeEndingIndex(t *testing.T) {
	if coreindexes.SafeEndingIndex(5, 3) != 3 {
		t.Error("within range should return lastTaking")
	}
	if coreindexes.SafeEndingIndex(3, 5) != 2 {
		t.Error("exceeding should return lastIndex")
	}
}

// TestHasIndexPlusRemoveIndex verifies find-and-remove.
func TestHasIndexPlusRemoveIndex(t *testing.T) {
	indexes := []int{1, 2, 3}
	if !coreindexes.HasIndexPlusRemoveIndex(&indexes, 2) {
		t.Error("2 should be found")
	}
	if len(indexes) != 2 {
		t.Errorf("expected 2 items, got %d", len(indexes))
	}
	if coreindexes.HasIndexPlusRemoveIndex(&indexes, 99) {
		t.Error("99 should not be found")
	}
}

// TestConstants verifies index constants.
func TestConstants(t *testing.T) {
	if coreindexes.First != 0 { t.Error("First should be 0") }
	if coreindexes.Second != 1 { t.Error("Second should be 1") }
	if coreindexes.Tenth != 9 { t.Error("Tenth should be 9") }
	if coreindexes.Index0 != 0 { t.Error("Index0 should be 0") }
	if coreindexes.Index20 != 20 { t.Error("Index20 should be 20") }
}
