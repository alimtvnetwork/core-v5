package coretestcasestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ── CaseV1 basic getters ──

func Test_Cov2_CaseV1_Input(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "test title",
		ArrangeInput:  "input",
		ExpectedInput: "expected",
	}

	actual := args.Map{
		"input":     c.Input(),
		"expected":  c.Expected(),
		"title":     c.CaseTitle(),
		"typeName":  c.ArrangeTypeName(),
	}
	expected := args.Map{
		"input":     "input",
		"expected":  "expected",
		"title":     "test title",
		"typeName":  "string",
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_Input", actual)
}

func Test_Cov2_CaseV1_ExpectedLines_String(t *testing.T) {
	c := coretestcases.CaseV1{
		ExpectedInput: "hello",
	}

	lines := c.ExpectedLines()

	actual := args.Map{
		"len":   len(lines),
		"first": lines[0],
	}
	expected := args.Map{
		"len":   1,
		"first": "hello",
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_ExpectedLines_String", actual)
}

func Test_Cov2_CaseV1_ExpectedLines_Slice(t *testing.T) {
	c := coretestcases.CaseV1{
		ExpectedInput: []string{"a", "b"},
	}

	lines := c.ExpectedLines()

	actual := args.Map{
		"len": len(lines),
	}
	expected := args.Map{
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_ExpectedLines_Slice", actual)
}

func Test_Cov2_CaseV1_SetActual(t *testing.T) {
	c := coretestcases.CaseV1{Title: "test", ActualInput: "result"}

	actual := args.Map{
		"actual": c.Actual(),
	}
	expected := args.Map{
		"actual": "result",
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_SetActual", actual)
}

func Test_Cov2_CaseV1_AsSimpleTestCaseWrapper(t *testing.T) {
	c := coretestcases.CaseV1{Title: "test"}
	wrapper := c.AsSimpleTestCaseWrapper()

	actual := args.Map{
		"notNil": wrapper != nil,
	}
	expected := args.Map{
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_AsSimpleTestCaseWrapper", actual)
}

// ── CaseV1 ShouldBeEqual ──

func Test_Cov2_CaseV1_ShouldBeEqual(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeEqual test",
		ExpectedInput: "hello",
	}

	c.ShouldBeEqual(t, 0, "hello")
}

func Test_Cov2_CaseV1_ShouldBeEqualFirst(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeEqualFirst test",
		ExpectedInput: "hello",
	}

	c.ShouldBeEqualFirst(t, "hello")
}

func Test_Cov2_CaseV1_ShouldBeTrimEqual(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeTrimEqual test",
		ExpectedInput: "hello",
	}

	c.ShouldBeTrimEqual(t, 0, "hello")
}

func Test_Cov2_CaseV1_ShouldBeTrimEqualFirst(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeTrimEqualFirst test",
		ExpectedInput: "hello",
	}

	c.ShouldBeTrimEqualFirst(t, "hello")
}

func Test_Cov2_CaseV1_ShouldBeSortedEqual(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeSortedEqual test",
		ExpectedInput: []string{"a", "b"},
	}

	c.ShouldBeSortedEqual(t, 0, "a", "b")
}

func Test_Cov2_CaseV1_ShouldBeSortedEqualFirst(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeSortedEqualFirst test",
		ExpectedInput: []string{"a", "b"},
	}

	c.ShouldBeSortedEqualFirst(t, "a", "b")
}

func Test_Cov2_CaseV1_ShouldContains(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldContains test",
		ExpectedInput: "hello",
	}

	c.ShouldContains(t, 0, "hello world")
}

func Test_Cov2_CaseV1_ShouldContainsFirst(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldContainsFirst test",
		ExpectedInput: "hello",
	}

	c.ShouldContainsFirst(t, "hello world")
}

func Test_Cov2_CaseV1_ShouldStartsWith(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldStartsWith test",
		ExpectedInput: "hello",
	}

	c.ShouldStartsWith(t, 0, "hello world")
}

func Test_Cov2_CaseV1_ShouldStartsWithFirst(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldStartsWithFirst test",
		ExpectedInput: "hello",
	}

	c.ShouldStartsWithFirst(t, "hello world")
}

func Test_Cov2_CaseV1_ShouldEndsWith(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldEndsWith test",
		ExpectedInput: "world",
	}

	c.ShouldEndsWith(t, 0, "hello world")
}

func Test_Cov2_CaseV1_ShouldEndsWithFirst(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldEndsWithFirst test",
		ExpectedInput: "world",
	}

	c.ShouldEndsWithFirst(t, "hello world")
}

func Test_Cov2_CaseV1_ShouldBeNotEqual(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeNotEqual test",
		ExpectedInput: "hello",
	}

	c.ShouldBeNotEqual(t, 0, "world")
}

func Test_Cov2_CaseV1_ShouldBeNotEqualFirst(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeNotEqualFirst test",
		ExpectedInput: "hello",
	}

	c.ShouldBeNotEqualFirst(t, "world")
}

// ── CaseV1 VerifyAll / VerifyError ──

func Test_Cov2_CaseV1_VerifyAllEqual(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "VerifyAllEqual test",
		ExpectedInput: "hello",
	}

	err := c.VerifyAllEqual(0, "hello")

	actual := args.Map{
		"noErr": err == nil,
	}
	expected := args.Map{
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_VerifyAllEqual", actual)
}

func Test_Cov2_CaseV1_VerifyError(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "VerifyError test",
		ExpectedInput: "hello",
	}

	err := c.VerifyError(0, stringcompareas.Equal, "hello")

	actual := args.Map{
		"noErr": err == nil,
	}
	expected := args.Map{
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_VerifyError", actual)
}

func Test_Cov2_CaseV1_VerifyFirst(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "VerifyFirst test",
		ExpectedInput: "hello",
	}

	err := c.VerifyFirst(0, stringcompareas.Equal, []string{"hello"})

	actual := args.Map{
		"noErr": err == nil,
	}
	expected := args.Map{
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_VerifyFirst", actual)
}

func Test_Cov2_CaseV1_SliceValidator(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "SliceValidator test",
		ExpectedInput: "hello",
	}

	sv := c.SliceValidator(stringcompareas.Equal, []string{"hello"})

	actual := args.Map{
		"hasActual":   len(sv.ActualLines) > 0,
		"hasExpected": len(sv.ExpectedLines) > 0,
	}
	expected := args.Map{
		"hasActual":   true,
		"hasExpected": true,
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_SliceValidator", actual)
}

// ── CaseV1 Map Assertions ──

func Test_Cov2_CaseV1_ShouldBeEqualMap(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeEqualMap test",
		ExpectedInput: args.Map{"key": "value"},
	}

	c.ShouldBeEqualMap(t, 0, args.Map{"key": "value"})
}

func Test_Cov2_CaseV1_ShouldBeEqualMapFirst(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeEqualMapFirst test",
		ExpectedInput: args.Map{"key": "value"},
	}

	c.ShouldBeEqualMapFirst(t, args.Map{"key": "value"})
}

func Test_Cov2_CaseV1_ExpectedAsMap(t *testing.T) {
	c := coretestcases.CaseV1{
		ExpectedInput: args.Map{"key": "value"},
	}

	m := c.ExpectedAsMap()

	actual := args.Map{
		"len": len(m),
	}
	expected := args.Map{
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_ExpectedAsMap", actual)
}

// ── GenericGherkins CompareWith ──

func Test_Cov2_GenericGherkins_CompareWith_Equal(t *testing.T) {
	g1 := &coretestcases.StringBoolGherkins{Title: "t", Feature: "f", When: "w"}
	g2 := &coretestcases.StringBoolGherkins{Title: "t", Feature: "f", When: "w"}

	isEqual, diff := g1.CompareWith(g2)

	actual := args.Map{
		"isEqual":   isEqual,
		"diffEmpty": diff == "",
	}
	expected := args.Map{
		"isEqual":   true,
		"diffEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "GenericGherkins_CompareWith_Equal", actual)
}

func Test_Cov2_GenericGherkins_CompareWith_Diff(t *testing.T) {
	g1 := &coretestcases.StringBoolGherkins{Title: "a"}
	g2 := &coretestcases.StringBoolGherkins{Title: "b"}

	isEqual, diff := g1.CompareWith(g2)

	actual := args.Map{
		"isEqual":    isEqual,
		"hasDiff":    diff != "",
	}
	expected := args.Map{
		"isEqual":    false,
		"hasDiff":    true,
	}
	expected.ShouldBeEqual(t, 0, "GenericGherkins_CompareWith_Diff", actual)
}

func Test_Cov2_GenericGherkins_CompareWith_BothNil(t *testing.T) {
	var g1, g2 *coretestcases.StringBoolGherkins

	isEqual, _ := g1.CompareWith(g2)

	actual := args.Map{
		"isEqual": isEqual,
	}
	expected := args.Map{
		"isEqual": true,
	}
	expected.ShouldBeEqual(t, 0, "GenericGherkins_CompareWith_BothNil", actual)
}

func Test_Cov2_GenericGherkins_CompareWith_OneNil(t *testing.T) {
	g1 := &coretestcases.StringBoolGherkins{Title: "a"}
	var g2 *coretestcases.StringBoolGherkins

	isEqual, diff := g1.CompareWith(g2)

	actual := args.Map{
		"isEqual": isEqual,
		"hasDiff": diff != "",
	}
	expected := args.Map{
		"isEqual": false,
		"hasDiff": true,
	}
	expected.ShouldBeEqual(t, 0, "GenericGherkins_CompareWith_OneNil", actual)
}

// ── GenericGherkins Typed Assertions ──

func Test_Cov2_GenericGherkins_ShouldMatchExpected(t *testing.T) {
	g := &coretestcases.StringBoolGherkins{
		Title:    "ShouldMatchExpected test",
		Expected: true,
	}

	g.ShouldMatchExpected(t, 0, true)
}

func Test_Cov2_GenericGherkins_ShouldMatchExpectedFirst(t *testing.T) {
	g := &coretestcases.StringBoolGherkins{
		Title:    "ShouldMatchExpectedFirst test",
		Expected: true,
	}

	g.ShouldMatchExpectedFirst(t, true)
}

// ── GenericGherkins TypedWrapper ──

func Test_Cov2_GenericGherkins_TypedWrapper(t *testing.T) {
	g := &coretestcases.StringBoolGherkins{
		Title:    "wrapper",
		Input:    "input",
		Expected: true,
	}

	g.SetTypedActual(false)

	actual := args.Map{
		"caseTitle":    g.CaseTitle(),
		"typedInput":   g.TypedInput(),
		"typedExpect":  g.TypedExpected(),
		"typedActual":  g.TypedActual(),
		"wrapperNotNil": g.AsTypedTestCaseWrapper() != nil,
	}
	expected := args.Map{
		"caseTitle":    "wrapper",
		"typedInput":   "input",
		"typedExpect":  true,
		"typedActual":  false,
		"wrapperNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "GenericGherkins_TypedWrapper", actual)
}

// ── CaseNilSafe ──

func Test_Cov2_CaseNilSafe_CaseTitle_Empty(t *testing.T) {
	c := coretestcases.CaseNilSafe{
		Func: (*testing.T).Name,
	}

	title := c.CaseTitle()

	actual := args.Map{
		"hasTitle": title != "",
	}
	expected := args.Map{
		"hasTitle": true,
	}
	expected.ShouldBeEqual(t, 0, "CaseNilSafe_CaseTitle_Empty", actual)
}

func Test_Cov2_CaseNilSafe_MethodName(t *testing.T) {
	c := coretestcases.CaseNilSafe{
		Title: "explicit title",
		Func:  (*testing.T).Name,
	}

	actual := args.Map{
		"title":      c.CaseTitle(),
		"methodName": c.MethodName() != "",
	}
	expected := args.Map{
		"title":      "explicit title",
		"methodName": true,
	}
	expected.ShouldBeEqual(t, 0, "CaseNilSafe_MethodName", actual)
}
