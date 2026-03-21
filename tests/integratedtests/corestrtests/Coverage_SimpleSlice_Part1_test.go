package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// SimpleSlice — Segment 8: Add variants, accessors, search, wrap (L1-700)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovSS1_01_Add(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items()
	ss.Add("a")
	if ss.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovSS1_02_AddSplit(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items()
	ss.AddSplit("a,b,c", ",")
	if ss.Length() != 3 {
		t.Fatalf("expected 3, got %d", ss.Length())
	}
}

func Test_CovSS1_03_AddIf(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items()
	ss.AddIf(false, "skip")
	if ss.Length() != 0 {
		t.Fatal("expected 0")
	}
	ss.AddIf(true, "keep")
	if ss.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovSS1_04_Adds_Append(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items()
	ss.Adds("a", "b")
	if ss.Length() != 2 {
		t.Fatal("expected 2")
	}
	ss.Adds()
	ss.Append("c")
	if ss.Length() != 3 {
		t.Fatal("expected 3")
	}
	ss.Append()
}

func Test_CovSS1_05_AppendFmt(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items()
	ss.AppendFmt("hello %s", "world")
	if ss.Length() != 1 {
		t.Fatal("expected 1")
	}
	// empty format and no values
	ss.AppendFmt("")
}

func Test_CovSS1_06_AppendFmtIf(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items()
	ss.AppendFmtIf(false, "skip %s", "x")
	if ss.Length() != 0 {
		t.Fatal("expected 0")
	}
	ss.AppendFmtIf(true, "keep %s", "x")
	if ss.Length() != 1 {
		t.Fatal("expected 1")
	}
	// empty format
	ss.AppendFmtIf(true, "")
}

func Test_CovSS1_07_AddAsTitleValue(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items()
	ss.AddAsTitleValue("Name", "Alice")
	if ss.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovSS1_08_AddAsCurlyTitleWrap(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items()
	ss.AddAsCurlyTitleWrap("Name", "Alice")
	if ss.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovSS1_09_AddAsCurlyTitleWrapIf(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items()
	ss.AddAsCurlyTitleWrapIf(false, "skip", "x")
	if ss.Length() != 0 {
		t.Fatal("expected 0")
	}
	ss.AddAsCurlyTitleWrapIf(true, "keep", "x")
	if ss.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovSS1_10_AddAsTitleValueIf(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items()
	ss.AddAsTitleValueIf(false, "skip", "x")
	if ss.Length() != 0 {
		t.Fatal("expected 0")
	}
	ss.AddAsTitleValueIf(true, "keep", "x")
	if ss.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovSS1_11_InsertAt(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a", "c")
	ss.InsertAt(1, "b")
	if ss.Length() != 3 {
		t.Fatal("expected 3")
	}
	// out of range
	ss.InsertAt(-1, "x")
	ss.InsertAt(100, "x")
}

func Test_CovSS1_12_AddStruct(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items()
	ss.AddStruct(true, struct{ Name string }{"Alice"})
	if ss.Length() != 1 {
		t.Fatal("expected 1")
	}
	ss.AddStruct(true, nil)
	if ss.Length() != 1 {
		t.Fatal("expected still 1")
	}
}

func Test_CovSS1_13_AddPointer(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items()
	v := "hello"
	ss.AddPointer(false, &v)
	if ss.Length() != 1 {
		t.Fatal("expected 1")
	}
	ss.AddPointer(false, nil)
}

func Test_CovSS1_14_AddsIf(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items()
	ss.AddsIf(false, "a", "b")
	if ss.Length() != 0 {
		t.Fatal("expected 0")
	}
	ss.AddsIf(true, "a", "b")
	if ss.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CovSS1_15_AddError(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items()
	ss.AddError(nil)
	if ss.Length() != 0 {
		t.Fatal("expected 0")
	}
	ss.AddError(fmt.Errorf("oops"))
	if ss.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovSS1_16_AsDefaultError_AsError(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("err1", "err2")
	e := ss.AsDefaultError()
	if e == nil {
		t.Fatal("expected error")
	}
	e2 := ss.AsError(",")
	if e2 == nil {
		t.Fatal("expected error")
	}
	// empty
	empty := corestr.New.SimpleSlice.Items()
	if empty.AsError(",") != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovSS1_17_First_Last_Dynamic(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a", "b", "c")
	if ss.First() != "a" {
		t.Fatal("expected a")
	}
	if ss.Last() != "c" {
		t.Fatal("expected c")
	}
	_ = ss.FirstDynamic()
	_ = ss.LastDynamic()
}

func Test_CovSS1_18_FirstOrDefault_LastOrDefault(t *testing.T) {
	empty := corestr.New.SimpleSlice.Items()
	if empty.FirstOrDefault() != "" {
		t.Fatal("expected empty")
	}
	if empty.LastOrDefault() != "" {
		t.Fatal("expected empty")
	}
	_ = empty.FirstOrDefaultDynamic()
	_ = empty.LastOrDefaultDynamic()

	ss := corestr.New.SimpleSlice.Items("a")
	if ss.FirstOrDefault() != "a" {
		t.Fatal("expected a")
	}
	if ss.LastOrDefault() != "a" {
		t.Fatal("expected a")
	}
}

func Test_CovSS1_19_Skip_SkipDynamic(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a", "b", "c")
	r := ss.Skip(1)
	if len(r) != 2 {
		t.Fatal("expected 2")
	}
	// skip all
	r2 := ss.Skip(10)
	if len(r2) != 0 {
		t.Fatal("expected 0")
	}
	_ = ss.SkipDynamic(1)
	_ = ss.SkipDynamic(10)
}

func Test_CovSS1_20_Take_TakeDynamic_Limit(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a", "b", "c")
	r := ss.Take(2)
	if len(r) != 2 {
		t.Fatal("expected 2")
	}
	// take all
	r2 := ss.Take(10)
	if len(r2) != 3 {
		t.Fatal("expected 3")
	}
	_ = ss.TakeDynamic(2)
	_ = ss.TakeDynamic(10)
	_ = ss.Limit(1)
	_ = ss.LimitDynamic(1)
}

func Test_CovSS1_21_Length_Count_CountFunc(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a", "bb", "ccc")
	if ss.Length() != 3 {
		t.Fatal("expected 3")
	}
	if ss.Count() != 3 {
		t.Fatal("expected 3")
	}
	c := ss.CountFunc(func(i int, s string) bool { return len(s) > 1 })
	if c != 2 {
		t.Fatal("expected 2")
	}
	// empty
	e := corestr.New.SimpleSlice.Items()
	if e.CountFunc(func(i int, s string) bool { return true }) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovSS1_22_IsEmpty_HasAnyItem(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items()
	if !ss.IsEmpty() {
		t.Fatal("expected empty")
	}
	if ss.HasAnyItem() {
		t.Fatal("expected no items")
	}
	ss.Add("a")
	if ss.IsEmpty() {
		t.Fatal("expected not empty")
	}
	if !ss.HasAnyItem() {
		t.Fatal("expected items")
	}
}

func Test_CovSS1_23_IsContains(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a", "b")
	if !ss.IsContains("a") {
		t.Fatal("expected true")
	}
	if ss.IsContains("z") {
		t.Fatal("expected false")
	}
	// empty
	e := corestr.New.SimpleSlice.Items()
	if e.IsContains("a") {
		t.Fatal("expected false")
	}
}

func Test_CovSS1_24_IsContainsFunc(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("abc", "def")
	found := ss.IsContainsFunc("abc", func(item, searching string) bool {
		return item == searching
	})
	if !found {
		t.Fatal("expected true")
	}
	// empty
	e := corestr.New.SimpleSlice.Items()
	if e.IsContainsFunc("x", func(a, b string) bool { return a == b }) {
		t.Fatal("expected false")
	}
}

func Test_CovSS1_25_IndexOf_IndexOfFunc(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a", "b", "c")
	if ss.IndexOf("b") != 1 {
		t.Fatal("expected 1")
	}
	if ss.IndexOf("z") != -1 {
		t.Fatal("expected -1")
	}
	idx := ss.IndexOfFunc("b", func(item, searching string) bool {
		return item == searching
	})
	if idx != 1 {
		t.Fatal("expected 1")
	}
	// empty
	e := corestr.New.SimpleSlice.Items()
	if e.IndexOf("a") != -1 {
		t.Fatal("expected -1")
	}
	if e.IndexOfFunc("a", func(a, b string) bool { return a == b }) != -1 {
		t.Fatal("expected -1")
	}
}

func Test_CovSS1_26_LastIndex_HasIndex(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a", "b")
	if ss.LastIndex() != 1 {
		t.Fatal("expected 1")
	}
	if !ss.HasIndex(0) {
		t.Fatal("expected true")
	}
	if ss.HasIndex(5) {
		t.Fatal("expected false")
	}
	if ss.HasIndex(-1) {
		t.Fatal("expected false")
	}
}

func Test_CovSS1_27_Strings_List(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a")
	if len(ss.Strings()) != 1 {
		t.Fatal("expected 1")
	}
	if len(ss.List()) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovSS1_28_WrapQuotes(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a")
	r := ss.WrapDoubleQuote()
	if (*r)[0] != `"a"` {
		t.Fatal("expected wrapped")
	}
	ss2 := corestr.New.SimpleSlice.Items("a")
	r2 := ss2.WrapSingleQuote()
	if (*r2)[0] != "'a'" {
		t.Fatal("expected wrapped")
	}
	ss3 := corestr.New.SimpleSlice.Items("a")
	_ = ss3.WrapTildaQuote()
	ss4 := corestr.New.SimpleSlice.Items("a")
	_ = ss4.WrapDoubleQuoteIfMissing()
	ss5 := corestr.New.SimpleSlice.Items("a")
	_ = ss5.WrapSingleQuoteIfMissing()
}

func Test_CovSS1_29_Transpile_TranspileJoin(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a", "b")
	r := ss.Transpile(func(s string) string { return s + "!" })
	if (*r)[0] != "a!" {
		t.Fatal("expected a!")
	}
	// empty
	e := corestr.New.SimpleSlice.Items()
	_ = e.Transpile(func(s string) string { return s })
	// TranspileJoin
	s := ss.TranspileJoin(func(s string) string { return s }, ",")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovSS1_30_Hashset(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a", "b", "a")
	hs := ss.Hashset()
	if hs.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CovSS1_31_Join_JoinLine_JoinSpace_JoinComma(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a", "b")
	if ss.Join(",") != "a,b" {
		t.Fatal("expected a,b")
	}
	_ = ss.JoinLine()
	_ = ss.JoinSpace()
	_ = ss.JoinComma()
	// empty
	e := corestr.New.SimpleSlice.Items()
	if e.Join(",") != "" {
		t.Fatal("expected empty")
	}
	if e.JoinLine() != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovSS1_32_JoinLineEofLine(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a", "b")
	r := ss.JoinLineEofLine()
	if r == "" {
		t.Fatal("expected non-empty")
	}
	// empty
	e := corestr.New.SimpleSlice.Items()
	if e.JoinLineEofLine() != "" {
		t.Fatal("expected empty")
	}
	// already has suffix
	ss2 := corestr.New.SimpleSlice.Items("a\n")
	_ = ss2.JoinLineEofLine()
}

func Test_CovSS1_33_JoinCsv_JoinCsvLine_JoinCsvString(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a", "b")
	_ = ss.JoinCsv()
	_ = ss.JoinCsvLine()
	s := ss.JoinCsvString(",")
	if s == "" {
		t.Fatal("expected non-empty")
	}
	// empty
	e := corestr.New.SimpleSlice.Items()
	if e.JoinCsvString(",") != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovSS1_34_CsvStrings(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a")
	r := ss.CsvStrings()
	if len(r) != 1 {
		t.Fatal("expected 1")
	}
	// empty
	e := corestr.New.SimpleSlice.Items()
	if len(e.CsvStrings()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovSS1_35_EachItemSplitBy(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a,b", "c,d")
	r := ss.EachItemSplitBy(",")
	if r.Length() != 4 {
		t.Fatalf("expected 4, got %d", r.Length())
	}
}

func Test_CovSS1_36_PrependJoin_AppendJoin(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("b")
	s := ss.PrependJoin(",", "a")
	if s != "a,b" {
		t.Fatalf("expected 'a,b', got '%s'", s)
	}
	s2 := ss.AppendJoin(",", "c")
	if s2 != "b,c" {
		t.Fatalf("expected 'b,c', got '%s'", s2)
	}
}

func Test_CovSS1_37_PrependAppend(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("b")
	ss.PrependAppend([]string{"a"}, []string{"c"})
	if ss.Length() != 3 {
		t.Fatal("expected 3")
	}
	// empty prepend/append
	ss.PrependAppend(nil, nil)
}

func Test_CovSS1_38_JoinWith(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a", "b")
	r := ss.JoinWith(",")
	if r == "" {
		t.Fatal("expected non-empty")
	}
	// empty
	e := corestr.New.SimpleSlice.Items()
	if e.JoinWith(",") != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovSS1_39_IsEqual(t *testing.T) {
	a := corestr.New.SimpleSlice.Items("a", "b")
	b := corestr.New.SimpleSlice.Items("a", "b")
	if !a.IsEqual(b) {
		t.Fatal("expected equal")
	}
	// nil
	if a.IsEqual(nil) {
		t.Fatal("expected false")
	}
	// diff length
	c := corestr.New.SimpleSlice.Items("a")
	if a.IsEqual(c) {
		t.Fatal("expected false")
	}
	// both empty
	e1 := corestr.New.SimpleSlice.Items()
	e2 := corestr.New.SimpleSlice.Items()
	if !e1.IsEqual(e2) {
		t.Fatal("expected true")
	}
}

func Test_CovSS1_40_IsEqualLines(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("a", "b")
	if !ss.IsEqualLines([]string{"a", "b"}) {
		t.Fatal("expected true")
	}
	if ss.IsEqualLines([]string{"a", "c"}) {
		t.Fatal("expected false")
	}
	if ss.IsEqualLines([]string{"a"}) {
		t.Fatal("expected false (diff length)")
	}
}

func Test_CovSS1_41_IsEqualUnorderedLines(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("b", "a")
	if !ss.IsEqualUnorderedLines([]string{"a", "b"}) {
		t.Fatal("expected true")
	}
	if ss.IsEqualUnorderedLines([]string{"a", "c"}) {
		t.Fatal("expected false")
	}
	// diff length
	if ss.IsEqualUnorderedLines([]string{"a"}) {
		t.Fatal("expected false")
	}
	// both empty
	e := corestr.New.SimpleSlice.Items()
	if !e.IsEqualUnorderedLines([]string{}) {
		t.Fatal("expected true")
	}
}

func Test_CovSS1_42_IsEqualUnorderedLinesClone(t *testing.T) {
	ss := corestr.New.SimpleSlice.Items("b", "a")
	if !ss.IsEqualUnorderedLinesClone([]string{"a", "b"}) {
		t.Fatal("expected true")
	}
	if ss.IsEqualUnorderedLinesClone([]string{"a", "c"}) {
		t.Fatal("expected false")
	}
	// diff length
	if ss.IsEqualUnorderedLinesClone([]string{"a"}) {
		t.Fatal("expected false")
	}
	// both empty
	e := corestr.New.SimpleSlice.Items()
	if !e.IsEqualUnorderedLinesClone([]string{}) {
		t.Fatal("expected true")
	}
}
