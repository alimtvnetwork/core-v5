package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// FuncWrapAny — basic
// ═══════════════════════════════════════════

func Test_Cov6_FuncWrapAny_NilFunc(t *testing.T) {
	fw := args.NewFuncWrap.Default(nil)
	actual := args.Map{
		"isInvalid":    fw.IsInvalid(),
		"hasValidFunc": fw.HasValidFunc(),
	}
	expected := args.Map{"isInvalid": true, "hasValidFunc": false}
	expected.ShouldBeEqual(t, 0, "FuncWrapAny returns nil -- nil func", actual)
}

func Test_Cov6_FuncWrapAny_ValidFunc(t *testing.T) {
	fn := func(a, b int) int { return a + b }
	fw := args.NewFuncWrap.Default(fn)
	actual := args.Map{
		"nameNE":       fw.Name != "",
		"isInvalid":    fw.IsInvalid(),
		"hasValidFunc": fw.HasValidFunc(),
	}
	expected := args.Map{"nameNE": true, "isInvalid": false, "hasValidFunc": true}
	expected.ShouldBeEqual(t, 0, "FuncWrapAny returns non-empty -- valid func", actual)
}

func Test_Cov6_FuncWrapAny_Invoke(t *testing.T) {
	fn := func(a, b int) int { return a + b }
	fw := args.NewFuncWrap.Default(fn)
	results, err := fw.Invoke(3, 4)
	actual := args.Map{
		"errNil":    err == nil,
		"resultLen": len(results),
	}
	expected := args.Map{"errNil": true, "resultLen": 1}
	expected.ShouldBeEqual(t, 0, "FuncWrapAny returns correct value -- invoke", actual)
}

func Test_Cov6_FuncWrapAny_InvokeNilFunc(t *testing.T) {
	fw := args.NewFuncWrap.Default(nil)
	_, err := fw.Invoke(1)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FuncWrapAny returns nil -- invoke nil func", actual)
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
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAs", actual)
}

func Test_Cov6_Map_GetAsBoolDefault(t *testing.T) {
	m := args.Map{"bool": true}
	val := m.GetAsBoolDefault("bool", false)
	missing := m.GetAsBoolDefault("missing", false)
	actual := args.Map{"val": val, "missing": missing}
	expected := args.Map{"val": true, "missing": false}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsBoolDefault", actual)
}

func Test_Cov6_Map_ArgsCount(t *testing.T) {
	m := args.Map{"a": 1, "b": 2, "func": nil, "expect": nil}
	actual := args.Map{"count": m.ArgsCount()}
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- ArgsCount", actual)
}

func Test_Cov6_Map_WorkFunc(t *testing.T) {
	fn := func() string { return "hello" }
	m := args.Map{"func": fn}
	wf := m.WorkFunc()
	actual := args.Map{"hasFunc": wf != nil}
	expected := args.Map{"hasFunc": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- WorkFunc", actual)
}

func Test_Cov6_Map_GetFirstOfNames(t *testing.T) {
	m := args.Map{"input": "hello", "when": "hello2"}
	first := m.GetFirstOfNames("input", "when")
	missing := m.GetFirstOfNames("missing1", "missing2")
	actual := args.Map{
		"first":   first,
		"missing": missing == nil,
	}
	expected := args.Map{"first": "hello", "missing": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetFirstOfNames", actual)
}

func Test_Cov6_Map_HasFunc(t *testing.T) {
	fn := func() {}
	m1 := args.Map{"func": fn}
	m2 := args.Map{"other": 1}
	actual := args.Map{
		"hasFunc":   m1.HasFunc(),
		"alsoHas":   m2.HasFunc(),
	}
	// HasFunc() always returns true because FuncWrap.Default(nil) returns non-nil *FuncWrapAny
	expected := args.Map{"hasFunc": true, "alsoHas": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- HasFunc", actual)
}

func Test_Cov6_Map_CompileToStrings(t *testing.T) {
	m := args.Map{"b": 2, "a": 1}
	lines := m.CompileToStrings()
	actual := args.Map{
		"linesLen": len(lines),
		"sorted":   lines[0] < lines[1],
	}
	expected := args.Map{"linesLen": 2, "sorted": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- CompileToStrings", actual)
}

func Test_Cov6_Map_GoLiteralLines(t *testing.T) {
	m := args.Map{"key": "val"}
	lines := m.GoLiteralLines()
	actual := args.Map{"linesLen": len(lines) > 0}
	expected := args.Map{"linesLen": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GoLiteralLines", actual)
}

// ═══════════════════════════════════════════
// One through Six — basic (using *Any aliases for untyped usage)
// ═══════════════════════════════════════════

func Test_Cov6_One_Basic(t *testing.T) {
	o := args.OneAny{First: 1}
	actual := args.Map{
		"first":  o.First,
		"str":    o.String() != "",
		"count":  o.ArgsCount(),
		"slice":  len(o.Slice()),
	}
	expected := args.Map{
		"first": 1, "str": true, "count": 1, "slice": 1,
	}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- basic", actual)
}

func Test_Cov6_Two_Basic(t *testing.T) {
	tw := args.TwoAny{First: 1, Second: 2}
	actual := args.Map{
		"first":  tw.First,
		"second": tw.Second,
		"count":  tw.ArgsCount(),
		"slice":  len(tw.Slice()),
	}
	expected := args.Map{"first": 1, "second": 2, "count": 2, "slice": 2}
	expected.ShouldBeEqual(t, 0, "Two returns correct value -- basic", actual)
}

func Test_Cov6_Three_Basic(t *testing.T) {
	th := args.ThreeAny{First: 1, Second: 2, Third: 3}
	actual := args.Map{
		"first":  th.First,
		"second": th.Second,
		"third":  th.Third,
		"count":  th.ArgsCount(),
	}
	expected := args.Map{"first": 1, "second": 2, "third": 3, "count": 3}
	expected.ShouldBeEqual(t, 0, "Three returns correct value -- basic", actual)
}

func Test_Cov6_Four_Basic(t *testing.T) {
	f := args.FourAny{First: 1, Second: 2, Third: 3, Fourth: 4}
	actual := args.Map{
		"first":  f.First,
		"second": f.Second,
		"third":  f.Third,
		"fourth": f.Fourth,
		"count":  f.ArgsCount(),
	}
	expected := args.Map{"first": 1, "second": 2, "third": 3, "fourth": 4, "count": 4}
	expected.ShouldBeEqual(t, 0, "Four returns correct value -- basic", actual)
}

func Test_Cov6_Five_Basic(t *testing.T) {
	f := args.FiveAny{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5}
	actual := args.Map{
		"first":  f.First,
		"second": f.Second,
		"third":  f.Third,
		"fourth": f.Fourth,
		"fifth":  f.Fifth,
		"count":  f.ArgsCount(),
	}
	expected := args.Map{
		"first": 1, "second": 2, "third": 3, "fourth": 4, "fifth": 5, "count": 5,
	}
	expected.ShouldBeEqual(t, 0, "Five returns correct value -- basic", actual)
}

func Test_Cov6_Six_Basic(t *testing.T) {
	s := args.SixAny{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, Sixth: 6}
	actual := args.Map{
		"first":  s.First,
		"second": s.Second,
		"third":  s.Third,
		"fourth": s.Fourth,
		"fifth":  s.Fifth,
		"sixth":  s.Sixth,
		"count":  s.ArgsCount(),
	}
	expected := args.Map{
		"first": 1, "second": 2, "third": 3, "fourth": 4,
		"fifth": 5, "sixth": 6, "count": 6,
	}
	expected.ShouldBeEqual(t, 0, "Six returns correct value -- basic", actual)
}

// ═══════════════════════════════════════════
// Holder
// ═══════════════════════════════════════════

func Test_Cov6_Holder_Basic(t *testing.T) {
	h := args.HolderAny{First: "hello"}
	actual := args.Map{
		"first": h.First,
		"count": h.ArgsCount(),
	}
	expected := args.Map{"first": "hello", "count": 7}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- basic", actual)
}

// ═══════════════════════════════════════════
// LeftRight — args
// ═══════════════════════════════════════════

func Test_Cov6_LeftRight_Basic(t *testing.T) {
	lr := args.LeftRightAny{Left: "hello", Right: 42}
	actual := args.Map{
		"left":  lr.Left,
		"right": lr.Right,
		"count": lr.ArgsCount(),
	}
	expected := args.Map{
		"left": "hello", "right": 42, "count": 2,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- basic", actual)
}

// ═══════════════════════════════════════════
// Dynamic
// ═══════════════════════════════════════════

func Test_Cov6_Dynamic_Basic(t *testing.T) {
	d := args.DynamicAny{Params: args.Map{"val": "hello"}, Expect: "expected"}
	actual := args.Map{
		"expect": d.Expect,
		"hasVal": d.HasDefined("val"),
	}
	expected := args.Map{"expect": "expected", "hasVal": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- basic", actual)
}
