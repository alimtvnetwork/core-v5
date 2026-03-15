package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Format / FormatUsingFmt ──

func Test_Cov7_Format_ToString(t *testing.T) {
	result := enumimpl.Format.ToString(42, "TestEnum")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Format ToString", actual)
}

func Test_Cov7_FormatUsingFmt_NameValue(t *testing.T) {
	result := enumimpl.FormatUsingFmt.NameValue("MyEnum", 5)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FormatUsingFmt NameValue", actual)
}

// ── PrependJoin ──

func Test_Cov7_PrependJoin_DotJoin(t *testing.T) {
	result := enumimpl.PrependJoin.DotJoin("prefix", "suffix")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "prefix.suffix"}
	expected.ShouldBeEqual(t, 0, "PrependJoin DotJoin", actual)
}

// ── JoinPrependUsingDot ──

func Test_Cov7_JoinPrependUsingDot(t *testing.T) {
	result := enumimpl.JoinPrependUsingDot("base", "ext")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "base.ext"}
	expected.ShouldBeEqual(t, 0, "JoinPrependUsingDot", actual)
}

// ── ConvAnyValToInteger ──

func Test_Cov7_ConvAnyValToInteger_Int(t *testing.T) {
	val, ok := enumimpl.ConvAnyValToInteger(42)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 42, "ok": true}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger int", actual)
}

func Test_Cov7_ConvAnyValToInteger_Int8(t *testing.T) {
	val, ok := enumimpl.ConvAnyValToInteger(int8(5))
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 5, "ok": true}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger int8", actual)
}

func Test_Cov7_ConvAnyValToInteger_String(t *testing.T) {
	_, ok := enumimpl.ConvAnyValToInteger("notAnInt")
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "ConvAnyValToInteger string", actual)
}

// ── NameWithValue ──

func Test_Cov7_NameWithValue(t *testing.T) {
	nv := enumimpl.NameWithValue{Name: "TestEnum", Value: 10}
	actual := args.Map{
		"name":  nv.Name,
		"value": nv.Value,
	}
	expected := args.Map{"name": "TestEnum", "value": 10}
	expected.ShouldBeEqual(t, 0, "NameWithValue struct", actual)
}

// ── AllNameValues ──

func Test_Cov7_AllNameValues_FromMap(t *testing.T) {
	m := map[string]int{"A": 1, "B": 2}
	result := enumimpl.AllNameValues(m)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AllNameValues from map", actual)
}

// ── UnsupportedNames ──

func Test_Cov7_UnsupportedNames(t *testing.T) {
	supported := []string{"A", "B"}
	all := []string{"A", "B", "C", "D"}
	result := enumimpl.UnsupportedNames(supported, all)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "UnsupportedNames", actual)
}

// ── KeyAnyVal ──

func Test_Cov7_KeyAnyVal(t *testing.T) {
	kv := enumimpl.KeyAnyVal{Key: "test", Value: 42}
	actual := args.Map{"key": kv.Key, "val": kv.Value}
	expected := args.Map{"key": "test", "val": 42}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal struct", actual)
}

// ── KeyAnyValues ──

func Test_Cov7_KeyAnyValues_Length(t *testing.T) {
	kvs := enumimpl.KeyAnyValues{
		{Key: "a", Value: 1},
		{Key: "b", Value: 2},
	}
	actual := args.Map{"len": len(kvs)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues length", actual)
}

// ── DiffLeftRight ──

func Test_Cov7_DiffLeftRight(t *testing.T) {
	dlr := enumimpl.DiffLeftRight{
		Left:  "leftVal",
		Right: "rightVal",
	}
	actual := args.Map{"left": dlr.Left, "right": dlr.Right}
	expected := args.Map{"left": "leftVal", "right": "rightVal"}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight struct", actual)
}
