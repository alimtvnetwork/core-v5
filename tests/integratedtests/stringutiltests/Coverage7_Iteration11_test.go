package stringutiltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

// ══════════════════════════════════════════════════════════════════════════════
// UsingNamerMapOptions — curly and non-curly with actual values
// ══════════════════════════════════════════════════════════════════════════════

// Note: UsingNamerMapOptions non-nil paths require in-package test (namer is unexported)

// ══════════════════════════════════════════════════════════════════════════════
// UsingBracketsWrappedTemplate, UsingQuotesWrappedTemplate — normal paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_I11_UsingBracketsWrappedTemplate_Normal(t *testing.T) {
	result := stringutil.ReplaceTemplate.UsingBracketsWrappedTemplate("prefix {brackets-wrapped} suffix", "VALUE")
	actual := args.Map{"has": len(result) > 0}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "UsingBracketsWrappedTemplate normal", actual)
}

func Test_I11_UsingQuotesWrappedTemplate_Normal(t *testing.T) {
	result := stringutil.ReplaceTemplate.UsingQuotesWrappedTemplate(`prefix "{quotes-wrapped}" suffix`, "VALUE")
	actual := args.Map{"has": len(result) > 0}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "UsingQuotesWrappedTemplate normal", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReplaceWhiteSpaces — with tabs and newlines
// ══════════════════════════════════════════════════════════════════════════════

func Test_I11_ReplaceWhiteSpaces_WithTabs(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.ReplaceWhiteSpaces("  a\tb\nc  ")}
	expected := args.Map{"v": "abc"}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpaces tabs", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReplaceWhiteSpacesToSingle — with newlines/tabs
// ══════════════════════════════════════════════════════════════════════════════

func Test_I11_ReplaceWhiteSpacesToSingle_WithNewlines(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.ReplaceWhiteSpacesToSingle("a\nb\tc")}
	expected := args.Map{"v": "abc"}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle newlines", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CurlyKeyUsingMap — normal path
// ══════════════════════════════════════════════════════════════════════════════

func Test_I11_CurlyKeyUsingMap_Normal(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyKeyUsingMap("{x}-{y}", map[string]string{"x": "1", "y": "2"})}
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "CurlyKeyUsingMap normal", actual)
}
