package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/smartystreets/goconvey/convey"
)

// =============================================================================
// HashsetsCollection
// =============================================================================

func Test_Cov57_HashsetsCollection_IsEmpty(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("HashsetsCollection IsEmpty", args.Map{"IsEmpty": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsEmpty": hc.IsEmpty()})
}

func Test_Cov57_HashsetsCollection_HasItems(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hs := corestr.New.Hashset.Strings([]string{"a"})
	hc.Add(hs)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("HashsetsCollection HasItems", args.Map{"HasItems": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"HasItems": hc.HasItems()})
}

func Test_Cov57_HashsetsCollection_Length(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("HashsetsCollection Length", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": hc.Length()})
}

func Test_Cov57_HashsetsCollection_Add(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hs := corestr.New.Hashset.Strings([]string{"a"})
	hc.Add(hs)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("HashsetsCollection Add", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": hc.Length()})
}

func Test_Cov57_HashsetsCollection_AddNonNil(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.AddNonNil(nil)
	hs := corestr.New.Hashset.Strings([]string{"a"})
	hc.AddNonNil(hs)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddNonNil", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": hc.Length()})
}

func Test_Cov57_HashsetsCollection_AddNonEmpty(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.AddNonEmpty(corestr.New.Hashset.Empty())
	hc.AddNonEmpty(corestr.New.Hashset.Strings([]string{"a"}))
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddNonEmpty", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": hc.Length()})
}

func Test_Cov57_HashsetsCollection_Adds(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hs1 := corestr.New.Hashset.Strings([]string{"a"})
	hs2 := corestr.New.Hashset.Strings([]string{"b"})
	hc.Adds(hs1, hs2)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Adds", args.Map{"Length": 2}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": hc.Length()})
}

func Test_Cov57_HashsetsCollection_Adds_Nil(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.Adds(nil)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Adds nil", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": hc.Length()})
}

func Test_Cov57_HashsetsCollection_StringsList(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
	result := hc.StringsList()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("StringsList", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(result)})
}

func Test_Cov57_HashsetsCollection_StringsList_Empty(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	result := hc.StringsList()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("StringsList empty", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(result)})
}

func Test_Cov57_HashsetsCollection_HasAll_True(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.Add(corestr.New.Hashset.Strings([]string{"a", "b"}))
	result := hc.HasAll("a", "b")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("HasAll true", args.Map{"HasAll": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"HasAll": result})
}

func Test_Cov57_HashsetsCollection_HasAll_Empty(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	result := hc.HasAll("a")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("HasAll empty", args.Map{"HasAll": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"HasAll": result})
}

func Test_Cov57_HashsetsCollection_ListDirectPtr(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
	result := hc.ListDirectPtr()
	convey.Convey("ListDirectPtr", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
		convey.So(len(*result), convey.ShouldEqual, 1)
	})
}

func Test_Cov57_HashsetsCollection_AddHashsetsCollection(t *testing.T) {
	hc1 := corestr.New.HashsetsCollection.Empty()
	hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
	hc2 := corestr.New.HashsetsCollection.Empty()
	hc2.Add(corestr.New.Hashset.Strings([]string{"b"}))
	hc1.AddHashsetsCollection(hc2)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddHashsetsCollection", args.Map{"Length": 2}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": hc1.Length()})
}

func Test_Cov57_HashsetsCollection_AddHashsetsCollection_Nil(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.AddHashsetsCollection(nil)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddHashsetsCollection nil", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": hc.Length()})
}

func Test_Cov57_HashsetsCollection_ConcatNew(t *testing.T) {
	hc1 := corestr.New.HashsetsCollection.Empty()
	hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
	hc2 := corestr.New.HashsetsCollection.Empty()
	hc2.Add(corestr.New.Hashset.Strings([]string{"b"}))
	result := hc1.ConcatNew(hc2)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("ConcatNew", args.Map{"Length": 2}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": result.Length()})
}

func Test_Cov57_HashsetsCollection_ConcatNew_NoArgs(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
	result := hc.ConcatNew()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("ConcatNew no args", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": result.Length()})
}

func Test_Cov57_HashsetsCollection_LastIndex(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LastIndex", args.Map{"LastIndex": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"LastIndex": hc.LastIndex()})
}

func Test_Cov57_HashsetsCollection_IsEqual(t *testing.T) {
	hc1 := corestr.New.HashsetsCollection.Empty()
	hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
	hc2 := corestr.New.HashsetsCollection.Empty()
	hc2.Add(corestr.New.Hashset.Strings([]string{"a"}))
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("IsEqual", args.Map{"IsEqual": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsEqual": hc1.IsEqualPtr(hc2)})
}

func Test_Cov57_HashsetsCollection_IsEqualPtr_DiffLength(t *testing.T) {
	hc1 := corestr.New.HashsetsCollection.Empty()
	hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
	hc2 := corestr.New.HashsetsCollection.Empty()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("IsEqualPtr diff length", args.Map{"IsEqual": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsEqual": hc1.IsEqualPtr(hc2)})
}

func Test_Cov57_HashsetsCollection_IsEqualPtr_BothEmpty(t *testing.T) {
	hc1 := corestr.New.HashsetsCollection.Empty()
	hc2 := corestr.New.HashsetsCollection.Empty()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("IsEqualPtr both empty", args.Map{"IsEqual": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsEqual": hc1.IsEqualPtr(hc2)})
}

func Test_Cov57_HashsetsCollection_IsEqualPtr_SamePtr(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("IsEqualPtr same ptr", args.Map{"IsEqual": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsEqual": hc.IsEqualPtr(hc)})
}

func Test_Cov57_HashsetsCollection_IsEqual_Val(t *testing.T) {
	hc1 := corestr.New.HashsetsCollection.Empty()
	hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
	hc2 := corestr.New.HashsetsCollection.Empty()
	hc2.Add(corestr.New.Hashset.Strings([]string{"a"}))
	convey.Convey("IsEqual val", t, func() {
		convey.So(hc1.IsEqual(*hc2), convey.ShouldBeTrue)
	})
}

func Test_Cov57_HashsetsCollection_Json(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
	jsonResult := hc.JsonPtr()
	target := corestr.New.HashsetsCollection.Empty()
	result := target.ParseInjectUsingJsonMust(jsonResult)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Json roundtrip", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": result.Length()})
}

func Test_Cov57_HashsetsCollection_String(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	convey.Convey("String empty", t, func() {
		convey.So(hc.String(), convey.ShouldContainSubstring, "No Element")
	})
}

func Test_Cov57_HashsetsCollection_String_NonEmpty(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
	convey.Convey("String non-empty", t, func() {
		convey.So(hc.String(), convey.ShouldNotBeEmpty)
	})
}

func Test_Cov57_HashsetsCollection_Join(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
	convey.Convey("Join", t, func() {
		convey.So(hc.Join(","), convey.ShouldEqual, "a")
	})
}

func Test_Cov57_HashsetsCollection_Serialize(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
	bytes, err := hc.Serialize()
	convey.Convey("Serialize", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(bytes, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov57_HashsetsCollection_Deserialize(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
	var target corestr.HashsetsCollection
	err := hc.Deserialize(&target)
	convey.Convey("Deserialize", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_Cov57_HashsetsCollection_Interfaces(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	convey.Convey("Interfaces", t, func() {
		convey.So(hc.AsJsonContractsBinder(), convey.ShouldNotBeNil)
		convey.So(hc.AsJsoner(), convey.ShouldNotBeNil)
		convey.So(hc.AsJsonParseSelfInjector(), convey.ShouldNotBeNil)
		convey.So(hc.AsJsonMarshaller(), convey.ShouldNotBeNil)
	})
}

func Test_Cov57_HashsetsCollection_IndexOf(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
	convey.Convey("IndexOf", t, func() {
		convey.So(hc.IndexOf(0), convey.ShouldNotBeNil)
	})
}

func Test_Cov57_HashsetsCollection_ListPtr(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Empty()
	convey.Convey("ListPtr", t, func() {
		convey.So(hc.ListPtr(), convey.ShouldNotBeNil)
		convey.So(hc.List(), convey.ShouldNotBeNil)
	})
}

// =============================================================================
// KeyValuePair
// =============================================================================

func Test_Cov57_KeyValuePair_Basics(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "name", Value: "test"}
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVP basics", args.Map{
			"KeyName":      "name",
			"ValueString":  "test",
			"IsKey":        true,
			"IsVal":        true,
			"HasKey":       true,
			"HasValue":     true,
			"IsKeyEmpty":   false,
			"IsValueEmpty": false,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"KeyName":      kv.KeyName(),
		"ValueString":  kv.ValueString(),
		"IsKey":        kv.IsKey("name"),
		"IsVal":        kv.IsVal("test"),
		"HasKey":       kv.HasKey(),
		"HasValue":     kv.HasValue(),
		"IsKeyEmpty":   kv.IsKeyEmpty(),
		"IsValueEmpty": kv.IsValueEmpty(),
	})
}

func Test_Cov57_KeyValuePair_ValueConversions(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "42"}
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVP conversions", args.Map{
			"ValueInt":  42,
			"DefInt":    42,
			"DefByte":   byte(42),
			"Float":     42.0,
			"DefFloat":  42.0,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"ValueInt":  kv.ValueInt(0),
		"DefInt":    kv.ValueDefInt(),
		"DefByte":   kv.ValueDefByte(),
		"Float":     kv.ValueFloat64(0),
		"DefFloat":  kv.ValueDefFloat64(),
	})
}

func Test_Cov57_KeyValuePair_ValueBool(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "true"}
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVP ValueBool", args.Map{"Bool": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Bool": kv.ValueBool()})
}

func Test_Cov57_KeyValuePair_ValueBool_Empty(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: ""}
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVP ValueBool empty", args.Map{"Bool": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Bool": kv.ValueBool()})
}

func Test_Cov57_KeyValuePair_Is(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVP Is", args.Map{"Is": true, "IsNot": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Is": kv.Is("k", "v"), "IsNot": kv.Is("k", "x")})
}

func Test_Cov57_KeyValuePair_Trim(t *testing.T) {
	kv := corestr.KeyValuePair{Key: " k ", Value: " v "}
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVP Trim", args.Map{"TrimKey": "k", "TrimVal": "v"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"TrimKey": kv.TrimKey(), "TrimVal": kv.TrimValue()})
}

func Test_Cov57_KeyValuePair_ValueByte(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "100"}
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVP ValueByte", args.Map{"Byte": byte(100)}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Byte": kv.ValueByte(0)})
}

func Test_Cov57_KeyValuePair_FormatString(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	result := kv.FormatString("%s=%s")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVP FormatString", "k=v"),
	}
	tc.ShouldBeEqual(0, t, result)
}

func Test_Cov57_KeyValuePair_String(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	convey.Convey("KVP String", t, func() {
		convey.So(kv.String(), convey.ShouldNotBeEmpty)
	})
}

func Test_Cov57_KeyValuePair_Clear(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	kv.Clear()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVP Clear", args.Map{"IsEmpty": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsEmpty": kv.IsKeyValueEmpty()})
}

func Test_Cov57_KeyValuePair_Dispose(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	kv.Dispose()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVP Dispose", args.Map{"IsEmpty": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsEmpty": kv.IsKeyValueEmpty()})
}

func Test_Cov57_KeyValuePair_ValueValid(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	vv := kv.ValueValid()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVP ValueValid", args.Map{"IsValid": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsValid": vv.IsValid})
}

func Test_Cov57_KeyValuePair_ValueValidOptions(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	vv := kv.ValueValidOptions(false, "err")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVP ValueValidOptions", args.Map{"IsValid": false, "Message": "err"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsValid": vv.IsValid, "Message": vv.Message})
}

func Test_Cov57_KeyValuePair_IsKeyValueAnyEmpty(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "", Value: "v"}
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("IsKeyValueAnyEmpty", args.Map{"Result": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Result": kv.IsKeyValueAnyEmpty()})
}

func Test_Cov57_KeyValuePair_IsVariableNameEqual(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("IsVariableNameEqual", args.Map{"Result": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Result": kv.IsVariableNameEqual("k")})
}

func Test_Cov57_KeyValuePair_IsValueEqual(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("IsValueEqual", args.Map{"Result": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Result": kv.IsValueEqual("v")})
}

func Test_Cov57_KeyValuePair_Json(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	convey.Convey("KVP Json", t, func() {
		convey.So(kv.JsonPtr(), convey.ShouldNotBeNil)
	})
}

func Test_Cov57_KeyValuePair_Serialize(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	bytes, err := kv.Serialize()
	convey.Convey("KVP Serialize", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(bytes, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov57_KeyValuePair_SerializeMust(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	convey.Convey("KVP SerializeMust", t, func() {
		convey.So(kv.SerializeMust(), convey.ShouldNotBeEmpty)
	})
}

func Test_Cov57_KeyValuePair_Compile(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	convey.Convey("KVP Compile", t, func() {
		convey.So(kv.Compile(), convey.ShouldEqual, kv.String())
	})
}

func Test_Cov57_KeyValuePair_VariableName(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VariableName", "k"),
	}
	tc.ShouldBeEqual(0, t, kv.VariableName())
}

// =============================================================================
// KeyValueCollection
// =============================================================================

func Test_Cov57_KeyValueCollection_Basics(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k1", "v1").Add("k2", "v2")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVC basics", args.Map{
			"Length":  2,
			"Count":  2,
			"HasAny": true,
			"HasKey": true,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"Length":  kvc.Length(),
		"Count":  kvc.Count(),
		"HasAny": kvc.HasAnyItem(),
		"HasKey": kvc.HasKey("k1"),
	})
}

func Test_Cov57_KeyValueCollection_AddIf(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddIf(true, "a", "b")
	kvc.AddIf(false, "c", "d")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVC AddIf", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": kvc.Length()})
}

func Test_Cov57_KeyValueCollection_FirstLast(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k1", "v1").Add("k2", "v2")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVC FirstLast", args.Map{
			"FirstKey": "k1",
			"LastKey":  "k2",
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"FirstKey": kvc.First().Key,
		"LastKey":  kvc.Last().Key,
	})
}

func Test_Cov57_KeyValueCollection_FirstOrDefault_Empty(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	convey.Convey("FirstOrDefault empty", t, func() {
		convey.So(kvc.FirstOrDefault(), convey.ShouldBeNil)
		convey.So(kvc.LastOrDefault(), convey.ShouldBeNil)
	})
}

func Test_Cov57_KeyValueCollection_Find(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("a", "1").Add("b", "2")
	result := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
		return kv, kv.Key == "a", false
	})
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVC Find", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(result)})
}

func Test_Cov57_KeyValueCollection_SafeValueAt(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("SafeValueAt", args.Map{"Value": "v", "Empty": ""}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Value": kvc.SafeValueAt(0), "Empty": kvc.SafeValueAt(99)})
}

func Test_Cov57_KeyValueCollection_SafeValuesAtIndexes(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	result := kvc.SafeValuesAtIndexes(0, 99)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("SafeValuesAtIndexes", args.Map{"Length": 2}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(result)})
}

func Test_Cov57_KeyValueCollection_StringsUsingFormat(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	result := kvc.StringsUsingFormat("%s=%s")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("StringsUsingFormat", "k=v"),
	}
	tc.ShouldBeEqual(0, t, result[0])
}

func Test_Cov57_KeyValueCollection_AllKeysValues(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k1", "v1").Add("k2", "v2")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AllKeys/AllValues", args.Map{"KeysLen": 2, "ValsLen": 2}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"KeysLen": len(kvc.AllKeys()), "ValsLen": len(kvc.AllValues())})
}

func Test_Cov57_KeyValueCollection_AllKeysSorted(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("b", "2").Add("a", "1")
	result := kvc.AllKeysSorted()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AllKeysSorted", args.Map{"First": "a"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"First": result[0]})
}

func Test_Cov57_KeyValueCollection_JoinMethods(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	convey.Convey("Join methods", t, func() {
		convey.So(kvc.Join(","), convey.ShouldNotBeEmpty)
		convey.So(kvc.JoinKeys(","), convey.ShouldEqual, "k")
		convey.So(kvc.JoinValues(","), convey.ShouldEqual, "v")
	})
}

func Test_Cov57_KeyValueCollection_AddMap(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddMap(map[string]string{"k": "v"})
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddMap", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": kvc.Length()})
}

func Test_Cov57_KeyValueCollection_AddMap_Nil(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddMap(nil)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddMap nil", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": kvc.Length()})
}

func Test_Cov57_KeyValueCollection_AddHashsetMap(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddHashsetMap(map[string]bool{"a": true})
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddHashsetMap", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": kvc.Length()})
}

func Test_Cov57_KeyValueCollection_AddHashsetMap_Nil(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddHashsetMap(nil)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddHashsetMap nil", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": kvc.Length()})
}

func Test_Cov57_KeyValueCollection_AddHashset(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	hs := corestr.New.Hashset.Strings([]string{"a"})
	kvc.AddHashset(hs)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddHashset", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": kvc.Length()})
}

func Test_Cov57_KeyValueCollection_AddHashset_Nil(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddHashset(nil)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddHashset nil", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": kvc.Length()})
}

func Test_Cov57_KeyValueCollection_AddsHashmap(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	hm := corestr.New.Hashmap.Cap(1)
	hm.AddOrUpdate("k", "v")
	kvc.AddsHashmap(hm)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddsHashmap", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": kvc.Length()})
}

func Test_Cov57_KeyValueCollection_AddsHashmap_Nil(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddsHashmap(nil)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddsHashmap nil", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": kvc.Length()})
}

func Test_Cov57_KeyValueCollection_AddsHashmaps(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	hm := corestr.New.Hashmap.Cap(1)
	hm.AddOrUpdate("k", "v")
	kvc.AddsHashmaps(hm)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddsHashmaps", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": kvc.Length()})
}

func Test_Cov57_KeyValueCollection_AddsHashmaps_Nil(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddsHashmaps()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddsHashmaps nil", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": kvc.Length()})
}

func Test_Cov57_KeyValueCollection_Hashmap(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	hm := kvc.Hashmap()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Hashmap", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": hm.Length()})
}

func Test_Cov57_KeyValueCollection_Map(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	m := kvc.Map()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Map", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(m)})
}

func Test_Cov57_KeyValueCollection_IsContains(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("IsContains", args.Map{"Contains": true, "Missing": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Contains": kvc.IsContains("k"), "Missing": kvc.IsContains("x")})
}

func Test_Cov57_KeyValueCollection_Get(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	val, ok := kvc.Get("k")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Get", args.Map{"Value": "v", "Ok": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Value": val, "Ok": ok})
}

func Test_Cov57_KeyValueCollection_Get_Missing(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	_, ok := kvc.Get("k")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("Get missing", args.Map{"Ok": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Ok": ok})
}

func Test_Cov57_KeyValueCollection_HasIndex(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("HasIndex", args.Map{"Has": true, "Missing": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Has": kvc.HasIndex(0), "Missing": kvc.HasIndex(99)})
}

func Test_Cov57_KeyValueCollection_Adds(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Adds(
		corestr.KeyValuePair{Key: "a", Value: "1"},
		corestr.KeyValuePair{Key: "b", Value: "2"},
	)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVC Adds", args.Map{"Length": 2}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": kvc.Length()})
}

func Test_Cov57_KeyValueCollection_Adds_Empty(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Adds()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("KVC Adds empty", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": kvc.Length()})
}

func Test_Cov57_KeyValueCollection_AddStringBySplit(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddStringBySplit("=", "key=value")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddStringBySplit", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": kvc.Length()})
}

func Test_Cov57_KeyValueCollection_AddStringBySplitTrim(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddStringBySplitTrim("=", " key = value ")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AddStringBySplitTrim", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": kvc.Length()})
}

func Test_Cov57_KeyValueCollection_Serialize(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	bytes, err := kvc.Serialize()
	convey.Convey("KVC Serialize", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(bytes, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov57_KeyValueCollection_SerializeMust(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	convey.Convey("KVC SerializeMust", t, func() {
		convey.So(kvc.SerializeMust(), convey.ShouldNotBeEmpty)
	})
}

func Test_Cov57_KeyValueCollection_String(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	convey.Convey("KVC String", t, func() {
		convey.So(kvc.String(), convey.ShouldNotBeEmpty)
		convey.So(kvc.Compile(), convey.ShouldEqual, kvc.String())
	})
}

// =============================================================================
// LeftRight
// =============================================================================

func Test_Cov57_LeftRight_New(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR New", args.Map{
			"Left": "a", "Right": "b", "IsValid": true,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"Left": lr.Left, "Right": lr.Right, "IsValid": lr.IsValid,
	})
}

func Test_Cov57_LeftRight_Invalid(t *testing.T) {
	lr := corestr.InvalidLeftRight("err")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR Invalid", args.Map{"IsValid": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsValid": lr.IsValid})
}

func Test_Cov57_LeftRight_InvalidNoMessage(t *testing.T) {
	lr := corestr.InvalidLeftRightNoMessage()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR InvalidNoMessage", args.Map{"IsValid": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsValid": lr.IsValid})
}

func Test_Cov57_LeftRight_Bytes(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	convey.Convey("LR Bytes", t, func() {
		convey.So(lr.LeftBytes(), convey.ShouldResemble, []byte("a"))
		convey.So(lr.RightBytes(), convey.ShouldResemble, []byte("b"))
	})
}

func Test_Cov57_LeftRight_Trim(t *testing.T) {
	lr := corestr.NewLeftRight(" a ", " b ")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR Trim", args.Map{"Left": "a", "Right": "b"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lr.LeftTrim(), "Right": lr.RightTrim()})
}

func Test_Cov57_LeftRight_EmptyChecks(t *testing.T) {
	lr := corestr.NewLeftRight("", "b")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR empty checks", args.Map{
			"LeftEmpty":  true,
			"RightEmpty": false,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"LeftEmpty":  lr.IsLeftEmpty(),
		"RightEmpty": lr.IsRightEmpty(),
	})
}

func Test_Cov57_LeftRight_WhitespaceChecks(t *testing.T) {
	lr := corestr.NewLeftRight("  ", "b")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR whitespace", args.Map{
			"LeftWS":  true,
			"RightWS": false,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"LeftWS":  lr.IsLeftWhitespace(),
		"RightWS": lr.IsRightWhitespace(),
	})
}

func Test_Cov57_LeftRight_ValidNonEmpty(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR valid non-empty", args.Map{
			"Left": true, "Right": true, "Safe": true,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"Left": lr.HasValidNonEmptyLeft(), "Right": lr.HasValidNonEmptyRight(), "Safe": lr.HasSafeNonEmpty(),
	})
}

func Test_Cov57_LeftRight_ValidNonWhitespace(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR valid non-ws", args.Map{"Left": true, "Right": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"Left": lr.HasValidNonWhitespaceLeft(), "Right": lr.HasValidNonWhitespaceRight(),
	})
}

func Test_Cov57_LeftRight_Is(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR Is", args.Map{
			"Is": true, "IsLeft": true, "IsRight": true,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"Is": lr.Is("a", "b"), "IsLeft": lr.IsLeft("a"), "IsRight": lr.IsRight("b"),
	})
}

func Test_Cov57_LeftRight_IsEqual(t *testing.T) {
	lr1 := corestr.NewLeftRight("a", "b")
	lr2 := corestr.NewLeftRight("a", "b")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR IsEqual", args.Map{"Equal": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Equal": lr1.IsEqual(lr2)})
}

func Test_Cov57_LeftRight_Clone(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	cloned := lr.Clone()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR Clone", args.Map{"Left": "a"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": cloned.Left})
}

func Test_Cov57_LeftRight_NonPtr(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	convey.Convey("NonPtr", t, func() {
		convey.So(lr.NonPtr().Left, convey.ShouldEqual, "a")
		convey.So(lr.Ptr(), convey.ShouldNotBeNil)
	})
}

func Test_Cov57_LeftRight_RegexMatch(t *testing.T) {
	lr := corestr.NewLeftRight("abc123", "xyz")
	re := regexp.MustCompile(`\d+`)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR Regex", args.Map{"LeftMatch": true, "RightMatch": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"LeftMatch": lr.IsLeftRegexMatch(re), "RightMatch": lr.IsRightRegexMatch(re),
	})
}

func Test_Cov57_LeftRight_RegexMatch_Nil(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR Regex nil", args.Map{"Left": false, "Right": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"Left": lr.IsLeftRegexMatch(nil), "Right": lr.IsRightRegexMatch(nil),
	})
}

func Test_Cov57_LeftRight_ClearDispose(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	lr.Clear()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR Clear", args.Map{"LeftEmpty": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"LeftEmpty": lr.IsLeftEmpty()})
}

func Test_Cov57_LeftRight_FromSplit(t *testing.T) {
	lr := corestr.LeftRightFromSplit("key=value", "=")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR FromSplit", args.Map{"Left": "key", "Right": "value"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lr.Left, "Right": lr.Right})
}

func Test_Cov57_LeftRight_FromSplitTrimmed(t *testing.T) {
	lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR FromSplitTrimmed", args.Map{"Left": "key", "Right": "value"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lr.Left, "Right": lr.Right})
}

func Test_Cov57_LeftRight_FromSplitFull(t *testing.T) {
	lr := corestr.LeftRightFromSplitFull("a:b:c", ":")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR FromSplitFull", args.Map{"Left": "a", "Right": "b:c"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lr.Left, "Right": lr.Right})
}

func Test_Cov57_LeftRight_FromSplitFullTrimmed(t *testing.T) {
	lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR FromSplitFullTrimmed", args.Map{"Left": "a"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lr.Left})
}

func Test_Cov57_LeftRight_UsingSlice(t *testing.T) {
	lr := corestr.LeftRightUsingSlice([]string{"a", "b"})
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR UsingSlice", args.Map{"Left": "a", "Right": "b"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lr.Left, "Right": lr.Right})
}

func Test_Cov57_LeftRight_UsingSlice_Single(t *testing.T) {
	lr := corestr.LeftRightUsingSlice([]string{"a"})
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR UsingSlice single", args.Map{"Left": "a", "IsValid": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lr.Left, "IsValid": lr.IsValid})
}

func Test_Cov57_LeftRight_UsingSlice_Empty(t *testing.T) {
	lr := corestr.LeftRightUsingSlice([]string{})
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR UsingSlice empty", args.Map{"IsValid": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsValid": lr.IsValid})
}

func Test_Cov57_LeftRight_UsingSlicePtr(t *testing.T) {
	lr := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR UsingSlicePtr", args.Map{"Left": "a"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lr.Left})
}

func Test_Cov57_LeftRight_TrimmedUsingSlice(t *testing.T) {
	lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR TrimmedUsingSlice", args.Map{"Left": "a", "Right": "b"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lr.Left, "Right": lr.Right})
}

func Test_Cov57_LeftRight_TrimmedUsingSlice_Nil(t *testing.T) {
	lr := corestr.LeftRightTrimmedUsingSlice(nil)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR TrimmedUsingSlice nil", args.Map{"IsValid": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsValid": lr.IsValid})
}

func Test_Cov57_LeftRight_TrimmedUsingSlice_Single(t *testing.T) {
	lr := corestr.LeftRightTrimmedUsingSlice([]string{" a "})
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LR TrimmedUsingSlice single", args.Map{"Left": " a ", "IsValid": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lr.Left, "IsValid": lr.IsValid})
}

// =============================================================================
// LeftMiddleRight
// =============================================================================

func Test_Cov57_LeftMiddleRight_New(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LMR New", args.Map{"Left": "a", "Middle": "b", "Right": "c"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lmr.Left, "Middle": lmr.Middle, "Right": lmr.Right})
}

func Test_Cov57_LeftMiddleRight_Invalid(t *testing.T) {
	lmr := corestr.InvalidLeftMiddleRight("err")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LMR Invalid", args.Map{"IsValid": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsValid": lmr.IsValid})
}

func Test_Cov57_LeftMiddleRight_InvalidNoMessage(t *testing.T) {
	lmr := corestr.InvalidLeftMiddleRightNoMessage()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LMR InvalidNoMessage", args.Map{"IsValid": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsValid": lmr.IsValid})
}

func Test_Cov57_LeftMiddleRight_Bytes(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	convey.Convey("LMR Bytes", t, func() {
		convey.So(lmr.LeftBytes(), convey.ShouldResemble, []byte("a"))
		convey.So(lmr.MiddleBytes(), convey.ShouldResemble, []byte("b"))
		convey.So(lmr.RightBytes(), convey.ShouldResemble, []byte("c"))
	})
}

func Test_Cov57_LeftMiddleRight_Trim(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight(" a ", " b ", " c ")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LMR Trim", args.Map{"Left": "a", "Mid": "b", "Right": "c"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lmr.LeftTrim(), "Mid": lmr.MiddleTrim(), "Right": lmr.RightTrim()})
}

func Test_Cov57_LeftMiddleRight_EmptyChecks(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("", "b", "")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LMR empty", args.Map{
			"LeftEmpty": true, "MidEmpty": false, "RightEmpty": true,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"LeftEmpty": lmr.IsLeftEmpty(), "MidEmpty": lmr.IsMiddleEmpty(), "RightEmpty": lmr.IsRightEmpty(),
	})
}

func Test_Cov57_LeftMiddleRight_WhitespaceChecks(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("  ", "b", "  ")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LMR ws", args.Map{
			"LeftWS": true, "MidWS": false, "RightWS": true,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"LeftWS": lmr.IsLeftWhitespace(), "MidWS": lmr.IsMiddleWhitespace(), "RightWS": lmr.IsRightWhitespace(),
	})
}

func Test_Cov57_LeftMiddleRight_ValidNonEmpty(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LMR valid non-empty", args.Map{
			"Left": true, "Mid": true, "Right": true, "Safe": true,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"Left": lmr.HasValidNonEmptyLeft(), "Mid": lmr.HasValidNonEmptyMiddle(),
		"Right": lmr.HasValidNonEmptyRight(), "Safe": lmr.HasSafeNonEmpty(),
	})
}

func Test_Cov57_LeftMiddleRight_ValidNonWhitespace(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LMR valid non-ws", args.Map{"Left": true, "Mid": true, "Right": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"Left": lmr.HasValidNonWhitespaceLeft(), "Mid": lmr.HasValidNonWhitespaceMiddle(),
		"Right": lmr.HasValidNonWhitespaceRight(),
	})
}

func Test_Cov57_LeftMiddleRight_IsAll(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LMR IsAll", args.Map{"IsAll": true, "Is": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsAll": lmr.IsAll("a", "b", "c"), "Is": lmr.Is("a", "c")})
}

func Test_Cov57_LeftMiddleRight_Clone(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	cloned := lmr.Clone()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LMR Clone", args.Map{"Left": "a"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": cloned.Left})
}

func Test_Cov57_LeftMiddleRight_ToLeftRight(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	lr := lmr.ToLeftRight()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LMR ToLeftRight", args.Map{"Left": "a", "Right": "c"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lr.Left, "Right": lr.Right})
}

func Test_Cov57_LeftMiddleRight_ClearDispose(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	lmr.Dispose()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LMR Dispose", args.Map{"LeftEmpty": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"LeftEmpty": lmr.IsLeftEmpty()})
}

func Test_Cov57_LeftMiddleRight_FromSplit(t *testing.T) {
	lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LMR FromSplit", args.Map{"Left": "a", "Mid": "b", "Right": "c"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lmr.Left, "Mid": lmr.Middle, "Right": lmr.Right})
}

func Test_Cov57_LeftMiddleRight_FromSplitTrimmed(t *testing.T) {
	lmr := corestr.LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LMR FromSplitTrimmed", args.Map{"Left": "a", "Mid": "b", "Right": "c"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lmr.Left, "Mid": lmr.Middle, "Right": lmr.Right})
}

func Test_Cov57_LeftMiddleRight_FromSplitN(t *testing.T) {
	lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d", ":")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LMR FromSplitN", args.Map{"Left": "a", "Mid": "b", "Right": "c:d"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lmr.Left, "Mid": lmr.Middle, "Right": lmr.Right})
}

func Test_Cov57_LeftMiddleRight_FromSplitNTrimmed(t *testing.T) {
	lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d ", ":")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("LMR FromSplitNTrimmed", args.Map{"Left": "a", "Mid": "b"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Left": lmr.Left, "Mid": lmr.Middle})
}

// =============================================================================
// ValidValue
// =============================================================================

func Test_Cov57_ValidValue_New(t *testing.T) {
	vv := corestr.NewValidValue("hello")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV New", args.Map{"Value": "hello", "IsValid": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Value": vv.Value, "IsValid": vv.IsValid})
}

func Test_Cov57_ValidValue_Empty(t *testing.T) {
	vv := corestr.NewValidValueEmpty()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV Empty", args.Map{"IsEmpty": true, "IsValid": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsEmpty": vv.IsEmpty(), "IsValid": vv.IsValid})
}

func Test_Cov57_ValidValue_Invalid(t *testing.T) {
	vv := corestr.InvalidValidValue("err")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV Invalid", args.Map{"IsValid": false, "Message": "err"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsValid": vv.IsValid, "Message": vv.Message})
}

func Test_Cov57_ValidValue_InvalidNoMessage(t *testing.T) {
	vv := corestr.InvalidValidValueNoMessage()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV InvalidNoMessage", args.Map{"IsValid": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsValid": vv.IsValid})
}

func Test_Cov57_ValidValue_Conversions(t *testing.T) {
	vv := corestr.NewValidValue("42")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV Conversions", args.Map{
			"Int": 42, "DefInt": 42, "Byte": byte(42), "Float": 42.0,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"Int": vv.ValueInt(0), "DefInt": vv.ValueDefInt(),
		"Byte": vv.ValueByte(0), "Float": vv.ValueFloat64(0),
	})
}

func Test_Cov57_ValidValue_Bool(t *testing.T) {
	vv := corestr.NewValidValue("true")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV Bool", args.Map{"Bool": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Bool": vv.ValueBool()})
}

func Test_Cov57_ValidValue_Bool_Empty(t *testing.T) {
	vv := corestr.NewValidValue("")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV Bool empty", args.Map{"Bool": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Bool": vv.ValueBool()})
}

func Test_Cov57_ValidValue_WhitespaceChecks(t *testing.T) {
	vv := corestr.NewValidValue("  ")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV ws", args.Map{
			"IsWS": true, "HasValidNonWS": false, "Trim": "",
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"IsWS": vv.IsWhitespace(), "HasValidNonWS": vv.HasValidNonWhitespace(), "Trim": vv.Trim(),
	})
}

func Test_Cov57_ValidValue_HasValidNonEmpty(t *testing.T) {
	vv := corestr.NewValidValue("a")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV HasValidNonEmpty", args.Map{"Result": true, "Safe": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Result": vv.HasValidNonEmpty(), "Safe": vv.HasSafeNonEmpty()})
}

func Test_Cov57_ValidValue_Is(t *testing.T) {
	vv := corestr.NewValidValue("a")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV Is", args.Map{"Is": true, "IsNot": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Is": vv.Is("a"), "IsNot": vv.Is("b")})
}

func Test_Cov57_ValidValue_IsAnyOf(t *testing.T) {
	vv := corestr.NewValidValue("b")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV IsAnyOf", args.Map{"Found": true, "NotFound": false, "Empty": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"Found": vv.IsAnyOf("a", "b"), "NotFound": vv.IsAnyOf("c"), "Empty": vv.IsAnyOf(),
	})
}

func Test_Cov57_ValidValue_IsContains(t *testing.T) {
	vv := corestr.NewValidValue("hello world")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV IsContains", args.Map{"Contains": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Contains": vv.IsContains("world")})
}

func Test_Cov57_ValidValue_IsAnyContains(t *testing.T) {
	vv := corestr.NewValidValue("hello world")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV IsAnyContains", args.Map{
			"Found": true, "NotFound": false, "Empty": true,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"Found": vv.IsAnyContains("world"), "NotFound": vv.IsAnyContains("xyz"), "Empty": vv.IsAnyContains(),
	})
}

func Test_Cov57_ValidValue_IsEqualNonSensitive(t *testing.T) {
	vv := corestr.NewValidValue("Hello")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV EqualFold", args.Map{"Result": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Result": vv.IsEqualNonSensitive("hello")})
}

func Test_Cov57_ValidValue_Regex(t *testing.T) {
	vv := corestr.NewValidValue("abc123")
	re := regexp.MustCompile(`\d+`)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV Regex", args.Map{
			"Matches": true, "FindStr": "123",
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"Matches": vv.IsRegexMatches(re), "FindStr": vv.RegexFindString(re),
	})
}

func Test_Cov57_ValidValue_Regex_Nil(t *testing.T) {
	vv := corestr.NewValidValue("abc")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV Regex nil", args.Map{"Matches": false, "FindStr": ""}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"Matches": vv.IsRegexMatches(nil), "FindStr": vv.RegexFindString(nil),
	})
}

func Test_Cov57_ValidValue_RegexFindAllStrings(t *testing.T) {
	vv := corestr.NewValidValue("a1b2c3")
	re := regexp.MustCompile(`\d`)
	result := vv.RegexFindAllStrings(re, -1)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV FindAllStrings", args.Map{"Length": 3}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(result)})
}

func Test_Cov57_ValidValue_RegexFindAllStrings_Nil(t *testing.T) {
	vv := corestr.NewValidValue("abc")
	result := vv.RegexFindAllStrings(nil, -1)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV FindAllStrings nil", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(result)})
}

func Test_Cov57_ValidValue_RegexFindAllStringsWithFlag(t *testing.T) {
	vv := corestr.NewValidValue("a1b2")
	re := regexp.MustCompile(`\d`)
	items, hasAny := vv.RegexFindAllStringsWithFlag(re, -1)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV FindAllWithFlag", args.Map{"Length": 2, "HasAny": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(items), "HasAny": hasAny})
}

func Test_Cov57_ValidValue_RegexFindAllStringsWithFlag_Nil(t *testing.T) {
	vv := corestr.NewValidValue("abc")
	_, hasAny := vv.RegexFindAllStringsWithFlag(nil, -1)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV FindAllWithFlag nil", args.Map{"HasAny": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"HasAny": hasAny})
}

func Test_Cov57_ValidValue_Split(t *testing.T) {
	vv := corestr.NewValidValue("a,b,c")
	result := vv.Split(",")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV Split", args.Map{"Length": 3}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(result)})
}

func Test_Cov57_ValidValue_Clone(t *testing.T) {
	vv := corestr.NewValidValue("a")
	cloned := vv.Clone()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV Clone", args.Map{"Value": "a"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Value": cloned.Value})
}

func Test_Cov57_ValidValue_String(t *testing.T) {
	vv := corestr.NewValidValue("a")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV String", "a"),
	}
	tc.ShouldBeEqual(0, t, vv.String())
}

func Test_Cov57_ValidValue_FullString(t *testing.T) {
	vv := corestr.NewValidValue("a")
	convey.Convey("VV FullString", t, func() {
		convey.So(vv.FullString(), convey.ShouldNotBeEmpty)
	})
}

func Test_Cov57_ValidValue_ClearDispose(t *testing.T) {
	vv := corestr.NewValidValue("a")
	vv.Dispose()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV Dispose", args.Map{"IsEmpty": true, "IsValid": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsEmpty": vv.IsEmpty(), "IsValid": vv.IsValid})
}

func Test_Cov57_ValidValue_ValueBytesOnce(t *testing.T) {
	vv := corestr.NewValidValue("abc")
	bytes1 := vv.ValueBytesOnce()
	bytes2 := vv.ValueBytesOnce()
	convey.Convey("VV ValueBytesOnce", t, func() {
		convey.So(bytes1, convey.ShouldResemble, []byte("abc"))
		convey.So(bytes2, convey.ShouldResemble, bytes1)
	})
}

func Test_Cov57_ValidValue_ValueBytesOncePtr(t *testing.T) {
	vv := corestr.NewValidValue("abc")
	convey.Convey("VV ValueBytesOncePtr", t, func() {
		convey.So(vv.ValueBytesOncePtr(), convey.ShouldResemble, []byte("abc"))
	})
}

func Test_Cov57_ValidValue_Serialize(t *testing.T) {
	vv := corestr.NewValidValue("a")
	bytes, err := vv.Serialize()
	convey.Convey("VV Serialize", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(bytes, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov57_ValidValue_NewUsingAny(t *testing.T) {
	vv := corestr.NewValidValueUsingAny(false, true, "hello")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV UsingAny", args.Map{"IsValid": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsValid": vv.IsValid})
}

func Test_Cov57_ValidValue_DefByte_Overflow(t *testing.T) {
	vv := corestr.NewValidValue("999")
	convey.Convey("VV DefByte overflow", t, func() {
		convey.So(vv.ValueDefByte(), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_Cov57_ValidValue_DefFloat(t *testing.T) {
	vv := corestr.NewValidValue("3.14")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VV DefFloat", args.Map{"Float": 3.14}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Float": vv.ValueDefFloat64()})
}

// =============================================================================
// ValidValues
// =============================================================================

func Test_Cov57_ValidValues_Basics(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.Add("a").Add("b")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs basics", args.Map{
			"Length": 2, "Count": 2, "HasAny": true, "IsEmpty": false,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"Length": vvs.Length(), "Count": vvs.Count(), "HasAny": vvs.HasAnyItem(), "IsEmpty": vvs.IsEmpty(),
	})
}

func Test_Cov57_ValidValues_AddFull(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.AddFull(false, "val", "msg")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs AddFull", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": vvs.Length()})
}

func Test_Cov57_ValidValues_SafeValueAt(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.Add("a")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs SafeValueAt", args.Map{"Value": "a", "Empty": ""}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Value": vvs.SafeValueAt(0), "Empty": vvs.SafeValueAt(99)})
}

func Test_Cov57_ValidValues_SafeValidValueAt(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.Add("a")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs SafeValidValueAt", args.Map{"Value": "a"}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Value": vvs.SafeValidValueAt(0)})
}

func Test_Cov57_ValidValues_SafeValuesAtIndexes(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.Add("a")
	result := vvs.SafeValuesAtIndexes(0, 99)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs SafeValuesAtIndexes", args.Map{"Length": 2}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(result)})
}

func Test_Cov57_ValidValues_SafeValidValuesAtIndexes(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.Add("a")
	result := vvs.SafeValidValuesAtIndexes(0)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs SafeValidValuesAtIndexes", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(result)})
}

func Test_Cov57_ValidValues_Strings(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.Add("a")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs Strings", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(vvs.Strings())})
}

func Test_Cov57_ValidValues_FullStrings(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.Add("a")
	convey.Convey("VVs FullStrings", t, func() {
		convey.So(len(vvs.FullStrings()), convey.ShouldEqual, 1)
	})
}

func Test_Cov57_ValidValues_String(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.Add("a")
	convey.Convey("VVs String", t, func() {
		convey.So(vvs.String(), convey.ShouldNotBeEmpty)
	})
}

func Test_Cov57_ValidValues_Find(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.Add("a").Add("b")
	result := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
		return vv, vv.Value == "a", false
	})
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs Find", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(result)})
}

func Test_Cov57_ValidValues_ConcatNew(t *testing.T) {
	vvs1 := corestr.EmptyValidValues()
	vvs1.Add("a")
	vvs2 := corestr.EmptyValidValues()
	vvs2.Add("b")
	result := vvs1.ConcatNew(false, vvs2)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs ConcatNew", args.Map{"Length": 2}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": result.Length()})
}

func Test_Cov57_ValidValues_ConcatNew_CloneOnEmpty(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.Add("a")
	result := vvs.ConcatNew(true)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs ConcatNew cloneOnEmpty", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": result.Length()})
}

func Test_Cov57_ValidValues_ConcatNew_NoClone(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.Add("a")
	result := vvs.ConcatNew(false)
	convey.Convey("VVs ConcatNew noClone returns self", t, func() {
		convey.So(result, convey.ShouldEqual, vvs)
	})
}

func Test_Cov57_ValidValues_AddHashsetMap(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.AddHashsetMap(map[string]bool{"a": true})
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs AddHashsetMap", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": vvs.Length()})
}

func Test_Cov57_ValidValues_AddHashsetMap_Nil(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.AddHashsetMap(nil)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs AddHashsetMap nil", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": vvs.Length()})
}

func Test_Cov57_ValidValues_AddHashset(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	hs := corestr.New.Hashset.Strings([]string{"a"})
	vvs.AddHashset(hs)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs AddHashset", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": vvs.Length()})
}

func Test_Cov57_ValidValues_AddHashset_Nil(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.AddHashset(nil)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs AddHashset nil", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": vvs.Length()})
}

func Test_Cov57_ValidValues_Hashmap(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.Add("a")
	hm := vvs.Hashmap()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs Hashmap", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": hm.Length()})
}

func Test_Cov57_ValidValues_Map(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.Add("a")
	m := vvs.Map()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs Map", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(m)})
}

func Test_Cov57_ValidValues_NewUsingValues(t *testing.T) {
	vvs := corestr.NewValidValuesUsingValues(
		corestr.ValidValue{Value: "a", IsValid: true},
	)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs NewUsingValues", args.Map{"Length": 1}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": vvs.Length()})
}

func Test_Cov57_ValidValues_NewUsingValues_Empty(t *testing.T) {
	vvs := corestr.NewValidValuesUsingValues()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs NewUsingValues empty", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": vvs.Length()})
}

func Test_Cov57_ValidValues_HasIndex(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.Add("a")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs HasIndex", args.Map{"Has": true, "Missing": false}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Has": vvs.HasIndex(0), "Missing": vvs.HasIndex(99)})
}

func Test_Cov57_ValidValues_AddValidValues(t *testing.T) {
	vvs1 := corestr.EmptyValidValues()
	vvs1.Add("a")
	vvs2 := corestr.EmptyValidValues()
	vvs2.Add("b")
	vvs1.AddValidValues(vvs2)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs AddValidValues", args.Map{"Length": 2}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": vvs1.Length()})
}

func Test_Cov57_ValidValues_AddValidValues_Nil(t *testing.T) {
	vvs := corestr.EmptyValidValues()
	vvs.AddValidValues(nil)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("VVs AddValidValues nil", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": vvs.Length()})
}

// =============================================================================
// ValueStatus
// =============================================================================

func Test_Cov57_ValueStatus_Invalid(t *testing.T) {
	vs := corestr.InvalidValueStatus("err")
	convey.Convey("VS Invalid", t, func() {
		convey.So(vs.ValueValid.IsValid, convey.ShouldBeFalse)
	})
}

func Test_Cov57_ValueStatus_InvalidNoMessage(t *testing.T) {
	vs := corestr.InvalidValueStatusNoMessage()
	convey.Convey("VS InvalidNoMessage", t, func() {
		convey.So(vs.ValueValid.IsValid, convey.ShouldBeFalse)
	})
}

func Test_Cov57_ValueStatus_Clone(t *testing.T) {
	vs := corestr.InvalidValueStatus("err")
	cloned := vs.Clone()
	convey.Convey("VS Clone", t, func() {
		convey.So(cloned.Index, convey.ShouldEqual, vs.Index)
	})
}

// =============================================================================
// TextWithLineNumber
// =============================================================================

func Test_Cov57_TextWithLineNumber(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("TWL", args.Map{
			"HasLine": true, "IsInvalid": false, "Length": 5, "IsEmpty": false, "IsEmptyText": false,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"HasLine": twl.HasLineNumber(), "IsInvalid": twl.IsInvalidLineNumber(),
		"Length": twl.Length(), "IsEmpty": twl.IsEmpty(), "IsEmptyText": twl.IsEmptyText(),
	})
}

func Test_Cov57_TextWithLineNumber_Empty(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("TWL empty", args.Map{
			"IsEmpty": true, "IsEmptyBoth": true,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"IsEmpty": twl.IsEmpty(), "IsEmptyBoth": twl.IsEmptyTextLineBoth(),
	})
}

// =============================================================================
// NonChainedLinkedListNodes
// =============================================================================

func Test_Cov57_NonChainedLinkedListNodes(t *testing.T) {
	nodes := corestr.NewNonChainedLinkedListNodes(2)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("NCLLN empty", args.Map{
			"IsEmpty": true, "Length": 0, "IsChainingApplied": false,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"IsEmpty": nodes.IsEmpty(), "Length": nodes.Length(), "IsChainingApplied": nodes.IsChainingApplied(),
	})
}

func Test_Cov57_NonChainedLinkedListNodes_Adds(t *testing.T) {
	nodes := corestr.NewNonChainedLinkedListNodes(2)
	ll := corestr.New.LinkedList.Strings("a", "b")
	nodes.Adds(ll.Head())
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("NCLLN Adds", args.Map{"Length": 1, "HasItems": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": nodes.Length(), "HasItems": nodes.HasItems()})
}

func Test_Cov57_NonChainedLinkedListNodes_FirstLast(t *testing.T) {
	nodes := corestr.NewNonChainedLinkedListNodes(2)
	ll := corestr.New.LinkedList.Strings("a", "b")
	nodes.Adds(ll.Head(), ll.Head().Next())
	convey.Convey("NCLLN FirstLast", t, func() {
		convey.So(nodes.First(), convey.ShouldNotBeNil)
		convey.So(nodes.Last(), convey.ShouldNotBeNil)
		convey.So(nodes.FirstOrDefault(), convey.ShouldNotBeNil)
		convey.So(nodes.LastOrDefault(), convey.ShouldNotBeNil)
	})
}

func Test_Cov57_NonChainedLinkedListNodes_FirstOrDefault_Empty(t *testing.T) {
	nodes := corestr.NewNonChainedLinkedListNodes(2)
	convey.Convey("NCLLN FirstOrDefault empty", t, func() {
		convey.So(nodes.FirstOrDefault(), convey.ShouldBeNil)
		convey.So(nodes.LastOrDefault(), convey.ShouldBeNil)
	})
}

func Test_Cov57_NonChainedLinkedListNodes_ApplyChaining(t *testing.T) {
	nodes := corestr.NewNonChainedLinkedListNodes(2)
	ll := corestr.New.LinkedList.Strings("a", "b")
	nodes.Adds(ll.Head(), ll.Head().Next())
	nodes.ApplyChaining()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("NCLLN ApplyChaining", args.Map{"IsChainingApplied": true}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"IsChainingApplied": nodes.IsChainingApplied()})
}

func Test_Cov57_NonChainedLinkedListNodes_Items(t *testing.T) {
	nodes := corestr.NewNonChainedLinkedListNodes(2)
	convey.Convey("NCLLN Items", t, func() {
		convey.So(nodes.Items(), convey.ShouldNotBeNil)
	})
}

// =============================================================================
// NonChainedLinkedCollectionNodes
// =============================================================================

func Test_Cov57_NonChainedLinkedCollectionNodes(t *testing.T) {
	nodes := corestr.NewNonChainedLinkedCollectionNodes(2)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("NCLCN empty", args.Map{
			"IsEmpty": true, "Length": 0, "IsChainingApplied": false,
		}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{
		"IsEmpty": nodes.IsEmpty(), "Length": nodes.Length(), "IsChainingApplied": nodes.IsChainingApplied(),
	})
}

func Test_Cov57_NonChainedLinkedCollectionNodes_FirstOrDefault_Empty(t *testing.T) {
	nodes := corestr.NewNonChainedLinkedCollectionNodes(2)
	convey.Convey("NCLCN FirstOrDefault empty", t, func() {
		convey.So(nodes.FirstOrDefault(), convey.ShouldBeNil)
		convey.So(nodes.LastOrDefault(), convey.ShouldBeNil)
	})
}

func Test_Cov57_NonChainedLinkedCollectionNodes_Items(t *testing.T) {
	nodes := corestr.NewNonChainedLinkedCollectionNodes(2)
	convey.Convey("NCLCN Items", t, func() {
		convey.So(nodes.Items(), convey.ShouldNotBeNil)
	})
}

// =============================================================================
// CloneSlice / CloneSliceIf
// =============================================================================

func Test_Cov57_CloneSlice(t *testing.T) {
	result := corestr.CloneSlice([]string{"a", "b"})
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("CloneSlice", args.Map{"Length": 2}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(result)})
}

func Test_Cov57_CloneSlice_Empty(t *testing.T) {
	result := corestr.CloneSlice(nil)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("CloneSlice empty", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(result)})
}

func Test_Cov57_CloneSliceIf_True(t *testing.T) {
	result := corestr.CloneSliceIf(true, "a", "b")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("CloneSliceIf true", args.Map{"Length": 2}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(result)})
}

func Test_Cov57_CloneSliceIf_False(t *testing.T) {
	result := corestr.CloneSliceIf(false, "a", "b")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("CloneSliceIf false", args.Map{"Length": 2}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(result)})
}

func Test_Cov57_CloneSliceIf_Empty(t *testing.T) {
	result := corestr.CloneSliceIf(true)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("CloneSliceIf empty", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": len(result)})
}

// =============================================================================
// AnyToString
// =============================================================================

func Test_Cov57_AnyToString(t *testing.T) {
	result := corestr.AnyToString(false, "hello")
	convey.Convey("AnyToString", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov57_AnyToString_WithFieldName(t *testing.T) {
	result := corestr.AnyToString(true, "hello")
	convey.Convey("AnyToString with field name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov57_AnyToString_Empty(t *testing.T) {
	result := corestr.AnyToString(false, "")
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AnyToString empty", ""),
	}
	tc.ShouldBeEqual(0, t, result)
}

// =============================================================================
// AllIndividualStringsOfStringsLength
// =============================================================================

func Test_Cov57_AllIndividualStringsOfStringsLength(t *testing.T) {
	items := [][]string{{"a", "b"}, {"c"}}
	result := corestr.AllIndividualStringsOfStringsLength(&items)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AllIndividualStringsOfStringsLength", args.Map{"Length": 3}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": result})
}

func Test_Cov57_AllIndividualStringsOfStringsLength_Nil(t *testing.T) {
	result := corestr.AllIndividualStringsOfStringsLength(nil)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AllIndividualStringsOfStringsLength nil", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": result})
}

// =============================================================================
// AllIndividualsLengthOfSimpleSlices
// =============================================================================

func Test_Cov57_AllIndividualsLengthOfSimpleSlices(t *testing.T) {
	ss1 := corestr.New.SimpleSlice.Strings([]string{"a"})
	ss2 := corestr.New.SimpleSlice.Strings([]string{"b", "c"})
	result := corestr.AllIndividualsLengthOfSimpleSlices(ss1, ss2)
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AllIndividualsLengthOfSimpleSlices", args.Map{"Length": 3}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": result})
}

func Test_Cov57_AllIndividualsLengthOfSimpleSlices_Nil(t *testing.T) {
	result := corestr.AllIndividualsLengthOfSimpleSlices()
	tc := coretestcases.CaseV1{
		BaseTestCase: coretestcases.BaseFields("AllIndividualsLengthOfSimpleSlices nil", args.Map{"Length": 0}),
	}
	tc.ShouldBeEqualMap(0, t, args.Map{"Length": result})
}
