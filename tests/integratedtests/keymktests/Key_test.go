package keymktests

import (
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/keymk"
)

func Test_Key_Compile_Verification(t *testing.T) {
	for caseIndex, testCase := range keyCompileTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		main, _ := input.GetAsString("main")
		chains, _ := input.GetAsStrings("chains")

		// Act
		key := keymk.NewKey.DefaultStrings(main, chains...)
		compiled := key.Compile()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, compiled)
	}
}
