package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── BasicByte via newBasicByteCreator ──

func Test_Cov3_BasicByte_Create_Default(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active", "Inactive"},
	)

	actual := args.Map{
		"min":          bb.Min(),
		"max":          bb.Max(),
		"isValidRange": bb.IsValidRange(1),
		"outOfRange":   bb.IsValidRange(5),
		"toString":     bb.ToEnumString(1),
		"typeName":     bb.TypeName(),
		"length":       bb.Length(),
		"count":        bb.Count(),
	}
	expected := args.Map{
		"min":          byte(0),
		"max":          byte(2),
		"isValidRange": true,
		"outOfRange":   false,
		"toString":     "Active",
		"typeName":     "enumimpltests.myEnum",
		"length":       3,
		"count":        3,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_Create_Default", actual)
}

func Test_Cov3_BasicByte_IsAnyOf(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active", "Inactive"},
	)

	actual := args.Map{
		"isAnyEmpty":   bb.IsAnyOf(1),
		"isAnyMatch":  bb.IsAnyOf(1, 0, 1, 2),
		"isAnyNoMatch": bb.IsAnyOf(1, 0, 2),
	}
	expected := args.Map{
		"isAnyEmpty":   true,
		"isAnyMatch":  true,
		"isAnyNoMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_IsAnyOf", actual)
}

func Test_Cov3_BasicByte_IsAnyNamesOf(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active", "Inactive"},
	)

	actual := args.Map{
		"matchName":   bb.IsAnyNamesOf(1, "Active"),
		"noMatchName": bb.IsAnyNamesOf(1, "Invalid"),
	}
	expected := args.Map{
		"matchName":   true,
		"noMatchName": false,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_IsAnyNamesOf", actual)
}

func Test_Cov3_BasicByte_GetValueByString(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active", "Inactive"},
	)

	actual := args.Map{
		"byName": bb.GetValueByString("Active"),
	}
	expected := args.Map{
		"byName": byte(1),
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_GetValueByString", actual)
}

func Test_Cov3_BasicByte_GetValueByName(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	val, err := bb.GetValueByName("Active")
	_, errNotFound := bb.GetValueByName("NotExist")

	actual := args.Map{
		"val":      val,
		"noErr":    err == nil,
		"hasError": errNotFound != nil,
	}
	expected := args.Map{
		"val":      byte(1),
		"noErr":    true,
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_GetValueByName", actual)
}

func Test_Cov3_BasicByte_GetStringValue(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	actual := args.Map{
		"val": bb.GetStringValue(0),
	}
	expected := args.Map{
		"val": "Invalid",
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_GetStringValue", actual)
}

func Test_Cov3_BasicByte_Ranges(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	actual := args.Map{
		"rangesLen": len(bb.Ranges()),
	}
	expected := args.Map{
		"rangesLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_Ranges", actual)
}

func Test_Cov3_BasicByte_Hashmap(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	hm := bb.Hashmap()
	hmPtr := bb.HashmapPtr()

	actual := args.Map{
		"hmHasItems":    len(hm) > 0,
		"hmPtrNotNil":   hmPtr != nil,
	}
	expected := args.Map{
		"hmHasItems":    true,
		"hmPtrNotNil":   true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_Hashmap", actual)
}

func Test_Cov3_BasicByte_ToEnumJsonBytes(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	jsonBytes, err := bb.ToEnumJsonBytes(0)
	_, errNotFound := bb.ToEnumJsonBytes(99)

	actual := args.Map{
		"hasBytes":  len(jsonBytes) > 0,
		"noErr":     err == nil,
		"notFound":  errNotFound != nil,
	}
	expected := args.Map{
		"hasBytes":  true,
		"noErr":     true,
		"notFound":  true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_ToEnumJsonBytes", actual)
}

func Test_Cov3_BasicByte_AppendPrependJoinValue(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	result := bb.AppendPrependJoinValue(".", 1, 0)

	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_AppendPrependJoinValue", actual)
}

func Test_Cov3_BasicByte_ToNumberString(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	actual := args.Map{
		"notEmpty": bb.ToNumberString(1) != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_ToNumberString", actual)
}

func Test_Cov3_BasicByte_UnmarshallToValue(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	// nil with mapping to first
	val1, err1 := bb.UnmarshallToValue(true, nil)
	// nil without mapping
	_, err2 := bb.UnmarshallToValue(false, nil)
	// empty string with mapping
	val3, err3 := bb.UnmarshallToValue(true, []byte(""))
	// valid name
	val4, err4 := bb.UnmarshallToValue(false, []byte("Active"))

	actual := args.Map{
		"nilMapped":    val1,
		"nilMappedErr": err1 == nil,
		"nilNoMap":     err2 != nil,
		"emptyMapped":  val3,
		"emptyNoErr":   err3 == nil,
		"validVal":     val4,
		"validNoErr":   err4 == nil,
	}
	expected := args.Map{
		"nilMapped":    byte(0),
		"nilMappedErr": true,
		"nilNoMap":     true,
		"emptyMapped":  byte(0),
		"emptyNoErr":   true,
		"validVal":     byte(1),
		"validNoErr":   true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_UnmarshallToValue", actual)
}

func Test_Cov3_BasicByte_EnumType(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	actual := args.Map{
		"enumType": bb.EnumType().String(),
	}
	expected := args.Map{
		"enumType": "Byte",
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_EnumType", actual)
}

func Test_Cov3_BasicByte_AsBasicByter(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	byter := bb.AsBasicByter()

	actual := args.Map{
		"notNil": byter != nil,
	}
	expected := args.Map{
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_AsBasicByter", actual)
}

// ── BasicByte with alias map ──

func Test_Cov3_BasicByte_WithAliasMap(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.DefaultWithAliasMap(
		myEnum(0),
		[]string{"Invalid", "Active"},
		map[string]byte{"on": 1},
	)

	val, err := bb.GetValueByName("on")

	actual := args.Map{
		"aliasVal": val,
		"noErr":    err == nil,
	}
	expected := args.Map{
		"aliasVal": byte(1),
		"noErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_WithAliasMap", actual)
}

func Test_Cov3_BasicByte_AllCases(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.DefaultAllCases(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	val, err := bb.GetValueByName("active")

	actual := args.Map{
		"lowerVal": val,
		"noErr":    err == nil,
	}
	expected := args.Map{
		"lowerVal": byte(1),
		"noErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_AllCases", actual)
}

// ── numberEnumBase methods via BasicByte ──

func Test_Cov3_NumberEnumBase_MinMaxAny(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	min, max := bb.MinMaxAny()

	actual := args.Map{
		"min": min,
		"max": max,
	}
	expected := args.Map{
		"min": byte(0),
		"max": byte(1),
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_MinMaxAny", actual)
}

func Test_Cov3_NumberEnumBase_ValueStrings(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	actual := args.Map{
		"minValStr":     bb.MinValueString() != "",
		"maxValStr":     bb.MaxValueString() != "",
		"minInt":        bb.MinInt(),
		"maxInt":        bb.MaxInt(),
		"rangesCsv":     bb.RangeNamesCsv() != "",
		"rangesInvalid": bb.RangesInvalidMessage() != "",
		"rangesErr":     bb.RangesInvalidErr() != nil,
	}
	expected := args.Map{
		"minValStr":     true,
		"maxValStr":     true,
		"minInt":        0,
		"maxInt":        1,
		"rangesCsv":     true,
		"rangesInvalid": true,
		"rangesErr":     true,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_ValueStrings", actual)
}

func Test_Cov3_NumberEnumBase_StringRanges(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	actual := args.Map{
		"strRangesLen":    len(bb.StringRanges()),
		"strRangesPtrLen": len(bb.StringRangesPtr()),
		"namesHashLen":    len(bb.NamesHashset()),
	}
	expected := args.Map{
		"strRangesLen":    2,
		"strRangesPtrLen": 2,
		"namesHashLen":    2,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_StringRanges", actual)
}

func Test_Cov3_NumberEnumBase_DynamicMap(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	dm := bb.RangesDynamicMap()
	dm2 := bb.DynamicMap()

	actual := args.Map{
		"dmLen":  len(dm),
		"dm2Len": len(dm2),
	}
	expected := args.Map{
		"dmLen":  2,
		"dm2Len": 2,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_DynamicMap", actual)
}

func Test_Cov3_NumberEnumBase_IntegerEnumRanges(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	ranges := bb.IntegerEnumRanges()

	actual := args.Map{
		"len": len(ranges),
	}
	expected := args.Map{
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_IntegerEnumRanges", actual)
}

func Test_Cov3_NumberEnumBase_AllNameValues(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	anv := bb.AllNameValues()

	actual := args.Map{
		"len": len(anv),
	}
	expected := args.Map{
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_AllNameValues", actual)
}

func Test_Cov3_NumberEnumBase_KeyAnyValues(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	kav := bb.KeyAnyValues()
	kvi := bb.KeyValIntegers()

	actual := args.Map{
		"kavLen": len(kav),
		"kviLen": len(kvi),
	}
	expected := args.Map{
		"kavLen": 2,
		"kviLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_KeyAnyValues", actual)
}

func Test_Cov3_NumberEnumBase_Loop(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	count := 0
	bb.Loop(func(index int, name string, anyVal any) (isBreak bool) {
		count++
		return false
	})

	intCount := 0
	bb.LoopInteger(func(index int, name string, anyVal int) (isBreak bool) {
		intCount++
		return false
	})

	actual := args.Map{
		"loopCount":    count,
		"intLoopCount": intCount,
	}
	expected := args.Map{
		"loopCount":    2,
		"intLoopCount": 2,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_Loop", actual)
}

func Test_Cov3_NumberEnumBase_LoopBreak(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active", "Inactive"},
	)

	count := 0
	bb.Loop(func(index int, name string, anyVal any) (isBreak bool) {
		count++
		return true // break after first
	})

	actual := args.Map{
		"breakCount": count,
	}
	expected := args.Map{
		"breakCount": 1,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_LoopBreak", actual)
}

func Test_Cov3_NumberEnumBase_Format(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	result := bb.Format("Enum of {type-name} - {name} - {value}", byte(1))

	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_Format", actual)
}

func Test_Cov3_NumberEnumBase_NameWithValue(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	nv := bb.NameWithValue(byte(1))
	nvOpt := bb.NameWithValueOption(byte(1), true)
	nvOptNo := bb.NameWithValueOption(byte(1), false)

	actual := args.Map{
		"nv":      nv != "",
		"nvOpt":   nvOpt != "",
		"nvOptNo": nvOptNo != "",
	}
	expected := args.Map{
		"nv":      true,
		"nvOpt":   true,
		"nvOptNo": true,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_NameWithValue", actual)
}

func Test_Cov3_NumberEnumBase_RangesMap(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	rm := bb.RangesMap()
	rism := bb.RangesIntegerStringMap()

	actual := args.Map{
		"rmLen":   len(rm),
		"rismLen": len(rism),
	}
	expected := args.Map{
		"rmLen":   2,
		"rismLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_RangesMap", actual)
}

func Test_Cov3_NumberEnumBase_OnlySupportedErr(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	noErr := bb.OnlySupportedErr("Invalid", "Active")
	hasErr := bb.OnlySupportedErr("Invalid")
	msgErr := bb.OnlySupportedMsgErr("context", "Invalid")

	actual := args.Map{
		"noErr":  noErr == nil,
		"hasErr": hasErr != nil,
		"msgErr": msgErr != nil,
	}
	expected := args.Map{
		"noErr":  true,
		"hasErr": true,
		"msgErr": true,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_OnlySupportedErr", actual)
}

func Test_Cov3_NumberEnumBase_JsonString(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	js := bb.JsonString(byte(1))
	es := bb.ToEnumString(byte(1))
	tn := bb.ToName(byte(1))

	actual := args.Map{
		"js": js != "",
		"es": es != "",
		"tn": tn != "",
	}
	expected := args.Map{
		"js": true,
		"es": true,
		"tn": true,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_JsonString", actual)
}

// ── BasicString ──

func Test_Cov3_BasicString_Create(t *testing.T) {
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active", "Inactive"},
	)

	actual := args.Map{
		"min":           bs.Min(),
		"max":           bs.Max(),
		"length":        bs.Length(),
		"hasAnyItem":    bs.HasAnyItem(),
		"maxIndex":      bs.MaxIndex(),
		"isValidActive": bs.IsValidRange("Active"),
		"isValidBad":    bs.IsValidRange("NotExist"),
	}
	expected := args.Map{
		"min":           bs.Min(),
		"max":           bs.Max(),
		"length":        3,
		"hasAnyItem":    true,
		"maxIndex":      2,
		"isValidActive": true,
		"isValidBad":    false,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_Create", actual)
}

func Test_Cov3_BasicString_IsAnyOf(t *testing.T) {
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	actual := args.Map{
		"empty":   bs.IsAnyOf("x"),
		"match":   bs.IsAnyOf("Active", "Active", "Invalid"),
		"noMatch": bs.IsAnyOf("Active", "Invalid"),
	}
	expected := args.Map{
		"empty":   true,
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_IsAnyOf", actual)
}

func Test_Cov3_BasicString_IsAnyNamesOf(t *testing.T) {
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	actual := args.Map{
		"match":   bs.IsAnyNamesOf("Active", "Active"),
		"noMatch": bs.IsAnyNamesOf("Active", "Invalid"),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_IsAnyNamesOf", actual)
}

func Test_Cov3_BasicString_GetNameByIndex(t *testing.T) {
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active", "Inactive"},
	)

	actual := args.Map{
		"valid":   bs.GetNameByIndex(1),
		"invalid": bs.GetNameByIndex(99),
	}
	expected := args.Map{
		"valid":   "Active",
		"invalid": "",
	}
	expected.ShouldBeEqual(t, 0, "BasicString_GetNameByIndex", actual)
}

func Test_Cov3_BasicString_GetIndexByName(t *testing.T) {
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	actual := args.Map{
		"found":   bs.GetIndexByName("Active"),
		"empty":   bs.GetIndexByName(""),
		"missing": bs.GetIndexByName("NotExist"),
	}
	expected := args.Map{
		"found":   1,
		"empty":   -1,
		"missing": -1,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_GetIndexByName", actual)
}

func Test_Cov3_BasicString_Ranges(t *testing.T) {
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	actual := args.Map{
		"rangesLen":   len(bs.Ranges()),
		"hashsetLen":  len(bs.Hashset()),
		"hashsetPtr":  bs.HashsetPtr() != nil,
		"integersLen": len(bs.RangesIntegers()),
		"nameIdxLen":  len(bs.NameWithIndexMap()),
	}
	expected := args.Map{
		"rangesLen":   2,
		"hashsetLen":  len(bs.Hashset()),
		"hashsetPtr":  true,
		"integersLen": 2,
		"nameIdxLen":  2,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_Ranges", actual)
}

func Test_Cov3_BasicString_GetValueByName(t *testing.T) {
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	val, err := bs.GetValueByName("Active")
	_, errNotFound := bs.GetValueByName("NotExist")

	actual := args.Map{
		"val":      val,
		"noErr":    err == nil,
		"hasError": errNotFound != nil,
	}
	expected := args.Map{
		"val":      "Active",
		"noErr":    true,
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_GetValueByName", actual)
}

func Test_Cov3_BasicString_ToEnumJsonBytes(t *testing.T) {
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	jsonBytes, err := bs.ToEnumJsonBytes("Active")
	_, errBad := bs.ToEnumJsonBytes("NotExist")

	actual := args.Map{
		"hasBytes":  len(jsonBytes) > 0,
		"noErr":     err == nil,
		"errOnBad":  errBad != nil,
	}
	expected := args.Map{
		"hasBytes":  true,
		"noErr":     true,
		"errOnBad":  true,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_ToEnumJsonBytes", actual)
}

func Test_Cov3_BasicString_UnmarshallToValue(t *testing.T) {
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	val1, err1 := bs.UnmarshallToValue(true, nil)
	_, err2 := bs.UnmarshallToValue(false, nil)
	val3, err3 := bs.UnmarshallToValue(true, []byte(""))
	val4, err4 := bs.UnmarshallToValue(false, []byte("Active"))

	actual := args.Map{
		"nilMapped":   val1,
		"nilMapErr":   err1 == nil,
		"nilNoMapErr": err2 != nil,
		"emptyVal":    val3,
		"emptyErr":    err3 == nil,
		"validVal":    val4,
		"validErr":    err4 == nil,
	}
	expected := args.Map{
		"nilMapped":   "Invalid",
		"nilMapErr":   true,
		"nilNoMapErr": true,
		"emptyVal":    "Invalid",
		"emptyErr":    true,
		"validVal":    "Active",
		"validErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_UnmarshallToValue", actual)
}

func Test_Cov3_BasicString_EnumType(t *testing.T) {
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	actual := args.Map{
		"enumType": bs.EnumType().String(),
	}
	expected := args.Map{
		"enumType": "String",
	}
	expected.ShouldBeEqual(t, 0, "BasicString_EnumType", actual)
}

func Test_Cov3_BasicString_AppendPrependJoinValue(t *testing.T) {
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	result := bs.AppendPrependJoinValue(".", "Active", "Invalid")

	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_AppendPrependJoinValue", actual)
}

func Test_Cov3_BasicString_OnlySupportedErr(t *testing.T) {
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	noErr := bs.OnlySupportedErr("Invalid", "Active")
	hasErr := bs.OnlySupportedErr("Invalid")
	msgErr := bs.OnlySupportedMsgErr("context", "Invalid")

	actual := args.Map{
		"noErr":  noErr == nil,
		"hasErr": hasErr != nil,
		"msgErr": msgErr != nil,
	}
	expected := args.Map{
		"noErr":  true,
		"hasErr": true,
		"msgErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_OnlySupportedErr", actual)
}

// ── differCheckerImpl ──

func Test_Cov3_DifferChecker_GetSingleDiffResult(t *testing.T) {
	leftResult := enumimpl.DefaultDiffCheckerImpl.GetSingleDiffResult(true, "L", "R")
	rightResult := enumimpl.DefaultDiffCheckerImpl.GetSingleDiffResult(false, "L", "R")

	actual := args.Map{
		"left":  leftResult,
		"right": rightResult,
	}
	expected := args.Map{
		"left":  "L",
		"right": "R",
	}
	expected.ShouldBeEqual(t, 0, "DifferChecker_GetSingleDiffResult", actual)
}

func Test_Cov3_DifferChecker_GetResultOnKeyMissing(t *testing.T) {
	result := enumimpl.DefaultDiffCheckerImpl.GetResultOnKeyMissingInRightExistInLeft("key", "val")

	actual := args.Map{
		"result": result,
	}
	expected := args.Map{
		"result": "val",
	}
	expected.ShouldBeEqual(t, 0, "DifferChecker_GetResultOnKeyMissing", actual)
}

func Test_Cov3_DifferChecker_IsEqual(t *testing.T) {
	actual := args.Map{
		"regardlessSame": enumimpl.DefaultDiffCheckerImpl.IsEqual(true, 1, 1),
		"regardlessDiff": enumimpl.DefaultDiffCheckerImpl.IsEqual(true, 1, 2),
		"strictSame":     enumimpl.DefaultDiffCheckerImpl.IsEqual(false, "a", "a"),
		"strictDiff":     enumimpl.DefaultDiffCheckerImpl.IsEqual(false, "a", "b"),
	}
	expected := args.Map{
		"regardlessSame": true,
		"regardlessDiff": false,
		"strictSame":     true,
		"strictDiff":     false,
	}
	expected.ShouldBeEqual(t, 0, "DifferChecker_IsEqual", actual)
}

func Test_Cov3_DifferChecker_AsDifferChecker(t *testing.T) {
	checker := enumimpl.DefaultDiffCheckerImpl.AsDifferChecker()

	actual := args.Map{
		"notNil": checker != nil,
	}
	expected := args.Map{
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "DifferChecker_AsDifferChecker", actual)
}

// ── leftRightDiffCheckerImpl ──

func Test_Cov3_LeftRightDiffChecker_GetSingleDiffResult(t *testing.T) {
	result := enumimpl.LeftRightDiffCheckerImpl.GetSingleDiffResult(true, "L", "R")

	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRightDiffChecker_GetSingleDiffResult", actual)
}

func Test_Cov3_LeftRightDiffChecker_GetResultOnKeyMissing(t *testing.T) {
	result := enumimpl.LeftRightDiffCheckerImpl.GetResultOnKeyMissingInRightExistInLeft("key", "val")

	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRightDiffChecker_GetResultOnKeyMissing", actual)
}

func Test_Cov3_LeftRightDiffChecker_IsEqual(t *testing.T) {
	actual := args.Map{
		"same": enumimpl.LeftRightDiffCheckerImpl.IsEqual(false, "a", "a"),
		"diff": enumimpl.LeftRightDiffCheckerImpl.IsEqual(false, "a", "b"),
	}
	expected := args.Map{
		"same": true,
		"diff": false,
	}
	expected.ShouldBeEqual(t, 0, "LeftRightDiffChecker_IsEqual", actual)
}

func Test_Cov3_LeftRightDiffChecker_AsChecker(t *testing.T) {
	checker := enumimpl.LeftRightDiffCheckerImpl.AsChecker()

	actual := args.Map{
		"notNil": checker != nil,
	}
	expected := args.Map{
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRightDiffChecker_AsChecker", actual)
}

// ── FormatUsingFmt ──

func Test_Cov3_FormatUsingFmt(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	result := bb.Format("{type-name}.{name}={value}", byte(1))

	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "FormatUsingFmt", actual)
}

// ── ConvEnumAnyValToInteger additional branches ──

func Test_Cov3_ConvEnumAnyValToInteger_Byte(t *testing.T) {
	type myByte byte
	result := enumimpl.ConvEnumAnyValToInteger(myByte(5))

	actual := args.Map{
		"isPositive": result >= 0,
	}
	expected := args.Map{
		"isPositive": true,
	}
	expected.ShouldBeEqual(t, 0, "ConvEnumAnyValToInteger_Byte", actual)
}

// ── BasicByte via CreateUsingMap ──

func Test_Cov3_BasicByte_CreateUsingMap(t *testing.T) {
	bb := enumimpl.New.BasicByte.CreateUsingMap(
		"testEnum",
		map[byte]string{0: "Invalid", 1: "Active"},
	)

	actual := args.Map{
		"typeName": bb.TypeName(),
		"length":   bb.Length(),
	}
	expected := args.Map{
		"typeName": "testEnum",
		"length":   2,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_CreateUsingMap", actual)
}

// ── BasicByte ExpectingEnumValueError ──

func Test_Cov3_BasicByte_ExpectingEnumValueError(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{"Invalid", "Active"},
	)

	noErr := bb.ExpectingEnumValueError("Active", byte(1))
	hasErr := bb.ExpectingEnumValueError("Invalid", byte(1))
	parseErr := bb.ExpectingEnumValueError("NotExist", byte(1))

	actual := args.Map{
		"matchNoErr":  noErr == nil,
		"mismatchErr": hasErr != nil,
		"parseErr":    parseErr != nil,
	}
	expected := args.Map{
		"matchNoErr":  true,
		"mismatchErr": true,
		"parseErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "BasicByte_ExpectingEnumValueError", actual)
}

// ── NamesHashset empty ──

func Test_Cov3_NumberEnumBase_NamesHashsetEmpty(t *testing.T) {
	type myEnum byte
	bb := enumimpl.New.BasicByte.Default(
		myEnum(0),
		[]string{},
	)

	actual := args.Map{
		"len": len(bb.NamesHashset()),
	}
	expected := args.Map{
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "numberEnumBase_NamesHashsetEmpty", actual)
}
