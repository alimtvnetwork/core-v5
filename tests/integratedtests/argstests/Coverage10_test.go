package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// args Coverage — One, Two, Map, LeftRight, String, Holder
// ══════════════════════════════════════════════════════════════════════════════

// --- One ---

func Test_CovArgs_01_One_Basic(t *testing.T) {
	o := &args.OneAny{First: "hello", Expect: 42}
	if o.FirstItem() != "hello" {
		t.Fatal("expected hello")
	}
	if o.Expected() != 42 {
		t.Fatal("expected 42")
	}
	if !o.HasFirst() {
		t.Fatal("expected true")
	}
	if !o.HasExpect() {
		t.Fatal("expected true")
	}
	if o.ArgsCount() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovArgs_02_One_ValidArgs_Args_Slice(t *testing.T) {
	o := &args.OneAny{First: "hello", Expect: 42}
	va := o.ValidArgs()
	if len(va) != 1 {
		t.Fatal("expected 1")
	}
	a := o.Args(1)
	if len(a) != 1 {
		t.Fatal("expected 1")
	}
	a0 := o.Args(0)
	if len(a0) != 0 {
		t.Fatal("expected 0")
	}
	sl := o.Slice()
	if len(sl) < 1 {
		t.Fatal("expected at least 1")
	}
	// cached
	sl2 := o.Slice()
	if len(sl2) != len(sl) {
		t.Fatal("expected same")
	}
}

func Test_CovArgs_03_One_GetByIndex_String(t *testing.T) {
	o := &args.OneAny{First: "hello"}
	_ = o.GetByIndex(0)
	_ = o.GetByIndex(99)
	s := o.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovArgs_04_One_LeftRight(t *testing.T) {
	o := &args.OneAny{First: "hello", Expect: 42}
	lr := o.LeftRight()
	if lr.Left != "hello" {
		t.Fatal("expected hello")
	}
}

func Test_CovArgs_05_One_ArgTwo(t *testing.T) {
	o := &args.OneAny{First: "hello", Expect: 42}
	a2 := o.ArgTwo()
	if a2.First != "hello" {
		t.Fatal("expected hello")
	}
}

func Test_CovArgs_06_One_AsInterfaces(t *testing.T) {
	o := args.OneAny{First: "hello"}
	_ = o.AsOneParameter()
	_ = o.AsArgBaseContractsBinder()
}

// --- Two ---

func Test_CovArgs_07_Two_Basic(t *testing.T) {
	tw := &args.TwoAny{First: "a", Second: "b", Expect: 1}
	if tw.FirstItem() != "a" {
		t.Fatal("expected a")
	}
	if tw.SecondItem() != "b" {
		t.Fatal("expected b")
	}
	if tw.Expected() != 1 {
		t.Fatal("expected 1")
	}
	if !tw.HasFirst() {
		t.Fatal("expected true")
	}
	if !tw.HasSecond() {
		t.Fatal("expected true")
	}
	if !tw.HasExpect() {
		t.Fatal("expected true")
	}
	if tw.ArgsCount() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CovArgs_08_Two_ValidArgs_Args_Slice(t *testing.T) {
	tw := &args.TwoAny{First: "a", Second: "b", Expect: 1}
	va := tw.ValidArgs()
	if len(va) != 2 {
		t.Fatal("expected 2")
	}
	a := tw.Args(2)
	if len(a) != 2 {
		t.Fatal("expected 2")
	}
	a1 := tw.Args(1)
	if len(a1) != 1 {
		t.Fatal("expected 1")
	}
	sl := tw.Slice()
	if len(sl) < 2 {
		t.Fatal("expected at least 2")
	}
}

func Test_CovArgs_09_Two_GetByIndex_String(t *testing.T) {
	tw := &args.TwoAny{First: "a", Second: "b"}
	_ = tw.GetByIndex(0)
	_ = tw.GetByIndex(99)
	s := tw.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovArgs_10_Two_LeftRight_ArgTwo(t *testing.T) {
	tw := &args.TwoAny{First: "a", Second: "b", Expect: 1}
	lr := tw.LeftRight()
	if lr.Left != "a" || lr.Right != "b" {
		t.Fatal("expected a,b")
	}
	a2 := tw.ArgTwo()
	if a2.First != "a" {
		t.Fatal("expected a")
	}
}

func Test_CovArgs_11_Two_AsInterfaces(t *testing.T) {
	tw := args.TwoAny{First: "a", Second: "b"}
	_ = tw.AsTwoParameter()
	_ = tw.AsArgBaseContractsBinder()
}

// --- Map ---

func Test_CovArgs_12_Map_Basic(t *testing.T) {
	m := args.Map{
		"first": "hello",
		"expected": 42,
	}
	if m.Length() != 2 {
		t.Fatal("expected 2")
	}
	if m.Expected() != 42 {
		t.Fatal("expected 42")
	}
	if !m.HasExpect() {
		t.Fatal("expected true")
	}
	if !m.HasFirst() {
		t.Fatal("expected true")
	}
	if m.FirstItem() != "hello" {
		t.Fatal("expected hello")
	}
}

func Test_CovArgs_13_Map_ArgsCount(t *testing.T) {
	m := args.Map{
		"first": "hello",
		"expected": 42,
		"func": func() {},
	}
	// ArgsCount excludes expected and func
	ac := m.ArgsCount()
	if ac != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovArgs_14_Map_Get_Has_HasDefined(t *testing.T) {
	m := args.Map{"key": "val"}
	v, ok := m.Get("key")
	if !ok || v != "val" {
		t.Fatal("expected val")
	}
	_, ok2 := m.Get("missing")
	if ok2 {
		t.Fatal("expected false")
	}
	if !m.Has("key") {
		t.Fatal("expected true")
	}
	if m.Has("missing") {
		t.Fatal("expected false")
	}
	if !m.HasDefined("key") {
		t.Fatal("expected true")
	}
	if m.HasDefined("missing") {
		t.Fatal("expected false")
	}
}

func Test_CovArgs_15_Map_HasDefinedAll_IsKeyInvalid_IsKeyMissing(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	if !m.HasDefinedAll("a", "b") {
		t.Fatal("expected true")
	}
	if m.HasDefinedAll("a", "missing") {
		t.Fatal("expected false")
	}
	if m.HasDefinedAll() {
		t.Fatal("expected false for empty names")
	}
	if m.IsKeyMissing("a") {
		t.Fatal("expected false")
	}
	if !m.IsKeyMissing("missing") {
		t.Fatal("expected true")
	}
}

func Test_CovArgs_16_Map_NilMap(t *testing.T) {
	var m args.Map
	_, ok := m.Get("key")
	if ok {
		t.Fatal("expected false")
	}
	if m.Has("key") {
		t.Fatal("expected false")
	}
	if m.HasDefined("key") {
		t.Fatal("expected false")
	}
	if m.HasDefinedAll("key") {
		t.Fatal("expected false")
	}
	if m.IsKeyInvalid("key") {
		t.Fatal("expected false")
	}
	if m.IsKeyMissing("key") {
		t.Fatal("expected false")
	}
}

func Test_CovArgs_17_Map_SortedKeys(t *testing.T) {
	m := args.Map{"b": 2, "a": 1}
	keys, err := m.SortedKeys()
	if err != nil || len(keys) != 2 {
		t.Fatal("expected 2 keys")
	}
	if keys[0] != "a" {
		t.Fatal("expected a first")
	}
	// empty map
	m2 := args.Map{}
	keys2, _ := m2.SortedKeys()
	if len(keys2) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovArgs_18_Map_SortedKeysMust(t *testing.T) {
	m := args.Map{"a": 1}
	keys := m.SortedKeysMust()
	if len(keys) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovArgs_19_Map_Items(t *testing.T) {
	m := args.Map{
		"first": "a",
		"second": "b",
		"third": "c",
		"fourth": "d",
		"fifth": "e",
		"sixth": "f",
		"seventh": "g",
	}
	if m.SecondItem() != "b" {
		t.Fatal("expected b")
	}
	if m.ThirdItem() != "c" {
		t.Fatal("expected c")
	}
	if m.FourthItem() != "d" {
		t.Fatal("expected d")
	}
	if m.FifthItem() != "e" {
		t.Fatal("expected e")
	}
	if m.SixthItem() != "f" {
		t.Fatal("expected f")
	}
	if m.Seventh() != "g" {
		t.Fatal("expected g")
	}
}

func Test_CovArgs_20_Map_GetLowerCase_GetDirectLower(t *testing.T) {
	m := args.Map{"key": "val"}
	v, ok := m.GetLowerCase("KEY")
	if !ok || v != "val" {
		t.Fatal("expected val")
	}
	v2 := m.GetDirectLower("KEY")
	if v2 != "val" {
		t.Fatal("expected val")
	}
	v3 := m.GetDirectLower("MISSING")
	if v3 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovArgs_21_Map_Expect_Actual_Arrange(t *testing.T) {
	m := args.Map{"expect": 1, "actual": 2, "arrange": 3}
	if m.Expect() != 1 {
		t.Fatal("expected 1")
	}
	if m.Actual() != 2 {
		t.Fatal("expected 2")
	}
	if m.Arrange() != 3 {
		t.Fatal("expected 3")
	}
}

func Test_CovArgs_22_Map_SetActual(t *testing.T) {
	m := args.Map{}
	m.SetActual("val")
	if m.Actual() != "val" {
		t.Fatal("expected val")
	}
}

func Test_CovArgs_23_Map_When_Title(t *testing.T) {
	m := args.Map{"when": "w", "title": "t"}
	if m.When() != "w" {
		t.Fatal("expected w")
	}
	if m.Title() != "t" {
		t.Fatal("expected t")
	}
}

func Test_CovArgs_24_Map_GetByIndex(t *testing.T) {
	m := args.Map{"a": 1}
	_ = m.GetByIndex(0)
	r := m.GetByIndex(99)
	if r != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovArgs_25_Map_Raw_Args_ValidArgs(t *testing.T) {
	m := args.Map{"first": "a", "second": "b"}
	raw := m.Raw()
	if len(raw) != 2 {
		t.Fatal("expected 2")
	}
	a := m.Args("first", "second")
	if len(a) != 2 {
		t.Fatal("expected 2")
	}
	va := m.ValidArgs()
	if len(va) == 0 {
		t.Fatal("expected non-empty")
	}
}

func Test_CovArgs_26_Map_GetFirstOfNames(t *testing.T) {
	m := args.Map{"a": 1}
	v := m.GetFirstOfNames("missing", "a")
	if v != 1 {
		t.Fatal("expected 1")
	}
	v2 := m.GetFirstOfNames()
	if v2 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovArgs_27_Map_GetAsStringSliceFirstOfNames(t *testing.T) {
	m := args.Map{"items": []string{"a", "b"}}
	r := m.GetAsStringSliceFirstOfNames("items")
	if len(r) != 2 {
		t.Fatal("expected 2")
	}
	r2 := m.GetAsStringSliceFirstOfNames("missing")
	if r2 != nil {
		t.Fatal("expected nil")
	}
	r3 := m.GetAsStringSliceFirstOfNames()
	if r3 != nil {
		t.Fatal("expected nil")
	}
}

// --- LeftRight ---

func Test_CovArgs_28_LeftRight(t *testing.T) {
	lr := &args.LeftRightAny{Left: "a", Right: "b", Expect: 1}
	if lr.FirstItem() != "a" {
		t.Fatal("expected a")
	}
	if lr.SecondItem() != "b" {
		t.Fatal("expected b")
	}
	if lr.Expected() != 1 {
		t.Fatal("expected 1")
	}
	if lr.ArgsCount() != 2 {
		t.Fatal("expected 2")
	}
	if !lr.HasFirst() {
		t.Fatal("expected true")
	}
	if !lr.HasSecond() {
		t.Fatal("expected true")
	}
	if !lr.HasLeft() {
		t.Fatal("expected true")
	}
	if !lr.HasRight() {
		t.Fatal("expected true")
	}
	if !lr.HasExpect() {
		t.Fatal("expected true")
	}
	va := lr.ValidArgs()
	if len(va) != 2 {
		t.Fatal("expected 2")
	}
	a := lr.Args(2)
	if len(a) != 2 {
		t.Fatal("expected 2")
	}
	sl := lr.Slice()
	if len(sl) < 2 {
		t.Fatal("expected at least 2")
	}
	_ = lr.GetByIndex(0)
	_ = lr.String()
	a2 := lr.ArgTwo()
	if a2.First != "a" {
		t.Fatal("expected a")
	}
	c := lr.Clone()
	if c.Left != "a" {
		t.Fatal("expected a")
	}
	lrV := args.LeftRightAny{Left: "a", Right: "b"}
	_ = lrV.AsTwoParameter()
	_ = lrV.AsArgBaseContractsBinder()
}

// --- Holder ---

func Test_CovArgs_29_Holder(t *testing.T) {
	fn := func(s string) string { return s }
	h := &args.HolderAny{
		First:    "a",
		Second:   "b",
		Third:    "c",
		Fourth:   "d",
		Fifth:    "e",
		Sixth:    "f",
		WorkFunc: fn,
		Expect:   "x",
	}
	if h.FirstItem() != "a" {
		t.Fatal("expected a")
	}
	if h.SecondItem() != "b" {
		t.Fatal("expected b")
	}
	if h.ThirdItem() != "c" {
		t.Fatal("expected c")
	}
	if h.FourthItem() != "d" {
		t.Fatal("expected d")
	}
	if h.FifthItem() != "e" {
		t.Fatal("expected e")
	}
	if h.SixthItem() != "f" {
		t.Fatal("expected f")
	}
	if h.Expected() != "x" {
		t.Fatal("expected x")
	}
	if h.ArgsCount() != 7 {
		t.Fatal("expected 7")
	}
	if !h.HasFirst() {
		t.Fatal("expected true")
	}
	if !h.HasSecond() {
		t.Fatal("expected true")
	}
	if !h.HasThird() {
		t.Fatal("expected true")
	}
	if !h.HasFourth() {
		t.Fatal("expected true")
	}
	if !h.HasFifth() {
		t.Fatal("expected true")
	}
	if !h.HasSixth() {
		t.Fatal("expected true")
	}
	if !h.HasFunc() {
		t.Fatal("expected true")
	}
	if !h.HasExpect() {
		t.Fatal("expected true")
	}
	_ = h.GetWorkFunc()
	_ = h.GetFuncName()
	va := h.ValidArgs()
	if len(va) != 6 {
		t.Fatal("expected 6")
	}
	a := h.Args(6)
	if len(a) != 6 {
		t.Fatal("expected 6")
	}
	sl := h.Slice()
	if len(sl) < 6 {
		t.Fatal("expected at least 6")
	}
	_ = h.GetByIndex(0)
	_ = h.String()
	_ = h.ArgTwo()
	_ = h.ArgThree()
	_ = h.ArgFour()
	_ = h.ArgFive()
	hv := args.HolderAny{First: "a"}
	_ = hv.AsSixthParameter()
	_ = hv.AsArgFuncContractsBinder()
}

// --- String ---

func Test_CovArgs_30_String(t *testing.T) {
	s := args.String("hello")
	if s.String() != "hello" {
		t.Fatal("expected hello")
	}
	if s.Length() != 5 {
		t.Fatal("expected 5")
	}
	if s.Count() != 5 {
		t.Fatal("expected 5")
	}
	if s.AscIILength() != 5 {
		t.Fatal("expected 5")
	}
	if s.IsEmpty() {
		t.Fatal("expected false")
	}
	if !s.HasCharacter() {
		t.Fatal("expected true")
	}
	if !s.IsDefined() {
		t.Fatal("expected true")
	}
	if s.IsEmptyOrWhitespace() {
		t.Fatal("expected false")
	}
	_ = s.Bytes()
	_ = s.Runes()
	_ = s.TrimSpace()
	_ = s.DoubleQuote()
	_ = s.DoubleQuoteQ()
	_ = s.SingleQuote()
	_ = s.ValueDoubleQuote()
	_ = s.ReplaceAll("h", "H")
	_ = s.Concat("world")
	_ = s.Join(",", "world")
	_ = s.Split(",")
	_ = s.Substring(0, 3)
}

func Test_CovArgs_31_String_Empty(t *testing.T) {
	s := args.String("")
	if !s.IsEmpty() {
		t.Fatal("expected true")
	}
	if s.HasCharacter() {
		t.Fatal("expected false")
	}
	if !s.IsEmptyOrWhitespace() {
		t.Fatal("expected true")
	}
}

func Test_CovArgs_32_String_TrimReplaceMap(t *testing.T) {
	s := args.String("Hello {name}")
	r := s.TrimReplaceMap(map[string]string{"{name}": "World"})
	if r.String() != "Hello World" {
		t.Fatalf("expected 'Hello World', got '%s'", r.String())
	}
}
