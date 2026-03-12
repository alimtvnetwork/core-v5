package mapdiffinternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/internal/mapdiffinternal"
)

// Cover: HashmapDiff nil receiver, DiffRaw nil branches, MapStringAnyDiff nil/empty branches

func Test_HashmapDiff_NilPtr_Length_Cov2(t *testing.T) {
	var h *mapdiffinternal.HashmapDiff
	if h.Length() != 0 {
		t.Error("nil should have length 0")
	}
}

func Test_HashmapDiff_Empty_Cov2(t *testing.T) {
	h := mapdiffinternal.HashmapDiff{}
	if !h.IsEmpty() {
		t.Error("should be empty")
	}
	if h.HasAnyItem() {
		t.Error("should not have items")
	}
	if h.LastIndex() != -1 {
		t.Error("expected -1")
	}
	keys := h.AllKeysSorted()
	if len(keys) != 0 {
		t.Error("expected 0 keys")
	}
}

func Test_HashmapDiff_NonEmpty_Cov2(t *testing.T) {
	h := mapdiffinternal.HashmapDiff{"a": "1", "b": "2"}
	if h.IsEmpty() {
		t.Error("should not be empty")
	}
	if !h.HasAnyItem() {
		t.Error("should have items")
	}
	if h.LastIndex() != 1 {
		t.Error("expected 1")
	}
	keys := h.AllKeysSorted()
	if len(keys) != 2 || keys[0] != "a" {
		t.Error("expected sorted keys")
	}
}

func Test_HashmapDiff_IsRawEqual_BothNil_Cov2(t *testing.T) {
	var h *mapdiffinternal.HashmapDiff
	if !h.IsRawEqual(nil) {
		t.Error("both nil should be equal")
	}
}

func Test_HashmapDiff_IsRawEqual_LeftNil_Cov2(t *testing.T) {
	var h *mapdiffinternal.HashmapDiff
	if h.IsRawEqual(map[string]string{"a": "1"}) {
		t.Error("left nil should not equal non-nil")
	}
}

func Test_HashmapDiff_IsRawEqual_DiffLength_Cov2(t *testing.T) {
	h := mapdiffinternal.HashmapDiff{"a": "1"}
	if h.IsRawEqual(map[string]string{"a": "1", "b": "2"}) {
		t.Error("different lengths should not be equal")
	}
}

func Test_HashmapDiff_IsRawEqual_MissingKey_Cov2(t *testing.T) {
	h := mapdiffinternal.HashmapDiff{"a": "1"}
	if h.IsRawEqual(map[string]string{"b": "1"}) {
		t.Error("missing key should not be equal")
	}
}

func Test_HashmapDiff_IsRawEqual_DiffValue_Cov2(t *testing.T) {
	h := mapdiffinternal.HashmapDiff{"a": "1"}
	if h.IsRawEqual(map[string]string{"a": "2"}) {
		t.Error("different values should not be equal")
	}
}

func Test_HashmapDiff_IsRawEqual_Equal_Cov2(t *testing.T) {
	h := mapdiffinternal.HashmapDiff{"a": "1"}
	if !h.IsRawEqual(map[string]string{"a": "1"}) {
		t.Error("should be equal")
	}
}

func Test_HashmapDiff_DiffRaw_BothNil_Cov2(t *testing.T) {
	var h *mapdiffinternal.HashmapDiff
	r := h.DiffRaw(nil)
	if len(r) != 0 {
		t.Error("expected empty")
	}
}

func Test_HashmapDiff_DiffRaw_LeftNilRightNot_Cov2(t *testing.T) {
	var h *mapdiffinternal.HashmapDiff
	right := map[string]string{"a": "1"}
	r := h.DiffRaw(right)
	if len(r) != 1 {
		t.Error("expected right map")
	}
}

func Test_HashmapDiff_DiffRaw_RightNil_Cov2(t *testing.T) {
	h := mapdiffinternal.HashmapDiff{"a": "1"}
	r := h.DiffRaw(nil)
	if len(r) != 1 {
		t.Error("expected left map")
	}
}

func Test_HashmapDiff_DiffRaw_DiffValues_Cov2(t *testing.T) {
	h := mapdiffinternal.HashmapDiff{"a": "1", "b": "2"}
	r := h.DiffRaw(map[string]string{"a": "1", "b": "3"})
	if len(r) != 1 || r["b"] != "2" {
		t.Error("expected diff on b")
	}
}

func Test_HashmapDiff_DiffRaw_RightExtra_Cov2(t *testing.T) {
	h := mapdiffinternal.HashmapDiff{"a": "1"}
	r := h.DiffRaw(map[string]string{"a": "1", "c": "3"})
	if len(r) != 1 || r["c"] != "3" {
		t.Error("expected c from right")
	}
}

func Test_HashmapDiff_DiffJsonMessage_NoDiff_Cov2(t *testing.T) {
	h := mapdiffinternal.HashmapDiff{"a": "1"}
	r := h.DiffJsonMessage(map[string]string{"a": "1"})
	if r != "" {
		t.Error("expected empty")
	}
}

func Test_HashmapDiff_DiffJsonMessage_WithDiff_Cov2(t *testing.T) {
	h := mapdiffinternal.HashmapDiff{"a": "1"}
	r := h.DiffJsonMessage(map[string]string{"a": "2"})
	if r == "" {
		t.Error("expected non-empty")
	}
}

func Test_HashmapDiff_ShouldDiffMessage_NoDiff_Cov2(t *testing.T) {
	h := mapdiffinternal.HashmapDiff{"a": "1"}
	r := h.ShouldDiffMessage("test", map[string]string{"a": "1"})
	if r != "" {
		t.Error("expected empty")
	}
}

func Test_HashmapDiff_ShouldDiffMessage_WithDiff_Cov2(t *testing.T) {
	h := mapdiffinternal.HashmapDiff{"a": "1"}
	r := h.ShouldDiffMessage("test", map[string]string{"a": "2"})
	if r == "" {
		t.Error("expected non-empty")
	}
}

func Test_HashmapDiff_LogShouldDiffMessage_NoDiff_Cov2(t *testing.T) {
	h := mapdiffinternal.HashmapDiff{"a": "1"}
	r := h.LogShouldDiffMessage("test", map[string]string{"a": "1"})
	if r != "" {
		t.Error("expected empty")
	}
}

// ============================================================================
// MapStringAnyDiff
// ============================================================================

func Test_MapStringAnyDiff_NilPtr_Length_Cov2(t *testing.T) {
	var m *mapdiffinternal.MapStringAnyDiff
	if m.Length() != 0 {
		t.Error("nil should be 0")
	}
}

func Test_MapStringAnyDiff_Empty_Cov2(t *testing.T) {
	m := mapdiffinternal.MapStringAnyDiff{}
	if !m.IsEmpty() {
		t.Error("should be empty")
	}
	if m.HasAnyItem() {
		t.Error("should not have items")
	}
	if m.LastIndex() != -1 {
		t.Error("expected -1")
	}
}

func Test_MapStringAnyDiff_Raw_Nil_Cov2(t *testing.T) {
	var m mapdiffinternal.MapStringAnyDiff
	r := m.Raw()
	if r == nil {
		t.Error("expected empty map not nil")
	}
}

func Test_MapStringAnyDiff_Raw_NonNil_Cov2(t *testing.T) {
	m := mapdiffinternal.MapStringAnyDiff{"a": 1}
	r := m.Raw()
	if len(r) != 1 {
		t.Error("expected 1")
	}
}

func Test_MapStringAnyDiff_HasAnyChanges_Cov2(t *testing.T) {
	m := mapdiffinternal.MapStringAnyDiff{"a": 1}
	if !m.HasAnyChanges(false, map[string]any{"a": 2}) {
		t.Error("should have changes")
	}
}

func Test_MapStringAnyDiff_IsRawEqual_RegardlessType_Cov2(t *testing.T) {
	m := mapdiffinternal.MapStringAnyDiff{"a": 1}
	if !m.IsRawEqual(true, map[string]any{"a": 1}) {
		t.Error("should be equal regardless type")
	}
}

func Test_MapStringAnyDiff_IsRawEqual_StrictType_Cov2(t *testing.T) {
	m := mapdiffinternal.MapStringAnyDiff{"a": 1}
	if !m.IsRawEqual(false, map[string]any{"a": 1}) {
		t.Error("should be equal")
	}
}

func Test_MapStringAnyDiff_DiffRaw_BothNil_Cov2(t *testing.T) {
	var m *mapdiffinternal.MapStringAnyDiff
	r := m.DiffRaw(false, nil)
	if len(r) != 0 {
		t.Error("expected empty")
	}
}

func Test_MapStringAnyDiff_DiffRaw_LeftNil_Cov2(t *testing.T) {
	var m *mapdiffinternal.MapStringAnyDiff
	r := m.DiffRaw(false, map[string]any{"a": 1})
	if len(r) != 1 {
		t.Error("expected right map")
	}
}

func Test_MapStringAnyDiff_DiffRaw_RightNil_Cov2(t *testing.T) {
	m := mapdiffinternal.MapStringAnyDiff{"a": 1}
	r := m.DiffRaw(false, nil)
	if len(r) != 1 {
		t.Error("expected left map")
	}
}

func Test_MapStringAnyDiff_ToStringsSliceOfDiffMap_StringVal_Cov2(t *testing.T) {
	m := mapdiffinternal.MapStringAnyDiff{}
	diff := map[string]any{"key": "stringval"}
	r := m.ToStringsSliceOfDiffMap(diff)
	if len(r) != 1 {
		t.Error("expected 1")
	}
}

func Test_MapStringAnyDiff_ToStringsSliceOfDiffMap_NonStringVal_Cov2(t *testing.T) {
	m := mapdiffinternal.MapStringAnyDiff{}
	diff := map[string]any{"key": 42}
	r := m.ToStringsSliceOfDiffMap(diff)
	if len(r) != 1 {
		t.Error("expected 1")
	}
}

func Test_MapStringAnyDiff_DiffJsonMessage_NoDiff_Cov2(t *testing.T) {
	m := mapdiffinternal.MapStringAnyDiff{"a": 1}
	r := m.DiffJsonMessage(false, map[string]any{"a": 1})
	if r != "" {
		t.Error("expected empty")
	}
}

func Test_MapStringAnyDiff_DiffJsonMessage_WithDiff_Cov2(t *testing.T) {
	m := mapdiffinternal.MapStringAnyDiff{"a": 1}
	r := m.DiffJsonMessage(false, map[string]any{"a": 2})
	if r == "" {
		t.Error("expected non-empty")
	}
}

func Test_MapStringAnyDiff_ShouldDiffMessage_Cov2(t *testing.T) {
	m := mapdiffinternal.MapStringAnyDiff{"a": 1}
	r := m.ShouldDiffMessage(false, "test", map[string]any{"a": 1})
	if r != "" {
		t.Error("expected empty for no diff")
	}
	r = m.ShouldDiffMessage(false, "test", map[string]any{"a": 2})
	if r == "" {
		t.Error("expected non-empty for diff")
	}
}

func Test_MapStringAnyDiff_LogShouldDiffMessage_NoDiff_Cov2(t *testing.T) {
	m := mapdiffinternal.MapStringAnyDiff{"a": 1}
	r := m.LogShouldDiffMessage(false, "test", map[string]any{"a": 1})
	if r != "" {
		t.Error("expected empty")
	}
}
