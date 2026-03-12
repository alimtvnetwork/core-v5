package coretestcasestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── GenericGherkins Getters ──

func Test_Cov_GenericGherkins_IsFailedToMatch(t *testing.T) {
	g := &coretestcases.StringBoolGherkins{IsMatching: true}
	g2 := &coretestcases.StringBoolGherkins{IsMatching: false}
	actual := args.Map{
		"matchingFails": g.IsFailedToMatch(),
		"notMatchFails": g2.IsFailedToMatch(),
	}
	expected := args.Map{
		"matchingFails": false,
		"notMatchFails": true,
	}
	expected.ShouldBeEqual(t, 0, "IsFailedToMatch", actual)
}

func Test_Cov_GenericGherkins_HasExtraArgs(t *testing.T) {
	g := &coretestcases.StringBoolGherkins{ExtraArgs: args.Map{"k": "v"}}
	g2 := &coretestcases.StringBoolGherkins{}
	var gNil *coretestcases.StringBoolGherkins
	actual := args.Map{
		"hasExtra":    g.HasExtraArgs(),
		"noExtra":     g2.HasExtraArgs(),
		"nilHasExtra": gNil.HasExtraArgs(),
	}
	expected := args.Map{
		"hasExtra":    true,
		"noExtra":     false,
		"nilHasExtra": false,
	}
	expected.ShouldBeEqual(t, 0, "HasExtraArgs", actual)
}

func Test_Cov_GenericGherkins_GetExtra(t *testing.T) {
	g := &coretestcases.StringBoolGherkins{ExtraArgs: args.Map{"k": "v"}}
	var gNil *coretestcases.StringBoolGherkins
	actual := args.Map{
		"found":    g.GetExtra("k"),
		"notFound": g.GetExtra("missing") == nil,
		"nilGet":   gNil.GetExtra("k") == nil,
	}
	expected := args.Map{
		"found":    "v",
		"notFound": true,
		"nilGet":   true,
	}
	expected.ShouldBeEqual(t, 0, "GetExtra", actual)
}

func Test_Cov_GenericGherkins_GetExtraAsString(t *testing.T) {
	g := &coretestcases.StringBoolGherkins{ExtraArgs: args.Map{"k": "v"}}
	var gNil *coretestcases.StringBoolGherkins
	val, ok := g.GetExtraAsString("k")
	nilVal, nilOk := gNil.GetExtraAsString("k")
	actual := args.Map{"val": val, "ok": ok, "nilVal": nilVal, "nilOk": nilOk}
	expected := args.Map{"val": "v", "ok": true, "nilVal": "", "nilOk": false}
	expected.ShouldBeEqual(t, 0, "GetExtraAsString", actual)
}

func Test_Cov_GenericGherkins_GetExtraAsBool(t *testing.T) {
	g := &coretestcases.StringBoolGherkins{ExtraArgs: args.Map{"k": true}}
	var gNil *coretestcases.StringBoolGherkins
	val, ok := g.GetExtraAsBool("k")
	nilVal, nilOk := gNil.GetExtraAsBool("k")
	actual := args.Map{"val": val, "ok": ok, "nilVal": nilVal, "nilOk": nilOk}
	expected := args.Map{"val": true, "ok": true, "nilVal": false, "nilOk": false}
	expected.ShouldBeEqual(t, 0, "GetExtraAsBool", actual)
}

func Test_Cov_GenericGherkins_GetExtraAsBoolDefault(t *testing.T) {
	g := &coretestcases.StringBoolGherkins{ExtraArgs: args.Map{"k": true}}
	var gNil *coretestcases.StringBoolGherkins
	val := g.GetExtraAsBoolDefault("k", false)
	nilVal := gNil.GetExtraAsBoolDefault("k", true)
	actual := args.Map{"val": val, "nilVal": nilVal}
	expected := args.Map{"val": true, "nilVal": true}
	expected.ShouldBeEqual(t, 0, "GetExtraAsBoolDefault", actual)
}

// ── GenericGherkins Formatting ──

func Test_Cov_GenericGherkins_ToString(t *testing.T) {
	g := &coretestcases.StringBoolGherkins{
		Feature: "f", Given: "g", When: "w", Then: "t",
	}
	result := g.ToString(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToString", actual)
}

func Test_Cov_GenericGherkins_String(t *testing.T) {
	g := &coretestcases.StringBoolGherkins{Feature: "f"}
	actual := args.Map{"notEmpty": g.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String", actual)
}

func Test_Cov_GenericGherkins_GetWithExpectation(t *testing.T) {
	g := &coretestcases.StringBoolGherkins{Feature: "f"}
	result := g.GetWithExpectation(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetWithExpectation", actual)
}

func Test_Cov_GenericGherkins_GetMessageConditional(t *testing.T) {
	g := &coretestcases.StringBoolGherkins{Feature: "f"}
	withExp := g.GetMessageConditional(true, 0)
	withoutExp := g.GetMessageConditional(false, 0)
	actual := args.Map{
		"withNotEmpty":    withExp != "",
		"withoutNotEmpty": withoutExp != "",
	}
	expected := args.Map{
		"withNotEmpty":    true,
		"withoutNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "GetMessageConditional", actual)
}

func Test_Cov_GenericGherkins_FullString(t *testing.T) {
	g := &coretestcases.StringBoolGherkins{
		Title: "t", Feature: "f", Given: "g", When: "w", Then: "th",
		ExtraArgs: args.Map{"k": "v"},
	}
	result := g.FullString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FullString", actual)
}

func Test_Cov_GenericGherkins_FullString_Nil(t *testing.T) {
	var g *coretestcases.StringBoolGherkins
	result := g.FullString()
	actual := args.Map{"isNilMsg": result == "<nil GenericGherkins>"}
	expected := args.Map{"isNilMsg": true}
	expected.ShouldBeEqual(t, 0, "FullString_Nil", actual)
}

// ── GenericGherkins CaseTitle ──

func Test_Cov_GenericGherkins_CaseTitle(t *testing.T) {
	gTitle := &coretestcases.StringBoolGherkins{Title: "myTitle", When: "myWhen"}
	gWhen := &coretestcases.StringBoolGherkins{When: "myWhen"}
	actual := args.Map{
		"titleResult": gTitle.CaseTitle(),
		"whenResult":  gWhen.CaseTitle(),
	}
	expected := args.Map{
		"titleResult": "myTitle",
		"whenResult":  "myWhen",
	}
	expected.ShouldBeEqual(t, 0, "CaseTitle", actual)
}
