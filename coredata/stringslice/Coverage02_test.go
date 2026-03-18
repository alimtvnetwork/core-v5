package stringslice

import (
	"testing"
)

func TestEmptyPtr(t *testing.T) {
	p := EmptyPtr()
	if p == nil || len(*p) != 0 { t.Fatal("unexpected") }
}

func TestIsEmptyPtr(t *testing.T) {
	if !IsEmptyPtr(nil) { t.Fatal("expected true") }
	s := []string{}
	if !IsEmptyPtr(&s) { t.Fatal("expected true") }
	s2 := []string{"a"}
	if IsEmptyPtr(&s2) { t.Fatal("expected false") }
}

func TestHasAnyItemPtr(t *testing.T) {
	if HasAnyItemPtr(nil) { t.Fatal("expected false") }
	s := []string{"a"}
	if !HasAnyItemPtr(&s) { t.Fatal("expected true") }
}

func TestFirstPtr(t *testing.T) {
	s := []string{"a", "b"}
	if FirstPtr(&s) != "a" { t.Fatal("unexpected") }
}

func TestFirstOrDefaultPtr(t *testing.T) {
	if FirstOrDefaultPtr(nil) != "" { t.Fatal("expected empty") }
	s := []string{"a"}
	if FirstOrDefaultPtr(&s) != "a" { t.Fatal("unexpected") }
}

func TestLastPtr(t *testing.T) {
	s := []string{"a", "b"}
	if LastPtr(&s) != "b" { t.Fatal("unexpected") }
}

func TestLastOrDefaultPtr(t *testing.T) {
	if LastOrDefaultPtr(nil) != "" { t.Fatal("expected empty") }
	s := []string{"a", "b"}
	if LastOrDefaultPtr(&s) != "b" { t.Fatal("unexpected") }
}

func TestClonePtr(t *testing.T) {
	r := ClonePtr(nil)
	if r == nil || len(*r) != 0 { t.Fatal("unexpected") }
	s := []string{"a"}
	r2 := ClonePtr(&s)
	if len(*r2) != 1 { t.Fatal("expected 1") }
}

func TestMakePtr(t *testing.T) {
	p := MakePtr(0, 10)
	if p == nil { t.Fatal("expected non-nil") }
}

func TestMakeDefaultPtr(t *testing.T) {
	p := MakeDefaultPtr(5)
	if p == nil { t.Fatal("expected non-nil") }
}

func TestMakeLen(t *testing.T) {
	s := MakeLen(5)
	if len(s) != 5 { t.Fatal("expected 5") }
}

func TestMakeLenPtr(t *testing.T) {
	p := MakeLenPtr(5)
	if p == nil || len(*p) != 5 { t.Fatal("unexpected") }
}

func TestLengthOfPointer(t *testing.T) {
	if LengthOfPointer(nil) != 0 { t.Fatal("expected 0") }
	s := []string{"a", "b"}
	if LengthOfPointer(&s) != 2 { t.Fatal("expected 2") }
}

func TestLastIndexPtr(t *testing.T) {
	if LastIndexPtr(nil) != -1 { t.Fatal("expected -1") }
	s := []string{"a", "b"}
	if LastIndexPtr(&s) != 1 { t.Fatal("expected 1") }
}

func TestLastSafeIndexPtr(t *testing.T) {
	if LastSafeIndexPtr(nil) != 0 { t.Fatal("expected 0") }
	s := []string{"a", "b"}
	if LastSafeIndexPtr(&s) != 1 { t.Fatal("expected 1") }
}

func TestNonEmptySlicePtr(t *testing.T) {
	r := NonEmptySlicePtr(nil)
	if len(r) != 0 { t.Fatal("expected 0") }
	r2 := NonEmptySlicePtr([]string{"a", "", "b"})
	if len(r2) != 2 { t.Fatal("expected 2") }
}

func TestNonWhitespacePtr(t *testing.T) {
	r := NonWhitespacePtr(nil)
	if r == nil || len(*r) != 0 { t.Fatal("unexpected") }
	s := []string{"a", "  ", "b"}
	r2 := NonWhitespacePtr(&s)
	if len(*r2) != 2 { t.Fatal("expected 2") }
}

func TestNonEmptyJoinPtr(t *testing.T) {
	r := NonEmptyJoinPtr(nil, ",")
	if r != "" { t.Fatal("expected empty") }
	s := []string{"a", "", "b"}
	r2 := NonEmptyJoinPtr(&s, ",")
	if r2 == "" { t.Fatal("expected non-empty") }
}

func TestNonWhitespaceJoin(t *testing.T) {
	r := NonWhitespaceJoin(nil, ",")
	if r != "" { t.Fatal("expected empty") }
	r2 := NonWhitespaceJoin([]string{"a", " ", "b"}, ",")
	if r2 == "" { t.Fatal("expected non-empty") }
}

func TestNonWhitespaceJoinPtr(t *testing.T) {
	r := NonWhitespaceJoinPtr(nil, ",")
	if r != "" { t.Fatal("expected empty") }
	s := []string{"a", " ", "b"}
	r2 := NonWhitespaceJoinPtr(&s, ",")
	if r2 == "" { t.Fatal("expected non-empty") }
}

func TestSafeIndexAtWith(t *testing.T) {
	r := SafeIndexAtWith(nil, 0, "def")
	if r != "def" { t.Fatal("expected def") }
	r2 := SafeIndexAtWith([]string{"a"}, 0, "def")
	if r2 != "a" { t.Fatal("expected a") }
}

func TestSafeIndexAtWithPtr(t *testing.T) {
	r := SafeIndexAtWithPtr(nil, 0, "def")
	if r != "def" { t.Fatal("expected def") }
	s := []string{"a"}
	r2 := SafeIndexAtWithPtr(&s, 0, "def")
	if r2 != "a" { t.Fatal("expected a") }
}

func TestSlicePtr(t *testing.T) {
	s := []string{"a"}
	p := SlicePtr(s)
	if p == nil { t.Fatal("expected non-nil") }
}

func TestNonEmptyIf(t *testing.T) {
	r := NonEmptyIf(true, []string{"a", "", "b"})
	if len(r) != 2 { t.Fatal("expected 2") }
	r2 := NonEmptyIf(false, []string{"a", "", "b"})
	if len(r2) != 3 { t.Fatal("expected 3") }
}

func TestNonNullStrings(t *testing.T) {
	a := "hello"
	r := NonNullStrings(&a, nil)
	if len(r) != 1 { t.Fatal("expected 1") }
	r2 := NonNullStrings()
	if len(r2) != 0 { t.Fatal("expected 0") }
}

func TestNonEmptyStrings(t *testing.T) {
	r := NonEmptyStrings("a", "", "b")
	if len(r) != 2 { t.Fatal("expected 2") }
	r2 := NonEmptyStrings()
	if len(r2) != 0 { t.Fatal("expected 0") }
}

func TestSafeRangeItems(t *testing.T) {
	r := SafeRangeItems([]string{"a", "b", "c"}, 0, 2)
	if len(r) != 2 { t.Fatal("expected 2") }
	r2 := SafeRangeItems(nil, 0, 2)
	if len(r2) != 0 { t.Fatal("expected 0") }
}

func TestSafeRangeItemsPtr(t *testing.T) {
	s := []string{"a", "b", "c"}
	r := SafeRangeItemsPtr(&s, 0, 2)
	if len(r) != 2 { t.Fatal("expected 2") }
	r2 := SafeRangeItemsPtr(nil, 0, 2)
	if len(r2) != 0 { t.Fatal("expected 0") }
}
