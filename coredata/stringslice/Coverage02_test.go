package stringslice

import (
	"testing"
)

func TestEmptyPtr(t *testing.T) {
	p := EmptyPtr()
	if len(p) != 0 { t.Fatal("unexpected") }
}

func TestIsEmptyPtr(t *testing.T) {
	if !IsEmptyPtr(nil) { t.Fatal("expected true") }
	if !IsEmptyPtr([]string{}) { t.Fatal("expected true") }
	if IsEmptyPtr([]string{"a"}) { t.Fatal("expected false") }
}

func TestHasAnyItemPtr(t *testing.T) {
	if HasAnyItemPtr(nil) { t.Fatal("expected false") }
	if !HasAnyItemPtr([]string{"a"}) { t.Fatal("expected true") }
}

func TestFirstPtr(t *testing.T) {
	if FirstPtr([]string{"a", "b"}) != "a" { t.Fatal("unexpected") }
}

func TestFirstOrDefaultPtr(t *testing.T) {
	if FirstOrDefaultPtr(nil) != "" { t.Fatal("expected empty") }
	if FirstOrDefaultPtr([]string{"a"}) != "a" { t.Fatal("unexpected") }
}

func TestLastPtr(t *testing.T) {
	if LastPtr([]string{"a", "b"}) != "b" { t.Fatal("unexpected") }
}

func TestLastOrDefaultPtr(t *testing.T) {
	if LastOrDefaultPtr(nil) != "" { t.Fatal("expected empty") }
	if LastOrDefaultPtr([]string{"a", "b"}) != "b" { t.Fatal("unexpected") }
}

func TestClonePtr(t *testing.T) {
	r := ClonePtr(nil)
	if len(r) != 0 { t.Fatal("unexpected") }
	r2 := ClonePtr([]string{"a"})
	if len(r2) != 1 { t.Fatal("expected 1") }
}

func TestMakePtr(t *testing.T) {
	p := MakePtr(0, 10)
	if cap(p) != 10 { t.Fatal("expected cap 10") }
}

func TestMakeDefaultPtr(t *testing.T) {
	p := MakeDefaultPtr(5)
	if cap(p) < 5 { t.Fatal("expected cap >= 5") }
}

func TestMakeLen(t *testing.T) {
	s := MakeLen(5)
	if len(s) != 5 { t.Fatal("expected 5") }
}

func TestMakeLenPtr(t *testing.T) {
	p := MakeLenPtr(5)
	if len(p) != 5 { t.Fatal("expected 5") }
}

func TestLengthOfPointer(t *testing.T) {
	if LengthOfPointer(nil) != 0 { t.Fatal("expected 0") }
	if LengthOfPointer([]string{"a", "b"}) != 2 { t.Fatal("expected 2") }
}

func TestLastIndexPtr(t *testing.T) {
	if LastIndexPtr([]string{"a", "b"}) != 1 { t.Fatal("expected 1") }
}

func TestLastSafeIndexPtr(t *testing.T) {
	_ = LastSafeIndexPtr(nil)
	if LastSafeIndexPtr([]string{"a", "b"}) != 1 { t.Fatal("expected 1") }
}

func TestNonEmptySlicePtr(t *testing.T) {
	r := NonEmptySlicePtr(nil)
	if len(r) != 0 { t.Fatal("expected 0") }
	r2 := NonEmptySlicePtr([]string{"a", "", "b"})
	if len(r2) != 2 { t.Fatal("expected 2") }
}

func TestNonWhitespacePtr(t *testing.T) {
	r := NonWhitespacePtr(nil)
	if len(r) != 0 { t.Fatal("unexpected") }
	r2 := NonWhitespacePtr([]string{"a", "  ", "b"})
	if len(r2) != 2 { t.Fatal("expected 2") }
}

func TestNonEmptyJoinPtr(t *testing.T) {
	r := NonEmptyJoinPtr(nil, ",")
	if r != "" { t.Fatal("expected empty") }
	r2 := NonEmptyJoinPtr([]string{"a", "", "b"}, ",")
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
	r2 := NonWhitespaceJoinPtr([]string{"a", " ", "b"}, ",")
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
	r2 := SafeIndexAtWithPtr([]string{"a"}, 0, "def")
	if r2 != "a" { t.Fatal("expected a") }
}

func TestSlicePtr(t *testing.T) {
	s := SlicePtr([]string{"a"})
	if len(s) != 1 { t.Fatal("expected 1") }
	s2 := SlicePtr(nil)
	if len(s2) != 0 { t.Fatal("expected 0") }
}

func TestSafeRangeItemsPtr(t *testing.T) {
	r := SafeRangeItemsPtr(nil, 0, 2)
	if len(r) != 0 { t.Fatal("expected 0") }
	r2 := SafeRangeItemsPtr([]string{"a", "b", "c"}, 0, 2)
	if len(r2) != 2 { t.Fatal("expected 2") }
}

func TestTrimmedEachWordsPtr(t *testing.T) {
	r := TrimmedEachWordsPtr(nil)
	if len(r) != 0 { t.Fatal("expected 0") }
	r2 := TrimmedEachWordsPtr([]string{" a ", "b"})
	if len(r2) != 2 { t.Fatal("expected 2") }
}

func TestFirstLastDefault(t *testing.T) {
	f, l := FirstLastDefault(nil)
	if f != "" || l != "" { t.Fatal("expected empty") }
	f2, l2 := FirstLastDefault([]string{"a", "b"})
	if f2 != "a" || l2 != "b" { t.Fatal("unexpected") }
}

func TestFirstLastDefaultPtr(t *testing.T) {
	f, l := FirstLastDefaultPtr(nil)
	if f != "" || l != "" { t.Fatal("expected empty") }
	f2, l2 := FirstLastDefaultPtr([]string{"a", "b"})
	if f2 != "a" || l2 != "b" { t.Fatal("unexpected") }
}

func TestFirstLastDefaultStatus(t *testing.T) {
	s := FirstLastDefaultStatus(nil)
	if s.IsValid { t.Fatal("expected invalid") }
	s2 := FirstLastDefaultStatus([]string{"a"})
	if !s2.HasFirst { t.Fatal("expected has first") }
	s3 := FirstLastDefaultStatus([]string{"a", "b"})
	if !s3.HasLast { t.Fatal("expected has last") }
}

func TestFirstLastDefaultStatusPtr(t *testing.T) {
	s := FirstLastDefaultStatusPtr(nil)
	if s.IsValid { t.Fatal("expected invalid") }
	s2 := FirstLastDefaultStatusPtr([]string{"a", "b"})
	if !s2.IsValid { t.Fatal("expected valid") }
}
