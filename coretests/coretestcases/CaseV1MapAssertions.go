package coretestcases

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

// ExpectedAsMap retrieves ExpectedInput as args.Map.
//
// Panics if ExpectedInput is not args.Map.
func (it CaseV1) ExpectedAsMap() args.Map {
	m, ok := it.ExpectedInput.(args.Map)

	if !ok {
		panic("ExpectedInput is not args.Map")
	}

	return m
}

// ShouldBeEqualMap compares actual args.Map against ExpectedInput args.Map.
//
// Both maps are compiled to sorted "key : value" string lines using
// CompileToStrings(), then compared line-by-line.
//
// On mismatch, diagnostics are shown in two forms:
//  1. Line-by-line comparison with aligned actual/expected labels
//     and standard header separators (============================>)
//  2. Copy-pasteable Go literal block via slog.Warn
func (it CaseV1) ShouldBeEqualMap(
	t *testing.T,
	caseIndex int,
	actual args.Map,
) {
	t.Helper()

	expectedMap := it.ExpectedAsMap()
	actualLines := actual.CompileToStrings()
	expectedLines := expectedMap.CompileToStrings()

	hasMismatch := errcore.HasAnyMismatchOnLines(actualLines, expectedLines)

	var validationErr error

	if hasMismatch {
		// Print copy-pasteable Go literal via slog
		slog.Warn("copy-pasteable expected (from actual):",
			"caseIndex", caseIndex,
			"title", it.Title,
			"actualGoLiteral", "\n"+actual.GoLiteralString(),
		)

		// Build map-specific diagnostic with header separators
		// and aligned actual/expected labels per line
		mapErrMsg := errcore.MapMismatchError(
			caseIndex,
			it.Title,
			actualLines,
			expectedLines,
		)

		validationErr = errors.New(mapErrMsg)
	}

	convey.Convey(
		it.Title, t, func() {
			convey.So(
				validationErr,
				should.BeNil,
			)
		},
	)
}

// ShouldBeEqualMapFirst asserts using ShouldBeEqualMap with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldBeEqualMapFirst(
	t *testing.T,
	actual args.Map,
) {
	t.Helper()

	it.ShouldBeEqualMap(
		t,
		0,
		actual,
	)
}
