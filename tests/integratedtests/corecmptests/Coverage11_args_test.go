package corecmptests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════
// args.Map — coverage for Map methods
// ══════════════════════════════════════════════════════════════════

func Test_ArgsMap_Length(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Map Length", actual)
}

func Test_ArgsMap_ArgsCount_NoFuncNoExpect(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	actual := args.Map{"count": m.ArgsCount()}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Map ArgsCount no func no expect", actual)
}

func Test_ArgsMap_Has_Present(t *testing.T) {
	m := args.Map{"key": "val"}
	actual := args.Map{"has": m.Has("key")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Map Has present", actual)
}

func Test_ArgsMap_Has_Missing(t *testing.T) {
	m := args.Map{"key": "val"}
	actual := args.Map{"has": m.Has("missing")}
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "Map Has missing", actual)
}

func Test_ArgsMap_Has_NilMap(t *testing.T) {
	var m args.Map
	actual := args.Map{"has": m.Has("key")}
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "Map Has nil map", actual)
}

func Test_ArgsMap_HasDefined_Present(t *testing.T) {
	m := args.Map{"key": "val"}
	actual := args.Map{"has": m.HasDefined("key")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Map HasDefined present", actual)
}

func Test_ArgsMap_HasDefined_NilMap(t *testing.T) {
	var m args.Map
	actual := args.Map{"has": m.HasDefined("key")}
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "Map HasDefined nil map", actual)
}

func Test_ArgsMap_HasDefinedAll_Present(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	actual := args.Map{"has": m.HasDefinedAll("a", "b")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Map HasDefinedAll present", actual)
}

func Test_ArgsMap_HasDefinedAll_NilMap(t *testing.T) {
	var m args.Map
	actual := args.Map{"has": m.HasDefinedAll("a")}
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "Map HasDefinedAll nil map", actual)
}

func Test_ArgsMap_HasDefinedAll_NoNames(t *testing.T) {
	m := args.Map{"a": 1}
	actual := args.Map{"has": m.HasDefinedAll()}
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "Map HasDefinedAll no names", actual)
}

func Test_ArgsMap_IsKeyInvalid_NilMap(t *testing.T) {
	var m args.Map
	actual := args.Map{"invalid": m.IsKeyInvalid("key")}
	expected := args.Map{"invalid": false}
	expected.ShouldBeEqual(t, 0, "Map IsKeyInvalid nil map", actual)
}

func Test_ArgsMap_IsKeyMissing_NilMap(t *testing.T) {
	var m args.Map
	actual := args.Map{"missing": m.IsKeyMissing("key")}
	expected := args.Map{"missing": false}
	expected.ShouldBeEqual(t, 0, "Map IsKeyMissing nil map", actual)
}

func Test_ArgsMap_IsKeyMissing_Present(t *testing.T) {
	m := args.Map{"key": "val"}
	actual := args.Map{"missing": m.IsKeyMissing("key")}
	expected := args.Map{"missing": false}
	expected.ShouldBeEqual(t, 0, "Map IsKeyMissing present", actual)
}

func Test_ArgsMap_IsKeyMissing_Absent(t *testing.T) {
	m := args.Map{"key": "val"}
	actual := args.Map{"missing": m.IsKeyMissing("other")}
	expected := args.Map{"missing": true}
	expected.ShouldBeEqual(t, 0, "Map IsKeyMissing absent", actual)
}

func Test_ArgsMap_Get_NilMap(t *testing.T) {
	var m args.Map
	_, isValid := m.Get("key")
	actual := args.Map{"isValid": isValid}
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "Map Get nil map", actual)
}

func Test_ArgsMap_Get_Present(t *testing.T) {
	m := args.Map{"key": "val"}
	item, isValid := m.Get("key")
	actual := args.Map{"item": item, "isValid": isValid}
	expected := args.Map{"item": "val", "isValid": true}
	expected.ShouldBeEqual(t, 0, "Map Get present", actual)
}

func Test_ArgsMap_Get_Missing(t *testing.T) {
	m := args.Map{"key": "val"}
	_, isValid := m.Get("missing")
	actual := args.Map{"isValid": isValid}
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "Map Get missing", actual)
}

func Test_ArgsMap_GetAsInt(t *testing.T) {
	m := args.Map{"num": 42}
	val, ok := m.GetAsInt("num")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 42, "ok": true}
	expected.ShouldBeEqual(t, 0, "Map GetAsInt", actual)
}

func Test_ArgsMap_GetAsInt_Missing(t *testing.T) {
	m := args.Map{}
	val, ok := m.GetAsInt("num")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 0, "ok": false}
	expected.ShouldBeEqual(t, 0, "Map GetAsInt missing", actual)
}

func Test_ArgsMap_GetAsIntDefault(t *testing.T) {
	m := args.Map{"num": 42}
	actual := args.Map{"val": m.GetAsIntDefault("num", 99)}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "Map GetAsIntDefault found", actual)
}

func Test_ArgsMap_GetAsIntDefault_Missing(t *testing.T) {
	m := args.Map{}
	actual := args.Map{"val": m.GetAsIntDefault("num", 99)}
	expected := args.Map{"val": 99}
	expected.ShouldBeEqual(t, 0, "Map GetAsIntDefault missing", actual)
}

func Test_ArgsMap_GetAsBool(t *testing.T) {
	m := args.Map{"flag": true}
	val, ok := m.GetAsBool("flag")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": true, "ok": true}
	expected.ShouldBeEqual(t, 0, "Map GetAsBool", actual)
}

func Test_ArgsMap_GetAsBool_Missing(t *testing.T) {
	m := args.Map{}
	val, ok := m.GetAsBool("flag")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": false, "ok": false}
	expected.ShouldBeEqual(t, 0, "Map GetAsBool missing", actual)
}

func Test_ArgsMap_GetAsBoolDefault(t *testing.T) {
	m := args.Map{}
	actual := args.Map{"val": m.GetAsBoolDefault("flag", true)}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "Map GetAsBoolDefault fallback", actual)
}

func Test_ArgsMap_GetAsString(t *testing.T) {
	m := args.Map{"name": "hello"}
	val, ok := m.GetAsString("name")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": "hello", "ok": true}
	expected.ShouldBeEqual(t, 0, "Map GetAsString", actual)
}

func Test_ArgsMap_GetAsStringDefault(t *testing.T) {
	m := args.Map{}
	actual := args.Map{"val": m.GetAsStringDefault("name")}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "Map GetAsStringDefault empty", actual)
}

func Test_ArgsMap_GetAsStrings(t *testing.T) {
	m := args.Map{"items": []string{"a", "b"}}
	items, ok := m.GetAsStrings("items")
	actual := args.Map{"len": len(items), "ok": ok}
	expected := args.Map{"len": 2, "ok": true}
	expected.ShouldBeEqual(t, 0, "Map GetAsStrings", actual)
}

func Test_ArgsMap_GetAsStrings_Missing(t *testing.T) {
	m := args.Map{}
	items, ok := m.GetAsStrings("items")
	actual := args.Map{"len": len(items), "ok": ok}
	expected := args.Map{"len": 0, "ok": false}
	expected.ShouldBeEqual(t, 0, "Map GetAsStrings missing", actual)
}

func Test_ArgsMap_GetAsAnyItems(t *testing.T) {
	m := args.Map{"items": []any{1, "two"}}
	items, ok := m.GetAsAnyItems("items")
	actual := args.Map{"len": len(items), "ok": ok}
	expected := args.Map{"len": 2, "ok": true}
	expected.ShouldBeEqual(t, 0, "Map GetAsAnyItems", actual)
}

func Test_ArgsMap_GetAsAnyItems_Missing(t *testing.T) {
	m := args.Map{}
	items, ok := m.GetAsAnyItems("items")
	actual := args.Map{"len": len(items), "ok": ok}
	expected := args.Map{"len": 0, "ok": false}
	expected.ShouldBeEqual(t, 0, "Map GetAsAnyItems missing", actual)
}

func Test_ArgsMap_GetDirectLower(t *testing.T) {
	m := args.Map{"name": "hello"}
	actual := args.Map{"val": m.GetDirectLower("NAME")}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Map GetDirectLower", actual)
}

func Test_ArgsMap_GetDirectLower_Missing(t *testing.T) {
	m := args.Map{"name": "hello"}
	isNil := m.GetDirectLower("MISSING") == nil
	actual := args.Map{"isNil": isNil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Map GetDirectLower missing", actual)
}

func Test_ArgsMap_When(t *testing.T) {
	m := args.Map{"when": "now"}
	actual := args.Map{"val": m.When()}
	expected := args.Map{"val": "now"}
	expected.ShouldBeEqual(t, 0, "Map When", actual)
}

func Test_ArgsMap_Title(t *testing.T) {
	m := args.Map{"title": "test"}
	actual := args.Map{"val": m.Title()}
	expected := args.Map{"val": "test"}
	expected.ShouldBeEqual(t, 0, "Map Title", actual)
}

func Test_ArgsMap_SetActual(t *testing.T) {
	m := args.Map{}
	m.SetActual("result")
	actual := args.Map{"val": m.Actual()}
	expected := args.Map{"val": "result"}
	expected.ShouldBeEqual(t, 0, "Map SetActual", actual)
}

func Test_ArgsMap_GetFirstOfNames_Found(t *testing.T) {
	m := args.Map{"p2": "val"}
	actual := args.Map{"val": m.GetFirstOfNames("p1", "p2", "p3")}
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "Map GetFirstOfNames found", actual)
}

func Test_ArgsMap_GetFirstOfNames_Empty(t *testing.T) {
	m := args.Map{"x": "val"}
	isNil := m.GetFirstOfNames() == nil
	actual := args.Map{"isNil": isNil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Map GetFirstOfNames empty", actual)
}

func Test_ArgsMap_SortedKeys_Empty(t *testing.T) {
	m := args.Map{}
	keys, err := m.SortedKeys()
	actual := args.Map{"len": len(keys), "isNil": err == nil}
	expected := args.Map{"len": 0, "isNil": true}
	expected.ShouldBeEqual(t, 0, "Map SortedKeys empty", actual)
}

func Test_ArgsMap_SortedKeys_NonEmpty(t *testing.T) {
	m := args.Map{"b": 2, "a": 1}
	keys, _ := m.SortedKeys()
	actual := args.Map{"first": keys[0], "second": keys[1]}
	expected := args.Map{"first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "Map SortedKeys sorted", actual)
}

func Test_ArgsMap_CompileToStrings(t *testing.T) {
	m := args.Map{"b": 2, "a": 1}
	lines := m.CompileToStrings()
	actual := args.Map{"len": len(lines), "first": lines[0]}
	expected := args.Map{"len": 2, "first": "a : 1"}
	expected.ShouldBeEqual(t, 0, "Map CompileToStrings", actual)
}

func Test_ArgsMap_CompileToStrings_Empty(t *testing.T) {
	m := args.Map{}
	lines := m.CompileToStrings()
	actual := args.Map{"len": len(lines)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Map CompileToStrings empty", actual)
}

func Test_ArgsMap_CompileToString(t *testing.T) {
	m := args.Map{"a": 1}
	result := m.CompileToString()
	actual := args.Map{"val": result}
	expected := args.Map{"val": "a : 1"}
	expected.ShouldBeEqual(t, 0, "Map CompileToString", actual)
}

func Test_ArgsMap_GoLiteralLines(t *testing.T) {
	m := args.Map{"name": "test"}
	lines := m.GoLiteralLines()
	actual := args.Map{"len": len(lines)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Map GoLiteralLines", actual)
}

func Test_ArgsMap_GoLiteralLines_Empty(t *testing.T) {
	m := args.Map{}
	lines := m.GoLiteralLines()
	actual := args.Map{"len": len(lines)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Map GoLiteralLines empty", actual)
}

func Test_ArgsMap_GoLiteralString(t *testing.T) {
	m := args.Map{"a": 1}
	result := m.GoLiteralString()
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Map GoLiteralString", actual)
}

func Test_ArgsMap_GetByIndex(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	result := m.GetByIndex(0)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Map GetByIndex valid", actual)
}

func Test_ArgsMap_GetByIndex_OOB(t *testing.T) {
	m := args.Map{"a": 1}
	result := m.GetByIndex(10)
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Map GetByIndex OOB", actual)
}

func Test_ArgsMap_Slice(t *testing.T) {
	m := args.Map{"a": 1}
	slice := m.Slice()
	actual := args.Map{"len": len(slice)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Map Slice", actual)
}

func Test_ArgsMap_String_NonEmpty(t *testing.T) {
	m := args.Map{"key": "val"}
	result := m.String()
	actual := args.Map{"hasMap": strings.Contains(result, "Map")}
	expected := args.Map{"hasMap": true}
	expected.ShouldBeEqual(t, 0, "Map String non-empty", actual)
}

func Test_ArgsMap_Raw(t *testing.T) {
	m := args.Map{"a": 1}
	raw := m.Raw()
	actual := args.Map{"len": len(raw)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Map Raw", actual)
}

func Test_ArgsMap_Args(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	result := m.Args("a", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Map Args", actual)
}

func Test_ArgsMap_Expect(t *testing.T) {
	m := args.Map{"expect": "val"}
	actual := args.Map{"val": m.Expect()}
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "Map Expect", actual)
}

func Test_ArgsMap_Arrange(t *testing.T) {
	m := args.Map{"arrange": "data"}
	actual := args.Map{"val": m.Arrange()}
	expected := args.Map{"val": "data"}
	expected.ShouldBeEqual(t, 0, "Map Arrange", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.One — coverage
// ══════════════════════════════════════════════════════════════════

func Test_ArgsOne_All(t *testing.T) {
	one := args.OneAny{First: "hello", Expect: 42}
	actual := args.Map{
		"first":     one.FirstItem(),
		"expected":  one.Expected(),
		"hasFirst":  one.HasFirst(),
		"hasExpect": one.HasExpect(),
		"argsCount": one.ArgsCount(),
	}
	expected := args.Map{
		"first":     "hello",
		"expected":  42,
		"hasFirst":  true,
		"hasExpect": true,
		"argsCount": 1,
	}
	expected.ShouldBeEqual(t, 0, "One all methods", actual)
}

func Test_ArgsOne_Slice(t *testing.T) {
	one := args.OneAny{First: "hello", Expect: 42}
	slice := one.Slice()
	actual := args.Map{"len": len(slice)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "One Slice", actual)
}

func Test_ArgsOne_SliceCached(t *testing.T) {
	one := args.OneAny{First: "hello"}
	_ = one.Slice()
	slice := one.Slice() // cached
	actual := args.Map{"len": len(slice)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "One Slice cached", actual)
}

func Test_ArgsOne_String(t *testing.T) {
	one := args.OneAny{First: "hello"}
	result := one.String()
	actual := args.Map{"hasOne": strings.Contains(result, "One")}
	expected := args.Map{"hasOne": true}
	expected.ShouldBeEqual(t, 0, "One String", actual)
}

func Test_ArgsOne_GetByIndex(t *testing.T) {
	one := args.OneAny{First: "hello"}
	actual := args.Map{"val": one.GetByIndex(0)}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "One GetByIndex", actual)
}

func Test_ArgsOne_GetByIndex_OOB(t *testing.T) {
	one := args.OneAny{First: "hello"}
	isNil := one.GetByIndex(10) == nil
	actual := args.Map{"isNil": isNil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "One GetByIndex OOB", actual)
}

func Test_ArgsOne_Args(t *testing.T) {
	one := args.OneAny{First: "hello"}
	result := one.Args(1)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "One Args upTo 1", actual)
}

func Test_ArgsOne_Args_Zero(t *testing.T) {
	one := args.OneAny{First: "hello"}
	result := one.Args(0)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "One Args upTo 0", actual)
}

func Test_ArgsOne_ValidArgs(t *testing.T) {
	one := args.OneAny{First: "hello"}
	result := one.ValidArgs()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "One ValidArgs", actual)
}

func Test_ArgsOne_LeftRight(t *testing.T) {
	one := args.OneAny{First: "hello", Expect: "exp"}
	lr := one.LeftRight()
	actual := args.Map{"left": lr.Left, "expect": lr.Expect}
	expected := args.Map{"left": "hello", "expect": "exp"}
	expected.ShouldBeEqual(t, 0, "One LeftRight", actual)
}

func Test_ArgsOne_ArgTwo(t *testing.T) {
	one := args.OneAny{First: "hello", Expect: "exp"}
	two := one.ArgTwo()
	actual := args.Map{"first": two.First, "expect": two.Expect}
	expected := args.Map{"first": "hello", "expect": "exp"}
	expected.ShouldBeEqual(t, 0, "One ArgTwo", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Two — coverage
// ══════════════════════════════════════════════════════════════════

func Test_ArgsTwo_All(t *testing.T) {
	two := args.TwoAny{First: "a", Second: "b", Expect: 1}
	actual := args.Map{
		"first":     two.FirstItem(),
		"second":    two.SecondItem(),
		"expected":  two.Expected(),
		"hasFirst":  two.HasFirst(),
		"hasSecond": two.HasSecond(),
		"argsCount": two.ArgsCount(),
	}
	expected := args.Map{
		"first": "a", "second": "b", "expected": 1,
		"hasFirst": true, "hasSecond": true, "argsCount": 2,
	}
	expected.ShouldBeEqual(t, 0, "Two all methods", actual)
}

func Test_ArgsTwo_Slice(t *testing.T) {
	two := args.TwoAny{First: "a", Second: "b"}
	actual := args.Map{"len": len(two.Slice())}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Two Slice", actual)
}

func Test_ArgsTwo_String(t *testing.T) {
	two := args.TwoAny{First: "a", Second: "b"}
	result := two.String()
	actual := args.Map{"hasTwo": strings.Contains(result, "Two")}
	expected := args.Map{"hasTwo": true}
	expected.ShouldBeEqual(t, 0, "Two String", actual)
}

func Test_ArgsTwo_Args(t *testing.T) {
	two := args.TwoAny{First: "a", Second: "b"}
	actual := args.Map{
		"args0": len(two.Args(0)),
		"args1": len(two.Args(1)),
		"args2": len(two.Args(2)),
	}
	expected := args.Map{"args0": 0, "args1": 1, "args2": 2}
	expected.ShouldBeEqual(t, 0, "Two Args", actual)
}

func Test_ArgsTwo_LeftRight(t *testing.T) {
	two := args.TwoAny{First: "a", Second: "b"}
	lr := two.LeftRight()
	actual := args.Map{"left": lr.Left, "right": lr.Right}
	expected := args.Map{"left": "a", "right": "b"}
	expected.ShouldBeEqual(t, 0, "Two LeftRight", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Three — coverage
// ══════════════════════════════════════════════════════════════════

func Test_ArgsThree_All(t *testing.T) {
	three := args.ThreeAny{First: "a", Second: "b", Third: "c", Expect: 1}
	actual := args.Map{
		"first":     three.FirstItem(),
		"second":    three.SecondItem(),
		"third":     three.ThirdItem(),
		"hasThird":  three.HasThird(),
		"argsCount": three.ArgsCount(),
	}
	expected := args.Map{
		"first": "a", "second": "b", "third": "c",
		"hasThird": true, "argsCount": 3,
	}
	expected.ShouldBeEqual(t, 0, "Three all methods", actual)
}

func Test_ArgsThree_ArgTwo(t *testing.T) {
	three := args.ThreeAny{First: "a", Second: "b", Third: "c"}
	two := three.ArgTwo()
	actual := args.Map{"first": two.First, "second": two.Second}
	expected := args.Map{"first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "Three ArgTwo", actual)
}

func Test_ArgsThree_ArgThree(t *testing.T) {
	three := args.ThreeAny{First: "a", Second: "b", Third: "c"}
	copy := three.ArgThree()
	actual := args.Map{"third": copy.Third}
	expected := args.Map{"third": "c"}
	expected.ShouldBeEqual(t, 0, "Three ArgThree", actual)
}

func Test_ArgsThree_Args(t *testing.T) {
	three := args.ThreeAny{First: "a", Second: "b", Third: "c"}
	actual := args.Map{"args3": len(three.Args(3))}
	expected := args.Map{"args3": 3}
	expected.ShouldBeEqual(t, 0, "Three Args", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Four — coverage
// ══════════════════════════════════════════════════════════════════

func Test_ArgsFour_All(t *testing.T) {
	four := args.FourAny{First: "a", Second: "b", Third: "c", Fourth: "d"}
	actual := args.Map{
		"fourth":    four.FourthItem(),
		"hasFourth": four.HasFourth(),
		"argsCount": four.ArgsCount(),
	}
	expected := args.Map{"fourth": "d", "hasFourth": true, "argsCount": 4}
	expected.ShouldBeEqual(t, 0, "Four all methods", actual)
}

func Test_ArgsFour_Args(t *testing.T) {
	four := args.FourAny{First: "a", Second: "b", Third: "c", Fourth: "d"}
	actual := args.Map{"args4": len(four.Args(4))}
	expected := args.Map{"args4": 4}
	expected.ShouldBeEqual(t, 0, "Four Args", actual)
}

func Test_ArgsFour_ArgTwo(t *testing.T) {
	four := args.FourAny{First: "a", Second: "b", Third: "c", Fourth: "d"}
	two := four.ArgTwo()
	actual := args.Map{"first": two.First}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "Four ArgTwo", actual)
}

func Test_ArgsFour_ArgThree(t *testing.T) {
	four := args.FourAny{First: "a", Second: "b", Third: "c", Fourth: "d"}
	three := four.ArgThree()
	actual := args.Map{"third": three.Third}
	expected := args.Map{"third": "c"}
	expected.ShouldBeEqual(t, 0, "Four ArgThree", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Five — coverage
// ══════════════════════════════════════════════════════════════════

func Test_ArgsFive_All(t *testing.T) {
	five := args.FiveAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e"}
	actual := args.Map{
		"fifth":    five.FifthItem(),
		"hasFifth": five.HasFifth(),
		"count":    five.ArgsCount(),
	}
	expected := args.Map{"fifth": "e", "hasFifth": true, "count": 5}
	expected.ShouldBeEqual(t, 0, "Five all methods", actual)
}

func Test_ArgsFive_Args(t *testing.T) {
	five := args.FiveAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e"}
	actual := args.Map{"args5": len(five.Args(5))}
	expected := args.Map{"args5": 5}
	expected.ShouldBeEqual(t, 0, "Five Args", actual)
}

func Test_ArgsFive_ArgFour(t *testing.T) {
	five := args.FiveAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e"}
	four := five.ArgFour()
	actual := args.Map{"fourth": four.Fourth}
	expected := args.Map{"fourth": "d"}
	expected.ShouldBeEqual(t, 0, "Five ArgFour", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Six — coverage
// ══════════════════════════════════════════════════════════════════

func Test_ArgsSix_All(t *testing.T) {
	six := args.SixAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f"}
	actual := args.Map{
		"sixth":    six.SixthItem(),
		"hasSixth": six.HasSixth(),
		"count":    six.ArgsCount(),
	}
	expected := args.Map{"sixth": "f", "hasSixth": true, "count": 6}
	expected.ShouldBeEqual(t, 0, "Six all methods", actual)
}

func Test_ArgsSix_Args(t *testing.T) {
	six := args.SixAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f"}
	actual := args.Map{"args6": len(six.Args(6))}
	expected := args.Map{"args6": 6}
	expected.ShouldBeEqual(t, 0, "Six Args", actual)
}

func Test_ArgsSix_ArgFive(t *testing.T) {
	six := args.SixAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f"}
	five := six.ArgFive()
	actual := args.Map{"fifth": five.Fifth}
	expected := args.Map{"fifth": "e"}
	expected.ShouldBeEqual(t, 0, "Six ArgFive", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.LeftRight — coverage
// ══════════════════════════════════════════════════════════════════

func Test_ArgsLeftRight_All(t *testing.T) {
	lr := args.LeftRightAny{Left: "a", Right: "b", Expect: 1}
	actual := args.Map{
		"left":      lr.FirstItem(),
		"right":     lr.SecondItem(),
		"hasLeft":   lr.HasLeft(),
		"hasRight":  lr.HasRight(),
		"hasFirst":  lr.HasFirst(),
		"hasSecond": lr.HasSecond(),
		"count":     lr.ArgsCount(),
	}
	expected := args.Map{
		"left": "a", "right": "b",
		"hasLeft": true, "hasRight": true,
		"hasFirst": true, "hasSecond": true,
		"count": 2,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight all methods", actual)
}

func Test_ArgsLeftRight_Clone(t *testing.T) {
	lr := args.LeftRightAny{Left: "a", Right: "b"}
	cloned := lr.Clone()
	actual := args.Map{"left": cloned.Left, "right": cloned.Right}
	expected := args.Map{"left": "a", "right": "b"}
	expected.ShouldBeEqual(t, 0, "LeftRight Clone", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.String — coverage
// ══════════════════════════════════════════════════════════════════

func Test_ArgsString_Methods(t *testing.T) {
	s := args.String("hello")
	actual := args.Map{
		"string":    s.String(),
		"length":    s.Length(),
		"count":     s.Count(),
		"isEmpty":   s.IsEmpty(),
		"isDefined": s.IsDefined(),
		"hasCh":     s.HasCharacter(),
		"asciiLen":  s.AscIILength(),
	}
	expected := args.Map{
		"string": "hello", "length": 5, "count": 5,
		"isEmpty": false, "isDefined": true, "hasCh": true,
		"asciiLen": 5,
	}
	expected.ShouldBeEqual(t, 0, "String basic methods", actual)
}

func Test_ArgsString_IsEmptyTrue(t *testing.T) {
	s := args.String("")
	actual := args.Map{"isEmpty": s.IsEmpty(), "isEW": s.IsEmptyOrWhitespace()}
	expected := args.Map{"isEmpty": true, "isEW": true}
	expected.ShouldBeEqual(t, 0, "String empty", actual)
}

func Test_ArgsString_TrimSpace(t *testing.T) {
	s := args.String("  hello  ")
	actual := args.Map{"val": s.TrimSpace().String()}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "String TrimSpace", actual)
}

func Test_ArgsString_ReplaceAll(t *testing.T) {
	s := args.String("hello world")
	actual := args.Map{"val": s.ReplaceAll("world", "go").String()}
	expected := args.Map{"val": "hello go"}
	expected.ShouldBeEqual(t, 0, "String ReplaceAll", actual)
}

func Test_ArgsString_Concat(t *testing.T) {
	s := args.String("hello")
	actual := args.Map{"val": s.Concat(" ", "world").String()}
	expected := args.Map{"val": "hello world"}
	expected.ShouldBeEqual(t, 0, "String Concat", actual)
}

func Test_ArgsString_Join(t *testing.T) {
	s := args.String("hello")
	actual := args.Map{"val": s.Join(",", "a", "b").String()}
	expected := args.Map{"val": "hello,a,b"}
	expected.ShouldBeEqual(t, 0, "String Join", actual)
}

func Test_ArgsString_Split(t *testing.T) {
	s := args.String("a,b,c")
	result := s.Split(",")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "String Split", actual)
}

func Test_ArgsString_Bytes(t *testing.T) {
	s := args.String("hi")
	actual := args.Map{"len": len(s.Bytes())}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "String Bytes", actual)
}

func Test_ArgsString_Runes(t *testing.T) {
	s := args.String("hi")
	actual := args.Map{"len": len(s.Runes())}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "String Runes", actual)
}

func Test_ArgsString_Substring(t *testing.T) {
	s := args.String("hello")
	actual := args.Map{"val": s.Substring(1, 4).String()}
	expected := args.Map{"val": "ell"}
	expected.ShouldBeEqual(t, 0, "String Substring", actual)
}

func Test_ArgsString_DoubleQuote(t *testing.T) {
	s := args.String("hi")
	result := s.DoubleQuote().String()
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "String DoubleQuote", actual)
}

func Test_ArgsString_DoubleQuoteQ(t *testing.T) {
	s := args.String("hi")
	result := s.DoubleQuoteQ().String()
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "String DoubleQuoteQ", actual)
}

func Test_ArgsString_SingleQuote(t *testing.T) {
	s := args.String("hi")
	result := s.SingleQuote().String()
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "String SingleQuote", actual)
}

func Test_ArgsString_ValueDoubleQuote(t *testing.T) {
	s := args.String("hi")
	result := s.ValueDoubleQuote().String()
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "String ValueDoubleQuote", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.FuncWrap — coverage
// ══════════════════════════════════════════════════════════════════

func sampleFunc(s string) int { return len(s) }

func Test_NewFuncWrap_Default(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	actual := args.Map{
		"isValid":   fw.IsValid(),
		"argsCount": fw.ArgsCount(),
		"retLen":    fw.ReturnLength(),
	}
	expected := args.Map{"isValid": true, "argsCount": 1, "retLen": 1}
	expected.ShouldBeEqual(t, 0, "FuncWrap Default", actual)
}

func Test_NewFuncWrap_Default_Nil(t *testing.T) {
	fw := args.NewFuncWrap.Default(nil)
	actual := args.Map{"isValid": fw.IsValid()}
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "FuncWrap Default nil", actual)
}

func Test_NewFuncWrap_Default_NotFunc(t *testing.T) {
	fw := args.NewFuncWrap.Default("not a func")
	actual := args.Map{"isValid": fw.IsValid(), "isInvalid": fw.IsInvalid()}
	expected := args.Map{"isValid": false, "isInvalid": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap Default not func", actual)
}

func Test_NewFuncWrap_Invalid(t *testing.T) {
	fw := args.NewFuncWrap.Invalid()
	actual := args.Map{"isInvalid": fw.IsInvalid()}
	expected := args.Map{"isInvalid": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap Invalid", actual)
}

func Test_FuncWrap_Invoke(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	results, err := fw.Invoke("hello")
	actual := args.Map{"result": results[0], "err": err == nil}
	expected := args.Map{"result": 5, "err": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap Invoke", actual)
}

func Test_FuncWrap_InvokeMust(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	results := fw.InvokeMust("hi")
	actual := args.Map{"result": results[0]}
	expected := args.Map{"result": 2}
	expected.ShouldBeEqual(t, 0, "FuncWrap InvokeMust", actual)
}

func Test_FuncWrap_GetFuncName(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	actual := args.Map{"name": fw.GetFuncName()}
	expected := args.Map{"name": "sampleFunc"}
	expected.ShouldBeEqual(t, 0, "FuncWrap GetFuncName", actual)
}

func Test_FuncWrap_GetFuncName_Nil(t *testing.T) {
	var fw *args.FuncWrapAny
	actual := args.Map{"name": fw.GetFuncName()}
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "FuncWrap GetFuncName nil", actual)
}

func Test_FuncWrap_HasValidFunc(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	actual := args.Map{"valid": fw.HasValidFunc()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap HasValidFunc", actual)
}

func Test_FuncWrap_ValidationError_Valid(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	actual := args.Map{"isNil": fw.ValidationError() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap ValidationError valid", actual)
}

func Test_FuncWrap_ValidationError_Nil(t *testing.T) {
	var fw *args.FuncWrapAny
	actual := args.Map{"hasErr": fw.ValidationError() != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap ValidationError nil", actual)
}

func Test_FuncWrap_ValidationError_Invalid(t *testing.T) {
	fw := args.NewFuncWrap.Default("not a func")
	actual := args.Map{"hasErr": fw.ValidationError() != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap ValidationError invalid", actual)
}

func Test_FuncWrap_InvalidError_Nil(t *testing.T) {
	var fw *args.FuncWrapAny
	actual := args.Map{"hasErr": fw.InvalidError() != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap InvalidError nil", actual)
}

func Test_FuncWrap_InvalidError_Valid(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	actual := args.Map{"isNil": fw.InvalidError() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap InvalidError valid", actual)
}

func Test_FuncWrap_GetInArgsTypes(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	types := fw.GetInArgsTypes()
	actual := args.Map{"len": len(types)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FuncWrap GetInArgsTypes", actual)
}

func Test_FuncWrap_GetOutArgsTypes(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	types := fw.GetOutArgsTypes()
	actual := args.Map{"len": len(types)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FuncWrap GetOutArgsTypes", actual)
}

func Test_FuncWrap_GetInArgsTypesNames(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	names := fw.GetInArgsTypesNames()
	actual := args.Map{"len": len(names), "first": names[0]}
	expected := args.Map{"len": 1, "first": "string"}
	expected.ShouldBeEqual(t, 0, "FuncWrap GetInArgsTypesNames", actual)
}

func Test_FuncWrap_GetOutArgsTypesNames(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	names := fw.GetOutArgsTypesNames()
	actual := args.Map{"len": len(names), "first": names[0]}
	expected := args.Map{"len": 1, "first": "int"}
	expected.ShouldBeEqual(t, 0, "FuncWrap GetOutArgsTypesNames", actual)
}

func Test_FuncWrap_InArgNames(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	names := fw.InArgNames()
	actual := args.Map{"len": len(names)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FuncWrap InArgNames", actual)
}

func Test_FuncWrap_OutArgNames(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	names := fw.OutArgNames()
	actual := args.Map{"len": len(names)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FuncWrap OutArgNames", actual)
}

func Test_FuncWrap_IsStringFunc(t *testing.T) {
	fw := args.NewFuncWrap.Default(func() string { return "" })
	actual := args.Map{"isString": fw.IsStringFunc()}
	expected := args.Map{"isString": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap IsStringFunc", actual)
}

func Test_FuncWrap_IsBoolFunc(t *testing.T) {
	fw := args.NewFuncWrap.Default(func() bool { return false })
	actual := args.Map{"isBool": fw.IsBoolFunc()}
	expected := args.Map{"isBool": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap IsBoolFunc", actual)
}

func Test_FuncWrap_IsVoidFunc(t *testing.T) {
	fw := args.NewFuncWrap.Default(func() {})
	actual := args.Map{"isVoid": fw.IsVoidFunc()}
	expected := args.Map{"isVoid": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap IsVoidFunc", actual)
}

func Test_FuncWrap_IsAnyFunc(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	actual := args.Map{"isAny": fw.IsAnyFunc()}
	expected := args.Map{"isAny": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap IsAnyFunc", actual)
}

func Test_FuncWrap_InvokeAsBool(t *testing.T) {
	fw := args.NewFuncWrap.Default(func() bool { return true })
	val, err := fw.InvokeAsBool()
	actual := args.Map{"val": val, "noErr": err == nil}
	expected := args.Map{"val": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap InvokeAsBool", actual)
}

func Test_FuncWrap_InvokeAsString(t *testing.T) {
	fw := args.NewFuncWrap.Default(func() string { return "hi" })
	val, err := fw.InvokeAsString()
	actual := args.Map{"val": val, "noErr": err == nil}
	expected := args.Map{"val": "hi", "noErr": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap InvokeAsString", actual)
}

func Test_FuncWrap_InvokeAsAny(t *testing.T) {
	fw := args.NewFuncWrap.Default(func() int { return 42 })
	val, err := fw.InvokeAsAny()
	actual := args.Map{"val": val, "noErr": err == nil}
	expected := args.Map{"val": 42, "noErr": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap InvokeAsAny", actual)
}

func Test_FuncWrap_VoidCall(t *testing.T) {
	called := false
	fw := args.NewFuncWrap.Default(func() { called = true })
	_, err := fw.VoidCall()
	actual := args.Map{"called": called, "noErr": err == nil}
	expected := args.Map{"called": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap VoidCall", actual)
}

func Test_FuncWrap_IsEqual_Same(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	actual := args.Map{"equal": fw.IsEqual(fw)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap IsEqual same", actual)
}

func Test_FuncWrap_IsNotEqual(t *testing.T) {
	fw1 := args.NewFuncWrap.Default(sampleFunc)
	fw2 := args.NewFuncWrap.Default(func() {})
	actual := args.Map{"notEqual": fw1.IsNotEqual(fw2)}
	expected := args.Map{"notEqual": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap IsNotEqual", actual)
}

func Test_FuncWrap_PkgPath(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	path := fw.PkgPath()
	actual := args.Map{"hasPath": len(path) > 0}
	expected := args.Map{"hasPath": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap PkgPath", actual)
}

func Test_FuncWrap_PkgNameOnly(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	name := fw.PkgNameOnly()
	actual := args.Map{"hasName": len(name) > 0}
	expected := args.Map{"hasName": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap PkgNameOnly", actual)
}

func Test_FuncWrap_FuncDirectInvokeName(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	name := fw.FuncDirectInvokeName()
	actual := args.Map{"hasName": len(name) > 0}
	expected := args.Map{"hasName": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap FuncDirectInvokeName", actual)
}

func Test_FuncWrap_GetType(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	actual := args.Map{"notNil": fw.GetType() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap GetType", actual)
}

func Test_FuncWrap_IsInTypeMatches(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	actual := args.Map{"matches": fw.IsInTypeMatches("hello")}
	expected := args.Map{"matches": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap IsInTypeMatches", actual)
}

func Test_FuncWrap_GetPascalCaseFuncName(t *testing.T) {
	fw := args.NewFuncWrap.Default(sampleFunc)
	actual := args.Map{"name": fw.GetPascalCaseFuncName()}
	expected := args.Map{"name": "SampleFunc"}
	expected.ShouldBeEqual(t, 0, "FuncWrap GetPascalCaseFuncName", actual)
}

func Test_FuncWrap_GetPascalCaseFuncName_Nil(t *testing.T) {
	var fw *args.FuncWrapAny
	actual := args.Map{"name": fw.GetPascalCaseFuncName()}
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "FuncWrap GetPascalCaseFuncName nil", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.NewFuncWrap — Map, Many, Single
// ══════════════════════════════════════════════════════════════════

func Test_NewFuncWrap_Map(t *testing.T) {
	fm := args.NewFuncWrap.Map(sampleFunc)
	actual := args.Map{"len": fm.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap Map", actual)
}

func Test_NewFuncWrap_Map_Empty(t *testing.T) {
	fm := args.NewFuncWrap.Map()
	actual := args.Map{"len": fm.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap Map empty", actual)
}

func Test_NewFuncWrap_Many(t *testing.T) {
	wraps := args.NewFuncWrap.Many(sampleFunc)
	actual := args.Map{"len": len(wraps)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap Many", actual)
}

func Test_NewFuncWrap_Many_Empty(t *testing.T) {
	wraps := args.NewFuncWrap.Many()
	actual := args.Map{"len": len(wraps)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap Many empty", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Empty — coverage
// ══════════════════════════════════════════════════════════════════

func Test_Empty_Map(t *testing.T) {
	m := args.Empty.Map()
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty Map", actual)
}

func Test_Empty_FuncWrap(t *testing.T) {
	fw := args.Empty.FuncWrap()
	actual := args.Map{"isInvalid": fw.IsInvalid()}
	expected := args.Map{"isInvalid": true}
	expected.ShouldBeEqual(t, 0, "Empty FuncWrap", actual)
}

func Test_Empty_FuncMap(t *testing.T) {
	fm := args.Empty.FuncMap()
	actual := args.Map{"isEmpty": fm.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Empty FuncMap", actual)
}

func Test_Empty_Holder(t *testing.T) {
	h := args.Empty.Holder()
	actual := args.Map{"count": h.ArgsCount()}
	expected := args.Map{"count": 7}
	expected.ShouldBeEqual(t, 0, "Empty Holder", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.FuncMap — coverage
// ══════════════════════════════════════════════════════════════════

func Test_FuncMap_Has(t *testing.T) {
	fm := args.NewFuncWrap.Map(sampleFunc)
	actual := args.Map{
		"has":        fm.Has("sampleFunc"),
		"hasMissing": fm.Has("missing"),
		"isContains": fm.IsContains("sampleFunc"),
	}
	expected := args.Map{"has": true, "hasMissing": false, "isContains": true}
	expected.ShouldBeEqual(t, 0, "FuncMap Has", actual)
}

func Test_FuncMap_Get(t *testing.T) {
	fm := args.NewFuncWrap.Map(sampleFunc)
	f := fm.Get("sampleFunc")
	actual := args.Map{"notNil": f != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FuncMap Get", actual)
}

func Test_FuncMap_Get_Missing(t *testing.T) {
	fm := args.NewFuncWrap.Map(sampleFunc)
	f := fm.Get("missing")
	actual := args.Map{"isNil": f == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FuncMap Get missing", actual)
}

func Test_FuncMap_IsValidFuncOf(t *testing.T) {
	fm := args.NewFuncWrap.Map(sampleFunc)
	actual := args.Map{"valid": fm.IsValidFuncOf("sampleFunc")}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "FuncMap IsValidFuncOf", actual)
}

func Test_FuncMap_IsInvalidFunc(t *testing.T) {
	fm := args.NewFuncWrap.Map(sampleFunc)
	actual := args.Map{"invalid": fm.IsInvalidFunc("missing")}
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "FuncMap IsInvalidFunc missing", actual)
}

func Test_FuncMap_HasAnyItem(t *testing.T) {
	fm := args.NewFuncWrap.Map(sampleFunc)
	actual := args.Map{"hasAny": fm.HasAnyItem()}
	expected := args.Map{"hasAny": true}
	expected.ShouldBeEqual(t, 0, "FuncMap HasAnyItem", actual)
}

func Test_FuncMap_ArgsCount(t *testing.T) {
	fm := args.NewFuncWrap.Map(sampleFunc)
	actual := args.Map{"count": fm.ArgsCount("sampleFunc")}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "FuncMap ArgsCount", actual)
}

func Test_FuncMap_ArgsCount_Missing(t *testing.T) {
	fm := args.NewFuncWrap.Map(sampleFunc)
	actual := args.Map{"count": fm.ArgsCount("missing")}
	expected := args.Map{"count": 0}
	expected.ShouldBeEqual(t, 0, "FuncMap ArgsCount missing", actual)
}

func Test_FuncMap_ReturnLength(t *testing.T) {
	fm := args.NewFuncWrap.Map(sampleFunc)
	actual := args.Map{"retLen": fm.ReturnLength("sampleFunc")}
	expected := args.Map{"retLen": 1}
	expected.ShouldBeEqual(t, 0, "FuncMap ReturnLength", actual)
}

func Test_FuncMap_ReturnLength_Missing(t *testing.T) {
	fm := args.NewFuncWrap.Map(sampleFunc)
	actual := args.Map{"retLen": fm.ReturnLength("missing")}
	expected := args.Map{"retLen": 0}
	expected.ShouldBeEqual(t, 0, "FuncMap ReturnLength missing", actual)
}

func Test_FuncMap_InvalidError_Empty(t *testing.T) {
	fm := args.FuncMap{}
	actual := args.Map{"hasErr": fm.InvalidError() != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FuncMap InvalidError empty", actual)
}

func Test_FuncMap_InvalidError_NonEmpty(t *testing.T) {
	fm := args.NewFuncWrap.Map(sampleFunc)
	actual := args.Map{"isNil": fm.InvalidError() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FuncMap InvalidError non-empty", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.FuncDetector — coverage
// ══════════════════════════════════════════════════════════════════

func Test_FuncDetector_GetFuncWrap_Func(t *testing.T) {
	fw := args.FuncDetector.GetFuncWrap(sampleFunc)
	actual := args.Map{"isValid": fw.IsValid()}
	expected := args.Map{"isValid": true}
	expected.ShouldBeEqual(t, 0, "FuncDetector from func", actual)
}

func Test_FuncDetector_GetFuncWrap_FuncWrap(t *testing.T) {
	original := args.NewFuncWrap.Default(sampleFunc)
	fw := args.FuncDetector.GetFuncWrap(original)
	actual := args.Map{"same": fw == original}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "FuncDetector from FuncWrap", actual)
}

func Test_FuncDetector_GetFuncWrap_Map(t *testing.T) {
	m := args.Map{"func": sampleFunc}
	fw := args.FuncDetector.GetFuncWrap(m)
	actual := args.Map{"isValid": fw.IsValid()}
	expected := args.Map{"isValid": true}
	expected.ShouldBeEqual(t, 0, "FuncDetector from Map", actual)
}
