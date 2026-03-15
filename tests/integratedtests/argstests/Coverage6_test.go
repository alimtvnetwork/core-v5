package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// FuncWrapAny — basic
// ═══════════════════════════════════════════

func Test_Cov6_FuncWrapAny_NilFunc(t *testing.T) {
	fw := args.New.FuncWrap.AnyFunc("nilTest", nil)
	actual := args.Map{
		"name":    fw.Name,
		"isNil":   fw.IsNilFunc(),
		"hasFunc": fw.HasFunc(),
	}
	expected := args.Map{"name": "nilTest", "isNil": true, "hasFunc": false}
	expected.ShouldBeEqual(t, 0, "FuncWrapAny nil func", actual)
}

func Test_Cov6_FuncWrapAny_ValidFunc(t *testing.T) {
	fn := func(a, b int) int { return a + b }
	fw := args.New.FuncWrap.AnyFunc("add", fn)
	actual := args.Map{
		"name":    fw.Name,
		"isNil":   fw.IsNilFunc(),
		"hasFunc": fw.HasFunc(),
	}
	expected := args.Map{"name": "add", "isNil": false, "hasFunc": true}
	expected.ShouldBeEqual(t, 0, "FuncWrapAny valid func", actual)
}

func Test_Cov6_FuncWrapAny_Invoke(t *testing.T) {
	fn := func(a, b int) int { return a + b }
	fw := args.New.FuncWrap.AnyFunc("add", fn)
	results, err := fw.Invoke(3, 4)
	actual := args.Map{
		"errNil":    err == nil,
		"resultLen": len(results),
	}
	expected := args.Map{"errNil": true, "resultLen": 1}
	expected.ShouldBeEqual(t, 0, "FuncWrapAny invoke", actual)
}

func Test_Cov6_FuncWrapAny_InvokeNilFunc(t *testing.T) {
	fw := args.New.FuncWrap.AnyFunc("nilTest", nil)
	_, err := fw.Invoke(1)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FuncWrapAny invoke nil func", actual)
}

// ═══════════════════════════════════════════
// Map — accessors
// ═══════════════════════════════════════════

func Test_Cov6_Map_GetAs(t *testing.T) {
	m := args.Map{
		"str": "hello", "int": 42, "bool": true,
		"strs": []string{"a", "b"},
	}
	strVal, strOK := m.GetAsString("str")
	intVal, intOK := m.GetAsInt("int")
	boolVal, boolOK := m.GetAsBool("bool")
	strsVal, strsOK := m.GetAsStrings("strs")
	_, missingOK := m.GetAsString("missing")
	actual := args.Map{
		"str":       strVal,
		"strOK":     strOK,
		"int":       intVal,
		"intOK":     intOK,
		"bool":      boolVal,
		"boolOK":    boolOK,
		"strsLen":   len(strsVal),
		"strsOK":    strsOK,
		"missingOK": missingOK,
	}
	expected := args.Map{
		"str": "hello", "strOK": true,
		"int": 42, "intOK": true,
		"bool": true, "boolOK": true,
		"strsLen": 2, "strsOK": true,
		"missingOK": false,
	}
	expected.ShouldBeEqual(t, 0, "Map GetAs", actual)
}

func Test_Cov6_Map_GetAsBoolDefault(t *testing.T) {
	m := args.Map{"bool": true}
	val := m.GetAsBoolDefault("bool")
	missing := m.GetAsBoolDefault("missing")
	actual := args.Map{"val": val, "missing": missing}
	expected := args.Map{"val": true, "missing": false}
	expected.ShouldBeEqual(t, 0, "Map GetAsBoolDefault", actual)
}

func Test_Cov6_Map_ArgsCount(t *testing.T) {
	m := args.Map{"a": 1, "b": 2, "func": nil, "expect": nil}
	actual := args.Map{"count": m.ArgsCount()}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "Map ArgsCount", actual)
}

func Test_Cov6_Map_WorkFunc(t *testing.T) {
	fn := func() string { return "hello" }
	m := args.Map{"func": fn}
	wf := m.WorkFunc()
	actual := args.Map{"hasFunc": wf != nil}
	expected := args.Map{"hasFunc": true}
	expected.ShouldBeEqual(t, 0, "Map WorkFunc", actual)
}

func Test_Cov6_Map_FirstItem(t *testing.T) {
	m := args.Map{"input": "hello", "when": "hello2"}
	first := m.FirstItem("input", "when")
	missing := m.FirstItem("missing1", "missing2")
	actual := args.Map{
		"first":   first,
		"missing": missing == nil,
	}
	expected := args.Map{"first": "hello", "missing": true}
	expected.ShouldBeEqual(t, 0, "Map FirstItem", actual)
}

func Test_Cov6_Map_HasFunc(t *testing.T) {
	fn := func() {}
	m1 := args.Map{"func": fn}
	m2 := args.Map{"other": 1}
	actual := args.Map{
		"hasFunc":  m1.HasFunc(),
		"noFunc":   m2.HasFunc(),
	}
	expected := args.Map{"hasFunc": true, "noFunc": false}
	expected.ShouldBeEqual(t, 0, "Map HasFunc", actual)
}

func Test_Cov6_Map_CompileToStrings(t *testing.T) {
	m := args.Map{"b": 2, "a": 1}
	lines := m.CompileToStrings()
	actual := args.Map{
		"linesLen": len(lines),
		"sorted":   lines[0] < lines[1],
	}
	expected := args.Map{"linesLen": 2, "sorted": true}
	expected.ShouldBeEqual(t, 0, "Map CompileToStrings", actual)
}

func Test_Cov6_Map_GoLiteralLines(t *testing.T) {
	m := args.Map{"key": "val"}
	lines := m.GoLiteralLines()
	actual := args.Map{"linesLen": len(lines) > 0}
	expected := args.Map{"linesLen": true}
	expected.ShouldBeEqual(t, 0, "Map GoLiteralLines", actual)
}

// ═══════════════════════════════════════════
// One through Six — basic
// ═══════════════════════════════════════════

func Test_Cov6_One_Basic(t *testing.T) {
	o := args.One{1}
	actual := args.Map{
		"first":  o.First(),
		"str":    o.String() != "",
		"len":    o.Length(),
		"count":  o.Count(),
		"all":    len(o.All()),
	}
	expected := args.Map{
		"first": 1, "str": true, "len": 1, "count": 1, "all": 1,
	}
	expected.ShouldBeEqual(t, 0, "One basic", actual)
}

func Test_Cov6_Two_Basic(t *testing.T) {
	tw := args.Two{1, 2}
	actual := args.Map{
		"first":  tw.First(),
		"second": tw.Second(),
		"len":    tw.Length(),
		"all":    len(tw.All()),
	}
	expected := args.Map{"first": 1, "second": 2, "len": 2, "all": 2}
	expected.ShouldBeEqual(t, 0, "Two basic", actual)
}

func Test_Cov6_Three_Basic(t *testing.T) {
	th := args.Three{1, 2, 3}
	actual := args.Map{
		"first":  th.First(),
		"second": th.Second(),
		"third":  th.Third(),
		"len":    th.Length(),
	}
	expected := args.Map{"first": 1, "second": 2, "third": 3, "len": 3}
	expected.ShouldBeEqual(t, 0, "Three basic", actual)
}

func Test_Cov6_Four_Basic(t *testing.T) {
	f := args.Four{1, 2, 3, 4}
	actual := args.Map{
		"first":  f.First(),
		"second": f.Second(),
		"third":  f.Third(),
		"fourth": f.Fourth(),
		"len":    f.Length(),
	}
	expected := args.Map{"first": 1, "second": 2, "third": 3, "fourth": 4, "len": 4}
	expected.ShouldBeEqual(t, 0, "Four basic", actual)
}

func Test_Cov6_Five_Basic(t *testing.T) {
	f := args.Five{1, 2, 3, 4, 5}
	actual := args.Map{
		"first":  f.First(),
		"second": f.Second(),
		"third":  f.Third(),
		"fourth": f.Fourth(),
		"fifth":  f.Fifth(),
		"len":    f.Length(),
	}
	expected := args.Map{
		"first": 1, "second": 2, "third": 3, "fourth": 4, "fifth": 5, "len": 5,
	}
	expected.ShouldBeEqual(t, 0, "Five basic", actual)
}

func Test_Cov6_Six_Basic(t *testing.T) {
	s := args.Six{1, 2, 3, 4, 5, 6}
	actual := args.Map{
		"first":  s.First(),
		"second": s.Second(),
		"third":  s.Third(),
		"fourth": s.Fourth(),
		"fifth":  s.Fifth(),
		"sixth":  s.Sixth(),
		"len":    s.Length(),
	}
	expected := args.Map{
		"first": 1, "second": 2, "third": 3, "fourth": 4,
		"fifth": 5, "sixth": 6, "len": 6,
	}
	expected.ShouldBeEqual(t, 0, "Six basic", actual)
}

// ═══════════════════════════════════════════
// Holder
// ═══════════════════════════════════════════

func Test_Cov6_Holder_Basic(t *testing.T) {
	h := args.Holder[string]{Value: "hello"}
	actual := args.Map{
		"val": h.Value,
	}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Holder basic", actual)
}

// ═══════════════════════════════════════════
// LeftRight — args
// ═══════════════════════════════════════════

func Test_Cov6_LeftRight_Basic(t *testing.T) {
	lr := args.LeftRight[string, int]{Left: "hello", Right: 42}
	lra := lr.LeftRightAny()
	actual := args.Map{
		"left":     lr.Left,
		"right":    lr.Right,
		"lraLeft":  lra.Left,
		"lraRight": lra.Right,
	}
	expected := args.Map{
		"left": "hello", "right": 42,
		"lraLeft": "hello", "lraRight": 42,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight basic", actual)
}

// ═══════════════════════════════════════════
// Dynamic
// ═══════════════════════════════════════════

func Test_Cov6_Dynamic_Basic(t *testing.T) {
	d := args.Dynamic[string]{Value: "hello"}
	actual := args.Map{
		"val": d.Value,
	}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Dynamic basic", actual)
}
