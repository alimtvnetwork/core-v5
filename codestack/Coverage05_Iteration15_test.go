package codestack

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/corejson"
)

// ══════════════════════════════════════════════════════════════════════════════
// FileWithLine — ParseInjectUsingJsonMust panic branch (line 92)
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_FileWithLine_ParseInjectUsingJsonMust_Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	badJson := corejson.NewPtr("{invalid")
	fw := &FileWithLine{}
	fw.ParseInjectUsingJsonMust(badJson)
}

// ══════════════════════════════════════════════════════════════════════════════
// Trace — ParseInjectUsingJson error + Must panic (lines 181, 197)
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_Trace_ParseInjectUsingJson_Error(t *testing.T) {
	badJson := corejson.NewPtr("{invalid")
	tr := &Trace{}
	_, err := tr.ParseInjectUsingJson(badJson)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I15_Trace_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	badJson := corejson.NewPtr("{invalid")
	tr := &Trace{}
	tr.ParseInjectUsingJsonMust(badJson)
}

// ══════════════════════════════════════════════════════════════════════════════
// TraceCollection — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_TraceCollection_AddsUsingSkip_BreakOnceInvalid(t *testing.T) {
	// isSkip && isBreakOnceInvalid => return it (line 73-75)
	tc := New.traces.Default()
	// Use a very high skip index to trigger invalid traces
	tc.AddsUsingSkip(true, true, 9999, 5)
}

func Test_I15_TraceCollection_AddsUsingSkipUsingFilter_BreakOnFilter(t *testing.T) {
	tc := New.traces.Default()
	count := 0
	tc.AddsUsingSkipUsingFilter(true, false, 0, DefaultStackCount, func(tr *Trace) (bool, bool) {
		count++
		if count >= 2 {
			return true, true // take and break
		}
		return true, false
	})
}

func Test_I15_TraceCollection_AddsUsingSkipUsingFilter_SkipInvalid(t *testing.T) {
	tc := New.traces.Default()
	// isSkip && isBreakOnceInvalid (line 117-119)
	tc.AddsUsingSkipUsingFilter(true, true, 9999, 5, func(tr *Trace) (bool, bool) {
		return true, false
	})
}

func Test_I15_TraceCollection_PagedItems_NegativeIndex(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for negative page index")
		}
	}()
	tc := New.traces.Default()
	tr := New.Create(0)
	tc.Adds(tr)
	tc.PagedItems(-1, 10)
}

func Test_I15_TraceCollection_FilterWithLimit_BreakBranch(t *testing.T) {
	tc := New.traces.Default()
	tr1 := New.Create(0)
	tr2 := New.Create(0)
	tc.Adds(tr1, tr2)

	result := tc.FilterWithLimit(10, func(tr *Trace) (bool, bool) {
		return true, true // take first and break
	})
	if len(result) != 1 {
		t.Fatalf("expected 1, got %d", len(result))
	}
}

func Test_I15_TraceCollection_IsEqualItems_NilIt(t *testing.T) {
	var tc *TraceCollection
	// nil == nil (line 810)
	if !tc.IsEqualItems() {
		t.Fatal("expected nil == nil to be true")
	}
}

func Test_I15_TraceCollection_IsEqualItems_NilVsNonNil(t *testing.T) {
	var tc *TraceCollection
	tr := New.Create(0)
	// nil vs non-nil (line 814)
	if tc.IsEqualItems(tr) {
		t.Fatal("expected nil != non-nil")
	}
}

func Test_I15_TraceCollection_ParseInjectUsingJson_Error(t *testing.T) {
	badJson := corejson.NewPtr("{invalid")
	tc := New.traces.Default()
	_, err := tc.ParseInjectUsingJson(badJson)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I15_TraceCollection_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	badJson := corejson.NewPtr("{invalid")
	tc := New.traces.Default()
	tc.ParseInjectUsingJsonMust(badJson)
}

// ══════════════════════════════════════════════════════════════════════════════
// dirGetter — runtime.Caller fails branch (lines 22, 49)
// These are hard to trigger since runtime.Caller rarely fails at valid skip.
// We at least exercise normal paths to ensure coverage of the "isOkay" branch.
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_DirGetter_Get_HighSkip(t *testing.T) {
	result := Dir.Get(99999)
	if result != "" {
		t.Fatal("expected empty for unreachable skip")
	}
}

func Test_I15_DirGetter_CurDir(t *testing.T) {
	result := Dir.CurDir()
	if result == "" {
		t.Fatal("expected non-empty")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// fileGetter — runtime.Caller fails branches (lines 15, 31, 74)
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_FileGetter_Name_HighSkip(t *testing.T) {
	result := File.Name(99999)
	// When runtime.Caller fails, isOkay=false AND file="" → return empty
	if result != "" {
		t.Fatal("expected empty for unreachable skip")
	}
}

func Test_I15_FileGetter_Path_HighSkip(t *testing.T) {
	result := File.Path(99999)
	if result != "" {
		t.Fatal("expected empty for unreachable skip")
	}
}

func Test_I15_FileGetter_CurrentFilePath(t *testing.T) {
	result := File.CurrentFilePath()
	if result == "" {
		t.Fatal("expected non-empty")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// newTraceCollection — Default, Using clone branch (lines 13, 21, 33, 38)
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_NewTraceCollection_Default(t *testing.T) {
	tc := New.traces.Default()
	if tc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I15_NewTraceCollection_Using_NilTraces(t *testing.T) {
	tc := New.traces.Using(false)
	if tc == nil || tc.Length() != 0 {
		t.Fatal("expected empty collection")
	}
}

func Test_I15_NewTraceCollection_Using_Clone(t *testing.T) {
	tr := New.Create(0)
	tc := New.traces.Using(true, tr)
	if tc == nil || tc.Length() != 1 {
		t.Fatal("expected 1 trace")
	}
}

func Test_I15_NewTraceCollection_Empty(t *testing.T) {
	tc := New.traces.Empty()
	if tc == nil {
		t.Fatal("expected non-nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Verify JSON roundtrip for Trace
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_Trace_JsonRoundtrip(t *testing.T) {
	tr := New.Create(0)
	b, err := json.Marshal(tr)
	if err != nil {
		t.Fatal(err)
	}
	var tr2 Trace
	if err := json.Unmarshal(b, &tr2); err != nil {
		t.Fatal(err)
	}
}
