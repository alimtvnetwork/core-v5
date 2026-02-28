package mutexbykeytests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/mutexbykey"
)

func Test_Get_Verification(t *testing.T) {
	for caseIndex, testCase := range getAndDeleteTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		key, _ := input.GetAsString("key")

		// Act
		mutex := mutexbykey.Get(key)
		isNotNil := fmt.Sprintf("%v", mutex != nil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, isNotNil)

		// Cleanup
		mutexbykey.Delete(key)
	}
}

func Test_Delete_Verification(t *testing.T) {
	for caseIndex, testCase := range deleteTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		key, _ := input.GetAsString("key")

		// pre-create only for the "existing key" case
		when := input.When()
		if fmt.Sprintf("%v", when) == "given existing key to delete" {
			mutexbykey.Get(key)
		}

		// Act
		deleted := mutexbykey.Delete(key)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", deleted))
	}
}
