package chmodhelpertests

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// isStackTraceLine returns true if the line is a stack trace artifact
// that should be stripped before comparison.
// See issues/chmodhelpertests-stack-trace-mismatch.md
func isStackTraceLine(line string) bool {
	trimmed := strings.TrimSpace(line)
	if trimmed == "Stack-Trace:" {
		return true
	}
	if strings.HasPrefix(trimmed, "- /") {
		return true
	}
	if strings.HasPrefix(trimmed, "- ErrorRefOnly") ||
		strings.HasPrefix(trimmed, "- getVerifyRwxInternalError") {
		return true
	}
	return false
}

func nonWhiteSortedLines(s string) []string {
	if s == "" {
		return []string{""}
	}

	lines := strings.Split(strings.TrimSpace(s), "\n")
	var filtered []string

	for _, line := range lines {
		if isStackTraceLine(line) {
			continue
		}
		tokens := strings.Fields(line)
		sort.Strings(tokens)
		filtered = append(filtered, strings.Join(tokens, " "))
	}

	if len(filtered) == 0 {
		return []string{""}
	}

	sort.Strings(filtered)

	return filtered
}

func assertNonWhiteSortedEqual(
	t *testing.T,
	testCase coretestcases.CaseV1,
	caseIndex int,
	actualErr error,
) {
	actStr := ""
	if actualErr != nil {
		actStr = actualErr.Error()
	}

	// Fix: handle both string and []string ExpectedInput types
	// See issues/chmodhelpertests-type-assertion-panic.md
	var expectedStr string
	switch v := testCase.ExpectedInput.(type) {
	case []string:
		expectedStr = strings.Join(v, "\n")
	case string:
		expectedStr = v
	default:
		expectedStr = fmt.Sprintf("%v", testCase.ExpectedInput)
	}

	actNorm := nonWhiteSortedLines(actStr)
	expNorm := nonWhiteSortedLines(expectedStr)

	normalizedCase := testCase
	normalizedCase.ExpectedInput = expNorm

	normalizedCase.ShouldBeEqual(t, caseIndex, actNorm...)
}
