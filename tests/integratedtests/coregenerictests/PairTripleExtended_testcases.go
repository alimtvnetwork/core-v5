package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Pair — IsEqual edge cases (mixed validity, same values diff validity)
// ==========================================================================

var pairIsEqualExtendedTestCases = []coretestcases.CaseV1{
	{
		Name: "IsEqual: same values, different validity → not equal",
		WantLines: []string{
			"false",
		},
	},
	{
		Name: "IsEqual: different right values → not equal",
		WantLines: []string{
			"false",
		},
	},
	{
		Name: "IsEqual: both invalid with same zero values → equal",
		WantLines: []string{
			"true",
		},
	},
	{
		Name: "IsEqual: Pair[int,int] same values → equal",
		WantLines: []string{
			"true",
		},
	},
	{
		Name: "IsEqual: Pair[int,int] different values → not equal",
		WantLines: []string{
			"false",
		},
	},
	{
		Name: "IsEqual: Pair[string,int] mixed types → equal",
		WantLines: []string{
			"true",
		},
	},
}

// ==========================================================================
// Pair — HasMessage edge cases
// ==========================================================================

var pairHasMessageTestCases = []coretestcases.CaseV1{
	{
		Name: "HasMessage: valid pair with no message → false",
		WantLines: []string{
			"false",
		},
	},
	{
		Name: "HasMessage: invalid pair with message → true",
		WantLines: []string{
			"true",
		},
	},
	{
		Name: "HasMessage: pair with whitespace-only message → true",
		WantLines: []string{
			"true",
		},
	},
	{
		Name: "HasMessage: nil pair → false",
		WantLines: []string{
			"false",
		},
	},
}

// ==========================================================================
// Pair — IsInvalid edge cases
// ==========================================================================

var pairIsInvalidTestCases = []coretestcases.CaseV1{
	{
		Name: "IsInvalid: valid pair → false",
		WantLines: []string{
			"false",
		},
	},
	{
		Name: "IsInvalid: invalid pair → true",
		WantLines: []string{
			"true",
		},
	},
	{
		Name: "IsInvalid: nil pair → true",
		WantLines: []string{
			"true",
		},
	},
}

// ==========================================================================
// Pair — String output
// ==========================================================================

var pairStringTestCases = []coretestcases.CaseV1{
	{
		Name: "String: valid Pair[string,string]",
		WantLines: []string{
			"{Left: hello, Right: world, IsValid: true}",
		},
	},
	{
		Name: "String: invalid Pair with zero values",
		WantLines: []string{
			"{Left: , Right: , IsValid: false}",
		},
	},
	{
		Name: "String: nil Pair → empty",
		WantLines: []string{
			"",
		},
	},
	{
		Name: "String: Pair[string,int]",
		WantLines: []string{
			"{Left: key, Right: 42, IsValid: true}",
		},
	},
}

// ==========================================================================
// Triple — IsEqual edge cases
// ==========================================================================

var tripleIsEqualExtendedTestCases = []coretestcases.CaseV1{
	{
		Name: "IsEqual: same values same validity → equal",
		WantLines: []string{
			"true",
		},
	},
	{
		Name: "IsEqual: same values different validity → not equal",
		WantLines: []string{
			"false",
		},
	},
	{
		Name: "IsEqual: different middle → not equal",
		WantLines: []string{
			"false",
		},
	},
	{
		Name: "IsEqual: both nil → equal",
		WantLines: []string{
			"true",
		},
	},
	{
		Name: "IsEqual: nil vs non-nil → not equal",
		WantLines: []string{
			"false",
		},
	},
}

// ==========================================================================
// Triple — HasMessage edge cases
// ==========================================================================

var tripleHasMessageTestCases = []coretestcases.CaseV1{
	{
		Name: "HasMessage: valid triple no message → false",
		WantLines: []string{
			"false",
		},
	},
	{
		Name: "HasMessage: invalid triple with message → true",
		WantLines: []string{
			"true",
		},
	},
	{
		Name: "HasMessage: nil triple → false",
		WantLines: []string{
			"false",
		},
	},
}

// ==========================================================================
// Triple — IsInvalid edge cases
// ==========================================================================

var tripleIsInvalidTestCases = []coretestcases.CaseV1{
	{
		Name: "IsInvalid: valid triple → false",
		WantLines: []string{
			"false",
		},
	},
	{
		Name: "IsInvalid: invalid triple → true",
		WantLines: []string{
			"true",
		},
	},
	{
		Name: "IsInvalid: nil triple → true",
		WantLines: []string{
			"true",
		},
	},
}

// ==========================================================================
// Triple — String output
// ==========================================================================

var tripleStringTestCases = []coretestcases.CaseV1{
	{
		Name: "String: valid Triple[string,string,string]",
		WantLines: []string{
			"{Left: a, Middle: b, Right: c, IsValid: true}",
		},
	},
	{
		Name: "String: invalid Triple with zero values",
		WantLines: []string{
			"{Left: , Middle: , Right: , IsValid: false}",
		},
	},
	{
		Name: "String: nil Triple → empty",
		WantLines: []string{
			"",
		},
	},
}

// ==========================================================================
// Pair — NewPairWithMessage
// ==========================================================================

var pairWithMessageTestCases = []coretestcases.CaseV1{
	{
		Name: "NewPairWithMessage valid with message",
		WantLines: []string{
			"hello",
			"world",
			"true",
			"ok",
		},
	},
	{
		Name: "NewPairWithMessage invalid with error message",
		WantLines: []string{
			"",
			"",
			"false",
			"failed",
		},
	},
}

// ==========================================================================
// Triple — NewTripleWithMessage
// ==========================================================================

var tripleWithMessageTestCases = []coretestcases.CaseV1{
	{
		Name: "NewTripleWithMessage valid with message",
		WantLines: []string{
			"a",
			"b",
			"c",
			"true",
			"success",
		},
	},
	{
		Name: "NewTripleWithMessage invalid with error",
		WantLines: []string{
			"",
			"",
			"",
			"false",
			"error occurred",
		},
	},
}

// ==========================================================================
// Pair — Dispose (alias for Clear)
// ==========================================================================

var pairDisposeTestCases = []coretestcases.CaseV1{
	{
		Name: "Dispose resets pair same as Clear",
		WantLines: []string{
			"",
			"",
			"false",
			"",
		},
	},
}

// ==========================================================================
// Triple — Dispose
// ==========================================================================

var tripleDisposeTestCases = []coretestcases.CaseV1{
	{
		Name: "Dispose resets triple same as Clear",
		WantLines: []string{
			"",
			"",
			"",
			"false",
			"",
		},
	},
}
