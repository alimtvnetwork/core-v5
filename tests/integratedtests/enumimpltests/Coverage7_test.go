package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ConvEnumAnyValToInteger ──

func Test_Cov7_ConvEnumAnyValToInteger_Int(t *testing.T) {
	val := enumimpl.ConvEnumAnyValToInteger(42)
	actual := args.Map{"val": val}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ConvEnumAnyValToInteger returns 42 -- int input", actual)
}

func Test_Cov7_ConvEnumAnyValToInteger_String(t *testing.T) {
	val := enumimpl.ConvEnumAnyValToInteger("notAnInt")
	actual := args.Map{"isMinInt": val < 0}
	expected := args.Map{"isMinInt": true}
	expected.ShouldBeEqual(t, 0, "ConvEnumAnyValToInteger returns MinInt -- string input", actual)
}

// ── NameWithValue (function) ──

func Test_Cov7_NameWithValue(t *testing.T) {
	result := enumimpl.NameWithValue(10)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameWithValue returns formatted string -- int input", actual)
}

// ── UnsupportedNames ──

func Test_Cov7_UnsupportedNames(t *testing.T) {
	allNames := []string{"A", "B", "C", "D"}
	result := enumimpl.UnsupportedNames(allNames, "A", "B")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "UnsupportedNames returns 2 -- two unsupported", actual)
}

func Test_Cov7_UnsupportedNames_AllSupported(t *testing.T) {
	allNames := []string{"A", "B"}
	result := enumimpl.UnsupportedNames(allNames, "A", "B")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "UnsupportedNames returns 0 -- all supported", actual)
}

// ── KeyAnyVal ──

func Test_Cov7_KeyAnyVal(t *testing.T) {
	kv := enumimpl.KeyAnyVal{Key: "test", AnyValue: 42}
	actual := args.Map{
		"key":      kv.Key,
		"valInt":   kv.ValInt(),
		"isString": kv.IsString(),
	}
	expected := args.Map{"key": "test", "valInt": 42, "isString": false}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal returns correct fields -- int value", actual)
}

func Test_Cov7_KeyAnyVal_StringValue(t *testing.T) {
	kv := enumimpl.KeyAnyVal{Key: "strKey", AnyValue: "hello"}
	actual := args.Map{
		"key":      kv.Key,
		"isString": kv.IsString(),
		"anyVal":   kv.AnyValString(),
	}
	expected := args.Map{"key": "strKey", "isString": true, "anyVal": "hello"}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal returns string type -- string value", actual)
}

// ── DiffLeftRight ──

func Test_Cov7_DiffLeftRight_Same(t *testing.T) {
	dlr := &enumimpl.DiffLeftRight{Left: "same", Right: "same"}
	actual := args.Map{
		"isSame":      dlr.IsSame(),
		"isNotEqual":  dlr.IsNotEqual(),
		"isEqual":     dlr.IsEqual(false),
		"diffStr":     dlr.DiffString(),
	}
	expected := args.Map{
		"isSame": true, "isNotEqual": false, "isEqual": true, "diffStr": "",
	}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight returns same -- equal values", actual)
}

func Test_Cov7_DiffLeftRight_Different(t *testing.T) {
	dlr := &enumimpl.DiffLeftRight{Left: "left", Right: "right"}
	actual := args.Map{
		"isSame":     dlr.IsSame(),
		"isNotEqual": dlr.IsNotEqual(),
		"hasMismatch": dlr.HasMismatch(false),
	}
	expected := args.Map{
		"isSame": false, "isNotEqual": true, "hasMismatch": true,
	}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight returns mismatch -- different values", actual)
}

func Test_Cov7_DiffLeftRight_RegardlessOfType(t *testing.T) {
	dlr := &enumimpl.DiffLeftRight{Left: 42, Right: 42}
	actual := args.Map{
		"isEqualRegardless": dlr.IsEqual(true),
		"isSameTypeSame":    dlr.IsSameTypeSame(),
	}
	expected := args.Map{
		"isEqualRegardless": true, "isSameTypeSame": true,
	}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight returns equal -- regardless of type", actual)
}

func Test_Cov7_DiffLeftRight_JsonString(t *testing.T) {
	dlr := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	actual := args.Map{"notEmpty": dlr.JsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight returns json -- serialized", actual)
}

func Test_Cov7_DiffLeftRight_SpecificFullString(t *testing.T) {
	dlr := &enumimpl.DiffLeftRight{Left: "x", Right: "y"}
	l, r := dlr.SpecificFullString()
	actual := args.Map{
		"leftNotEmpty":  l != "",
		"rightNotEmpty": r != "",
	}
	expected := args.Map{"leftNotEmpty": true, "rightNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight returns full strings -- both sides", actual)
}

// ── DefaultDiffCheckerImpl ──

func Test_Cov7_DefaultDiffChecker_IsEqual(t *testing.T) {
	checker := enumimpl.DefaultDiffCheckerImpl
	result := checker.IsEqual(false, 42, 42)
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "DefaultDiffChecker returns true -- equal values", actual)
}

func Test_Cov7_DefaultDiffChecker_IsEqual_Regardless(t *testing.T) {
	checker := enumimpl.DefaultDiffCheckerImpl
	result := checker.IsEqual(true, 42, 42)
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "DefaultDiffChecker returns true -- regardless mode", actual)
}

func Test_Cov7_LeftRightDiffChecker_IsEqual(t *testing.T) {
	checker := enumimpl.LeftRightDiffCheckerImpl
	result := checker.IsEqual(false, "a", "a")
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "LeftRightDiffChecker returns true -- equal strings", actual)
}
