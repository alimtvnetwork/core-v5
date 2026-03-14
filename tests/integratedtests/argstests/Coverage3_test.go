package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Map basic methods ──

func Test_Cov3_Map_Length(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	actual := args.Map{"length": m.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "Map.Length returns 2 -- two entries", actual)
}

func Test_Cov3_Map_ArgsCount(t *testing.T) {
	m := args.Map{"a": 1, "b": 2, "expected": 3}
	actual := args.Map{"count": m.ArgsCount()}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "Map.ArgsCount excludes expected -- 3 entries minus 1", actual)
}

func Test_Cov3_Map_Expected(t *testing.T) {
	m := args.Map{"expected": 42}
	actual := args.Map{"val": m.Expected()}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "Map.Expected returns value -- key 'expected'", actual)
}

func Test_Cov3_Map_HasExpect(t *testing.T) {
	m1 := args.Map{"expected": 42}
	m2 := args.Map{"a": 1}
	actual := args.Map{"has": m1.HasExpect(), "notHas": m2.HasExpect()}
	expected := args.Map{"has": true, "notHas": false}
	expected.ShouldBeEqual(t, 0, "Map.HasExpect returns correct -- with and without", actual)
}

func Test_Cov3_Map_HasFunc(t *testing.T) {
	m1 := args.Map{"func": func() {}}
	m2 := args.Map{"a": 1}
	actual := args.Map{"has": m1.HasFunc(), "notHas": m2.HasFunc()}
	expected := args.Map{"has": true, "notHas": false}
	expected.ShouldBeEqual(t, 0, "Map.HasFunc returns correct -- with and without", actual)
}

func Test_Cov3_Map_GetAs(t *testing.T) {
	m := args.Map{"name": "hello", "count": 42, "flag": true}
	name, _ := m.GetAsString("name")
	count, _ := m.GetAsInt("count")
	flag, _ := m.GetAsBool("flag")
	actual := args.Map{"name": name, "count": count, "flag": flag}
	expected := args.Map{"name": "hello", "count": 42, "flag": true}
	expected.ShouldBeEqual(t, 0, "Map.GetAs* returns correct types -- string, int, bool", actual)
}

func Test_Cov3_Map_GetAsStringSlice(t *testing.T) {
	m := args.Map{"items": []string{"a", "b"}}
	items, _ := m.GetAsStringSlice("items")
	actual := args.Map{"len": len(items)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Map.GetAsStringSlice returns correct -- 2 items", actual)
}

func Test_Cov3_Map_GetAsBytes(t *testing.T) {
	m := args.Map{"data": []byte{1, 2, 3}}
	data, _ := m.GetAsBytes("data")
	actual := args.Map{"len": len(data)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Map.GetAsBytes returns correct -- 3 bytes", actual)
}

func Test_Cov3_Map_WorkFunc(t *testing.T) {
	fn := func() string { return "hello" }
	m := args.Map{"func": fn}
	actual := args.Map{"notNil": m.WorkFunc() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Map.WorkFunc returns non-nil -- has func", actual)
}

func Test_Cov3_Map_GetFirstOfNames(t *testing.T) {
	m := args.Map{"name": "hello"}
	val := m.GetFirstOfNames("missing", "name", "other")
	actual := args.Map{"val": val}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Map.GetFirstOfNames finds first match -- name key", actual)
}

func Test_Cov3_Map_SortedKeys(t *testing.T) {
	m := args.Map{"c": 3, "a": 1, "b": 2}
	keys := m.SortedKeys()
	actual := args.Map{"first": keys[0], "last": keys[2]}
	expected := args.Map{"first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "Map.SortedKeys returns sorted -- 3 keys", actual)
}

func Test_Cov3_Map_String(t *testing.T) {
	m := args.Map{"key": "value"}
	s := m.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Map.String returns non-empty -- single entry", actual)
}

// ── One[T1] basic methods ──

func Test_Cov3_One_Basic(t *testing.T) {
	one := &args.One[string]{First: "hello", Expect: 42}
	actual := args.Map{
		"first":    one.FirstItem(),
		"expected": one.Expected(),
		"hasFirst": one.HasFirst(),
		"hasExpect": one.HasExpect(),
		"count":    one.ArgsCount(),
	}
	expected := args.Map{
		"first":    "hello",
		"expected": 42,
		"hasFirst": true,
		"hasExpect": true,
		"count":    1,
	}
	expected.ShouldBeEqual(t, 0, "One basic getters -- string first", actual)
}

func Test_Cov3_One_ArgTwo(t *testing.T) {
	one := &args.One[string]{First: "hello", Expect: 42}
	two := one.ArgTwo()
	actual := args.Map{"first": two.First, "expect": two.Expect}
	expected := args.Map{"first": "hello", "expect": 42}
	expected.ShouldBeEqual(t, 0, "One.ArgTwo returns copy -- same data", actual)
}

func Test_Cov3_One_Args(t *testing.T) {
	one := &args.One[string]{First: "hello"}
	a := one.Args()
	actual := args.Map{"len": len(a)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "One.Args returns 1 -- single first", actual)
}

func Test_Cov3_One_Slice(t *testing.T) {
	one := &args.One[string]{First: "hello"}
	s := one.Slice()
	actual := args.Map{"len": len(s)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "One.Slice returns 1 -- single first", actual)
}

func Test_Cov3_One_String(t *testing.T) {
	one := &args.One[string]{First: "hello"}
	s := one.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "One.String returns non-empty -- has first", actual)
}

func Test_Cov3_One_GetByIndex(t *testing.T) {
	one := &args.One[string]{First: "hello"}
	val := one.GetByIndex(0)
	actual := args.Map{"val": val}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "One.GetByIndex returns first -- index 0", actual)
}

func Test_Cov3_One_LeftRight(t *testing.T) {
	one := &args.One[string]{First: "hello", Expect: "world"}
	lr := one.LeftRight()
	actual := args.Map{"left": lr.Left, "right": lr.Right}
	expected := args.Map{"left": "hello", "right": "world"}
	expected.ShouldBeEqual(t, 0, "One.LeftRight returns left=first right=expect -- set", actual)
}

func Test_Cov3_One_AsOneParameter(t *testing.T) {
	one := &args.One[string]{First: "hello"}
	param := one.AsOneParameter()
	actual := args.Map{"notNil": param != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "One.AsOneParameter returns non-nil -- valid", actual)
}

// ── Two[T1, T2] basic methods ──

func Test_Cov3_Two_Basic(t *testing.T) {
	two := &args.Two[string, int]{First: "hello", Second: 42, Expect: true}
	actual := args.Map{
		"first":    two.FirstItem(),
		"second":   two.SecondItem(),
		"expected": two.Expected(),
		"hasFirst": two.HasFirst(),
		"hasSecond": two.HasSecond(),
		"hasExpect": two.HasExpect(),
		"count":    two.ArgsCount(),
	}
	expected := args.Map{
		"first":    "hello",
		"second":   42,
		"expected": true,
		"hasFirst": true,
		"hasSecond": true,
		"hasExpect": true,
		"count":    2,
	}
	expected.ShouldBeEqual(t, 0, "Two basic getters -- string and int", actual)
}

func Test_Cov3_Two_Args(t *testing.T) {
	two := &args.Two[string, int]{First: "hello", Second: 42}
	a := two.Args()
	actual := args.Map{"len": len(a)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Two.Args returns 2 -- first and second", actual)
}

func Test_Cov3_Two_Slice(t *testing.T) {
	two := &args.Two[string, int]{First: "hello", Second: 42}
	s := two.Slice()
	actual := args.Map{"len": len(s)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Two.Slice returns 2 -- first and second", actual)
}

func Test_Cov3_Two_GetByIndex(t *testing.T) {
	two := &args.Two[string, int]{First: "hello", Second: 42}
	actual := args.Map{
		"idx0": two.GetByIndex(0),
		"idx1": two.GetByIndex(1),
	}
	expected := args.Map{
		"idx0": "hello",
		"idx1": 42,
	}
	expected.ShouldBeEqual(t, 0, "Two.GetByIndex returns correct -- index 0 and 1", actual)
}

func Test_Cov3_Two_String(t *testing.T) {
	two := &args.Two[string, int]{First: "hello", Second: 42}
	s := two.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Two.String returns non-empty -- has items", actual)
}

func Test_Cov3_Two_LeftRight(t *testing.T) {
	two := &args.Two[string, int]{First: "hello", Second: 42}
	lr := two.LeftRight()
	actual := args.Map{"left": lr.Left, "right": lr.Right}
	expected := args.Map{"left": "hello", "right": 42}
	expected.ShouldBeEqual(t, 0, "Two.LeftRight returns first and second -- set", actual)
}

// ── Three[T1, T2, T3] basic methods ──

func Test_Cov3_Three_Basic(t *testing.T) {
	three := &args.Three[string, int, bool]{First: "hello", Second: 42, Third: true, Expect: "yes"}
	actual := args.Map{
		"first":    three.FirstItem(),
		"second":   three.SecondItem(),
		"third":    three.ThirdItem(),
		"expected": three.Expected(),
		"hasFirst": three.HasFirst(),
		"hasSecond": three.HasSecond(),
		"hasThird": three.HasThird(),
		"count":    three.ArgsCount(),
	}
	expected := args.Map{
		"first":    "hello",
		"second":   42,
		"third":    true,
		"expected": "yes",
		"hasFirst": true,
		"hasSecond": true,
		"hasThird": true,
		"count":    3,
	}
	expected.ShouldBeEqual(t, 0, "Three basic getters -- string, int, bool", actual)
}

func Test_Cov3_Three_Args(t *testing.T) {
	three := &args.Three[string, int, bool]{First: "a", Second: 1, Third: true}
	a := three.Args()
	actual := args.Map{"len": len(a)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Three.Args returns 3 -- all three", actual)
}

func Test_Cov3_Three_GetByIndex(t *testing.T) {
	three := &args.Three[string, int, bool]{First: "a", Second: 1, Third: true}
	actual := args.Map{
		"idx0": three.GetByIndex(0),
		"idx1": three.GetByIndex(1),
		"idx2": three.GetByIndex(2),
	}
	expected := args.Map{
		"idx0": "a",
		"idx1": 1,
		"idx2": true,
	}
	expected.ShouldBeEqual(t, 0, "Three.GetByIndex returns correct -- all indexes", actual)
}

func Test_Cov3_Three_String(t *testing.T) {
	three := &args.Three[string, int, bool]{First: "a", Second: 1, Third: true}
	s := three.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Three.String returns non-empty -- has items", actual)
}

// ── FuncWrap (via args.NewFuncWrap) ──

func Test_Cov3_FuncWrap_Basic(t *testing.T) {
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Func(fn)
	actual := args.Map{
		"isValid": fw.IsValid(),
		"name":    len(fw.Name()) > 0,
	}
	expected := args.Map{
		"isValid": true,
		"name":    true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap basic -- valid func", actual)
}

func Test_Cov3_FuncWrap_InOutArgs(t *testing.T) {
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Func(fn)
	actual := args.Map{
		"inCount":  fw.InArgsCount(),
		"outCount": fw.OutArgsCount(),
	}
	expected := args.Map{
		"inCount":  1,
		"outCount": 1,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap in/out args -- func(string)int", actual)
}

func Test_Cov3_FuncWrap_IsStringFunc(t *testing.T) {
	fn := func() string { return "hello" }
	fw := args.NewFuncWrap.Func(fn)
	actual := args.Map{"isString": fw.IsStringFunc()}
	expected := args.Map{"isString": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap.IsStringFunc returns true -- func()string", actual)
}

func Test_Cov3_FuncWrap_IsBoolFunc(t *testing.T) {
	fn := func() bool { return true }
	fw := args.NewFuncWrap.Func(fn)
	actual := args.Map{"isBool": fw.IsBoolFunc()}
	expected := args.Map{"isBool": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap.IsBoolFunc returns true -- func()bool", actual)
}

func Test_Cov3_FuncWrap_IsVoidFunc(t *testing.T) {
	fn := func() {}
	fw := args.NewFuncWrap.Func(fn)
	actual := args.Map{"isVoid": fw.IsVoidFunc()}
	expected := args.Map{"isVoid": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap.IsVoidFunc returns true -- func()", actual)
}

func Test_Cov3_FuncWrap_IsErrorFunc(t *testing.T) {
	fn := func() error { return nil }
	fw := args.NewFuncWrap.Func(fn)
	actual := args.Map{"isError": fw.IsErrorFunc()}
	expected := args.Map{"isError": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap.IsErrorFunc returns true -- func()error", actual)
}

func Test_Cov3_FuncWrap_String(t *testing.T) {
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Func(fn)
	s := fw.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap.String returns non-empty -- valid func", actual)
}

// ── Holder ──

func Test_Cov3_Holder_Basic(t *testing.T) {
	fn := func(s string) int { return len(s) }
	h := &args.Holder[string]{
		First:  "hello",
		Expect: 5,
		Func:   fn,
	}
	actual := args.Map{
		"first":    h.FirstItem(),
		"expected": h.Expected(),
		"hasFirst": h.HasFirst(),
		"hasExpect": h.HasExpect(),
		"hasFunc":  h.HasFunc(),
		"count":    h.ArgsCount(),
	}
	expected := args.Map{
		"first":    "hello",
		"expected": 5,
		"hasFirst": true,
		"hasExpect": true,
		"hasFunc":  true,
		"count":    1,
	}
	expected.ShouldBeEqual(t, 0, "Holder basic getters -- string first with func", actual)
}

func Test_Cov3_Holder_GetFuncName(t *testing.T) {
	fn := func() string { return "hello" }
	h := &args.Holder[string]{Func: fn}
	name := h.GetFuncName()
	actual := args.Map{"hasContent": len(name) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Holder.GetFuncName returns non-empty -- has func", actual)
}

func Test_Cov3_Holder_FuncWrap(t *testing.T) {
	fn := func() string { return "hello" }
	h := &args.Holder[string]{Func: fn}
	fw := h.FuncWrap()
	actual := args.Map{"isValid": fw.IsValid()}
	expected := args.Map{"isValid": true}
	expected.ShouldBeEqual(t, 0, "Holder.FuncWrap returns valid -- has func", actual)
}

// ── LeftRight ──

func Test_Cov3_LeftRight_Basic(t *testing.T) {
	lr := &args.LeftRight{Left: "hello", Right: "world"}
	actual := args.Map{
		"left":    lr.Left,
		"right":   lr.Right,
	}
	expected := args.Map{
		"left":    "hello",
		"right":   "world",
	}
	expected.ShouldBeEqual(t, 0, "LeftRight basic -- both set", actual)
}

func Test_Cov3_LeftRight_GetByIndex(t *testing.T) {
	lr := &args.LeftRight{Left: "L", Right: "R"}
	actual := args.Map{
		"idx0": lr.GetByIndex(0),
		"idx1": lr.GetByIndex(1),
	}
	expected := args.Map{
		"idx0": "L",
		"idx1": "R",
	}
	expected.ShouldBeEqual(t, 0, "LeftRight.GetByIndex returns correct -- both indexes", actual)
}

// ── String args ──

func Test_Cov3_String_Arg(t *testing.T) {
	sa := args.String{Value: "hello world", TrimAll: true}
	actual := args.Map{
		"val":     sa.Value,
		"trimAll": sa.TrimAll,
	}
	expected := args.Map{
		"val":     "hello world",
		"trimAll": true,
	}
	expected.ShouldBeEqual(t, 0, "String arg basic -- value and trimAll", actual)
}
