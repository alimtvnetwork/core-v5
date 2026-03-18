package stringslice

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	if !IsEmpty(nil) { t.Fatal("expected true") }
	if !IsEmpty([]string{}) { t.Fatal("expected true") }
	if IsEmpty([]string{"a"}) { t.Fatal("expected false") }
}

func TestHasAnyItem(t *testing.T) {
	if HasAnyItem(nil) { t.Fatal("expected false") }
	if !HasAnyItem([]string{"a"}) { t.Fatal("expected true") }
}

func TestEmpty(t *testing.T) {
	e := Empty()
	if len(e) != 0 { t.Fatal("expected empty") }
}

func TestFirst(t *testing.T) {
	if First([]string{"a", "b"}) != "a" { t.Fatal("unexpected") }
}

func TestFirstOrDefault(t *testing.T) {
	if FirstOrDefault(nil) != "" { t.Fatal("expected empty") }
	if FirstOrDefault([]string{"a"}) != "a" { t.Fatal("unexpected") }
}

func TestLast(t *testing.T) {
	if Last([]string{"a", "b"}) != "b" { t.Fatal("unexpected") }
}

func TestLastOrDefault(t *testing.T) {
	if LastOrDefault(nil) != "" { t.Fatal("expected empty") }
	if LastOrDefault([]string{"a", "b"}) != "b" { t.Fatal("unexpected") }
}

func TestClone(t *testing.T) {
	c := Clone(nil)
	if len(c) != 0 { t.Fatal("expected empty") }
	c2 := Clone([]string{"a", "b"})
	if len(c2) != 2 { t.Fatal("expected 2") }
}

func TestCloneIf(t *testing.T) {
	_ = CloneIf(false, 0, nil)
	_ = CloneIf(false, 0, []string{"a"})
	_ = CloneIf(true, 5, []string{"a"})
}

func TestMake(t *testing.T) {
	s := Make(0, 10)
	if cap(s) != 10 { t.Fatal("expected cap 10") }
}

func TestMergeNew(t *testing.T) {
	r := MergeNew([]string{"a"}, "b", "c")
	if len(r) != 3 { t.Fatal("expected 3") }
	r2 := MergeNew(nil)
	if len(r2) != 0 { t.Fatal("expected 0") }
}

func TestNonEmptySlice(t *testing.T) {
	r := NonEmptySlice([]string{"a", "", "b"})
	if len(r) != 2 { t.Fatal("expected 2") }
	r2 := NonEmptySlice(nil)
	if len(r2) != 0 { t.Fatal("expected 0") }
}

func TestNonWhitespace(t *testing.T) {
	r := NonWhitespace([]string{"a", "  ", "b"})
	if len(r) != 2 { t.Fatal("expected 2") }
	r2 := NonWhitespace(nil)
	if len(r2) != 0 { t.Fatal("expected 0") }
	r3 := NonWhitespace([]string{})
	if len(r3) != 0 { t.Fatal("expected 0") }
}

func TestIndexAt(t *testing.T) {
	if IndexAt([]string{"a", "b"}, 1) != "b" { t.Fatal("unexpected") }
}

func TestSafeIndexAt(t *testing.T) {
	if SafeIndexAt(nil, 0) != "" { t.Fatal("expected empty") }
	if SafeIndexAt([]string{"a"}, -1) != "" { t.Fatal("expected empty") }
	if SafeIndexAt([]string{"a"}, 5) != "" { t.Fatal("expected empty") }
	if SafeIndexAt([]string{"a"}, 0) != "a" { t.Fatal("unexpected") }
}

func TestInPlaceReverse(t *testing.T) {
	r := InPlaceReverse(nil)
	if len(*r) != 0 { t.Fatal("expected empty") }
	s1 := []string{"a"}
	InPlaceReverse(&s1)
	s2 := []string{"a", "b"}
	InPlaceReverse(&s2)
	if s2[0] != "b" { t.Fatal("expected b") }
	s3 := []string{"a", "b", "c"}
	InPlaceReverse(&s3)
	if s3[0] != "c" { t.Fatal("expected c") }
}

func TestPrependNew(t *testing.T) {
	r := PrependNew([]string{"b"}, "a")
	if len(*r) != 2 || (*r)[0] != "a" { t.Fatal("unexpected") }
	r2 := PrependNew(nil)
	if len(*r2) != 0 { t.Fatal("expected 0") }
}

func TestAppendLineNew(t *testing.T) {
	r := AppendLineNew([]string{"a"}, "b")
	if len(r) != 2 { t.Fatal("expected 2") }
}

func TestSortIf(t *testing.T) {
	s := []string{"b", "a"}
	SortIf(true, s)
	if s[0] != "a" { t.Fatal("expected a") }
	SortIf(false, s)
}

func TestExpandBySplit(t *testing.T) {
	r := ExpandBySplit([]string{"a,b", "c"}, ",")
	if len(r) != 3 { t.Fatal("expected 3") }
	r2 := ExpandBySplit(nil, ",")
	if len(r2) != 0 { t.Fatal("expected 0") }
}

func TestNonEmptyJoin(t *testing.T) {
	r := NonEmptyJoin([]string{"a", "", "b"}, ",")
	if r == "" { t.Fatal("expected non-empty") }
	r2 := NonEmptyJoin(nil, ",")
	if r2 != "" { t.Fatal("expected empty") }
	r3 := NonEmptyJoin([]string{}, ",")
	if r3 != "" { t.Fatal("expected empty") }
}

func TestJoinWith(t *testing.T) {
	r := JoinWith(",", "a", "b")
	if r == "" { t.Fatal("expected non-empty") }
	r2 := JoinWith(",")
	if r2 != "" { t.Fatal("expected empty") }
}

func TestJoins(t *testing.T) {
	r := Joins(",", "a", "b")
	if r != "a,b" { t.Fatal("unexpected") }
	r2 := Joins(",")
	if r2 != "" { t.Fatal("expected empty") }
}

func TestMakeDefault(t *testing.T) {
	s := MakeDefault(5)
	if cap(s) < 5 { t.Fatal("expected cap >= 5") }
}

func TestCloneUsingCap(t *testing.T) {
	r := CloneUsingCap(5, []string{"a"})
	if len(r) != 1 { t.Fatal("expected 1") }
	r2 := CloneUsingCap(0, nil)
	if len(r2) != 0 { t.Fatal("expected 0") }
}

func TestMergeNewSimple(t *testing.T) {
	r := MergeNewSimple([]string{"a"}, []string{"b"})
	if len(r) != 2 { t.Fatal("expected 2") }
}

func TestPrependLineNew(t *testing.T) {
	r := PrependLineNew([]string{"b"}, "a")
	if len(*r) != 2 { t.Fatal("expected 2") }
}
