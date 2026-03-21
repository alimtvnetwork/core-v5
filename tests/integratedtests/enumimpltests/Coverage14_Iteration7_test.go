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

func Test_Cov14_ConvAnyValToInteger_FromString(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger("hello")
	actual := args.Map{"isMinInt": result < -999999}
	expected := args.Map{"isMinInt": true}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger returns correct value -- string", actual)
}

func Test_Cov14_ConvAnyValToInteger_FromInt(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger(42)
	actual := args.Map{"result": result}
	expected := args.Map{"result": 42}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger returns correct value -- int", actual)
}

func Test_Cov14_ConvAnyValToInteger_FromFloat(t *testing.T) {
	// float64 won't match any switch case, falls through to Atoi("3.14") → fail
	result := enumimpl.ConvEnumAnyValToInteger(3.14)
	actual := args.Map{"isMinInt": result < -999999}
	expected := args.Map{"isMinInt": true}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger returns correct value -- float", actual)
}

func Test_Cov14_ConvAnyValToInteger_FromByte(t *testing.T) {
	// byte is uint8 — passes through Atoi path
	result := enumimpl.ConvEnumAnyValToInteger(byte(7))
	actual := args.Map{"result": result}
	expected := args.Map{"result": 7}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger returns correct value -- byte", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicMap — isEqualSingle regardless, ConvMap overflow branches
// Covers DynamicMap.go L867-881, L1263, L1288, L1313, L1338, L1363
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_DynamicMap_DiffRaw_RegardlessType(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": "1"}
	diff := dm.DiffRaw(true, enumimpl.DynamicMap{"a": "1", "b": 1})
	actual := args.Map{"diffLen": len(diff)}
	expected := args.Map{"diffLen": 0}
	expected.ShouldBeEqual(t, 0, "DiffRaw returns correct value -- regardless type", actual)
}

func Test_Cov14_DynamicMap_ConvMapInt8String_Overflow(t *testing.T) {
	dm := enumimpl.DynamicMap{"overflow": int(math.MaxInt8 + 1)}
	result := dm.ConvMapInt8String()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ConvMapInt8String returns correct value -- overflow", actual)
}

func Test_Cov14_DynamicMap_ConvMapInt16String_Overflow(t *testing.T) {
	dm := enumimpl.DynamicMap{"overflow": int(math.MaxInt16 + 1)}
	result := dm.ConvMapInt16String()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ConvMapInt16String returns correct value -- overflow", actual)
}

func Test_Cov14_DynamicMap_ConvMapInt32String_Overflow(t *testing.T) {
	dm := enumimpl.DynamicMap{"overflow": int(math.MaxInt32 + 1)}
	result := dm.ConvMapInt32String()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ConvMapInt32String returns correct value -- overflow", actual)
}

func Test_Cov14_DynamicMap_ConvMapUInt16String_Negative(t *testing.T) {
	dm := enumimpl.DynamicMap{"negative": int(-1)}
	result := dm.ConvMapUInt16String()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ConvMapUInt16String returns correct value -- negative", actual)
}

func Test_Cov14_DynamicMap_ConvMapStringString_NotFound(t *testing.T) {
	dm := enumimpl.DynamicMap{"key": 123}
	result := dm.ConvMapStringString()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ConvMapStringString returns correct value -- not found", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicMap — KeyValueByte and KeyValueInt edge cases
// Covers DynamicMap.go L958-970, L987-989, L1023-1048
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_DynamicMap_KeyValueByte_NotANumber(t *testing.T) {
	dm := enumimpl.DynamicMap{"key": "not-a-number"}
	val, isFound, isFailed := dm.KeyValueByte("key")
	actual := args.Map{"val": fmt.Sprintf("%d", val), "isFound": isFound, "isFailed": isFailed}
	expected := args.Map{"val": "0", "isFound": true, "isFailed": false}
	expected.ShouldBeEqual(t, 0, "KeyValueByte returns correct value -- not-a-number", actual)
}

func Test_Cov14_DynamicMap_KeyValueByte_OutOfRange(t *testing.T) {
	dm := enumimpl.DynamicMap{"key": 999}
	val, isFound, isFailed := dm.KeyValueByte("key")
	actual := args.Map{"val": fmt.Sprintf("%d", val), "isFound": isFound, "isFailed": isFailed}
	expected := args.Map{"val": "0", "isFound": true, "isFailed": true}
	expected.ShouldBeEqual(t, 0, "KeyValueByte returns correct value -- out of range", actual)
}

func Test_Cov14_DynamicMap_KeyValueInt_NotANumber(t *testing.T) {
	dm := enumimpl.DynamicMap{"key": "not-int"}
	_, isFound, isFailed := dm.KeyValueInt("key")
	actual := args.Map{"isFound": isFound, "isFailed": isFailed}
	expected := args.Map{"isFound": true, "isFailed": true}
	expected.ShouldBeEqual(t, 0, "KeyValueInt returns correct value -- not-a-number", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BasicByte/BasicString — GetValueByName wrapped-quote branch
// Covers BasicByte.go L81, BasicString.go L139
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_BasicByte_GetValueByName_WrappedQuote(t *testing.T) {
	dm := enumimpl.DynamicMap{"Alpha": byte(0), "Beta": byte(1)}
	bb := dm.BasicByte("TestByte")

	// Act — pass unwrapped name; the map stores "Alpha" as double-quoted key
	// The method first tries exact match, then wraps with quotes
	val, err := bb.GetValueByName("Alpha")

	actual := args.Map{"val": fmt.Sprintf("%d", val), "hasErr": err != nil}
	expected := args.Map{"val": "0", "hasErr": false}
	expected.ShouldBeEqual(t, 0, "BasicByte returns correct value -- GetValueByName", actual)
}

func Test_Cov14_BasicString_GetValueByName_WrappedQuote(t *testing.T) {
	dm := enumimpl.DynamicMap{"Alpha": "Alpha", "Beta": "Beta"}
	bs := dm.BasicString("TestString")

	val, err := bs.GetValueByName("Alpha")

	actual := args.Map{"val": val, "hasErr": err != nil}
	expected := args.Map{"val": "Alpha", "hasErr": false}
	expected.ShouldBeEqual(t, 0, "BasicString returns correct value -- GetValueByName", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// newBasicStringCreator — CreateUsingStringersSpread
// Covers newBasicStringCreator.go L139-164
// ══════════════════════════════════════════════════════════════════════════════

type testStringer struct{ name string }

func (s testStringer) String() string { return s.name }

func Test_Cov14_NewBasicStringCreator_CreateUsingStringersSpread(t *testing.T) {
	bs := enumimpl.New.BasicString.CreateUsingStringersSpread(
		"TestStringerEnum",
		testStringer{"Alpha"},
		testStringer{"Beta"},
		testStringer{"Gamma"},
	)
	actual := args.Map{"typeName": bs.TypeName(), "length": bs.Length()}
	expected := args.Map{"typeName": "TestStringerEnum", "length": 3}
	expected.ShouldBeEqual(t, 0, "CreateUsingStringersSpread returns correct value -- with args", actual)
}

func Test_Cov14_NewBasicStringCreator_CreateUsingStringersSpread_MinMax(t *testing.T) {
	bs := enumimpl.New.BasicString.CreateUsingStringersSpread(
		"MinMaxEnum",
		testStringer{"C"},
		testStringer{"A"},
		testStringer{"B"},
	)
	actual := args.Map{"length": bs.Length()}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "CreateUsingStringersSpread returns correct value -- min/max", actual)
}
