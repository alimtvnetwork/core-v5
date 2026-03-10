package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ==========================================
// SliceValidator.IsValid
// ==========================================

func Test_SliceValidator_IsValid_ExactMatch(t *testing.T) {
	tc := svIsValidExactMatchTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"a", "b", "c"}, ExpectedLines: []string{"a", "b", "c"},
	}

	actual := args.Map{"isValid": v.IsValid(true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_IsValid_Mismatch(t *testing.T) {
	tc := svIsValidMismatchTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"a", "b"}, ExpectedLines: []string{"a", "x"},
	}

	actual := args.Map{"isValid": v.IsValid(true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_IsValid_LengthMismatch(t *testing.T) {
	tc := svIsValidLengthMismatchTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"a"}, ExpectedLines: []string{"a", "b"},
	}

	actual := args.Map{"isValid": v.IsValid(true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_IsValid_BothNil(t *testing.T) {
	tc := svIsValidBothNilTestCase
	v := corevalidator.SliceValidator{
		CompareAs: stringcompareas.Equal, ActualLines: nil, ExpectedLines: nil,
	}

	actual := args.Map{"isValid": v.IsValid(true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_IsValid_OneNil(t *testing.T) {
	tc := svIsValidOneNilTestCase
	v := corevalidator.SliceValidator{
		CompareAs: stringcompareas.Equal, ActualLines: nil, ExpectedLines: []string{"a"},
	}

	actual := args.Map{"isValid": v.IsValid(true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver tests migrated to SliceValidator_NilReceiver_testcases.go)

func Test_SliceValidator_IsValid_BothEmpty(t *testing.T) {
	tc := svIsValidBothEmptyTestCase
	v := corevalidator.SliceValidator{
		CompareAs: stringcompareas.Equal, ActualLines: []string{}, ExpectedLines: []string{},
	}

	actual := args.Map{"isValid": v.IsValid(true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// SliceValidator.IsValid — with Trim
// ==========================================

func Test_SliceValidator_IsValid_TrimMatch(t *testing.T) {
	tc := svIsValidTrimMatchTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultTrimCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"  hello  ", " world "}, ExpectedLines: []string{"hello", "world"},
	}

	actual := args.Map{"isValid": v.IsValid(true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// SliceValidator.IsValid — Contains
// ==========================================

func Test_SliceValidator_IsValid_Contains(t *testing.T) {
	tc := svIsValidContainsTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Contains,
		ActualLines: []string{"hello world", "foo bar"}, ExpectedLines: []string{"ello", "bar"},
	}

	actual := args.Map{"isValid": v.IsValid(true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// SliceValidator — helper methods
// ==========================================

func Test_SliceValidator_ActualLinesLength(t *testing.T) {
	tc := svActualLinesLengthTestCase
	v := corevalidator.SliceValidator{ActualLines: []string{"a", "b"}}

	actual := args.Map{"length": v.ActualLinesLength()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver test migrated to SliceValidator_NilReceiver_testcases.go)

func Test_SliceValidator_ExpectingLinesLength(t *testing.T) {
	tc := svExpectingLinesLengthTestCase
	v := corevalidator.SliceValidator{ExpectedLines: []string{"a", "b", "c"}}

	actual := args.Map{"length": v.ExpectingLinesLength()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_IsUsedAlready_False(t *testing.T) {
	tc := svIsUsedAlreadyFalseTestCase
	v := corevalidator.SliceValidator{ExpectedLines: []string{"a"}}

	actual := args.Map{"isUsed": v.IsUsedAlready()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_IsUsedAlready_TrueAfterComparing(t *testing.T) {
	tc := svIsUsedAlreadyTrueTestCase
	v := corevalidator.SliceValidator{
		CompareAs: stringcompareas.Equal, ExpectedLines: []string{"a"}, ActualLines: []string{"a"},
	}
	_ = v.ComparingValidators()

	actual := args.Map{"isUsed": v.IsUsedAlready()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_MethodName(t *testing.T) {
	tc := svMethodNameTestCase
	v := corevalidator.SliceValidator{CompareAs: stringcompareas.Contains}

	actual := args.Map{"name": v.MethodName()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// SliceValidator.SetActual / SetActualVsExpected
// ==========================================

func Test_SliceValidator_SetActual(t *testing.T) {
	tc := svSetActualTestCase
	v := corevalidator.SliceValidator{
		CompareAs: stringcompareas.Equal, ExpectedLines: []string{"a"},
	}
	v.SetActual([]string{"a"})

	actual := args.Map{"length": v.ActualLinesLength()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_SetActualVsExpected(t *testing.T) {
	tc := svSetActualVsExpectedTestCase
	v := corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}
	v.SetActualVsExpected([]string{"a"}, []string{"b"})

	actual := args.Map{
		"actualLen":   v.ActualLinesLength(),
		"expectedLen": v.ExpectingLinesLength(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// SliceValidator.IsValidOtherLines
// ==========================================

func Test_SliceValidator_IsValidOtherLines_Match(t *testing.T) {
	tc := svIsValidOtherLinesMatchTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ExpectedLines: []string{"a", "b"},
	}

	actual := args.Map{"isValid": v.IsValidOtherLines(true, []string{"a", "b"})}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_IsValidOtherLines_Mismatch(t *testing.T) {
	tc := svIsValidOtherLinesMismatchTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ExpectedLines: []string{"a", "b"},
	}

	actual := args.Map{"isValid": v.IsValidOtherLines(true, []string{"a", "x"})}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// SliceValidator.AllVerifyError
// ==========================================

func Test_SliceValidator_AllVerifyError_Pass(t *testing.T) {
	tc := svAllVerifyErrorPassTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"a", "b"}, ExpectedLines: []string{"a", "b"},
	}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test", IsCaseSensitive: true}

	actual := args.Map{"hasError": v.AllVerifyError(params) != nil}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_AllVerifyError_Fail(t *testing.T) {
	tc := svAllVerifyErrorFailTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"a", "x"}, ExpectedLines: []string{"a", "b"},
	}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test", IsCaseSensitive: true}

	actual := args.Map{"hasError": v.AllVerifyError(params) != nil}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver test migrated to SliceValidator_NilReceiver_testcases.go)

func Test_SliceValidator_AllVerifyError_SkipEmpty(t *testing.T) {
	tc := svAllVerifyErrorSkipEmptyTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{}, ExpectedLines: []string{"a"},
	}
	params := &corevalidator.Parameter{CaseIndex: 0, IsSkipCompareOnActualEmpty: true}

	actual := args.Map{"hasError": v.AllVerifyError(params) != nil}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// SliceValidator.VerifyFirstError
// ==========================================

func Test_SliceValidator_VerifyFirstError_Pass(t *testing.T) {
	tc := svVerifyFirstErrorPassTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"a", "b"}, ExpectedLines: []string{"a", "b"},
	}
	params := &corevalidator.Parameter{CaseIndex: 0, IsCaseSensitive: true}

	actual := args.Map{"hasError": v.VerifyFirstError(params) != nil}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver test migrated to SliceValidator_NilReceiver_testcases.go)

// ==========================================
// SliceValidator.Dispose
// ==========================================

func Test_SliceValidator_Dispose(t *testing.T) {
	tc := svDisposeTestCase
	v := corevalidator.SliceValidator{
		CompareAs: stringcompareas.Equal, ActualLines: []string{"a"}, ExpectedLines: []string{"a"},
	}
	v.Dispose()

	actual := args.Map{
		"actualNil":   v.ActualLines == nil,
		"expectedNil": v.ExpectedLines == nil,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver test migrated to SliceValidator_NilReceiver_testcases.go)

// ==========================================
// SliceValidator — case insensitive
// ==========================================

func Test_SliceValidator_IsValid_CaseInsensitive(t *testing.T) {
	tc := svCaseInsensitiveTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"Hello", "WORLD"}, ExpectedLines: []string{"hello", "world"},
	}

	actual := args.Map{"isValid": v.IsValid(false)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_IsValid_CaseSensitiveFail(t *testing.T) {
	tc := svCaseSensitiveFailTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"Hello"}, ExpectedLines: []string{"hello"},
	}

	actual := args.Map{"isValid": v.IsValid(true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// NewSliceValidatorUsingErr
// ==========================================

func Test_NewSliceValidatorUsingErr_NilError(t *testing.T) {
	tc := svNewUsingErrNilTestCase
	v := corevalidator.NewSliceValidatorUsingErr(
		nil, "expected\nlines",
		true, false, false,
		stringcompareas.Equal,
	)

	actual := args.Map{
		"isNotNil":  v != nil,
		"actualLen": v.ActualLinesLength(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}
