package stringslicetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Empty ──

func Test_Cov7_Empty(t *testing.T) {
	result := stringslice.Empty()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty", actual)
}

// ── IsEmpty ──

func Test_Cov7_IsEmpty_Empty(t *testing.T) {
	actual := args.Map{"result": stringslice.IsEmpty([]string{})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty empty", actual)
}

func Test_Cov7_IsEmpty_NonEmpty(t *testing.T) {
	actual := args.Map{"result": stringslice.IsEmpty([]string{"a"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEmpty non-empty", actual)
}

// ── HasAnyItem ──

func Test_Cov7_HasAnyItem_Empty(t *testing.T) {
	actual := args.Map{"result": stringslice.HasAnyItem([]string{})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasAnyItem empty", actual)
}

func Test_Cov7_HasAnyItem_NonEmpty(t *testing.T) {
	actual := args.Map{"result": stringslice.HasAnyItem([]string{"a"})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem non-empty", actual)
}

// ── First ──

func Test_Cov7_First(t *testing.T) {
	actual := args.Map{"result": stringslice.First([]string{"a", "b"})}
	expected := args.Map{"result": "a"}
	expected.ShouldBeEqual(t, 0, "First", actual)
}

// ── Last ──

func Test_Cov7_Last(t *testing.T) {
	actual := args.Map{"result": stringslice.Last([]string{"a", "b"})}
	expected := args.Map{"result": "b"}
	expected.ShouldBeEqual(t, 0, "Last", actual)
}

// ── IndexAt ──

func Test_Cov7_IndexAt(t *testing.T) {
	actual := args.Map{"result": stringslice.IndexAt([]string{"a", "b", "c"}, 1)}
	expected := args.Map{"result": "b"}
	expected.ShouldBeEqual(t, 0, "IndexAt", actual)
}

// ── SafeIndexAt ──

func Test_Cov7_SafeIndexAt_Valid(t *testing.T) {
	actual := args.Map{"result": stringslice.SafeIndexAt([]string{"a", "b"}, 0)}
	expected := args.Map{"result": "a"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt valid", actual)
}

func Test_Cov7_SafeIndexAt_OutOfRange(t *testing.T) {
	actual := args.Map{"result": stringslice.SafeIndexAt([]string{"a"}, 5)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt out of range", actual)
}

func Test_Cov7_SafeIndexAt_Negative(t *testing.T) {
	actual := args.Map{"result": stringslice.SafeIndexAt([]string{"a"}, -1)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt negative", actual)
}

func Test_Cov7_SafeIndexAt_Empty(t *testing.T) {
	actual := args.Map{"result": stringslice.SafeIndexAt([]string{}, 0)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt empty", actual)
}

// ── MergeNew ──

func Test_Cov7_MergeNew(t *testing.T) {
	result := stringslice.MergeNew([]string{"a"}, "b", "c")
	actual := args.Map{"len": len(result), "first": result[0], "last": result[2]}
	expected := args.Map{"len": 3, "first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "MergeNew", actual)
}

func Test_Cov7_MergeNew_EmptyFirst(t *testing.T) {
	result := stringslice.MergeNew([]string{}, "b")
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 1, "first": "b"}
	expected.ShouldBeEqual(t, 0, "MergeNew empty first", actual)
}

// ── MergeNewSimple ──

func Test_Cov7_MergeNewSimple_Empty(t *testing.T) {
	result := stringslice.MergeNewSimple()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple empty", actual)
}

func Test_Cov7_MergeNewSimple_Multiple(t *testing.T) {
	result := stringslice.MergeNewSimple([]string{"a"}, []string{"b", "c"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple multiple", actual)
}

// ── AppendLineNew ──

func Test_Cov7_AppendLineNew(t *testing.T) {
	result := stringslice.AppendLineNew([]string{"a"}, "b")
	actual := args.Map{"len": len(result), "last": result[1]}
	expected := args.Map{"len": 2, "last": "b"}
	expected.ShouldBeEqual(t, 0, "AppendLineNew", actual)
}

// ── PrependNew ──

func Test_Cov7_PrependNew(t *testing.T) {
	result := stringslice.PrependNew([]string{"c"}, "a", "b")
	actual := args.Map{"len": len(*result), "first": (*result)[0], "last": (*result)[2]}
	expected := args.Map{"len": 3, "first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "PrependNew", actual)
}

// ── InPlaceReverse ──

func Test_Cov7_InPlaceReverse_Nil(t *testing.T) {
	result := stringslice.InPlaceReverse(nil)
	actual := args.Map{"len": len(*result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse nil", actual)
}

func Test_Cov7_InPlaceReverse_Single(t *testing.T) {
	s := []string{"a"}
	result := stringslice.InPlaceReverse(&s)
	actual := args.Map{"first": (*result)[0]}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse single", actual)
}

func Test_Cov7_InPlaceReverse_Two(t *testing.T) {
	s := []string{"a", "b"}
	result := stringslice.InPlaceReverse(&s)
	actual := args.Map{"first": (*result)[0], "second": (*result)[1]}
	expected := args.Map{"first": "b", "second": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse two", actual)
}

func Test_Cov7_InPlaceReverse_Three(t *testing.T) {
	s := []string{"a", "b", "c"}
	result := stringslice.InPlaceReverse(&s)
	actual := args.Map{"first": (*result)[0], "last": (*result)[2]}
	expected := args.Map{"first": "c", "last": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse three", actual)
}

// ── SortIf ──

func Test_Cov7_SortIf_True(t *testing.T) {
	result := stringslice.SortIf(true, []string{"c", "a", "b"})
	actual := args.Map{"first": result[0], "last": result[2]}
	expected := args.Map{"first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "SortIf true", actual)
}

func Test_Cov7_SortIf_False(t *testing.T) {
	result := stringslice.SortIf(false, []string{"c", "a", "b"})
	actual := args.Map{"first": result[0]}
	expected := args.Map{"first": "c"}
	expected.ShouldBeEqual(t, 0, "SortIf false", actual)
}

// ── NonEmptySlice ──

func Test_Cov7_NonEmptySlice_Empty(t *testing.T) {
	result := stringslice.NonEmptySlice([]string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice empty", actual)
}

func Test_Cov7_NonEmptySlice_Mixed(t *testing.T) {
	result := stringslice.NonEmptySlice([]string{"a", "", "b"})
	actual := args.Map{"len": len(result), "first": result[0], "second": result[1]}
	expected := args.Map{"len": 2, "first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice mixed", actual)
}

// ── NonEmptyJoin ──

func Test_Cov7_NonEmptyJoin_Nil(t *testing.T) {
	result := stringslice.NonEmptyJoin(nil, ",")
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin nil", actual)
}

func Test_Cov7_NonEmptyJoin_Empty(t *testing.T) {
	result := stringslice.NonEmptyJoin([]string{}, ",")
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin empty", actual)
}

func Test_Cov7_NonEmptyJoin_WithEmpty(t *testing.T) {
	result := stringslice.NonEmptyJoin([]string{"a", "", "b"}, ",")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin with empty", actual)
}

// ── ExpandBySplit ──

func Test_Cov7_ExpandBySplit_Empty(t *testing.T) {
	result := stringslice.ExpandBySplit([]string{}, ",")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit empty", actual)
}

func Test_Cov7_ExpandBySplit_NonEmpty(t *testing.T) {
	result := stringslice.ExpandBySplit([]string{"a,b", "c"}, ",")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit non-empty", actual)
}

// ── CloneIf ──

func Test_Cov7_CloneIf_Clone(t *testing.T) {
	result := stringslice.CloneIf(true, 0, []string{"a"})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 1, "first": "a"}
	expected.ShouldBeEqual(t, 0, "CloneIf clone", actual)
}

func Test_Cov7_CloneIf_NoClone(t *testing.T) {
	result := stringslice.CloneIf(false, 0, []string{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CloneIf no clone", actual)
}

func Test_Cov7_CloneIf_NilNoClone(t *testing.T) {
	result := stringslice.CloneIf(false, 0, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CloneIf nil no clone", actual)
}

// ── JoinWith ──

func Test_Cov7_JoinWith_Empty(t *testing.T) {
	result := stringslice.JoinWith(",")
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "JoinWith empty", actual)
}

func Test_Cov7_JoinWith_Items(t *testing.T) {
	result := stringslice.JoinWith(",", "a", "b")
	actual := args.Map{"result": result}
	expected := args.Map{"result": ",a,b"}
	expected.ShouldBeEqual(t, 0, "JoinWith items", actual)
}

// ── Joins ──

func Test_Cov7_Joins_Empty(t *testing.T) {
	result := stringslice.Joins(",")
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Joins empty", actual)
}

func Test_Cov7_Joins_Items(t *testing.T) {
	result := stringslice.Joins(",", "a", "b")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "Joins items", actual)
}
