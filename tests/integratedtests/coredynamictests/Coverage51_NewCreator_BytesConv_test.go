package coredynamictests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// New.Collection creators (newCreator → newCollectionCreator → generic/typed)
// =============================================================================

func Test_Cov51_New_Collection_String_Empty(t *testing.T) {
	c := coredynamic.New.Collection.String.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Empty", actual)
}

func Test_Cov51_New_Collection_String_Cap(t *testing.T) {
	c := coredynamic.New.Collection.String.Cap(10)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Cap", actual)
}

func Test_Cov51_New_Collection_String_From(t *testing.T) {
	c := coredynamic.New.Collection.String.From([]string{"a", "b"})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.From", actual)
}

func Test_Cov51_New_Collection_String_Clone(t *testing.T) {
	c := coredynamic.New.Collection.String.Clone([]string{"a"})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Clone", actual)
}

func Test_Cov51_New_Collection_String_Items(t *testing.T) {
	c := coredynamic.New.Collection.String.Items("a", "b", "c")
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Items", actual)
}

func Test_Cov51_New_Collection_String_Create(t *testing.T) {
	c := coredynamic.New.Collection.String.Create([]string{"x"})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Create", actual)
}

func Test_Cov51_New_Collection_String_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.String.LenCap(3, 10)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.LenCap", actual)
}

func Test_Cov51_New_Collection_Int_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Int.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.Empty", actual)
}

func Test_Cov51_New_Collection_Int_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.Int.LenCap(5, 10)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 5}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.LenCap", actual)
}

func Test_Cov51_New_Collection_Int64_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.Int64.LenCap(2, 8)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int64.LenCap", actual)
}

func Test_Cov51_New_Collection_Byte_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.Byte.LenCap(4, 16)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "New.Collection.Byte.LenCap", actual)
}

func Test_Cov51_New_Collection_Any_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Any.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Any.Empty", actual)
}

func Test_Cov51_New_Collection_Any_Cap(t *testing.T) {
	c := coredynamic.New.Collection.Any.Cap(5)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Any.Cap", actual)
}

func Test_Cov51_New_Collection_Bool_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Bool.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Bool.Empty", actual)
}

func Test_Cov51_New_Collection_Float32_From(t *testing.T) {
	c := coredynamic.New.Collection.Float32.From([]float32{1.5, 2.5})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.Float32.From", actual)
}

func Test_Cov51_New_Collection_Float64_Items(t *testing.T) {
	c := coredynamic.New.Collection.Float64.Items(1.0, 2.0)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.Float64.Items", actual)
}

func Test_Cov51_New_Collection_ByteSlice_Empty(t *testing.T) {
	c := coredynamic.New.Collection.ByteSlice.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.ByteSlice.Empty", actual)
}

func Test_Cov51_New_Collection_AnyMap_Empty(t *testing.T) {
	c := coredynamic.New.Collection.AnyMap.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.AnyMap.Empty", actual)
}

func Test_Cov51_New_Collection_StringMap_Empty(t *testing.T) {
	c := coredynamic.New.Collection.StringMap.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.StringMap.Empty", actual)
}

func Test_Cov51_New_Collection_IntMap_Empty(t *testing.T) {
	c := coredynamic.New.Collection.IntMap.Empty()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.IntMap.Empty", actual)
}

// =============================================================================
// BytesConverter — remaining branches
// =============================================================================

func Test_Cov51_BytesConverter_ToBool_Valid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("true"))
	r, err := bc.ToBool()
	actual := args.Map{"noErr": err == nil, "r": r}
	expected := args.Map{"noErr": true, "r": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToBool valid", actual)
}

func Test_Cov51_BytesConverter_ToBoolMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("false"))
	actual := args.Map{"r": bc.ToBoolMust()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToBoolMust", actual)
}

func Test_Cov51_BytesConverter_SafeCastString_Empty(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte{})
	actual := args.Map{"r": bc.SafeCastString()}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "BytesConverter SafeCastString empty", actual)
}

func Test_Cov51_BytesConverter_SafeCastString_Valid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("hello"))
	actual := args.Map{"r": bc.SafeCastString()}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesConverter SafeCastString valid", actual)
}

func Test_Cov51_BytesConverter_CastString_Empty(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte{})
	_, err := bc.CastString()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter CastString empty", actual)
}

func Test_Cov51_BytesConverter_CastString_Valid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("hello"))
	s, err := bc.CastString()
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesConverter CastString valid", actual)
}

func Test_Cov51_BytesConverter_ToString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	s, err := bc.ToString()
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToString", actual)
}

func Test_Cov51_BytesConverter_ToStringMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"world"`))
	actual := args.Map{"r": bc.ToStringMust()}
	expected := args.Map{"r": "world"}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStringMust", actual)
}

func Test_Cov51_BytesConverter_ToStrings(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	ss, err := bc.ToStrings()
	actual := args.Map{"noErr": err == nil, "len": len(ss)}
	expected := args.Map{"noErr": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStrings", actual)
}

func Test_Cov51_BytesConverter_ToStringsMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["x"]`))
	ss := bc.ToStringsMust()
	actual := args.Map{"len": len(ss)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStringsMust", actual)
}

func Test_Cov51_BytesConverter_ToInt64(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("42"))
	v, err := bc.ToInt64()
	actual := args.Map{"noErr": err == nil, "v": v}
	expected := args.Map{"noErr": true, "v": int64(42)}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToInt64", actual)
}

func Test_Cov51_BytesConverter_ToInt64Must(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("99"))
	actual := args.Map{"v": bc.ToInt64Must()}
	expected := args.Map{"v": int64(99)}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToInt64Must", actual)
}

func Test_Cov51_BytesConverter_ToHashmap_Valid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`{"key":"val"}`))
	h, err := bc.ToHashmap()
	actual := args.Map{"noErr": err == nil, "notNil": h != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToHashmap valid", actual)
}

func Test_Cov51_BytesConverter_ToHashmap_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToHashmap()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToHashmap invalid", actual)
}

func Test_Cov51_BytesConverter_ToHashset_Valid(t *testing.T) {
	b, _ := json.Marshal(map[string]bool{"a": true, "b": true})
	bc := coredynamic.NewBytesConverter(b)
	h, err := bc.ToHashset()
	actual := args.Map{"noErr": err == nil, "notNil": h != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToHashset valid", actual)
}

func Test_Cov51_BytesConverter_ToHashset_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToHashset()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToHashset invalid", actual)
}

func Test_Cov51_BytesConverter_ToCollection_Valid(t *testing.T) {
	b, _ := json.Marshal([]string{"a"})
	bc := coredynamic.NewBytesConverter(b)
	c, err := bc.ToCollection()
	actual := args.Map{"noErr": err == nil, "notNil": c != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToCollection valid", actual)
}

func Test_Cov51_BytesConverter_ToCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToCollection invalid", actual)
}

func Test_Cov51_BytesConverter_ToSimpleSlice_Valid(t *testing.T) {
	b, _ := json.Marshal([]string{"a"})
	bc := coredynamic.NewBytesConverter(b)
	s, err := bc.ToSimpleSlice()
	actual := args.Map{"noErr": err == nil, "notNil": s != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToSimpleSlice valid", actual)
}

func Test_Cov51_BytesConverter_ToSimpleSlice_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToSimpleSlice()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToSimpleSlice invalid", actual)
}

func Test_Cov51_BytesConverter_ToKeyValCollection_Valid(t *testing.T) {
	b, _ := json.Marshal(map[string]any{"Items": []map[string]any{{"Key": "a", "Value": 1}}})
	bc := coredynamic.NewBytesConverter(b)
	c, err := bc.ToKeyValCollection()
	actual := args.Map{"noErr": err == nil, "notNil": c != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToKeyValCollection valid", actual)
}

func Test_Cov51_BytesConverter_ToKeyValCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToKeyValCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToKeyValCollection invalid", actual)
}

func Test_Cov51_BytesConverter_ToAnyCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToAnyCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToAnyCollection invalid", actual)
}

func Test_Cov51_BytesConverter_ToMapAnyItems_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToMapAnyItems()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToMapAnyItems invalid", actual)
}

func Test_Cov51_BytesConverter_ToDynamicCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToDynamicCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToDynamicCollection invalid", actual)
}

func Test_Cov51_BytesConverter_ToJsonResultCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToJsonResultCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToJsonResultCollection invalid", actual)
}

func Test_Cov51_BytesConverter_ToJsonMapResults_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToJsonMapResults()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToJsonMapResults invalid", actual)
}

func Test_Cov51_BytesConverter_ToBytesCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToBytesCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToBytesCollection invalid", actual)
}

func Test_Cov51_BytesConverter_Deserialize(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	var s string
	err := bc.Deserialize(&s)
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesConverter Deserialize", actual)
}

func Test_Cov51_BytesConverter_DeserializeMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"world"`))
	var s string
	bc.DeserializeMust(&s)
	actual := args.Map{"r": s}
	expected := args.Map{"r": "world"}
	expected.ShouldBeEqual(t, 0, "BytesConverter DeserializeMust", actual)
}

func Test_Cov51_NewBytesConverterUsingJsonResult_Valid(t *testing.T) {
	r := corejson.New("test")
	bc, err := coredynamic.NewBytesConverterUsingJsonResult(r.Ptr())
	actual := args.Map{"noErr": err == nil, "notNil": bc != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "NewBytesConverterUsingJsonResult valid", actual)
}

func Test_Cov51_NewBytesConverterUsingJsonResult_Invalid(t *testing.T) {
	r := corejson.Result{}
	_, err := coredynamic.NewBytesConverterUsingJsonResult(&r)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewBytesConverterUsingJsonResult invalid", actual)
}

// =============================================================================
// DynamicCollectionModel
// =============================================================================

func Test_Cov51_DynamicCollectionModel(t *testing.T) {
	dcm := coredynamic.DynamicCollectionModel{
		Items: []coredynamic.Dynamic{
			*coredynamic.NewDynamicPtr("a", true),
		},
	}
	actual := args.Map{"len": len(dcm.Items)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollectionModel", actual)
}
