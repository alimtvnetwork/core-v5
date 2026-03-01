package ostypetests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/coretestcases"
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
		groupName := group.Name()
		isUnix := fmt.Sprintf("%v", group.IsUnix())
		isWindows := fmt.Sprintf("%v", group.IsWindows())

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, groupName, isUnix, isWindows)
	}
}

func Test_Variation_Identity_Verification(t *testing.T) {
	for caseIndex, testCase := range variationIdentityTestCases {
		// Arrange
		input := testCase.ArrangeInput.(ostype.Variation)

		// Act
		isWindows := fmt.Sprintf("%v", input.IsWindows())
		isLinux := fmt.Sprintf("%v", input.IsLinux())
		isDarwin := fmt.Sprintf("%v", input.IsDarwinOrMacOs())
		isValid := fmt.Sprintf("%v", input.IsValid())

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, isWindows, isLinux, isDarwin, isValid)
	}
}
