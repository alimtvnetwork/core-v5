package corestr

import (
	"errors"
	"testing"
)

func TestSimpleSlice_Basic(t *testing.T) {
	s := New.SimpleSlice.Empty()
	if !s.IsEmpty() || s.HasAnyItem() {
		t.Fatal("should be empty")
	}
	if s.Length() != 0 || s.Count() != 0 || s.LastIndex() != -1 {
		t.Fatal("wrong length")
	}
}

func TestSimpleSlice_NilReceiver(t *testing.T) {
	var s *SimpleSlice
	if s.Length() != 0 || !s.IsEmpty() {
		t.Fatal("expected 0/empty")
	}
}

func TestSimpleSlice_Add(t *testing.T) {
	s := New.SimpleSlice.Empty()
	s.Add("a").Add("b")
	if s.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestSimpleSlice_AddSplit(t *testing.T) {
	s := New.SimpleSlice.Empty()
	s.AddSplit("a,b,c", ",")
	if s.Length() != 3 {
		t.Fatal("expected 3")
	}
}

func TestSimpleSlice_AddIf(t *testing.T) {
	s := New.SimpleSlice.Empty()
	s.AddIf(false, "skip")
	s.AddIf(true, "add")
	if s.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestSimpleSlice_Adds(t *testing.T) {
	s := New.SimpleSlice.Empty()
	s.Adds()
	s.Adds("a", "b")
	if s.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestSimpleSlice_Append(t *testing.T) {
	s := New.SimpleSlice.Empty()
	s.Append()
	s.Append("a")
	if s.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestSimpleSlice_AppendFmt(t *testing.T) {
	s := New.SimpleSlice.Empty()
	s.AppendFmt("", /* empty */)
	s.AppendFmt("hello %s", "world")
	if s.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestSimpleSlice_AppendFmtIf(t *testing.T) {
	s := New.SimpleSlice.Empty()
	s.AppendFmtIf(false, "skip")
	s.AppendFmtIf(true, "val=%d", 42)
	if s.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestSimpleSlice_AddAsTitleValue(t *testing.T) {
	s := New.SimpleSlice.Empty()
	s.AddAsTitleValue("key", "val")
	if s.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestSimpleSlice_AddAsCurlyTitleWrap(t *testing.T) {
	s := New.SimpleSlice.Empty()
	s.AddAsCurlyTitleWrap("key", "val")
	if s.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestSimpleSlice_AddAsCurlyTitleWrapIf(t *testing.T) {
	s := New.SimpleSlice.Empty()
	s.AddAsCurlyTitleWrapIf(false, "k", "v")
	s.AddAsCurlyTitleWrapIf(true, "k", "v")
	if s.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestSimpleSlice_AddAsTitleValueIf(t *testing.T) {
	s := New.SimpleSlice.Empty()
	s.AddAsTitleValueIf(false, "k", "v")
	s.AddAsTitleValueIf(true, "k", "v")
	if s.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestSimpleSlice_InsertAt(t *testing.T) {
	s := New.SimpleSlice.Lines("a", "c")
	s.InsertAt(1, "b")
	if s.Length() != 3 || (*s)[1] != "b" {
		t.Fatal("unexpected")
	}
	// out of range
	s.InsertAt(-1, "x")
	s.InsertAt(100, "x")
}

func TestSimpleSlice_AddsIf(t *testing.T) {
	s := New.SimpleSlice.Empty()
	s.AddsIf(false, "skip")
	s.AddsIf(true, "a", "b")
	if s.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestSimpleSlice_AddError(t *testing.T) {
	s := New.SimpleSlice.Empty()
	s.AddError(nil)
	s.AddError(errors.New("e"))
	if s.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestSimpleSlice_AddStruct(t *testing.T) {
	s := New.SimpleSlice.Empty()
	s.AddStruct(false, nil)
	s.AddStruct(false, "hello")
	if s.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestSimpleSlice_AddPointer(t *testing.T) {
	s := New.SimpleSlice.Empty()
	s.AddPointer(false, nil)
	val := "hello"
	s.AddPointer(false, &val)
	if s.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestSimpleSlice_AsError(t *testing.T) {
	s := New.SimpleSlice.Empty()
	if s.AsDefaultError() != nil {
		t.Fatal("expected nil")
	}
	if s.AsError(",") != nil {
		t.Fatal("expected nil")
	}
	s.Add("e")
	if s.AsDefaultError() == nil {
		t.Fatal("expected non-nil")
	}
}

func TestSimpleSlice_FirstLast(t *testing.T) {
	s := New.SimpleSlice.Lines("a", "b", "c")
	if s.First() != "a" || s.Last() != "c" {
		t.Fatal("unexpected")
	}
	if s.FirstDynamic() != "a" || s.LastDynamic() != "c" {
		t.Fatal("unexpected")
	}
	if s.FirstOrDefault() != "a" || s.LastOrDefault() != "c" {
		t.Fatal("unexpected")
	}
	if s.FirstOrDefaultDynamic() != "a" || s.LastOrDefaultDynamic() != "c" {
		t.Fatal("unexpected")
	}
}

func TestSimpleSlice_FirstOrDefault_Empty(t *testing.T) {
	s := New.SimpleSlice.Empty()
	if s.FirstOrDefault() != "" || s.LastOrDefault() != "" {
		t.Fatal("expected empty")
	}
}

func TestSimpleSlice_SkipTake(t *testing.T) {
	s := New.SimpleSlice.Lines("a", "b", "c")
	if len(s.Skip(1)) != 2 {
		t.Fatal("expected 2")
	}
	if len(s.Skip(100)) != 0 {
		t.Fatal("expected 0")
	}
	if len(s.Take(2)) != 2 {
		t.Fatal("expected 2")
	}
	if len(s.Take(100)) != 3 {
		t.Fatal("expected 3")
	}
	_ = s.SkipDynamic(1)
	_ = s.TakeDynamic(1)
	_ = s.LimitDynamic(1)
	_ = s.Limit(1)
}

func TestSimpleSlice_CountFunc(t *testing.T) {
	s := New.SimpleSlice.Lines("a", "bb", "ccc")
	count := s.CountFunc(func(i int, item string) bool { return len(item) > 1 })
	if count != 2 {
		t.Fatal("expected 2")
	}
	empty := New.SimpleSlice.Empty()
	if empty.CountFunc(func(i int, item string) bool { return true }) != 0 {
		t.Fatal("expected 0")
	}
}

func TestSimpleSlice_IsContains(t *testing.T) {
	s := New.SimpleSlice.Lines("a", "b")
	if !s.IsContains("a") || s.IsContains("c") {
		t.Fatal("unexpected")
	}
	empty := New.SimpleSlice.Empty()
	if empty.IsContains("a") {
		t.Fatal("unexpected")
	}
}

func TestSimpleSlice_IsContainsFunc(t *testing.T) {
	s := New.SimpleSlice.Lines("hello", "world")
	found := s.IsContainsFunc("hello", func(item, searching string) bool { return item == searching })
	if !found {
		t.Fatal("expected found")
	}
	empty := New.SimpleSlice.Empty()
	if empty.IsContainsFunc("a", func(item, searching string) bool { return true }) {
		t.Fatal("unexpected")
	}
}

func TestSimpleSlice_IndexOf(t *testing.T) {
	s := New.SimpleSlice.Lines("a", "b", "c")
	if s.IndexOf("b") != 1 || s.IndexOf("z") != -1 {
		t.Fatal("unexpected")
	}
	empty := New.SimpleSlice.Empty()
	if empty.IndexOf("a") != -1 {
		t.Fatal("expected -1")
	}
}

func TestSimpleSlice_IndexOfFunc(t *testing.T) {
	s := New.SimpleSlice.Lines("a", "b")
	idx := s.IndexOfFunc("b", func(item, searching string) bool { return item == searching })
	if idx != 1 {
		t.Fatal("expected 1")
	}
	empty := New.SimpleSlice.Empty()
	if empty.IndexOfFunc("a", func(item, searching string) bool { return true }) != -1 {
		t.Fatal("expected -1")
	}
}

func TestSimpleSlice_HasIndex(t *testing.T) {
	s := New.SimpleSlice.Lines("a", "b")
	if !s.HasIndex(0) || !s.HasIndex(1) || s.HasIndex(2) || s.HasIndex(-1) {
		t.Fatal("unexpected")
	}
}

func TestSimpleSlice_StringsList(t *testing.T) {
	s := New.SimpleSlice.Lines("a")
	if len(s.Strings()) != 1 || len(s.List()) != 1 {
		t.Fatal("expected 1")
	}
}

func TestSimpleSlice_WrapQuotes(t *testing.T) {
	s := New.SimpleSlice.Lines("a")
	_ = s.WrapDoubleQuote()
	s2 := New.SimpleSlice.Lines("a")
	_ = s2.WrapSingleQuote()
	s3 := New.SimpleSlice.Lines("a")
	_ = s3.WrapTildaQuote()
	s4 := New.SimpleSlice.Lines("a")
	_ = s4.WrapDoubleQuoteIfMissing()
	s5 := New.SimpleSlice.Lines("a")
	_ = s5.WrapSingleQuoteIfMissing()
}
