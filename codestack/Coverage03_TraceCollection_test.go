package codestack

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

func TestTraceCollection_BasicOps(t *testing.T) {
	tc := New.traces.Empty()
	if !tc.IsEmpty() {
		t.Fatal("should be empty")
	}
	if tc.HasAnyItem() {
		t.Fatal("should not have items")
	}
	if tc.Length() != 0 {
		t.Fatal("expected 0")
	}
	if tc.Count() != 0 {
		t.Fatal("expected 0")
	}
	if tc.LastIndex() != -1 {
		t.Fatal("expected -1")
	}
}

func TestTraceCollection_Add(t *testing.T) {
	tc := New.traces.Empty()
	trace := New.Create(0)
	tc.Add(trace)
	if tc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestTraceCollection_Adds(t *testing.T) {
	tc := New.traces.Empty()
	t1 := New.Create(0)
	t2 := New.Create(0)
	tc.Adds(t1, t2)
	if tc.Length() != 2 {
		t.Fatal("expected 2")
	}

	// empty adds
	tc.Adds()
	if tc.Length() != 2 {
		t.Fatal("expected 2 still")
	}
}

func TestTraceCollection_AddsPtr(t *testing.T) {
	tc := New.traces.Empty()
	p := New.Ptr(0)
	tc.AddsPtr(false, p, nil)
	if tc.Length() != 1 {
		t.Fatal("expected 1")
	}

	tc2 := New.traces.Empty()
	tc2.AddsPtr(true, p, nil)
	if tc2.Length() != 1 {
		t.Fatal("expected 1")
	}

	tc3 := New.traces.Empty()
	tc3.AddsPtr(false)
	if tc3.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func TestTraceCollection_AddsIf(t *testing.T) {
	tc := New.traces.Empty()
	tr := New.Create(0)
	tc.AddsIf(true, tr)
	if tc.Length() != 1 {
		t.Fatal("expected 1")
	}
	tc.AddsIf(false, tr)
	if tc.Length() != 1 {
		t.Fatal("expected 1 still")
	}
}

func TestTraceCollection_FirstLast(t *testing.T) {
	tc := New.traces.Empty()
	if tc.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
	if tc.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}
	_ = tc.FirstOrDefaultDynamic()
	_ = tc.LastOrDefaultDynamic()

	tr := New.Create(0)
	tc.Add(tr)
	tc.Add(New.Create(0))

	first := tc.First()
	_ = first
	last := tc.Last()
	_ = last
	_ = tc.FirstDynamic()
	_ = tc.LastDynamic()
	if tc.FirstOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
	if tc.LastOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
}

func TestTraceCollection_SkipTakeLimit(t *testing.T) {
	tc := New.traces.Empty()
	for i := 0; i < 5; i++ {
		tc.Add(New.Create(0))
	}

	s := tc.Skip(2)
	if len(s) != 3 {
		t.Fatal("expected 3")
	}
	_ = tc.SkipDynamic(2)
	_ = tc.SkipDynamic(10) // over length
	_ = tc.SkipCollection(1)

	tk := tc.Take(3)
	if len(tk) != 3 {
		t.Fatal("expected 3")
	}
	_ = tc.TakeDynamic(2)
	_ = tc.TakeCollection(2)

	_ = tc.Limit(3)
	_ = tc.LimitCollection(3)
	_ = tc.LimitDynamic(2)
	_ = tc.SafeLimitCollection(3)
	_ = tc.SafeLimitCollection(100) // over length
}

func TestTraceCollection_HasIndex(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	if !tc.HasIndex(0) {
		t.Fatal("should have index 0")
	}
	if tc.HasIndex(1) {
		t.Fatal("should not have index 1")
	}
}

func TestTraceCollection_Strings(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	s := tc.Strings()
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func TestTraceCollection_Filter(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	tc.Add(New.Create(0))

	filtered := tc.Filter(func(tr *Trace) (bool, bool) {
		return true, false
	})
	if len(filtered) != 2 {
		t.Fatal("expected 2")
	}

	// with break
	filtered2 := tc.Filter(func(tr *Trace) (bool, bool) {
		return true, true
	})
	if len(filtered2) != 1 {
		t.Fatal("expected 1")
	}
}

func TestTraceCollection_FilterWithLimit(t *testing.T) {
	tc := New.traces.Empty()
	for i := 0; i < 10; i++ {
		tc.Add(New.Create(0))
	}
	filtered := tc.FilterWithLimit(3, func(tr *Trace) (bool, bool) {
		return true, false
	})
	if len(filtered) != 3 {
		t.Fatal("expected 3")
	}
}

func TestTraceCollection_FilterTraceCollection(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	fc := tc.FilterTraceCollection(func(tr *Trace) (bool, bool) {
		return true, false
	})
	if fc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestTraceCollection_FilterByNames(t *testing.T) {
	tc := New.traces.Empty()
	tr := New.Create(0)
	tc.Add(tr)

	_ = tc.FilterPackageNameTraceCollection(tr.PackageName)
	_ = tc.SkipFilterPackageNameTraceCollection("nonexistent")
	_ = tc.FilterMethodNameTraceCollection(tr.MethodName)
	_ = tc.SkipFilterMethodNameTraceCollection("nonexistent")
	_ = tc.FilterFullMethodNameTraceCollection(tr.PackageMethodName)
	_ = tc.SkipFilterFullMethodNameTraceCollection("nonexistent")
	_ = tc.SkipFilterFilenameTraceCollection("nonexistent")
}

func TestTraceCollection_FileWithLines(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	fwl := tc.FileWithLines()
	if len(fwl) != 1 {
		t.Fatal("expected 1")
	}
	fwls := tc.FileWithLinesStrings()
	if len(fwls) != 1 {
		t.Fatal("expected 1")
	}
}

func TestTraceCollection_StringsUsingFmt(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	s := tc.StringsUsingFmt(func(tr *Trace) string {
		return tr.PackageName
	})
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func TestTraceCollection_JoinUsingFmt(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	s := tc.JoinUsingFmt(func(tr *Trace) string {
		return tr.PackageName
	}, ",")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestTraceCollection_ShortStrings(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	ss := tc.ShortStrings()
	if len(ss) != 1 {
		t.Fatal("expected 1")
	}
}

func TestTraceCollection_JoinShortStrings(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	s := tc.JoinShortStrings(",")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestTraceCollection_Reverse(t *testing.T) {
	tc := New.traces.Empty()
	tc.Reverse() // empty

	tc.Add(New.Create(0))
	tc.Reverse() // single

	tc.Add(New.Create(0))
	tc.Reverse() // two

	tc.Add(New.Create(0))
	tc.Reverse() // three
}

func TestTraceCollection_JsonStrings(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	js := tc.JsonStrings()
	if len(js) != 1 {
		t.Fatal("expected 1")
	}
}

func TestTraceCollection_Joins(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	_ = tc.JoinFileWithLinesStrings(",")
	_ = tc.JoinJsonStrings(",")
	_ = tc.Join(",")
	_ = tc.JoinLines()
	_ = tc.JoinCsv()
	_ = tc.JoinCsvLine()
}

func TestTraceCollection_CodeStacksString(t *testing.T) {
	tc := New.traces.Empty()
	if tc.CodeStacksString() != "" {
		t.Fatal("expected empty")
	}
	tc.Add(New.Create(0))
	s := tc.CodeStacksString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestTraceCollection_FileWithLinesString(t *testing.T) {
	tc := New.traces.Empty()
	if tc.FileWithLinesString() != "" {
		t.Fatal("expected empty")
	}
	tc.Add(New.Create(0))
	if tc.FileWithLinesString() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestTraceCollection_CodeStacksStringLimit(t *testing.T) {
	tc := New.traces.Empty()
	if tc.CodeStacksStringLimit(5) != "" {
		t.Fatal("expected empty")
	}
	tc.Add(New.Create(0))
	s := tc.CodeStacksStringLimit(5)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestTraceCollection_IsEqual(t *testing.T) {
	tc1 := New.traces.Empty()
	tc2 := New.traces.Empty()
	if !tc1.IsEqual(tc2) {
		t.Fatal("should be equal")
	}

	var nilTC *TraceCollection
	if !nilTC.IsEqual(nil) {
		t.Fatal("both nil should be equal")
	}
	if nilTC.IsEqual(tc1) {
		t.Fatal("nil vs non-nil should not be equal")
	}
	if tc1.IsEqual(nil) {
		t.Fatal("non-nil vs nil should not be equal")
	}
}

func TestTraceCollection_IsEqualItems(t *testing.T) {
	tc := New.traces.Empty()
	tr := New.Create(0)
	tc.Add(tr)
	if !tc.IsEqualItems(tr) {
		t.Fatal("should be equal")
	}

	var nilTC *TraceCollection
	if !nilTC.IsEqualItems() {
		t.Fatal("nil vs nil should be equal")
	}
}

func TestTraceCollection_JsonString(t *testing.T) {
	tc := TraceCollection{}
	if tc.JsonString() != "" {
		t.Fatal("expected empty for empty collection")
	}
	tc.Items = append(tc.Items, New.Create(0))
	s := tc.JsonString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestTraceCollection_String(t *testing.T) {
	tc := TraceCollection{}
	if tc.String() != "" {
		t.Fatal("expected empty")
	}
	tc.Items = append(tc.Items, New.Create(0))
	if tc.String() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestTraceCollection_CsvStrings(t *testing.T) {
	tc := TraceCollection{}
	if len(tc.CsvStrings()) != 0 {
		t.Fatal("expected empty")
	}
	tc.Items = append(tc.Items, New.Create(0))
	if len(tc.CsvStrings()) != 1 {
		t.Fatal("expected 1")
	}
}

func TestTraceCollection_Json(t *testing.T) {
	tc := TraceCollection{Items: []Trace{New.Create(0)}}
	j := tc.Json()
	if j.HasError() {
		t.Fatal(j.Error)
	}
	jp := tc.JsonPtr()
	if jp == nil {
		t.Fatal("expected non-nil")
	}
}

func TestTraceCollection_JsonModel(t *testing.T) {
	tc := TraceCollection{Items: []Trace{New.Create(0)}}
	_ = tc.JsonModel()
	_ = tc.JsonModelAny()
}

func TestTraceCollection_Serializer(t *testing.T) {
	tc := TraceCollection{Items: []Trace{New.Create(0)}}
	b, err := tc.Serializer()
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func TestTraceCollection_StackTracesBytes(t *testing.T) {
	tc := TraceCollection{}
	if len(tc.StackTracesBytes()) != 0 {
		t.Fatal("expected empty")
	}
	tc.Items = append(tc.Items, New.Create(0))
	if len(tc.StackTracesBytes()) == 0 {
		t.Fatal("expected non-empty")
	}
}

func TestTraceCollection_ParseInjectUsingJson(t *testing.T) {
	tc := TraceCollection{Items: []Trace{New.Create(0)}}
	jr := corejson.NewPtr(tc)
	target := New.traces.Empty()
	_, err := target.ParseInjectUsingJson(jr)
	if err != nil {
		t.Fatal(err)
	}

	badJr := corejson.NewResult.UsingBytes([]byte("invalid"))
	_, err2 := target.ParseInjectUsingJson(badJr.Ptr())
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func TestTraceCollection_ParseInjectUsingJsonMust(t *testing.T) {
	tc := TraceCollection{Items: []Trace{New.Create(0)}}
	jr := corejson.NewPtr(tc)
	target := New.traces.Empty()
	result := target.ParseInjectUsingJsonMust(jr)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func TestTraceCollection_JsonParseSelfInject(t *testing.T) {
	tc := TraceCollection{Items: []Trace{New.Create(0)}}
	jr := corejson.NewPtr(tc)
	target := New.traces.Empty()
	err := target.JsonParseSelfInject(jr)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTraceCollection_AsInterfaces(t *testing.T) {
	tc := TraceCollection{}
	_ = tc.AsJsonContractsBinder()
	_ = tc.AsJsoner()
	_ = tc.AsJsonParseSelfInjector()
}

func TestTraceCollection_ClearDispose(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	tc.Clear()

	tc2 := New.traces.Empty()
	tc2.Add(New.Create(0))
	tc2.Dispose()

	var nilTC *TraceCollection
	nilTC.Dispose()
}

func TestTraceCollection_Clone(t *testing.T) {
	tc := TraceCollection{Items: []Trace{New.Create(0)}}
	cl := tc.Clone()
	if cl.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestTraceCollection_ClonePtr(t *testing.T) {
	tc := &TraceCollection{Items: []Trace{New.Create(0)}}
	cp := tc.ClonePtr()
	if cp == nil {
		t.Fatal("expected non-nil")
	}

	var nilTC *TraceCollection
	if nilTC.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
}

func TestTraceCollection_Paging(t *testing.T) {
	tc := New.traces.Empty()
	for i := 0; i < 25; i++ {
		tc.Add(New.Create(0))
	}
	if tc.GetPagesSize(10) != 3 {
		t.Fatal("expected 3")
	}
	if tc.GetPagesSize(0) != 0 {
		t.Fatal("expected 0")
	}
	paged := tc.GetPagedCollection(10)
	if len(paged) != 3 {
		t.Fatal("expected 3")
	}
}

func TestTraceCollection_ConcatNew(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	tc2 := tc.ConcatNew(New.Create(0))
	if tc2.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestTraceCollection_ConcatNewPtr(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	tc2 := tc.ConcatNewPtr(New.Ptr(0), nil)
	if tc2.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestTraceCollection_StackTraces(t *testing.T) {
	tc := TraceCollection{Items: []Trace{New.Create(0)}}
	s := tc.StackTraces()
	if s == "" {
		t.Fatal("expected non-empty")
	}
	jr := tc.StackTracesJsonResult()
	if jr == nil {
		t.Fatal("expected non-nil")
	}
}

func TestTraceCollection_NewStackTraces(t *testing.T) {
	tc := TraceCollection{}
	s := tc.NewStackTraces(0)
	_ = s
	s2 := tc.NewDefaultStackTraces()
	_ = s2
	jr := tc.NewStackTracesJsonResult(0)
	_ = jr
	jr2 := tc.NewDefaultStackTracesJsonResult()
	_ = jr2
}

func TestTraceCollection_NilLength(t *testing.T) {
	var nilTC *TraceCollection
	if nilTC.Length() != 0 {
		t.Fatal("expected 0")
	}
}
