package chmodhelpertests

import (
	"sort"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

func nonWhiteSortedLines(s string) []string {
	if s == "" {
		return []string{""}
	}

	lines := strings.Split(strings.TrimSpace(s), "\n")
	result := make([]string, len(lines))

	for i, line := range lines {
		tokens := strings.Fields(line)
		sort.Strings(tokens)
		result[i] = strings.Join(tokens, " ")
	}

	sort.Strings(result)

	return result
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
