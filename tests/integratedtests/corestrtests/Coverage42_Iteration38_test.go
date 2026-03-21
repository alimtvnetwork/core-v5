package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ─── Collection: AddHashmapsKeysValues ──────

func Test_Cov42_Collection_AddHashmapsKeysValues_Valid(t *testing.T) {
	col := corestr.New.Collection.Empty()
	hm := corestr.New.Hashmap.KeyValue("k1", "v1")
	col.AddHashmapsKeysValues(hm)
	tc := coretestcases.CaseV1{
		Name:     "AddHashmapsKeysValues adds both",
		Expected: 2,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_AddHashmapsKeysValues_Nil(t *testing.T) {
	col := corestr.New.Collection.Empty()
	col.AddHashmapsKeysValues(nil)
	tc := coretestcases.CaseV1{
		Name:     "AddHashmapsKeysValues nil",
		Expected: 0,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: AddHashmapsKeysValuesUsingFilter ──────

func Test_Cov42_Collection_AddHashmapsKeysValuesUsingFilter_Accept(t *testing.T) {
	col := corestr.New.Collection.Empty()
	hm := corestr.New.Hashmap.KeyValue("k1", "v1")
	filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
		return pair.Value, true, false
	}
	col.AddHashmapsKeysValuesUsingFilter(filter, hm)
	tc := coretestcases.CaseV1{
		Name:     "AddHashmapsKeysValuesUsingFilter accept",
		Expected: true,
		Actual:   col.Has("v1"),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_AddHashmapsKeysValuesUsingFilter_Break(t *testing.T) {
	col := corestr.New.Collection.Empty()
	hm := corestr.New.Hashmap.KeyValue("k1", "v1")
	filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
		return pair.Value, false, true
	}
	col.AddHashmapsKeysValuesUsingFilter(filter, hm)
	tc := coretestcases.CaseV1{
		Name:     "AddHashmapsKeysValuesUsingFilter break",
		Expected: 0,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_AddHashmapsKeysValuesUsingFilter_Nil(t *testing.T) {
	col := corestr.New.Collection.Empty()
	col.AddHashmapsKeysValuesUsingFilter(nil, nil)
	tc := coretestcases.CaseV1{
		Name:     "AddHashmapsKeysValuesUsingFilter nil",
		Expected: 0,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: AddPointerCollectionsLock ──────

func Test_Cov42_Collection_AddPointerCollectionsLock(t *testing.T) {
	col := corestr.New.Collection.Empty()
	other := corestr.New.Collection.Strings([]string{"a", "b"})
	col.AddPointerCollectionsLock(other)
	tc := coretestcases.CaseV1{
		Name:     "AddPointerCollectionsLock",
		Expected: 2,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: AppendCollectionPtr ──────

func Test_Cov42_Collection_AppendCollectionPtr(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"x"})
	other := corestr.New.Collection.Strings([]string{"y", "z"})
	col.AppendCollectionPtr(other)
	tc := coretestcases.CaseV1{
		Name:     "AppendCollectionPtr",
		Expected: 3,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: Single ──────

func Test_Cov42_Collection_Single_OneItem(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"only"})
	tc := coretestcases.CaseV1{
		Name:     "Single with one item",
		Expected: "only",
		Actual:   col.Single(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: SortedListDsc ──────

func Test_Cov42_Collection_SortedListDsc(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"apple", "cherry", "banana"})
	sorted := col.SortedListDsc()
	tc := coretestcases.CaseV1{
		Name:     "SortedListDsc first item",
		Expected: "cherry",
		Actual:   sorted[0],
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: HasUsingSensitivity ──────

func Test_Cov42_Collection_HasUsingSensitivity_CaseSensitive(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"Hello"})
	tc := coretestcases.CaseV1{
		Name:     "HasUsingSensitivity case sensitive miss",
		Expected: false,
		Actual:   col.HasUsingSensitivity("hello", true),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_HasUsingSensitivity_CaseInsensitive(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"Hello"})
	tc := coretestcases.CaseV1{
		Name:     "HasUsingSensitivity case insensitive match",
		Expected: true,
		Actual:   col.HasUsingSensitivity("hello", false),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: AddStringsAsync ──────

func Test_Cov42_Collection_AddStringsAsync_Empty(t *testing.T) {
	col := corestr.New.Collection.Empty()
	wg := &sync.WaitGroup{}
	col.AddStringsAsync(wg, []string{})
	tc := coretestcases.CaseV1{
		Name:     "AddStringsAsync empty",
		Expected: 0,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: AddNonEmptyStrings / AddNonEmptyStringsSlice ──────

func Test_Cov42_Collection_AddNonEmptyStrings_Valid(t *testing.T) {
	col := corestr.New.Collection.Empty()
	col.AddNonEmptyStrings("a", "b")
	tc := coretestcases.CaseV1{
		Name:     "AddNonEmptyStrings valid",
		Expected: 2,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_AddNonEmptyStrings_Empty(t *testing.T) {
	col := corestr.New.Collection.Empty()
	col.AddNonEmptyStrings()
	tc := coretestcases.CaseV1{
		Name:     "AddNonEmptyStrings no args",
		Expected: 0,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_AddNonEmptyStringsSlice_Valid(t *testing.T) {
	col := corestr.New.Collection.Empty()
	col.AddNonEmptyStringsSlice([]string{"x", "y"})
	tc := coretestcases.CaseV1{
		Name:     "AddNonEmptyStringsSlice valid",
		Expected: 2,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_AddNonEmptyStringsSlice_Empty(t *testing.T) {
	col := corestr.New.Collection.Empty()
	col.AddNonEmptyStringsSlice([]string{})
	tc := coretestcases.CaseV1{
		Name:     "AddNonEmptyStringsSlice empty",
		Expected: 0,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: AddFuncResult ──────

func Test_Cov42_Collection_AddFuncResult_Valid(t *testing.T) {
	col := corestr.New.Collection.Empty()
	col.AddFuncResult(func() string { return "hello" })
	tc := coretestcases.CaseV1{
		Name:     "AddFuncResult valid",
		Expected: "hello",
		Actual:   col.First(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_AddFuncResult_Nil(t *testing.T) {
	col := corestr.New.Collection.Empty()
	col.AddFuncResult(nil)
	tc := coretestcases.CaseV1{
		Name:     "AddFuncResult nil",
		Expected: 0,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: AddStringsByFuncChecking ──────

func Test_Cov42_Collection_AddStringsByFuncChecking_Filter(t *testing.T) {
	col := corestr.New.Collection.Empty()
	col.AddStringsByFuncChecking(
		[]string{"apple", "ban", "cherry"},
		func(line string) bool { return len(line) > 3 },
	)
	tc := coretestcases.CaseV1{
		Name:     "AddStringsByFuncChecking filters",
		Expected: 2,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: ExpandSlicePlusAdd ──────

func Test_Cov42_Collection_ExpandSlicePlusAdd(t *testing.T) {
	col := corestr.New.Collection.Empty()
	col.ExpandSlicePlusAdd(
		[]string{"a,b", "c,d"},
		func(line string) []string {
			return []string{line}
		},
	)
	tc := coretestcases.CaseV1{
		Name:     "ExpandSlicePlusAdd",
		Expected: 2,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: MergeSlicesOfSlice ──────

func Test_Cov42_Collection_MergeSlicesOfSlice(t *testing.T) {
	col := corestr.New.Collection.Empty()
	col.MergeSlicesOfSlice([]string{"a", "b"}, []string{"c"})
	tc := coretestcases.CaseV1{
		Name:     "MergeSlicesOfSlice",
		Expected: 3,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: CharCollectionMap ──────

func Test_Cov42_Collection_CharCollectionMap(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"apple", "banana", "avocado"})
	ccm := col.CharCollectionMap()
	tc := coretestcases.CaseV1{
		Name:     "CharCollectionMap groups by first char",
		Expected: 2,
		Actual:   ccm.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: CsvLines / CsvLinesOptions / Csv / CsvOptions ──────

func Test_Cov42_Collection_CsvLines(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b"})
	lines := col.CsvLines()
	tc := coretestcases.CaseV1{
		Name:     "CsvLines",
		Expected: 2,
		Actual:   len(lines),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_CsvLinesOptions(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a"})
	lines := col.CsvLinesOptions(true)
	tc := coretestcases.CaseV1{
		Name:     "CsvLinesOptions",
		Expected: 1,
		Actual:   len(lines),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_Csv_Empty(t *testing.T) {
	col := corestr.New.Collection.Empty()
	tc := coretestcases.CaseV1{
		Name:     "Csv empty",
		Expected: "",
		Actual:   col.Csv(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_Csv_NonEmpty(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b"})
	csv := col.Csv()
	tc := coretestcases.CaseV1{
		Name:     "Csv non-empty",
		Expected: true,
		Actual:   len(csv) > 0,
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_CsvOptions_Empty(t *testing.T) {
	col := corestr.New.Collection.Empty()
	tc := coretestcases.CaseV1{
		Name:     "CsvOptions empty",
		Expected: "",
		Actual:   col.CsvOptions(true),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_CsvOptions_NonEmpty(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"x"})
	csv := col.CsvOptions(false)
	tc := coretestcases.CaseV1{
		Name:     "CsvOptions non-empty",
		Expected: true,
		Actual:   len(csv) > 0,
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: IsContainsPtr ──────

func Test_Cov42_Collection_IsContainsPtr_Found(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"hello"})
	s := "hello"
	tc := coretestcases.CaseV1{
		Name:     "IsContainsPtr found",
		Expected: true,
		Actual:   col.IsContainsPtr(&s),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_IsContainsPtr_Nil(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"hello"})
	tc := coretestcases.CaseV1{
		Name:     "IsContainsPtr nil",
		Expected: false,
		Actual:   col.IsContainsPtr(nil),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_IsContainsPtr_Empty(t *testing.T) {
	col := corestr.New.Collection.Empty()
	s := "hello"
	tc := coretestcases.CaseV1{
		Name:     "IsContainsPtr empty collection",
		Expected: false,
		Actual:   col.IsContainsPtr(&s),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: GetHashsetPlusHasAll ──────

func Test_Cov42_Collection_GetHashsetPlusHasAll_True(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
	hs, hasAll := col.GetHashsetPlusHasAll([]string{"a", "b"})
	tc := coretestcases.CaseV1{
		Name:     "GetHashsetPlusHasAll true",
		Expected: true,
		Actual:   hasAll && hs != nil,
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_GetHashsetPlusHasAll_False(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b"})
	_, hasAll := col.GetHashsetPlusHasAll([]string{"a", "c"})
	tc := coretestcases.CaseV1{
		Name:     "GetHashsetPlusHasAll false",
		Expected: false,
		Actual:   hasAll,
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: IsContainsAllSlice ──────

func Test_Cov42_Collection_IsContainsAllSlice_True(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
	tc := coretestcases.CaseV1{
		Name:     "IsContainsAllSlice true",
		Expected: true,
		Actual:   col.IsContainsAllSlice([]string{"a", "b"}),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_IsContainsAllSlice_False(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b"})
	tc := coretestcases.CaseV1{
		Name:     "IsContainsAllSlice false",
		Expected: false,
		Actual:   col.IsContainsAllSlice([]string{"a", "c"}),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_IsContainsAllSlice_Empty(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{
		Name:     "IsContainsAllSlice empty items",
		Expected: true,
		Actual:   col.IsContainsAllSlice([]string{}),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_IsContainsAllSlice_EmptyCollection(t *testing.T) {
	col := corestr.New.Collection.Empty()
	tc := coretestcases.CaseV1{
		Name:     "IsContainsAllSlice empty collection",
		Expected: false,
		Actual:   col.IsContainsAllSlice([]string{"a"}),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: IsContainsAll / IsContainsAllLock ──────

func Test_Cov42_Collection_IsContainsAll_True(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
	tc := coretestcases.CaseV1{
		Name:     "IsContainsAll true",
		Expected: true,
		Actual:   col.IsContainsAll("a", "b"),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_IsContainsAllLock_True(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
	tc := coretestcases.CaseV1{
		Name:     "IsContainsAllLock true",
		Expected: true,
		Actual:   col.IsContainsAllLock("a", "b"),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_IsContainsAllLock_Nil(t *testing.T) {
	col := corestr.New.Collection.Empty()
	tc := coretestcases.CaseV1{
		Name:     "IsContainsAllLock nil",
		Expected: false,
		Actual:   col.IsContainsAllLock("x"),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: New (instance method) ──────

func Test_Cov42_Collection_New_Empty(t *testing.T) {
	col := corestr.New.Collection.Empty()
	newCol := col.New()
	tc := coretestcases.CaseV1{
		Name:     "Collection.New empty",
		Expected: 0,
		Actual:   newCol.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_New_WithArgs(t *testing.T) {
	col := corestr.New.Collection.Empty()
	newCol := col.New("a", "b")
	tc := coretestcases.CaseV1{
		Name:     "Collection.New with args",
		Expected: 2,
		Actual:   newCol.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── newCollectionCreator: CloneStrings / LineUsingSep / LineDefault / StringsPlusCap / CapStrings / LenCap ──────

func Test_Cov42_newCollectionCreator_CloneStrings(t *testing.T) {
	items := []string{"a", "b"}
	col := corestr.New.Collection.CloneStrings(items)
	items[0] = "changed"
	tc := coretestcases.CaseV1{
		Name:     "CloneStrings is independent",
		Expected: "a",
		Actual:   col.First(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_newCollectionCreator_LineUsingSep(t *testing.T) {
	col := corestr.New.Collection.LineUsingSep(",", "a,b,c")
	tc := coretestcases.CaseV1{
		Name:     "LineUsingSep",
		Expected: 3,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_newCollectionCreator_LineDefault(t *testing.T) {
	col := corestr.New.Collection.LineDefault("a\nb")
	tc := coretestcases.CaseV1{
		Name:     "LineDefault",
		Expected: true,
		Actual:   col.Length() >= 1,
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_newCollectionCreator_StringsPlusCap_ZeroCap(t *testing.T) {
	col := corestr.New.Collection.StringsPlusCap(0, []string{"a"})
	tc := coretestcases.CaseV1{
		Name:     "StringsPlusCap zero cap",
		Expected: 1,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_newCollectionCreator_StringsPlusCap_WithCap(t *testing.T) {
	col := corestr.New.Collection.StringsPlusCap(10, []string{"a", "b"})
	tc := coretestcases.CaseV1{
		Name:     "StringsPlusCap with cap",
		Expected: 2,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_newCollectionCreator_CapStrings_ZeroCap(t *testing.T) {
	col := corestr.New.Collection.CapStrings(0, []string{"x"})
	tc := coretestcases.CaseV1{
		Name:     "CapStrings zero cap",
		Expected: 1,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_newCollectionCreator_CapStrings_WithCap(t *testing.T) {
	col := corestr.New.Collection.CapStrings(5, []string{"x"})
	tc := coretestcases.CaseV1{
		Name:     "CapStrings with cap",
		Expected: 1,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_newCollectionCreator_LenCap(t *testing.T) {
	col := corestr.New.Collection.LenCap(3, 10)
	tc := coretestcases.CaseV1{
		Name:     "LenCap creates with length",
		Expected: 3,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_newCollectionCreator_Create(t *testing.T) {
	col := corestr.New.Collection.Create([]string{"a"})
	tc := coretestcases.CaseV1{
		Name:     "Create wraps slice",
		Expected: "a",
		Actual:   col.First(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_newCollectionCreator_StringsOptions_Clone(t *testing.T) {
	items := []string{"x"}
	col := corestr.New.Collection.StringsOptions(true, items)
	items[0] = "changed"
	tc := coretestcases.CaseV1{
		Name:     "StringsOptions clone",
		Expected: "x",
		Actual:   col.First(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_newCollectionCreator_StringsOptions_NoClone_Empty(t *testing.T) {
	col := corestr.New.Collection.StringsOptions(false, []string{})
	tc := coretestcases.CaseV1{
		Name:     "StringsOptions no clone empty",
		Expected: true,
		Actual:   col.IsEmpty(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: AppendAnys with nil items ──────

func Test_Cov42_Collection_AppendAnys_WithNilItem(t *testing.T) {
	col := corestr.New.Collection.Empty()
	col.AppendAnys("hello", nil, "world")
	tc := coretestcases.CaseV1{
		Name:     "AppendAnys skips nil",
		Expected: 2,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: AppendAnysUsingFilter ──────

func Test_Cov42_Collection_AppendAnysUsingFilter_Accept(t *testing.T) {
	col := corestr.New.Collection.Empty()
	filter := func(str string, index int) (string, bool, bool) {
		return str, true, false
	}
	col.AppendAnysUsingFilter(filter, "a", "b")
	tc := coretestcases.CaseV1{
		Name:     "AppendAnysUsingFilter accept",
		Expected: 2,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_AppendAnysUsingFilter_Break(t *testing.T) {
	col := corestr.New.Collection.Empty()
	filter := func(str string, index int) (string, bool, bool) {
		return str, true, true
	}
	col.AppendAnysUsingFilter(filter, "a", "b")
	tc := coretestcases.CaseV1{
		Name:     "AppendAnysUsingFilter break after first",
		Expected: 1,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_AppendAnysUsingFilter_Skip(t *testing.T) {
	col := corestr.New.Collection.Empty()
	filter := func(str string, index int) (string, bool, bool) {
		return str, false, false
	}
	col.AppendAnysUsingFilter(filter, "a")
	tc := coretestcases.CaseV1{
		Name:     "AppendAnysUsingFilter skip",
		Expected: 0,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: AppendAnysUsingFilterLock ──────

func Test_Cov42_Collection_AppendAnysUsingFilterLock_Accept(t *testing.T) {
	col := corestr.New.Collection.Empty()
	filter := func(str string, index int) (string, bool, bool) {
		return str, true, false
	}
	col.AppendAnysUsingFilterLock(filter, "x")
	tc := coretestcases.CaseV1{
		Name:     "AppendAnysUsingFilterLock accept",
		Expected: 1,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_AppendAnysUsingFilterLock_Nil(t *testing.T) {
	col := corestr.New.Collection.Empty()
	col.AppendAnysUsingFilterLock(nil, nil)
	tc := coretestcases.CaseV1{
		Name:     "AppendAnysUsingFilterLock nil args",
		Expected: 0,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: AppendNonEmptyAnys ──────

func Test_Cov42_Collection_AppendNonEmptyAnys_Valid(t *testing.T) {
	col := corestr.New.Collection.Empty()
	col.AppendNonEmptyAnys("hello", nil, "world")
	tc := coretestcases.CaseV1{
		Name:     "AppendNonEmptyAnys skips nil",
		Expected: 2,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_AppendNonEmptyAnys_Nil(t *testing.T) {
	col := corestr.New.Collection.Empty()
	col.AppendNonEmptyAnys(nil)
	tc := coretestcases.CaseV1{
		Name:     "AppendNonEmptyAnys nil args",
		Expected: 0,
		Actual:   col.Length(),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: GetAllExceptCollection / GetAllExcept ──────

func Test_Cov42_Collection_GetAllExceptCollection_WithExclude(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
	exclude := corestr.New.Collection.Strings([]string{"b"})
	result := col.GetAllExceptCollection(exclude)
	tc := coretestcases.CaseV1{
		Name:     "GetAllExceptCollection excludes",
		Expected: 2,
		Actual:   len(result),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_GetAllExceptCollection_NilExclude(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b"})
	result := col.GetAllExceptCollection(nil)
	tc := coretestcases.CaseV1{
		Name:     "GetAllExceptCollection nil returns copy",
		Expected: 2,
		Actual:   len(result),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_GetAllExcept_Valid(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
	result := col.GetAllExcept([]string{"c"})
	tc := coretestcases.CaseV1{
		Name:     "GetAllExcept excludes",
		Expected: 2,
		Actual:   len(result),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_GetAllExcept_Nil(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a"})
	result := col.GetAllExcept(nil)
	tc := coretestcases.CaseV1{
		Name:     "GetAllExcept nil returns copy",
		Expected: 1,
		Actual:   len(result),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: SummaryString / SummaryStringWithHeader ──────

func Test_Cov42_Collection_SummaryString(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a"})
	s := col.SummaryString(1)
	tc := coretestcases.CaseV1{
		Name:     "SummaryString",
		Expected: true,
		Actual:   len(s) > 0,
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_SummaryStringWithHeader_NonEmpty(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"x"})
	s := col.SummaryStringWithHeader("header")
	tc := coretestcases.CaseV1{
		Name:     "SummaryStringWithHeader non-empty",
		Expected: true,
		Actual:   len(s) > 0,
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_SummaryStringWithHeader_Empty(t *testing.T) {
	col := corestr.New.Collection.Empty()
	s := col.SummaryStringWithHeader("header")
	tc := coretestcases.CaseV1{
		Name:     "SummaryStringWithHeader empty",
		Expected: true,
		Actual:   len(s) > 0,
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: Joins with extra items ──────

func Test_Cov42_Collection_Joins_WithExtra(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b"})
	result := col.Joins(",", "c")
	tc := coretestcases.CaseV1{
		Name:     "Joins with extra items",
		Expected: true,
		Actual:   len(result) > 0,
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_Joins_NoExtra(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b"})
	result := col.Joins(",")
	tc := coretestcases.CaseV1{
		Name:     "Joins no extra",
		Expected: "a,b",
		Actual:   result,
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: Serialize / Deserialize ──────

func Test_Cov42_Collection_Serialize(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a"})
	bytes, err := col.Serialize()
	tc := coretestcases.CaseV1{
		Name:     "Serialize success",
		Expected: true,
		Actual:   err == nil && len(bytes) > 0,
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_Deserialize(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"hello"})
	var target []string
	err := col.Deserialize(&target)
	tc := coretestcases.CaseV1{
		Name:     "Deserialize success",
		Expected: true,
		Actual:   err == nil && len(target) == 1,
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: NonEmptyList / NonEmptyListPtr ──────

func Test_Cov42_Collection_NonEmptyList(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "", "b"})
	list := col.NonEmptyList()
	tc := coretestcases.CaseV1{
		Name:     "NonEmptyList filters empty",
		Expected: 2,
		Actual:   len(list),
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_NonEmptyListPtr(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "", "b"})
	listPtr := col.NonEmptyListPtr()
	tc := coretestcases.CaseV1{
		Name:     "NonEmptyListPtr non-nil",
		Expected: true,
		Actual:   listPtr != nil,
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

// ─── Collection: StringLock ──────

func Test_Cov42_Collection_StringLock_NonEmpty(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a"})
	s := col.StringLock()
	tc := coretestcases.CaseV1{
		Name:     "StringLock non-empty",
		Expected: true,
		Actual:   len(s) > 0,
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}

func Test_Cov42_Collection_StringLock_Empty(t *testing.T) {
	col := corestr.New.Collection.Empty()
	s := col.StringLock()
	tc := coretestcases.CaseV1{
		Name:     "StringLock empty",
		Expected: true,
		Actual:   len(s) > 0,
		Args:     args.Map{},
	}
	tc.ShouldBeEqual(t)
}
