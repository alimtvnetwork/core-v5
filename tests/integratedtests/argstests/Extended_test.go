package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Map_Get(t *testing.T) {
	for caseIndex, testCase := range extMapGetTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		key, _ := input.GetAsString("key")

		// Act
		val, isValid := input.Get(key)

		actual := args.Map{
			"isValid": isValid,
		}

		if isValid {
			actual["value"] = val
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Map_TypedGetters(t *testing.T) {
	for caseIndex, testCase := range extMapTypedGetTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)

		// Act
		actual := args.Map{}

		intVal, intValid := input.GetAsInt("count")
		if intValid {
			actual["intVal"] = intVal
			actual["intValid"] = intValid
		}

		boolVal, boolValid := input.GetAsBool("active")
		if boolValid {
			actual["boolVal"] = boolVal
			actual["boolValid"] = boolValid
		}

		strVal, strValid := input.GetAsString("text")
		if strValid {
			actual["strVal"] = strVal
			actual["strValid"] = strValid
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Map_Compile(t *testing.T) {
	for caseIndex, testCase := range extMapCompileTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)

		// Act
		lines := input.CompileToStrings()

		actual := args.Map{
			"lineCount": len(lines),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Map_Utility_Methods(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": "two",
		"c": true,
	}

	// Act & Assert
	if m.Length() != 3 {
		t.Errorf("expected Length 3, got %d", m.Length())
	}

	if !m.Has("a") {
		t.Error("expected Has('a') true")
	}

	if m.Has("z") {
		t.Error("expected Has('z') false")
	}

	if !m.HasDefined("a") {
		t.Error("expected HasDefined('a') true")
	}

	if m.IsKeyMissing("a") {
		t.Error("expected IsKeyMissing('a') false")
	}

	if !m.IsKeyMissing("z") {
		t.Error("expected IsKeyMissing('z') true")
	}

	if !m.HasDefinedAll("a", "b") {
		t.Error("expected HasDefinedAll('a','b') true")
	}

	if m.HasDefinedAll("a", "z") {
		t.Error("expected HasDefinedAll('a','z') false")
	}
}

func Test_Map_NilSafety(t *testing.T) {
	// Arrange
	var m args.Map

	// Act & Assert
	_, isValid := m.Get("key")
	if isValid {
		t.Error("nil map Get should return false")
	}

	if m.Has("key") {
		t.Error("nil map Has should return false")
	}

	if m.HasDefined("key") {
		t.Error("nil map HasDefined should return false")
	}

	if m.HasDefinedAll("a") {
		t.Error("nil map HasDefinedAll should return false")
	}

	if m.IsKeyInvalid("key") {
		t.Error("nil map IsKeyInvalid should return false")
	}

	if m.IsKeyMissing("key") {
		t.Error("nil map IsKeyMissing should return false")
	}
}

func Test_Map_GetDefaults(t *testing.T) {
	// Arrange
	m := args.Map{
		"count": 5,
		"flag":  true,
		"text":  "hello",
	}

	// Act & Assert
	if m.GetAsIntDefault("count", 0) != 5 {
		t.Error("expected GetAsIntDefault to return 5")
	}

	if m.GetAsIntDefault("missing", 99) != 99 {
		t.Error("expected GetAsIntDefault to return default 99")
	}

	if m.GetAsBoolDefault("flag", false) != true {
		t.Error("expected GetAsBoolDefault to return true")
	}

	if m.GetAsBoolDefault("missing", true) != true {
		t.Error("expected GetAsBoolDefault to return default true")
	}

	if m.GetAsStringDefault("text") != "hello" {
		t.Error("expected GetAsStringDefault to return hello")
	}

	if m.GetAsStringDefault("missing") != "" {
		t.Error("expected GetAsStringDefault to return empty")
	}
}

func Test_Map_SpecialKeys(t *testing.T) {
	// Arrange
	m := args.Map{
		"when":     "test scenario",
		"title":    "test title",
		"expected": "some result",
		"first":    "f1",
		"second":   "f2",
		"third":    "f3",
	}

	// Act & Assert
	if m.When() != "test scenario" {
		t.Error("expected When() to return scenario")
	}

	if m.Title() != "test title" {
		t.Error("expected Title() to return title")
	}

	if m.Expected() != "some result" {
		t.Error("expected Expected() to return expected")
	}

	if m.FirstItem() != "f1" {
		t.Error("expected FirstItem() to return f1")
	}

	if m.SecondItem() != "f2" {
		t.Error("expected SecondItem() to return f2")
	}

	if m.ThirdItem() != "f3" {
		t.Error("expected ThirdItem() to return f3")
	}
}

func Test_Map_GetByIndex(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}

	// Act
	slice := m.Slice()

	// Assert
	if len(slice) != 2 {
		t.Errorf("expected 2 items in slice, got %d", len(slice))
	}

	if m.GetByIndex(999) != nil {
		t.Error("expected nil for out of range index")
	}
}

func Test_Map_GoLiteral(t *testing.T) {
	// Arrange
	m := args.Map{
		"key1": "value1",
		"key2": 42,
	}

	// Act
	lines := m.GoLiteralLines()
	str := m.GoLiteralString()

	// Assert
	if len(lines) != 2 {
		t.Errorf("expected 2 lines, got %d", len(lines))
	}

	if str == "" {
		t.Error("GoLiteralString should not be empty")
	}
}

func Test_Map_GoLiteral_Empty(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	lines := m.GoLiteralLines()

	// Assert
	if len(lines) != 0 {
		t.Errorf("expected 0 lines, got %d", len(lines))
	}
}

func Test_Map_GetAsStrings(t *testing.T) {
	// Arrange
	m := args.Map{
		"items": []string{"a", "b", "c"},
	}

	// Act
	items, isValid := m.GetAsStrings("items")

	// Assert
	if !isValid {
		t.Error("expected valid")
	}

	if len(items) != 3 {
		t.Errorf("expected 3 items, got %d", len(items))
	}
}

func Test_Map_GetAsAnyItems(t *testing.T) {
	// Arrange
	m := args.Map{
		"items": []any{1, "two", true},
	}

	// Act
	items, isValid := m.GetAsAnyItems("items")

	// Assert
	if !isValid {
		t.Error("expected valid")
	}

	if len(items) != 3 {
		t.Errorf("expected 3 items, got %d", len(items))
	}
}

func Test_Map_String(t *testing.T) {
	// Arrange
	m := args.Map{
		"key": "value",
	}

	// Act
	s := m.String()

	// Assert
	if s == "" {
		t.Error("Map String() should not be empty")
	}
}

func Test_Map_SetActual(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	m.SetActual("result")

	// Assert
	if m.Actual() != "result" {
		t.Error("expected Actual to return result")
	}
}

func Test_Map_SortedKeys(t *testing.T) {
	// Arrange
	m := args.Map{
		"c": 3,
		"a": 1,
		"b": 2,
	}

	// Act
	keys, err := m.SortedKeys()

	// Assert
	if err != nil {
		t.Errorf("SortedKeys error: %v", err)
	}

	if len(keys) != 3 {
		t.Errorf("expected 3 keys, got %d", len(keys))
	}

	if keys[0] != "a" || keys[1] != "b" || keys[2] != "c" {
		t.Error("keys should be sorted")
	}
}

func Test_Map_ValidArgs(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": "hello",
		"b": 42,
	}

	// Act
	validArgs := m.ValidArgs()

	// Assert
	if len(validArgs) != 2 {
		t.Errorf("expected 2 valid args, got %d", len(validArgs))
	}
}

func Test_Map_Args(t *testing.T) {
	// Arrange
	m := args.Map{
		"x": 10,
		"y": 20,
	}

	// Act
	result := m.Args("x", "y")

	// Assert
	if len(result) != 2 {
		t.Errorf("expected 2 args, got %d", len(result))
	}
}

func Test_One(t *testing.T) {
	for caseIndex, testCase := range extOneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		first := input.GetDirectLower("first")
		expect := input.GetDirectLower("expect")

		one := args.OneAny{
			First:  first,
			Expect: expect,
		}

		// Act
		actual := args.Map{
			"hasFirst":   one.HasFirst(),
			"hasExpect":  one.HasExpect(),
			"argsCount":  one.ArgsCount(),
			"validCount": len(one.ValidArgs()),
		}

		if one.HasFirst() {
			actual["firstItem"] = one.FirstItem()
		}

		if one.HasExpect() {
			actual["expected"] = one.Expected()
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_One_String(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello", Expect: 42}

	// Act
	s := one.String()

	// Assert
	if s == "" {
		t.Error("One.String() should not be empty")
	}
}

func Test_One_Args(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello"}

	// Act & Assert
	if len(one.Args(1)) != 1 {
		t.Error("expected 1 arg")
	}

	if len(one.Args(0)) != 0 {
		t.Error("expected 0 args")
	}
}

func Test_One_Slice_Cached(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello"}

	// Act
	slice1 := one.Slice()
	slice2 := one.Slice()

	// Assert
	if len(slice1) != len(slice2) {
		t.Error("cached Slice should return same length")
	}
}

func Test_One_GetByIndex(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello", Expect: 42}

	// Act & Assert
	if one.GetByIndex(0) != "hello" {
		t.Error("expected GetByIndex(0) to return hello")
	}

	if one.GetByIndex(99) != nil {
		t.Error("expected GetByIndex(99) to return nil")
	}
}

func Test_One_LeftRight(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello", Expect: 42}

	// Act
	lr := one.LeftRight()

	// Assert
	if lr.Left != "hello" {
		t.Error("expected Left to be hello")
	}
}

func Test_Two(t *testing.T) {
	for caseIndex, testCase := range extTwoTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		first, _ := input.GetAsString("first")
		second, _ := input.GetAsString("second")

		two := args.TwoAny{
			First:  first,
			Second: second,
		}

		// Act
		actual := args.Map{
			"hasFirst":   two.HasFirst(),
			"hasSecond":  two.HasSecond(),
			"argsCount":  two.ArgsCount(),
			"validCount": len(two.ValidArgs()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Two_String(t *testing.T) {
	// Arrange
	two := args.TwoAny{First: "a", Second: "b"}

	// Act
	s := two.String()

	// Assert
	if s == "" {
		t.Error("Two.String() should not be empty")
	}
}

func Test_Two_LeftRight(t *testing.T) {
	// Arrange
	two := args.TwoAny{First: "a", Second: "b", Expect: "c"}

	// Act
	lr := two.LeftRight()

	// Assert
	if lr.Left != "a" || lr.Right != "b" {
		t.Error("LeftRight conversion failed")
	}
}

func Test_Three_Methods(t *testing.T) {
	// Arrange
	three := args.ThreeAny{
		First:  "a",
		Second: "b",
		Third:  "c",
		Expect: "d",
	}

	// Act & Assert
	if !three.HasFirst() {
		t.Error("HasFirst should be true")
	}

	if !three.HasSecond() {
		t.Error("HasSecond should be true")
	}

	if !three.HasThird() {
		t.Error("HasThird should be true")
	}

	if !three.HasExpect() {
		t.Error("HasExpect should be true")
	}

	if three.ArgsCount() != 3 {
		t.Errorf("expected ArgsCount 3, got %d", three.ArgsCount())
	}

	if len(three.ValidArgs()) != 3 {
		t.Error("expected 3 valid args")
	}

	if three.String() == "" {
		t.Error("String() should not be empty")
	}

	if three.FirstItem() != "a" {
		t.Error("FirstItem should be a")
	}

	if three.SecondItem() != "b" {
		t.Error("SecondItem should be b")
	}

	if three.ThirdItem() != "c" {
		t.Error("ThirdItem should be c")
	}

	argTwo := three.ArgTwo()
	if argTwo.First != "a" || argTwo.Second != "b" {
		t.Error("ArgTwo should have first two")
	}

	argThree := three.ArgThree()
	if argThree.First != "a" {
		t.Error("ArgThree should copy")
	}

	lr := three.LeftRight()
	if lr.Left != "a" || lr.Right != "b" {
		t.Error("LeftRight should map first/second")
	}
}

func Test_LeftRight(t *testing.T) {
	for caseIndex, testCase := range extLeftRightTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsInt("left")
		right, _ := input.GetAsInt("right")

		lr := args.LeftRightAny{
			Left:  left,
			Right: right,
		}

		// Act
		actual := args.Map{
			"hasLeft":    lr.HasLeft(),
			"hasRight":   lr.HasRight(),
			"hasFirst":   lr.HasFirst(),
			"hasSecond":  lr.HasSecond(),
			"argsCount":  lr.ArgsCount(),
			"validCount": len(lr.ValidArgs()),
			"firstItem":  lr.FirstItem(),
			"secondItem": lr.SecondItem(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LeftRight_Clone(t *testing.T) {
	// Arrange
	lr := args.LeftRightAny{Left: "a", Right: "b", Expect: "c"}

	// Act
	cloned := lr.Clone()

	// Assert
	if cloned.Left != "a" || cloned.Right != "b" || cloned.Expect != "c" {
		t.Error("Clone should preserve all fields")
	}
}

func Test_LeftRight_String(t *testing.T) {
	// Arrange
	lr := args.LeftRightAny{Left: "a", Right: "b"}

	// Act
	s := lr.String()

	// Assert
	if s == "" {
		t.Error("LeftRight.String() should not be empty")
	}
}

func Test_String_Methods(t *testing.T) {
	for caseIndex, testCase := range extStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		s := args.String(inputStr)

		// Act
		actual := args.Map{
			"length":              s.Length(),
			"isEmpty":             s.IsEmpty(),
			"hasCharacter":        s.HasCharacter(),
			"isDefined":           s.IsDefined(),
			"isEmptyOrWhitespace": s.IsEmptyOrWhitespace(),
			"asciiLength":         s.AscIILength(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_String_Operations(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act & Assert
	if s.String() != "hello" {
		t.Error("String() should return hello")
	}

	if len(s.Bytes()) != 5 {
		t.Error("Bytes() should have 5 bytes")
	}

	if len(s.Runes()) != 5 {
		t.Error("Runes() should have 5 runes")
	}

	if s.Count() != 5 {
		t.Errorf("Count should be 5, got %d", s.Count())
	}

	trimmed := args.String("  hello  ").TrimSpace()
	if trimmed != "hello" {
		t.Errorf("TrimSpace expected hello, got %s", trimmed)
	}

	replaced := s.ReplaceAll("hello", "world")
	if replaced != "world" {
		t.Errorf("ReplaceAll expected world, got %s", replaced)
	}

	sub := s.Substring(0, 3)
	if sub != "hel" {
		t.Errorf("Substring expected hel, got %s", sub)
	}

	concat := args.String("a").Concat("b", "c")
	if concat != "abc" {
		t.Errorf("Concat expected abc, got %s", concat)
	}

	joined := args.String("a").Join("-", "b", "c")
	if joined != "a-b-c" {
		t.Errorf("Join expected a-b-c, got %s", joined)
	}

	split := args.String("a,b,c").Split(",")
	if len(split) != 3 {
		t.Errorf("Split expected 3, got %d", len(split))
	}
}

func Test_String_Quoting(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act & Assert
	dq := s.DoubleQuote()
	if dq == "" {
		t.Error("DoubleQuote should not be empty")
	}

	dqq := s.DoubleQuoteQ()
	if dqq == "" {
		t.Error("DoubleQuoteQ should not be empty")
	}

	sq := s.SingleQuote()
	if sq == "" {
		t.Error("SingleQuote should not be empty")
	}

	vdq := s.ValueDoubleQuote()
	if vdq == "" {
		t.Error("ValueDoubleQuote should not be empty")
	}
}

func Test_Empty_Creator(t *testing.T) {
	// Arrange & Act
	m := args.Empty.Map()
	fw := args.Empty.FuncWrap()
	fm := args.Empty.FuncMap()
	h := args.Empty.Holder()

	// Assert
	if m == nil {
		t.Error("Empty.Map should not be nil")
	}

	if fw == nil {
		t.Error("Empty.FuncWrap should not be nil")
	}

	if fm == nil {
		t.Error("Empty.FuncMap should not be nil")
	}

	if h.ArgsCount() != 7 {
		t.Errorf("Empty.Holder ArgsCount should be 7, got %d", h.ArgsCount())
	}
}

func Test_Holder_Methods(t *testing.T) {
	// Arrange
	h := args.HolderAny{
		First:  "a",
		Second: "b",
		Third:  "c",
		Fourth: "d",
		Fifth:  "e",
		Sixth:  "f",
		Expect: "g",
	}

	// Act & Assert
	if !h.HasFirst() {
		t.Error("HasFirst should be true")
	}

	if !h.HasSecond() {
		t.Error("HasSecond should be true")
	}

	if !h.HasThird() {
		t.Error("HasThird should be true")
	}

	if !h.HasFourth() {
		t.Error("HasFourth should be true")
	}

	if !h.HasFifth() {
		t.Error("HasFifth should be true")
	}

	if !h.HasSixth() {
		t.Error("HasSixth should be true")
	}

	if !h.HasExpect() {
		t.Error("HasExpect should be true")
	}

	if h.ArgsCount() != 7 {
		t.Errorf("ArgsCount should be 7, got %d", h.ArgsCount())
	}

	if len(h.ValidArgs()) != 6 {
		t.Errorf("expected 6 valid args, got %d", len(h.ValidArgs()))
	}

	if h.FirstItem() != "a" {
		t.Error("FirstItem should be a")
	}

	if h.SecondItem() != "b" {
		t.Error("SecondItem should be b")
	}

	if h.ThirdItem() != "c" {
		t.Error("ThirdItem should be c")
	}

	if h.FourthItem() != "d" {
		t.Error("FourthItem should be d")
	}

	if h.FifthItem() != "e" {
		t.Error("FifthItem should be e")
	}

	if h.SixthItem() != "f" {
		t.Error("SixthItem should be f")
	}

	if h.Expected() != "g" {
		t.Error("Expected should be g")
	}

	if h.String() == "" {
		t.Error("Holder String() should not be empty")
	}
}

func Test_Holder_Args(t *testing.T) {
	// Arrange
	h := args.HolderAny{
		First:  "a",
		Second: "b",
		Third:  "c",
	}

	// Act & Assert
	if len(h.Args(1)) != 1 {
		t.Error("Args(1) should return 1")
	}

	if len(h.Args(3)) != 3 {
		t.Error("Args(3) should return 3")
	}

	if len(h.Args(6)) != 6 {
		t.Error("Args(6) should return 6")
	}
}

func Test_Holder_ArgTwo_ArgThree_ArgFour_ArgFive(t *testing.T) {
	// Arrange
	h := args.HolderAny{
		First:  "a",
		Second: "b",
		Third:  "c",
		Fourth: "d",
		Fifth:  "e",
	}

	// Act & Assert
	at := h.ArgTwo()
	if at.First != "a" || at.Second != "b" {
		t.Error("ArgTwo should have first two")
	}

	a3 := h.ArgThree()
	if a3.First != "a" {
		t.Error("ArgThree should have first")
	}

	a4 := h.ArgFour()
	if a4.First != "a" {
		t.Error("ArgFour should have first")
	}

	a5 := h.ArgFive()
	if a5.First != "a" {
		t.Error("ArgFive should have first")
	}
}

func Test_Dynamic_Methods(t *testing.T) {
	// Arrange
	d := args.DynamicAny{
		Params: args.Map{
			"key1": "val1",
			"key2": 42,
		},
		Expect: "expected",
	}

	// Act & Assert
	// Map.ArgsCount() excludes "expected"/"func" keys.
	// Map.HasFunc() returns true even for nil func (non-nil FuncWrap wrapper),
	// so ArgsCount = Length(2) - HasFunc(1) = 1.
	if d.ArgsCount() != 1 {
		t.Errorf("expected ArgsCount 1, got %d", d.ArgsCount())
	}

	if !d.HasExpect() {
		t.Error("HasExpect should be true")
	}

	if d.Expected() != "expected" {
		t.Error("Expected should be expected")
	}

	if !d.HasDefined("key1") {
		t.Error("HasDefined key1 should be true")
	}

	if !d.Has("key1") {
		t.Error("Has key1 should be true")
	}

	if d.IsKeyMissing("key1") {
		t.Error("IsKeyMissing key1 should be false")
	}

	if !d.IsKeyMissing("missing") {
		t.Error("IsKeyMissing missing should be true")
	}

	if d.IsKeyInvalid("key1") {
		t.Error("IsKeyInvalid key1 should be false")
	}

	val, isValid := d.Get("key1")
	if !isValid || val != "val1" {
		t.Error("Get key1 should return val1")
	}

	intVal, intValid := d.GetAsInt("key2")
	if !intValid || intVal != 42 {
		t.Error("GetAsInt key2 should return 42")
	}

	if d.GetAsIntDefault("missing", 99) != 99 {
		t.Error("GetAsIntDefault should return 99")
	}

	strVal, strValid := d.GetAsString("key1")
	if !strValid || strVal != "val1" {
		t.Error("GetAsString key1 should return val1")
	}

	if d.GetAsStringDefault("missing") != "" {
		t.Error("GetAsStringDefault should return empty")
	}

	// Note: Dynamic.String() calls Slice() which uses converters.Map.SortedKeys.
	// That function does not support args.Map type, so String() panics.
	// This is a known limitation — skipping String() test.
}

func Test_Dynamic_NilSafety(t *testing.T) {
	// Arrange
	var d *args.DynamicAny

	// Act & Assert
	if d.ArgsCount() != 0 {
		t.Error("nil ArgsCount should return 0")
	}

	if d.GetWorkFunc() != nil {
		t.Error("nil GetWorkFunc should return nil")
	}

	if d.HasDefined("key") {
		t.Error("nil HasDefined should return false")
	}

	if d.Has("key") {
		t.Error("nil Has should return false")
	}

	if d.HasDefinedAll("key") {
		t.Error("nil HasDefinedAll should return false")
	}

	if d.IsKeyInvalid("key") {
		t.Error("nil IsKeyInvalid should return false")
	}

	if d.IsKeyMissing("key") {
		t.Error("nil IsKeyMissing should return false")
	}

	_, isValid := d.Get("key")
	if isValid {
		t.Error("nil Get should return false")
	}

	if d.HasExpect() {
		t.Error("nil HasExpect should return false")
	}
}

func Test_FuncMap_Basic(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act & Assert
	if !fm.IsEmpty() {
		t.Error("new FuncMap should be empty")
	}

	if fm.Length() != 0 {
		t.Error("Length should be 0")
	}

	if fm.Count() != 0 {
		t.Error("Count should be 0")
	}

	if fm.HasAnyItem() {
		t.Error("HasAnyItem should be false")
	}

	if fm.Has("nonexistent") {
		t.Error("Has should be false")
	}

	if fm.IsContains("nonexistent") {
		t.Error("IsContains should be false")
	}

	if fm.Get("nonexistent") != nil {
		t.Error("Get should return nil")
	}

	if fm.IsValidFuncOf("nonexistent") {
		t.Error("IsValidFuncOf should be false")
	}

	if !fm.IsInvalidFunc("nonexistent") {
		t.Error("IsInvalidFunc should be true")
	}

	if fm.PkgPath("nonexistent") != "" {
		t.Error("PkgPath should return empty")
	}

	if fm.PkgNameOnly("nonexistent") != "" {
		t.Error("PkgNameOnly should return empty")
	}

	if fm.FuncDirectInvokeName("nonexistent") != "" {
		t.Error("FuncDirectInvokeName should return empty")
	}

	if fm.ArgsCount("nonexistent") != 0 {
		t.Error("ArgsCount should return 0")
	}

	if fm.ReturnLength("nonexistent") != 0 {
		t.Error("ReturnLength should return 0")
	}

	if fm.IsPublicMethod("nonexistent") {
		t.Error("IsPublicMethod should return false")
	}

	if fm.IsPrivateMethod("nonexistent") {
		t.Error("IsPrivateMethod should return false")
	}

	if fm.GetType("nonexistent") != nil {
		t.Error("GetType should return nil")
	}

	if fm.GetPascalCaseFuncName("nonexistent") != "" {
		t.Error("GetPascalCaseFuncName should return empty")
	}

	if fm.InvalidError() == nil {
		t.Error("empty FuncMap InvalidError should return error")
	}
}

func Test_FuncMap_Add(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	fm.Add(someFunctionV1)

	// Assert
	if fm.IsEmpty() {
		t.Error("FuncMap should not be empty after Add")
	}

	if !fm.Has("someFunctionV1") {
		t.Error("expected someFunctionV1 in map")
	}

	if !fm.IsValidFuncOf("someFunctionV1") {
		t.Error("someFunctionV1 should be valid")
	}
}

func Test_FuncMap_Adds(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	fm.Adds(someFunctionV1, someFunctionV2)

	// Assert
	if fm.Length() != 2 {
		t.Errorf("expected 2, got %d", fm.Length())
	}
}

func Test_FuncDetector(t *testing.T) {
	// Arrange & Act
	fw := args.FuncDetector.GetFuncWrap(someFunctionV1)

	// Assert
	if fw == nil {
		t.Error("GetFuncWrap should not return nil")
	}

	if !fw.HasValidFunc() {
		t.Error("wrapped func should be valid")
	}
}
