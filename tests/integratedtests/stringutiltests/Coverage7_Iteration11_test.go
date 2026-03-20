package stringutiltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

// ══════════════════════════════════════════════════════════════════════════════
// UsingNamerMapOptions — curly and non-curly with actual values
// ══════════════════════════════════════════════════════════════════════════════

type testNamer11 struct{ name string }

func (n testNamer11) Name() string { return n.name }

func Test_I11_UsingNamerMapOptions_Curly(t *testing.T) {
	m := map[interface{ Name() string }]string{
		testNamer11{"a"}: "1",
		testNamer11{"b"}: "2",
	}
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingNamerMapOptions(true, "{a}-{b}", m)}
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "UsingNamerMapOptions curly", actual)
}

func Test_I11_UsingNamerMapOptions_NonCurly(t *testing.T) {
	m := map[interface{ Name() string }]string{
		testNamer11{"A"}: "1",
	}
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingNamerMapOptions(false, "A-B", m)}
	expected := args.Map{"v": "1-B"}
	expected.ShouldBeEqual(t, 0, "UsingNamerMapOptions non-curly", actual)
}

func Test_I11_UsingNamerMapOptions_EmptyMap(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingNamerMapOptions(true, "hello", nil)}
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "UsingNamerMapOptions empty map", actual)
}

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
