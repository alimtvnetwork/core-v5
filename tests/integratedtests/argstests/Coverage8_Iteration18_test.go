package argstests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// --- Map ---

func Test_I18_Map_Length(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	if m.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_I18_Map_ArgsCount(t *testing.T) {
	m := args.Map{"first": 1, "second": 2}
	if m.ArgsCount() != 2 {
		t.Fatal("expected 2 args")
	}
}

func Test_I18_Map_ArgsCount_WithFuncAndExpect(t *testing.T) {
	fn := func() {}
	m := args.Map{"first": 1, "func": fn, "expected": "x"}
	count := m.ArgsCount()
	if count != 1 {
		t.Fatalf("expected 1 arg (excluding func and expect), got %d", count)
	}
}

func Test_I18_Map_HasFirst(t *testing.T) {
	m := args.Map{"first": "hello"}
	if !m.HasFirst() {
		t.Fatal("expected HasFirst")
	}

	m2 := args.Map{}
	if m2.HasFirst() {
		t.Fatal("expected no first")
	}
}

func Test_I18_Map_Items(t *testing.T) {
	m := args.Map{
		"first":  "a",
		"second": "b",
		"third":  "c",
		"fourth": "d",
		"fifth":  "e",
		"sixth":  "f",
	}
	if m.FirstItem() != "a" {
		t.Fatal("first mismatch")
	}
	if m.SecondItem() != "b" {
		t.Fatal("second mismatch")
	}
	if m.ThirdItem() != "c" {
		t.Fatal("third mismatch")
	}
	if m.FourthItem() != "d" {
		t.Fatal("fourth mismatch")
	}
	if m.FifthItem() != "e" {
		t.Fatal("fifth mismatch")
	}
	if m.SixthItem() != "f" {
		t.Fatal("sixth mismatch")
	}
}

func Test_I18_Map_Seventh(t *testing.T) {
	m := args.Map{"seventh": "g"}
	if m.Seventh() != "g" {
		t.Fatal("seventh mismatch")
	}
}

func Test_I18_Map_Expected(t *testing.T) {
	m := args.Map{"expected": "x"}
	if m.Expected() != "x" {
		t.Fatal("expected mismatch")
	}
}

func Test_I18_Map_HasExpect(t *testing.T) {
	m := args.Map{"expected": "x"}
	if !m.HasExpect() {
		t.Fatal("expected HasExpect")
	}
}

func Test_I18_Map_GetByIndex(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	item := m.GetByIndex(0)
	if item == nil {
		t.Fatal("expected item at index 0")
	}

	nilItem := m.GetByIndex(100)
	if nilItem != nil {
		t.Fatal("expected nil for out of range")
	}
}

func Test_I18_Map_HasFunc(t *testing.T) {
	fn := func() {}
	m := args.Map{"func": fn}
	if !m.HasFunc() {
		t.Fatal("expected HasFunc")
	}
}

func Test_I18_Map_GetFuncName(t *testing.T) {
	m := args.Map{}
	name := m.GetFuncName()
	if name != "" {
		t.Fatal("expected empty func name")
	}
}

func Test_I18_Map_HasDefined(t *testing.T) {
	m := args.Map{"key": "val"}
	if !m.HasDefined("key") {
		t.Fatal("expected defined")
	}
	if m.HasDefined("missing") {
		t.Fatal("expected not defined")
	}

	var nilMap args.Map
	if nilMap.HasDefined("key") {
		t.Fatal("expected false for nil map")
	}
}

func Test_I18_Map_Has(t *testing.T) {
	m := args.Map{"key": "val"}
	if !m.Has("key") {
		t.Fatal("expected has")
	}
	if m.Has("missing") {
		t.Fatal("expected not has")
	}
	var nilMap args.Map
	if nilMap.Has("key") {
		t.Fatal("expected false for nil")
	}
}

func Test_I18_Map_HasDefinedAll(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	if !m.HasDefinedAll("a", "b") {
		t.Fatal("expected all defined")
	}
	if m.HasDefinedAll("a", "missing") {
		t.Fatal("expected false for missing key")
	}

	var nilMap args.Map
	if nilMap.HasDefinedAll("a") {
		t.Fatal("expected false for nil")
	}
}

func Test_I18_Map_IsKeyInvalid(t *testing.T) {
	m := args.Map{"key": "val"}
	if m.IsKeyInvalid("key") {
		t.Fatal("expected valid key")
	}
	if !m.IsKeyInvalid("missing") {
		t.Fatal("expected invalid for missing")
	}
}

func Test_I18_Map_IsKeyMissing(t *testing.T) {
	m := args.Map{"key": "val"}
	if m.IsKeyMissing("key") {
		t.Fatal("expected not missing")
	}
	if !m.IsKeyMissing("missing") {
		t.Fatal("expected missing")
	}
}

func Test_I18_Map_SortedKeys(t *testing.T) {
	m := args.Map{"b": 2, "a": 1}
	keys, err := m.SortedKeys()
	if err != nil || len(keys) != 2 || keys[0] != "a" {
		t.Fatal("unexpected sorted keys")
	}
}

func Test_I18_Map_SortedKeys_Empty(t *testing.T) {
	m := args.Map{}
	keys, err := m.SortedKeys()
	if err != nil || len(keys) != 0 {
		t.Fatal("expected empty sorted keys")
	}
}

func Test_I18_Map_When(t *testing.T) {
	m := args.Map{"when": "now"}
	if m.When() != "now" {
		t.Fatal("expected when=now")
	}
}

func Test_I18_Map_Title(t *testing.T) {
	m := args.Map{"title": "test"}
	if m.Title() != "test" {
		t.Fatal("expected title=test")
	}
}

func Test_I18_Map_Get(t *testing.T) {
	m := args.Map{"key": "val"}
	item, valid := m.Get("key")
	if !valid || item != "val" {
		t.Fatal("unexpected")
	}

	_, valid = m.Get("missing")
	if valid {
		t.Fatal("expected invalid for missing")
	}

	var nilMap args.Map
	_, valid = nilMap.Get("key")
	if valid {
		t.Fatal("expected invalid for nil")
	}
}

func Test_I18_Map_GetLowerCase(t *testing.T) {
	m := args.Map{"key": "val"}
	item, valid := m.GetLowerCase("KEY")
	if !valid || item != "val" {
		t.Fatal("expected lower case match")
	}
}

func Test_I18_Map_GetDirectLower(t *testing.T) {
	m := args.Map{"key": "val"}
	item := m.GetDirectLower("KEY")
	if item != "val" {
		t.Fatal("expected lower case match")
	}

	nilItem := m.GetDirectLower("MISSING")
	if nilItem != nil {
		t.Fatal("expected nil for missing")
	}
}

func Test_I18_Map_Expect(t *testing.T) {
	m := args.Map{"expect": "x"}
	if m.Expect() != "x" {
		t.Fatal("expected x")
	}
}

func Test_I18_Map_Actual(t *testing.T) {
	m := args.Map{"actual": "y"}
	if m.Actual() != "y" {
		t.Fatal("expected y")
	}
}

func Test_I18_Map_Arrange(t *testing.T) {
	m := args.Map{"arrange": "z"}
	if m.Arrange() != "z" {
		t.Fatal("expected z")
	}
}

func Test_I18_Map_SetActual(t *testing.T) {
	m := args.Map{}
	m.SetActual("val")
	if m.Actual() != "val" {
		t.Fatal("expected set actual")
	}
}

func Test_I18_Map_WorkFunc(t *testing.T) {
	fn := func() {}
	m := args.Map{"func": fn}
	wf := m.WorkFunc()
	if wf == nil {
		t.Fatal("expected non-nil work func")
	}
}

func Test_I18_Map_GetWorkFunc(t *testing.T) {
	fn := func() {}
	m := args.Map{"func": fn}
	wf := m.GetWorkFunc()
	if wf == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I18_Map_GetFirstOfNames_Empty(t *testing.T) {
	m := args.Map{"a": 1}
	if m.GetFirstOfNames() != nil {
		t.Fatal("expected nil for empty names")
	}
}

func Test_I18_Map_GetAsStringSliceFirstOfNames(t *testing.T) {
	m := args.Map{"items": []string{"a", "b"}}
	s := m.GetAsStringSliceFirstOfNames("items")
	if len(s) != 2 {
		t.Fatal("expected 2 items")
	}
}

func Test_I18_Map_GetAsStringSliceFirstOfNames_Empty(t *testing.T) {
	m := args.Map{}
	s := m.GetAsStringSliceFirstOfNames()
	if s != nil {
		t.Fatal("expected nil")
	}
}

func Test_I18_Map_GetAsStringSliceFirstOfNames_Undefined(t *testing.T) {
	m := args.Map{}
	s := m.GetAsStringSliceFirstOfNames("missing")
	if s != nil {
		t.Fatal("expected nil for undefined")
	}
}

func Test_I18_Map_GetAsInt(t *testing.T) {
	m := args.Map{"n": 42}
	n, ok := m.GetAsInt("n")
	if !ok || n != 42 {
		t.Fatal("expected 42")
	}

	_, ok = m.GetAsInt("missing")
	if ok {
		t.Fatal("expected not ok")
	}
}

func Test_I18_Map_GetAsIntDefault(t *testing.T) {
	m := args.Map{"n": 42}
	n := m.GetAsIntDefault("n", 0)
	if n != 42 {
		t.Fatal("expected 42")
	}

	n = m.GetAsIntDefault("missing", 99)
	if n != 99 {
		t.Fatal("expected default 99")
	}
}

func Test_I18_Map_GetAsBool(t *testing.T) {
	m := args.Map{"b": true}
	b, ok := m.GetAsBool("b")
	if !ok || !b {
		t.Fatal("expected true")
	}
}

func Test_I18_Map_GetAsBoolDefault(t *testing.T) {
	m := args.Map{}
	b := m.GetAsBoolDefault("b", true)
	if !b {
		t.Fatal("expected default true")
	}
}

func Test_I18_Map_GetAsString(t *testing.T) {
	m := args.Map{"s": "hello"}
	s, ok := m.GetAsString("s")
	if !ok || s != "hello" {
		t.Fatal("expected hello")
	}
}

func Test_I18_Map_GetAsStringDefault(t *testing.T) {
	m := args.Map{}
	s := m.GetAsStringDefault("s")
	if s != "" {
		t.Fatal("expected empty default")
	}
}

func Test_I18_Map_GetAsStrings(t *testing.T) {
	m := args.Map{"s": []string{"a", "b"}}
	s, ok := m.GetAsStrings("s")
	if !ok || len(s) != 2 {
		t.Fatal("unexpected")
	}
}

func Test_I18_Map_GetAsAnyItems(t *testing.T) {
	m := args.Map{"items": []any{1, "two"}}
	items, ok := m.GetAsAnyItems("items")
	if !ok || len(items) != 2 {
		t.Fatal("unexpected")
	}
}

func Test_I18_Map_ValidArgs(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	validArgs := m.ValidArgs()
	if len(validArgs) != 2 {
		t.Fatal("expected 2 valid args")
	}
}

func Test_I18_Map_Args(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	a := m.Args("a", "b")
	if len(a) != 2 {
		t.Fatal("expected 2 args")
	}
}

func Test_I18_Map_Raw(t *testing.T) {
	m := args.Map{"a": 1}
	raw := m.Raw()
	if raw == nil {
		t.Fatal("expected non-nil raw")
	}
}

func Test_I18_Map_Slice(t *testing.T) {
	m := args.Map{"a": 1}
	s := m.Slice()
	if len(s) != 1 {
		t.Fatal("expected 1 item")
	}
}

func Test_I18_Map_String(t *testing.T) {
	m := args.Map{"a": 1}
	s := m.String()
	if s == "" {
		t.Fatal("expected non-empty string")
	}
}

func Test_I18_Map_GetFirstFuncNameOf(t *testing.T) {
	m := args.Map{}
	name := m.GetFirstFuncNameOf("func")
	if name != "" {
		t.Fatal("expected empty name")
	}
}

func Test_I18_Map_WorkFuncName(t *testing.T) {
	m := args.Map{}
	name := m.WorkFuncName()
	if name != "" {
		t.Fatal("expected empty")
	}
}

// --- Dynamic ---

func Test_I18_Dynamic_NilReceiver(t *testing.T) {
	var d *args.DynamicAny
	if d.ArgsCount() != 0 {
		t.Fatal("expected 0")
	}
	if d.GetWorkFunc() != nil {
		t.Fatal("expected nil")
	}
	if d.HasFirst() {
		t.Fatal("expected false")
	}
	if d.HasDefined("x") {
		t.Fatal("expected false")
	}
	if d.Has("x") {
		t.Fatal("expected false")
	}
	if d.HasDefinedAll("x") {
		t.Fatal("expected false")
	}
	if d.IsKeyInvalid("x") {
		t.Fatal("expected false for nil")
	}
	if d.IsKeyMissing("x") {
		t.Fatal("expected false for nil")
	}
	_, valid := d.Get("x")
	if valid {
		t.Fatal("expected invalid for nil")
	}
	if d.HasExpect() {
		t.Fatal("expected false")
	}
}

func Test_I18_Dynamic_AllMethods(t *testing.T) {
	d := &args.DynamicAny{
		Params: args.Map{
			"first":  "a",
			"second": "b",
			"third":  "c",
			"fourth": "d",
			"fifth":  "e",
			"sixth":  "f",
		},
		Expect: "expected",
	}

	if d.FirstItem() != "a" {
		t.Fatal("first mismatch")
	}
	if d.SecondItem() != "b" {
		t.Fatal("second mismatch")
	}
	if d.ThirdItem() != "c" {
		t.Fatal("third mismatch")
	}
	if d.FourthItem() != "d" {
		t.Fatal("fourth mismatch")
	}
	if d.FifthItem() != "e" {
		t.Fatal("fifth mismatch")
	}
	if d.SixthItem() != "f" {
		t.Fatal("sixth mismatch")
	}
	if d.Expected() != "expected" {
		t.Fatal("expected mismatch")
	}

	if d.GetByIndex(0) == nil {
		t.Fatal("expected item at 0")
	}
	if !d.HasFunc() == false {
		// just exercise
	}
	_ = d.GetFuncName()
	_ = d.FuncWrap()

	if d.HasDefined("first") != true {
		t.Fatal("expected defined")
	}
	if d.Has("missing") {
		t.Fatal("expected not has")
	}
	if !d.HasDefinedAll("first", "second") {
		t.Fatal("expected all defined")
	}
	if d.IsKeyInvalid("first") {
		t.Fatal("expected valid")
	}
	if !d.IsKeyMissing("missing") {
		t.Fatal("expected missing")
	}
	_, _ = d.GetLowerCase("FIRST")
	_ = d.GetDirectLower("FIRST")
	_ = d.Actual()
	_ = d.Arrange()

	_, _ = d.GetAsInt("first")
	_ = d.GetAsIntDefault("first", 0)
	_, _ = d.GetAsString("first")
	_ = d.GetAsStringDefault("first")
	_, _ = d.GetAsStrings("first")
	_, _ = d.GetAsAnyItems("first")

	_ = d.ValidArgs()
	_ = d.Args("first")
	_ = d.Slice()
	s := d.String()
	if s == "" {
		t.Fatal("expected non-empty string")
	}

	// second call should return cached
	s2 := d.String()
	if s != s2 {
		t.Fatal("expected cached string")
	}

	_ = d.AsArgsMapper()
	_ = d.AsArgFuncNameContractsBinder()
	_ = d.AsArgBaseContractsBinder()
}

// --- DynamicFunc ---

func Test_I18_DynamicFunc_NilReceiver(t *testing.T) {
	var df *args.DynamicFuncAny
	if df.ArgsCount() != 0 {
		t.Fatal("expected 0")
	}
	if df.Length() != 0 {
		t.Fatal("expected 0")
	}
	if df.HasDefined("x") {
		t.Fatal("expected false")
	}
	if df.Has("x") {
		t.Fatal("expected false")
	}
	if df.HasDefinedAll("x") {
		t.Fatal("expected false")
	}
	if df.IsKeyInvalid("x") {
		t.Fatal("expected false for nil")
	}
	if df.IsKeyMissing("x") {
		t.Fatal("expected false for nil")
	}
	_, valid := df.Get("x")
	if valid {
		t.Fatal("expected invalid")
	}
	if df.HasFunc() {
		t.Fatal("expected false")
	}
	if df.HasExpect() {
		t.Fatal("expected false")
	}
}

func Test_I18_DynamicFunc_AllMethods(t *testing.T) {
	fn := func(s string) string { return s }
	df := &args.DynamicFuncAny{
		Params: args.Map{
			"first":  "a",
			"second": "b",
			"third":  "c",
			"fourth": "d",
			"fifth":  "e",
			"sixth":  "f",
		},
		WorkFunc: fn,
		Expect:   "expected",
	}

	_ = df.GetWorkFunc()
	_ = df.HasFirst()
	_ = df.GetByIndex(0)
	_ = df.GetByIndex(100)
	_ = df.FirstItem()
	_ = df.SecondItem()
	_ = df.ThirdItem()
	_ = df.FourthItem()
	_ = df.FifthItem()
	_ = df.SixthItem()
	_ = df.Expected()
	_ = df.When()
	_ = df.Title()
	_, _ = df.GetLowerCase("FIRST")
	_ = df.GetDirectLower("FIRST")
	_ = df.GetDirectLower("MISSING")
	_ = df.Actual()
	_ = df.Arrange()
	_, _ = df.GetAsInt("first")
	_, _ = df.GetAsString("first")
	_, _ = df.GetAsStrings("first")
	_, _ = df.GetAsAnyItems("first")
	_ = df.HasFunc()
	_ = df.HasExpect()
	_ = df.GetFuncName()
	_ = df.FuncWrap()
	_ = df.ValidArgs()
	_ = df.Args("first")

	s := df.Slice()
	if len(s) == 0 {
		t.Fatal("expected non-empty slice")
	}

	str := df.String()
	if str == "" {
		t.Fatal("expected non-empty string")
	}

	// cached
	str2 := df.String()
	if str != str2 {
		t.Fatal("expected cached")
	}

	_ = df.AsArgsMapper()
	_ = df.AsArgFuncNameContractsBinder()
	_ = df.AsArgBaseContractsBinder()
}

// --- One ---

func Test_I18_One_AllMethods(t *testing.T) {
	o := &args.OneAny{First: "hello", Expect: "world"}
	if o.FirstItem() != "hello" {
		t.Fatal("first mismatch")
	}
	if o.Expected() != "world" {
		t.Fatal("expected mismatch")
	}
	_ = o.ArgTwo()
	if !o.HasFirst() {
		t.Fatal("expected HasFirst")
	}
	if !o.HasExpect() {
		t.Fatal("expected HasExpect")
	}
	if len(o.ValidArgs()) != 1 {
		t.Fatal("expected 1 valid arg")
	}
	if len(o.Args(1)) != 1 {
		t.Fatal("expected 1 arg")
	}
	if o.ArgsCount() != 1 {
		t.Fatal("expected 1")
	}
	_ = o.Slice()
	_ = o.GetByIndex(0)
	s := o.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
	_ = o.LeftRight()
	_ = args.OneAny{First: "x"}.AsOneParameter()
	_ = args.OneAny{First: "x"}.AsArgBaseContractsBinder()
}

// --- Two ---

func Test_I18_Two_AllMethods(t *testing.T) {
	tw := &args.TwoAny{First: "a", Second: "b", Expect: "c"}
	_ = tw.FirstItem()
	_ = tw.SecondItem()
	_ = tw.Expected()
	_ = tw.ArgTwo()
	_ = tw.HasFirst()
	_ = tw.HasSecond()
	_ = tw.HasExpect()
	_ = tw.ValidArgs()
	_ = tw.ArgsCount()
	_ = tw.Args(2)
	_ = tw.Slice()
	_ = tw.GetByIndex(0)
	_ = tw.String()
	_ = tw.LeftRight()
	_ = args.TwoAny{}.AsTwoParameter()
	_ = args.TwoAny{}.AsArgBaseContractsBinder()
}

// --- Three ---

func Test_I18_Three_AllMethods(t *testing.T) {
	th := &args.ThreeAny{First: "a", Second: "b", Third: "c", Expect: "d"}
	_ = th.ArgsCount()
	_ = th.FirstItem()
	_ = th.SecondItem()
	_ = th.ThirdItem()
	_ = th.Expected()
	_ = th.ArgTwo()
	_ = th.ArgThree()
	_ = th.HasFirst()
	_ = th.HasSecond()
	_ = th.HasThird()
	_ = th.HasExpect()
	_ = th.ValidArgs()
	_ = th.Args(3)
	_ = th.Slice()
	_ = th.GetByIndex(0)
	_ = th.String()
	_ = th.LeftRight()
	_ = args.ThreeAny{}.AsThreeParameter()
	_ = args.ThreeAny{}.AsArgBaseContractsBinder()
}

// --- Four ---

func Test_I18_Four_AllMethods(t *testing.T) {
	f := &args.FourAny{First: "a", Second: "b", Third: "c", Fourth: "d", Expect: "e"}
	_ = f.ArgsCount()
	_ = f.FirstItem()
	_ = f.SecondItem()
	_ = f.ThirdItem()
	_ = f.FourthItem()
	_ = f.Expected()
	_ = f.ArgTwo()
	_ = f.ArgThree()
	_ = f.HasFirst()
	_ = f.HasSecond()
	_ = f.HasThird()
	_ = f.HasFourth()
	_ = f.HasExpect()
	_ = f.ValidArgs()
	_ = f.Args(4)
	_ = f.Slice()
	_ = f.GetByIndex(0)
	_ = f.String()
	_ = args.FourAny{}.AsFourParameter()
	_ = args.FourAny{}.AsArgBaseContractsBinder()
}

// --- Five ---

func Test_I18_Five_AllMethods(t *testing.T) {
	f := &args.FiveAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Expect: "f"}
	_ = f.ArgsCount()
	_ = f.FirstItem()
	_ = f.SecondItem()
	_ = f.ThirdItem()
	_ = f.FourthItem()
	_ = f.FifthItem()
	_ = f.Expected()
	_ = f.ArgTwo()
	_ = f.ArgThree()
	_ = f.ArgFour()
	_ = f.HasFirst()
	_ = f.HasSecond()
	_ = f.HasThird()
	_ = f.HasFourth()
	_ = f.HasFifth()
	_ = f.HasExpect()
	_ = f.ValidArgs()
	_ = f.Args(5)
	_ = f.Slice()
	_ = f.GetByIndex(0)
	_ = f.String()
	_ = args.FiveAny{}.AsFifthParameter()
	_ = args.FiveAny{}.AsArgBaseContractsBinder()
}

// --- Six ---

func Test_I18_Six_AllMethods(t *testing.T) {
	s := &args.SixAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f", Expect: "g"}
	_ = s.ArgsCount()
	_ = s.FirstItem()
	_ = s.SecondItem()
	_ = s.ThirdItem()
	_ = s.FourthItem()
	_ = s.FifthItem()
	_ = s.SixthItem()
	_ = s.Expected()
	_ = s.ArgTwo()
	_ = s.ArgThree()
	_ = s.ArgFour()
	_ = s.ArgFive()
	_ = s.HasFirst()
	_ = s.HasSecond()
	_ = s.HasThird()
	_ = s.HasFourth()
	_ = s.HasFifth()
	_ = s.HasSixth()
	_ = s.HasExpect()
	_ = s.ValidArgs()
	_ = s.Args(6)
	_ = s.Slice()
	_ = s.GetByIndex(0)
	_ = s.String()
	_ = args.SixAny{}.AsSixthParameter()
	_ = args.SixAny{}.AsArgBaseContractsBinder()
}

// --- LeftRight ---

func Test_I18_LeftRight_AllMethods(t *testing.T) {
	lr := &args.LeftRightAny{Left: "a", Right: "b", Expect: "c"}
	_ = lr.ArgsCount()
	_ = lr.FirstItem()
	_ = lr.SecondItem()
	_ = lr.Expected()
	_ = lr.ArgTwo()
	_ = lr.HasFirst()
	_ = lr.HasSecond()
	_ = lr.HasLeft()
	_ = lr.HasRight()
	_ = lr.HasExpect()
	_ = lr.ValidArgs()
	_ = lr.Args(2)
	_ = lr.Slice()
	_ = lr.GetByIndex(0)
	_ = lr.String()
	clone := lr.Clone()
	if clone.Left != "a" {
		t.Fatal("clone mismatch")
	}
	_ = args.LeftRightAny{}.AsTwoParameter()
	_ = args.LeftRightAny{}.AsArgBaseContractsBinder()
}

// --- Holder ---

func Test_I18_Holder_AllMethods(t *testing.T) {
	fn := func(s string) string { return s }
	h := &args.HolderAny{
		First: "a", Second: "b", Third: "c",
		Fourth: "d", Fifth: "e", Sixth: "f",
		WorkFunc: fn, Expect: "exp",
	}
	_ = h.GetWorkFunc()
	_ = h.ArgsCount()
	_ = h.FirstItem()
	_ = h.SecondItem()
	_ = h.ThirdItem()
	_ = h.FourthItem()
	_ = h.FifthItem()
	_ = h.SixthItem()
	_ = h.Expected()
	_ = h.ArgTwo()
	_ = h.ArgThree()
	_ = h.ArgFour()
	_ = h.ArgFive()
	_ = h.HasFirst()
	_ = h.HasSecond()
	_ = h.HasThird()
	_ = h.HasFourth()
	_ = h.HasFifth()
	_ = h.HasSixth()
	_ = h.HasFunc()
	_ = h.HasExpect()
	_ = h.GetFuncName()
	_ = h.FuncWrap()
	_ = h.ValidArgs()
	_ = h.Args(6)
	_ = h.Slice()
	_ = h.GetByIndex(0)
	_ = h.String()

	hVal := args.HolderAny{}
	_ = hVal.AsSixthParameter()
	_ = hVal.AsArgFuncContractsBinder()
}

// --- String ---

func Test_I18_String_AllMethods(t *testing.T) {
	s := args.String("hello")
	_ = s.Concat(" world")
	_ = s.Join("-", "a", "b")
	_ = s.Split("l")
	_ = s.DoubleQuote()
	_ = s.DoubleQuoteQ()
	_ = s.SingleQuote()
	_ = s.ValueDoubleQuote()
	_ = s.String()
	_ = s.Bytes()
	_ = s.Runes()
	_ = s.Length()
	_ = s.Count()
	_ = s.IsEmptyOrWhitespace()
	_ = s.TrimSpace()
	_ = s.ReplaceAll("h", "H")
	_ = s.Substring(0, 3)
	if s.IsEmpty() {
		t.Fatal("expected not empty")
	}
	if !s.HasCharacter() {
		t.Fatal("expected has character")
	}
	if !s.IsDefined() {
		t.Fatal("expected defined")
	}
	_ = s.AscIILength()

	empty := args.String("")
	if !empty.IsEmpty() {
		t.Fatal("expected empty")
	}
	if !empty.IsEmptyOrWhitespace() {
		t.Fatal("expected empty or whitespace")
	}
	ws := args.String("   ")
	if !ws.IsEmptyOrWhitespace() {
		t.Fatal("expected whitespace only")
	}
}

func Test_I18_String_TrimReplaceMap(t *testing.T) {
	s := args.String("hello {name}")
	result := s.TrimReplaceMap(map[string]string{"{name}": "world"})
	if !strings.Contains(string(result), "world") {
		t.Fatal("expected replacement")
	}
}

// --- Empty ---

func Test_I18_Empty_Map(t *testing.T) {
	m := args.Empty.Map()
	if m == nil || len(m) != 0 {
		t.Fatal("expected empty map")
	}
}

func Test_I18_Empty_FuncWrap(t *testing.T) {
	fw := args.Empty.FuncWrap()
	if fw == nil || !fw.IsInvalid() {
		t.Fatal("expected invalid func wrap")
	}
}

func Test_I18_Empty_FuncMap(t *testing.T) {
	fm := args.Empty.FuncMap()
	if fm == nil || len(fm) != 0 {
		t.Fatal("expected empty func map")
	}
}

func Test_I18_Empty_Holder(t *testing.T) {
	h := args.Empty.Holder()
	if h.HasFirst() {
		t.Fatal("expected no first")
	}
}

// --- FuncWrap ---

func Test_I18_NewTypedFuncWrap_Valid(t *testing.T) {
	fn := func(s string) int { return len(s) }
	fw := args.NewTypedFuncWrap(fn)
	if fw.IsInvalid() {
		t.Fatal("expected valid")
	}
	if fw.ArgsCount() != 1 {
		t.Fatal("expected 1 arg")
	}
}

func Test_I18_NewTypedFuncWrap_Nil(t *testing.T) {
	var fn func()
	fw := args.NewTypedFuncWrap(fn)
	if !fw.IsInvalid() {
		t.Fatal("expected invalid for nil func")
	}
}

func Test_I18_NewTypedFuncWrap_NotFunc(t *testing.T) {
	fw := args.NewTypedFuncWrap("not a func")
	if !fw.IsInvalid() {
		t.Fatal("expected invalid for non-func")
	}
}

func Test_I18_FuncWrap_GetFuncName_Nil(t *testing.T) {
	var fw *args.FuncWrapAny
	if fw.GetFuncName() != "" {
		t.Fatal("expected empty for nil")
	}
}

func Test_I18_FuncWrap_GetPascalCaseFuncName_Nil(t *testing.T) {
	var fw *args.FuncWrapAny
	if fw.GetPascalCaseFuncName() != "" {
		t.Fatal("expected empty for nil")
	}
}

func Test_I18_FuncWrap_IsEqual(t *testing.T) {
	fn := func() {}
	fw1 := args.NewTypedFuncWrap(fn)
	fw2 := args.NewTypedFuncWrap(fn)

	if fw1.IsNotEqual(fw2) {
		t.Fatal("expected equal")
	}

	var nilFw *args.FuncWrapAny
	if !nilFw.IsEqual(nil) {
		t.Fatal("expected nil == nil")
	}
	if nilFw.IsEqual(fw1) {
		t.Fatal("expected nil != non-nil")
	}
}

func Test_I18_FuncWrap_PkgPath(t *testing.T) {
	fn := func() {}
	fw := args.NewTypedFuncWrap(fn)
	p := fw.PkgPath()
	_ = p
	// second call for cached path
	p2 := fw.PkgPath()
	_ = p2
}

func Test_I18_FuncWrap_PkgNameOnly(t *testing.T) {
	fn := func() {}
	fw := args.NewTypedFuncWrap(fn)
	p := fw.PkgNameOnly()
	_ = p
	p2 := fw.PkgNameOnly()
	_ = p2
}

func Test_I18_FuncWrap_FuncDirectInvokeName(t *testing.T) {
	fn := func() {}
	fw := args.NewTypedFuncWrap(fn)
	n := fw.FuncDirectInvokeName()
	_ = n
	n2 := fw.FuncDirectInvokeName()
	_ = n2
}

func Test_I18_FuncWrap_IsEqualValue(t *testing.T) {
	fn := func() {}
	fw1 := args.NewTypedFuncWrap(fn)
	fw2 := args.NewTypedFuncWrap(fn)
	if !fw1.IsEqualValue(*fw2) {
		t.Fatal("expected equal value")
	}
}

// --- FuncMap ---

func Test_I18_FuncMap_Basic(t *testing.T) {
	fm := args.FuncMap{}
	if !fm.IsEmpty() {
		t.Fatal("expected empty")
	}
	if fm.Length() != 0 {
		t.Fatal("expected 0")
	}
	if fm.Count() != 0 {
		t.Fatal("expected 0")
	}
	if fm.HasAnyItem() {
		t.Fatal("expected no items")
	}
	if fm.Has("x") {
		t.Fatal("expected not has")
	}
	if fm.IsContains("x") {
		t.Fatal("expected not contains")
	}
	if fm.Get("x") != nil {
		t.Fatal("expected nil")
	}
}

func Test_I18_FuncMap_AddAndInvoke(t *testing.T) {
	fn := func(n int) int { return n * 2 }
	fm := args.FuncMap{}
	fm.Add(fn)

	if fm.IsEmpty() {
		t.Fatal("expected non-empty")
	}

	name := fm.GetPascalCaseFuncName("x")
	_ = name
}

func Test_I18_FuncMap_InvalidError(t *testing.T) {
	fm := args.FuncMap{}
	err := fm.InvalidError()
	if err == nil {
		t.Fatal("expected error for empty map")
	}
}

// --- NewFuncWrap creator ---

func Test_I18_NewFuncWrap_Default_Nil(t *testing.T) {
	fw := args.NewFuncWrap.Default(nil)
	if fw == nil || !fw.IsInvalid() {
		t.Fatal("expected invalid for nil")
	}
}

func Test_I18_NewFuncWrap_Default_Valid(t *testing.T) {
	fn := func() {}
	fw := args.NewFuncWrap.Default(fn)
	if fw == nil || fw.IsInvalid() {
		t.Fatal("expected valid")
	}
}

// --- FuncDetector ---

func Test_I18_FuncDetector(t *testing.T) {
	fn := func() {}
	if !args.FuncDetector.IsFunc(fn) {
		t.Fatal("expected is func")
	}
	if args.FuncDetector.IsFunc("not a func") {
		t.Fatal("expected not func")
	}
}
