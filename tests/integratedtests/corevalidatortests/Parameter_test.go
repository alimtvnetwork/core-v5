package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
)

func Test_Parameter_IsIgnoreCase(t *testing.T) {
	for caseIndex, tc := range parameterIsIgnoreCaseTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isCaseSensitive, _ := input.GetAsBool("isCaseSensitive")
		p := corevalidator.Parameter{IsCaseSensitive: isCaseSensitive}

		// Act
		actual := args.Map{
			"isIgnoreCase": p.IsIgnoreCase(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Parameter_DefaultValues(t *testing.T) {
	tc := parameterDefaultValuesTestCase

	// Arrange
	p := corevalidator.Parameter{}

	// Act
	actual := args.Map{
		"caseIndex":                  p.CaseIndex,
		"header":                     p.Header,
		"isSkipCompareOnActualEmpty": p.IsSkipCompareOnActualEmpty,
		"isAttachUserInputs":         p.IsAttachUserInputs,
		"isCaseSensitive":            p.IsCaseSensitive,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
