package coreargstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Map basic methods ──

func Test_Cov_Map_Length(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Map Length", actual)
}

func Test_Cov_Map_ArgsCount(t *testing.T) {
	m := args.Map{"a": 1, "func": func() {}, "expected": "x"}
	actual := args.Map{"argsCount": m.ArgsCount()}
	expected := args.Map{"argsCount": 1}
	expected.ShouldBeEqual(t, 0, "Map ArgsCount", actual)
}

func Test_Cov_Map_Has_HasDefined(t *testing.T) {
	m := args.Map{"a": 1, "b": nil}
	actual := args.Map{
		"hasA":        m.Has("a"),
		"hasC":        m.Has("c"),
		"definedA":    m.HasDefined("a"),
		"definedB":    m.HasDefined("b"),
		"definedC":    m.HasDefined("c"),
		"nilHas":      args.Map(nil).Has("a"),
		"nilDefined":  args.Map(nil).HasDefined("a"),
	}
	expected := args.Map{
		"hasA": true, "hasC": false,
		"definedA": true, "definedB": false, "definedC": false,
		"nilHas": false, "nilDefined": false,
	}
	expected.ShouldBeEqual(t, 0, "Map Has/HasDefined", actual)
}

func Test_Cov_Map_HasDefinedAll(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	actual := args.Map{
		"allDef":     m.HasDefinedAll("a", "b"),
		"oneMissing": m.HasDefinedAll("a", "c"),
		"nil":        args.Map(nil).HasDefinedAll("a"),
		"empty":      m.HasDefinedAll(),
	}
	expected := args.Map{"allDef": true, "oneMissing": false, "nil": false, "empty": false}
	expected.ShouldBeEqual(t, 0, "Map HasDefinedAll", actual)
}

func Test_Cov_Map_IsKeyInvalid_IsKeyMissing(t *testing.T) {
	m := args.Map{"a": 1, "b": nil}
	actual := args.Map{
		"invalidA":   m.IsKeyInvalid("a"),
		"invalidB":   m.IsKeyInvalid("b"),
		"invalidC":   m.IsKeyInvalid("c"),
		"missingA":   m.IsKeyMissing("a"),
		"missingC":   m.IsKeyMissing("c"),
		"nilInvalid": args.Map(nil).IsKeyInvalid("a"),
		"nilMissing": args.Map(nil).IsKeyMissing("a"),
	}
	expected := args.Map{
		"invalidA": false, "invalidB": true, "invalidC": true,
		"missingA": false, "missingC": true,
		"nilInvalid": false, "nilMissing": false,
	}
	expected.ShouldBeEqual(t, 0, "Map IsKeyInvalid/IsKeyMissing", actual)
}

func Test_Cov_Map_Get(t *testing.T) {
	m := args.Map{"a": 1}
	val, ok := m.Get("a")
	_, notOk := m.Get("missing")
	nilVal, nilOk := args.Map(nil).Get("a")
	actual := args.Map{"val": val, "ok": ok, "notOk": notOk, "nilVal": nilVal == nil, "nilOk": nilOk}
	expected := args.Map{"val": 1, "ok": true, "notOk": false, "nilVal": true, "nilOk": false}
	expected.ShouldBeEqual(t, 0, "Map Get", actual)
}

func Test_Cov_Map_GetLowerCase(t *testing.T) {
	m := args.Map{"hello": "world"}
	val, ok := m.GetLowerCase("HELLO")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": "world", "ok": true}
	expected.ShouldBeEqual(t, 0, "Map GetLowerCase", actual)
}

func Test_Cov_Map_GetDirectLower(t *testing.T) {
	m := args.Map{"hello": "world"}
	actual := args.Map{
		"found":   m.GetDirectLower("HELLO"),
		"missing": m.GetDirectLower("OTHER") == nil,
	}
	expected := args.Map{"found": "world", "missing": true}
	expected.ShouldBeEqual(t, 0, "Map GetDirectLower", actual)
}

func Test_Cov_Map_GetAsInt(t *testing.T) {
	m := args.Map{"a": 42, "b": "not int"}
	v, ok := m.GetAsInt("a")
	_, notOk := m.GetAsInt("b")
	_, missingOk := m.GetAsInt("c")
	actual := args.Map{"v": v, "ok": ok, "notOk": notOk, "missingOk": missingOk}
	expected := args.Map{"v": 42, "ok": true, "notOk": false, "missingOk": false}
	expected.ShouldBeEqual(t, 0, "Map GetAsInt", actual)
}

func Test_Cov_Map_GetAsIntDefault(t *testing.T) {
	m := args.Map{"a": 42}
	actual := args.Map{
		"found":   m.GetAsIntDefault("a", 0),
		"missing": m.GetAsIntDefault("b", 99),
	}
	expected := args.Map{"found": 42, "missing": 99}
	expected.ShouldBeEqual(t, 0, "Map GetAsIntDefault", actual)
}

func Test_Cov_Map_GetAsString(t *testing.T) {
	m := args.Map{"a": "hello"}
	v, ok := m.GetAsString("a")
	_, notOk := m.GetAsString("b")
	actual := args.Map{"v": v, "ok": ok, "notOk": notOk}
	expected := args.Map{"v": "hello", "ok": true, "notOk": false}
	expected.ShouldBeEqual(t, 0, "Map GetAsString", actual)
}

func Test_Cov_Map_GetAsStringDefault(t *testing.T) {
	m := args.Map{"a": "hello"}
	actual := args.Map{
		"found":   m.GetAsStringDefault("a"),
		"missing": m.GetAsStringDefault("b"),
	}
	expected := args.Map{"found": "hello", "missing": ""}
	expected.ShouldBeEqual(t, 0, "Map GetAsStringDefault", actual)
}

func Test_Cov_Map_GetAsBool(t *testing.T) {
	m := args.Map{"a": true}
	v, ok := m.GetAsBool("a")
	_, notOk := m.GetAsBool("b")
	actual := args.Map{"v": v, "ok": ok, "notOk": notOk}
	expected := args.Map{"v": true, "ok": true, "notOk": false}
	expected.ShouldBeEqual(t, 0, "Map GetAsBool", actual)
}

func Test_Cov_Map_NamedAccessors(t *testing.T) {
	m := args.Map{
		"first": 1, "second": 2, "third": 3,
		"fourth": 4, "fifth": 5, "sixth": 6, "seventh": 7,
		"when": "w", "title": "t",
	}
	actual := args.Map{
		"first": m.FirstItem(), "second": m.SecondItem(), "third": m.ThirdItem(),
		"fourth": m.FourthItem(), "fifth": m.FifthItem(), "sixth": m.SixthItem(),
		"seventh": m.Seventh(), "when": m.When(), "title": m.Title(),
		"hasFirst": m.HasFirst(),
	}
	expected := args.Map{
		"first": 1, "second": 2, "third": 3,
		"fourth": 4, "fifth": 5, "sixth": 6,
		"seventh": 7, "when": "w", "title": "t",
		"hasFirst": true,
	}
	expected.ShouldBeEqual(t, 0, "Map NamedAccessors", actual)
}

func Test_Cov_Map_AltKeyAccessors(t *testing.T) {
	m := args.Map{"f1": "a", "p2": "b", "p3": "c", "f4": "d", "f5": "e", "f6": "f", "f7": "g"}
	actual := args.Map{
		"first": m.FirstItem(), "second": m.SecondItem(), "third": m.ThirdItem(),
		"fourth": m.FourthItem(), "fifth": m.FifthItem(), "sixth": m.SixthItem(),
		"seventh": m.Seventh(),
	}
	expected := args.Map{
		"first": "a", "second": "b", "third": "c",
		"fourth": "d", "fifth": "e", "sixth": "f",
		"seventh": "g",
	}
	expected.ShouldBeEqual(t, 0, "Map AltKeyAccessors", actual)
}

func Test_Cov_Map_Expect_Actual_Arrange(t *testing.T) {
	m := args.Map{"expect": "e", "actual": "a", "arrange": "r"}
	actual := args.Map{
		"expect":  m.Expect(),
		"actual":  m.Actual(),
		"arrange": m.Arrange(),
	}
	expected := args.Map{"expect": "e", "actual": "a", "arrange": "r"}
	expected.ShouldBeEqual(t, 0, "Map Expect/Actual/Arrange", actual)
}

func Test_Cov_Map_SetActual(t *testing.T) {
	m := args.Map{}
	m.SetActual("val")
	actual := args.Map{"actual": m["actual"]}
	expected := args.Map{"actual": "val"}
	expected.ShouldBeEqual(t, 0, "Map SetActual", actual)
}

func Test_Cov_Map_SortedKeys(t *testing.T) {
	m := args.Map{"b": 2, "a": 1}
	keys, err := m.SortedKeys()
	emptyKeys, emptyErr := args.Map{}.SortedKeys()
	actual := args.Map{
		"first": keys[0], "second": keys[1], "noErr": err == nil,
		"emptyLen": len(emptyKeys), "emptyNoErr": emptyErr == nil,
	}
	expected := args.Map{
		"first": "a", "second": "b", "noErr": true,
		"emptyLen": 0, "emptyNoErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Map SortedKeys", actual)
}

func Test_Cov_Map_SortedKeysMust(t *testing.T) {
	m := args.Map{"b": 2, "a": 1}
	keys := m.SortedKeysMust()
	actual := args.Map{"first": keys[0]}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "Map SortedKeysMust", actual)
}

func Test_Cov_Map_GetByIndex(t *testing.T) {
	m := args.Map{"a": 1}
	actual := args.Map{
		"found":   m.GetByIndex(0) != nil,
		"outRange": m.GetByIndex(99) == nil,
	}
	expected := args.Map{"found": true, "outRange": true}
	expected.ShouldBeEqual(t, 0, "Map GetByIndex", actual)
}

func Test_Cov_Map_GetFirstOfNames(t *testing.T) {
	m := args.Map{"b": 2}
	actual := args.Map{
		"found":   m.GetFirstOfNames("a", "b"),
		"empty":   m.GetFirstOfNames() == nil,
		"missing": m.GetFirstOfNames("x", "y") == nil,
	}
	expected := args.Map{"found": 2, "empty": true, "missing": true}
	expected.ShouldBeEqual(t, 0, "Map GetFirstOfNames", actual)
}

func Test_Cov_Map_Raw(t *testing.T) {
	m := args.Map{"a": 1}
	raw := m.Raw()
	actual := args.Map{"len": len(raw)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Map Raw", actual)
}

func Test_Cov_Map_Args(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	result := m.Args("a", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Map Args", actual)
}

func Test_Cov_Map_ValidArgs(t *testing.T) {
	m := args.Map{"a": 1, "b": nil, "func": func() {}}
	result := m.ValidArgs()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Map ValidArgs", actual)
}

func Test_Cov_Map_Expected(t *testing.T) {
	m1 := args.Map{"expected": "a"}
	m2 := args.Map{"expects": "b"}
	m3 := args.Map{"expect": "c"}
	actual := args.Map{
		"expected": m1.Expected(), "expects": m2.Expected(), "expect": m3.Expected(),
		"hasExpect": m1.HasExpect(),
	}
	expected := args.Map{"expected": "a", "expects": "b", "expect": "c", "hasExpect": true}
	expected.ShouldBeEqual(t, 0, "Map Expected", actual)
}

func Test_Cov_Map_WorkFunc_HasFunc(t *testing.T) {
	fn := func() {}
	m := args.Map{"func": fn}
	actual := args.Map{
		"hasFunc":  m.HasFunc(),
		"noFunc":   args.Map{}.HasFunc(),
		"funcName": m.GetFuncName() != "",
	}
	expected := args.Map{"hasFunc": true, "noFunc": false, "funcName": true}
	expected.ShouldBeEqual(t, 0, "Map WorkFunc/HasFunc", actual)
}

func Test_Cov_Map_Compile(t *testing.T) {
	m := args.Map{"a": 1, "b": "hello"}
	result := m.CompileToStrings()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Map CompileToStrings", actual)
}

func Test_Cov_Map_Slice_String(t *testing.T) {
	m := args.Map{"a": 1}
	actual := args.Map{
		"sliceLen":  len(m.Slice()) > 0,
		"strNotEmpty": m.String() != "",
	}
	expected := args.Map{"sliceLen": true, "strNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "Map Slice/String", actual)
}

func Test_Cov_Map_GoLiteralString(t *testing.T) {
	m := args.Map{"a": 1}
	actual := args.Map{"notEmpty": m.GoLiteralString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Map GoLiteralString", actual)
}

// ── Dynamic ──

func Test_Cov_Dynamic_Methods(t *testing.T) {
	d := &args.DynamicAny{
		Params: args.Map{"first": 1, "second": 2},
		Expect: "expected",
	}
	actual := args.Map{
		"argsCount":  d.ArgsCount(),
		"expected":   d.Expected(),
		"hasExpect":  d.HasExpect(),
		"hasFirst":   d.HasFirst(),
		"firstItem":  d.FirstItem(),
		"secondItem": d.SecondItem(),
		"hasDefined": d.HasDefined("first"),
		"has":        d.Has("first"),
	}
	expected := args.Map{
		"argsCount": 2, "expected": "expected", "hasExpect": true,
		"hasFirst": true, "firstItem": 1, "secondItem": 2,
		"hasDefined": true, "has": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic Methods", actual)
}

func Test_Cov_Dynamic_Nil(t *testing.T) {
	var d *args.DynamicAny
	actual := args.Map{
		"argsCount":  d.ArgsCount(),
		"hasFirst":   d.HasFirst(),
		"hasDefined": d.HasDefined("a"),
		"has":        d.Has("a"),
		"hasExpect":  d.HasExpect(),
		"workFunc":   d.GetWorkFunc() == nil,
	}
	expected := args.Map{
		"argsCount": 0, "hasFirst": false, "hasDefined": false,
		"has": false, "hasExpect": false, "workFunc": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic Nil", actual)
}

func Test_Cov_Dynamic_Get_Methods(t *testing.T) {
	d := &args.DynamicAny{Params: args.Map{"num": 42, "str": "hello", "actual": "act", "arrange": "arr"}}
	v, ok := d.Get("num")
	intV, intOk := d.GetAsInt("num")
	strV, strOk := d.GetAsString("str")
	actual := args.Map{
		"v": v, "ok": ok, "intV": intV, "intOk": intOk,
		"strV": strV, "strOk": strOk,
		"strDefault": d.GetAsStringDefault("str"),
		"intDefault": d.GetAsIntDefault("num", 0),
		"intMissing": d.GetAsIntDefault("x", 99),
		"actual":     d.Actual(),
		"arrange":    d.Arrange(),
	}
	expected := args.Map{
		"v": 42, "ok": true, "intV": 42, "intOk": true,
		"strV": "hello", "strOk": true,
		"strDefault": "hello", "intDefault": 42, "intMissing": 99,
		"actual": "act", "arrange": "arr",
	}
	expected.ShouldBeEqual(t, 0, "Dynamic Get Methods", actual)
}

func Test_Cov_Dynamic_HasDefinedAll_IsKeyInvalid_IsKeyMissing(t *testing.T) {
	d := &args.DynamicAny{Params: args.Map{"a": 1, "b": nil}}
	actual := args.Map{
		"allDef":     d.HasDefinedAll("a"),
		"allMissing": d.HasDefinedAll("a", "c"),
		"invalidB":   d.IsKeyInvalid("b"),
		"missingC":   d.IsKeyMissing("c"),
	}
	expected := args.Map{"allDef": true, "allMissing": false, "invalidB": true, "missingC": true}
	expected.ShouldBeEqual(t, 0, "Dynamic HasDefinedAll/IsKeyInvalid/IsKeyMissing", actual)
}

func Test_Cov_Dynamic_GetAsStrings_GetAsAnyItems(t *testing.T) {
	d := &args.DynamicAny{Params: args.Map{"strs": []string{"a"}, "anys": []any{1}}}
	strs, sOk := d.GetAsStrings("strs")
	anys, aOk := d.GetAsAnyItems("anys")
	_, sMiss := d.GetAsStrings("x")
	_, aMiss := d.GetAsAnyItems("x")
	actual := args.Map{
		"strsLen": len(strs), "sOk": sOk, "anysLen": len(anys), "aOk": aOk,
		"sMiss": sMiss, "aMiss": aMiss,
	}
	expected := args.Map{
		"strsLen": 1, "sOk": true, "anysLen": 1, "aOk": true,
		"sMiss": false, "aMiss": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic GetAsStrings/GetAsAnyItems", actual)
}

func Test_Cov_Dynamic_Slice_String_Contracts(t *testing.T) {
	d := &args.DynamicAny{Params: args.Map{"a": 1}, Expect: "e"}
	actual := args.Map{
		"sliceLen":  len(d.Slice()) > 0,
		"strNotEmpty": d.String() != "",
		"mapper":    d.AsArgsMapper() != nil,
		"funcName":  d.AsArgFuncNameContractsBinder() != nil,
		"base":      d.AsArgBaseContractsBinder() != nil,
	}
	expected := args.Map{
		"sliceLen": true, "strNotEmpty": true,
		"mapper": true, "funcName": true, "base": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic Slice/String/Contracts", actual)
}

// ── DynamicFunc ──

func Test_Cov_DynamicFunc_Methods(t *testing.T) {
	fn := func(s string) int { return len(s) }
	d := &args.DynamicFuncAny{
		Params:   args.Map{"first": "hello"},
		WorkFunc: fn,
		Expect:   5,
	}
	actual := args.Map{
		"argsCount":  d.ArgsCount(),
		"length":     d.Length(),
		"hasFirst":   d.HasFirst(),
		"firstItem":  d.FirstItem(),
		"expected":   d.Expected(),
		"workFunc":   d.GetWorkFunc() != nil,
		"hasDefined": d.HasDefined("first"),
		"has":        d.Has("first"),
	}
	expected := args.Map{
		"argsCount": 1, "length": 1, "hasFirst": true,
		"firstItem": "hello", "expected": 5,
		"workFunc": true, "hasDefined": true, "has": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc Methods", actual)
}

func Test_Cov_DynamicFunc_Nil(t *testing.T) {
	var d *args.DynamicFuncAny
	actual := args.Map{"argsCount": d.ArgsCount(), "length": d.Length()}
	expected := args.Map{"argsCount": 0, "length": 0}
	expected.ShouldBeEqual(t, 0, "DynamicFunc Nil", actual)
}

func Test_Cov_DynamicFunc_Contracts(t *testing.T) {
	d := args.DynamicFuncAny{Params: args.Map{"a": 1}, WorkFunc: func() {}, Expect: "e"}
	actual := args.Map{
		"mapper":   d.AsArgsMapper() != nil,
		"funcName": d.AsArgFuncNameContractsBinder() != nil,
		"base":     d.AsArgBaseContractsBinder() != nil,
	}
	expected := args.Map{"mapper": true, "funcName": true, "base": true}
	expected.ShouldBeEqual(t, 0, "DynamicFunc Contracts", actual)
}

// ── FuncWrap ──

func Test_Cov_FuncWrap_Valid(t *testing.T) {
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Default(fn)
	actual := args.Map{
		"hasValid":    fw.HasValidFunc(),
		"isInvalid":   fw.IsInvalid(),
		"isValid":     fw.IsValid(),
		"name":        fw.GetFuncName() != "",
		"pascal":      fw.GetPascalCaseFuncName() != "",
		"argsCount":   fw.ArgsCount(),
		"returnLen":   fw.ReturnLength(),
		"outCount":    fw.OutArgsCount(),
		"typeNotNil":  fw.GetType() != nil,
		"pkgPath":     fw.PkgPath() != "",
		"pkgNameOnly": fw.PkgNameOnly() != "",
		"directName":  fw.FuncDirectInvokeName() != "",
	}
	expected := args.Map{
		"hasValid": true, "isInvalid": false, "isValid": true,
		"name": true, "pascal": true, "argsCount": 1, "returnLen": 1,
		"outCount": 1, "typeNotNil": true,
		"pkgPath": true, "pkgNameOnly": true, "directName": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap Valid", actual)
}

func Test_Cov_FuncWrap_Invalid(t *testing.T) {
	fw := args.NewFuncWrap.Default(nil)
	actual := args.Map{
		"isInvalid":  fw.IsInvalid(),
		"argsCount":  fw.ArgsCount(),
		"returnLen":  fw.ReturnLength(),
		"typeIsNil":  fw.GetType() == nil,
		"pkgPath":    fw.PkgPath(),
		"pkgName":    fw.PkgNameOnly(),
		"directName": fw.FuncDirectInvokeName(),
	}
	expected := args.Map{
		"isInvalid": true, "argsCount": -1, "returnLen": -1,
		"typeIsNil": true, "pkgPath": "", "pkgName": "", "directName": "",
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap Invalid", actual)
}

func Test_Cov_FuncWrap_Nil(t *testing.T) {
	var fw *args.FuncWrapAny
	actual := args.Map{
		"name":   fw.GetFuncName(),
		"pascal": fw.GetPascalCaseFuncName(),
	}
	expected := args.Map{"name": "", "pascal": ""}
	expected.ShouldBeEqual(t, 0, "FuncWrap Nil", actual)
}

func Test_Cov_FuncWrap_IsEqual(t *testing.T) {
	fn := func(s string) int { return len(s) }
	fw1 := args.NewFuncWrap.Default(fn)
	fw2 := args.NewFuncWrap.Default(fn)
	fw3 := args.NewFuncWrap.Default(func() {})
	var nilFw *args.FuncWrapAny
	actual := args.Map{
		"equalSame":    fw1.IsEqual(fw2),
		"notEqual":     fw1.IsNotEqual(fw3),
		"nilNil":       nilFw.IsEqual(nilFw),
		"nilNonNil":    nilFw.IsEqual(fw1),
		"equalVal":     fw1.IsEqualValue(*fw2),
		"isPublic":     fw1.IsPublicMethod(),
		"isPrivate":    fw1.IsPrivateMethod(),
	}
	expected := args.Map{
		"equalSame": true, "notEqual": true, "nilNil": true,
		"nilNonNil": false, "equalVal": true,
		"isPublic": true, "isPrivate": false,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap IsEqual", actual)
}

func Test_Cov_FuncWrap_ArgsTypes(t *testing.T) {
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Default(fn)
	actual := args.Map{
		"inLen":      len(fw.GetInArgsTypes()),
		"outLen":     len(fw.GetOutArgsTypes()),
		"inNames":    len(fw.GetInArgsTypesNames()) > 0,
		"outNames":   len(fw.GetOutArgsTypesNames()) > 0,
		"inArgNames": len(fw.InArgNames()) > 0,
		"outArgNames": len(fw.OutArgNames()) > 0,
	}
	expected := args.Map{
		"inLen": 1, "outLen": 1, "inNames": true,
		"outNames": true, "inArgNames": true, "outArgNames": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap ArgsTypes", actual)
}

func Test_Cov_FuncWrap_Invoke(t *testing.T) {
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Default(fn)
	results, err := fw.Invoke("hello")
	actual := args.Map{
		"noErr":  err == nil,
		"result": results[0],
	}
	expected := args.Map{"noErr": true, "result": 5}
	expected.ShouldBeEqual(t, 0, "FuncWrap Invoke", actual)
}

func Test_Cov_FuncWrap_InvokeMultiArg(t *testing.T) {
	fn := func(a, b string) string { return a + b }
	fw := args.NewFuncWrap.Default(fn)
	results, err := fw.Invoke("hello", "world")
	actual := args.Map{"noErr": err == nil, "result": results[0]}
	expected := args.Map{"noErr": true, "result": "helloworld"}
	expected.ShouldBeEqual(t, 0, "FuncWrap Invoke multi-arg", actual)
}

func Test_Cov_FuncWrap_Validate(t *testing.T) {
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Default(fn)
	err := fw.ValidateMethodArgs([]any{"hello"})
	errMismatch := fw.ValidateMethodArgs([]any{"a", "b"})
	actual := args.Map{"noErr": err == nil, "hasErr": errMismatch != nil}
	expected := args.Map{"noErr": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap Validate", actual)
}

func Test_Cov_FuncWrap_VerifyArgs(t *testing.T) {
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Default(fn)
	inOk, inErr := fw.VerifyInArgs([]any{"hello"})
	outOk, outErr := fw.VerifyOutArgs([]any{5})
	actual := args.Map{
		"inOk": inOk, "inErr": inErr == nil,
		"outOk": outOk, "outErr": outErr == nil,
	}
	expected := args.Map{"inOk": true, "inErr": true, "outOk": true, "outErr": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap VerifyArgs", actual)
}

func Test_Cov_FuncWrap_InArgNamesEachLine(t *testing.T) {
	fn := func(a, b string) int { return len(a) }
	fw := args.NewFuncWrap.Default(fn)
	inLines := fw.InArgNamesEachLine()
	outLines := fw.OutArgNamesEachLine()
	actual := args.Map{"inLen": len(inLines) > 0, "outLen": len(outLines) > 0}
	expected := args.Map{"inLen": true, "outLen": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap InArgNamesEachLine", actual)
}

func Test_Cov_FuncWrap_InOutArgsMap(t *testing.T) {
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Default(fn)
	inMap := fw.InArgsMap()
	outMap := fw.OutArgsMap()
	actual := args.Map{"inLen": len(inMap) > 0, "outLen": len(outMap) > 0}
	expected := args.Map{"inLen": true, "outLen": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap InOutArgsMap", actual)
}

// ── NewTypedFuncWrap ──

func Test_Cov_NewTypedFuncWrap_Valid(t *testing.T) {
	fn := func(s string) int { return len(s) }
	fw := args.NewTypedFuncWrap(fn)
	actual := args.Map{"isValid": fw.IsValid(), "name": fw.GetFuncName() != ""}
	expected := args.Map{"isValid": true, "name": true}
	expected.ShouldBeEqual(t, 0, "NewTypedFuncWrap Valid", actual)
}

func Test_Cov_NewTypedFuncWrap_NonFunc(t *testing.T) {
	fw := args.NewTypedFuncWrap(42)
	actual := args.Map{"isInvalid": fw.IsInvalid()}
	expected := args.Map{"isInvalid": true}
	expected.ShouldBeEqual(t, 0, "NewTypedFuncWrap NonFunc", actual)
}

// ── Holder ──

func Test_Cov_Holder(t *testing.T) {
	h := &args.HolderAny{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, Hashmap: map[string]any{"k": "v"}}
	actual := args.Map{
		"first": h.First, "second": h.Second, "third": h.Third,
		"fourth": h.Fourth, "fifth": h.Fifth, "hmLen": len(h.Hashmap),
	}
	expected := args.Map{
		"first": 1, "second": 2, "third": 3,
		"fourth": 4, "fifth": 5, "hmLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "Holder", actual)
}

// ── LeftRight ──

func Test_Cov_LeftRight(t *testing.T) {
	lr := args.LeftRight{Left: "l", Right: "r"}
	actual := args.Map{"left": lr.Left, "right": lr.Right}
	expected := args.Map{"left": "l", "right": "r"}
	expected.ShouldBeEqual(t, 0, "LeftRight", actual)
}

// ── One through Six ──

func Test_Cov_One(t *testing.T) {
	o := args.One{First: 1}
	actual := args.Map{"first": o.FirstItem(), "str": o.String() != ""}
	expected := args.Map{"first": 1, "str": true}
	expected.ShouldBeEqual(t, 0, "One", actual)
}

func Test_Cov_Two(t *testing.T) {
	o := args.Two{First: 1, Second: 2}
	actual := args.Map{"first": o.FirstItem(), "second": o.SecondItem(), "str": o.String() != ""}
	expected := args.Map{"first": 1, "second": 2, "str": true}
	expected.ShouldBeEqual(t, 0, "Two", actual)
}

func Test_Cov_Three(t *testing.T) {
	o := args.Three{First: 1, Second: 2, Third: 3}
	actual := args.Map{"first": o.FirstItem(), "second": o.SecondItem(), "third": o.ThirdItem(), "str": o.String() != ""}
	expected := args.Map{"first": 1, "second": 2, "third": 3, "str": true}
	expected.ShouldBeEqual(t, 0, "Three", actual)
}

func Test_Cov_Four(t *testing.T) {
	o := args.Four{First: 1, Second: 2, Third: 3, Fourth: 4}
	actual := args.Map{"fourth": o.FourthItem(), "str": o.String() != ""}
	expected := args.Map{"fourth": 4, "str": true}
	expected.ShouldBeEqual(t, 0, "Four", actual)
}

func Test_Cov_Five(t *testing.T) {
	o := args.Five{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5}
	actual := args.Map{"fifth": o.FifthItem(), "str": o.String() != ""}
	expected := args.Map{"fifth": 5, "str": true}
	expected.ShouldBeEqual(t, 0, "Five", actual)
}

func Test_Cov_Six(t *testing.T) {
	o := args.Six{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, Sixth: 6}
	actual := args.Map{"sixth": o.SixthItem(), "str": o.String() != ""}
	expected := args.Map{"sixth": 6, "str": true}
	expected.ShouldBeEqual(t, 0, "Six", actual)
}

// ── OneFunc through SixFunc ──

func Test_Cov_OneFunc(t *testing.T) {
	fn := func() {}
	o := args.OneFunc{First: 1, Func: fn}
	actual := args.Map{"first": o.FirstItem(), "hasFunc": o.GetWorkFunc() != nil}
	expected := args.Map{"first": 1, "hasFunc": true}
	expected.ShouldBeEqual(t, 0, "OneFunc", actual)
}

func Test_Cov_TwoFunc(t *testing.T) {
	fn := func() {}
	o := args.TwoFunc{First: 1, Second: 2, Func: fn}
	actual := args.Map{"second": o.SecondItem(), "hasFunc": o.GetWorkFunc() != nil}
	expected := args.Map{"second": 2, "hasFunc": true}
	expected.ShouldBeEqual(t, 0, "TwoFunc", actual)
}

func Test_Cov_FuncMap(t *testing.T) {
	fn := func() {}
	fm := args.FuncMap{Params: args.Map{"a": 1}, Func: fn}
	actual := args.Map{"hasFunc": fm.GetWorkFunc() != nil, "firstNotNil": fm.Params != nil}
	expected := args.Map{"hasFunc": true, "firstNotNil": true}
	expected.ShouldBeEqual(t, 0, "FuncMap", actual)
}

// ── Empty creator ──

func Test_Cov_Empty_Map(t *testing.T) {
	m := args.Empty.Map()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty Map", actual)
}

func Test_Cov_Map_GetAsStringSliceFirstOfNames(t *testing.T) {
	m := args.Map{"lines": []string{"a", "b"}}
	result := m.GetAsStringSliceFirstOfNames("lines")
	nilResult := m.GetAsStringSliceFirstOfNames("x")
	emptyResult := m.GetAsStringSliceFirstOfNames()
	actual := args.Map{
		"len": len(result), "nilResult": nilResult == nil, "emptyResult": emptyResult == nil,
	}
	expected := args.Map{"len": 2, "nilResult": true, "emptyResult": true}
	expected.ShouldBeEqual(t, 0, "Map GetAsStringSliceFirstOfNames", actual)
}

func Test_Cov_Map_WorkFuncName(t *testing.T) {
	fn := func() {}
	m := args.Map{"func": fn}
	actual := args.Map{"notEmpty": m.WorkFuncName() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Map WorkFuncName", actual)
}

func Test_Cov_Map_GetFirstFuncNameOf(t *testing.T) {
	fn := func() {}
	m := args.Map{"workFunc": fn}
	actual := args.Map{"notEmpty": m.GetFirstFuncNameOf("workFunc") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Map GetFirstFuncNameOf", actual)
}
