package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Four ──

func Test_Cov4_Four_Basic(t *testing.T) {
	four := &args.Four[string, int, bool, float64]{First: "a", Second: 1, Third: true, Fourth: 3.14, Expect: "yes"}
	actual := args.Map{
		"first":     four.FirstItem(),
		"second":    four.SecondItem(),
		"third":     four.ThirdItem(),
		"fourth":    four.FourthItem(),
		"expected":  four.Expected(),
		"hasFirst":  four.HasFirst(),
		"hasSecond": four.HasSecond(),
		"hasThird":  four.HasThird(),
		"hasFourth": four.HasFourth(),
		"count":     four.ArgsCount(),
	}
	expected := args.Map{
		"first": "a", "second": 1, "third": true, "fourth": 3.14, "expected": "yes",
		"hasFirst": true, "hasSecond": true, "hasThird": true, "hasFourth": true, "count": 4,
	}
	expected.ShouldBeEqual(t, 0, "Four basic -- all types", actual)
}

func Test_Cov4_Four_GetByIndex(t *testing.T) {
	four := &args.Four[string, int, bool, float64]{First: "a", Second: 1, Third: true, Fourth: 3.14}
	actual := args.Map{
		"idx0": four.GetByIndex(0), "idx1": four.GetByIndex(1),
		"idx2": four.GetByIndex(2), "idx3": four.GetByIndex(3),
	}
	expected := args.Map{"idx0": "a", "idx1": 1, "idx2": true, "idx3": 3.14}
	expected.ShouldBeEqual(t, 0, "Four GetByIndex", actual)
}

func Test_Cov4_Four_Slice(t *testing.T) {
	four := &args.Four[string, int, bool, float64]{First: "a", Second: 1, Third: true, Fourth: 3.14}
	actual := args.Map{"len": len(four.Slice())}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "Four Slice", actual)
}

func Test_Cov4_Four_String(t *testing.T) {
	four := &args.Four[string, int, bool, float64]{First: "a", Second: 1, Third: true, Fourth: 3.14}
	actual := args.Map{"hasContent": len(four.String()) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Four String", actual)
}

// ── Five ──

func Test_Cov4_Five_Basic(t *testing.T) {
	five := &args.Five[string, int, bool, float64, byte]{First: "a", Second: 1, Third: true, Fourth: 3.14, Fifth: byte(5), Expect: "yes"}
	actual := args.Map{
		"first":  five.FirstItem(), "second": five.SecondItem(), "third": five.ThirdItem(),
		"fourth": five.FourthItem(), "fifth": five.FifthItem(), "expected": five.Expected(),
		"count":  five.ArgsCount(),
	}
	expected := args.Map{
		"first": "a", "second": 1, "third": true, "fourth": 3.14, "fifth": byte(5), "expected": "yes", "count": 5,
	}
	expected.ShouldBeEqual(t, 0, "Five basic", actual)
}

func Test_Cov4_Five_GetByIndex(t *testing.T) {
	five := &args.Five[string, int, bool, float64, byte]{First: "a", Second: 1, Third: true, Fourth: 3.14, Fifth: byte(5)}
	actual := args.Map{"idx4": five.GetByIndex(4)}
	expected := args.Map{"idx4": byte(5)}
	expected.ShouldBeEqual(t, 0, "Five GetByIndex", actual)
}

// ── Six ──

func Test_Cov4_Six_Basic(t *testing.T) {
	six := &args.Six[string, int, bool, float64, byte, uint]{
		First: "a", Second: 1, Third: true, Fourth: 3.14, Fifth: byte(5), Sixth: uint(6), Expect: "yes",
	}
	actual := args.Map{
		"first": six.FirstItem(), "sixth": six.SixthItem(), "expected": six.Expected(), "count": six.ArgsCount(),
	}
	expected := args.Map{"first": "a", "sixth": uint(6), "expected": "yes", "count": 6}
	expected.ShouldBeEqual(t, 0, "Six basic", actual)
}

func Test_Cov4_Six_GetByIndex(t *testing.T) {
	six := &args.Six[string, int, bool, float64, byte, uint]{
		First: "a", Second: 1, Third: true, Fourth: 3.14, Fifth: byte(5), Sixth: uint(6),
	}
	actual := args.Map{"idx5": six.GetByIndex(5)}
	expected := args.Map{"idx5": uint(6)}
	expected.ShouldBeEqual(t, 0, "Six GetByIndex", actual)
}

// ── Dynamic ──

func Test_Cov4_Dynamic_Basic(t *testing.T) {
	d := &args.Dynamic{First: "hello", Expect: 42}
	actual := args.Map{
		"first":  d.FirstItem(), "expected": d.Expected(),
		"hasFirst": d.HasFirst(), "hasExpect": d.HasExpect(), "count": d.ArgsCount(),
	}
	expected := args.Map{"first": "hello", "expected": 42, "hasFirst": true, "hasExpect": true, "count": 1}
	expected.ShouldBeEqual(t, 0, "Dynamic basic", actual)
}

func Test_Cov4_Dynamic_GetByIndex(t *testing.T) {
	d := &args.Dynamic{First: "hello"}
	actual := args.Map{"idx0": d.GetByIndex(0)}
	expected := args.Map{"idx0": "hello"}
	expected.ShouldBeEqual(t, 0, "Dynamic GetByIndex", actual)
}

// ── FuncMap ──

func Test_Cov4_FuncMap_Basic(t *testing.T) {
	fm := &args.FuncMap{Func: func() {}, Expect: "yes"}
	actual := args.Map{
		"hasFunc":   fm.HasFunc(),
		"hasExpect": fm.HasExpect(),
		"expected":  fm.Expected(),
	}
	expected := args.Map{"hasFunc": true, "hasExpect": true, "expected": "yes"}
	expected.ShouldBeEqual(t, 0, "FuncMap basic", actual)
}

// ── Map.Get missing key ──

func Test_Cov4_Map_Get_MissingKey(t *testing.T) {
	m := args.Map{"a": 1}
	val, ok := m.Get("missing")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": nil, "ok": false}
	expected.ShouldBeEqual(t, 0, "Map Get missing key", actual)
}

func Test_Cov4_Map_GetFirstOfNames_None(t *testing.T) {
	m := args.Map{"a": 1}
	val := m.GetFirstOfNames("x", "y", "z")
	actual := args.Map{"isNil": val == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Map GetFirstOfNames none found", actual)
}

// ── Map compile ──

func Test_Cov4_MapCompile_ToGoLiteral(t *testing.T) {
	m := args.Map{"key": "value"}
	literal := m.ToGoLiteral()
	actual := args.Map{"notEmpty": len(literal) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Map ToGoLiteral", actual)
}

// ── Holder nil func ──

func Test_Cov4_Holder_NilFunc(t *testing.T) {
	h := &args.Holder[func() string]{First: "hello"}
	actual := args.Map{"hasFunc": h.HasFunc(), "funcName": h.GetFuncName()}
	expected := args.Map{"hasFunc": false, "funcName": ""}
	expected.ShouldBeEqual(t, 0, "Holder nil func", actual)
}

// ── Empty creator ──

func Test_Cov4_Empty_One(t *testing.T) {
	o := args.EmptyCreator.One()
	actual := args.Map{"notNil": o != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Empty One", actual)
}

func Test_Cov4_Empty_Two(t *testing.T) {
	tw := args.EmptyCreator.Two()
	actual := args.Map{"notNil": tw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Empty Two", actual)
}

func Test_Cov4_Empty_Map(t *testing.T) {
	m := args.EmptyCreator.Map()
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty Map", actual)
}

// ── OneFunc / TwoFunc / ThreeFunc ──

func Test_Cov4_OneFunc(t *testing.T) {
	of := &args.OneFunc[string, func() string]{First: "a", WorkFunc: func() string { return "hello" }}
	actual := args.Map{"first": of.FirstItem(), "hasFunc": of.HasFunc()}
	expected := args.Map{"first": "a", "hasFunc": true}
	expected.ShouldBeEqual(t, 0, "OneFunc", actual)
}

func Test_Cov4_TwoFunc(t *testing.T) {
	tf := &args.TwoFunc[string, int, func() string]{First: "a", Second: 1, WorkFunc: func() string { return "hello" }}
	actual := args.Map{"first": tf.FirstItem(), "second": tf.SecondItem(), "hasFunc": tf.HasFunc()}
	expected := args.Map{"first": "a", "second": 1, "hasFunc": true}
	expected.ShouldBeEqual(t, 0, "TwoFunc", actual)
}

func Test_Cov4_ThreeFunc(t *testing.T) {
	tf := &args.ThreeFunc[string, int, bool, func() string]{
		First: "a", Second: 1, Third: true, WorkFunc: func() string { return "hello" },
	}
	actual := args.Map{"first": tf.FirstItem(), "hasFunc": tf.HasFunc()}
	expected := args.Map{"first": "a", "hasFunc": true}
	expected.ShouldBeEqual(t, 0, "ThreeFunc", actual)
}

// ── FuncWrap with args ──

func Test_Cov4_FuncWrap_WithArgs(t *testing.T) {
	fn := func(a, b int) int { return a + b }
	fw := args.NewFuncWrap.Default(fn)
	actual := args.Map{
		"inCount":  fw.InArgsCount(),
		"outCount": fw.OutArgsCount(),
		"isValid":  fw.IsValid(),
		"isIntFunc": fw.IsIntFunc(),
	}
	expected := args.Map{"inCount": 2, "outCount": 1, "isValid": true, "isIntFunc": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap with args", actual)
}

func Test_Cov4_FuncWrap_Nil(t *testing.T) {
	fw := args.NewFuncWrap.Default(nil)
	actual := args.Map{"isValid": fw.IsValid()}
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "FuncWrap nil", actual)
}
