package args

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/errcore"
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

// ShouldBeEqual compares this Map (expected) against actual Map.
//
// Both maps are compiled to sorted "key : value" string lines using
// CompileToStrings(), then compared line-by-line.
//
// On mismatch, a copy-pasteable Go literal block is printed showing
// each entry on its own indexed line in Go literal format.
//
// Usage:
//
//	expected := args.Map{"result": 42}
//	expected.ShouldBeEqual(t, 0, "TestTitle", actual)
func (it Map) ShouldBeEqual(
	t *testing.T,
	caseIndex int,
	title string,
	actual Map,
) {
	t.Helper()

	actualLines := actual.CompileToStrings()
	expectedLines := it.CompileToStrings()

	hasMismatch := errcore.HasAnyMismatchOnLines(actualLines, expectedLines)

	var validationErr error

	if hasMismatch {
		errcore.PrintDiffOnMismatch(caseIndex, title, actualLines, expectedLines)

		mapErrMsg := errcore.MapMismatchError(
			t.Name(),
			caseIndex,
			title,
			actual.GoLiteralLines(),
			it.GoLiteralLines(),
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
