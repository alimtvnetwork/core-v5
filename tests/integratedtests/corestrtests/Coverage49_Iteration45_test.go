package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ═══════════════════════════════════════════════════════════════
// Collection — remaining: string, join, csv, json, resize, dispose
// ═══════════════════════════════════════════════════════════════

func Test_Cov49_Collection_String(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	tc := coretestcases.CaseV1{Name: "String", Expected: true, Actual: len(c.String()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_String_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	tc := coretestcases.CaseV1{Name: "String empty", Expected: true, Actual: len(c.String()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_StringLock(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "StringLock", Expected: true, Actual: len(c.StringLock()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_StringLock_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	tc := coretestcases.CaseV1{Name: "StringLock empty", Expected: true, Actual: len(c.StringLock()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_SummaryString(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "SummaryString", Expected: true, Actual: len(c.SummaryString(1)) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_SummaryStringWithHeader_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	tc := coretestcases.CaseV1{Name: "SummaryStringWithHeader empty", Expected: true, Actual: len(c.SummaryStringWithHeader("H")) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_SummaryStringWithHeader_HasItems(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "SummaryStringWithHeader has", Expected: true, Actual: len(c.SummaryStringWithHeader("H")) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_Join(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	tc := coretestcases.CaseV1{Name: "Join", Expected: "a,b", Actual: c.Join(","), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_Join_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	tc := coretestcases.CaseV1{Name: "Join empty", Expected: "", Actual: c.Join(","), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_JoinLine(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	tc := coretestcases.CaseV1{Name: "JoinLine", Expected: true, Actual: len(c.JoinLine()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_JoinLine_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	tc := coretestcases.CaseV1{Name: "JoinLine empty", Expected: "", Actual: c.JoinLine(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_Joins_NoExtra(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	tc := coretestcases.CaseV1{Name: "Joins no extra", Expected: "a,b", Actual: c.Joins(","), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_Joins_WithExtra(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	result := c.Joins(",", "b", "c")
	tc := coretestcases.CaseV1{Name: "Joins with extra", Expected: true, Actual: len(result) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_NonEmptyJoins(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "", "b"})
	result := c.NonEmptyJoins(",")
	tc := coretestcases.CaseV1{Name: "NonEmptyJoins", Expected: "a,b", Actual: result, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_NonWhitespaceJoins(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "  ", "b"})
	result := c.NonWhitespaceJoins(",")
	tc := coretestcases.CaseV1{Name: "NonWhitespaceJoins", Expected: "a,b", Actual: result, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_Csv(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	result := c.Csv()
	tc := coretestcases.CaseV1{Name: "Csv", Expected: true, Actual: len(result) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_Csv_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	tc := coretestcases.CaseV1{Name: "Csv empty", Expected: "", Actual: c.Csv(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_CsvOptions(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	result := c.CsvOptions(true)
	tc := coretestcases.CaseV1{Name: "CsvOptions", Expected: true, Actual: len(result) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_CsvOptions_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	tc := coretestcases.CaseV1{Name: "CsvOptions empty", Expected: "", Actual: c.CsvOptions(false), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_CsvLines(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	result := c.CsvLines()
	tc := coretestcases.CaseV1{Name: "CsvLines", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_CsvLinesOptions(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	result := c.CsvLinesOptions(true)
	tc := coretestcases.CaseV1{Name: "CsvLinesOptions", Expected: 1, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_GetAllExcept(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
	result := c.GetAllExcept([]string{"b"})
	tc := coretestcases.CaseV1{Name: "GetAllExcept", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_GetAllExcept_Nil(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	result := c.GetAllExcept(nil)
	tc := coretestcases.CaseV1{Name: "GetAllExcept nil", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_CharCollectionMap(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"hello", "hi", "abc"})
	ccm := c.CharCollectionMap()
	tc := coretestcases.CaseV1{Name: "CharCollectionMap", Expected: true, Actual: ccm != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_AddCapacity(t *testing.T) {
	c := corestr.New.Collection.Cap(2)
	c.AddCapacity(10)
	tc := coretestcases.CaseV1{Name: "AddCapacity", Expected: true, Actual: c.Capacity() >= 10, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_AddCapacity_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AddCapacity()
	tc := coretestcases.CaseV1{Name: "AddCapacity empty", Expected: true, Actual: c != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_Resize_Bigger(t *testing.T) {
	c := corestr.New.Collection.Cap(2)
	c.Resize(20)
	tc := coretestcases.CaseV1{Name: "Resize bigger", Expected: true, Actual: c.Capacity() >= 20, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_Resize_Smaller(t *testing.T) {
	c := corestr.New.Collection.Cap(20)
	c.Resize(5)
	tc := coretestcases.CaseV1{Name: "Resize smaller noop", Expected: true, Actual: c.Capacity() >= 20, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_Clear(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	c.Clear()
	tc := coretestcases.CaseV1{Name: "Clear", Expected: 0, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_Clear_Nil(t *testing.T) {
	var c *corestr.Collection
	result := c.Clear()
	tc := coretestcases.CaseV1{Name: "Clear nil", Expected: true, Actual: result == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_Dispose(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	c.Dispose()
	tc := coretestcases.CaseV1{Name: "Dispose", Expected: true, Actual: c.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_Dispose_Nil(t *testing.T) {
	var c *corestr.Collection
	c.Dispose() // should not panic
	tc := coretestcases.CaseV1{Name: "Dispose nil", Expected: true, Actual: true, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_JsonModel(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "JsonModel", Expected: 1, Actual: len(c.JsonModel()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_JsonModelAny(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "JsonModelAny", Expected: true, Actual: c.JsonModelAny() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_MarshalJSON(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	data, err := c.MarshalJSON()
	tc := coretestcases.CaseV1{Name: "MarshalJSON", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_UnmarshalJSON(t *testing.T) {
	c := &corestr.Collection{}
	err := c.UnmarshalJSON([]byte(`["a","b"]`))
	tc := coretestcases.CaseV1{Name: "UnmarshalJSON", Expected: true, Actual: err == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "UnmarshalJSON length", Expected: 2, Actual: c.Length(), Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov49_Collection_Serialize(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	data, err := c.Serialize()
	tc := coretestcases.CaseV1{Name: "Serialize", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_AsJsonContractsBinder(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "AsJsonContractsBinder", Expected: true, Actual: c.AsJsonContractsBinder() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_AsJsonMarshaller(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "AsJsonMarshaller", Expected: true, Actual: c.AsJsonMarshaller() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_Collection_ExpandSlicePlusAdd(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.ExpandSlicePlusAdd([]string{"a,b"}, func(s string) []string {
		return []string{s + "_expanded"}
	})
	tc := coretestcases.CaseV1{Name: "ExpandSlicePlusAdd", Expected: 1, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// CharHashsetMap — remaining: json, serialize, clear, remove
// ═══════════════════════════════════════════════════════════════

func Test_Cov49_CharHashsetMap_GetHashsetByChar(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	hs := chm.GetHashsetByChar('h')
	tc := coretestcases.CaseV1{Name: "CHM GetHashsetByChar", Expected: true, Actual: hs != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_HashsetByCharLock_Found(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	hs := chm.HashsetByCharLock('h')
	tc := coretestcases.CaseV1{Name: "CHM HashsetByCharLock found", Expected: true, Actual: hs.Has("hello"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_HashsetByCharLock_NotFound(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	hs := chm.HashsetByCharLock('z')
	tc := coretestcases.CaseV1{Name: "CHM HashsetByCharLock not found", Expected: true, Actual: hs.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_HashsetByStringFirstChar(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	hs := chm.HashsetByStringFirstChar("hello")
	tc := coretestcases.CaseV1{Name: "CHM HashsetByStringFirstChar", Expected: true, Actual: hs != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_HashsetByStringFirstCharLock(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	hs := chm.HashsetByStringFirstCharLock("hello")
	tc := coretestcases.CaseV1{Name: "CHM HashsetByStringFirstCharLock", Expected: true, Actual: hs.Has("hello"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_JsonModel(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	tc := coretestcases.CaseV1{Name: "CHM JsonModel", Expected: true, Actual: chm.JsonModel() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_JsonModelAny(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	tc := coretestcases.CaseV1{Name: "CHM JsonModelAny", Expected: true, Actual: chm.JsonModelAny() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_AsJsonContractsBinder(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	tc := coretestcases.CaseV1{Name: "CHM AsJsonContractsBinder", Expected: true, Actual: chm.AsJsonContractsBinder() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_AsJsoner(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	tc := coretestcases.CaseV1{Name: "CHM AsJsoner", Expected: true, Actual: chm.AsJsoner() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_AsJsonMarshaller(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	tc := coretestcases.CaseV1{Name: "CHM AsJsonMarshaller", Expected: true, Actual: chm.AsJsonMarshaller() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_AsJsonParseSelfInjector(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	tc := coretestcases.CaseV1{Name: "CHM AsJsonParseSelfInjector", Expected: true, Actual: chm.AsJsonParseSelfInjector() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_RemoveAll(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	chm.RemoveAll()
	tc := coretestcases.CaseV1{Name: "CHM RemoveAll", Expected: true, Actual: chm.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_RemoveAll_Empty(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.RemoveAll()
	tc := coretestcases.CaseV1{Name: "CHM RemoveAll empty", Expected: true, Actual: chm.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_Clear(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	chm.Clear()
	tc := coretestcases.CaseV1{Name: "CHM Clear", Expected: true, Actual: chm.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_StringLock(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	tc := coretestcases.CaseV1{Name: "CHM StringLock", Expected: true, Actual: len(chm.StringLock()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_SummaryStringLock(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	tc := coretestcases.CaseV1{Name: "CHM SummaryStringLock", Expected: true, Actual: len(chm.SummaryStringLock()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_Print_True(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	chm.Print(true) // should not panic
	tc := coretestcases.CaseV1{Name: "CHM Print true", Expected: true, Actual: true, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_PrintLock_False(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.PrintLock(false)
	tc := coretestcases.CaseV1{Name: "CHM PrintLock false", Expected: true, Actual: true, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_PrintLock_True(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	chm.PrintLock(true)
	tc := coretestcases.CaseV1{Name: "CHM PrintLock true", Expected: true, Actual: true, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_AddHashsetLock_New(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	hs := corestr.New.Hashset.StringsSpreadItems("abc")
	result := chm.AddHashsetLock("a", hs)
	tc := coretestcases.CaseV1{Name: "CHM AddHashsetLock new", Expected: true, Actual: result != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_AddHashsetLock_Existing(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("abc")
	hs := corestr.New.Hashset.StringsSpreadItems("axy")
	result := chm.AddHashsetLock("a", hs)
	tc := coretestcases.CaseV1{Name: "CHM AddHashsetLock existing", Expected: true, Actual: result != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_AddHashsetLock_NilHashset(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	result := chm.AddHashsetLock("a", nil)
	tc := coretestcases.CaseV1{Name: "CHM AddHashsetLock nil", Expected: true, Actual: result != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_AddSameCharsCollectionLock_New(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	col := corestr.New.Collection.Strings([]string{"abc"})
	result := chm.AddSameCharsCollectionLock("a", col)
	tc := coretestcases.CaseV1{Name: "CHM AddSameCharsCollectionLock new", Expected: true, Actual: result != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_AddSameCharsCollectionLock_Existing(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("abc")
	col := corestr.New.Collection.Strings([]string{"axy"})
	result := chm.AddSameCharsCollectionLock("a", col)
	tc := coretestcases.CaseV1{Name: "CHM AddSameCharsCollectionLock existing", Expected: true, Actual: result != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_AddSameCharsCollectionLock_Nil(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	result := chm.AddSameCharsCollectionLock("a", nil)
	tc := coretestcases.CaseV1{Name: "CHM AddSameCharsCollectionLock nil", Expected: true, Actual: result != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov49_CharHashsetMap_AddSameCharsCollectionLock_ExistingNil(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("abc")
	result := chm.AddSameCharsCollectionLock("a", nil)
	tc := coretestcases.CaseV1{Name: "CHM AddSameCharsCollectionLock existing nil", Expected: true, Actual: result != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}
