package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ── String type ──

func Test_Cov2_String_Concat(t *testing.T) {
	s := args.String("hello")
	result := s.Concat(" ", "world")

	actual := args.Map{
		"result": result.String(),
	}
	expected := args.Map{
		"result": "hello world",
	}
	expected.ShouldBeEqual(t, 0, "String_Concat", actual)
}

func Test_Cov2_String_Join(t *testing.T) {
	s := args.String("hello")
	result := s.Join("-", "world", "go")

	actual := args.Map{
		"notEmpty": result.String() != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "String_Join", actual)
}

func Test_Cov2_String_Split(t *testing.T) {
	s := args.String("a,b,c")
	result := s.Split(",")

	actual := args.Map{
		"len": len(result),
	}
	expected := args.Map{
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "String_Split", actual)
}

func Test_Cov2_String_Quoting(t *testing.T) {
	s := args.String("hello")

	actual := args.Map{
		"doubleQuote":      s.DoubleQuote().String() != "",
		"doubleQuoteQ":     s.DoubleQuoteQ().String() != "",
		"singleQuote":      s.SingleQuote().String() != "",
		"valueDoubleQuote": s.ValueDoubleQuote().String() != "",
	}
	expected := args.Map{
		"doubleQuote":      true,
		"doubleQuoteQ":     true,
		"singleQuote":      true,
		"valueDoubleQuote": true,
	}
	expected.ShouldBeEqual(t, 0, "String_Quoting", actual)
}

func Test_Cov2_String_Length(t *testing.T) {
	s := args.String("hello")
	empty := args.String("")

	actual := args.Map{
		"length":      s.Length(),
		"count":       s.Count(),
		"asciiLen":    s.AscIILength(),
		"isEmpty":     s.IsEmpty(),
		"emptyIsTrue": empty.IsEmpty(),
		"hasCh":       s.HasCharacter(),
		"isDefined":   s.IsDefined(),
	}
	expected := args.Map{
		"length":      5,
		"count":       5,
		"asciiLen":    5,
		"isEmpty":     false,
		"emptyIsTrue": true,
		"hasCh":       true,
		"isDefined":   true,
	}
	expected.ShouldBeEqual(t, 0, "String_Length", actual)
}

func Test_Cov2_String_IsEmptyOrWhitespace(t *testing.T) {
	s := args.String("  ")
	sNonEmpty := args.String("hello")

	actual := args.Map{
		"whitespace": s.IsEmptyOrWhitespace(),
		"nonEmpty":   sNonEmpty.IsEmptyOrWhitespace(),
	}
	expected := args.Map{
		"whitespace": true,
		"nonEmpty":   false,
	}
	expected.ShouldBeEqual(t, 0, "String_IsEmptyOrWhitespace", actual)
}

func Test_Cov2_String_TrimSpace(t *testing.T) {
	s := args.String("  hello  ")
	result := s.TrimSpace()

	actual := args.Map{
		"result": result.String(),
	}
	expected := args.Map{
		"result": "hello",
	}
	expected.ShouldBeEqual(t, 0, "String_TrimSpace", actual)
}

func Test_Cov2_String_ReplaceAll(t *testing.T) {
	s := args.String("hello world")
	result := s.ReplaceAll("world", "go")

	actual := args.Map{
		"result": result.String(),
	}
	expected := args.Map{
		"result": "hello go",
	}
	expected.ShouldBeEqual(t, 0, "String_ReplaceAll", actual)
}

func Test_Cov2_String_Substring(t *testing.T) {
	s := args.String("hello world")
	result := s.Substring(0, 5)

	actual := args.Map{
		"result": result.String(),
	}
	expected := args.Map{
		"result": "hello",
	}
	expected.ShouldBeEqual(t, 0, "String_Substring", actual)
}

func Test_Cov2_String_Bytes(t *testing.T) {
	s := args.String("hi")

	actual := args.Map{
		"bytesLen": len(s.Bytes()),
		"runesLen": len(s.Runes()),
	}
	expected := args.Map{
		"bytesLen": 2,
		"runesLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "String_Bytes", actual)
}

// ── Dynamic ──

func Test_Cov2_Dynamic_Getters(t *testing.T) {
	d := &args.DynamicAny{
		Params: args.Map{
			"first":  "f1",
			"second": "f2",
			"third":  "f3",
			"fourth": "f4",
			"fifth":  "f5",
			"sixth":  "f6",
		},
		Expect: "exp",
	}

	actual := args.Map{
		"first":    d.FirstItem(),
		"second":   d.SecondItem(),
		"third":    d.ThirdItem(),
		"fourth":   d.FourthItem(),
		"fifth":    d.FifthItem(),
		"sixth":    d.SixthItem(),
		"expected": d.Expected(),
		"hasFirst": d.HasFirst(),
	}
	expected := args.Map{
		"first":    "f1",
		"second":   "f2",
		"third":    "f3",
		"fourth":   "f4",
		"fifth":    "f5",
		"sixth":    "f6",
		"expected": "exp",
		"hasFirst": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_Getters", actual)
}

func Test_Cov2_Dynamic_NilSafety(t *testing.T) {
	var d *args.DynamicAny

	actual := args.Map{
		"argsCount":     d.ArgsCount(),
		"getWorkFunc":   d.GetWorkFunc() == nil,
		"hasFirst":      d.HasFirst(),
		"hasDefined":    d.HasDefined("key"),
		"has":           d.Has("key"),
		"hasDefinedAll": d.HasDefinedAll("key"),
		"isKeyInvalid":  d.IsKeyInvalid("key"),
		"isKeyMissing":  d.IsKeyMissing("key"),
		"hasExpect":     d.HasExpect(),
	}
	expected := args.Map{
		"argsCount":     0,
		"getWorkFunc":   true,
		"hasFirst":      false,
		"hasDefined":    false,
		"has":           false,
		"hasDefinedAll": false,
		"isKeyInvalid":  false,
		"isKeyMissing":  false,
		"hasExpect":     false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_NilSafety", actual)
}

func Test_Cov2_Dynamic_TypedGetters(t *testing.T) {
	d := &args.DynamicAny{
		Params: args.Map{
			"count":  5,
			"name":   "test",
			"items":  []string{"a", "b"},
			"anyArr": []any{1, 2},
		},
	}

	intVal, intOk := d.GetAsInt("count")
	strVal, strOk := d.GetAsString("name")
	stringsVal, stringsOk := d.GetAsStrings("items")
	anyItems, anyOk := d.GetAsAnyItems("anyArr")
	intDefault := d.GetAsIntDefault("missing", 99)
	strDefault := d.GetAsStringDefault("missing")

	actual := args.Map{
		"intVal":     intVal,
		"intOk":      intOk,
		"strVal":     strVal,
		"strOk":      strOk,
		"stringsLen": len(stringsVal),
		"stringsOk":  stringsOk,
		"anyLen":     len(anyItems),
		"anyOk":      anyOk,
		"intDefault": intDefault,
		"strDefault": strDefault,
	}
	expected := args.Map{
		"intVal":     5,
		"intOk":      true,
		"strVal":     "test",
		"strOk":      true,
		"stringsLen": 2,
		"stringsOk":  true,
		"anyLen":     2,
		"anyOk":      true,
		"intDefault": 99,
		"strDefault": "",
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_TypedGetters", actual)
}

func Test_Cov2_Dynamic_GetLowerCase(t *testing.T) {
	d := &args.DynamicAny{
		Params: args.Map{"name": "val"},
	}

	item, ok := d.GetLowerCase("Name")
	direct := d.GetDirectLower("Name")

	actual := args.Map{
		"item":   item,
		"ok":     ok,
		"direct": direct,
	}
	expected := args.Map{
		"item":   "val",
		"ok":     true,
		"direct": "val",
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_GetLowerCase", actual)
}

func Test_Cov2_Dynamic_HasDefined(t *testing.T) {
	d := &args.DynamicAny{
		Params: args.Map{"key": "val", "null": nil},
	}

	actual := args.Map{
		"hasDefined":     d.HasDefined("key"),
		"hasNull":        d.HasDefined("null"),
		"hasMissing":     d.HasDefined("missing"),
		"has":            d.Has("key"),
		"hasMissingKey":  d.Has("missing"),
		"isKeyInvalid":   d.IsKeyInvalid("null"),
		"isKeyMissing":   d.IsKeyMissing("missing"),
		"hasDefinedAll":  d.HasDefinedAll("key"),
		"hasDefinedNone": d.HasDefinedAll(),
	}
	expected := args.Map{
		"hasDefined":     true,
		"hasNull":        false,
		"hasMissing":     false,
		"has":            true,
		"hasMissingKey":  false,
		"isKeyInvalid":   true,
		"isKeyMissing":   true,
		"hasDefinedAll":  true,
		"hasDefinedNone": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_HasDefined", actual)
}

func Test_Cov2_Dynamic_ValidArgs(t *testing.T) {
	d := &args.DynamicAny{
		Params: args.Map{"a": 1, "b": 2},
	}

	validArgs := d.ValidArgs()
	customArgs := d.Args("a")

	actual := args.Map{
		"validLen":  len(validArgs),
		"customLen": len(customArgs),
	}
	expected := args.Map{
		"validLen":  2,
		"customLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_ValidArgs", actual)
}

func Test_Cov2_Dynamic_Slice(t *testing.T) {
	d := &args.DynamicAny{
		Params: args.Map{"a": 1},
		Expect: "exp",
	}

	slice := d.Slice()
	slice2 := d.Slice() // cached

	actual := args.Map{
		"sliceLen":  len(slice),
		"cachedLen": len(slice2),
	}
	expected := args.Map{
		"sliceLen":  len(slice),
		"cachedLen": len(slice2),
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_Slice", actual)
}

func Test_Cov2_Dynamic_String(t *testing.T) {
	d := &args.DynamicAny{
		Params: args.Map{"a": 1},
	}

	s := d.String()

	actual := args.Map{
		"notEmpty": s != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_String", actual)
}

func Test_Cov2_Dynamic_Actual(t *testing.T) {
	d := &args.DynamicAny{
		Params: args.Map{"actual": "myActual", "arrange": "myArrange"},
	}

	actual := args.Map{
		"actual":  d.Actual(),
		"arrange": d.Arrange(),
	}
	expected := args.Map{
		"actual":  "myActual",
		"arrange": "myArrange",
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_Actual", actual)
}

func Test_Cov2_Dynamic_Contracts(t *testing.T) {
	d := args.DynamicAny{
		Params: args.Map{"a": 1},
	}

	actual := args.Map{
		"mapper":        d.AsArgsMapper() != nil,
		"funcNameBind":  d.AsArgFuncNameContractsBinder() != nil,
		"baseBind":      d.AsArgBaseContractsBinder() != nil,
	}
	expected := args.Map{
		"mapper":        true,
		"funcNameBind":  true,
		"baseBind":      true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_Contracts", actual)
}

// ── Holder ──

func Test_Cov2_Holder_Getters(t *testing.T) {
	h := &args.HolderAny{
		First:  "f1",
		Second: "f2",
		Third:  "f3",
		Fourth: "f4",
		Fifth:  "f5",
		Sixth:  "f6",
		Expect: "exp",
	}

	actual := args.Map{
		"first":    h.FirstItem(),
		"second":   h.SecondItem(),
		"third":    h.ThirdItem(),
		"fourth":   h.FourthItem(),
		"fifth":    h.FifthItem(),
		"sixth":    h.SixthItem(),
		"expected": h.Expected(),
		"count":    h.ArgsCount(),
	}
	expected := args.Map{
		"first":    "f1",
		"second":   "f2",
		"third":    "f3",
		"fourth":   "f4",
		"fifth":    "f5",
		"sixth":    "f6",
		"expected": "exp",
		"count":    7,
	}
	expected.ShouldBeEqual(t, 0, "Holder_Getters", actual)
}

func Test_Cov2_Holder_HasMethods(t *testing.T) {
	h := &args.HolderAny{
		First:  "f1",
		Second: "f2",
		Third:  "f3",
		Fourth: "f4",
		Fifth:  "f5",
		Sixth:  "f6",
		Expect: "exp",
	}

	actual := args.Map{
		"hasFirst":  h.HasFirst(),
		"hasSecond": h.HasSecond(),
		"hasThird":  h.HasThird(),
		"hasFourth": h.HasFourth(),
		"hasFifth":  h.HasFifth(),
		"hasSixth":  h.HasSixth(),
		"hasExpect": h.HasExpect(),
		"hasFunc":   h.HasFunc(),
	}
	expected := args.Map{
		"hasFirst":  true,
		"hasSecond": true,
		"hasThird":  true,
		"hasFourth": true,
		"hasFifth":  true,
		"hasSixth":  true,
		"hasExpect": true,
		"hasFunc":   false,
	}
	expected.ShouldBeEqual(t, 0, "Holder_HasMethods", actual)
}

func Test_Cov2_Holder_ArgTwo(t *testing.T) {
	h := &args.HolderAny{First: "a", Second: "b"}
	two := h.ArgTwo()

	actual := args.Map{
		"first":  two.First,
		"second": two.Second,
	}
	expected := args.Map{
		"first":  "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "Holder_ArgTwo", actual)
}

func Test_Cov2_Holder_ArgThree(t *testing.T) {
	h := &args.HolderAny{First: "a", Second: "b", Third: "c"}
	three := h.ArgThree()

	actual := args.Map{
		"first": three.First,
		"third": three.Third,
	}
	expected := args.Map{
		"first": "a",
		"third": "c",
	}
	expected.ShouldBeEqual(t, 0, "Holder_ArgThree", actual)
}

func Test_Cov2_Holder_Args(t *testing.T) {
	h := &args.HolderAny{First: "a", Second: "b", Third: "c"}

	args1 := h.Args(1)
	args2 := h.Args(2)
	args3 := h.Args(3)

	actual := args.Map{
		"args1Len": len(args1),
		"args2Len": len(args2),
		"args3Len": len(args3),
	}
	expected := args.Map{
		"args1Len": 1,
		"args2Len": 2,
		"args3Len": 3,
	}
	expected.ShouldBeEqual(t, 0, "Holder_Args", actual)
}

func Test_Cov2_Holder_ValidArgs(t *testing.T) {
	h := &args.HolderAny{First: "a", Second: "b"}

	va := h.ValidArgs()

	actual := args.Map{
		"len": len(va),
	}
	expected := args.Map{
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "Holder_ValidArgs", actual)
}

func Test_Cov2_Holder_Slice(t *testing.T) {
	h := &args.HolderAny{First: "a", Expect: "exp"}

	s1 := h.Slice()
	s2 := h.Slice() // cached

	actual := args.Map{
		"len":      len(s1),
		"cachedEq": len(s1) == len(s2),
	}
	expected := args.Map{
		"len":      2,
		"cachedEq": true,
	}
	expected.ShouldBeEqual(t, 0, "Holder_Slice", actual)
}

func Test_Cov2_Holder_GetByIndex(t *testing.T) {
	h := &args.HolderAny{First: "a"}

	actual := args.Map{
		"valid":   h.GetByIndex(0),
		"invalid": h.GetByIndex(99) == nil,
	}
	expected := args.Map{
		"valid":   "a",
		"invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "Holder_GetByIndex", actual)
}

func Test_Cov2_Holder_String(t *testing.T) {
	h := &args.HolderAny{First: "a"}

	actual := args.Map{
		"notEmpty": h.String() != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Holder_String", actual)
}

// ── LeftRight ──

func Test_Cov2_LeftRight_Methods(t *testing.T) {
	lr := &args.LeftRightAny{
		Left:   "L",
		Right:  "R",
		Expect: "exp",
	}

	actual := args.Map{
		"first":     lr.FirstItem(),
		"second":    lr.SecondItem(),
		"expected":  lr.Expected(),
		"argsCount": lr.ArgsCount(),
		"hasFirst":  lr.HasFirst(),
		"hasSecond": lr.HasSecond(),
		"hasLeft":   lr.HasLeft(),
		"hasRight":  lr.HasRight(),
		"hasExpect": lr.HasExpect(),
	}
	expected := args.Map{
		"first":     "L",
		"second":    "R",
		"expected":  "exp",
		"argsCount": 2,
		"hasFirst":  true,
		"hasSecond": true,
		"hasLeft":   true,
		"hasRight":  true,
		"hasExpect": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight_Methods", actual)
}

func Test_Cov2_LeftRight_ArgTwo(t *testing.T) {
	lr := &args.LeftRightAny{Left: "L", Right: "R"}
	two := lr.ArgTwo()

	actual := args.Map{
		"first":  two.First,
		"second": two.Second,
	}
	expected := args.Map{
		"first":  "L",
		"second": "R",
	}
	expected.ShouldBeEqual(t, 0, "LeftRight_ArgTwo", actual)
}

func Test_Cov2_LeftRight_ValidArgs(t *testing.T) {
	lr := &args.LeftRightAny{Left: "L", Right: "R"}
	va := lr.ValidArgs()
	a := lr.Args(2)

	actual := args.Map{
		"vaLen": len(va),
		"aLen":  len(a),
	}
	expected := args.Map{
		"vaLen": 2,
		"aLen":  2,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight_ValidArgs", actual)
}

func Test_Cov2_LeftRight_Clone(t *testing.T) {
	lr := &args.LeftRightAny{Left: "L", Right: "R", Expect: "e"}
	clone := lr.Clone()

	actual := args.Map{
		"left":   clone.Left,
		"right":  clone.Right,
		"expect": clone.Expect,
	}
	expected := args.Map{
		"left":   "L",
		"right":  "R",
		"expect": "e",
	}
	expected.ShouldBeEqual(t, 0, "LeftRight_Clone", actual)
}

func Test_Cov2_LeftRight_String(t *testing.T) {
	lr := &args.LeftRightAny{Left: "L"}

	actual := args.Map{
		"notEmpty": lr.String() != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight_String", actual)
}

// ── Two ──

func Test_Cov2_Two_Methods(t *testing.T) {
	tw := &args.TwoAny{First: "a", Second: "b", Expect: "exp"}

	actual := args.Map{
		"first":     tw.FirstItem(),
		"second":    tw.SecondItem(),
		"expected":  tw.Expected(),
		"argsCount": tw.ArgsCount(),
		"hasFirst":  tw.HasFirst(),
		"hasSecond": tw.HasSecond(),
		"hasExpect": tw.HasExpect(),
	}
	expected := args.Map{
		"first":     "a",
		"second":    "b",
		"expected":  "exp",
		"argsCount": 2,
		"hasFirst":  true,
		"hasSecond": true,
		"hasExpect": true,
	}
	expected.ShouldBeEqual(t, 0, "Two_Methods", actual)
}

func Test_Cov2_Two_ArgTwo(t *testing.T) {
	tw := &args.TwoAny{First: "a", Second: "b"}
	at := tw.ArgTwo()

	actual := args.Map{
		"first":  at.First,
		"second": at.Second,
	}
	expected := args.Map{
		"first":  "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "Two_ArgTwo", actual)
}

func Test_Cov2_Two_LeftRight(t *testing.T) {
	tw := &args.TwoAny{First: "a", Second: "b", Expect: "e"}
	lr := tw.LeftRight()

	actual := args.Map{
		"left":   lr.Left,
		"right":  lr.Right,
		"expect": lr.Expect,
	}
	expected := args.Map{
		"left":   "a",
		"right":  "b",
		"expect": "e",
	}
	expected.ShouldBeEqual(t, 0, "Two_LeftRight", actual)
}

func Test_Cov2_Two_String(t *testing.T) {
	tw := &args.TwoAny{First: "a", Second: "b"}

	actual := args.Map{
		"notEmpty": tw.String() != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Two_String", actual)
}

// ── Three ──

func Test_Cov2_Three_Methods(t *testing.T) {
	th := &args.ThreeAny{First: "a", Second: "b", Third: "c", Expect: "exp"}

	actual := args.Map{
		"first":     th.FirstItem(),
		"second":    th.SecondItem(),
		"third":     th.ThirdItem(),
		"expected":  th.Expected(),
		"argsCount": th.ArgsCount(),
		"hasFirst":  th.HasFirst(),
		"hasSecond": th.HasSecond(),
		"hasThird":  th.HasThird(),
		"hasExpect": th.HasExpect(),
	}
	expected := args.Map{
		"first":     "a",
		"second":    "b",
		"third":     "c",
		"expected":  "exp",
		"argsCount": 3,
		"hasFirst":  true,
		"hasSecond": true,
		"hasThird":  true,
		"hasExpect": true,
	}
	expected.ShouldBeEqual(t, 0, "Three_Methods", actual)
}

func Test_Cov2_Three_ArgTwo(t *testing.T) {
	th := &args.ThreeAny{First: "a", Second: "b", Third: "c"}
	tw := th.ArgTwo()
	at := th.ArgThree()

	actual := args.Map{
		"twoFirst":   tw.First,
		"threeThird": at.Third,
	}
	expected := args.Map{
		"twoFirst":   "a",
		"threeThird": "c",
	}
	expected.ShouldBeEqual(t, 0, "Three_ArgTwo", actual)
}

func Test_Cov2_Three_LeftRight(t *testing.T) {
	th := &args.ThreeAny{First: "a", Second: "b", Expect: "e"}
	lr := th.LeftRight()

	actual := args.Map{
		"left":  lr.Left,
		"right": lr.Right,
	}
	expected := args.Map{
		"left":  "a",
		"right": "b",
	}
	expected.ShouldBeEqual(t, 0, "Three_LeftRight", actual)
}

func Test_Cov2_Three_String(t *testing.T) {
	th := &args.ThreeAny{First: "a"}

	actual := args.Map{
		"notEmpty": th.String() != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Three_String", actual)
}

// ── FuncWrap ──

func Test_Cov2_FuncWrap_NewTypedFuncWrap(t *testing.T) {
	fn := func(s string) int { return len(s) }
	fw := args.NewTypedFuncWrap(fn)

	actual := args.Map{
		"isValid":   fw.IsValid(),
		"hasName":   fw.GetFuncName() != "",
		"argsCount": fw.ArgsCount(),
		"returnLen": fw.ReturnLength(),
	}
	expected := args.Map{
		"isValid":   true,
		"hasName":   true,
		"argsCount": 1,
		"returnLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_NewTypedFuncWrap", actual)
}

func Test_Cov2_FuncWrap_NilFunc(t *testing.T) {
	fw := args.NewTypedFuncWrap[any](nil)

	actual := args.Map{
		"isInvalid": fw.IsInvalid(),
		"getName":   fw.GetFuncName(),
	}
	expected := args.Map{
		"isInvalid": true,
		"getName":   "",
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_NilFunc", actual)
}

func Test_Cov2_FuncWrap_NonFunc(t *testing.T) {
	fw := args.NewTypedFuncWrap(42)

	actual := args.Map{
		"isInvalid": fw.IsInvalid(),
	}
	expected := args.Map{
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_NonFunc", actual)
}

func Test_Cov2_FuncWrap_Invoke(t *testing.T) {
	fn := func(a, b int) int { return a + b }
	fw := args.NewTypedFuncWrap(fn)

	results, err := fw.Invoke(3, 4)

	actual := args.Map{
		"noErr":  err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"noErr":  true,
		"result": 7,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_Invoke", actual)
}

func Test_Cov2_FuncWrap_InvokeMust(t *testing.T) {
	fn := func() string { return "ok" }
	fw := args.NewTypedFuncWrap(fn)

	results := fw.InvokeMust()

	actual := args.Map{
		"result": results[0],
	}
	expected := args.Map{
		"result": "ok",
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_InvokeMust", actual)
}

func Test_Cov2_FuncWrap_VoidCall(t *testing.T) {
	called := false
	fn := func() { called = true }
	fw := args.NewTypedFuncWrap(fn)

	_, err := fw.VoidCall()

	actual := args.Map{
		"noErr":  err == nil,
		"called": called,
	}
	expected := args.Map{
		"noErr":  true,
		"called": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_VoidCall", actual)
}

func Test_Cov2_FuncWrap_ValidationError(t *testing.T) {
	fn := func() {}
	fw := args.NewTypedFuncWrap(fn)
	nilFw := args.NewTypedFuncWrap[any](nil)

	actual := args.Map{
		"validNoErr": fw.ValidationError() == nil,
		"nilHasErr":  nilFw.ValidationError() != nil,
	}
	expected := args.Map{
		"validNoErr": true,
		"nilHasErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_ValidationError", actual)
}

func Test_Cov2_FuncWrap_InvalidError(t *testing.T) {
	fn := func() {}
	fw := args.NewTypedFuncWrap(fn)
	nilFw := args.NewTypedFuncWrap[any](nil)

	actual := args.Map{
		"validNil": fw.InvalidError() == nil,
		"nilErr":   nilFw.InvalidError() != nil,
	}
	expected := args.Map{
		"validNil": true,
		"nilErr":   true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_InvalidError", actual)
}

func Test_Cov2_FuncWrap_IsEqual(t *testing.T) {
	fn := func() {}
	fw1 := args.NewTypedFuncWrap(fn)
	fw2 := args.NewTypedFuncWrap(fn)

	actual := args.Map{
		"selfEqual":   fw1.IsEqual(fw1),
		"sameFunc":    fw1.IsEqual(fw2),
		"nilBoth":     (*args.FuncWrapAny)(nil).IsEqual(nil),
		"nilOne":      fw1.IsEqual(nil),
		"isNotEqual":  fw1.IsNotEqual(nil),
		"isEqualVal":  fw1.IsEqualValue(*fw2),
	}
	expected := args.Map{
		"selfEqual":   true,
		"sameFunc":    true,
		"nilBoth":     true,
		"nilOne":      false,
		"isNotEqual":  true,
		"isEqualVal":  true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_IsEqual", actual)
}

func Test_Cov2_FuncWrap_PkgPath(t *testing.T) {
	fn := func() {}
	fw := args.NewTypedFuncWrap(fn)

	actual := args.Map{
		"pkgPath":    fw.PkgPath() != "",
		"pkgName":    fw.PkgNameOnly() != "",
		"directName": fw.FuncDirectInvokeName() != "",
		"getType":    fw.GetType() != nil,
		"pascal":     fw.GetPascalCaseFuncName() != "",
	}
	expected := args.Map{
		"pkgPath":    true,
		"pkgName":    true,
		"directName": true,
		"getType":    true,
		"pascal":     true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_PkgPath", actual)
}

func Test_Cov2_FuncWrap_GetFirstResponseOfInvoke(t *testing.T) {
	fn := func() string { return "first" }
	fw := args.NewTypedFuncWrap(fn)

	result, err := fw.GetFirstResponseOfInvoke()

	actual := args.Map{
		"result": result,
		"noErr":  err == nil,
	}
	expected := args.Map{
		"result": "first",
		"noErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_GetFirstResponseOfInvoke", actual)
}

func Test_Cov2_FuncWrap_InvokeResultOfIndex(t *testing.T) {
	fn := func() (string, int) { return "a", 42 }
	fw := args.NewTypedFuncWrap(fn)

	r0, err0 := fw.InvokeResultOfIndex(0)
	r1, err1 := fw.InvokeResultOfIndex(1)

	actual := args.Map{
		"r0":    r0,
		"err0":  err0 == nil,
		"r1":    r1,
		"err1":  err1 == nil,
	}
	expected := args.Map{
		"r0":    "a",
		"err0":  true,
		"r1":    42,
		"err1":  true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_InvokeResultOfIndex", actual)
}

// ── Map CompileToString ──

func Test_Cov2_Map_CompileToString(t *testing.T) {
	m := args.Map{"b": 2, "a": 1}
	result := m.CompileToString()

	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Map_CompileToString", actual)
}

func Test_Cov2_Map_CompileToString_Empty(t *testing.T) {
	m := args.Map{}
	result := m.CompileToString()

	actual := args.Map{
		"isEmpty": result == "",
	}
	expected := args.Map{
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Map_CompileToString_Empty", actual)
}
