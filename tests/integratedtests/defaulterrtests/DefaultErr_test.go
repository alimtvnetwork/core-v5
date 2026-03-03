package defaulterrtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/defaulterr"
)

func Test_DefaultErr_AllSentinels(t *testing.T) {
	errorMap := map[string]error{
		"Marshalling":                       defaulterr.Marshalling,
		"UnMarshalling":                     defaulterr.UnMarshalling,
		"OutOfRange":                        defaulterr.OutOfRange,
		"CannotProcessNilOrEmpty":           defaulterr.CannotProcessNilOrEmpty,
		"NegativeDataCannotProcess":         defaulterr.NegativeDataCannotProcess,
		"NilResult":                         defaulterr.NilResult,
		"UnexpectedValue":                   defaulterr.UnexpectedValue,
		"CannotRemoveFromEmptyCollection":   defaulterr.CannotRemoveFromEmptyCollection,
		"MarshallingFailedDueToNilOrEmpty":  defaulterr.MarshallingFailedDueToNilOrEmpty,
		"UnmarshallingFailedDueToNilOrEmpty": defaulterr.UnmarshallingFailedDueToNilOrEmpty,
		"KeyNotExistInMap":                  defaulterr.KeyNotExistInMap,
	}

	for caseIndex, tc := range defaultErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		errorName, _ := input.GetAsString("error")
		err := errorMap[errorName]

		// Act
		isNotNil := fmt.Sprintf("%v", err != nil)
		hasMessage := fmt.Sprintf("%v", err != nil && err.Error() != "")

		actLines := []string{isNotNil, hasMessage}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
