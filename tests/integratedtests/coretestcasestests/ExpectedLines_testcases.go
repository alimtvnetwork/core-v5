package coretestcasestests

import (
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// --- ExpectedLines: int ---

var expectedLinesIntTestCase = coretestcases.CaseV1{
	Title:         "ExpectedLines converts int to string",
	ExpectedInput: 42,
}

var expectedLinesIntExpected = []string{"42"}

// --- ExpectedLines: bool true ---

var expectedLinesBoolTrueTestCase = coretestcases.CaseV1{
	Title:         "ExpectedLines converts bool true",
	ExpectedInput: true,
}

var expectedLinesBoolTrueExpected = []string{"true"}

// --- ExpectedLines: bool false ---

var expectedLinesBoolFalseTestCase = coretestcases.CaseV1{
	Title:         "ExpectedLines converts bool false",
	ExpectedInput: false,
}

var expectedLinesBoolFalseExpected = []string{"false"}

// --- ExpectedLines: []int (falls through to PrettyJSON) ---

var expectedLinesIntSliceTestCase = coretestcases.CaseV1{
	Title:         "ExpectedLines converts []int via PrettyJSON",
	ExpectedInput: []int{10, 20, 30},
}

var expectedLinesIntSliceExpected = []string{
	"[",
	"  10,",
	"  20,",
	"  30",
	"]",
}

// --- ExpectedLines: []bool (falls through to PrettyJSON) ---

var expectedLinesBoolSliceTestCase = coretestcases.CaseV1{
	Title:         "ExpectedLines converts []bool via PrettyJSON",
	ExpectedInput: []bool{true, false, true},
}

var expectedLinesBoolSliceExpected = []string{
	"[",
	"  true,",
	"  false,",
	"  true",
	"]",
}

// --- ExpectedLines: string (existing behavior) ---

var expectedLinesStringTestCase = coretestcases.CaseV1{
	Title:         "ExpectedLines wraps string into slice",
	ExpectedInput: "hello",
}

var expectedLinesStringExpected = []string{"hello"}

// --- ExpectedLines: []string (existing behavior) ---

var expectedLinesStringSliceTestCase = coretestcases.CaseV1{
	Title:         "ExpectedLines returns []string as-is",
	ExpectedInput: []string{"a", "b", "c"},
}

var expectedLinesStringSliceExpected = []string{"a", "b", "c"}

// --- ExpectedLines: map[string]int ---

var expectedLinesMapStringIntTestCase = coretestcases.CaseV1{
	Title: "ExpectedLines converts map[string]int sorted",
	ExpectedInput: map[string]int{
		"age":   30,
		"count": 5,
	},
}

var expectedLinesMapStringIntExpected = []string{
	"age : 30",
	"count : 5",
}

// --- Verification wrapper test case for assertion ---

func newExpectedLinesVerificationCase(title string, expected []string) coretests.BaseTestCase {
	return coretests.BaseTestCase{
		Title:         title,
		ExpectedInput: expected,
	}
}
