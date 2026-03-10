package coredynamictestwrappers

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/errcore"
)

type FromToTestWrapper struct {
	Header                          string
	From, To, ExpectedValue, actual any
	IsUsePointerInFrom              bool
	IsExpectingError                bool
	HasPanic                        bool
	Validator                       corevalidator.TextValidator
}

func (it FromToTestWrapper) CaseTitle() string {
	return it.Header
}

func (it FromToTestWrapper) Input() any {
	return it.From
}

func (it FromToTestWrapper) Expected() any {
	return it.ExpectedValue
}

func (it FromToTestWrapper) ToFieldToDraftType() *coretests.DraftType {
	return coretests.AnyToDraftType(it.To)
}

func (it FromToTestWrapper) ToFieldToBytes() []byte {
	return coretests.AnyToBytes(it.To)
}

func (it FromToTestWrapper) ExpectedFieldToDraftType() *coretests.DraftType {
	return coretests.AnyToDraftType(it.ExpectedValue)
}

func (it FromToTestWrapper) ExpectedFieldToBytes() []byte {
	return coretests.AnyToBytes(it.ExpectedValue)
}

func (it FromToTestWrapper) SetActual(actual any) {
	it.actual = actual
}

func (it FromToTestWrapper) Actual() any {
	return it.actual
}

// ShouldBeEqual asserts actLines match expectedLines using
// the wrapper's Header as the test title.
func (it FromToTestWrapper) ShouldBeEqual(
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

func (it FromToTestWrapper) AsSimpleTestCaseWrapper() coretests.SimpleTestCaseWrapper {
	return &it
}

func (it *FromToTestWrapper) AsSimpleTestCaseWrapperContractsBinder() coretests.SimpleTestCaseWrapperContractsBinder {
	return it
}
