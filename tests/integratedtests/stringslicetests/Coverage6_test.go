package stringslicetests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Clone ──

func Test_Cov6_Clone_NonEmpty(t *testing.T) {
	result := stringslice.Clone([]string{"a", "b"})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "a"}
	expected.ShouldBeEqual(t, 0, "Clone non-empty", actual)
}

func Test_Cov6_Clone_Empty(t *testing.T) {
	result := stringslice.Clone(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clone nil", actual)
}

// ── FirstLastDefault ──

func Test_Cov6_FirstLastDefault_SingleElement(t *testing.T) {
	first, last := stringslice.FirstLastDefault([]string{"only"})
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "only", "last": ""}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault single element", actual)
}

func Test_Cov6_FirstLastDefault_Empty(t *testing.T) {
	first, last := stringslice.FirstLastDefault(nil)
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "", "last": ""}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault empty", actual)
}

func Test_Cov6_FirstLastDefault_Multi(t *testing.T) {
	first, last := stringslice.FirstLastDefault([]string{"a", "b", "c"})
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault multi", actual)
}

// ── SafeIndexAt ──

func Test_Cov6_SafeIndexAt(t *testing.T) {
	slice := []string{"a", "b", "c"}
	actual := args.Map{
		"at0":        stringslice.SafeIndexAt(slice, 0),
		"at2":        stringslice.SafeIndexAt(slice, 2),
		"atNeg":      stringslice.SafeIndexAt(slice, -1),
		"atOutBound": stringslice.SafeIndexAt(slice, 10),
		"emptySlice": stringslice.SafeIndexAt(nil, 0),
	}
	expected := args.Map{
		"at0": "a", "at2": "c",
		"atNeg": "", "atOutBound": "", "emptySlice": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt", actual)
}

// ── NonEmptyJoin ──

func Test_Cov6_NonEmptyJoin(t *testing.T) {
	result := stringslice.NonEmptyJoin([]string{"a", "", "b"}, ",")
	resultNil := stringslice.NonEmptyJoin(nil, ",")
	resultEmpty := stringslice.NonEmptyJoin([]string{}, ",")
	actual := args.Map{
		"result":    result,
		"nilResult": resultNil,
		"emptyRes":  resultEmpty,
	}
	expected := args.Map{
		"result": "a,b", "nilResult": "", "emptyRes": "",
	}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin", actual)
}

// ── MergeNew ──

func Test_Cov6_MergeNew(t *testing.T) {
	result := stringslice.MergeNew([]string{"a"}, "b", "c")
	resultEmpty := stringslice.MergeNew(nil, "b")
	actual := args.Map{
		"len":      len(result),
		"emptyLen": len(resultEmpty),
	}
	expected := args.Map{"len": 3, "emptyLen": 1}
	expected.ShouldBeEqual(t, 0, "MergeNew", actual)
}

// ── InPlaceReverse ──

func Test_Cov6_InPlaceReverse(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}
	result := stringslice.InPlaceReverse(&slice)
	actual := args.Map{"first": (*result)[0], "last": (*result)[3]}
	expected := args.Map{"first": "d", "last": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse 4 elements", actual)
}

func Test_Cov6_InPlaceReverse_Two(t *testing.T) {
	slice := []string{"a", "b"}
	result := stringslice.InPlaceReverse(&slice)
	actual := args.Map{"first": (*result)[0], "last": (*result)[1]}
	expected := args.Map{"first": "b", "last": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse 2 elements", actual)
}

func Test_Cov6_InPlaceReverse_SingleAndNil(t *testing.T) {
	single := []string{"only"}
	r1 := stringslice.InPlaceReverse(&single)
	r2 := stringslice.InPlaceReverse(nil)
	actual := args.Map{
		"singleFirst": (*r1)[0],
		"nilLen":      len(*r2),
	}
	expected := args.Map{"singleFirst": "only", "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse single and nil", actual)
}

// ── SortIf ──

func Test_Cov6_SortIf(t *testing.T) {
	slice := []string{"c", "a", "b"}
	sorted := stringslice.SortIf(true, slice)
	notSorted := stringslice.SortIf(false, []string{"c", "a"})
	actual := args.Map{
		"sortedFirst":    sorted[0],
		"notSortedFirst": notSorted[0],
	}
	expected := args.Map{"sortedFirst": "a", "notSortedFirst": "c"}
	expected.ShouldBeEqual(t, 0, "SortIf", actual)
}

// ── ExpandByFunc ──

func Test_Cov6_ExpandByFunc(t *testing.T) {
	result := stringslice.ExpandByFunc(
		[]string{"a,b", "c,d"},
		func(line string) []string { return strings.Split(line, ",") },
	)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc", actual)
}

func Test_Cov6_ExpandByFunc_Empty(t *testing.T) {
	result := stringslice.ExpandByFunc(nil, func(line string) []string { return nil })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc empty", actual)
}

// ── AllElemLengthSlices ──

func Test_Cov6_AllElemLengthSlices(t *testing.T) {
	result := stringslice.AllElemLengthSlices(
		[]string{"a", "b"},
		nil,
		[]string{"c"},
	)
	actual := args.Map{"total": result}
	expected := args.Map{"total": 3}
	expected.ShouldBeEqual(t, 0, "AllElemLengthSlices", actual)
}

func Test_Cov6_AllElemLengthSlices_Empty(t *testing.T) {
	result := stringslice.AllElemLengthSlices()
	actual := args.Map{"total": result}
	expected := args.Map{"total": 0}
	expected.ShouldBeEqual(t, 0, "AllElemLengthSlices empty", actual)
}

// ── PrependLineNew ──

func Test_Cov6_PrependLineNew(t *testing.T) {
	result := stringslice.PrependLineNew("first", []string{"second", "third"})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 3, "first": "first"}
	expected.ShouldBeEqual(t, 0, "PrependLineNew", actual)
}

// ── AppendLineNew ──

func Test_Cov6_AppendLineNew(t *testing.T) {
	result := stringslice.AppendLineNew([]string{"a", "b"}, "c")
	actual := args.Map{"len": len(result), "last": result[2]}
	expected := args.Map{"len": 3, "last": "c"}
	expected.ShouldBeEqual(t, 0, "AppendLineNew", actual)
}
