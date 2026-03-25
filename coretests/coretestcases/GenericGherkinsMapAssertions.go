package coretestcases

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

// ShouldBeEqualMap compares actual args.Map against Expected args.Map
// on a MapGherkins (GenericGherkins[args.Map, args.Map]).
//
// Both maps are compiled to sorted "key : value" string lines using
// CompileToStrings(), then compared line-by-line.
//
// On mismatch, a copy-pasteable Go literal block is printed showing
// each entry on its own line in Go literal format.
//
// Usage:
//
//	tc.ShouldBeEqualMap(t, caseIndex, actual)
func (it *GenericGherkins[TInput, TExpect]) ShouldBeEqualMap(
	t *testing.T,
	caseIndex int,
	actual args.Map,
) {
	t.Helper()

	expectedMap, ok := any(it.Expected).(args.Map)
	if !ok {
		t.Fatalf("Expected is not args.Map in test case: %s", it.Title)
		return
	}

	title := it.Title
	if title == "" {
		title = it.When
	}

	actualLines := actual.CompileToStrings()
	expectedLines := expectedMap.CompileToStrings()

	hasMismatch := errcore.HasAnyMismatchOnLines(actualLines, expectedLines)

	var validationErr error

	if hasMismatch {
		// Print line-by-line diff for detailed comparison
		errcore.PrintDiffOnMismatch(caseIndex, title, actualLines, expectedLines)

		// Build map-specific diagnostic with Go literal format (copy-pasteable)
		mapErrMsg := errcore.MapMismatchError(
			t.Name(),
			caseIndex,
			title,
			actual.GoLiteralLines(),
			expectedMap.GoLiteralLines(),
		)

		validationErr = errors.New(mapErrMsg)
	}

	convey.Convey(
		title, t, func() {
			convey.So(
				validationErr,
				should.BeNil,
			)
		},
	)
}

// ShouldBeEqualMapFirst asserts using ShouldBeEqualMap with caseIndex=0.
// Use for named single test cases (non-loop).
func (it *GenericGherkins[TInput, TExpect]) ShouldBeEqualMapFirst(
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
