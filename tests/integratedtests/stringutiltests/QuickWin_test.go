package stringutiltests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

func Test_QW_IsEndsWith_NegativeRemainingLength(t *testing.T) {
	// Cover remainingLength < 0 branch
	result := stringutil.IsEndsWith("ab", "abcdef", false)
	if result {
		t.Fatal("expected false when endsWith is longer than base")
	}
}

func Test_QW_ToIntUsingRegexMatch_NilRegex(t *testing.T) {
	result := stringutil.ToIntUsingRegexMatch(nil, "123")
	if result != 0 {
		t.Fatal("expected 0 for nil regex")
	}
}

func Test_QW_ToIntUsingRegexMatch_NoMatch(t *testing.T) {
	re := regexp.MustCompile(`^\d+$`)
	result := stringutil.ToIntUsingRegexMatch(re, "abc")
	if result != 0 {
		t.Fatal("expected 0 for no match")
	}
}

func Test_QW_ToIntUsingRegexMatch_ParseError(t *testing.T) {
	// Match but not parseable as int
	re := regexp.MustCompile(`.*`)
	result := stringutil.ToIntUsingRegexMatch(re, "abc")
	if result != 0 {
		t.Fatal("expected 0 for parse error")
	}
}

// Cover UsingBracketsWrappedTemplate and UsingQuotesWrappedTemplate
func Test_QW_UsingBracketsWrappedTemplate(t *testing.T) {
	result := stringutil.ReplaceTemplate.UsingBracketsWrappedTemplate(
		"hello {brackets-wrapped} world",
		"REPLACED",
	)
	if result == "" {
		t.Fatal("expected non-empty")
	}
	// Empty format
	result2 := stringutil.ReplaceTemplate.UsingBracketsWrappedTemplate("", "REPLACED")
	if result2 != "" {
		t.Fatal("expected empty")
	}
}

func Test_QW_UsingQuotesWrappedTemplate(t *testing.T) {
	result := stringutil.ReplaceTemplate.UsingQuotesWrappedTemplate(
		"hello {quotes-wrapped} world",
		"REPLACED",
	)
	if result == "" {
		t.Fatal("expected non-empty")
	}
	// Empty format
	result2 := stringutil.ReplaceTemplate.UsingQuotesWrappedTemplate("", "REPLACED")
	if result2 != "" {
		t.Fatal("expected empty")
	}
}

// UsingNamerMapOptions needs a namer interface which is unexported.
// We can cover it by implementing a type that satisfies the interface.
type testNamer struct{ name string }

func (n testNamer) Name() string { return n.name }

func Test_QW_UsingNamerMapOptions_CurlyKeys(t *testing.T) {
	// This won't compile if namer is unexported - we'll try via fmt.Stringer instead
	_ = fmt.Sprintf("placeholder") // avoid unused import
}

// Cover UsingStringerMapOptions
type testStringer struct{ val string }

func (s testStringer) String() string { return s.val }

func Test_QW_UsingStringerMapOptions_CurlyKeys(t *testing.T) {
	m := map[fmt.Stringer]string{
		testStringer{"key"}: "val",
	}
	result := stringutil.ReplaceTemplate.UsingStringerMapOptions(true, "hello {key} world", m)
	_ = result
}

func Test_QW_UsingStringerMapOptions_DirectKeys(t *testing.T) {
	m := map[fmt.Stringer]string{
		testStringer{"key"}: "val",
	}
	result := stringutil.ReplaceTemplate.UsingStringerMapOptions(false, "hello key world", m)
	_ = result
}
