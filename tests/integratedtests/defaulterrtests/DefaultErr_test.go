package defaulterrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/defaulterr"
)

func Test_DefaultErr_AllSentinels(t *testing.T) {
	errorMap := map[string]error{
		"Marshalling":                        defaulterr.Marshalling,
		"UnMarshalling":                      defaulterr.UnMarshalling,
		"OutOfRange":                         defaulterr.OutOfRange,
		"CannotProcessNilOrEmpty":            defaulterr.CannotProcessNilOrEmpty,
		"NegativeDataCannotProcess":          defaulterr.NegativeDataCannotProcess,
		"NilResult":                          defaulterr.NilResult,
		"UnexpectedValue":                    defaulterr.UnexpectedValue,
		"CannotRemoveFromEmptyCollection":    defaulterr.CannotRemoveFromEmptyCollection,
		"MarshallingFailedDueToNilOrEmpty":   defaulterr.MarshallingFailedDueToNilOrEmpty,
		"UnmarshallingFailedDueToNilOrEmpty": defaulterr.UnmarshallingFailedDueToNilOrEmpty,
		"KeyNotExistInMap":                   defaulterr.KeyNotExistInMap,
	}

	for caseIndex, tc := range defaultErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		errorName, _ := input.GetAsString("error")
		err := errorMap[errorName]

		// Act
		actual := args.Map{
			"isNotNil":   err != nil,
			"hasMessage": err != nil && err.Error() != "",
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
