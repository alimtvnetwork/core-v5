package corevalidatortestwrappers

import (
	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

var SegmentValidatorTestCases = []SegmentValidatorWrapper{
	{
		Header:                "",
		IsSkipOnContentsEmpty: false,
		IsCaseSensitive:       true,
		RangeSegmentsValidator: corevalidator.RangeSegmentsValidator{
			VerifierSegments: []corevalidator.RangesSegment{

				{
					RangeInt:      corerange.RangeInt{},
					ExpectedLines: nil,
					CompareAs:     stringcompareas.Equal,
					Condition:     corevalidator.DefaultDisabledCoreCondition,
				},
			},
		},
	},
}
