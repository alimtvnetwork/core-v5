package anycmptests

import (
	"testing"

	"github.com/alimtvnetwork/core/anycmp"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Cmp_SamePointer_Verification(t *testing.T) {
	for caseIndex, testCase := range cmpSamePointerTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pairRaw, _ := input.Get("pair")
		pair := pairRaw.(args.TwoAny)

		// Act
		result := anycmp.Cmp(pair.First, pair.Second)

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Cmp_BothNil_Verification(t *testing.T) {
	for caseIndex, testCase := range cmpBothNilTestCases {
		// Arrange & Act
		result := anycmp.Cmp(nil, nil)

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Cmp_OneNil_Verification(t *testing.T) {
	for caseIndex, testCase := range cmpOneNilTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pairRaw, _ := input.Get("pair")
		pair := pairRaw.(args.TwoAny)

		// Act
		result := anycmp.Cmp(pair.First, pair.Second)

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Cmp_TypedNilBothNull_Verification(t *testing.T) {
	for caseIndex, testCase := range cmpTypedNilBothNullTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pairRaw, _ := input.Get("pair")
		pair := pairRaw.(args.TwoAny)

		// Act
		result := anycmp.Cmp(pair.First, pair.Second)

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Cmp_TypedNilOneSide_Verification(t *testing.T) {
	for caseIndex, testCase := range cmpTypedNilOneSideTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pairRaw, _ := input.Get("pair")
		pair := pairRaw.(args.TwoAny)

		// Act
		result := anycmp.Cmp(pair.First, pair.Second)

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Cmp_BothNonNil_Verification(t *testing.T) {
	for caseIndex, testCase := range cmpBothNonNilTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pairRaw, _ := input.Get("pair")
		pair := pairRaw.(args.TwoAny)

		// Act
		result := anycmp.Cmp(pair.First, pair.Second)

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
