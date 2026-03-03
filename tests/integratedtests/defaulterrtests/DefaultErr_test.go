package defaulterrtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/defaulterr"
	"gitlab.com/auk-go/core/errcore"
)

func verifyDefaultErr(t *testing.T, caseIndex int, title string, err error) {
	isNotNil := fmt.Sprintf("%v", err != nil)
	hasMessage := fmt.Sprintf("%v", err != nil && err.Error() != "")

	actLines := []string{isNotNil, hasMessage}
	expectedLines := []string{"true", "true"}

	errcore.AssertDiffOnMismatch(t, caseIndex, title, actLines, expectedLines)
}

func Test_DefaultErr_Marshalling(t *testing.T) {
	verifyDefaultErr(t, 0, "Marshalling error is not nil", defaulterr.Marshalling)
}

func Test_DefaultErr_UnMarshalling(t *testing.T) {
	verifyDefaultErr(t, 0, "UnMarshalling error is not nil", defaulterr.UnMarshalling)
}

func Test_DefaultErr_OutOfRange(t *testing.T) {
	verifyDefaultErr(t, 0, "OutOfRange error is not nil", defaulterr.OutOfRange)
}

func Test_DefaultErr_CannotProcessNilOrEmpty(t *testing.T) {
	verifyDefaultErr(t, 0, "CannotProcessNilOrEmpty error is not nil", defaulterr.CannotProcessNilOrEmpty)
}

func Test_DefaultErr_NegativeDataCannotProcess(t *testing.T) {
	verifyDefaultErr(t, 0, "NegativeDataCannotProcess error is not nil", defaulterr.NegativeDataCannotProcess)
}

func Test_DefaultErr_NilResult(t *testing.T) {
	verifyDefaultErr(t, 0, "NilResult error is not nil", defaulterr.NilResult)
}

func Test_DefaultErr_UnexpectedValue(t *testing.T) {
	verifyDefaultErr(t, 0, "UnexpectedValue error is not nil", defaulterr.UnexpectedValue)
}

func Test_DefaultErr_CannotRemoveFromEmptyCollection(t *testing.T) {
	verifyDefaultErr(t, 0, "CannotRemoveFromEmptyCollection error is not nil", defaulterr.CannotRemoveFromEmptyCollection)
}

func Test_DefaultErr_MarshallingFailedDueToNilOrEmpty(t *testing.T) {
	verifyDefaultErr(t, 0, "MarshallingFailedDueToNilOrEmpty error is not nil", defaulterr.MarshallingFailedDueToNilOrEmpty)
}

func Test_DefaultErr_UnmarshallingFailedDueToNilOrEmpty(t *testing.T) {
	verifyDefaultErr(t, 0, "UnmarshallingFailedDueToNilOrEmpty error is not nil", defaulterr.UnmarshallingFailedDueToNilOrEmpty)
}

func Test_DefaultErr_KeyNotExistInMap(t *testing.T) {
	verifyDefaultErr(t, 0, "KeyNotExistInMap error is not nil", defaulterr.KeyNotExistInMap)
}
