package coreappendtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coreappend"
)

// ==========================================
// PrependAppendAnyItemsToStringsSkipOnNil
// ==========================================

func Test_PrependAppendToStrings_AllNonNil(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		"PRE", "POST", "a", "b",
	)
	if len(result) != 4 {
		t.Errorf("expected 4, got %d", len(result))
	}
	if result[0] != "PRE" || result[len(result)-1] != "POST" {
		t.Errorf("expected PRE...POST, got %v", result)
	}
}

func Test_PrependAppendToStrings_NilPrepend(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		nil, "POST", "a",
	)
	if len(result) != 2 {
		t.Errorf("expected 2 (skip nil prepend), got %d", len(result))
	}
	if result[0] != "a" || result[1] != "POST" {
		t.Errorf("unexpected result: %v", result)
	}
}

func Test_PrependAppendToStrings_NilAppend(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		"PRE", nil, "a",
	)
	if len(result) != 2 {
		t.Errorf("expected 2 (skip nil append), got %d", len(result))
	}
	if result[0] != "PRE" || result[1] != "a" {
		t.Errorf("unexpected result: %v", result)
	}
}

func Test_PrependAppendToStrings_BothNil(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		nil, nil, "a",
	)
	if len(result) != 1 {
		t.Errorf("expected 1, got %d", len(result))
	}
}

func Test_PrependAppendToStrings_NilInMiddle(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		"PRE", "POST", "a", nil, "b",
	)
	if len(result) != 4 {
		t.Errorf("expected 4 (skip nil middle), got %d", len(result))
	}
}

func Test_PrependAppendToStrings_NoItems(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		"PRE", "POST",
	)
	if len(result) != 2 {
		t.Errorf("expected 2 (just pre+post), got %d", len(result))
	}
}

func Test_PrependAppendToStrings_AllNil(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		nil, nil,
	)
	if len(result) != 0 {
		t.Errorf("expected 0, got %d", len(result))
	}
}

// ==========================================
// AppendAnyItemsToStringSkipOnNil
// ==========================================

func Test_AppendToString_Basic(t *testing.T) {
	result := coreappend.AppendAnyItemsToStringSkipOnNil(
		",", "SUFFIX", "a", "b",
	)
	if result != "a,b,SUFFIX" {
		t.Errorf("expected 'a,b,SUFFIX', got '%s'", result)
	}
}

func Test_AppendToString_NilAppend(t *testing.T) {
	result := coreappend.AppendAnyItemsToStringSkipOnNil(
		",", nil, "a",
	)
	if result != "a" {
		t.Errorf("expected 'a', got '%s'", result)
	}
}

// ==========================================
// PrependAnyItemsToStringSkipOnNil
// ==========================================

func Test_PrependToString_Basic(t *testing.T) {
	result := coreappend.PrependAnyItemsToStringSkipOnNil(
		",", "PREFIX", "a", "b",
	)
	if result != "PREFIX,a,b" {
		t.Errorf("expected 'PREFIX,a,b', got '%s'", result)
	}
}

func Test_PrependToString_NilPrepend(t *testing.T) {
	result := coreappend.PrependAnyItemsToStringSkipOnNil(
		",", nil, "a",
	)
	if result != "a" {
		t.Errorf("expected 'a', got '%s'", result)
	}
}

// ==========================================
// PrependAppendAnyItemsToStringSkipOnNil (joined)
// ==========================================

func Test_PrependAppendToString_Joined(t *testing.T) {
	result := coreappend.PrependAppendAnyItemsToStringSkipOnNil(
		"-", "PRE", "POST", "mid",
	)
	if result != "PRE-mid-POST" {
		t.Errorf("expected 'PRE-mid-POST', got '%s'", result)
	}
}

// ==========================================
// PrependAppendAnyItemsToStringsUsingFunc
// ==========================================

func Test_PrependAppendUsingFunc_Basic(t *testing.T) {
	compiler := func(item any) string {
		return fmt.Sprintf("[%v]", item)
	}
	result := coreappend.PrependAppendAnyItemsToStringsUsingFunc(
		false, compiler, "pre", "post", "a", "b",
	)
	if len(result) != 4 {
		t.Errorf("expected 4, got %d", len(result))
	}
	if result[0] != "[pre]" || result[3] != "[post]" {
		t.Errorf("unexpected: %v", result)
	}
}

func Test_PrependAppendUsingFunc_SkipEmpty(t *testing.T) {
	compiler := func(item any) string {
		if item == nil {
			return ""
		}
		return fmt.Sprintf("%v", item)
	}
	result := coreappend.PrependAppendAnyItemsToStringsUsingFunc(
		true, compiler, nil, nil, "a", nil, "b",
	)
	// prepend=nil→"" skipped, append=nil→"" skipped, nil middle skipped
	if len(result) != 2 {
		t.Errorf("expected 2 (skip empties), got %d: %v", len(result), result)
	}
}

func Test_PrependAppendUsingFunc_NoSkipEmpty(t *testing.T) {
	compiler := func(item any) string {
		if item == nil {
			return ""
		}
		return fmt.Sprintf("%v", item)
	}
	result := coreappend.PrependAppendAnyItemsToStringsUsingFunc(
		false, compiler, nil, nil, "a",
	)
	// prepend="" included, append="" included, nil middle skipped
	if len(result) != 3 {
		t.Errorf("expected 3 (include empties), got %d: %v", len(result), result)
	}
}

// ==========================================
// MapStringStringAppendMapStringToAnyItems
// ==========================================

func Test_MapAppend_Basic(t *testing.T) {
	mainMap := map[string]string{"a": "1"}
	appendMap := map[string]any{"b": 2, "c": "three"}
	result := coreappend.MapStringStringAppendMapStringToAnyItems(false, mainMap, appendMap)
	if len(result) != 3 {
		t.Errorf("expected 3, got %d", len(result))
	}
}

func Test_MapAppend_EmptyAppend(t *testing.T) {
	mainMap := map[string]string{"a": "1"}
	result := coreappend.MapStringStringAppendMapStringToAnyItems(false, mainMap, map[string]any{})
	if len(result) != 1 {
		t.Errorf("expected 1, got %d", len(result))
	}
}

func Test_MapAppend_SkipEmpty(t *testing.T) {
	mainMap := map[string]string{}
	appendMap := map[string]any{"a": "", "b": "val"}
	result := coreappend.MapStringStringAppendMapStringToAnyItems(true, mainMap, appendMap)
	// "a" has value "" which after Sprintf becomes "" → skipped
	if _, has := result["a"]; has {
		t.Error("SkipEmpty should skip empty string values")
	}
	if result["b"] != "val" {
		t.Errorf("expected 'val', got '%s'", result["b"])
	}
}

func Test_MapAppend_OverwriteExisting(t *testing.T) {
	mainMap := map[string]string{"k": "old"}
	appendMap := map[string]any{"k": "new"}
	result := coreappend.MapStringStringAppendMapStringToAnyItems(false, mainMap, appendMap)
	if result["k"] != "new" {
		t.Errorf("expected overwrite to 'new', got '%s'", result["k"])
	}
}
