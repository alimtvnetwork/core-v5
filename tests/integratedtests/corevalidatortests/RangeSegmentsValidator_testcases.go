package corevalidatortests

import (
	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
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
		Title: "LengthOfVerifierSegments returns 0 -- no segments",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title:            "empty",
			VerifierSegments: nil,
		},
		ExpectedInput: args.Map{"length": 0},
	},
	{
		Title: "LengthOfVerifierSegments returns 1 -- one segment",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "one",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(0, 2),
			},
		},
		ExpectedInput: args.Map{"length": 1},
	},
	{
		Title: "LengthOfVerifierSegments returns 2 -- two segments",
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
		Title: "Validators returns HeaderSliceValidators -- one segment input",
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
		Title: "VerifyAll returns nil -- matching segment",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "match",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(0, 3),
			},
		},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "VerifyAll returns error -- mismatching segment",
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
		Title: "VerifySimple returns nil -- matching segment range 1-3",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "simple-match",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(1, 3),
			},
		},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "VerifySimple returns error -- mismatched segment range 0-2",
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
		Title: "VerifyFirst returns nil -- matching segment range 0-2",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "first-match",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(0, 2),
			},
		},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "VerifyFirst returns error -- mismatched segment range 0-2",
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
		Title: "VerifyUpto returns nil -- matching segment within length",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "upto-match",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(0, 3),
			},
		},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title: "VerifyUpto returns error -- mismatched segment range 0-2",
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
		Title: "VerifyFirstDefault returns nil -- matching segment with Title as header",
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
		Title: "VerifyUptoDefault returns nil -- matching segment with Title as header",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "default-upto",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(0, 3),
			},
		},
		ExpectedInput: args.Map{"hasError": false},
	},
}

// ==========================================
// SetActual
// ==========================================

var rangeSegmentsValidatorSetActualTestCases = []coretestcases.CaseV1{
	{
		Title: "SetActual returns self and match true -- matching segment propagated",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "set-actual-match",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(0, 3),
			},
		},
		ExpectedInput: args.Map{
			"returnsSelf": true,
			"isMatch":     true,
		},
	},
	{
		Title: "SetActual returns self and match false -- mismatch segment propagated",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "set-actual-mismatch",
			VerifierSegments: []corevalidator.RangesSegment{
				newMismatchRangeSegment(0, 2),
			},
		},
		ExpectedInput: args.Map{
			"returnsSelf": true,
			"isMatch":     false,
		},
	},
}

// ==========================================
// SetActualOnAll (via Validators)
// ==========================================

var rangeSegmentsValidatorSetActualOnAllTestCases = []coretestcases.CaseV1{
	{
		Title: "SetActualOnAll returns match true -- all segments matching",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "set-all-match",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(0, 3),
				newMatchingRangeSegment(3, 5),
			},
		},
		ExpectedInput: args.Map{
			"validatorCount": 2,
			"isMatch":        true,
		},
	},
	{
		Title: "SetActualOnAll returns match false -- one segment mismatched",
		ArrangeInput: &corevalidator.RangeSegmentsValidator{
			Title: "set-all-mismatch",
			VerifierSegments: []corevalidator.RangesSegment{
				newMatchingRangeSegment(0, 3),
				newMismatchRangeSegment(3, 5),
			},
		},
		ExpectedInput: args.Map{
			"validatorCount": 2,
			"isMatch":        false,
		},
	},
}
