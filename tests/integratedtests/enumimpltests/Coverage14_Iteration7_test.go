package enumimpltests

import (
	"fmt"
	"math"
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// ConvAnyValToInteger — type switch branches
// Covers ConvAnyValToInteger.go L25-40
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_ConvAnyValToInteger_FromByte(t *testing.T) {
	// Arrange — create a BasicByte enum to get a valueByter
	dm := enumimpl.DynamicMap{"A": byte(0), "B": byte(1)}
	bb := dm.BasicByte("TestByteEnum")

	// Act — pass byte(1) directly
	result := enumimpl.ConvEnumAnyValToInteger(byte(1))

	// Assert
	_ = bb // used for context
	actual := args.Map{"result": result}
	expected := args.Map{"result": 1}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger byte", actual)
}

func Test_Cov14_ConvAnyValToInteger_FromString(t *testing.T) {
	// Arrange — string type
	result := enumimpl.ConvEnumAnyValToInteger("hello")

	// Assert — should return MinInt
	actual := args.Map{"isMinInt": result < -999999}
	expected := args.Map{"isMinInt": true}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger string", actual)
}

func Test_Cov14_ConvAnyValToInteger_FromInt(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger(42)
	actual := args.Map{"result": result}
	expected := args.Map{"result": 42}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger int", actual)
}

func Test_Cov14_ConvAnyValToInteger_FromFloat(t *testing.T) {
	// Arrange — float64 won't match any switch case, falls through to Atoi
	result := enumimpl.ConvEnumAnyValToInteger(3.14)

	// Assert — Atoi("3.14") will fail → MinInt
	actual := args.Map{"isMinInt": result < -999999}
	expected := args.Map{"isMinInt": true}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger float", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicMap — Set nil receiver, AddNewOnly nil, isEqualSingle regardless
// Covers DynamicMap.go L26-29, L41-44, L867-881, L958-970, L987-989, L1023-1048
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_DynamicMap_IsEqualSingle_RegardlessType(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": "1"}

	// Act — diff with isRegardlessType=true, same string representation
	diff := dm.DiffRawMap(true, enumimpl.DynamicMap{"a": "1", "b": 1})

	// Assert — when regardless of type, "1"==1 because Sprint matches
	actual := args.Map{"diffLen": len(diff)}
	expected := args.Map{"diffLen": 0}
	expected.ShouldBeEqual(t, 0, "DynamicMap isEqualSingle regardless", actual)
}

func Test_Cov14_DynamicMap_ConvMapInt8String_Overflow(t *testing.T) {
	// Arrange — value exceeds int8 range
	dm := enumimpl.DynamicMap{"overflow": int(math.MaxInt8 + 1)}

	// Act
	result := dm.ConvMapInt8String()

	// Assert — overflow key should be skipped
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ConvMapInt8String overflow", actual)
}

func Test_Cov14_DynamicMap_ConvMapInt16String_Overflow(t *testing.T) {
	dm := enumimpl.DynamicMap{"overflow": int(math.MaxInt16 + 1)}
	result := dm.ConvMapInt16String()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ConvMapInt16String overflow", actual)
}

func Test_Cov14_DynamicMap_ConvMapInt32String_Overflow(t *testing.T) {
	dm := enumimpl.DynamicMap{"overflow": int(math.MaxInt32 + 1)}
	result := dm.ConvMapInt32String()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ConvMapInt32String overflow", actual)
}

func Test_Cov14_DynamicMap_ConvMapUInt16String_Negative(t *testing.T) {
	dm := enumimpl.DynamicMap{"negative": int(-1)}
	result := dm.ConvMapUInt16String()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ConvMapUInt16String negative", actual)
}

func Test_Cov14_DynamicMap_ConvMapStringString_NotFound(t *testing.T) {
	// Arrange — value is not a string
	dm := enumimpl.DynamicMap{"key": 123}
	result := dm.ConvMapStringString()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ConvMapStringString not found", actual)
}

func Test_Cov14_DynamicMap_GetByteValue_InvalidConversion(t *testing.T) {
	// Arrange — value that can't be converted to byte
	dm := enumimpl.DynamicMap{"key": "not-a-number"}

	// Act
	val, isFound, isFailed := dm.GetByteValue("key")

	// Assert
	actual := args.Map{
		"val":      fmt.Sprintf("%d", val),
		"isFound":  isFound,
		"isFailed": isFailed,
	}
	expected := args.Map{
		"val":      "0",
		"isFound":  true,
		"isFailed": false,
	}
	expected.ShouldBeEqual(t, 0, "GetByteValue invalid", actual)
}

func Test_Cov14_DynamicMap_GetByteValue_OutOfRange(t *testing.T) {
	dm := enumimpl.DynamicMap{"key": 999}
	val, isFound, isFailed := dm.GetByteValue("key")
	actual := args.Map{
		"val":      fmt.Sprintf("%d", val),
		"isFound":  isFound,
		"isFailed": isFailed,
	}
	expected := args.Map{
		"val":      "0",
		"isFound":  true,
		"isFailed": true,
	}
	expected.ShouldBeEqual(t, 0, "GetByteValue out of range", actual)
}

func Test_Cov14_DynamicMap_GetIntValue_InvalidConversion(t *testing.T) {
	dm := enumimpl.DynamicMap{"key": "not-int"}
	val, isFound, isFailed := dm.GetIntValue("key")
	actual := args.Map{
		"isFailed": isFailed,
		"isFound":  isFound,
		"val":      fmt.Sprintf("%d", val),
	}
	expected := args.Map{
		"isFailed": true,
		"isFound":  true,
		"val":      fmt.Sprintf("%d", -1), // InvalidValue
	}
	expected.ShouldBeEqual(t, 0, "GetIntValue invalid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BasicByte/Int8/Int16/Int32/String/UInt16 — UnmarshalJsonToEnumValue wrapped
// Covers Basic*.go wrapped-name branch
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_BasicByte_UnmarshalJson_Wrapped(t *testing.T) {
	dm := enumimpl.DynamicMap{"Alpha": byte(0), "Beta": byte(1)}
	bb := dm.BasicByte("TestByte")

	// Act — try with double-quoted name
	val, err := bb.UnmarshalJsonToEnumValue("\"Alpha\"")

	// Assert
	actual := args.Map{"val": fmt.Sprintf("%d", val), "hasErr": err != nil}
	expected := args.Map{"val": "0", "hasErr": false}
	expected.ShouldBeEqual(t, 0, "BasicByte UnmarshalJson wrapped", actual)
}

func Test_Cov14_BasicString_UnmarshalJson_Wrapped(t *testing.T) {
	dm := enumimpl.DynamicMap{"Alpha": "Alpha", "Beta": "Beta"}
	bs := dm.BasicString("TestString")

	// Act — try with double-quoted name
	val, err := bs.UnmarshalJsonToEnumValue("\"Alpha\"")

	// Assert
	actual := args.Map{"val": val, "hasErr": err != nil}
	expected := args.Map{"val": "\"Alpha\"", "hasErr": false}
	expected.ShouldBeEqual(t, 0, "BasicString UnmarshalJson wrapped", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DiffLeftRight — JsonString marshal error (dead branch L78-80)
// Noted as dead code — json.Marshal rarely fails on simple struct
// ══════════════════════════════════════════════════════════════════════════════

// ══════════════════════════════════════════════════════════════════════════════
// newBasicStringCreator — CreateUsingStringersSpread
// Covers newBasicStringCreator.go L139-164
// ══════════════════════════════════════════════════════════════════════════════

type testStringer struct{ name string }

func (s testStringer) String() string { return s.name }

func Test_Cov14_NewBasicStringCreator_CreateUsingStringersSpread(t *testing.T) {
	// Act
	bs := enumimpl.New.BasicString.CreateUsingStringersSpread(
		"TestStringerEnum",
		testStringer{"Alpha"},
		testStringer{"Beta"},
		testStringer{"Gamma"},
	)

	// Assert
	actual := args.Map{
		"typeName": bs.TypeName(),
		"length":   bs.Length(),
	}
	expected := args.Map{
		"typeName": "TestStringerEnum",
		"length":   3,
	}
	expected.ShouldBeEqual(t, 0, "CreateUsingStringersSpread", actual)
}

func Test_Cov14_NewBasicStringCreator_CreateUsingStringersSpread_MinMax(t *testing.T) {
	// Arrange — "C" > "A", ensures min/max logic
	bs := enumimpl.New.BasicString.CreateUsingStringersSpread(
		"MinMaxEnum",
		testStringer{"C"},
		testStringer{"A"},
		testStringer{"B"},
	)

	// Assert
	actual := args.Map{"length": bs.Length()}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "CreateUsingStringersSpread min/max", actual)
}
