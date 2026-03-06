package coredynamictests

import (
	"fmt"
	"reflect"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: IsInvalid
// ==========================================

func Test_CastedResult_IsInvalid_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsInvalidTestCases {
		actual := args.Map{"result": tc.CR.IsInvalid()}

		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: IsNotNull
// ==========================================

func Test_CastedResult_IsNotNull_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsNotNullTestCases {
		actual := args.Map{"result": tc.CR.IsNotNull()}

		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: IsNotPointer
// ==========================================

func Test_CastedResult_IsNotPointer_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsNotPointerTestCases {
		actual := args.Map{"result": tc.CR.IsNotPointer()}

		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: IsNotMatchingAcceptedType
// ==========================================

func Test_CastedResult_IsNotMatchingAcceptedType_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsNotMatchingAcceptedTypeTestCases {
		actual := args.Map{"result": tc.CR.IsNotMatchingAcceptedType()}

		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: IsSourceKind
// ==========================================

func Test_CastedResult_IsSourceKind_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsSourceKindTestCases {
		actual := args.Map{"result": tc.CR.IsSourceKind(tc.CheckKind)}

		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: HasError
// ==========================================

func Test_CastedResult_HasError_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultHasErrorTestCases {
		actual := args.Map{"result": tc.CR.HasError()}

		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: HasAnyIssues
// ==========================================

func Test_CastedResult_HasAnyIssues_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultHasAnyIssuesTestCases {
		actual := args.Map{"result": tc.CR.HasAnyIssues()}

		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: SourceReflectType
// ==========================================

func Test_CastedResult_SourceReflectType_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultSourceReflectTypeTestCases {
		actual := args.Map{
			"typeName":     tc.CR.SourceReflectType.Name(),
			"isStringKind": tc.CR.IsSourceKind(reflect.String),
		}

		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Casted
// ==========================================

func Test_CastedResult_CastedValue_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultCastedValueTestCases {
		actual := args.Map{
			"castedValue":  fmt.Sprintf("%v", tc.CR.Casted),
			"hasAnyIssues": tc.CR.HasAnyIssues(),
		}

		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: IsSourcePointer
// ==========================================

func Test_CastedResult_IsSourcePointer_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsSourcePointerTestCases {
		actual := args.Map{"result": tc.CR.IsSourcePointer}

		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
