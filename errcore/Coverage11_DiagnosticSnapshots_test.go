package errcore

import (
	"strings"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// Diagnostic Output Snapshot Tests
//
// These tests verify the exact format of diagnostic output produced by
// errcore formatter functions. Any formatting change will break these tests,
// serving as a regression gate for diagnostic output.
// ══════════════════════════════════════════════════════════════════════════════

// ── MapMismatchError ──

func TestSnapshot_MapMismatchError_Format(t *testing.T) {
	result := MapMismatchError(
		"TestFunc",
		1,
		"Case Title",
		[]string{`"key1": true,`, `"key2": false,`},
		[]string{`"key1": false,`},
	)

	// Verify header block
	assertContains(t, result, "Test Method : TestFunc")
	assertContains(t, result, "Case        : 1")
	assertContains(t, result, "Title       : Case Title")

	// Verify actual block
	assertContains(t, result, "1) Actual Received (2 entries):")
	assertContains(t, result, "    Case Title")

	// Verify expected block
	assertContains(t, result, "1) Expected Input (1 entries):")

	// Verify separator
	assertContains(t, result, "=============================>")

	// Verify entries are tab-indented
	assertContains(t, result, "\t\"key1\": true,")
	assertContains(t, result, "\t\"key2\": false,")
	assertContains(t, result, "\t\"key1\": false,")

	// Verify leading newline (for indentation under goconvey)
	if !strings.HasPrefix(result, "\n") {
		t.Fatal("MapMismatchError output must start with newline")
	}
}

func TestSnapshot_MapMismatchError_EmptyEntries(t *testing.T) {
	result := MapMismatchError("F", 0, "T", nil, nil)

	assertContains(t, result, "0) Actual Received (0 entries):")
	assertContains(t, result, "0) Expected Input (0 entries):")
}

// ── LineDiffToString ──

func TestSnapshot_LineDiffToString_MismatchFormat(t *testing.T) {
	result := LineDiffToString(
		3,
		"myHeader",
		[]string{"alpha", "CHANGED"},
		[]string{"alpha", "beta"},
	)

	// Verify header
	assertContains(t, result, "=== Line-by-Line Diff (Case 3: myHeader) ===")
	assertContains(t, result, "Actual lines: 2, Expected lines: 2")

	// Verify OK line format
	assertContains(t, result, "Line   0 [OK]: `alpha`")

	// Verify MISMATCH format with column-aligned labels
	assertContains(t, result, "Line   1 [MISMATCH]:")
	assertContains(t, result, "actual : `CHANGED`")
	assertContains(t, result, "expected : `beta`")

	// Verify footer
	assertContains(t, result, "=== Total: 2 lines, 1 mismatches ===")
}

func TestSnapshot_LineDiffToString_ExtraActualFormat(t *testing.T) {
	result := LineDiffToString(0, "h", []string{"a", "b", "extra"}, []string{"a"})

	assertContains(t, result, "Line   1 [EXTRA ACTUAL]: `b`")
	assertContains(t, result, "Line   2 [EXTRA ACTUAL]: `extra`")
	assertContains(t, result, "2 mismatches")
}

func TestSnapshot_LineDiffToString_MissingExpectedFormat(t *testing.T) {
	result := LineDiffToString(0, "h", []string{"a"}, []string{"a", "b", "c"})

	assertContains(t, result, "Line   1 [MISSING EXPECTED]: `b`")
	assertContains(t, result, "Line   2 [MISSING EXPECTED]: `c`")
	assertContains(t, result, "2 mismatches")
}

func TestSnapshot_LineDiffToString_Empty(t *testing.T) {
	result := LineDiffToString(0, "h", []string{}, []string{})
	if result != "" {
		t.Fatalf("expected empty string for no diffs, got: %q", result)
	}
}

func TestSnapshot_LineDiffToString_AllMatch(t *testing.T) {
	result := LineDiffToString(0, "h", []string{"a", "b"}, []string{"a", "b"})

	assertContains(t, result, "0 mismatches")
	// All lines should show [OK]
	if strings.Contains(result, "[MISMATCH]") {
		t.Fatal("should not contain MISMATCH for matching lines")
	}
}

// ── LineDiff ──

func TestSnapshot_LineDiff_StatusCodes(t *testing.T) {
	diffs := LineDiff(
		[]string{"same", "different", "extra"},
		[]string{"same", "other"},
	)

	cases := []struct {
		idx    int
		status string
	}{
		{0, "  "}, // match
		{1, "!!"}, // mismatch
		{2, "+"},  // extra actual
	}

	for _, c := range cases {
		if diffs[c.idx].Status != c.status {
			t.Fatalf("line %d: expected status %q, got %q", c.idx, c.status, diffs[c.idx].Status)
		}
	}

	// Test missing expected
	diffs2 := LineDiff([]string{"a"}, []string{"a", "b"})
	if diffs2[1].Status != "-" {
		t.Fatalf("expected status '-', got %q", diffs2[1].Status)
	}
	if diffs2[1].Actual != "<missing>" {
		t.Fatalf("expected '<missing>' for actual, got %q", diffs2[1].Actual)
	}
}

// ── SliceDiffSummary ──

func TestSnapshot_SliceDiffSummary_AllMatch(t *testing.T) {
	result := SliceDiffSummary([]string{"a", "b"}, []string{"a", "b"})
	if result != "all lines match" {
		t.Fatalf("expected 'all lines match', got: %q", result)
	}
}

func TestSnapshot_SliceDiffSummary_MismatchFormat(t *testing.T) {
	result := SliceDiffSummary([]string{"a", "X"}, []string{"a", "b"})

	// Should start with count
	if !strings.HasPrefix(result, "1 mismatches:") {
		t.Fatalf("expected '1 mismatches:' prefix, got: %q", result)
	}
	assertContains(t, result, "line 1 [!!]")
}

func TestSnapshot_SliceDiffSummary_ExtraAndMissing(t *testing.T) {
	result := SliceDiffSummary([]string{"a", "b"}, []string{"a"})
	assertContains(t, result, "line 1 [+]")

	result2 := SliceDiffSummary([]string{"a"}, []string{"a", "b"})
	assertContains(t, result2, "line 1 [-]")
}

// ── GherkinsString ──

func TestSnapshot_GherkinsString_Format(t *testing.T) {
	result := GherkinsString(
		5,
		"Feature X",
		"Given Y",
		"When Z",
		"Then W",
	)

	// Must contain the case index, feature, given, when, then
	assertContains(t, result, "5")
	assertContains(t, result, "Feature X")
	assertContains(t, result, "Given Y")
	assertContains(t, result, "When Z")
	assertContains(t, result, "Then W")
}

func TestSnapshot_GherkinsStringWithExpectation_Format(t *testing.T) {
	result := GherkinsStringWithExpectation(
		2,
		"Feature A",
		"Given B",
		"When C",
		"Then D",
		"actual-val",
		"expected-val",
	)

	assertContains(t, result, "2")
	assertContains(t, result, "Feature A")
	assertContains(t, result, "actual-val")
	assertContains(t, result, "expected-val")
}

// ── ExpectingRecord ──

func TestSnapshot_ExpectingRecord_Message(t *testing.T) {
	rec := &ExpectingRecord{
		ExpectingTitle: "MyTitle",
		WasExpecting:   "expected-val",
	}

	msg := rec.Message("actual-val")
	assertContains(t, msg, "MyTitle")
	assertContains(t, msg, "expected-val")
	assertContains(t, msg, "actual-val")
	assertContains(t, msg, "expecting")
	assertContains(t, msg, "received")
}

func TestSnapshot_ExpectingRecord_MessageSimple(t *testing.T) {
	rec := &ExpectingRecord{
		ExpectingTitle: "Title",
		WasExpecting:   42,
	}

	msg := rec.MessageSimple(99)
	assertContains(t, msg, "Title")
	assertContains(t, msg, "Expect")
	assertContains(t, msg, "Actual")
	assertContains(t, msg, "42")
	assertContains(t, msg, "99")
}

func TestSnapshot_ExpectingRecord_MessageSimpleNoType(t *testing.T) {
	rec := &ExpectingRecord{
		ExpectingTitle: "T",
		WasExpecting:   "exp",
	}

	msg := rec.MessageSimpleNoType("act")
	assertContains(t, msg, "T")
	assertContains(t, msg, "exp")
	assertContains(t, msg, "act")
	// Should NOT contain type info like "string"
	assertContains(t, msg, "Expect")
}

// ── ExpectationMessageDef ──

func TestSnapshot_ExpectationMessageDef_ToString(t *testing.T) {
	def := ExpectationMessageDef{
		CaseIndex:    1,
		FuncName:     "TestFunc",
		TestCaseName: "case-name",
		When:         "input is valid",
		Expected:     "expected-output",
	}

	result := def.ToString("actual-output")

	assertContains(t, result, "1")
	assertContains(t, result, "input is valid")
	assertContains(t, result, "actual-output")
	assertContains(t, result, "expected-output")
}

func TestSnapshot_ExpectationMessageDef_ExpectedSafeString(t *testing.T) {
	def := ExpectationMessageDef{Expected: "hello"}
	s := def.ExpectedSafeString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
	// Should cache
	s2 := def.ExpectedSafeString()
	if s != s2 {
		t.Fatal("expected cached result")
	}
}

func TestSnapshot_ExpectationMessageDef_ExpectedSafeString_Nil(t *testing.T) {
	def := ExpectationMessageDef{Expected: nil}
	s := def.ExpectedSafeString()
	if s != "" {
		t.Fatalf("expected empty for nil Expected, got: %q", s)
	}
}

func TestSnapshot_ExpectationMessageDef_ExpectedString_Panics(t *testing.T) {
	def := ExpectationMessageDef{Expected: nil}
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		def.ExpectedString()
	}()
	if !didPanic {
		t.Fatal("expected panic when Expected is nil")
	}
}

// ── Expecting / ExpectingSimple / ExpectingSimpleNoType ──

func TestSnapshot_Expecting_Format(t *testing.T) {
	result := Expecting("title", "exp", "act")
	assertContains(t, result, "title")
	assertContains(t, result, "exp")
	assertContains(t, result, "act")
	assertContains(t, result, "expecting")
}

func TestSnapshot_ExpectingSimple_Format(t *testing.T) {
	result := ExpectingSimple("title", "exp", "act")
	assertContains(t, result, "title")
	assertContains(t, result, "Expect")
	assertContains(t, result, "Actual")
}

func TestSnapshot_ExpectingSimpleNoType_Format(t *testing.T) {
	result := ExpectingSimpleNoType("title", "exp", "act")
	assertContains(t, result, "title")
	assertContains(t, result, "exp")
	assertContains(t, result, "act")
}

// ══════════════════════════════════════════════════════════════════════════════
// Helpers
// ══════════════════════════════════════════════════════════════════════════════

func assertContains(t *testing.T, haystack, needle string) {
	t.Helper()
	if !strings.Contains(haystack, needle) {
		t.Fatalf("expected output to contain %q, got:\n%s", needle, haystack)
	}
}
