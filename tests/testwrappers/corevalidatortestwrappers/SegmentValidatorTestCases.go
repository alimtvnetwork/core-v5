package corevalidatortestwrappers

import (
	"gitlab.com/evatix-go/core/coredata/corerange"
	"gitlab.com/evatix-go/core/corevalidator"
	"gitlab.com/evatix-go/core/enums/stringcompareas"
)

var SegmentValidatorTestCases = []SegmentValidatorWrapper{
	{
		Header:                "",
		IsSkipOnContentsEmpty: false,
		IsCaseSensitive:       true,
		SimpleSliceRangeValidator: corevalidator.SimpleSliceRangeValidator{
			VerifierSegments: []corevalidator.RangesSegment{
				{
					RangeInt:      corerange.RangeInt{},
					ExpectedLines: nil,
					CompareAs:     stringcompareas.Equal,
					ValidatorCoreCondition: corevalidator.ValidatorCoreCondition{
						IsTrimCompare:        false,
						IsUniqueWordOnly:     false,
						IsNonEmptyWhitespace: false,
						IsSortStringsBySpace: false,
					},
				},
			},
		},
	},
}
