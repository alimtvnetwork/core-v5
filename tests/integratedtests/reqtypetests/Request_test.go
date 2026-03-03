package reqtypetests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/reqtype"
)

func Test_Request_Identity_Verification(t *testing.T) {
	for caseIndex, testCase := range requestIdentityTestCases {
		// Arrange
		input := testCase.ArrangeInput.(reqtype.Request)

		// Act
		name := input.Name()
		isValid := fmt.Sprintf("%v", input.IsValid())
		isInvalid := fmt.Sprintf("%v", input.IsInvalid())

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			name,
			isValid,
			isInvalid,
		)
	}
}

func Test_Request_LogicalGroups_Verification(t *testing.T) {
	for caseIndex, testCase := range requestLogicalGroupTestCases {
		// Arrange
		input := testCase.ArrangeInput.(reqtype.Request)

		// Act
		isCreate := fmt.Sprintf("%v", input.IsCreateLogically())
		isDrop := fmt.Sprintf("%v", input.IsDropLogically())
		isCrud := fmt.Sprintf("%v", input.IsCrudOnlyLogically())
		isReadOrEdit := fmt.Sprintf("%v", input.IsReadOrEditLogically())

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			isCreate,
			isDrop,
			isCrud,
			isReadOrEdit,
		)
	}
}

func Test_Request_HttpMethods_Verification(t *testing.T) {
	for caseIndex, testCase := range requestHttpTestCases {
		// Arrange
		input := testCase.ArrangeInput.(reqtype.Request)

		// Act
		isGet := fmt.Sprintf("%v", input.IsGetHttp())
		isPost := fmt.Sprintf("%v", input.IsPostHttp())
		isPut := fmt.Sprintf("%v", input.IsPutHttp())
		isDelete := fmt.Sprintf("%v", input.IsDeleteHttp())
		isPatch := fmt.Sprintf("%v", input.IsPatchHttp())

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			isGet,
			isPost,
			isPut,
			isDelete,
			isPatch,
		)
	}
}
