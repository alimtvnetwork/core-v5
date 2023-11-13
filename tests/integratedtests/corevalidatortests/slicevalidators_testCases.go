package corevalidatortests

import (
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/tests/testwrappers/corevalidatortestwrappers"
)

var sliceValidatorsTestCases = []corevalidatortestwrappers.SliceValidatorsWrapper{
	{
		Header: "Comparing all flag to false, and comparing equal.",
		ComparingLines: []string{
			"alim      alim 2 alim 4",
		},
		Validators: corevalidator.SliceValidators{
			Validators: []corevalidator.SliceValidator{},
		},
		IsSkipOnContentsEmpty: false,
		IsCaseSensitive:       true,
		ExpectationLines: []string{
			"",
			"0 )   Header: `Comparing all flag to false, and comparing equal.`",
			"----- Method: `\"Equal\"`, Line Index: 0",
			"",
			"--------------- Actual:",
			"`\"alim      alim 2 alim 4\"`",
			"",
			"--- Expected or Search:",
			"`\"   alim      alim 2 alim 3                 \"`",
			"",
			"Additional: `Search Input: [`   alim      alim 2 alim 3                 `], CompareMethod: [`Equal`], IsTrimCompare: [`false`], IsSplitByWhitespace: [`false`], IsUniqueWordOnly: [`false`], IsNonEmptyWhitespace: [`false`], IsSortStringsBySpace: [`false`]`",
		},
	},
}
