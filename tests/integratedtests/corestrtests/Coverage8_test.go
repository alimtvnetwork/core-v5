package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ValidValue ──

func Test_Cov8_ValidValue_Valid(t *testing.T) {
	vv := corestr.NewValidValue("hello")
	actual := args.Map{
		"value":   vv.Value,
		"isValid": vv.IsValid,
		"isEmpty": vv.IsEmpty(),
	}
	expected := args.Map{
		"value": "hello", "isValid": true, "isEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "ValidValue valid", actual)
}

func Test_Cov8_ValidValue_Empty(t *testing.T) {
	vv := corestr.NewValidValue("")
	actual := args.Map{
		"isValid": vv.IsValid,
		"isEmpty": vv.IsEmpty(),
	}
	expected := args.Map{"isValid": false, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "ValidValue empty", actual)
}

// ── LeftRight ──

func Test_Cov8_LeftRight(t *testing.T) {
	lr := corestr.LeftRight{Left: "l", Right: "r"}
	actual := args.Map{
		"left":  lr.Left,
		"right": lr.Right,
	}
	expected := args.Map{"left": "l", "right": "r"}
	expected.ShouldBeEqual(t, 0, "LeftRight struct", actual)
}

// ── LeftMiddleRight ──

func Test_Cov8_LeftMiddleRight(t *testing.T) {
	lmr := corestr.LeftMiddleRight{Left: "l", Middle: "m", Right: "r"}
	actual := args.Map{
		"left":   lmr.Left,
		"middle": lmr.Middle,
		"right":  lmr.Right,
	}
	expected := args.Map{"left": "l", "middle": "m", "right": "r"}
	expected.ShouldBeEqual(t, 0, "LeftMiddleRight struct", actual)
}

// ── LeftRightFromSplit ──

func Test_Cov8_LeftRightFromSplit(t *testing.T) {
	lr := corestr.NewLeftRightFromSplit("key=value", "=")
	actual := args.Map{
		"left":  lr.Left,
		"right": lr.Right,
	}
	expected := args.Map{"left": "key", "right": "value"}
	expected.ShouldBeEqual(t, 0, "LeftRightFromSplit", actual)
}

func Test_Cov8_LeftRightFromSplit_NoSep(t *testing.T) {
	lr := corestr.NewLeftRightFromSplit("nosep", "=")
	actual := args.Map{"left": lr.Left, "right": lr.Right}
	expected := args.Map{"left": "nosep", "right": ""}
	expected.ShouldBeEqual(t, 0, "LeftRightFromSplit no separator", actual)
}

// ── LeftMiddleRightFromSplit ──

func Test_Cov8_LeftMiddleRightFromSplit(t *testing.T) {
	lmr := corestr.NewLeftMiddleRightFromSplit("a:b:c", ":")
	actual := args.Map{
		"left":   lmr.Left,
		"middle": lmr.Middle,
		"right":  lmr.Right,
	}
	expected := args.Map{"left": "a", "middle": "b", "right": "c"}
	expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplit", actual)
}

// ── ValueStatus ──

func Test_Cov8_ValueStatus(t *testing.T) {
	vs := corestr.ValueStatus{
		Value:   "test",
		IsValid: true,
	}
	actual := args.Map{
		"value":   vs.Value,
		"isValid": vs.IsValid,
	}
	expected := args.Map{"value": "test", "isValid": true}
	expected.ShouldBeEqual(t, 0, "ValueStatus struct", actual)
}

// ── TextWithLineNumber ──

func Test_Cov8_TextWithLineNumber(t *testing.T) {
	tln := corestr.TextWithLineNumber{
		LineNumber: 5,
		Text:       "hello",
	}
	actual := args.Map{
		"lineNo": tln.LineNumber,
		"text":   tln.Text,
	}
	expected := args.Map{"lineNo": 5, "text": "hello"}
	expected.ShouldBeEqual(t, 0, "TextWithLineNumber struct", actual)
}

// ── KeyValuePair ──

func Test_Cov8_KeyValuePair(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	actual := args.Map{
		"key":      kv.Key,
		"value":    kv.Value,
		"notEmpty": kv.String() != "",
	}
	expected := args.Map{"key": "k", "value": "v", "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyValuePair struct", actual)
}

// ── CloneSlice ──

func Test_Cov8_CloneSlice(t *testing.T) {
	original := []string{"a", "b", "c"}
	cloned := corestr.CloneSlice(original)
	actual := args.Map{
		"len":   len(cloned),
		"first": cloned[0],
	}
	expected := args.Map{"len": 3, "first": "a"}
	expected.ShouldBeEqual(t, 0, "CloneSlice", actual)
}

func Test_Cov8_CloneSlice_Nil(t *testing.T) {
	cloned := corestr.CloneSlice(nil)
	actual := args.Map{"isNil": cloned == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "CloneSlice nil", actual)
}

// ── CloneSliceIf ──

func Test_Cov8_CloneSliceIf_True(t *testing.T) {
	original := []string{"a", "b"}
	cloned := corestr.CloneSliceIf(true, original)
	actual := args.Map{"len": len(cloned)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CloneSliceIf true", actual)
}

func Test_Cov8_CloneSliceIf_False(t *testing.T) {
	original := []string{"a", "b"}
	result := corestr.CloneSliceIf(false, original)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CloneSliceIf false", actual)
}

// ── SimpleStringOnce ──

func Test_Cov8_SimpleStringOnce(t *testing.T) {
	so := corestr.NewSimpleStringOnce("hello")
	actual := args.Map{
		"value":   so.Value(),
		"string":  so.String(),
		"isEmpty": so.IsEmpty(),
	}
	expected := args.Map{
		"value": "hello", "string": "hello", "isEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "SimpleStringOnce", actual)
}
