package reqtypetests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/reqtype"
)

func Test_Request_Identity_Verification(t *testing.T) {
	for caseIndex, testCase := range requestIdentityTestCases {
		// Arrange
		input := testCase.ArrangeInput.(reqtype.Request)

		// Act
		actual := args.Map{
			"name":      input.Name(),
			"isValid":   fmt.Sprintf("%v", input.IsValid()),
			"isInvalid": fmt.Sprintf("%v", input.IsInvalid()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Request_LogicalGroups_Verification(t *testing.T) {
	for caseIndex, testCase := range requestLogicalGroupTestCases {
		// Arrange
		input := testCase.ArrangeInput.(reqtype.Request)

		// Act
		actual := args.Map{
			"isCreateLogically": fmt.Sprintf("%v", input.IsCreateLogically()),
			"isDropLogically":   fmt.Sprintf("%v", input.IsDropLogically()),
			"isCrudOnly":        fmt.Sprintf("%v", input.IsCrudOnlyLogically()),
			"isReadOrEdit":      fmt.Sprintf("%v", input.IsReadOrEditLogically()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Request_HttpMethods_Verification(t *testing.T) {
	for caseIndex, testCase := range requestHttpTestCases {
		// Arrange
		input := testCase.ArrangeInput.(reqtype.Request)

		// Act
		actual := args.Map{
			"isGet":    fmt.Sprintf("%v", input.IsGetHttp()),
			"isPost":   fmt.Sprintf("%v", input.IsPostHttp()),
			"isPut":    fmt.Sprintf("%v", input.IsPutHttp()),
			"isDelete": fmt.Sprintf("%v", input.IsDeleteHttp()),
			"isPatch":  fmt.Sprintf("%v", input.IsPatchHttp()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
