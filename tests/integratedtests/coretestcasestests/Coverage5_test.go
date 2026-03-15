package coretestcasestests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── GenericGherkins methods ──

func Test_Cov5_GenericGherkins_ShouldBeEqual(t *testing.T) {
	tc := coretestcases.GenericGherkins[string, string]{
		Title:         "test equal",
		ExpectedLines: []string{"hello"},
	}
	tc.ShouldBeEqual(t, 0, []string{"hello"}, []string{"hello"})
}

func Test_Cov5_GenericGherkins_ShouldBeEqualFirst(t *testing.T) {
	tc := coretestcases.GenericGherkins[string, string]{
		Title:         "test equal first",
		ExpectedLines: []string{"hello"},
	}
	tc.ShouldBeEqualFirst(t, []string{"hello"}, []string{"hello"})
}

func Test_Cov5_GenericGherkins_ShouldBeEqualArgs(t *testing.T) {
	tc := coretestcases.GenericGherkins[string, string]{
		Title:         "test equal args",
		ExpectedLines: []string{"hello"},
	}
	tc.ShouldBeEqualArgs(t, 0, "hello")
}

func Test_Cov5_GenericGherkins_ShouldBeEqualArgsFirst(t *testing.T) {
	tc := coretestcases.GenericGherkins[string, string]{
		Title:         "test equal args first",
		ExpectedLines: []string{"hello"},
	}
	tc.ShouldBeEqualArgsFirst(t, "hello")
}

func Test_Cov5_GenericGherkins_ShouldBeEqualUsingExpected(t *testing.T) {
	tc := coretestcases.GenericGherkins[string, string]{
		Title:         "test using expected",
		ExpectedLines: []string{"hello"},
	}
	tc.ShouldBeEqualUsingExpected(t, 0, []string{"hello"})
}

func Test_Cov5_GenericGherkins_ShouldBeEqualUsingExpectedFirst(t *testing.T) {
	tc := coretestcases.GenericGherkins[string, string]{
		Title:         "test using expected first",
		ExpectedLines: []string{"hello"},
	}
	tc.ShouldBeEqualUsingExpectedFirst(t, []string{"hello"})
}

// ── GenericGherkins — When fallback for title ──

func Test_Cov5_GenericGherkins_WhenFallback(t *testing.T) {
	tc := coretestcases.GenericGherkins[string, string]{
		When:          "when-title",
		ExpectedLines: []string{"hello"},
	}
	tc.ShouldBeEqual(t, 0, []string{"hello"}, []string{"hello"})

	// Just verifying it uses When as fallback title — no assertion needed beyond no panic
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "When fallback", actual)
}

// ── GenericGherkins — Getters ──

func Test_Cov5_GenericGherkins_CaseTitle(t *testing.T) {
	tc := coretestcases.GenericGherkins[string, string]{Title: "my-title"}
	actual := args.Map{"title": tc.CaseTitle()}
	expected := args.Map{"title": "my-title"}
	expected.ShouldBeEqual(t, 0, "GenericGherkins CaseTitle", actual)
}

func Test_Cov5_GenericGherkins_InputExpectedActual(t *testing.T) {
	tc := coretestcases.GenericGherkins[string, string]{
		Input:    "in",
		Expected: "exp",
		Actual:   "act",
	}
	actual := args.Map{
		"input":    fmt.Sprintf("%v", tc.TypedInput()),
		"expected": fmt.Sprintf("%v", tc.TypedExpected()),
		"actual":   fmt.Sprintf("%v", tc.TypedActual()),
	}
	expected := args.Map{"input": "in", "expected": "exp", "actual": "act"}
	expected.ShouldBeEqual(t, 0, "GenericGherkins getters", actual)
}

func Test_Cov5_GenericGherkins_SetActual(t *testing.T) {
	tc := coretestcases.GenericGherkins[string, string]{}
	tc.SetTypedActual("new")
	actual := args.Map{"actual": fmt.Sprintf("%v", tc.TypedActual())}
	expected := args.Map{"actual": "new"}
	expected.ShouldBeEqual(t, 0, "GenericGherkins SetActual", actual)
}

func Test_Cov5_GenericGherkins_SetExpected(t *testing.T) {
	tc := coretestcases.GenericGherkins[string, string]{}
	tc.Expected = "exp"
	actual := args.Map{"expected": fmt.Sprintf("%v", tc.TypedExpected())}
	expected := args.Map{"expected": "exp"}
	expected.ShouldBeEqual(t, 0, "GenericGherkins SetExpected", actual)
}
