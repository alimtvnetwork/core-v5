package enumimpltests

import (
	"gitlab.com/auk-go/core/coretests"
)

var dynamicMapDiffMessageTestCasesV2 = []EnumImplDynamicMapTestWrapper{
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map diff string compiled must be same",
			ArrangeInput: LeftRightDynamicMapWithDefaultChecker{
				LeftRightDynamicMap: LeftRightDynamicMap{
					Left: map[string]interface{}{
						"exist":                        1,
						"not-exist-in-right":           3,
						"exist-in-left-right-diff-val": 5,
					},
					Right: map[string]interface{}{
						"exist":                        1,
						"not-exist-in-left":            2,
						"exist-in-left-right-diff-val": 6,
					},
				},
				DifferChecker: checker1,
			},
			ActualInput: nil,
			ExpectedInput: []string{
				"Dynamic map diff string compiled must be same",
				"",
				"Difference Between Map:",
				"",
				"{",
				"- Left Map - Has Diff from Right Map:",
				"{",
				"",
				"\"exist-in-left-right-diff-val\":\"{\"Left\":5,\"Right\":6}\",",
				"\"not-exist-in-left\":\"{\"Left\":null,\"Right\":2}\"",
				"",
				"}",
				"",
				"",
				"- Right Map - Has Diff from Left Map:",
				"{",
				"",
				"\"exist-in-left-right-diff-val\":\"{\"Left\":5,\"Right\":6}\",",
				"\"not-exist-in-right\":\"3 (type:int) - left - key is missing!\"",
				"",
				"}}",
			},
			VerifyTypeOf:    typeVerifyOfForDynamicMapSimpleDiffTestCases,
			HasError:        false,
			IsValidateError: true,
		},
	},
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map diff - no changes",
			ArrangeInput: LeftRightDynamicMapWithDefaultChecker{
				LeftRightDynamicMap: LeftRightDynamicMap{
					Left: map[string]interface{}{
						"a":  1,
						"b":  3,
						"cl": 5,
					},
					Right: map[string]interface{}{
						"a":  1,
						"b":  3,
						"cl": 5,
					},
				},
				DifferChecker: checker1,
			},
			ActualInput:     nil,
			ExpectedInput:   "",
			VerifyTypeOf:    simpleDiffTestCases,
			HasError:        false,
			IsValidateError: true,
		},
	},
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map diff - right hand key missing",
			ArrangeInput: LeftRightDynamicMapWithDefaultChecker{
				LeftRightDynamicMap: LeftRightDynamicMap{
					Left: map[string]interface{}{
						"a":  1,
						"b":  3,
						"cl": 5,
					},
					Right: map[string]interface{}{
						"a": 1,
						"b": 3,
					},
				},
				DifferChecker: checker1,
			},
			ActualInput: nil,
			ExpectedInput: []string{
				"Dynamic map diff - right hand key missing",
				"",
				"Difference Between Map:",
				"",
				"{",
				"- Right Map - Has Diff from Left Map:",
				"{",
				"",
				"\"cl\":\"5 (type:int) - left - key is missing!\"",
				"",
				"}}",
			},
			VerifyTypeOf:    typeVerifyOfForDynamicMapSimpleDiffTestCases,
			HasError:        false,
			IsValidateError: true,
		},
	},
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map diff - left hand key missing",
			ArrangeInput: LeftRightDynamicMapWithDefaultChecker{
				LeftRightDynamicMap: LeftRightDynamicMap{
					Left: map[string]interface{}{
						"a": 1,
						"b": 3,
					},
					Right: map[string]interface{}{
						"a":  1,
						"b":  3,
						"cl": 5,
					},
				},
				DifferChecker: checker1,
			},
			ActualInput: nil,
			ExpectedInput: "Dynamic map diff - left hand key missing\n\n" +
				"Difference Between Map:\n\n{\n" +
				"- Left Map - Has Diff from Right Map:" +
				"\n{\n\n\"" +
				"cl\":\"{\"Left\":null,\"Right\":5}\"\n\n}\n}",
			VerifyTypeOf:    typeVerifyOfForDynamicMapSimpleDiffTestCases,
			HasError:        false,
			IsValidateError: true,
		},
	},
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "Dynamic map diff - left hand key missing",
			ArrangeInput: LeftRightDynamicMapWithDefaultChecker{
				LeftRightDynamicMap: LeftRightDynamicMap{
					Left: map[string]interface{}{
						"a": 1,
						"b": 3,
					},
					Right: map[string]interface{}{
						"a":  1,
						"b":  3,
						"cl": 5,
					},
				},
				DifferChecker: checker1,
			},
			ActualInput: nil,
			ExpectedInput: "Dynamic map diff - left hand key missing\n\n" +
				"Difference Between Map:\n\n{\n" +
				"- Left Map - Has Diff from Right Map:" +
				"\n{\n\n\"" +
				"cl\":\"{\"Left\":null,\"Right\":5}\"\n\n}\n}",
			VerifyTypeOf:    typeVerifyOfForDynamicMapSimpleDiffTestCases,
			HasError:        false,
			IsValidateError: true,
		},
	},
}
