package chmodhelpertests

import (
	"sort"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coretests/coretestcases"
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

	expectedLines := testCase.ExpectedInput.([]string)
	expectedStr := strings.Join(expectedLines, "\n")

	actNorm := nonWhiteSortedLines(actStr)
	expNorm := nonWhiteSortedLines(expectedStr)

	normalizedCase := testCase
	normalizedCase.ExpectedInput = expNorm

	normalizedCase.ShouldBeEqual(t, caseIndex, actNorm...)
}
