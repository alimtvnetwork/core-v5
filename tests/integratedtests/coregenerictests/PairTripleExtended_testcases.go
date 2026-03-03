package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Pair — IsEqual edge cases
// ==========================================================================

var pairIsEqualSameValuesDiffValidityTestCase = coretestcases.CaseV1{
	Title:         "IsEqual: same values, different validity → not equal",
	ExpectedInput: []string{"false"},
}

var pairIsEqualDiffRightTestCase = coretestcases.CaseV1{
	Title:         "IsEqual: different right values → not equal",
	ExpectedInput: []string{"false"},
}

var pairIsEqualBothInvalidZeroTestCase = coretestcases.CaseV1{
	Title:         "IsEqual: both invalid with same zero values → equal",
	ExpectedInput: []string{"true"},
}

var pairIsEqualIntSameTestCase = coretestcases.CaseV1{
	Title:         "IsEqual: Pair[int,int] same values → equal",
	ExpectedInput: []string{"true"},
}

var pairIsEqualIntDiffTestCase = coretestcases.CaseV1{
	Title:         "IsEqual: Pair[int,int] different values → not equal",
	ExpectedInput: []string{"false"},
}

var pairIsEqualMixedTypesTestCase = coretestcases.CaseV1{
	Title:         "IsEqual: Pair[string,int] mixed types → equal",
	ExpectedInput: []string{"true"},
}

// ==========================================================================
// Pair — HasMessage edge cases
// ==========================================================================

var pairHasMessageValidNoMsgTestCase = coretestcases.CaseV1{
	Title:         "HasMessage: valid pair with no message → false",
	ExpectedInput: []string{"false"},
}

var pairHasMessageInvalidWithMsgTestCase = coretestcases.CaseV1{
	Title:         "HasMessage: invalid pair with message → true",
	ExpectedInput: []string{"true"},
}

var pairHasMessageWhitespaceTestCase = coretestcases.CaseV1{
	Title:         "HasMessage: pair with whitespace-only message → true",
	ExpectedInput: []string{"true"},
}

var pairHasMessageNilTestCase = coretestcases.CaseV1{
	Title:         "HasMessage: nil pair → false",
	ExpectedInput: []string{"false"},
}

// ==========================================================================
// Pair — IsInvalid edge cases
// ==========================================================================

var pairIsInvalidValidTestCase = coretestcases.CaseV1{
	Title:         "IsInvalid: valid pair → false",
	ExpectedInput: []string{"false"},
}

var pairIsInvalidInvalidTestCase = coretestcases.CaseV1{
	Title:         "IsInvalid: invalid pair → true",
	ExpectedInput: []string{"true"},
}

var pairIsInvalidNilTestCase = coretestcases.CaseV1{
	Title:         "IsInvalid: nil pair → true",
	ExpectedInput: []string{"true"},
}

// ==========================================================================
// Pair — String output
// ==========================================================================

var pairStringValidTestCase = coretestcases.CaseV1{
	Title:         "String: valid Pair[string,string]",
	ExpectedInput: []string{"{Left: hello, Right: world, IsValid: true}"},
}

var pairStringInvalidZeroTestCase = coretestcases.CaseV1{
	Title:         "String: invalid Pair with zero values",
	ExpectedInput: []string{"{Left: , Right: , IsValid: false}"},
}

var pairStringNilTestCase = coretestcases.CaseV1{
	Title:         "String: nil Pair → empty",
	ExpectedInput: []string{""},
}

var pairStringMixedTypeTestCase = coretestcases.CaseV1{
	Title:         "String: Pair[string,int]",
	ExpectedInput: []string{"{Left: key, Right: 42, IsValid: true}"},
}

// ==========================================================================
// Triple — IsEqual edge cases
// ==========================================================================

var tripleIsEqualSameTestCase = coretestcases.CaseV1{
	Title:         "IsEqual: same values same validity → equal",
	ExpectedInput: []string{"true"},
}

var tripleIsEqualDiffValidityTestCase = coretestcases.CaseV1{
	Title:         "IsEqual: same values different validity → not equal",
	ExpectedInput: []string{"false"},
}

var tripleIsEqualDiffMiddleTestCase = coretestcases.CaseV1{
	Title:         "IsEqual: different middle → not equal",
	ExpectedInput: []string{"false"},
}

var tripleIsEqualBothNilTestCase = coretestcases.CaseV1{
	Title:         "IsEqual: both nil → equal",
	ExpectedInput: []string{"true"},
}

var tripleIsEqualNilVsNonNilTestCase = coretestcases.CaseV1{
	Title:         "IsEqual: nil vs non-nil → not equal",
	ExpectedInput: []string{"false"},
}

// ==========================================================================
// Triple — HasMessage edge cases
// ==========================================================================

var tripleHasMessageValidNoMsgTestCase = coretestcases.CaseV1{
	Title:         "HasMessage: valid triple no message → false",
	ExpectedInput: []string{"false"},
}

var tripleHasMessageInvalidWithMsgTestCase = coretestcases.CaseV1{
	Title:         "HasMessage: invalid triple with message → true",
	ExpectedInput: []string{"true"},
}

var tripleHasMessageNilTestCase = coretestcases.CaseV1{
	Title:         "HasMessage: nil triple → false",
	ExpectedInput: []string{"false"},
}

// ==========================================================================
// Triple — IsInvalid edge cases
// ==========================================================================

var tripleIsInvalidValidTestCase = coretestcases.CaseV1{
	Title:         "IsInvalid: valid triple → false",
	ExpectedInput: []string{"false"},
}

var tripleIsInvalidInvalidTestCase = coretestcases.CaseV1{
	Title:         "IsInvalid: invalid triple → true",
	ExpectedInput: []string{"true"},
}

var tripleIsInvalidNilTestCase = coretestcases.CaseV1{
	Title:         "IsInvalid: nil triple → true",
	ExpectedInput: []string{"true"},
}

// ==========================================================================
// Triple — String output
// ==========================================================================

var tripleStringValidTestCase = coretestcases.CaseV1{
	Title:         "String: valid Triple[string,string,string]",
	ExpectedInput: []string{"{Left: a, Middle: b, Right: c, IsValid: true}"},
}

var tripleStringInvalidZeroTestCase = coretestcases.CaseV1{
	Title:         "String: invalid Triple with zero values",
	ExpectedInput: []string{"{Left: , Middle: , Right: , IsValid: false}"},
}

var tripleStringNilTestCase = coretestcases.CaseV1{
	Title:         "String: nil Triple → empty",
	ExpectedInput: []string{""},
}

// ==========================================================================
// Pair — NewPairWithMessage
// ==========================================================================

var pairWithMessageValidTestCase = coretestcases.CaseV1{
	Title:         "NewPairWithMessage valid with message",
	ExpectedInput: []string{"hello", "world", "true", "ok"},
}

var pairWithMessageInvalidTestCase = coretestcases.CaseV1{
	Title:         "NewPairWithMessage invalid with error message",
	ExpectedInput: []string{"", "", "false", "failed"},
}

// ==========================================================================
// Triple — NewTripleWithMessage
// ==========================================================================

var tripleWithMessageValidTestCase = coretestcases.CaseV1{
	Title:         "NewTripleWithMessage valid with message",
	ExpectedInput: []string{"a", "b", "c", "true", "success"},
}

var tripleWithMessageInvalidTestCase = coretestcases.CaseV1{
	Title:         "NewTripleWithMessage invalid with error",
	ExpectedInput: []string{"", "", "", "false", "error occurred"},
}

// ==========================================================================
// Pair — Dispose
// ==========================================================================

var pairDisposeTestCase = coretestcases.CaseV1{
	Title:         "Dispose resets pair same as Clear",
	ExpectedInput: []string{"", "", "false", ""},
}

// ==========================================================================
// Triple — Dispose
// ==========================================================================

var tripleDisposeTestCase = coretestcases.CaseV1{
	Title:         "Dispose resets triple same as Clear",
	ExpectedInput: []string{"", "", "", "false", ""},
}
