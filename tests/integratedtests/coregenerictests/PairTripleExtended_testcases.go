package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Pair — IsEqual edge cases
// ==========================================================================

var pairIsEqualSameValuesDiffValidityTestCase = coretestcases.CaseV1{
	Name:      "IsEqual: same values, different validity → not equal",
	WantLines: []string{"false"},
}

var pairIsEqualDiffRightTestCase = coretestcases.CaseV1{
	Name:      "IsEqual: different right values → not equal",
	WantLines: []string{"false"},
}

var pairIsEqualBothInvalidZeroTestCase = coretestcases.CaseV1{
	Name:      "IsEqual: both invalid with same zero values → equal",
	WantLines: []string{"true"},
}

var pairIsEqualIntSameTestCase = coretestcases.CaseV1{
	Name:      "IsEqual: Pair[int,int] same values → equal",
	WantLines: []string{"true"},
}

var pairIsEqualIntDiffTestCase = coretestcases.CaseV1{
	Name:      "IsEqual: Pair[int,int] different values → not equal",
	WantLines: []string{"false"},
}

var pairIsEqualMixedTypesTestCase = coretestcases.CaseV1{
	Name:      "IsEqual: Pair[string,int] mixed types → equal",
	WantLines: []string{"true"},
}

// ==========================================================================
// Pair — HasMessage edge cases
// ==========================================================================

var pairHasMessageValidNoMsgTestCase = coretestcases.CaseV1{
	Name:      "HasMessage: valid pair with no message → false",
	WantLines: []string{"false"},
}

var pairHasMessageInvalidWithMsgTestCase = coretestcases.CaseV1{
	Name:      "HasMessage: invalid pair with message → true",
	WantLines: []string{"true"},
}

var pairHasMessageWhitespaceTestCase = coretestcases.CaseV1{
	Name:      "HasMessage: pair with whitespace-only message → true",
	WantLines: []string{"true"},
}

var pairHasMessageNilTestCase = coretestcases.CaseV1{
	Name:      "HasMessage: nil pair → false",
	WantLines: []string{"false"},
}

// ==========================================================================
// Pair — IsInvalid edge cases
// ==========================================================================

var pairIsInvalidValidTestCase = coretestcases.CaseV1{
	Name:      "IsInvalid: valid pair → false",
	WantLines: []string{"false"},
}

var pairIsInvalidInvalidTestCase = coretestcases.CaseV1{
	Name:      "IsInvalid: invalid pair → true",
	WantLines: []string{"true"},
}

var pairIsInvalidNilTestCase = coretestcases.CaseV1{
	Name:      "IsInvalid: nil pair → true",
	WantLines: []string{"true"},
}

// ==========================================================================
// Pair — String output
// ==========================================================================

var pairStringValidTestCase = coretestcases.CaseV1{
	Name:      "String: valid Pair[string,string]",
	WantLines: []string{"{Left: hello, Right: world, IsValid: true}"},
}

var pairStringInvalidZeroTestCase = coretestcases.CaseV1{
	Name:      "String: invalid Pair with zero values",
	WantLines: []string{"{Left: , Right: , IsValid: false}"},
}

var pairStringNilTestCase = coretestcases.CaseV1{
	Name:      "String: nil Pair → empty",
	WantLines: []string{""},
}

var pairStringMixedTypeTestCase = coretestcases.CaseV1{
	Name:      "String: Pair[string,int]",
	WantLines: []string{"{Left: key, Right: 42, IsValid: true}"},
}

// ==========================================================================
// Triple — IsEqual edge cases
// ==========================================================================

var tripleIsEqualSameTestCase = coretestcases.CaseV1{
	Name:      "IsEqual: same values same validity → equal",
	WantLines: []string{"true"},
}

var tripleIsEqualDiffValidityTestCase = coretestcases.CaseV1{
	Name:      "IsEqual: same values different validity → not equal",
	WantLines: []string{"false"},
}

var tripleIsEqualDiffMiddleTestCase = coretestcases.CaseV1{
	Name:      "IsEqual: different middle → not equal",
	WantLines: []string{"false"},
}

var tripleIsEqualBothNilTestCase = coretestcases.CaseV1{
	Name:      "IsEqual: both nil → equal",
	WantLines: []string{"true"},
}

var tripleIsEqualNilVsNonNilTestCase = coretestcases.CaseV1{
	Name:      "IsEqual: nil vs non-nil → not equal",
	WantLines: []string{"false"},
}

// ==========================================================================
// Triple — HasMessage edge cases
// ==========================================================================

var tripleHasMessageValidNoMsgTestCase = coretestcases.CaseV1{
	Name:      "HasMessage: valid triple no message → false",
	WantLines: []string{"false"},
}

var tripleHasMessageInvalidWithMsgTestCase = coretestcases.CaseV1{
	Name:      "HasMessage: invalid triple with message → true",
	WantLines: []string{"true"},
}

var tripleHasMessageNilTestCase = coretestcases.CaseV1{
	Name:      "HasMessage: nil triple → false",
	WantLines: []string{"false"},
}

// ==========================================================================
// Triple — IsInvalid edge cases
// ==========================================================================

var tripleIsInvalidValidTestCase = coretestcases.CaseV1{
	Name:      "IsInvalid: valid triple → false",
	WantLines: []string{"false"},
}

var tripleIsInvalidInvalidTestCase = coretestcases.CaseV1{
	Name:      "IsInvalid: invalid triple → true",
	WantLines: []string{"true"},
}

var tripleIsInvalidNilTestCase = coretestcases.CaseV1{
	Name:      "IsInvalid: nil triple → true",
	WantLines: []string{"true"},
}

// ==========================================================================
// Triple — String output
// ==========================================================================

var tripleStringValidTestCase = coretestcases.CaseV1{
	Name:      "String: valid Triple[string,string,string]",
	WantLines: []string{"{Left: a, Middle: b, Right: c, IsValid: true}"},
}

var tripleStringInvalidZeroTestCase = coretestcases.CaseV1{
	Name:      "String: invalid Triple with zero values",
	WantLines: []string{"{Left: , Middle: , Right: , IsValid: false}"},
}

var tripleStringNilTestCase = coretestcases.CaseV1{
	Name:      "String: nil Triple → empty",
	WantLines: []string{""},
}

// ==========================================================================
// Pair — NewPairWithMessage
// ==========================================================================

var pairWithMessageValidTestCase = coretestcases.CaseV1{
	Name:      "NewPairWithMessage valid with message",
	WantLines: []string{"hello", "world", "true", "ok"},
}

var pairWithMessageInvalidTestCase = coretestcases.CaseV1{
	Name:      "NewPairWithMessage invalid with error message",
	WantLines: []string{"", "", "false", "failed"},
}

// ==========================================================================
// Triple — NewTripleWithMessage
// ==========================================================================

var tripleWithMessageValidTestCase = coretestcases.CaseV1{
	Name:      "NewTripleWithMessage valid with message",
	WantLines: []string{"a", "b", "c", "true", "success"},
}

var tripleWithMessageInvalidTestCase = coretestcases.CaseV1{
	Name:      "NewTripleWithMessage invalid with error",
	WantLines: []string{"", "", "", "false", "error occurred"},
}

// ==========================================================================
// Pair — Dispose
// ==========================================================================

var pairDisposeTestCase = coretestcases.CaseV1{
	Name:      "Dispose resets pair same as Clear",
	WantLines: []string{"", "", "false", ""},
}

// ==========================================================================
// Triple — Dispose
// ==========================================================================

var tripleDisposeTestCase = coretestcases.CaseV1{
	Name:      "Dispose resets triple same as Clear",
	WantLines: []string{"", "", "", "false", ""},
}
