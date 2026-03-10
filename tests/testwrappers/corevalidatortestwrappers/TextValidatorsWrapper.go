package corevalidatortestwrappers

import (
	"testing"

	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/errcore"
)

type TextValidatorsWrapper struct {
	Header string
	// ComparingLines is actually the actual data from
	// test but here it is the test cases for the expectation
	ComparingLines        []string
	Validators            corevalidator.TextValidators
	IsSkipOnContentsEmpty bool
	IsCaseSensitive       bool
	ExpectationLines      []string
}

// ShouldBeEqual asserts actLines match expectedLines using
// the wrapper's Header as the test title.
func (it TextValidatorsWrapper) ShouldBeEqual(
	t *testing.T,
	caseIndex int,
	actLines []string,
	expectedLines []string,
) {
	t.Helper()

	errcore.AssertDiffOnMismatch(
		t,
		caseIndex,
		it.Header,
		actLines,
		expectedLines,
	)
}
