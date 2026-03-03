package coredynamictests

import (
	"fmt"
	"reflect"
	"testing"

	"gitlab.com/auk-go/core/errcore"
)

// ==========================================
// Test: IsInvalid
// ==========================================

func Test_CastedResult_IsInvalid_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsInvalidTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.CR.IsInvalid())}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}

// ==========================================
// Test: IsNotNull
// ==========================================

func Test_CastedResult_IsNotNull_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsNotNullTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.CR.IsNotNull())}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}

// ==========================================
// Test: IsNotPointer
// ==========================================

func Test_CastedResult_IsNotPointer_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsNotPointerTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.CR.IsNotPointer())}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}

// ==========================================
// Test: IsNotMatchingAcceptedType
// ==========================================

func Test_CastedResult_IsNotMatchingAcceptedType_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsNotMatchingAcceptedTypeTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.CR.IsNotMatchingAcceptedType())}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}

// ==========================================
// Test: IsSourceKind
// ==========================================

func Test_CastedResult_IsSourceKind_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsSourceKindTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.CR.IsSourceKind(tc.CheckKind))}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}

// ==========================================
// Test: HasError
// ==========================================

func Test_CastedResult_HasError_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultHasErrorTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.CR.HasError())}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}

// ==========================================
// Test: HasAnyIssues
// ==========================================

func Test_CastedResult_HasAnyIssues_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultHasAnyIssuesTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.CR.HasAnyIssues())}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}

// ==========================================
// Test: SourceReflectType
// ==========================================

func Test_CastedResult_SourceReflectType_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultSourceReflectTypeTestCases {
		actLines := []string{
			tc.CR.SourceReflectType.Name(),
			fmt.Sprintf("%v", tc.CR.IsSourceKind(reflect.String)),
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}

// ==========================================
// Test: Casted
// ==========================================

func Test_CastedResult_CastedValue_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultCastedValueTestCases {
		actLines := []string{
			fmt.Sprintf("%v", tc.CR.Casted),
			fmt.Sprintf("%v", tc.CR.HasAnyIssues()),
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}

// ==========================================
// Test: IsSourcePointer
// ==========================================

func Test_CastedResult_IsSourcePointer_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsSourcePointerTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.CR.IsSourcePointer)}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}
