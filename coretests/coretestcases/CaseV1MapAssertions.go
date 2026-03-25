package coretestcases

import (
	"errors"
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
// On mismatch, a copy-pasteable Go literal block is printed showing
// each entry on its own indexed line in Go literal format:
//
//	Actual Received (2 entries):
//	  0: "containsName": false,
//	  1: "hasError":      false,
//
//	Expected Input (1 entries):
//	  0: "hasError": false,
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
		// Print line-by-line diff for detailed comparison
		errcore.PrintDiffOnMismatch(caseIndex, it.Title, actualLines, expectedLines)

		// Build map-specific diagnostic with Go literal format (copy-pasteable)
		mapErrMsg := errcore.MapMismatchError(
			t.Name(),
			caseIndex,
			it.Title,
			actual.GoLiteralLines(),
			expectedMap.GoLiteralLines(),
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
