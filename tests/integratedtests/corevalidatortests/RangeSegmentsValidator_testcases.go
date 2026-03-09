package corevalidatortests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coredata/corerange"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

// ==========================================
// Shared helpers
// ==========================================

// actualLines provides a 5-element slice for range-based slicing.
var rangeSegActualLines = []string{"line0", "line1", "line2", "line3", "line4"}

func newMatchingRangeSegment(start, end int) corevalidator.RangesSegment {
	return corevalidator.RangesSegment{
		RangeInt: corerange.RangeInt{
			Start: start,
			End:   end,
		},
		ExpectedLines: rangeSegActualLines[start:end],
		CompareAs:     stringcompareas.Equal,
		Condition:     corevalidator.DefaultDisabledCoreCondition,
	}
}

func newMismatchRangeSegment(start, end int) corevalidator.RangesSegment {
	return corevalidator.RangesSegment{
		RangeInt: corerange.RangeInt{
			Start: start,
			End:   end,
		},
		ExpectedLines: []string{"WRONG", "DATA"},
		CompareAs:     stringcompareas.Equal,
		Condition:     corevalidator.DefaultDisabledCoreCondition,
	}
}

// ==========================================
// LengthOfVerifierSegments
// ==========================================

var rangeSegmentsValidatorLengthTestCases = []coretestcases.CaseV1{
	{
		Title: "No segments returns 0",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title:            "empty",
			VerifierSegments: nil,
		},
		ExpectedInput: args.Map{"length": 0},
	},
	{
		Title: "One segment returns 1",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "one",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(0, 2),
			},
		},
		ExpectedInput: args.Map{"length": 1},
	},
	{
		Title: "Two segments returns 2",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "two",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(0, 2),
				newMatchingRangeSegment(2, 4),
			},
		},
		ExpectedInput: args.Map{"length": 2},
	},
}

// ==========================================
// Validators
// ==========================================

var rangeSegmentsValidatorValidatorsTestCases = []coretestcases.CaseV1{
	{
		Title: "Produces HeaderSliceValidators from segments",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "test",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(0, 3),
			},
		},
		ExpectedInput: args.Map{
			"hasValidators": true,
		},
	},
}

// ==========================================
// VerifyAll
// ==========================================

var rangeSegmentsValidatorVerifyAllTestCases = []coretestcases.CaseV1{
	{
		Title: "Matching segment returns nil error",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "match",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(0, 3),
			},
		},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "Mismatching segment returns error",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "mismatch",
			VerifierSegments: []corevalidator.RangesSegment{
				newMismatchRangeSegment(0, 2),
			},
		},
		ExpectedInput: args.Map{"hasError": true},
	},
}

// ==========================================
// VerifySimple
// ==========================================

var rangeSegmentsValidatorVerifySimpleTestCases = []coretestcases.CaseV1{
	{
		Title: "Matching returns nil error",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "simple-match",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(1, 3),
			},
		},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "Mismatch returns error",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "simple-mismatch",
			VerifierSegments: []corevalidator.RangesSegment{
				newMismatchRangeSegment(0, 2),
			},
		},
		ExpectedInput: args.Map{"hasError": true},
	},
}

// ==========================================
// VerifyFirst
// ==========================================

var rangeSegmentsValidatorVerifyFirstTestCases = []coretestcases.CaseV1{
	{
		Title: "Matching returns nil error",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "first-match",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(0, 2),
			},
		},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "Mismatch returns error",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "first-mismatch",
			VerifierSegments: []corevalidator.RangesSegment{
				newMismatchRangeSegment(0, 2),
			},
		},
		ExpectedInput: args.Map{"hasError": true},
	},
}

// ==========================================
// VerifyUpto
// ==========================================

var rangeSegmentsValidatorVerifyUptoTestCases = []coretestcases.CaseV1{
	{
		Title: "Matching within upto length returns nil error",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "upto-match",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(0, 3),
			},
		},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "Mismatch returns error",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "upto-mismatch",
			VerifierSegments: []corevalidator.RangesSegment{
				newMismatchRangeSegment(0, 2),
			},
		},
		ExpectedInput: args.Map{"hasError": true},
	},
}

// ==========================================
// VerifyFirstDefault / VerifyUptoDefault
// ==========================================

var rangeSegmentsValidatorVerifyFirstDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "Default uses Title as header, matching returns nil",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "default-first",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(0, 2),
			},
		},
		ExpectedInput: args.Map{"hasError": false},
	},
}

var rangeSegmentsValidatorVerifyUptoDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "Default uses Title as header, matching returns nil",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "default-upto",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(0, 3),
			},
		},
		ExpectedInput: args.Map{"hasError": false},
	},
}
