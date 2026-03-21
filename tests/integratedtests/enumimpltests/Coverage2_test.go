package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ConvEnumAnyValToInteger ──

func Test_Cov2_ConvEnumAnyValToInteger_String(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger("hello")
	actual := args.Map{"isMinInt": result < 0}
	expected := args.Map{"isMinInt": true}
	expected.ShouldBeEqual(t, 0, "ConvToInt_String returns correct value -- with args", actual)
}

func Test_Cov2_ConvEnumAnyValToInteger_Int(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger(42)
	actual := args.Map{"val": result}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ConvToInt_Int returns correct value -- with args", actual)
}

func Test_Cov2_ConvEnumAnyValToInteger_Fallback(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger(3.14)
	actual := args.Map{"isInt": result != 0}
	expected := args.Map{"isInt": true}
	expected.ShouldBeEqual(t, 0, "ConvToInt_Fallback returns correct value -- with args", actual)
}

// ── PrependJoin / JoinPrependUsingDot ──

func Test_Cov2_PrependJoin(t *testing.T) {
	result := enumimpl.PrependJoin(".", "prefix", "a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrependJoin returns correct value -- with args", actual)
}

func Test_Cov2_JoinPrependUsingDot(t *testing.T) {
	result := enumimpl.JoinPrependUsingDot("prefix", "a")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JoinPrependUsingDot returns correct value -- with args", actual)
}

// ── KeyAnyVal ──

func Test_Cov2_KeyAnyVal_Methods(t *testing.T) {
	kv := enumimpl.KeyAnyVal{Key: "name", AnyValue: 42}
	kvStr := enumimpl.KeyAnyVal{Key: "name", AnyValue: "hello"}

	actual := args.Map{
		"key":         kv.KeyString(),
		"anyVal":      kv.AnyVal() != nil,
		"anyValStr":   kv.AnyValString() != "",
		"wrapKey":     kv.WrapKey() != "",
		"wrapVal":     kv.WrapValue() != "",
		"isString":    kv.IsString(),
		"valInt":      kv.ValInt(),
		"string":      kv.String() != "",
		"strIsString": kvStr.IsString(),
		"strString":   kvStr.String() != "",
	}
	expected := args.Map{
		"key":         "name",
		"anyVal":      true,
		"anyValStr":   true,
		"wrapKey":     true,
		"wrapVal":     true,
		"isString":    false,
		"valInt":      42,
		"string":      true,
		"strIsString": true,
		"strString":   true,
	}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal returns correct value -- with args", actual)
}

func Test_Cov2_KeyAnyVal_KeyValInteger(t *testing.T) {
	kv := enumimpl.KeyAnyVal{Key: "test", AnyValue: 5}
	kvi := kv.KeyValInteger()
	actual := args.Map{"key": kvi.Key, "val": kvi.ValueInteger}
	expected := args.Map{"key": "test", "val": 5}
	expected.ShouldBeEqual(t, 0, "KeyValInteger_conv returns correct value -- with args", actual)
}

// ── KeyValInteger ──

func Test_Cov2_KeyValInteger_Methods(t *testing.T) {
	kvi := enumimpl.KeyValInteger{Key: "test", ValueInteger: 5}
	kviStr := enumimpl.KeyValInteger{Key: "test", ValueInteger: -9223372036854775808}

	actual := args.Map{
		"wrapKey":   kvi.WrapKey() != "",
		"wrapVal":   kvi.WrapValue() != "",
		"isString":  kvi.IsString(),
		"string":    kvi.String() != "",
		"anyKey":    kvi.KeyAnyVal().Key,
		"strIsStr":  kviStr.IsString(),
		"strString": kviStr.String() != "",
	}
	expected := args.Map{
		"wrapKey":   true,
		"wrapVal":   true,
		"isString":  false,
		"string":    true,
		"anyKey":    "test",
		"strIsStr":  true,
		"strString": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValInteger returns correct value -- with args", actual)
}

// ── AllNameValues ──

func Test_Cov2_AllNameValues(t *testing.T) {
	names := []string{"Invalid", "Active"}
	values := []int{0, 1}
	result := enumimpl.AllNameValues(names, values)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AllNameValues returns non-empty -- with args", actual)
}

// ── KeyAnyValues ──

func Test_Cov2_KeyAnyValues(t *testing.T) {
	result := enumimpl.KeyAnyValues([]string{"a", "b"}, []int{1, 2})
	empty := enumimpl.KeyAnyValues([]string{}, []int{})
	actual := args.Map{"len": len(result), "emptyLen": len(empty)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues returns non-empty -- with args", actual)
}

// ── IntegersRangesOfAnyVal ──

func Test_Cov2_IntegersRangesOfAnyVal(t *testing.T) {
	result := enumimpl.IntegersRangesOfAnyVal([]int{3, 1, 2})
	actual := args.Map{"len": len(result), "first": result[0], "last": result[2]}
	expected := args.Map{"len": 3, "first": 1, "last": 3}
	expected.ShouldBeEqual(t, 0, "IntegersRangesOfAnyVal returns correct value -- with args", actual)
}

// ── DynamicMap extra coverage ──

func Test_Cov2_DynamicMap_IsValueString(t *testing.T) {
	dmStr := &enumimpl.DynamicMap{"a": "hello"}
	dmInt := &enumimpl.DynamicMap{"a": 1}
	actual := args.Map{
		"strIsStr": dmStr.IsValueString(),
		"intIsStr": dmInt.IsValueString(),
	}
	expected := args.Map{
		"strIsStr": true,
		"intIsStr": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_IsValueString returns non-empty -- with args", actual)
}

func Test_Cov2_DynamicMap_SortedKeyValues(t *testing.T) {
	dm := &enumimpl.DynamicMap{"b": 2, "a": 1}
	result := dm.SortedKeyValues()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SortedKeyValues returns non-empty -- with args", actual)
}

func Test_Cov2_DynamicMap_SortedKeyAnyValues(t *testing.T) {
	dm := &enumimpl.DynamicMap{"b": "y", "a": "x"}
	result := dm.SortedKeyAnyValues()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SortedKeyAnyValues returns non-empty -- with args", actual)
}

func Test_Cov2_DynamicMap_JsonString_NonNil(t *testing.T) {
	dm := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	result := dm.JsonString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight_JsonString returns correct value -- with args", actual)
}
