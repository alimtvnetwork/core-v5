package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── BasicByte ──

func Test_Cov11_BasicByte_Methods(t *testing.T) {
	bb := enumimpl.New.BasicByte.Create(1, "TestByte")
	actual := args.Map{
		"value":    bb.Value(),
		"name":     bb.Name(),
		"isValid":  bb.IsValid(),
		"hasName":  bb.HasName(),
		"str":      bb.String() != "",
		"valInt":   bb.ValueInt(),
		"valInt32": bb.ValueInt32(),
	}
	expected := args.Map{
		"value": byte(1), "name": "TestByte", "isValid": true,
		"hasName": true, "str": true, "valInt": 1, "valInt32": int32(1),
	}
	expected.ShouldBeEqual(t, 0, "BasicByte methods", actual)
}

func Test_Cov11_BasicByte_MaxByte(t *testing.T) {
	bb := enumimpl.New.BasicByte.Create(255, "MaxByte")
	actual := args.Map{"val": bb.Value(), "name": bb.Name()}
	expected := args.Map{"val": byte(255), "name": "MaxByte"}
	expected.ShouldBeEqual(t, 0, "BasicByte max", actual)
}

func Test_Cov11_BasicByte_IsEqual(t *testing.T) {
	a := enumimpl.New.BasicByte.Create(1, "A")
	b := enumimpl.New.BasicByte.Create(1, "A")
	c := enumimpl.New.BasicByte.Create(2, "B")
	actual := args.Map{"equal": a.IsEqual(b), "notEqual": a.IsEqual(c)}
	expected := args.Map{"equal": true, "notEqual": false}
	expected.ShouldBeEqual(t, 0, "BasicByte IsEqual", actual)
}

// ── BasicInt8 ──

func Test_Cov11_BasicInt8_Methods(t *testing.T) {
	bi := enumimpl.New.BasicInt8.Create(5, "TestInt8")
	actual := args.Map{
		"value":    bi.Value(),
		"name":     bi.Name(),
		"str":      bi.String() != "",
		"valInt":   bi.ValueInt(),
		"valInt32": bi.ValueInt32(),
	}
	expected := args.Map{
		"value": int8(5), "name": "TestInt8", "str": true,
		"valInt": 5, "valInt32": int32(5),
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8 methods", actual)
}

// ── BasicInt16 ──

func Test_Cov11_BasicInt16_Methods(t *testing.T) {
	bi := enumimpl.New.BasicInt16.Create(100, "TestInt16")
	actual := args.Map{
		"value":    bi.Value(),
		"name":     bi.Name(),
		"str":      bi.String() != "",
		"valInt":   bi.ValueInt(),
		"valInt32": bi.ValueInt32(),
	}
	expected := args.Map{
		"value": int16(100), "name": "TestInt16", "str": true,
		"valInt": 100, "valInt32": int32(100),
	}
	expected.ShouldBeEqual(t, 0, "BasicInt16 methods", actual)
}

// ── BasicInt32 ──

func Test_Cov11_BasicInt32_Methods(t *testing.T) {
	bi := enumimpl.New.BasicInt32.Create(1000, "TestInt32")
	actual := args.Map{
		"value":    bi.Value(),
		"name":     bi.Name(),
		"str":      bi.String() != "",
		"valInt":   bi.ValueInt(),
		"valInt32": bi.ValueInt32(),
	}
	expected := args.Map{
		"value": int32(1000), "name": "TestInt32", "str": true,
		"valInt": 1000, "valInt32": int32(1000),
	}
	expected.ShouldBeEqual(t, 0, "BasicInt32 methods", actual)
}

// ── BasicUInt16 ──

func Test_Cov11_BasicUInt16_Methods(t *testing.T) {
	bu := enumimpl.New.BasicUInt16.Create(200, "TestUInt16")
	actual := args.Map{
		"value":    bu.Value(),
		"name":     bu.Name(),
		"str":      bu.String() != "",
		"valInt":   bu.ValueInt(),
		"valInt32": bu.ValueInt32(),
	}
	expected := args.Map{
		"value": uint16(200), "name": "TestUInt16", "str": true,
		"valInt": 200, "valInt32": int32(200),
	}
	expected.ShouldBeEqual(t, 0, "BasicUInt16 methods", actual)
}

// ── BasicString ──

func Test_Cov11_BasicString_Methods(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("val", "TestStr")
	actual := args.Map{
		"value": bs.Value(),
		"name":  bs.Name(),
		"str":   bs.String() != "",
	}
	expected := args.Map{
		"value": "val", "name": "TestStr", "str": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicString methods", actual)
}

// ── Format ──

func Test_Cov11_Format_NameValue(t *testing.T) {
	result := enumimpl.Format.NameValue("TestEnum", 5)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Format.NameValue", actual)
}

func Test_Cov11_Format_ValueName(t *testing.T) {
	result := enumimpl.Format.ValueName(5, "TestEnum")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Format.ValueName", actual)
}

// ── ConvAnyValToInteger ──

func Test_Cov11_ConvAnyValToInteger_Int(t *testing.T) {
	result := enumimpl.ConvAnyValToInteger(42)
	actual := args.Map{"val": result}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger int", actual)
}

func Test_Cov11_ConvAnyValToInteger_Int8(t *testing.T) {
	result := enumimpl.ConvAnyValToInteger(int8(5))
	actual := args.Map{"val": result}
	expected := args.Map{"val": 5}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger int8", actual)
}

func Test_Cov11_ConvAnyValToInteger_Int16(t *testing.T) {
	result := enumimpl.ConvAnyValToInteger(int16(100))
	actual := args.Map{"val": result}
	expected := args.Map{"val": 100}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger int16", actual)
}

func Test_Cov11_ConvAnyValToInteger_Int32(t *testing.T) {
	result := enumimpl.ConvAnyValToInteger(int32(1000))
	actual := args.Map{"val": result}
	expected := args.Map{"val": 1000}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger int32", actual)
}

func Test_Cov11_ConvAnyValToInteger_Byte(t *testing.T) {
	result := enumimpl.ConvAnyValToInteger(byte(10))
	actual := args.Map{"val": result}
	expected := args.Map{"val": 10}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger byte", actual)
}

func Test_Cov11_ConvAnyValToInteger_UInt16(t *testing.T) {
	result := enumimpl.ConvAnyValToInteger(uint16(200))
	actual := args.Map{"val": result}
	expected := args.Map{"val": 200}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger uint16", actual)
}

func Test_Cov11_ConvAnyValToInteger_String(t *testing.T) {
	result := enumimpl.ConvAnyValToInteger("not-a-number")
	actual := args.Map{"val": result}
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger string", actual)
}

// ── NameWithValue ──

func Test_Cov11_NameWithValue(t *testing.T) {
	nv := enumimpl.NameWithValue{Name: "test", Value: 42}
	actual := args.Map{"name": nv.Name, "val": nv.Value}
	expected := args.Map{"name": "test", "val": 42}
	expected.ShouldBeEqual(t, 0, "NameWithValue", actual)
}

// ── AllNameValues ──

func Test_Cov11_AllNameValues(t *testing.T) {
	items := []enumimpl.NameWithValue{
		{Name: "a", Value: 1},
		{Name: "b", Value: 2},
	}
	anv := enumimpl.AllNameValues{Items: items}
	actual := args.Map{"len": len(anv.Items)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AllNameValues", actual)
}

// ── JoinPrependUsingDot ──

func Test_Cov11_JoinPrependUsingDot(t *testing.T) {
	result := enumimpl.JoinPrependUsingDot("prefix", "name")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JoinPrependUsingDot", actual)
}

// ── PrependJoin ──

func Test_Cov11_PrependJoin(t *testing.T) {
	result := enumimpl.PrependJoin(".", "prefix", "name")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrependJoin", actual)
}

// ── OnlySupportedErr ──

func Test_Cov11_OnlySupportedErr(t *testing.T) {
	err := enumimpl.OnlySupportedErr("TestEnum", "badVal")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr", actual)
}

// ── UnsupportedNames ──

func Test_Cov11_UnsupportedNames(t *testing.T) {
	result := enumimpl.UnsupportedNames("TestEnum", "bad1", "bad2")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "UnsupportedNames", actual)
}

// ── DynamicMap ──

func Test_Cov11_DynamicMap(t *testing.T) {
	dm := enumimpl.DynamicMap{
		EnumTypeName: "TestEnum",
	}
	actual := args.Map{"name": dm.EnumTypeName}
	expected := args.Map{"name": "TestEnum"}
	expected.ShouldBeEqual(t, 0, "DynamicMap", actual)
}

// ── KeyAnyVal ──

func Test_Cov11_KeyAnyVal(t *testing.T) {
	kv := enumimpl.KeyAnyVal{
		Key:   "k1",
		Value: 42,
	}
	actual := args.Map{"key": kv.Key, "val": kv.Value}
	expected := args.Map{"key": "k1", "val": 42}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal", actual)
}

// ── KeyAnyValues ──

func Test_Cov11_KeyAnyValues(t *testing.T) {
	kvs := enumimpl.KeyAnyValues{
		Items: []enumimpl.KeyAnyVal{
			{Key: "k1", Value: 1},
			{Key: "k2", Value: 2},
		},
	}
	actual := args.Map{"len": len(kvs.Items)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues", actual)
}

// ── FormatUsingFmt ──

func Test_Cov11_FormatUsingFmt(t *testing.T) {
	result := enumimpl.FormatUsingFmt("Enum", 5)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FormatUsingFmt", actual)
}

// ── KeyValInteger ──

func Test_Cov11_KeyValInteger(t *testing.T) {
	kv := enumimpl.KeyValInteger{Key: "k", Value: 42}
	actual := args.Map{"key": kv.Key, "val": kv.Value}
	expected := args.Map{"key": "k", "val": 42}
	expected.ShouldBeEqual(t, 0, "KeyValInteger", actual)
}

// ── DiffLeftRight ──

func Test_Cov11_DiffLeftRight(t *testing.T) {
	d := enumimpl.DiffLeftRight{
		LeftName:  "left",
		LeftVal:   1,
		RightName: "right",
		RightVal:  2,
	}
	actual := args.Map{
		"leftName": d.LeftName, "leftVal": d.LeftVal,
		"rightName": d.RightName, "rightVal": d.RightVal,
	}
	expected := args.Map{
		"leftName": "left", "leftVal": 1,
		"rightName": "right", "rightVal": 2,
	}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight", actual)
}
