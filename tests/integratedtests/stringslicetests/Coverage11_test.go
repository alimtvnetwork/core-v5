package stringslicetests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Empty / EmptyPtr ──

func Test_Cov11_Empty(t *testing.T) {
	result := stringslice.Empty()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty returns empty slice", actual)
}

func Test_Cov11_EmptyPtr(t *testing.T) {
	result := stringslice.EmptyPtr()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "EmptyPtr returns empty slice", actual)
}

// ── IsEmpty / IsEmptyPtr / HasAnyItem / HasAnyItemPtr ──

func Test_Cov11_IsEmpty_True(t *testing.T) {
	actual := args.Map{"empty": stringslice.IsEmpty([]string{})}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty true on empty", actual)
}

func Test_Cov11_IsEmpty_False(t *testing.T) {
	actual := args.Map{"empty": stringslice.IsEmpty([]string{"a"})}
	expected := args.Map{"empty": false}
	expected.ShouldBeEqual(t, 0, "IsEmpty false on non-empty", actual)
}

func Test_Cov11_IsEmptyPtr(t *testing.T) {
	actual := args.Map{"empty": stringslice.IsEmptyPtr(nil)}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr true on nil", actual)
}

func Test_Cov11_HasAnyItem(t *testing.T) {
	actual := args.Map{"has": stringslice.HasAnyItem([]string{"x"})}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem true", actual)
}

func Test_Cov11_HasAnyItemPtr(t *testing.T) {
	actual := args.Map{"has": stringslice.HasAnyItemPtr([]string{"x"})}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItemPtr true", actual)
}

func Test_Cov11_HasAnyItemPtr_Empty(t *testing.T) {
	actual := args.Map{"has": stringslice.HasAnyItemPtr(nil)}
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "HasAnyItemPtr false on nil", actual)
}

// ── LengthOfPointer ──

func Test_Cov11_LengthOfPointer(t *testing.T) {
	actual := args.Map{"len": stringslice.LengthOfPointer([]string{"a", "b"})}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfPointer", actual)
}

// ── Make / MakePtr / MakeLen / MakeLenPtr / MakeDefault / MakeDefaultPtr ──

func Test_Cov11_Make(t *testing.T) {
	s := stringslice.Make(3, 5)
	actual := args.Map{"len": len(s), "cap": cap(s)}
	expected := args.Map{"len": 3, "cap": 5}
	expected.ShouldBeEqual(t, 0, "Make", actual)
}

func Test_Cov11_MakePtr(t *testing.T) {
	s := stringslice.MakePtr(3, 5)
	actual := args.Map{"len": len(s), "cap": cap(s)}
	expected := args.Map{"len": 3, "cap": 5}
	expected.ShouldBeEqual(t, 0, "MakePtr", actual)
}

func Test_Cov11_MakeLen(t *testing.T) {
	s := stringslice.MakeLen(4)
	actual := args.Map{"len": len(s)}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "MakeLen", actual)
}

func Test_Cov11_MakeLenPtr(t *testing.T) {
	s := stringslice.MakeLenPtr(4)
	actual := args.Map{"len": len(s)}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "MakeLenPtr", actual)
}

func Test_Cov11_MakeDefault(t *testing.T) {
	s := stringslice.MakeDefault(10)
	actual := args.Map{"len": len(s), "cap": cap(s)}
	expected := args.Map{"len": 0, "cap": 10}
	expected.ShouldBeEqual(t, 0, "MakeDefault", actual)
}

func Test_Cov11_MakeDefaultPtr(t *testing.T) {
	s := stringslice.MakeDefaultPtr(10)
	actual := args.Map{"len": len(s), "cap": cap(s)}
	expected := args.Map{"len": 0, "cap": 10}
	expected.ShouldBeEqual(t, 0, "MakeDefaultPtr", actual)
}

// ── Clone / ClonePtr / CloneUsingCap ──

func Test_Cov11_Clone_NonEmpty(t *testing.T) {
	original := []string{"a", "b"}
	cloned := stringslice.Clone(original)
	original[0] = "X"
	actual := args.Map{"len": len(cloned), "first": cloned[0]}
	expected := args.Map{"len": 2, "first": "a"}
	expected.ShouldBeEqual(t, 0, "Clone non-empty", actual)
}

func Test_Cov11_Clone_Empty(t *testing.T) {
	cloned := stringslice.Clone(nil)
	actual := args.Map{"len": len(cloned)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clone nil", actual)
}

func Test_Cov11_ClonePtr(t *testing.T) {
	original := []string{"a", "b"}
	cloned := stringslice.ClonePtr(original)
	original[0] = "X"
	actual := args.Map{"len": len(cloned), "first": cloned[0]}
	expected := args.Map{"len": 2, "first": "a"}
	expected.ShouldBeEqual(t, 0, "ClonePtr", actual)
}

func Test_Cov11_ClonePtr_Empty(t *testing.T) {
	cloned := stringslice.ClonePtr(nil)
	actual := args.Map{"len": len(cloned)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ClonePtr nil", actual)
}

func Test_Cov11_CloneUsingCap_NonEmpty(t *testing.T) {
	cloned := stringslice.CloneUsingCap(5, []string{"a", "b"})
	actual := args.Map{"len": len(cloned), "capGt": cap(cloned) >= 7}
	expected := args.Map{"len": 2, "capGt": true}
	expected.ShouldBeEqual(t, 0, "CloneUsingCap non-empty", actual)
}

func Test_Cov11_CloneUsingCap_Empty(t *testing.T) {
	cloned := stringslice.CloneUsingCap(5, nil)
	actual := args.Map{"len": len(cloned)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CloneUsingCap empty", actual)
}

// ── CloneIf ──

func Test_Cov11_CloneIf_True(t *testing.T) {
	original := []string{"a"}
	cloned := stringslice.CloneIf(true, 0, original)
	original[0] = "X"
	actual := args.Map{"first": cloned[0]}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "CloneIf true clones", actual)
}

func Test_Cov11_CloneIf_False_NonNil(t *testing.T) {
	original := []string{"a"}
	result := stringslice.CloneIf(false, 0, original)
	actual := args.Map{"first": result[0]}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "CloneIf false returns same", actual)
}

func Test_Cov11_CloneIf_False_Nil(t *testing.T) {
	result := stringslice.CloneIf(false, 0, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CloneIf false nil", actual)
}

// ── CloneSimpleSliceToPointers ──

func Test_Cov11_CloneSimpleSliceToPointers_NonEmpty(t *testing.T) {
	result := stringslice.CloneSimpleSliceToPointers([]string{"a", "b"})
	actual := args.Map{"len": len(*result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CloneSimpleSliceToPointers non-empty", actual)
}

func Test_Cov11_CloneSimpleSliceToPointers_Empty(t *testing.T) {
	result := stringslice.CloneSimpleSliceToPointers(nil)
	actual := args.Map{"len": len(*result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CloneSimpleSliceToPointers empty", actual)
}

// ── JoinWith / Joins ──

func Test_Cov11_JoinWith_NonEmpty(t *testing.T) {
	result := stringslice.JoinWith(", ", "a", "b")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ", a, b"}
	expected.ShouldBeEqual(t, 0, "JoinWith non-empty", actual)
}

func Test_Cov11_JoinWith_Empty(t *testing.T) {
	result := stringslice.JoinWith(", ")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "JoinWith empty", actual)
}

func Test_Cov11_Joins_NonEmpty(t *testing.T) {
	result := stringslice.Joins(", ", "a", "b")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "a, b"}
	expected.ShouldBeEqual(t, 0, "Joins non-empty", actual)
}

func Test_Cov11_Joins_Empty(t *testing.T) {
	result := stringslice.Joins(", ")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "Joins empty", actual)
}

// ── First / FirstPtr / FirstOrDefault / FirstOrDefaultPtr / FirstOrDefaultWith ──

func Test_Cov11_First(t *testing.T) {
	actual := args.Map{"val": stringslice.First([]string{"a", "b"})}
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "First", actual)
}

func Test_Cov11_FirstPtr(t *testing.T) {
	actual := args.Map{"val": stringslice.FirstPtr([]string{"a", "b"})}
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "FirstPtr", actual)
}

func Test_Cov11_FirstOrDefault_NonEmpty(t *testing.T) {
	actual := args.Map{"val": stringslice.FirstOrDefault([]string{"a"})}
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault non-empty", actual)
}

func Test_Cov11_FirstOrDefault_Empty(t *testing.T) {
	actual := args.Map{"val": stringslice.FirstOrDefault(nil)}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault empty", actual)
}

func Test_Cov11_FirstOrDefaultPtr_NonEmpty(t *testing.T) {
	actual := args.Map{"val": stringslice.FirstOrDefaultPtr([]string{"a"})}
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultPtr non-empty", actual)
}

func Test_Cov11_FirstOrDefaultPtr_Empty(t *testing.T) {
	actual := args.Map{"val": stringslice.FirstOrDefaultPtr(nil)}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultPtr empty", actual)
}

func Test_Cov11_FirstOrDefaultWith_Found(t *testing.T) {
	result, ok := stringslice.FirstOrDefaultWith([]string{"a"}, "def")
	actual := args.Map{"val": result, "ok": ok}
	expected := args.Map{"val": "a", "ok": true}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultWith found", actual)
}

func Test_Cov11_FirstOrDefaultWith_NotFound(t *testing.T) {
	result, ok := stringslice.FirstOrDefaultWith(nil, "def")
	actual := args.Map{"val": result, "ok": ok}
	expected := args.Map{"val": "def", "ok": false}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultWith not found", actual)
}

// ── Last / LastPtr / LastOrDefault / LastOrDefaultPtr ──

func Test_Cov11_Last(t *testing.T) {
	actual := args.Map{"val": stringslice.Last([]string{"a", "b"})}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "Last", actual)
}

func Test_Cov11_LastPtr(t *testing.T) {
	actual := args.Map{"val": stringslice.LastPtr([]string{"a", "b"})}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "LastPtr", actual)
}

func Test_Cov11_LastOrDefault_NonEmpty(t *testing.T) {
	actual := args.Map{"val": stringslice.LastOrDefault([]string{"a", "b"})}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "LastOrDefault non-empty", actual)
}

func Test_Cov11_LastOrDefault_Empty(t *testing.T) {
	actual := args.Map{"val": stringslice.LastOrDefault(nil)}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "LastOrDefault empty", actual)
}

func Test_Cov11_LastOrDefaultPtr_NonEmpty(t *testing.T) {
	actual := args.Map{"val": stringslice.LastOrDefaultPtr([]string{"a", "b"})}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "LastOrDefaultPtr non-empty", actual)
}

func Test_Cov11_LastOrDefaultPtr_Empty(t *testing.T) {
	actual := args.Map{"val": stringslice.LastOrDefaultPtr(nil)}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "LastOrDefaultPtr empty", actual)
}

// ── LastIndexPtr / LastSafeIndexPtr ──

func Test_Cov11_LastIndexPtr(t *testing.T) {
	actual := args.Map{"idx": stringslice.LastIndexPtr([]string{"a", "b", "c"})}
	expected := args.Map{"idx": 2}
	expected.ShouldBeEqual(t, 0, "LastIndexPtr", actual)
}

func Test_Cov11_LastSafeIndexPtr_Valid(t *testing.T) {
	actual := args.Map{"idx": stringslice.LastSafeIndexPtr([]string{"a", "b"})}
	expected := args.Map{"idx": 1}
	expected.ShouldBeEqual(t, 0, "LastSafeIndexPtr valid", actual)
}

func Test_Cov11_LastSafeIndexPtr_Empty(t *testing.T) {
	actual := args.Map{"idx": stringslice.LastSafeIndexPtr(nil)}
	expected := args.Map{"idx": -1}
	expected.ShouldBeEqual(t, 0, "LastSafeIndexPtr empty", actual)
}

// ── FirstLastDefault / FirstLastDefaultPtr ──

func Test_Cov11_FirstLastDefault_Empty(t *testing.T) {
	first, last := stringslice.FirstLastDefault(nil)
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "", "last": ""}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault empty", actual)
}

func Test_Cov11_FirstLastDefault_One(t *testing.T) {
	first, last := stringslice.FirstLastDefault([]string{"a"})
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "a", "last": ""}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault one", actual)
}

func Test_Cov11_FirstLastDefault_Two(t *testing.T) {
	first, last := stringslice.FirstLastDefault([]string{"a", "b"})
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "a", "last": "b"}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault two", actual)
}

func Test_Cov11_FirstLastDefaultPtr_Empty(t *testing.T) {
	first, last := stringslice.FirstLastDefaultPtr(nil)
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "", "last": ""}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultPtr empty", actual)
}

func Test_Cov11_FirstLastDefaultPtr_NonEmpty(t *testing.T) {
	first, last := stringslice.FirstLastDefaultPtr([]string{"a", "b"})
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "a", "last": "b"}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultPtr non-empty", actual)
}

// ── FirstLastDefaultStatus / FirstLastDefaultStatusPtr / FirstLastStatus ──

func Test_Cov11_FirstLastDefaultStatus_Empty(t *testing.T) {
	result := stringslice.FirstLastDefaultStatus(nil)
	actual := args.Map{"valid": result.IsValid, "hasFirst": result.HasFirst, "hasLast": result.HasLast}
	expected := args.Map{"valid": false, "hasFirst": false, "hasLast": false}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus empty", actual)
}

func Test_Cov11_FirstLastDefaultStatus_One(t *testing.T) {
	result := stringslice.FirstLastDefaultStatus([]string{"a"})
	actual := args.Map{"valid": result.IsValid, "hasFirst": result.HasFirst, "hasLast": result.HasLast, "first": result.First}
	expected := args.Map{"valid": false, "hasFirst": true, "hasLast": false, "first": "a"}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus one", actual)
}

func Test_Cov11_FirstLastDefaultStatus_Two(t *testing.T) {
	result := stringslice.FirstLastDefaultStatus([]string{"a", "b"})
	actual := args.Map{"valid": result.IsValid, "hasFirst": result.HasFirst, "hasLast": result.HasLast}
	expected := args.Map{"valid": true, "hasFirst": true, "hasLast": true}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus two", actual)
}

func Test_Cov11_FirstLastDefaultStatusPtr_Empty(t *testing.T) {
	result := stringslice.FirstLastDefaultStatusPtr(nil)
	actual := args.Map{"valid": result.IsValid}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatusPtr empty", actual)
}

func Test_Cov11_FirstLastDefaultStatusPtr_NonEmpty(t *testing.T) {
	result := stringslice.FirstLastDefaultStatusPtr([]string{"a", "b"})
	actual := args.Map{"valid": result.IsValid}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatusPtr non-empty", actual)
}

func Test_Cov11_InvalidFirstLastStatus(t *testing.T) {
	result := stringslice.InvalidFirstLastStatus()
	actual := args.Map{"valid": result.IsValid, "first": result.First, "last": result.Last}
	expected := args.Map{"valid": false, "first": "", "last": ""}
	expected.ShouldBeEqual(t, 0, "InvalidFirstLastStatus", actual)
}

func Test_Cov11_InvalidFirstLastStatusForLast(t *testing.T) {
	result := stringslice.InvalidFirstLastStatusForLast("x")
	actual := args.Map{"hasFirst": result.HasFirst, "hasLast": result.HasLast, "first": result.First}
	expected := args.Map{"hasFirst": true, "hasLast": false, "first": "x"}
	expected.ShouldBeEqual(t, 0, "InvalidFirstLastStatusForLast", actual)
}

// ── IndexAt ──

func Test_Cov11_IndexAt(t *testing.T) {
	actual := args.Map{"val": stringslice.IndexAt([]string{"a", "b", "c"}, 1)}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "IndexAt", actual)
}

// ── IndexesDefault ──

func Test_Cov11_IndexesDefault_NonEmpty(t *testing.T) {
	result := stringslice.IndexesDefault([]string{"a", "b", "c"}, 0, 2)
	actual := args.Map{"len": len(result), "first": result[0], "second": result[1]}
	expected := args.Map{"len": 2, "first": "a", "second": "c"}
	expected.ShouldBeEqual(t, 0, "IndexesDefault non-empty", actual)
}

func Test_Cov11_IndexesDefault_Empty(t *testing.T) {
	result := stringslice.IndexesDefault(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "IndexesDefault empty", actual)
}

// ── InvalidIndexValuesDetail ──

func Test_Cov11_InvalidIndexValuesDetail(t *testing.T) {
	result := stringslice.InvalidIndexValuesDetail()
	actual := args.Map{"valid": result.IsValid, "missing": result.IsAnyMissing}
	expected := args.Map{"valid": false, "missing": true}
	expected.ShouldBeEqual(t, 0, "InvalidIndexValuesDetail", actual)
}

// ── SafeIndexAt / SafeIndexAtWith / SafeIndexAtUsingLastIndex ──

func Test_Cov11_SafeIndexAt_Valid(t *testing.T) {
	actual := args.Map{"val": stringslice.SafeIndexAt([]string{"a", "b"}, 1)}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt valid", actual)
}

func Test_Cov11_SafeIndexAt_OutOfRange(t *testing.T) {
	actual := args.Map{"val": stringslice.SafeIndexAt([]string{"a"}, 5)}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt out of range", actual)
}

func Test_Cov11_SafeIndexAt_Negative(t *testing.T) {
	actual := args.Map{"val": stringslice.SafeIndexAt([]string{"a"}, -1)}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt negative", actual)
}

func Test_Cov11_SafeIndexAt_Empty(t *testing.T) {
	actual := args.Map{"val": stringslice.SafeIndexAt(nil, 0)}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt empty", actual)
}

func Test_Cov11_SafeIndexAtWith_Valid(t *testing.T) {
	actual := args.Map{"val": stringslice.SafeIndexAtWith([]string{"a", "b"}, 1, "def")}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWith valid", actual)
}

func Test_Cov11_SafeIndexAtWith_OutOfRange(t *testing.T) {
	actual := args.Map{"val": stringslice.SafeIndexAtWith([]string{"a"}, 5, "def")}
	expected := args.Map{"val": "def"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWith out of range", actual)
}

func Test_Cov11_SafeIndexAtWith_Negative(t *testing.T) {
	actual := args.Map{"val": stringslice.SafeIndexAtWith([]string{"a"}, -1, "def")}
	expected := args.Map{"val": "def"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWith negative", actual)
}

func Test_Cov11_SafeIndexAtWithPtr_Valid(t *testing.T) {
	actual := args.Map{"val": stringslice.SafeIndexAtWithPtr([]string{"a", "b"}, 1, "def")}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWithPtr valid", actual)
}

func Test_Cov11_SafeIndexAtWithPtr_OutOfRange(t *testing.T) {
	actual := args.Map{"val": stringslice.SafeIndexAtWithPtr(nil, 0, "def")}
	expected := args.Map{"val": "def"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWithPtr out of range", actual)
}

func Test_Cov11_SafeIndexAtUsingLastIndex_Valid(t *testing.T) {
	actual := args.Map{"val": stringslice.SafeIndexAtUsingLastIndex([]string{"a", "b"}, 1, 0)}
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex valid", actual)
}

func Test_Cov11_SafeIndexAtUsingLastIndex_Invalid(t *testing.T) {
	actual := args.Map{"val": stringslice.SafeIndexAtUsingLastIndex([]string{"a"}, 0, 0)}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex lastIndex=0", actual)
}

func Test_Cov11_SafeIndexAtUsingLastIndex_NegativeIndex(t *testing.T) {
	actual := args.Map{"val": stringslice.SafeIndexAtUsingLastIndex([]string{"a"}, 1, -1)}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex negative index", actual)
}

func Test_Cov11_SafeIndexAtUsingLastIndex_IndexGtLastIndex(t *testing.T) {
	actual := args.Map{"val": stringslice.SafeIndexAtUsingLastIndex([]string{"a"}, 1, 5)}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex index > lastIndex", actual)
}

func Test_Cov11_SafeIndexAtUsingLastIndexPtr_Valid(t *testing.T) {
	actual := args.Map{"val": stringslice.SafeIndexAtUsingLastIndexPtr([]string{"a", "b"}, 1, 0)}
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndexPtr valid", actual)
}

func Test_Cov11_SafeIndexAtUsingLastIndexPtr_LastIndexZero(t *testing.T) {
	actual := args.Map{"val": stringslice.SafeIndexAtUsingLastIndexPtr([]string{"a"}, 0, 0)}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndexPtr lastIndex=0", actual)
}

func Test_Cov11_SafeIndexAtUsingLastIndexPtr_NegativeLastIndex(t *testing.T) {
	actual := args.Map{"val": stringslice.SafeIndexAtUsingLastIndexPtr([]string{"a"}, -1, 0)}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndexPtr negative lastIndex", actual)
}

// ── SafeIndexes / SafeIndexesPtr / SafeIndexesDefaultWithDetail ──

func Test_Cov11_SafeIndexes_Valid(t *testing.T) {
	result := stringslice.SafeIndexes([]string{"a", "b", "c"}, 0, 2)
	actual := args.Map{"first": result[0], "second": result[1]}
	expected := args.Map{"first": "a", "second": "c"}
	expected.ShouldBeEqual(t, 0, "SafeIndexes valid", actual)
}

func Test_Cov11_SafeIndexes_OutOfRange(t *testing.T) {
	result := stringslice.SafeIndexes([]string{"a"}, 0, 5)
	actual := args.Map{"first": result[0], "second": result[1]}
	expected := args.Map{"first": "a", "second": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexes out of range", actual)
}

func Test_Cov11_SafeIndexes_EmptySlice(t *testing.T) {
	result := stringslice.SafeIndexes(nil, 0)
	actual := args.Map{"first": result[0]}
	expected := args.Map{"first": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexes empty slice", actual)
}

func Test_Cov11_SafeIndexesPtr_NonEmpty(t *testing.T) {
	result := stringslice.SafeIndexesPtr([]string{"a", "b"}, 0, 1)
	actual := args.Map{"first": result[0], "second": result[1]}
	expected := args.Map{"first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexesPtr non-empty", actual)
}

func Test_Cov11_SafeIndexesPtr_Empty(t *testing.T) {
	result := stringslice.SafeIndexesPtr(nil, 0)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SafeIndexesPtr empty", actual)
}

func Test_Cov11_SafeIndexesDefaultWithDetail_Valid(t *testing.T) {
	result := stringslice.SafeIndexesDefaultWithDetail([]string{"a", "b", "c"}, 0, 2)
	actual := args.Map{"valid": result.IsValid, "missing": result.IsAnyMissing, "len": len(result.Values)}
	expected := args.Map{"valid": true, "missing": false, "len": 2}
	expected.ShouldBeEqual(t, 0, "SafeIndexesDefaultWithDetail valid", actual)
}

func Test_Cov11_SafeIndexesDefaultWithDetail_WithMissing(t *testing.T) {
	result := stringslice.SafeIndexesDefaultWithDetail([]string{"a"}, 0, 5)
	actual := args.Map{"valid": result.IsValid, "missing": result.IsAnyMissing, "missingLen": len(result.MissingIndexes)}
	expected := args.Map{"valid": true, "missing": true, "missingLen": 1}
	expected.ShouldBeEqual(t, 0, "SafeIndexesDefaultWithDetail with missing", actual)
}

func Test_Cov11_SafeIndexesDefaultWithDetail_Empty(t *testing.T) {
	result := stringslice.SafeIndexesDefaultWithDetail(nil, 0)
	actual := args.Map{"valid": result.IsValid}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "SafeIndexesDefaultWithDetail empty", actual)
}

// ── SafeIndexRanges ──

func Test_Cov11_SafeIndexRanges_Valid(t *testing.T) {
	result := stringslice.SafeIndexRanges([]string{"a", "b", "c", "d"}, 1, 3)
	actual := args.Map{"len": len(result), "first": result[0], "last": result[2]}
	expected := args.Map{"len": 3, "first": "b", "last": "d"}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges valid", actual)
}

func Test_Cov11_SafeIndexRanges_NegativeRange(t *testing.T) {
	result := stringslice.SafeIndexRanges([]string{"a"}, 5, 2)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges negative range", actual)
}

func Test_Cov11_SafeIndexRanges_OutOfRange(t *testing.T) {
	result := stringslice.SafeIndexRanges([]string{"a", "b"}, 0, 5)
	actual := args.Map{"len": len(result), "first": result[0], "second": result[1]}
	expected := args.Map{"len": 6, "first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges out of range pads empty", actual)
}

func Test_Cov11_SafeIndexRanges_EmptySlice(t *testing.T) {
	result := stringslice.SafeIndexRanges(nil, 0, 0)
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 1, "first": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges empty slice", actual)
}

// ── SafeRangeItems / SafeRangeItemsPtr ──

func Test_Cov11_SafeRangeItems_Valid(t *testing.T) {
	result := stringslice.SafeRangeItems([]string{"a", "b", "c"}, 0, 2)
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "a"}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems valid", actual)
}

func Test_Cov11_SafeRangeItems_Nil(t *testing.T) {
	result := stringslice.SafeRangeItems(nil, 0, 2)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems nil", actual)
}

func Test_Cov11_SafeRangeItems_Empty(t *testing.T) {
	result := stringslice.SafeRangeItems([]string{}, 0, 2)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems empty", actual)
}

func Test_Cov11_SafeRangeItems_StartGtLastIndex(t *testing.T) {
	result := stringslice.SafeRangeItems([]string{"a"}, 5, 10)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems start > lastIndex", actual)
}

func Test_Cov11_SafeRangeItems_EndGtLastIndex(t *testing.T) {
	result := stringslice.SafeRangeItems([]string{"a", "b", "c"}, 0, -1)
	actual := args.Map{"len": len(result), "first": result[0], "second": result[1]}
	expected := args.Map{"len": 2, "first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems end=-1 clips to lastIndex", actual)
}

func Test_Cov11_SafeRangeItems_StartInvalid(t *testing.T) {
	result := stringslice.SafeRangeItems([]string{"a", "b", "c"}, -1, 2)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems start=-1 uses [:end]", actual)
}

func Test_Cov11_SafeRangeItemsPtr_NonEmpty(t *testing.T) {
	result := stringslice.SafeRangeItemsPtr([]string{"a", "b"}, 0, 1)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SafeRangeItemsPtr non-empty", actual)
}

func Test_Cov11_SafeRangeItemsPtr_Empty(t *testing.T) {
	result := stringslice.SafeRangeItemsPtr(nil, 0, 1)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeRangeItemsPtr empty", actual)
}

// ── SlicePtr ──

func Test_Cov11_SlicePtr_NonEmpty(t *testing.T) {
	result := stringslice.SlicePtr([]string{"a"})
	actual := args.Map{"first": result[0]}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "SlicePtr non-empty", actual)
}

func Test_Cov11_SlicePtr_Empty(t *testing.T) {
	result := stringslice.SlicePtr(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SlicePtr empty", actual)
}

// ── InPlaceReverse ──

func Test_Cov11_InPlaceReverse_Nil(t *testing.T) {
	result := stringslice.InPlaceReverse(nil)
	actual := args.Map{"len": len(*result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse nil", actual)
}

func Test_Cov11_InPlaceReverse_Single(t *testing.T) {
	s := []string{"a"}
	result := stringslice.InPlaceReverse(&s)
	actual := args.Map{"first": (*result)[0]}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse single", actual)
}

func Test_Cov11_InPlaceReverse_Two(t *testing.T) {
	s := []string{"a", "b"}
	result := stringslice.InPlaceReverse(&s)
	actual := args.Map{"first": (*result)[0], "second": (*result)[1]}
	expected := args.Map{"first": "b", "second": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse two", actual)
}

func Test_Cov11_InPlaceReverse_Three(t *testing.T) {
	s := []string{"a", "b", "c"}
	result := stringslice.InPlaceReverse(&s)
	actual := args.Map{"first": (*result)[0], "last": (*result)[2]}
	expected := args.Map{"first": "c", "last": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse three", actual)
}

// ── MergeNew / MergeNewSimple / MergeSlicesOfSlices ──

func Test_Cov11_MergeNew(t *testing.T) {
	result := stringslice.MergeNew([]string{"a"}, "b", "c")
	actual := args.Map{"len": len(result), "first": result[0], "last": result[2]}
	expected := args.Map{"len": 3, "first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "MergeNew", actual)
}

func Test_Cov11_MergeNew_EmptyFirst(t *testing.T) {
	result := stringslice.MergeNew(nil, "b")
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 1, "first": "b"}
	expected.ShouldBeEqual(t, 0, "MergeNew empty first", actual)
}

func Test_Cov11_MergeNew_EmptyAdditional(t *testing.T) {
	result := stringslice.MergeNew([]string{"a"})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 1, "first": "a"}
	expected.ShouldBeEqual(t, 0, "MergeNew empty additional", actual)
}

func Test_Cov11_MergeNewSimple(t *testing.T) {
	result := stringslice.MergeNewSimple([]string{"a"}, []string{"b"}, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple", actual)
}

func Test_Cov11_MergeNewSimple_Empty(t *testing.T) {
	result := stringslice.MergeNewSimple()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple empty", actual)
}

func Test_Cov11_MergeSlicesOfSlices(t *testing.T) {
	result := stringslice.MergeSlicesOfSlices([]string{"a"}, nil, []string{"b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices", actual)
}

func Test_Cov11_MergeSlicesOfSlices_Empty(t *testing.T) {
	result := stringslice.MergeSlicesOfSlices()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices empty", actual)
}

func Test_Cov11_AllElemLengthSlices(t *testing.T) {
	result := stringslice.AllElemLengthSlices([]string{"a", "b"}, nil, []string{"c"})
	actual := args.Map{"count": result}
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "AllElemLengthSlices", actual)
}

func Test_Cov11_AllElemLengthSlices_Empty(t *testing.T) {
	result := stringslice.AllElemLengthSlices()
	actual := args.Map{"count": result}
	expected := args.Map{"count": 0}
	expected.ShouldBeEqual(t, 0, "AllElemLengthSlices empty", actual)
}

// ── AppendLineNew / PrependLineNew / PrependNew ──

func Test_Cov11_AppendLineNew(t *testing.T) {
	result := stringslice.AppendLineNew([]string{"a"}, "b")
	actual := args.Map{"len": len(result), "last": result[1]}
	expected := args.Map{"len": 2, "last": "b"}
	expected.ShouldBeEqual(t, 0, "AppendLineNew", actual)
}

func Test_Cov11_PrependLineNew(t *testing.T) {
	result := stringslice.PrependLineNew("x", []string{"a"})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "x"}
	expected.ShouldBeEqual(t, 0, "PrependLineNew", actual)
}

func Test_Cov11_PrependNew(t *testing.T) {
	result := stringslice.PrependNew([]string{"a"}, "x", "y")
	actual := args.Map{"len": len(*result), "first": (*result)[0], "last": (*result)[2]}
	expected := args.Map{"len": 3, "first": "x", "last": "a"}
	expected.ShouldBeEqual(t, 0, "PrependNew", actual)
}

func Test_Cov11_PrependNew_EmptyPrepend(t *testing.T) {
	result := stringslice.PrependNew([]string{"a"})
	actual := args.Map{"len": len(*result), "first": (*result)[0]}
	expected := args.Map{"len": 1, "first": "a"}
	expected.ShouldBeEqual(t, 0, "PrependNew empty prepend", actual)
}

func Test_Cov11_PrependNew_EmptySlice(t *testing.T) {
	result := stringslice.PrependNew(nil, "x")
	actual := args.Map{"len": len(*result), "first": (*result)[0]}
	expected := args.Map{"len": 1, "first": "x"}
	expected.ShouldBeEqual(t, 0, "PrependNew empty slice", actual)
}

// ── AppendAnyItemsWithStrings ──

func Test_Cov11_AppendAnyItemsWithStrings(t *testing.T) {
	result := stringslice.AppendAnyItemsWithStrings(false, false, []string{"a"}, "b", 42)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings", actual)
}

func Test_Cov11_AppendAnyItemsWithStrings_SkipEmpty(t *testing.T) {
	result := stringslice.AppendAnyItemsWithStrings(false, true, []string{"a"}, nil, "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings skip nil", actual)
}

func Test_Cov11_AppendAnyItemsWithStrings_NoAdditional(t *testing.T) {
	result := stringslice.AppendAnyItemsWithStrings(false, false, []string{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings no additional", actual)
}

// ── AppendStringsWithAnyItems ──

func Test_Cov11_AppendStringsWithAnyItems(t *testing.T) {
	result := stringslice.AppendStringsWithAnyItems(false, false, []any{"a"}, "b", "c")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems", actual)
}

func Test_Cov11_AppendStringsWithAnyItems_SkipEmpty(t *testing.T) {
	result := stringslice.AppendStringsWithAnyItems(false, true, []any{"a"}, "", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems skip empty", actual)
}

func Test_Cov11_AppendStringsWithAnyItems_NoAdditional(t *testing.T) {
	result := stringslice.AppendStringsWithAnyItems(false, false, []any{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems no additional", actual)
}

// ── AppendStringsWithMainSlice ──

func Test_Cov11_AppendStringsWithMainSlice(t *testing.T) {
	main := []string{"a"}
	result := stringslice.AppendStringsWithMainSlice(false, main, "b", "c")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice", actual)
}

func Test_Cov11_AppendStringsWithMainSlice_SkipEmpty(t *testing.T) {
	main := []string{"a"}
	result := stringslice.AppendStringsWithMainSlice(true, main, "", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice skip empty", actual)
}

func Test_Cov11_AppendStringsWithMainSlice_NoAdditional(t *testing.T) {
	main := []string{"a"}
	result := stringslice.AppendStringsWithMainSlice(false, main)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice no additional", actual)
}

// ── AnyItemsCloneIf / AnyItemsCloneUsingCap ──

func Test_Cov11_AnyItemsCloneIf_True(t *testing.T) {
	result := stringslice.AnyItemsCloneIf(true, 0, []any{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf true", actual)
}

func Test_Cov11_AnyItemsCloneIf_False_NonNil(t *testing.T) {
	original := []any{"a"}
	result := stringslice.AnyItemsCloneIf(false, 0, original)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf false non-nil", actual)
}

func Test_Cov11_AnyItemsCloneIf_False_Nil(t *testing.T) {
	result := stringslice.AnyItemsCloneIf(false, 0, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf false nil", actual)
}

func Test_Cov11_AnyItemsCloneUsingCap_NonEmpty(t *testing.T) {
	result := stringslice.AnyItemsCloneUsingCap(5, []any{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneUsingCap non-empty", actual)
}

func Test_Cov11_AnyItemsCloneUsingCap_Empty(t *testing.T) {
	result := stringslice.AnyItemsCloneUsingCap(5, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneUsingCap empty", actual)
}

// ── NonEmpty / NonEmptySlicePtr / NonEmptyStrings / NonNullStrings ──

func Test_Cov11_NonEmptySlice(t *testing.T) {
	result := stringslice.NonEmptySlice([]string{"a", "", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice", actual)
}

func Test_Cov11_NonEmptySlice_Empty(t *testing.T) {
	result := stringslice.NonEmptySlice(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice empty", actual)
}

func Test_Cov11_NonEmptySlicePtr(t *testing.T) {
	result := stringslice.NonEmptySlicePtr([]string{"a", "", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptySlicePtr", actual)
}

func Test_Cov11_NonEmptySlicePtr_Empty(t *testing.T) {
	result := stringslice.NonEmptySlicePtr(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptySlicePtr empty", actual)
}

func Test_Cov11_NonEmptyStrings(t *testing.T) {
	result := stringslice.NonEmptyStrings([]string{"a", "", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings", actual)
}

func Test_Cov11_NonEmptyStrings_Nil(t *testing.T) {
	result := stringslice.NonEmptyStrings(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings nil", actual)
}

func Test_Cov11_NonEmptyStrings_Empty(t *testing.T) {
	result := stringslice.NonEmptyStrings([]string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings empty", actual)
}

func Test_Cov11_NonNullStrings(t *testing.T) {
	result := stringslice.NonNullStrings([]string{"a", "", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonNullStrings", actual)
}

func Test_Cov11_NonNullStrings_Nil(t *testing.T) {
	result := stringslice.NonNullStrings(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonNullStrings nil", actual)
}

// ── NonEmptyIf ──

func Test_Cov11_NonEmptyIf_True(t *testing.T) {
	result := stringslice.NonEmptyIf(true, []string{"a", "", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptyIf true", actual)
}

func Test_Cov11_NonEmptyIf_False(t *testing.T) {
	result := stringslice.NonEmptyIf(false, []string{"a", "", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptyIf false uses NonNullStrings", actual)
}

// ── NonEmptyJoin / NonEmptyJoinPtr ──

func Test_Cov11_NonEmptyJoin(t *testing.T) {
	result := stringslice.NonEmptyJoin([]string{"a", "", "b"}, ", ")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "a, b"}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin", actual)
}

func Test_Cov11_NonEmptyJoin_Nil(t *testing.T) {
	result := stringslice.NonEmptyJoin(nil, ", ")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin nil", actual)
}

func Test_Cov11_NonEmptyJoin_Empty(t *testing.T) {
	result := stringslice.NonEmptyJoin([]string{}, ", ")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin empty", actual)
}

func Test_Cov11_NonEmptyJoinPtr(t *testing.T) {
	result := stringslice.NonEmptyJoinPtr([]string{"a", "", "b"}, ", ")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "a, b"}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoinPtr", actual)
}

func Test_Cov11_NonEmptyJoinPtr_Empty(t *testing.T) {
	result := stringslice.NonEmptyJoinPtr(nil, ", ")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoinPtr empty", actual)
}

// ── NonWhitespace / NonWhitespacePtr / NonWhitespaceJoin / NonWhitespaceJoinPtr ──

func Test_Cov11_NonWhitespace(t *testing.T) {
	result := stringslice.NonWhitespace([]string{"a", "  ", "", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonWhitespace", actual)
}

func Test_Cov11_NonWhitespace_Nil(t *testing.T) {
	result := stringslice.NonWhitespace(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespace nil", actual)
}

func Test_Cov11_NonWhitespace_Empty(t *testing.T) {
	result := stringslice.NonWhitespace([]string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespace empty", actual)
}

func Test_Cov11_NonWhitespacePtr(t *testing.T) {
	result := stringslice.NonWhitespacePtr([]string{"a", "  ", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonWhitespacePtr", actual)
}

func Test_Cov11_NonWhitespacePtr_Empty(t *testing.T) {
	result := stringslice.NonWhitespacePtr(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespacePtr empty", actual)
}

func Test_Cov11_NonWhitespaceJoin(t *testing.T) {
	result := stringslice.NonWhitespaceJoin([]string{"a", "  ", "b"}, ", ")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "a, b"}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin", actual)
}

func Test_Cov11_NonWhitespaceJoin_Nil(t *testing.T) {
	result := stringslice.NonWhitespaceJoin(nil, ", ")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin nil", actual)
}

func Test_Cov11_NonWhitespaceJoin_Empty(t *testing.T) {
	result := stringslice.NonWhitespaceJoin([]string{}, ", ")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin empty", actual)
}

func Test_Cov11_NonWhitespaceJoinPtr(t *testing.T) {
	result := stringslice.NonWhitespaceJoinPtr([]string{"a", "  ", "b"}, ", ")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "a, b"}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoinPtr", actual)
}

func Test_Cov11_NonWhitespaceJoinPtr_Empty(t *testing.T) {
	result := stringslice.NonWhitespaceJoinPtr(nil, ", ")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoinPtr empty", actual)
}

// ── TrimmedEachWords / TrimmedEachWordsPtr / TrimmedEachWordsIf ──

func Test_Cov11_TrimmedEachWords(t *testing.T) {
	result := stringslice.TrimmedEachWords([]string{"  a  ", " ", "  b  "})
	actual := args.Map{"len": len(result), "first": result[0], "second": result[1]}
	expected := args.Map{"len": 2, "first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWords", actual)
}

func Test_Cov11_TrimmedEachWords_Nil(t *testing.T) {
	result := stringslice.TrimmedEachWords(nil)
	isNil := result == nil
	actual := args.Map{"isNil": isNil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWords nil returns nil", actual)
}

func Test_Cov11_TrimmedEachWords_Empty(t *testing.T) {
	result := stringslice.TrimmedEachWords([]string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWords empty", actual)
}

func Test_Cov11_TrimmedEachWordsPtr(t *testing.T) {
	result := stringslice.TrimmedEachWordsPtr([]string{"  a  "})
	actual := args.Map{"first": result[0]}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsPtr", actual)
}

func Test_Cov11_TrimmedEachWordsPtr_Empty(t *testing.T) {
	result := stringslice.TrimmedEachWordsPtr(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsPtr empty", actual)
}

func Test_Cov11_TrimmedEachWordsIf_True(t *testing.T) {
	result := stringslice.TrimmedEachWordsIf(true, []string{"  a  ", " ", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsIf true", actual)
}

func Test_Cov11_TrimmedEachWordsIf_False(t *testing.T) {
	result := stringslice.TrimmedEachWordsIf(false, []string{"a", "", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsIf false uses NonNullStrings", actual)
}

// ── SortIf ──

func Test_Cov11_SortIf_True(t *testing.T) {
	s := []string{"c", "a", "b"}
	result := stringslice.SortIf(true, s)
	actual := args.Map{"first": result[0]}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "SortIf true sorts", actual)
}

func Test_Cov11_SortIf_False(t *testing.T) {
	s := []string{"c", "a", "b"}
	result := stringslice.SortIf(false, s)
	actual := args.Map{"first": result[0]}
	expected := args.Map{"first": "c"}
	expected.ShouldBeEqual(t, 0, "SortIf false no sort", actual)
}

// ── SplitContentsByWhitespace ──

func Test_Cov11_SplitContentsByWhitespace(t *testing.T) {
	result := stringslice.SplitContentsByWhitespace("hello  world   test")
	actual := args.Map{"len": len(result), "first": result[0], "last": result[2]}
	expected := args.Map{"len": 3, "first": "hello", "last": "test"}
	expected.ShouldBeEqual(t, 0, "SplitContentsByWhitespace", actual)
}

// ── SplitTrimmedNonEmpty / SplitTrimmedNonEmptyAll ──

func Test_Cov11_SplitTrimmedNonEmpty(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmpty(" a , b , c ", ",", -1)
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 3, "first": "a"}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmpty", actual)
}

func Test_Cov11_SplitTrimmedNonEmptyAll(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmptyAll(" a , b , c ", ",")
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 3, "first": "a"}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmptyAll", actual)
}

// ── RegexTrimmedSplitNonEmptyAll ──

func Test_Cov11_RegexTrimmedSplitNonEmptyAll(t *testing.T) {
	re := regexp.MustCompile(`[,;]`)
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, " a , b ; c ")
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 3, "first": "a"}
	expected.ShouldBeEqual(t, 0, "RegexTrimmedSplitNonEmptyAll", actual)
}

// ── ExpandByFunc / ExpandBySplit / ExpandBySplits ──

func Test_Cov11_ExpandByFunc(t *testing.T) {
	result := stringslice.ExpandByFunc([]string{"a,b", "c,d"}, func(line string) []string {
		return stringslice.SplitTrimmedNonEmptyAll(line, ",")
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc", actual)
}

func Test_Cov11_ExpandByFunc_Empty(t *testing.T) {
	result := stringslice.ExpandByFunc(nil, func(line string) []string { return nil })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc empty", actual)
}

func Test_Cov11_ExpandBySplit(t *testing.T) {
	result := stringslice.ExpandBySplit([]string{"a,b", "c"}, ",")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit", actual)
}

func Test_Cov11_ExpandBySplit_Empty(t *testing.T) {
	result := stringslice.ExpandBySplit(nil, ",")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit empty", actual)
}

func Test_Cov11_ExpandBySplits(t *testing.T) {
	result := stringslice.ExpandBySplits([]string{"a,b;c"}, ",", ";")
	actual := args.Map{"lenGt": len(result) > 0}
	expected := args.Map{"lenGt": true}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits", actual)
}

func Test_Cov11_ExpandBySplits_Empty(t *testing.T) {
	result := stringslice.ExpandBySplits(nil, ",")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits empty", actual)
}

func Test_Cov11_ExpandBySplits_NoSplitters(t *testing.T) {
	result := stringslice.ExpandBySplits([]string{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits no splitters", actual)
}

// ── ProcessAsync / ProcessOptionAsync ──

func Test_Cov11_ProcessAsync(t *testing.T) {
	result := stringslice.ProcessAsync(func(index int, item any) string {
		return "x"
	}, "a", "b")
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "x"}
	expected.ShouldBeEqual(t, 0, "ProcessAsync", actual)
}

func Test_Cov11_ProcessAsync_Empty(t *testing.T) {
	result := stringslice.ProcessAsync(func(index int, item any) string { return "" })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ProcessAsync empty", actual)
}

func Test_Cov11_ProcessOptionAsync_SkipEmpty(t *testing.T) {
	result := stringslice.ProcessOptionAsync(true, func(index int, item any) string {
		if index == 0 {
			return ""
		}
		return "x"
	}, "a", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync skip empty", actual)
}

func Test_Cov11_ProcessOptionAsync_ReturnAll(t *testing.T) {
	result := stringslice.ProcessOptionAsync(false, func(index int, item any) string {
		if index == 0 {
			return ""
		}
		return "x"
	}, "a", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync return all", actual)
}

func Test_Cov11_ProcessOptionAsync_Empty(t *testing.T) {
	result := stringslice.ProcessOptionAsync(true, func(index int, item any) string { return "" })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync empty", actual)
}

// ── LinesProcess / LinesSimpleProcess / LinesSimpleProcessNoEmpty / LinesAsyncProcess ──

func Test_Cov11_LinesProcess(t *testing.T) {
	result := stringslice.LinesProcess([]string{"a", "b", "c"}, func(i int, lineIn string) (string, bool, bool) {
		return lineIn + "!", true, false
	})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 3, "first": "a!"}
	expected.ShouldBeEqual(t, 0, "LinesProcess", actual)
}

func Test_Cov11_LinesProcess_WithBreak(t *testing.T) {
	result := stringslice.LinesProcess([]string{"a", "b", "c"}, func(i int, lineIn string) (string, bool, bool) {
		if i == 1 {
			return "", false, true
		}
		return lineIn, true, false
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LinesProcess with break", actual)
}

func Test_Cov11_LinesProcess_SkipItem(t *testing.T) {
	result := stringslice.LinesProcess([]string{"a", "b"}, func(i int, lineIn string) (string, bool, bool) {
		return lineIn, i == 0, false
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LinesProcess skip item", actual)
}

func Test_Cov11_LinesProcess_Empty(t *testing.T) {
	result := stringslice.LinesProcess(nil, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesProcess empty", actual)
}

func Test_Cov11_LinesSimpleProcess(t *testing.T) {
	result := stringslice.LinesSimpleProcess([]string{"a", "b"}, func(lineIn string) string {
		return lineIn + "!"
	})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "a!"}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcess", actual)
}

func Test_Cov11_LinesSimpleProcess_Empty(t *testing.T) {
	result := stringslice.LinesSimpleProcess(nil, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcess empty", actual)
}

func Test_Cov11_LinesSimpleProcessNoEmpty(t *testing.T) {
	result := stringslice.LinesSimpleProcessNoEmpty([]string{"a", "b"}, func(lineIn string) string {
		if lineIn == "b" {
			return ""
		}
		return lineIn
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcessNoEmpty", actual)
}

func Test_Cov11_LinesSimpleProcessNoEmpty_Empty(t *testing.T) {
	result := stringslice.LinesSimpleProcessNoEmpty(nil, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcessNoEmpty empty", actual)
}

func Test_Cov11_LinesAsyncProcess(t *testing.T) {
	result := stringslice.LinesAsyncProcess([]string{"a", "b"}, func(i int, lineIn string) string {
		return lineIn + "!"
	})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "a!"}
	expected.ShouldBeEqual(t, 0, "LinesAsyncProcess", actual)
}

func Test_Cov11_LinesAsyncProcess_Empty(t *testing.T) {
	result := stringslice.LinesAsyncProcess(nil, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesAsyncProcess empty", actual)
}

// ── AnyLinesProcessAsyncUsingProcessor ──

func Test_Cov11_AnyLinesProcessAsync_Slice(t *testing.T) {
	result := stringslice.AnyLinesProcessAsyncUsingProcessor(
		[]string{"a", "b"},
		func(i int, lineIn any) string { return "x" },
	)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync slice", actual)
}

func Test_Cov11_AnyLinesProcessAsync_Nil(t *testing.T) {
	result := stringslice.AnyLinesProcessAsyncUsingProcessor(nil, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync nil", actual)
}

func Test_Cov11_AnyLinesProcessAsync_NotSlice(t *testing.T) {
	result := stringslice.AnyLinesProcessAsyncUsingProcessor("hello", nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync not slice", actual)
}

func Test_Cov11_AnyLinesProcessAsync_EmptySlice(t *testing.T) {
	result := stringslice.AnyLinesProcessAsyncUsingProcessor(
		[]string{},
		func(i int, lineIn any) string { return "x" },
	)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync empty slice", actual)
}

func Test_Cov11_AnyLinesProcessAsync_Array(t *testing.T) {
	result := stringslice.AnyLinesProcessAsyncUsingProcessor(
		[2]string{"a", "b"},
		func(i int, lineIn any) string { return "x" },
	)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync array", actual)
}
