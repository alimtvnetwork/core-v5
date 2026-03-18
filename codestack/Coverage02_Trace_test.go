package codestack

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

func TestTrace_Create(t *testing.T) {
	trace := New.Create(0)
	if !trace.IsOkay {
		t.Fatal("should be okay")
	}
	if trace.PackageName == "" {
		t.Fatal("expected package name")
	}
}

func TestTrace_Default(t *testing.T) {
	trace := New.Default()
	if !trace.IsOkay {
		t.Fatal("should be okay")
	}
}

func TestTrace_SkipOne(t *testing.T) {
	trace := New.SkipOne()
	// may or may not be okay depending on call depth
	_ = trace
}

func TestTrace_Ptr(t *testing.T) {
	p := New.Ptr(0)
	if p == nil {
		t.Fatal("expected non-nil")
	}
}

func TestTrace_Message(t *testing.T) {
	trace := New.Create(0)
	m := trace.Message()
	if m == "" {
		t.Fatal("expected non-empty")
	}
	// call again for cache hit
	m2 := trace.Message()
	if m != m2 {
		t.Fatal("cache miss")
	}
}

func TestTrace_ShortString(t *testing.T) {
	trace := New.Create(0)
	s := trace.ShortString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
	// call again for cache hit
	s2 := trace.ShortString()
	if s != s2 {
		t.Fatal("cache miss")
	}
}

func TestTrace_IsNil_IsNotNil(t *testing.T) {
	trace := New.Ptr(0)
	if trace.IsNil() {
		t.Fatal("should not be nil")
	}
	if !trace.IsNotNil() {
		t.Fatal("should be not nil")
	}

	var nilT *Trace
	if !nilT.IsNil() {
		t.Fatal("should be nil")
	}
	if nilT.IsNotNil() {
		t.Fatal("should not be not nil")
	}
}

func TestTrace_HasIssues(t *testing.T) {
	trace := New.Create(0)
	if trace.HasIssues() {
		t.Fatal("should not have issues")
	}

	badTrace := Trace{}
	if !badTrace.HasIssues() {
		t.Fatal("should have issues")
	}

	var nilT *Trace
	if !nilT.HasIssues() {
		t.Fatal("nil should have issues")
	}
}

func TestTrace_String(t *testing.T) {
	trace := New.Create(0)
	s := trace.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}

	var nilT *Trace
	if nilT.String() != "" {
		t.Fatal("expected empty for nil")
	}
}

func TestTrace_StringUsingFmt(t *testing.T) {
	trace := New.Create(0)
	s := trace.StringUsingFmt(func(tr Trace) string {
		return tr.PackageName
	})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestTrace_FileWithLine(t *testing.T) {
	trace := New.Create(0)
	fwl := trace.FileWithLine()
	if fwl.FilePath == "" {
		t.Fatal("expected file path")
	}
}

func TestTrace_FullFilePath(t *testing.T) {
	trace := New.Create(0)
	if trace.FullFilePath() == "" {
		t.Fatal("expected file path")
	}
}

func TestTrace_FileName(t *testing.T) {
	trace := New.Create(0)
	if trace.FileName() == "" {
		t.Fatal("expected file name")
	}
}

func TestTrace_LineNumber(t *testing.T) {
	trace := New.Create(0)
	if trace.LineNumber() == 0 {
		t.Fatal("expected non-zero line")
	}
}

func TestTrace_FileWithLineString(t *testing.T) {
	trace := New.Create(0)
	s := trace.FileWithLineString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestTrace_JsonModel(t *testing.T) {
	trace := New.Create(0)
	m := trace.JsonModel()
	if m.PackageName == "" {
		t.Fatal("unexpected")
	}
}

func TestTrace_JsonModelAny(t *testing.T) {
	trace := New.Create(0)
	a := trace.JsonModelAny()
	if a == nil {
		t.Fatal("expected non-nil")
	}
}

func TestTrace_Dispose(t *testing.T) {
	trace := New.Create(0)
	trace.Dispose()
	if trace.PackageName != "" {
		t.Fatal("expected empty after dispose")
	}

	var nilT *Trace
	nilT.Dispose() // should not panic
}

func TestTrace_JsonString(t *testing.T) {
	trace := New.Create(0)
	s := trace.JsonString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestTrace_Json_JsonPtr(t *testing.T) {
	trace := New.Create(0)
	j := trace.Json()
	if j.HasError() {
		t.Fatal(j.Error)
	}
	jp := trace.JsonPtr()
	if jp == nil {
		t.Fatal("expected non-nil")
	}
}

func TestTrace_ParseInjectUsingJson(t *testing.T) {
	trace := New.Create(0)
	jr := corejson.NewPtr(trace)
	target := &Trace{}
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

func TestTrace_ParseInjectUsingJsonMust(t *testing.T) {
	trace := New.Create(0)
	jr := corejson.NewPtr(trace)
	target := &Trace{}
	result := target.ParseInjectUsingJsonMust(jr)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func TestTrace_JsonParseSelfInject(t *testing.T) {
	trace := New.Create(0)
	jr := corejson.NewPtr(trace)
	target := &Trace{}
	err := target.JsonParseSelfInject(jr)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTrace_Clone(t *testing.T) {
	trace := New.Create(0)
	c := trace.Clone()
	if c.PackageName != trace.PackageName {
		t.Fatal("clone mismatch")
	}
}

func TestTrace_ClonePtr(t *testing.T) {
	trace := New.Ptr(0)
	cp := trace.ClonePtr()
	if cp == nil {
		t.Fatal("expected non-nil")
	}

	var nilT *Trace
	if nilT.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
}

func TestTrace_AsFileLiner(t *testing.T) {
	trace := New.Ptr(0)
	liner := trace.AsFileLiner()
	if liner == nil {
		t.Fatal("expected non-nil")
	}
}
