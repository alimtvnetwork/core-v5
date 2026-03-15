package stringutiltests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

// ── IsEmpty / IsNotEmpty / IsDefined / IsBlank ──

func Test_Cov4_IsEmpty(t *testing.T) {
	actual := args.Map{"empty": stringutil.IsEmpty(""), "notEmpty": stringutil.IsEmpty("x")}
	expected := args.Map{"empty": true, "notEmpty": false}
	expected.ShouldBeEqual(t, 0, "IsEmpty", actual)
}

func Test_Cov4_IsNotEmpty(t *testing.T) {
	actual := args.Map{"notEmpty": stringutil.IsNotEmpty("x"), "empty": stringutil.IsNotEmpty("")}
	expected := args.Map{"notEmpty": true, "empty": false}
	expected.ShouldBeEqual(t, 0, "IsNotEmpty", actual)
}

func Test_Cov4_IsEmptyPtr(t *testing.T) {
	empty := ""
	text := "hello"
	actual := args.Map{
		"nil":   stringutil.IsEmptyPtr(nil),
		"empty": stringutil.IsEmptyPtr(&empty),
		"text":  stringutil.IsEmptyPtr(&text),
	}
	expected := args.Map{"nil": true, "empty": true, "text": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr", actual)
}

func Test_Cov4_IsNullOrEmptyPtr(t *testing.T) {
	empty := ""
	text := "hello"
	actual := args.Map{
		"nil":   stringutil.IsNullOrEmptyPtr(nil),
		"empty": stringutil.IsNullOrEmptyPtr(&empty),
		"text":  stringutil.IsNullOrEmptyPtr(&text),
	}
	expected := args.Map{"nil": true, "empty": true, "text": false}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmptyPtr", actual)
}

func Test_Cov4_IsBlank(t *testing.T) {
	actual := args.Map{
		"empty":  stringutil.IsBlank(""),
		"space":  stringutil.IsBlank(" "),
		"nl":     stringutil.IsBlank("\n"),
		"text":   stringutil.IsBlank("x"),
		"tabs":   stringutil.IsBlank("\t  "),
	}
	expected := args.Map{"empty": true, "space": true, "nl": true, "text": false, "tabs": true}
	expected.ShouldBeEqual(t, 0, "IsBlank", actual)
}

func Test_Cov4_IsEmptyOrWhitespace(t *testing.T) {
	actual := args.Map{
		"empty": stringutil.IsEmptyOrWhitespace(""),
		"space": stringutil.IsEmptyOrWhitespace("  "),
		"text":  stringutil.IsEmptyOrWhitespace("x"),
	}
	expected := args.Map{"empty": true, "space": true, "text": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespace", actual)
}

func Test_Cov4_IsContains(t *testing.T) {
	lines := []string{"hello", "world"}
	actual := args.Map{
		"found":    stringutil.IsContains(lines, "world", 0, true),
		"notFound": stringutil.IsContains(lines, "foo", 0, true),
	}
	expected := args.Map{"found": true, "notFound": false}
	expected.ShouldBeEqual(t, 0, "IsContains", actual)
}

// ── AnyToString ──

func Test_Cov4_AnyToString(t *testing.T) {
	actual := args.Map{
		"nil":    stringutil.AnyToString(nil),
		"string": stringutil.AnyToString("hello"),
		"int":    stringutil.AnyToString(42) != "",
	}
	expected := args.Map{"nil": "", "string": "hello", "int": true}
	expected.ShouldBeEqual(t, 0, "AnyToString", actual)
}

// ── SplitLeftRightType / SplitLeftRightTypeTrimmed / SplitLeftRightsTrims ──

func Test_Cov4_SplitLeftRightType(t *testing.T) {
	result := stringutil.SplitLeftRightType("key=value", "=")
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightType", actual)
}

func Test_Cov4_SplitLeftRightTypeTrimmed(t *testing.T) {
	result := stringutil.SplitLeftRightTypeTrimmed(" key = value ", "=")
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTypeTrimmed", actual)
}

func Test_Cov4_SplitLeftRightsTrims(t *testing.T) {
	result := stringutil.SplitLeftRightsTrims("=", "a=1", "b=2")
	emptyResult := stringutil.SplitLeftRightsTrims("=")
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightsTrims", actual)
}

// ── SplitContentsByWhitespaceConditions ──

func Test_Cov4_SplitContentsByWhitespaceConditions(t *testing.T) {
	result1 := stringutil.SplitContentsByWhitespaceConditions("hello world", true, true, true, false, false)
	result2 := stringutil.SplitContentsByWhitespaceConditions("hello world hello", false, true, false, true, true)
	result3 := stringutil.SplitContentsByWhitespaceConditions("  a  b  ", false, false, false, false, false)
	actual := args.Map{
		"trimSorted":    len(result1) > 0,
		"uniqueLower":   len(result2) > 0,
		"noFlags":       len(result3) > 0,
	}
	expected := args.Map{"trimSorted": true, "uniqueLower": true, "noFlags": true}
	expected.ShouldBeEqual(t, 0, "SplitContentsByWhitespaceConditions", actual)
}

// ── ToIntUsingRegexMatch ──

func Test_Cov4_ToIntUsingRegexMatch(t *testing.T) {
	re := regexp.MustCompile(`^\d+$`)
	actual := args.Map{
		"valid":    stringutil.ToIntUsingRegexMatch(re, "42"),
		"invalid":  stringutil.ToIntUsingRegexMatch(re, "abc"),
		"nilRegex": stringutil.ToIntUsingRegexMatch(nil, "42"),
	}
	expected := args.Map{"valid": 42, "invalid": 0, "nilRegex": 0}
	expected.ShouldBeEqual(t, 0, "ToIntUsingRegexMatch", actual)
}

// ── ReplaceTemplate ──

func Test_Cov4_ReplaceTemplate_CurlyOne(t *testing.T) {
	result := stringutil.ReplaceWhiteSpacesToSingle("Hello  World   !")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "Hello World !"}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle", actual)
}

func Test_Cov4_ReplaceTemplate_CurlyTwo(t *testing.T) {
	result := stringutil.Replace.CurlyTwo("{a} and {b}", "a", 1, "b", 2)
	empty := stringutil.Replace.CurlyTwo("", "a", 1, "b", 2)
	actual := args.Map{"notEmpty": result != "", "empty": empty}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "CurlyTwo", actual)
}

func Test_Cov4_ReplaceTemplate_Curly(t *testing.T) {
	result := stringutil.Replace.Curly("{x}", map[string]string{"x": "val"})
	empty := stringutil.Replace.Curly("", map[string]string{"x": "val"})
	actual := args.Map{"result": result, "empty": empty}
	expected := args.Map{"result": "val", "empty": ""}
	expected.ShouldBeEqual(t, 0, "Curly", actual)
}

func Test_Cov4_ReplaceTemplate_DirectOne(t *testing.T) {
	result := stringutil.Replace.DirectOne("Hello NAME!", "NAME", "World")
	empty := stringutil.Replace.DirectOne("", "NAME", "World")
	actual := args.Map{"result": result, "empty": empty}
	expected := args.Map{"result": "Hello World!", "empty": ""}
	expected.ShouldBeEqual(t, 0, "DirectOne", actual)
}

func Test_Cov4_ReplaceTemplate_DirectTwoItem(t *testing.T) {
	result := stringutil.Replace.DirectTwoItem("A and B", "A", 1, "B", 2)
	empty := stringutil.Replace.DirectTwoItem("", "A", 1, "B", 2)
	actual := args.Map{"notEmpty": result != "", "empty": empty}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "DirectTwoItem", actual)
}

func Test_Cov4_ReplaceTemplate_CurlyTwoItem(t *testing.T) {
	result := stringutil.Replace.CurlyTwoItem("{a} and {b}", "a", 1, "b", 2)
	empty := stringutil.Replace.CurlyTwoItem("", "a", 1, "b", 2)
	actual := args.Map{"notEmpty": result != "", "empty": empty}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "CurlyTwoItem", actual)
}

func Test_Cov4_ReplaceTemplate_DirectKeyUsingMap(t *testing.T) {
	result := stringutil.Replace.DirectKeyUsingMap("KEY=val", map[string]string{"KEY": "replaced"})
	emptyMap := stringutil.Replace.DirectKeyUsingMap("text", map[string]string{})
	emptyFmt := stringutil.Replace.DirectKeyUsingMap("", map[string]string{"x": "y"})
	actual := args.Map{"result": result, "emptyMap": emptyMap, "emptyFmt": emptyFmt}
	expected := args.Map{"result": "replaced=val", "emptyMap": "text", "emptyFmt": ""}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingMap", actual)
}

func Test_Cov4_ReplaceTemplate_DirectKeyUsingMapTrim(t *testing.T) {
	result := stringutil.Replace.DirectKeyUsingMapTrim("  KEY  ", map[string]string{"KEY": "val"})
	actual := args.Map{"result": result}
	expected := args.Map{"result": "val"}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingMapTrim", actual)
}

func Test_Cov4_ReplaceTemplate_CurlyKeyUsingMap(t *testing.T) {
	result := stringutil.Replace.CurlyKeyUsingMap("{x} y", map[string]string{"x": "a"})
	emptyMap := stringutil.Replace.CurlyKeyUsingMap("text", map[string]string{})
	actual := args.Map{"result": result, "emptyMap": emptyMap}
	expected := args.Map{"result": "a y", "emptyMap": "text"}
	expected.ShouldBeEqual(t, 0, "CurlyKeyUsingMap", actual)
}

func Test_Cov4_ReplaceTemplate_ReplaceWhiteSpaces(t *testing.T) {
	result := stringutil.Replace.ReplaceWhiteSpaces("  hello  world  ")
	empty := stringutil.Replace.ReplaceWhiteSpaces("   ")
	actual := args.Map{"result": result, "empty": empty}
	expected := args.Map{"result": "helloworld", "empty": ""}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpaces", actual)
}

func Test_Cov4_ReplaceTemplate_ReplaceWhiteSpacesToSingle(t *testing.T) {
	result := stringutil.Replace.ReplaceWhiteSpacesToSingle("  hello   world  ")
	empty := stringutil.Replace.ReplaceWhiteSpacesToSingle("   ")
	actual := args.Map{"result": result, "empty": empty}
	expected := args.Map{"result": "hello world", "empty": ""}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle", actual)
}

func Test_Cov4_ReplaceTemplate_UsingWrappedTemplate(t *testing.T) {
	result := stringutil.Replace.UsingWrappedTemplate("test {wrapped} here", "replaced")
	empty := stringutil.Replace.UsingWrappedTemplate("", "replaced")
	actual := args.Map{"notEmpty": result != "", "empty": empty}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "UsingWrappedTemplate", actual)
}

func Test_Cov4_ReplaceTemplate_UsingBracketsWrappedTemplate(t *testing.T) {
	result := stringutil.Replace.UsingBracketsWrappedTemplate("test {brackets-wrapped} here", "replaced")
	empty := stringutil.Replace.UsingBracketsWrappedTemplate("", "replaced")
	actual := args.Map{"notEmpty": result != "", "empty": empty}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "UsingBracketsWrappedTemplate", actual)
}

func Test_Cov4_ReplaceTemplate_UsingQuotesWrappedTemplate(t *testing.T) {
	result := stringutil.Replace.UsingQuotesWrappedTemplate(`test "{quotes-wrapped}" here`, "replaced")
	empty := stringutil.Replace.UsingQuotesWrappedTemplate("", "replaced")
	actual := args.Map{"notEmpty": result != "", "empty": empty}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "UsingQuotesWrappedTemplate", actual)
}

func Test_Cov4_ReplaceTemplate_UsingValueTemplate(t *testing.T) {
	result := stringutil.Replace.UsingValueTemplate("test {value} here", "replaced")
	empty := stringutil.Replace.UsingValueTemplate("", "replaced")
	actual := args.Map{"notEmpty": result != "", "empty": empty}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "UsingValueTemplate", actual)
}

func Test_Cov4_ReplaceTemplate_UsingValueWithFieldsTemplate(t *testing.T) {
	result := stringutil.Replace.UsingValueWithFieldsTemplate("test {value-fields} here", "replaced")
	empty := stringutil.Replace.UsingValueWithFieldsTemplate("", "replaced")
	actual := args.Map{"notEmpty": result != "", "empty": empty}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "UsingValueWithFieldsTemplate", actual)
}

func Test_Cov4_ReplaceTemplate_DirectKeyUsingKeyVal(t *testing.T) {
	result := stringutil.Replace.DirectKeyUsingKeyVal(
		"A and B",
		stringutil.KeyValReplacer{Key: "A", Value: "1"},
		stringutil.KeyValReplacer{Key: "B", Value: "2"},
	)
	emptyKv := stringutil.Replace.DirectKeyUsingKeyVal("text")
	emptyFmt := stringutil.Replace.DirectKeyUsingKeyVal("")
	actual := args.Map{"notEmpty": result != "", "emptyKv": emptyKv, "emptyFmt": emptyFmt}
	expected := args.Map{"notEmpty": true, "emptyKv": "text", "emptyFmt": ""}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingKeyVal", actual)
}
