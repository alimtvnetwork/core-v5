package coreversiontests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coreversion"
)

func Test_Creation_Verification(t *testing.T) {
	for caseIndex, testCase := range versionCreationTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]*coreversion.Version)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, input := range inputs {
			if input.IsInvalid() {
				actualSlice.AppendFmt(
					defaultInvalidCreationFmt,
					i)
			} else {
				actualSlice.AppendFmt(
					defaultCreationFmt,
					i,
					input.String(),
					input.VersionCompact,
					input.VersionDisplay())
			}
		}

		finalActLines := actualSlice.Strings()
		finalCase := testCase.AsCaseV1()

		// Assert
		finalCase.AssertEqual(
			t,
			caseIndex,
			finalActLines...)
	}
}
