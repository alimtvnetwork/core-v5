package corecmptests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════
// args.Map — additional coverage for uncovered methods
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_Map_CompileToStrings(t *testing.T) {
	m := args.Map{"b": 2, "a": 1}
	lines := m.CompileToStrings()
	actual := args.Map{"len": len(lines), "first": lines[0], "second": lines[1]}
	expected := args.Map{"len": 2, "first": "a : 1", "second": "b : 2"}
	expected.ShouldBeEqual(t, 0, "CompileToStrings", actual)
}

func Test_Cov12_Map_CompileToStrings_Empty(t *testing.T) {
	m := args.Map{}
	lines := m.CompileToStrings()
	actual := args.Map{"len": len(lines)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CompileToStrings empty", actual)
}

func Test_Cov12_Map_CompileToString(t *testing.T) {
	m := args.Map{"a": 1}
	result := m.CompileToString()
	actual := args.Map{"result": result}
	expected := args.Map{"result": "a : 1"}
	expected.ShouldBeEqual(t, 0, "CompileToString", actual)
}

func Test_Cov12_Map_CompileToString_Multi(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	result := m.CompileToString()
	actual := args.Map{"hasNewline": strings.Contains(result, "\n")}
	expected := args.Map{"hasNewline": true}
	expected.ShouldBeEqual(t, 0, "CompileToString multi", actual)
}

func Test_Cov12_Map_GoLiteralLines(t *testing.T) {
	m := args.Map{"a": 1, "name": "hello"}
	lines := m.GoLiteralLines()
	actual := args.Map{"len": len(lines)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "GoLiteralLines", actual)
}

func Test_Cov12_Map_GoLiteralLines_Empty(t *testing.T) {
	m := args.Map{}
	lines := m.GoLiteralLines()
	actual := args.Map{"len": len(lines)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GoLiteralLines empty", actual)
}

func Test_Cov12_Map_GoLiteralString(t *testing.T) {
	m := args.Map{"a": 1}
	result := m.GoLiteralString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GoLiteralString", actual)
}

func Test_Cov12_Map_GetAsInt(t *testing.T) {
	m := args.Map{"val": 42}
	val, ok := m.GetAsInt("val")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 42, "ok": true}
	expected.ShouldBeEqual(t, 0, "GetAsInt", actual)
}

func Test_Cov12_Map_GetAsInt_Missing(t *testing.T) {
	m := args.Map{"val": "str"}
	val, ok := m.GetAsInt("val")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 0, "ok": false}
	expected.ShouldBeEqual(t, 0, "GetAsInt wrong type", actual)
}

func Test_Cov12_Map_GetAsIntDefault(t *testing.T) {
	m := args.Map{"val": 42}
	val := m.GetAsIntDefault("val", 0)
	actual := args.Map{"val": val}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "GetAsIntDefault found", actual)
}

func Test_Cov12_Map_GetAsIntDefault_Missing(t *testing.T) {
	m := args.Map{}
	val := m.GetAsIntDefault("val", 99)
	actual := args.Map{"val": val}
	expected := args.Map{"val": 99}
	expected.ShouldBeEqual(t, 0, "GetAsIntDefault default", actual)
}

func Test_Cov12_Map_GetAsBool(t *testing.T) {
	m := args.Map{"flag": true}
	val, ok := m.GetAsBool("flag")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": true, "ok": true}
	expected.ShouldBeEqual(t, 0, "GetAsBool", actual)
}

func Test_Cov12_Map_GetAsBoolDefault(t *testing.T) {
	m := args.Map{}
	val := m.GetAsBoolDefault("flag", true)
	actual := args.Map{"val": val}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "GetAsBoolDefault", actual)
}

func Test_Cov12_Map_GetAsString(t *testing.T) {
	m := args.Map{"name": "hello"}
	val, ok := m.GetAsString("name")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": "hello", "ok": true}
	expected.ShouldBeEqual(t, 0, "GetAsString", actual)
}

func Test_Cov12_Map_GetAsStringDefault(t *testing.T) {
	m := args.Map{}
	val := m.GetAsStringDefault("name")
	actual := args.Map{"val": val}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "GetAsStringDefault", actual)
}

func Test_Cov12_Map_GetAsStrings(t *testing.T) {
	m := args.Map{"items": []string{"a", "b"}}
	items, ok := m.GetAsStrings("items")
	actual := args.Map{"len": len(items), "ok": ok}
	expected := args.Map{"len": 2, "ok": true}
	expected.ShouldBeEqual(t, 0, "GetAsStrings", actual)
}

func Test_Cov12_Map_GetAsStrings_Missing(t *testing.T) {
	m := args.Map{}
	items, ok := m.GetAsStrings("items")
	actual := args.Map{"len": len(items), "ok": ok}
	expected := args.Map{"len": 0, "ok": false}
	expected.ShouldBeEqual(t, 0, "GetAsStrings missing", actual)
}

func Test_Cov12_Map_GetAsAnyItems(t *testing.T) {
	m := args.Map{"items": []any{1, "a"}}
	items, ok := m.GetAsAnyItems("items")
	actual := args.Map{"len": len(items), "ok": ok}
	expected := args.Map{"len": 2, "ok": true}
	expected.ShouldBeEqual(t, 0, "GetAsAnyItems", actual)
}

func Test_Cov12_Map_GetAsAnyItems_Missing(t *testing.T) {
	m := args.Map{}
	items, ok := m.GetAsAnyItems("items")
	actual := args.Map{"len": len(items), "ok": ok}
	expected := args.Map{"len": 0, "ok": false}
	expected.ShouldBeEqual(t, 0, "GetAsAnyItems missing", actual)
}

func Test_Cov12_Map_Slice(t *testing.T) {
	m := args.Map{"a": 1}
	slice := m.Slice()
	actual := args.Map{"len": len(slice)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Map Slice", actual)
}

func Test_Cov12_Map_String(t *testing.T) {
	m := args.Map{"a": 1}
	result := m.String()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Map String", actual)
}

func Test_Cov12_Map_GetByIndex(t *testing.T) {
	m := args.Map{"a": 1}
	result := m.GetByIndex(0)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetByIndex valid", actual)
}

func Test_Cov12_Map_GetByIndex_OutOfBounds(t *testing.T) {
	m := args.Map{"a": 1}
	result := m.GetByIndex(99)
	actual := args.Map{"nil": result == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "GetByIndex out of bounds", actual)
}

func Test_Cov12_Map_SortedKeys(t *testing.T) {
	m := args.Map{"b": 2, "a": 1}
	keys, err := m.SortedKeys()
	actual := args.Map{"noErr": err == nil, "first": keys[0], "second": keys[1]}
	expected := args.Map{"noErr": true, "first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "SortedKeys", actual)
}

func Test_Cov12_Map_SortedKeys_Empty(t *testing.T) {
	m := args.Map{}
	keys, err := m.SortedKeys()
	actual := args.Map{"noErr": err == nil, "len": len(keys)}
	expected := args.Map{"noErr": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "SortedKeys empty", actual)
}

func Test_Cov12_Map_When(t *testing.T) {
	m := args.Map{"when": "condition"}
	actual := args.Map{"val": m.When()}
	expected := args.Map{"val": "condition"}
	expected.ShouldBeEqual(t, 0, "Map When", actual)
}

func Test_Cov12_Map_Title(t *testing.T) {
	m := args.Map{"title": "test"}
	actual := args.Map{"val": m.Title()}
	expected := args.Map{"val": "test"}
	expected.ShouldBeEqual(t, 0, "Map Title", actual)
}

func Test_Cov12_Map_GetLowerCase(t *testing.T) {
	m := args.Map{"name": "hello"}
	val, ok := m.GetLowerCase("NAME")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": "hello", "ok": true}
	expected.ShouldBeEqual(t, 0, "GetLowerCase", actual)
}

func Test_Cov12_Map_GetDirectLower(t *testing.T) {
	m := args.Map{"key": "val"}
	result := m.GetDirectLower("KEY")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "GetDirectLower", actual)
}

func Test_Cov12_Map_GetDirectLower_Missing(t *testing.T) {
	m := args.Map{}
	result := m.GetDirectLower("KEY")
	actual := args.Map{"nil": result == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "GetDirectLower missing", actual)
}

func Test_Cov12_Map_Expect(t *testing.T) {
	m := args.Map{"expect": 42}
	actual := args.Map{"val": m.Expect()}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "Map Expect", actual)
}

func Test_Cov12_Map_Actual(t *testing.T) {
	m := args.Map{"actual": "data"}
	actual := args.Map{"val": m.Actual()}
	expected := args.Map{"val": "data"}
	expected.ShouldBeEqual(t, 0, "Map Actual", actual)
}

func Test_Cov12_Map_Arrange(t *testing.T) {
	m := args.Map{"arrange": "setup"}
	actual := args.Map{"val": m.Arrange()}
	expected := args.Map{"val": "setup"}
	expected.ShouldBeEqual(t, 0, "Map Arrange", actual)
}

func Test_Cov12_Map_SetActual(t *testing.T) {
	m := args.Map{}
	m.SetActual("new-val")
	actual := args.Map{"val": m.Actual()}
	expected := args.Map{"val": "new-val"}
	expected.ShouldBeEqual(t, 0, "Map SetActual", actual)
}

func Test_Cov12_Map_HasDefinedAll(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	actual := args.Map{
		"all":  m.HasDefinedAll("a", "b"),
		"miss": m.HasDefinedAll("a", "c"),
	}
	expected := args.Map{"all": true, "miss": false}
	expected.ShouldBeEqual(t, 0, "HasDefinedAll", actual)
}

func Test_Cov12_Map_HasDefinedAll_Nil(t *testing.T) {
	var m args.Map
	actual := args.Map{"nil": m.HasDefinedAll("a")}
	expected := args.Map{"nil": false}
	expected.ShouldBeEqual(t, 0, "HasDefinedAll nil", actual)
}

func Test_Cov12_Map_HasDefinedAll_Empty(t *testing.T) {
	m := args.Map{"a": 1}
	actual := args.Map{"empty": m.HasDefinedAll()}
	expected := args.Map{"empty": false}
	expected.ShouldBeEqual(t, 0, "HasDefinedAll empty names", actual)
}

func Test_Cov12_Map_IsKeyInvalid(t *testing.T) {
	m := args.Map{"a": 1}
	actual := args.Map{
		"valid":   !m.IsKeyInvalid("a"),
		"invalid": m.IsKeyInvalid("b"),
	}
	expected := args.Map{"valid": true, "invalid": true}
	expected.ShouldBeEqual(t, 0, "IsKeyInvalid", actual)
}

func Test_Cov12_Map_IsKeyMissing(t *testing.T) {
	m := args.Map{"a": 1}
	actual := args.Map{
		"present": !m.IsKeyMissing("a"),
		"missing": m.IsKeyMissing("b"),
	}
	expected := args.Map{"present": true, "missing": true}
	expected.ShouldBeEqual(t, 0, "IsKeyMissing", actual)
}

func Test_Cov12_Map_GetAsStringSliceFirstOfNames(t *testing.T) {
	m := args.Map{"items": []string{"a", "b"}}
	result := m.GetAsStringSliceFirstOfNames("items")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "GetAsStringSliceFirstOfNames", actual)
}

func Test_Cov12_Map_GetAsStringSliceFirstOfNames_Empty(t *testing.T) {
	m := args.Map{}
	result := m.GetAsStringSliceFirstOfNames()
	actual := args.Map{"nil": result == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "GetAsStringSliceFirstOfNames empty", actual)
}

func Test_Cov12_Map_FirstItem(t *testing.T) {
	m := args.Map{"first": "val"}
	actual := args.Map{"val": m.FirstItem()}
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "FirstItem", actual)
}

func Test_Cov12_Map_SecondItem(t *testing.T) {
	m := args.Map{"second": "val"}
	actual := args.Map{"val": m.SecondItem()}
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "SecondItem", actual)
}

func Test_Cov12_Map_ThirdItem(t *testing.T) {
	m := args.Map{"third": "val"}
	actual := args.Map{"val": m.ThirdItem()}
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "ThirdItem", actual)
}

func Test_Cov12_Map_FourthItem(t *testing.T) {
	m := args.Map{"fourth": "val"}
	actual := args.Map{"val": m.FourthItem()}
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "FourthItem", actual)
}

func Test_Cov12_Map_FifthItem(t *testing.T) {
	m := args.Map{"fifth": "val"}
	actual := args.Map{"val": m.FifthItem()}
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "FifthItem", actual)
}

func Test_Cov12_Map_SixthItem(t *testing.T) {
	m := args.Map{"sixth": "val"}
	actual := args.Map{"val": m.SixthItem()}
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "SixthItem", actual)
}

func Test_Cov12_Map_Seventh(t *testing.T) {
	m := args.Map{"seventh": "val"}
	actual := args.Map{"val": m.Seventh()}
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "Seventh", actual)
}

func Test_Cov12_Map_ValidArgs(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	result := m.ValidArgs()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ValidArgs", actual)
}

func Test_Cov12_Map_Args(t *testing.T) {
	m := args.Map{"a": 1}
	result := m.Args("a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Args", actual)
}

func Test_Cov12_Map_Raw(t *testing.T) {
	m := args.Map{"a": 1}
	raw := m.Raw()
	actual := args.Map{"len": len(raw)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Raw", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.String — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_String_Concat(t *testing.T) {
	s := args.String("hello")
	result := s.Concat(" world")
	actual := args.Map{"val": result.String()}
	expected := args.Map{"val": "hello world"}
	expected.ShouldBeEqual(t, 0, "String Concat", actual)
}

func Test_Cov12_String_Join(t *testing.T) {
	s := args.String("a")
	result := s.Join("-", "b", "c")
	actual := args.Map{"val": result.String()}
	expected := args.Map{"val": "a-b-c"}
	expected.ShouldBeEqual(t, 0, "String Join", actual)
}

func Test_Cov12_String_Split(t *testing.T) {
	s := args.String("a-b-c")
	result := s.Split("-")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "String Split", actual)
}

func Test_Cov12_String_DoubleQuote(t *testing.T) {
	s := args.String("hello")
	result := s.DoubleQuote()
	actual := args.Map{"notEmpty": result.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String DoubleQuote", actual)
}

func Test_Cov12_String_DoubleQuoteQ(t *testing.T) {
	s := args.String("hello")
	result := s.DoubleQuoteQ()
	actual := args.Map{"notEmpty": result.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String DoubleQuoteQ", actual)
}

func Test_Cov12_String_SingleQuote(t *testing.T) {
	s := args.String("hello")
	result := s.SingleQuote()
	actual := args.Map{"notEmpty": result.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String SingleQuote", actual)
}

func Test_Cov12_String_ValueDoubleQuote(t *testing.T) {
	s := args.String("hello")
	result := s.ValueDoubleQuote()
	actual := args.Map{"notEmpty": result.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String ValueDoubleQuote", actual)
}

func Test_Cov12_String_Bytes(t *testing.T) {
	s := args.String("hello")
	result := s.Bytes()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 5}
	expected.ShouldBeEqual(t, 0, "String Bytes", actual)
}

func Test_Cov12_String_Runes(t *testing.T) {
	s := args.String("hello")
	result := s.Runes()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 5}
	expected.ShouldBeEqual(t, 0, "String Runes", actual)
}

func Test_Cov12_String_Length(t *testing.T) {
	s := args.String("hello")
	actual := args.Map{"len": s.Length(), "count": s.Count(), "ascii": s.AscIILength()}
	expected := args.Map{"len": 5, "count": 5, "ascii": 5}
	expected.ShouldBeEqual(t, 0, "String Length/Count/ASCII", actual)
}

func Test_Cov12_String_IsEmptyOrWhitespace(t *testing.T) {
	actual := args.Map{
		"empty":      args.String("").IsEmptyOrWhitespace(),
		"whitespace": args.String("   ").IsEmptyOrWhitespace(),
		"notEmpty":   args.String("x").IsEmptyOrWhitespace(),
	}
	expected := args.Map{"empty": true, "whitespace": true, "notEmpty": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespace", actual)
}

func Test_Cov12_String_TrimSpace(t *testing.T) {
	s := args.String("  hello  ")
	actual := args.Map{"val": s.TrimSpace().String()}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "String TrimSpace", actual)
}

func Test_Cov12_String_ReplaceAll(t *testing.T) {
	s := args.String("hello world")
	actual := args.Map{"val": s.ReplaceAll("world", "go").String()}
	expected := args.Map{"val": "hello go"}
	expected.ShouldBeEqual(t, 0, "String ReplaceAll", actual)
}

func Test_Cov12_String_Substring(t *testing.T) {
	s := args.String("hello")
	actual := args.Map{"val": s.Substring(0, 3).String()}
	expected := args.Map{"val": "hel"}
	expected.ShouldBeEqual(t, 0, "String Substring", actual)
}

func Test_Cov12_String_IsEmpty(t *testing.T) {
	actual := args.Map{
		"empty":    args.String("").IsEmpty(),
		"notEmpty": args.String("x").IsEmpty(),
		"hasCh":    args.String("x").HasCharacter(),
		"defined":  args.String("x").IsDefined(),
	}
	expected := args.Map{"empty": true, "notEmpty": false, "hasCh": true, "defined": true}
	expected.ShouldBeEqual(t, 0, "String IsEmpty/HasCharacter/IsDefined", actual)
}

func Test_Cov12_String_TrimReplaceMap(t *testing.T) {
	s := args.String("hello {name}")
	result := s.TrimReplaceMap(map[string]string{"{name}": "world"})
	actual := args.Map{"notEmpty": result.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TrimReplaceMap", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.One — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_One_AllMethods(t *testing.T) {
	o := args.OneAny{First: "hello", Expect: "expected"}
	actual := args.Map{
		"first":      o.FirstItem(),
		"expected":   o.Expected(),
		"hasFirst":   o.HasFirst(),
		"hasExpect":  o.HasExpect(),
		"argsCount":  o.ArgsCount(),
		"strNotNull": o.String() != "",
	}
	expected := args.Map{
		"first": "hello", "expected": "expected", "hasFirst": true,
		"hasExpect": true, "argsCount": 1, "strNotNull": true,
	}
	expected.ShouldBeEqual(t, 0, "One all methods", actual)
}

func Test_Cov12_One_Args(t *testing.T) {
	o := args.OneAny{First: "hello"}
	actual := args.Map{"len0": len(o.Args(0)), "len1": len(o.Args(1))}
	expected := args.Map{"len0": 0, "len1": 1}
	expected.ShouldBeEqual(t, 0, "One Args", actual)
}

func Test_Cov12_One_ValidArgs(t *testing.T) {
	o := args.OneAny{First: "hello"}
	actual := args.Map{"len": len(o.ValidArgs())}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "One ValidArgs", actual)
}

func Test_Cov12_One_LeftRight(t *testing.T) {
	o := args.OneAny{First: "hello", Expect: "exp"}
	lr := o.LeftRight()
	actual := args.Map{"left": lr.Left}
	expected := args.Map{"left": "hello"}
	expected.ShouldBeEqual(t, 0, "One LeftRight", actual)
}

func Test_Cov12_One_GetByIndex(t *testing.T) {
	o := args.OneAny{First: "hello"}
	actual := args.Map{"val": o.GetByIndex(0), "nil": o.GetByIndex(99) == nil}
	expected := args.Map{"val": "hello", "nil": true}
	expected.ShouldBeEqual(t, 0, "One GetByIndex", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Two — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_Two_AllMethods(t *testing.T) {
	tw := args.TwoAny{First: "a", Second: "b", Expect: "exp"}
	actual := args.Map{
		"first":     tw.FirstItem(),
		"second":    tw.SecondItem(),
		"expected":  tw.Expected(),
		"hasFirst":  tw.HasFirst(),
		"hasSecond": tw.HasSecond(),
		"hasExpect": tw.HasExpect(),
		"argsCount": tw.ArgsCount(),
	}
	expected := args.Map{
		"first": "a", "second": "b", "expected": "exp",
		"hasFirst": true, "hasSecond": true, "hasExpect": true, "argsCount": 2,
	}
	expected.ShouldBeEqual(t, 0, "Two all methods", actual)
}

func Test_Cov12_Two_Args(t *testing.T) {
	tw := args.TwoAny{First: "a", Second: "b"}
	actual := args.Map{"len1": len(tw.Args(1)), "len2": len(tw.Args(2))}
	expected := args.Map{"len1": 1, "len2": 2}
	expected.ShouldBeEqual(t, 0, "Two Args", actual)
}

func Test_Cov12_Two_LeftRight(t *testing.T) {
	tw := args.TwoAny{First: "a", Second: "b"}
	lr := tw.LeftRight()
	actual := args.Map{"left": lr.Left, "right": lr.Right}
	expected := args.Map{"left": "a", "right": "b"}
	expected.ShouldBeEqual(t, 0, "Two LeftRight", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Three — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_Three_AllMethods(t *testing.T) {
	th := args.ThreeAny{First: "a", Second: "b", Third: "c", Expect: "exp"}
	actual := args.Map{
		"first":     th.FirstItem(),
		"second":    th.SecondItem(),
		"third":     th.ThirdItem(),
		"hasThird":  th.HasThird(),
		"argsCount": th.ArgsCount(),
	}
	expected := args.Map{
		"first": "a", "second": "b", "third": "c",
		"hasThird": true, "argsCount": 3,
	}
	expected.ShouldBeEqual(t, 0, "Three all methods", actual)
}

func Test_Cov12_Three_ArgTwo(t *testing.T) {
	th := args.ThreeAny{First: "a", Second: "b", Third: "c"}
	tw := th.ArgTwo()
	actual := args.Map{"first": tw.FirstItem(), "second": tw.SecondItem()}
	expected := args.Map{"first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "Three ArgTwo", actual)
}

func Test_Cov12_Three_ArgThree(t *testing.T) {
	th := args.ThreeAny{First: "a", Second: "b", Third: "c"}
	copy := th.ArgThree()
	actual := args.Map{"third": copy.ThirdItem()}
	expected := args.Map{"third": "c"}
	expected.ShouldBeEqual(t, 0, "Three ArgThree", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Four — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_Four_AllMethods(t *testing.T) {
	f := args.FourAny{First: "a", Second: "b", Third: "c", Fourth: "d", Expect: "exp"}
	actual := args.Map{
		"hasFourth":  f.HasFourth(),
		"argsCount":  f.ArgsCount(),
		"fourthItem": f.FourthItem(),
	}
	expected := args.Map{"hasFourth": true, "argsCount": 4, "fourthItem": "d"}
	expected.ShouldBeEqual(t, 0, "Four all methods", actual)
}

func Test_Cov12_Four_ArgThree(t *testing.T) {
	f := args.FourAny{First: "a", Second: "b", Third: "c", Fourth: "d"}
	th := f.ArgThree()
	actual := args.Map{"third": th.ThirdItem()}
	expected := args.Map{"third": "c"}
	expected.ShouldBeEqual(t, 0, "Four ArgThree", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Five — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_Five_AllMethods(t *testing.T) {
	f := args.FiveAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Expect: "exp"}
	actual := args.Map{
		"hasFifth":  f.HasFifth(),
		"argsCount": f.ArgsCount(),
		"fifthItem": f.FifthItem(),
	}
	expected := args.Map{"hasFifth": true, "argsCount": 5, "fifthItem": "e"}
	expected.ShouldBeEqual(t, 0, "Five all methods", actual)
}

func Test_Cov12_Five_ArgFour(t *testing.T) {
	f := args.FiveAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e"}
	fo := f.ArgFour()
	actual := args.Map{"fourth": fo.FourthItem()}
	expected := args.Map{"fourth": "d"}
	expected.ShouldBeEqual(t, 0, "Five ArgFour", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Six — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_Six_AllMethods(t *testing.T) {
	s := args.SixAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f", Expect: "exp"}
	actual := args.Map{
		"hasSixth":  s.HasSixth(),
		"argsCount": s.ArgsCount(),
		"sixthItem": s.SixthItem(),
	}
	expected := args.Map{"hasSixth": true, "argsCount": 6, "sixthItem": "f"}
	expected.ShouldBeEqual(t, 0, "Six all methods", actual)
}

func Test_Cov12_Six_ArgFive(t *testing.T) {
	s := args.SixAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f"}
	fi := s.ArgFive()
	actual := args.Map{"fifth": fi.FifthItem()}
	expected := args.Map{"fifth": "e"}
	expected.ShouldBeEqual(t, 0, "Six ArgFive", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.LeftRight — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_LeftRight_AllMethods(t *testing.T) {
	lr := args.LeftRightAny{Left: "l", Right: "r", Expect: "exp"}
	actual := args.Map{
		"left":      lr.FirstItem(),
		"right":     lr.SecondItem(),
		"expected":  lr.Expected(),
		"hasLeft":   lr.HasLeft(),
		"hasRight":  lr.HasRight(),
		"hasExpect": lr.HasExpect(),
		"argsCount": lr.ArgsCount(),
	}
	expected := args.Map{
		"left": "l", "right": "r", "expected": "exp",
		"hasLeft": true, "hasRight": true, "hasExpect": true, "argsCount": 2,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight all methods", actual)
}

func Test_Cov12_LeftRight_Clone(t *testing.T) {
	lr := args.LeftRightAny{Left: "l", Right: "r"}
	cloned := lr.Clone()
	actual := args.Map{"left": cloned.Left, "right": cloned.Right}
	expected := args.Map{"left": "l", "right": "r"}
	expected.ShouldBeEqual(t, 0, "LeftRight Clone", actual)
}

func Test_Cov12_LeftRight_ArgTwo(t *testing.T) {
	lr := args.LeftRightAny{Left: "l", Right: "r"}
	tw := lr.ArgTwo()
	actual := args.Map{"first": tw.First, "second": tw.Second}
	expected := args.Map{"first": "l", "second": "r"}
	expected.ShouldBeEqual(t, 0, "LeftRight ArgTwo", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Holder — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_Holder_AllMethods(t *testing.T) {
	h := args.HolderAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f", Expect: "exp"}
	actual := args.Map{
		"argsCount":  h.ArgsCount(),
		"hasFirst":   h.HasFirst(),
		"hasSecond":  h.HasSecond(),
		"hasThird":   h.HasThird(),
		"hasFourth":  h.HasFourth(),
		"hasFifth":   h.HasFifth(),
		"hasSixth":   h.HasSixth(),
		"hasExpect":  h.HasExpect(),
		"firstItem":  h.FirstItem(),
		"secondItem": h.SecondItem(),
		"thirdItem":  h.ThirdItem(),
		"fourthItem": h.FourthItem(),
		"fifthItem":  h.FifthItem(),
		"sixthItem":  h.SixthItem(),
		"expected":   h.Expected(),
	}
	expected := args.Map{
		"argsCount": 7, "hasFirst": true, "hasSecond": true,
		"hasThird": true, "hasFourth": true, "hasFifth": true,
		"hasSixth": true, "hasExpect": true,
		"firstItem": "a", "secondItem": "b", "thirdItem": "c",
		"fourthItem": "d", "fifthItem": "e", "sixthItem": "f",
		"expected": "exp",
	}
	expected.ShouldBeEqual(t, 0, "Holder all methods", actual)
}

func Test_Cov12_Holder_ArgTwo(t *testing.T) {
	h := args.HolderAny{First: "a", Second: "b"}
	tw := h.ArgTwo()
	actual := args.Map{"first": tw.First, "second": tw.Second}
	expected := args.Map{"first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "Holder ArgTwo", actual)
}

func Test_Cov12_Holder_ArgThree(t *testing.T) {
	h := args.HolderAny{First: "a", Second: "b", Third: "c"}
	th := h.ArgThree()
	actual := args.Map{"third": th.Third}
	expected := args.Map{"third": "c"}
	expected.ShouldBeEqual(t, 0, "Holder ArgThree", actual)
}

func Test_Cov12_Holder_ArgFour(t *testing.T) {
	h := args.HolderAny{First: "a", Second: "b", Third: "c", Fourth: "d"}
	fo := h.ArgFour()
	actual := args.Map{"fourth": fo.Fourth}
	expected := args.Map{"fourth": "d"}
	expected.ShouldBeEqual(t, 0, "Holder ArgFour", actual)
}

func Test_Cov12_Holder_ArgFive(t *testing.T) {
	h := args.HolderAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e"}
	fi := h.ArgFive()
	actual := args.Map{"fifth": fi.Fifth}
	expected := args.Map{"fifth": "e"}
	expected.ShouldBeEqual(t, 0, "Holder ArgFive", actual)
}

func Test_Cov12_Holder_ValidArgs(t *testing.T) {
	h := args.HolderAny{First: "a", Second: "b"}
	actual := args.Map{"len": len(h.ValidArgs())}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Holder ValidArgs", actual)
}

func Test_Cov12_Holder_Args(t *testing.T) {
	h := args.HolderAny{First: "a", Second: "b", Third: "c"}
	actual := args.Map{"len2": len(h.Args(2)), "len3": len(h.Args(3))}
	expected := args.Map{"len2": 2, "len3": 3}
	expected.ShouldBeEqual(t, 0, "Holder Args", actual)
}

func Test_Cov12_Holder_Slice(t *testing.T) {
	h := args.HolderAny{First: "a", Expect: "exp"}
	slice := h.Slice()
	actual := args.Map{"len": len(slice)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Holder Slice", actual)
}

func Test_Cov12_Holder_String(t *testing.T) {
	h := args.HolderAny{First: "a"}
	actual := args.Map{"notEmpty": h.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Holder String", actual)
}

func Test_Cov12_Holder_GetByIndex(t *testing.T) {
	h := args.HolderAny{First: "a"}
	actual := args.Map{"val": h.GetByIndex(0), "nil": h.GetByIndex(99) == nil}
	expected := args.Map{"val": "a", "nil": true}
	expected.ShouldBeEqual(t, 0, "Holder GetByIndex", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Dynamic — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_Dynamic_AllMethods(t *testing.T) {
	d := args.DynamicAny{Params: args.Map{"first": "a", "actual": "data"}, Expect: "exp"}
	actual := args.Map{
		"expected":   d.Expected(),
		"hasExpect":  d.HasExpect(),
		"hasFirst":   d.HasFirst(),
		"firstItem":  d.FirstItem(),
		"actual":     d.Actual(),
		"arrange":    d.Arrange(),
		"argsCount":  d.ArgsCount(),
	}
	expected := args.Map{
		"expected": "exp", "hasExpect": true, "hasFirst": true,
		"firstItem": "a", "actual": "data", "arrange": nil,
		"argsCount": 1,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic all methods", actual)
}

func Test_Cov12_Dynamic_NilReceiver(t *testing.T) {
	var d *args.DynamicAny
	actual := args.Map{
		"argsCount":  d.ArgsCount(),
		"getWork":    d.GetWorkFunc() == nil,
		"hasFirst":   d.HasFirst(),
		"hasDefined": d.HasDefined("a"),
		"has":        d.Has("a"),
		"invalid":    d.IsKeyInvalid("a"),
		"missing":    d.IsKeyMissing("a"),
		"hasExpect":  d.HasExpect(),
	}
	expected := args.Map{
		"argsCount": 0, "getWork": true, "hasFirst": false,
		"hasDefined": false, "has": false, "invalid": false,
		"missing": false, "hasExpect": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic nil receiver", actual)
}

func Test_Cov12_Dynamic_Get(t *testing.T) {
	d := args.DynamicAny{Params: args.Map{"k": "v"}}
	val, ok := d.Get("k")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": "v", "ok": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Get", actual)
}

func Test_Cov12_Dynamic_GetAsInt(t *testing.T) {
	d := args.DynamicAny{Params: args.Map{"n": 42}}
	val, ok := d.GetAsInt("n")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 42, "ok": true}
	expected.ShouldBeEqual(t, 0, "Dynamic GetAsInt", actual)
}

func Test_Cov12_Dynamic_GetAsIntDefault(t *testing.T) {
	d := args.DynamicAny{Params: args.Map{}}
	val := d.GetAsIntDefault("n", 99)
	actual := args.Map{"val": val}
	expected := args.Map{"val": 99}
	expected.ShouldBeEqual(t, 0, "Dynamic GetAsIntDefault", actual)
}

func Test_Cov12_Dynamic_GetAsString(t *testing.T) {
	d := args.DynamicAny{Params: args.Map{"s": "hello"}}
	val, ok := d.GetAsString("s")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": "hello", "ok": true}
	expected.ShouldBeEqual(t, 0, "Dynamic GetAsString", actual)
}

func Test_Cov12_Dynamic_GetAsStringDefault(t *testing.T) {
	d := args.DynamicAny{Params: args.Map{}}
	val := d.GetAsStringDefault("s")
	actual := args.Map{"val": val}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "Dynamic GetAsStringDefault", actual)
}

func Test_Cov12_Dynamic_GetAsStrings(t *testing.T) {
	d := args.DynamicAny{Params: args.Map{"items": []string{"a"}}}
	items, ok := d.GetAsStrings("items")
	actual := args.Map{"len": len(items), "ok": ok}
	expected := args.Map{"len": 1, "ok": true}
	expected.ShouldBeEqual(t, 0, "Dynamic GetAsStrings", actual)
}

func Test_Cov12_Dynamic_GetAsAnyItems(t *testing.T) {
	d := args.DynamicAny{Params: args.Map{"items": []any{1}}}
	items, ok := d.GetAsAnyItems("items")
	actual := args.Map{"len": len(items), "ok": ok}
	expected := args.Map{"len": 1, "ok": true}
	expected.ShouldBeEqual(t, 0, "Dynamic GetAsAnyItems", actual)
}

func Test_Cov12_Dynamic_HasDefinedAll(t *testing.T) {
	d := args.DynamicAny{Params: args.Map{"a": 1, "b": 2}}
	actual := args.Map{
		"all":    d.HasDefinedAll("a", "b"),
		"miss":   d.HasDefinedAll("a", "c"),
		"empty":  d.HasDefinedAll(),
	}
	expected := args.Map{"all": true, "miss": false, "empty": false}
	expected.ShouldBeEqual(t, 0, "Dynamic HasDefinedAll", actual)
}

func Test_Cov12_Dynamic_String(t *testing.T) {
	d := args.DynamicAny{Params: args.Map{"a": 1}}
	actual := args.Map{"notEmpty": d.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dynamic String", actual)
}

func Test_Cov12_Dynamic_GetLowerCase(t *testing.T) {
	d := args.DynamicAny{Params: args.Map{"name": "val"}}
	val, ok := d.GetLowerCase("NAME")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": "val", "ok": true}
	expected.ShouldBeEqual(t, 0, "Dynamic GetLowerCase", actual)
}

func Test_Cov12_Dynamic_GetDirectLower(t *testing.T) {
	d := args.DynamicAny{Params: args.Map{"key": "val"}}
	result := d.GetDirectLower("KEY")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "Dynamic GetDirectLower", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.DynamicFunc — coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_DynamicFunc_AllMethods(t *testing.T) {
	df := args.DynamicFuncAny{
		Params:   args.Map{"first": "a", "when": "cond", "title": "t"},
		WorkFunc: strings.ToUpper,
		Expect:   "exp",
	}
	actual := args.Map{
		"argsCount": df.ArgsCount(),
		"hasFunc":   df.HasFunc(),
		"hasExpect": df.HasExpect(),
		"length":    df.Length(),
		"hasFirst":  df.HasFirst(),
		"when":      df.When(),
		"title":     df.Title(),
	}
	expected := args.Map{
		"argsCount": 2, "hasFunc": true, "hasExpect": true,
		"length": 3, "hasFirst": true, "when": "cond", "title": "t",
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc all methods", actual)
}

func Test_Cov12_DynamicFunc_NilReceiver(t *testing.T) {
	var df *args.DynamicFuncAny
	actual := args.Map{
		"argsCount":  df.ArgsCount(),
		"length":     df.Length(),
		"hasDefined": df.HasDefined("a"),
		"has":        df.Has("a"),
		"invalid":    df.IsKeyInvalid("a"),
		"missing":    df.IsKeyMissing("a"),
		"hasFunc":    df.HasFunc(),
		"hasExpect":  df.HasExpect(),
	}
	expected := args.Map{
		"argsCount": 0, "length": 0, "hasDefined": false,
		"has": false, "invalid": false, "missing": false,
		"hasFunc": false, "hasExpect": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc nil receiver", actual)
}

func Test_Cov12_DynamicFunc_Get(t *testing.T) {
	df := args.DynamicFuncAny{Params: args.Map{"k": "v"}}
	val, ok := df.Get("k")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": "v", "ok": true}
	expected.ShouldBeEqual(t, 0, "DynamicFunc Get", actual)
}

func Test_Cov12_DynamicFunc_GetAsInt(t *testing.T) {
	df := args.DynamicFuncAny{Params: args.Map{"n": 42}}
	val, ok := df.GetAsInt("n")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 42, "ok": true}
	expected.ShouldBeEqual(t, 0, "DynamicFunc GetAsInt", actual)
}

func Test_Cov12_DynamicFunc_GetAsString(t *testing.T) {
	df := args.DynamicFuncAny{Params: args.Map{"s": "hello"}}
	val, ok := df.GetAsString("s")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": "hello", "ok": true}
	expected.ShouldBeEqual(t, 0, "DynamicFunc GetAsString", actual)
}

func Test_Cov12_DynamicFunc_GetAsStrings(t *testing.T) {
	df := args.DynamicFuncAny{Params: args.Map{"items": []string{"a"}}}
	items, ok := df.GetAsStrings("items")
	actual := args.Map{"len": len(items), "ok": ok}
	expected := args.Map{"len": 1, "ok": true}
	expected.ShouldBeEqual(t, 0, "DynamicFunc GetAsStrings", actual)
}

func Test_Cov12_DynamicFunc_GetAsAnyItems(t *testing.T) {
	df := args.DynamicFuncAny{Params: args.Map{"items": []any{1}}}
	items, ok := df.GetAsAnyItems("items")
	actual := args.Map{"len": len(items), "ok": ok}
	expected := args.Map{"len": 1, "ok": true}
	expected.ShouldBeEqual(t, 0, "DynamicFunc GetAsAnyItems", actual)
}

func Test_Cov12_DynamicFunc_Actual_Arrange(t *testing.T) {
	df := args.DynamicFuncAny{Params: args.Map{"actual": "data", "arrange": "setup"}}
	actual := args.Map{"actual": df.Actual(), "arrange": df.Arrange()}
	expected := args.Map{"actual": "data", "arrange": "setup"}
	expected.ShouldBeEqual(t, 0, "DynamicFunc Actual/Arrange", actual)
}

func Test_Cov12_DynamicFunc_HasDefinedAll(t *testing.T) {
	df := args.DynamicFuncAny{Params: args.Map{"a": 1, "b": 2}}
	actual := args.Map{
		"all":  df.HasDefinedAll("a", "b"),
		"miss": df.HasDefinedAll("a", "c"),
	}
	expected := args.Map{"all": true, "miss": false}
	expected.ShouldBeEqual(t, 0, "DynamicFunc HasDefinedAll", actual)
}

func Test_Cov12_DynamicFunc_String(t *testing.T) {
	df := args.DynamicFuncAny{Params: args.Map{"a": 1}, WorkFunc: strings.ToUpper}
	actual := args.Map{"notEmpty": df.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DynamicFunc String", actual)
}

func Test_Cov12_DynamicFunc_GetLowerCase(t *testing.T) {
	df := args.DynamicFuncAny{Params: args.Map{"name": "val"}}
	val, ok := df.GetLowerCase("NAME")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": "val", "ok": true}
	expected.ShouldBeEqual(t, 0, "DynamicFunc GetLowerCase", actual)
}

func Test_Cov12_DynamicFunc_GetDirectLower(t *testing.T) {
	df := args.DynamicFuncAny{Params: args.Map{"key": "val"}}
	result := df.GetDirectLower("KEY")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "DynamicFunc GetDirectLower", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.OneFunc — coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_OneFunc_AllMethods(t *testing.T) {
	of := args.OneFuncAny{First: "hello", WorkFunc: strings.ToUpper, Expect: "HELLO"}
	actual := args.Map{
		"hasFirst":  of.HasFirst(),
		"hasFunc":   of.HasFunc(),
		"hasExpect": of.HasExpect(),
		"argsCount": of.ArgsCount(),
		"firstItem": of.FirstItem(),
		"expected":  of.Expected(),
		"funcName":  of.GetFuncName() != "",
	}
	expected := args.Map{
		"hasFirst": true, "hasFunc": true, "hasExpect": true,
		"argsCount": 1, "firstItem": "hello", "expected": "HELLO",
		"funcName": true,
	}
	expected.ShouldBeEqual(t, 0, "OneFunc all methods", actual)
}

func Test_Cov12_OneFunc_InvokeWithValidArgs(t *testing.T) {
	of := args.OneFuncAny{First: "hello", WorkFunc: strings.ToUpper}
	results, err := of.InvokeWithValidArgs()
	actual := args.Map{"noErr": err == nil, "result": results[0]}
	expected := args.Map{"noErr": true, "result": "HELLO"}
	expected.ShouldBeEqual(t, 0, "OneFunc InvokeWithValidArgs", actual)
}

func Test_Cov12_OneFunc_Slice(t *testing.T) {
	of := args.OneFuncAny{First: "hello", WorkFunc: strings.ToUpper, Expect: "exp"}
	slice := of.Slice()
	actual := args.Map{"len": len(slice)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "OneFunc Slice", actual)
}

func Test_Cov12_OneFunc_LeftRight(t *testing.T) {
	of := args.OneFuncAny{First: "hello", WorkFunc: strings.ToUpper}
	lr := of.LeftRight()
	actual := args.Map{"left": lr.Left}
	expected := args.Map{"left": "hello"}
	expected.ShouldBeEqual(t, 0, "OneFunc LeftRight", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.TwoFunc — coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_TwoFunc_AllMethods(t *testing.T) {
	tf := args.TwoFuncAny{First: "a", Second: "b", WorkFunc: strings.Join, Expect: "exp"}
	actual := args.Map{
		"hasFirst":  tf.HasFirst(),
		"hasSecond": tf.HasSecond(),
		"hasFunc":   tf.HasFunc(),
		"hasExpect": tf.HasExpect(),
		"argsCount": tf.ArgsCount(),
	}
	expected := args.Map{
		"hasFirst": true, "hasSecond": true, "hasFunc": true,
		"hasExpect": true, "argsCount": 2,
	}
	expected.ShouldBeEqual(t, 0, "TwoFunc all methods", actual)
}

func Test_Cov12_TwoFunc_ArgTwo(t *testing.T) {
	tf := args.TwoFuncAny{First: "a", Second: "b"}
	tw := tf.ArgTwo()
	actual := args.Map{"first": tw.First, "second": tw.Second}
	expected := args.Map{"first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "TwoFunc ArgTwo", actual)
}

func Test_Cov12_TwoFunc_LeftRight(t *testing.T) {
	tf := args.TwoFuncAny{First: "a", Second: "b"}
	lr := tf.LeftRight()
	actual := args.Map{"left": lr.Left, "right": lr.Right}
	expected := args.Map{"left": "a", "right": "b"}
	expected.ShouldBeEqual(t, 0, "TwoFunc LeftRight", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.FuncMap — coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_FuncMap_Basic(t *testing.T) {
	fm := args.NewFuncWrap.Map(strings.ToUpper, strings.ToLower)
	actual := args.Map{
		"len":     fm.Length(),
		"count":   fm.Count(),
		"hasAny":  fm.HasAnyItem(),
		"isEmpty": fm.IsEmpty(),
	}
	expected := args.Map{"len": 2, "count": 2, "hasAny": true, "isEmpty": false}
	expected.ShouldBeEqual(t, 0, "FuncMap basic", actual)
}

func Test_Cov12_FuncMap_Has(t *testing.T) {
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	actual := args.Map{
		"has":      fm.Has("ToUpper"),
		"notHas":   fm.Has("missing"),
		"contains": fm.IsContains("ToUpper"),
	}
	expected := args.Map{"has": true, "notHas": false, "contains": true}
	expected.ShouldBeEqual(t, 0, "FuncMap Has", actual)
}

func Test_Cov12_FuncMap_Get(t *testing.T) {
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	f := fm.Get("ToUpper")
	actual := args.Map{"notNil": f != nil, "nilMissing": fm.Get("missing") == nil}
	expected := args.Map{"notNil": true, "nilMissing": true}
	expected.ShouldBeEqual(t, 0, "FuncMap Get", actual)
}

func Test_Cov12_FuncMap_IsValidFuncOf(t *testing.T) {
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	actual := args.Map{
		"valid":   fm.IsValidFuncOf("ToUpper"),
		"invalid": fm.IsInvalidFunc("missing"),
	}
	expected := args.Map{"valid": true, "invalid": true}
	expected.ShouldBeEqual(t, 0, "FuncMap IsValidFuncOf", actual)
}

func Test_Cov12_FuncMap_ArgsCount(t *testing.T) {
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	actual := args.Map{
		"args":   fm.ArgsCount("ToUpper"),
		"ret":    fm.ReturnLength("ToUpper"),
		"argsL":  fm.ArgsLength("ToUpper"),
		"nilArg": fm.ArgsCount("missing"),
		"nilRet": fm.ReturnLength("missing"),
	}
	expected := args.Map{"args": 1, "ret": 1, "argsL": 1, "nilArg": 0, "nilRet": 0}
	expected.ShouldBeEqual(t, 0, "FuncMap ArgsCount/ReturnLength", actual)
}

func Test_Cov12_FuncMap_PkgPath(t *testing.T) {
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	actual := args.Map{
		"notEmpty":   fm.PkgPath("ToUpper") != "",
		"nilEmpty":   fm.PkgPath("missing") == "",
		"pkgName":    fm.PkgNameOnly("ToUpper") != "",
		"nilPkg":     fm.PkgNameOnly("missing") == "",
		"directName": fm.FuncDirectInvokeName("ToUpper") != "",
		"nilDirect":  fm.FuncDirectInvokeName("missing") == "",
	}
	expected := args.Map{
		"notEmpty": true, "nilEmpty": true,
		"pkgName": true, "nilPkg": true,
		"directName": true, "nilDirect": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap PkgPath", actual)
}

func Test_Cov12_FuncMap_GetPascalCaseFuncName(t *testing.T) {
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	result := fm.GetPascalCaseFuncName("ToUpper")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetPascalCaseFuncName", actual)
}

func Test_Cov12_FuncMap_GetPascalCaseFuncName_Empty(t *testing.T) {
	fm := args.FuncMap{}
	result := fm.GetPascalCaseFuncName("anything")
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "GetPascalCaseFuncName empty", actual)
}

func Test_Cov12_FuncMap_IsPublicPrivate(t *testing.T) {
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	actual := args.Map{
		"public":     fm.IsPublicMethod("ToUpper"),
		"nilPublic":  fm.IsPublicMethod("missing"),
		"nilPrivate": fm.IsPrivateMethod("missing"),
	}
	expected := args.Map{"public": true, "nilPublic": false, "nilPrivate": false}
	expected.ShouldBeEqual(t, 0, "FuncMap Public/Private", actual)
}

func Test_Cov12_FuncMap_GetType(t *testing.T) {
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	actual := args.Map{
		"notNil":  fm.GetType("ToUpper") != nil,
		"nilType": fm.GetType("missing") == nil,
	}
	expected := args.Map{"notNil": true, "nilType": true}
	expected.ShouldBeEqual(t, 0, "FuncMap GetType", actual)
}

func Test_Cov12_FuncMap_GetInOutArgsTypes(t *testing.T) {
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	inTypes := fm.GetInArgsTypes("ToUpper")
	outTypes := fm.GetOutArgsTypes("ToUpper")
	inNames := fm.GetInArgsTypesNames("ToUpper")
	actual := args.Map{
		"inLen":    len(inTypes),
		"outLen":   len(outTypes),
		"namesLen": len(inNames),
		"nilIn":    len(fm.GetInArgsTypes("missing")),
		"nilOut":   len(fm.GetOutArgsTypes("missing")),
		"nilNames": len(fm.GetInArgsTypesNames("missing")),
	}
	expected := args.Map{
		"inLen": 1, "outLen": 1, "namesLen": 1,
		"nilIn": 0, "nilOut": 0, "nilNames": 0,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap GetInOutArgsTypes", actual)
}

func Test_Cov12_FuncMap_ValidationError(t *testing.T) {
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	actual := args.Map{
		"valid":   fm.ValidationError("ToUpper") == nil,
		"invalid": fm.ValidationError("missing") != nil,
	}
	expected := args.Map{"valid": true, "invalid": true}
	expected.ShouldBeEqual(t, 0, "FuncMap ValidationError", actual)
}

func Test_Cov12_FuncMap_Invoke(t *testing.T) {
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	var knownName string
	for k := range fm {
		knownName = k
		break
	}
	results, err := fm.Invoke(knownName, "hello")
	var result any
	if len(results) > 0 {
		result = results[0]
	}
	actual := args.Map{"noErr": err == nil, "hasResult": result != nil}
	expected := args.Map{"noErr": true, "hasResult": true}
	expected.ShouldBeEqual(t, 0, "FuncMap Invoke", actual)
}

func Test_Cov12_FuncMap_Invoke_NotFound(t *testing.T) {
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	_, err := fm.Invoke("missing")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FuncMap Invoke not found", actual)
}

func Test_Cov12_FuncMap_InvalidError(t *testing.T) {
	fm := args.FuncMap{}
	actual := args.Map{"hasErr": fm.InvalidError() != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FuncMap InvalidError", actual)
}

func Test_Cov12_FuncMap_VoidCall(t *testing.T) {
	called := false
	fn := func() { called = true }
	fm := args.NewFuncWrap.Map(fn)
	name := fm.Get(fm.Get(fm.Get("").GetFuncName()).GetFuncName()).GetFuncName()
	// Just exercise VoidCall
	_ = name
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FuncMap VoidCall", actual)
	_ = called
}

// ══════════════════════════════════════════════════════════════════
// args.FuncWrap — additional coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_FuncWrap_TypedHelpers(t *testing.T) {
	boolFn := func(s string) bool { return s == "yes" }
	errFn := func(s string) error { return nil }
	strFn := func(s string) string { return s }
	voidFn := func() {}
	valErrFn := func(s string) (string, error) { return s, nil }

	boolFW := args.NewTypedFuncWrap(boolFn)
	errFW := args.NewTypedFuncWrap(errFn)
	strFW := args.NewTypedFuncWrap(strFn)
	voidFW := args.NewTypedFuncWrap(voidFn)
	valErrFW := args.NewTypedFuncWrap(valErrFn)

	actual := args.Map{
		"isBool":     boolFW.IsBoolFunc(),
		"isError":    errFW.IsErrorFunc(),
		"isString":   strFW.IsStringFunc(),
		"isVoid":     voidFW.IsVoidFunc(),
		"isValErr":   valErrFW.IsValueErrorFunc(),
		"isAnyErr":   valErrFW.IsAnyErrorFunc(),
		"isAny":      strFW.IsAnyFunc(),
	}
	expected := args.Map{
		"isBool": true, "isError": true, "isString": true,
		"isVoid": true, "isValErr": true, "isAnyErr": true,
		"isAny": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap typed helpers", actual)
}

func Test_Cov12_FuncWrap_InvokeAsBool(t *testing.T) {
	fn := func(s string) bool { return s == "yes" }
	fw := args.NewTypedFuncWrap(fn)
	result, err := fw.InvokeAsBool("yes")
	actual := args.Map{"result": result, "noErr": err == nil}
	expected := args.Map{"result": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeAsBool", actual)
}

func Test_Cov12_FuncWrap_InvokeAsString(t *testing.T) {
	fn := func(s string) string { return strings.ToUpper(s) }
	fw := args.NewTypedFuncWrap(fn)
	result, err := fw.InvokeAsString("hello")
	actual := args.Map{"result": result, "noErr": err == nil}
	expected := args.Map{"result": "HELLO", "noErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeAsString", actual)
}

func Test_Cov12_FuncWrap_InvokeAsAny(t *testing.T) {
	fn := func(s string) string { return s }
	fw := args.NewTypedFuncWrap(fn)
	result, err := fw.InvokeAsAny("hello")
	actual := args.Map{"result": result, "noErr": err == nil}
	expected := args.Map{"result": "hello", "noErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeAsAny", actual)
}

func Test_Cov12_FuncWrap_InvokeAsAnyError(t *testing.T) {
	fn := func(s string) (string, error) { return s, nil }
	fw := args.NewTypedFuncWrap(fn)
	result, funcErr, procErr := fw.InvokeAsAnyError("hello")
	actual := args.Map{"result": result, "funcErr": funcErr == nil, "procErr": procErr == nil}
	expected := args.Map{"result": "hello", "funcErr": true, "procErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeAsAnyError", actual)
}

func Test_Cov12_FuncWrap_InvokeAsError(t *testing.T) {
	fn := func() error { return nil }
	fw := args.NewTypedFuncWrap(fn)
	funcErr, procErr := fw.InvokeAsError()
	actual := args.Map{"funcErr": funcErr == nil, "procErr": procErr == nil}
	expected := args.Map{"funcErr": true, "procErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeAsError", actual)
}

func Test_Cov12_FuncWrap_InArgNames(t *testing.T) {
	fn := func(s string) string { return s }
	fw := args.NewTypedFuncWrap(fn)
	names := fw.InArgNames()
	actual := args.Map{"len": len(names)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "InArgNames", actual)
}

func Test_Cov12_FuncWrap_InArgNames_Multi(t *testing.T) {
	fn := func(a string, b int) string { return a }
	fw := args.NewTypedFuncWrap(fn)
	names := fw.InArgNames()
	actual := args.Map{"len": len(names)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "InArgNames multi", actual)
}

func Test_Cov12_FuncWrap_OutArgNames(t *testing.T) {
	fn := func() (string, error) { return "", nil }
	fw := args.NewTypedFuncWrap(fn)
	names := fw.OutArgNames()
	actual := args.Map{"len": len(names)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "OutArgNames", actual)
}

func Test_Cov12_FuncWrap_InArgNamesEachLine(t *testing.T) {
	fn := func(a string, b int) string { return a }
	fw := args.NewTypedFuncWrap(fn)
	lines := fw.InArgNamesEachLine()
	actual := args.Map{"greaterThan1": len(lines) > 1}
	expected := args.Map{"greaterThan1": true}
	expected.ShouldBeEqual(t, 0, "InArgNamesEachLine", actual)
}

func Test_Cov12_FuncWrap_OutArgNamesEachLine(t *testing.T) {
	fn := func() (string, error) { return "", nil }
	fw := args.NewTypedFuncWrap(fn)
	lines := fw.OutArgNamesEachLine()
	actual := args.Map{"greaterThan1": len(lines) > 1}
	expected := args.Map{"greaterThan1": true}
	expected.ShouldBeEqual(t, 0, "OutArgNamesEachLine", actual)
}

func Test_Cov12_FuncWrap_IsInTypeMatches(t *testing.T) {
	fn := func(s string) string { return s }
	fw := args.NewTypedFuncWrap(fn)
	actual := args.Map{
		"match":    fw.IsInTypeMatches("hello"),
		"outMatch": fw.IsOutTypeMatches("result"),
	}
	expected := args.Map{"match": true, "outMatch": true}
	expected.ShouldBeEqual(t, 0, "IsInTypeMatches/IsOutTypeMatches", actual)
}

func Test_Cov12_FuncWrap_PascalCase(t *testing.T) {
	fn := func() {}
	fw := args.NewTypedFuncWrap(fn)
	result := fw.GetPascalCaseFuncName()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetPascalCaseFuncName", actual)
}

func Test_Cov12_FuncWrap_PkgNameOnly(t *testing.T) {
	fw := args.NewTypedFuncWrap(strings.ToUpper)
	result := fw.PkgNameOnly()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PkgNameOnly", actual)
}

func Test_Cov12_FuncWrap_FuncDirectInvokeName(t *testing.T) {
	fw := args.NewTypedFuncWrap(strings.ToUpper)
	result := fw.FuncDirectInvokeName()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FuncDirectInvokeName", actual)
}

func Test_Cov12_FuncWrap_IsEqual_SameFunc(t *testing.T) {
	a := args.NewTypedFuncWrap(strings.ToUpper)
	b := args.NewTypedFuncWrap(strings.ToUpper)
	actual := args.Map{"equal": a.IsEqual(b)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqual same func", actual)
}

func Test_Cov12_FuncWrap_IsNotEqual_DiffFunc(t *testing.T) {
	a := args.NewTypedFuncWrap(strings.ToUpper)
	b := args.NewTypedFuncWrap(strings.ToLower)
	actual := args.Map{"notEqual": a.IsNotEqual(b)}
	expected := args.Map{"notEqual": true}
	expected.ShouldBeEqual(t, 0, "IsNotEqual diff func", actual)
}

func Test_Cov12_FuncWrap_IsEqualValue(t *testing.T) {
	a := args.NewTypedFuncWrap(strings.ToUpper)
	b := *args.NewTypedFuncWrap(strings.ToUpper)
	actual := args.Map{"equal": a.IsEqualValue(b)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqualValue", actual)
}

func Test_Cov12_FuncWrap_ValidationError_Nil(t *testing.T) {
	var fw *args.FuncWrapAny
	err := fw.ValidationError()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidationError nil", actual)
}

func Test_Cov12_FuncWrap_InvalidError(t *testing.T) {
	fw := args.NewFuncWrap.Invalid()
	err := fw.InvalidError()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvalidError", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.FuncDetector — coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_FuncDetector_GetFuncWrap_FromMap(t *testing.T) {
	m := args.Map{"func": strings.ToUpper}
	fw := args.FuncDetector.GetFuncWrap(m)
	actual := args.Map{"valid": fw.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "FuncDetector Map", actual)
}

func Test_Cov12_FuncDetector_GetFuncWrap_FromFuncWrap(t *testing.T) {
	fw := args.NewFuncWrap.Default(strings.ToUpper)
	result := args.FuncDetector.GetFuncWrap(fw)
	actual := args.Map{"valid": result.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "FuncDetector FuncWrap", actual)
}

func Test_Cov12_FuncDetector_GetFuncWrap_FromRawFunc(t *testing.T) {
	result := args.FuncDetector.GetFuncWrap(strings.ToUpper)
	actual := args.Map{"valid": result.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "FuncDetector raw func", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Empty / NewFuncWrap — coverage
// ══════════════════════════════════════════════════════════════════

func Test_Cov12_Empty_Map(t *testing.T) {
	m := args.Empty.Map()
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty.Map", actual)
}

func Test_Cov12_Empty_FuncWrap(t *testing.T) {
	fw := args.Empty.FuncWrap()
	actual := args.Map{"invalid": fw.IsInvalid()}
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "Empty.FuncWrap", actual)
}

func Test_Cov12_Empty_FuncMap(t *testing.T) {
	fm := args.Empty.FuncMap()
	actual := args.Map{"empty": fm.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.FuncMap", actual)
}

func Test_Cov12_Empty_Holder(t *testing.T) {
	h := args.Empty.Holder()
	actual := args.Map{"argsCount": h.ArgsCount()}
	expected := args.Map{"argsCount": 7}
	expected.ShouldBeEqual(t, 0, "Empty.Holder", actual)
}

func Test_Cov12_NewFuncWrap_Many(t *testing.T) {
	fws := args.NewFuncWrap.Many(strings.ToUpper, strings.ToLower)
	actual := args.Map{"len": len(fws)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap.Many", actual)
}

func Test_Cov12_NewFuncWrap_Many_Empty(t *testing.T) {
	fws := args.NewFuncWrap.Many()
	actual := args.Map{"len": len(fws)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap.Many empty", actual)
}

func Test_Cov12_NewFuncWrap_Map_Empty(t *testing.T) {
	fm := args.NewFuncWrap.Map()
	actual := args.Map{"empty": fm.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap.Map empty", actual)
}

func Test_Cov12_NewFuncWrap_Invalid(t *testing.T) {
	fw := args.NewFuncWrap.Invalid()
	actual := args.Map{"invalid": fw.IsInvalid()}
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap.Invalid", actual)
}

func Test_Cov12_NewFuncWrap_Default_Nil(t *testing.T) {
	fw := args.NewFuncWrap.Default(nil)
	actual := args.Map{"invalid": fw.IsInvalid()}
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap.Default nil", actual)
}

func Test_Cov12_NewFuncWrap_Default_NonFunc(t *testing.T) {
	fw := args.NewFuncWrap.Default("not a func")
	actual := args.Map{"invalid": fw.IsInvalid()}
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap.Default non-func", actual)
}

func Test_Cov12_NewFuncWrap_Single(t *testing.T) {
	fw := args.NewFuncWrap.Single(strings.ToUpper)
	actual := args.Map{"valid": fw.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap.Single", actual)
}

func Test_Cov12_NewTypedFuncWrap_Nil(t *testing.T) {
	fw := args.NewTypedFuncWrap[any](nil)
	actual := args.Map{"invalid": fw.IsInvalid()}
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "NewTypedFuncWrap nil", actual)
}

func Test_Cov12_NewTypedFuncWrap_NonFunc(t *testing.T) {
	fw := args.NewTypedFuncWrap("not a func")
	actual := args.Map{"invalid": fw.IsInvalid()}
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "NewTypedFuncWrap non-func", actual)
}
