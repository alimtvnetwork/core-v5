package stringutiltests

import (
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
