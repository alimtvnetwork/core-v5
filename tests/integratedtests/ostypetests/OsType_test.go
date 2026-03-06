package ostypetests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/ostype"
)

func Test_GetVariant_Verification(t *testing.T) {
	for caseIndex, testCase := range getVariantTestCases {
		// Arrange
		input := testCase.ArrangeInput.(string)

		// Act
		variant := ostype.GetVariant(input)
		name := variant.Name()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, name)
	}
}

func Test_GetGroup_Verification(t *testing.T) {
	for caseIndex, testCase := range getGroupTestCases {
		// Arrange
		input := testCase.ArrangeInput.(string)

		// Act
		group := ostype.GetGroup(input)
		name := group.Name()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, name)
	}
}

func Test_Variation_Group_Verification(t *testing.T) {
	for caseIndex, testCase := range variationGroupTestCases {
		// Arrange
		input := testCase.ArrangeInput.(ostype.Variation)

		// Act
		group := input.Group()
		actual := args.Map{
			"groupName": group.Name(),
			"isUnix":    fmt.Sprintf("%v", group.IsUnix()),
			"isWindows": fmt.Sprintf("%v", group.IsWindows()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variation_Identity_Verification(t *testing.T) {
	for caseIndex, testCase := range variationIdentityTestCases {
		// Arrange
		input := testCase.ArrangeInput.(ostype.Variation)

		// Act
		actual := args.Map{
			"isWindows": fmt.Sprintf("%v", input.IsWindows()),
			"isLinux":   fmt.Sprintf("%v", input.IsLinux()),
			"isDarwin":  fmt.Sprintf("%v", input.IsDarwinOrMacOs()),
			"isValid":   fmt.Sprintf("%v", input.IsValid()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
