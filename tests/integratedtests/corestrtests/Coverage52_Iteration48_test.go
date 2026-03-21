package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ═══════════════════════════════════════════════════════════════
// CharCollectionMap
// ═══════════════════════════════════════════════════════════════

func Test_Cov52_CharCollectionMap_Add(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("hello")
	tc := coretestcases.CaseV1{Name: "CCM Add", Expected: true, Actual: ccm.Has("hello"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AddLock(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.AddLock("world")
	tc := coretestcases.CaseV1{Name: "CCM AddLock", Expected: true, Actual: ccm.Has("world"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AddStrings(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.AddStrings("abc", "axy", "bcd")
	tc := coretestcases.CaseV1{Name: "CCM AddStrings", Expected: true, Actual: ccm.Has("abc"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AddStrings_Empty(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.AddStrings()
	tc := coretestcases.CaseV1{Name: "CCM AddStrings empty", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_IsEmpty(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	tc := coretestcases.CaseV1{Name: "CCM IsEmpty", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_HasItems(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("x")
	tc := coretestcases.CaseV1{Name: "CCM HasItems", Expected: true, Actual: ccm.HasItems(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_IsEmptyLock(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	tc := coretestcases.CaseV1{Name: "CCM IsEmptyLock", Expected: true, Actual: ccm.IsEmptyLock(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_Length(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	ccm.Add("bcd")
	tc := coretestcases.CaseV1{Name: "CCM Length", Expected: 2, Actual: ccm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_LengthLock(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	tc := coretestcases.CaseV1{Name: "CCM LengthLock", Expected: 1, Actual: ccm.LengthLock(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_Has_NotFound(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	tc := coretestcases.CaseV1{Name: "CCM Has not found", Expected: false, Actual: ccm.Has("z"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_HasWithCollection(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	found, col := ccm.HasWithCollection("abc")
	tc := coretestcases.CaseV1{Name: "CCM HasWithCollection", Expected: true, Actual: found, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "CCM HasWithCollection col", Expected: true, Actual: col != nil, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_HasWithCollection_Empty(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	found, _ := ccm.HasWithCollection("z")
	tc := coretestcases.CaseV1{Name: "CCM HasWithCollection empty", Expected: false, Actual: found, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_HasWithCollectionLock(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	found, _ := ccm.HasWithCollectionLock("abc")
	tc := coretestcases.CaseV1{Name: "CCM HasWithCollectionLock", Expected: true, Actual: found, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_LengthOf(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	ccm.Add("axy")
	tc := coretestcases.CaseV1{Name: "CCM LengthOf", Expected: 2, Actual: ccm.LengthOf('a'), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_LengthOf_Missing(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	tc := coretestcases.CaseV1{Name: "CCM LengthOf missing", Expected: 0, Actual: ccm.LengthOf('z'), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_LengthOfLock(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	tc := coretestcases.CaseV1{Name: "CCM LengthOfLock", Expected: 1, Actual: ccm.LengthOfLock('a'), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_LengthOfCollectionFromFirstChar(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	tc := coretestcases.CaseV1{Name: "CCM LengthOfCollFromFirstChar", Expected: 1, Actual: ccm.LengthOfCollectionFromFirstChar("axy"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AllLengthsSum(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	ccm.Add("bcd")
	tc := coretestcases.CaseV1{Name: "CCM AllLengthsSum", Expected: 2, Actual: ccm.AllLengthsSum(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AllLengthsSumLock(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	tc := coretestcases.CaseV1{Name: "CCM AllLengthsSumLock", Expected: 1, Actual: ccm.AllLengthsSumLock(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_GetChar(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	tc := coretestcases.CaseV1{Name: "CCM GetChar", Expected: byte('h'), Actual: ccm.GetChar("hello"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_GetChar_Empty(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	tc := coretestcases.CaseV1{Name: "CCM GetChar empty", Expected: byte(0), Actual: ccm.GetChar(""), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_GetMap(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	tc := coretestcases.CaseV1{Name: "CCM GetMap", Expected: true, Actual: ccm.GetMap() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_GetCopyMapLock(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	tc := coretestcases.CaseV1{Name: "CCM GetCopyMapLock", Expected: true, Actual: ccm.GetCopyMapLock() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_GetCollection(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	col := ccm.GetCollection("axy", false)
	tc := coretestcases.CaseV1{Name: "CCM GetCollection", Expected: true, Actual: col != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_GetCollection_AddNew(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	col := ccm.GetCollection("xyz", true)
	tc := coretestcases.CaseV1{Name: "CCM GetCollection addNew", Expected: true, Actual: col != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_GetCollection_Nil(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	col := ccm.GetCollection("xyz", false)
	tc := coretestcases.CaseV1{Name: "CCM GetCollection nil", Expected: true, Actual: col == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_GetCollectionLock(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	col := ccm.GetCollectionLock("axy", false)
	tc := coretestcases.CaseV1{Name: "CCM GetCollectionLock", Expected: true, Actual: col != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_GetCollectionByChar(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	col := ccm.GetCollectionByChar('a')
	tc := coretestcases.CaseV1{Name: "CCM GetCollectionByChar", Expected: true, Actual: col != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_IsEquals_Same(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	tc := coretestcases.CaseV1{Name: "CCM IsEquals same", Expected: true, Actual: ccm.IsEquals(ccm), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_IsEquals_Nil(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	tc := coretestcases.CaseV1{Name: "CCM IsEquals nil", Expected: false, Actual: ccm.IsEquals(nil), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_IsEquals_Equal(t *testing.T) {
	ccm1 := corestr.New.CharCollectionMap.Cap(5)
	ccm1.Add("abc")
	ccm2 := corestr.New.CharCollectionMap.Cap(5)
	ccm2.Add("abc")
	tc := coretestcases.CaseV1{Name: "CCM IsEquals equal", Expected: true, Actual: ccm1.IsEquals(ccm2), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_IsEqualsLock(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	tc := coretestcases.CaseV1{Name: "CCM IsEqualsLock", Expected: true, Actual: ccm.IsEqualsLock(ccm), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_IsEqualsCaseSensitive(t *testing.T) {
	ccm1 := corestr.New.CharCollectionMap.Cap(5)
	ccm1.Add("abc")
	ccm2 := corestr.New.CharCollectionMap.Cap(5)
	ccm2.Add("abc")
	tc := coretestcases.CaseV1{Name: "CCM IsEqualsCaseSensitive", Expected: true, Actual: ccm1.IsEqualsCaseSensitive(true, ccm2), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_List(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	ccm.Add("bcd")
	tc := coretestcases.CaseV1{Name: "CCM List", Expected: 2, Actual: len(ccm.List()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_ListLock(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	tc := coretestcases.CaseV1{Name: "CCM ListLock", Expected: 1, Actual: len(ccm.ListLock()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_SortedListAsc(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("bcd")
	ccm.Add("abc")
	list := ccm.SortedListAsc()
	tc := coretestcases.CaseV1{Name: "CCM SortedListAsc", Expected: "abc", Actual: list[0], Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_SortedListAsc_Empty(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	tc := coretestcases.CaseV1{Name: "CCM SortedListAsc empty", Expected: 0, Actual: len(ccm.SortedListAsc()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_String(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	tc := coretestcases.CaseV1{Name: "CCM String", Expected: true, Actual: len(ccm.String()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_StringLock(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	tc := coretestcases.CaseV1{Name: "CCM StringLock", Expected: true, Actual: len(ccm.StringLock()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_SummaryString(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	tc := coretestcases.CaseV1{Name: "CCM SummaryString", Expected: true, Actual: len(ccm.SummaryString()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_SummaryStringLock(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	tc := coretestcases.CaseV1{Name: "CCM SummaryStringLock", Expected: true, Actual: len(ccm.SummaryStringLock()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AddSameStartingCharItems(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.AddSameStartingCharItems('a', []string{"abc", "axy"}, false)
	tc := coretestcases.CaseV1{Name: "CCM AddSameStartingCharItems", Expected: true, Actual: ccm.Has("abc"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AddSameStartingCharItems_Existing(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	ccm.AddSameStartingCharItems('a', []string{"axy"}, false)
	tc := coretestcases.CaseV1{Name: "CCM AddSameStartingCharItems existing", Expected: 2, Actual: ccm.LengthOf('a'), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AddSameStartingCharItems_Empty(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.AddSameStartingCharItems('a', []string{}, false)
	tc := coretestcases.CaseV1{Name: "CCM AddSameStartingCharItems empty", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AddCollectionItems(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	col := corestr.New.Collection.Strings("abc", "bcd")
	ccm.AddCollectionItems(col)
	tc := coretestcases.CaseV1{Name: "CCM AddCollectionItems", Expected: 2, Actual: ccm.AllLengthsSum(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AddCollectionItems_Nil(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.AddCollectionItems(nil)
	tc := coretestcases.CaseV1{Name: "CCM AddCollectionItems nil", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AddHashmapsValues(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	hm := corestr.New.Hashmap.Cap(2)
	hm.Add("k", "abc")
	ccm.AddHashmapsValues(hm)
	tc := coretestcases.CaseV1{Name: "CCM AddHashmapsValues", Expected: true, Actual: ccm.Has("abc"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AddHashmapsValues_Nil(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.AddHashmapsValues(nil)
	tc := coretestcases.CaseV1{Name: "CCM AddHashmapsValues nil", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_HashsetByChar(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	hs := ccm.HashsetByChar('a')
	tc := coretestcases.CaseV1{Name: "CCM HashsetByChar", Expected: true, Actual: hs != nil && hs.Has("abc"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_HashsetByChar_Missing(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	hs := ccm.HashsetByChar('z')
	tc := coretestcases.CaseV1{Name: "CCM HashsetByChar missing", Expected: true, Actual: hs == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_HashsetByCharLock(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	hs := ccm.HashsetByCharLock('a')
	tc := coretestcases.CaseV1{Name: "CCM HashsetByCharLock", Expected: true, Actual: hs != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_HashsetByStringFirstChar(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	hs := ccm.HashsetByStringFirstChar("axy")
	tc := coretestcases.CaseV1{Name: "CCM HashsetByStringFirstChar", Expected: true, Actual: hs != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_HashsetsCollection(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	ccm.Add("bcd")
	hsc := ccm.HashsetsCollection()
	tc := coretestcases.CaseV1{Name: "CCM HashsetsCollection", Expected: true, Actual: hsc != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_HashsetsCollection_Empty(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	hsc := ccm.HashsetsCollection()
	tc := coretestcases.CaseV1{Name: "CCM HashsetsCollection empty", Expected: true, Actual: hsc != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_HashsetsCollectionByChars(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	hsc := ccm.HashsetsCollectionByChars('a')
	tc := coretestcases.CaseV1{Name: "CCM HashsetsCollByChars", Expected: true, Actual: hsc != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_HashsetsCollectionByStringFirstChar(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	hsc := ccm.HashsetsCollectionByStringFirstChar("abc")
	tc := coretestcases.CaseV1{Name: "CCM HashsetsCollByStringFirstChar", Expected: true, Actual: hsc != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_Resize(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(2)
	ccm.Add("abc")
	ccm.Resize(10)
	tc := coretestcases.CaseV1{Name: "CCM Resize", Expected: true, Actual: ccm.Has("abc"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_Resize_Smaller(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(10)
	ccm.Add("abc")
	ccm.Resize(1)
	tc := coretestcases.CaseV1{Name: "CCM Resize smaller", Expected: true, Actual: ccm.Has("abc"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AddLength(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(2)
	ccm.AddLength(5, 3)
	tc := coretestcases.CaseV1{Name: "CCM AddLength", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AddLength_Empty(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(2)
	ccm.AddLength()
	tc := coretestcases.CaseV1{Name: "CCM AddLength empty", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_JsonModel(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	dm := ccm.JsonModel()
	tc := coretestcases.CaseV1{Name: "CCM JsonModel", Expected: true, Actual: dm != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_JsonModelAny(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	tc := coretestcases.CaseV1{Name: "CCM JsonModelAny", Expected: true, Actual: ccm.JsonModelAny() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_MarshalJSON(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	data, err := ccm.MarshalJSON()
	tc := coretestcases.CaseV1{Name: "CCM MarshalJSON", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_UnmarshalJSON(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	data, _ := ccm.MarshalJSON()
	ccm2 := corestr.New.CharCollectionMap.Cap(5)
	err := ccm2.UnmarshalJSON(data)
	tc := coretestcases.CaseV1{Name: "CCM UnmarshalJSON", Expected: true, Actual: err == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_Json(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	j := ccm.Json()
	tc := coretestcases.CaseV1{Name: "CCM Json", Expected: true, Actual: j.HasSafeItems(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_JsonPtr(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	j := ccm.JsonPtr()
	tc := coretestcases.CaseV1{Name: "CCM JsonPtr", Expected: true, Actual: j != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AsJsonMarshaller(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	tc := coretestcases.CaseV1{Name: "CCM AsJsonMarshaller", Expected: true, Actual: ccm.AsJsonMarshaller() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AsJsoner(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	tc := coretestcases.CaseV1{Name: "CCM AsJsoner", Expected: true, Actual: ccm.AsJsoner() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AsJsonContractsBinder(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	tc := coretestcases.CaseV1{Name: "CCM AsJsonContractsBinder", Expected: true, Actual: ccm.AsJsonContractsBinder() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AsJsonParseSelfInjector(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	tc := coretestcases.CaseV1{Name: "CCM AsJsonParseSelfInjector", Expected: true, Actual: ccm.AsJsonParseSelfInjector() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AddSameCharsCollection(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	col := corestr.New.Collection.Strings("abc", "axy")
	result := ccm.AddSameCharsCollection("abc", col)
	tc := coretestcases.CaseV1{Name: "CCM AddSameCharsCollection", Expected: true, Actual: result != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AddSameCharsCollection_NilCol(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	result := ccm.AddSameCharsCollection("abc", nil)
	tc := coretestcases.CaseV1{Name: "CCM AddSameCharsCollection nil col", Expected: true, Actual: result != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_AddSameCharsCollectionLock(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	col := corestr.New.Collection.Strings("abc")
	result := ccm.AddSameCharsCollectionLock("abc", col)
	tc := coretestcases.CaseV1{Name: "CCM AddSameCharsCollectionLock", Expected: true, Actual: result != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_Clear(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	ccm.Clear()
	tc := coretestcases.CaseV1{Name: "CCM Clear", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_Clear_Empty(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Clear()
	tc := coretestcases.CaseV1{Name: "CCM Clear empty", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_Dispose(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	ccm.Dispose()
	tc := coretestcases.CaseV1{Name: "CCM Dispose", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_Print(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	ccm.Print(false)
	tc := coretestcases.CaseV1{Name: "CCM Print skip", Expected: true, Actual: true, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_PrintLock(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	ccm.Add("abc")
	ccm.PrintLock(false)
	tc := coretestcases.CaseV1{Name: "CCM PrintLock skip", Expected: true, Actual: true, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_GetCharsGroups(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	result := ccm.GetCharsGroups([]string{"abc", "axy", "bcd"})
	tc := coretestcases.CaseV1{Name: "CCM GetCharsGroups", Expected: true, Actual: result != nil && result.Has("abc"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov52_CharCollectionMap_GetCharsGroups_Empty(t *testing.T) {
	ccm := corestr.New.CharCollectionMap.Cap(5)
	result := ccm.GetCharsGroups([]string{})
	tc := coretestcases.CaseV1{Name: "CCM GetCharsGroups empty", Expected: true, Actual: result == ccm, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}
