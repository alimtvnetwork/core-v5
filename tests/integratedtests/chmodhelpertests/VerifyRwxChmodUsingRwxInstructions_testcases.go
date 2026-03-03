package chmodhelpertests

import (
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

var verifyRwxChmodUsingRwxInstructionsTestCases = []coretestcases.CaseV1{
	{
		Title: "rwx - missing paths",
		ArrangeInput: chmodhelpertestwrappers.VerifyRwxChmodUsingRwxInstructionsWrapper{
			RwxInstruction: chmodins.RwxInstruction{
				RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
					Owner: "rwx",
					Group: "rwx",
					Other: "---",
				},
				Condition: chmodins.Condition{
					IsSkipOnInvalid:   false,
					IsContinueOnError: false,
					IsRecursive:       false,
				},
			},
			Locations: chmodhelpertestwrappers.SimpleLocations,
		},
		ExpectedInput: []string{
			"Path missing or having other access issues! Ref(s) { " +
				"\"[/temp/core/test-cases-3s /temp/core/test-cases-3x]\" }",
		},
	},
	{
		Title: "rwx - expectation failed",
		ArrangeInput: chmodhelpertestwrappers.VerifyRwxChmodUsingRwxInstructionsWrapper{
			RwxInstruction: chmodins.RwxInstruction{
				RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
					Owner: "rwx",
					Group: "r-x",
					Other: "---",
				},
				Condition: chmodins.Condition{
					IsSkipOnInvalid:   true,
					IsContinueOnError: true,
					IsRecursive:       false,
				},
			},
			Locations: chmodhelpertestwrappers.SimpleLocations,
		},
		ExpectedInput: []string{
			"Path:/temp/core/test-cases-2 - " +
				"Expect [\"rwxr-x---\"] != [\"rwxr-xr--\"] Actual",
			"Path:/temp/core/test-cases-3 - " +
				"Expect [\"rwxr-x---\"] != [\"rwxr-xr--\"] Actual",
		},
	},
	{
		Title: "Recursive not supported",
		ArrangeInput: chmodhelpertestwrappers.VerifyRwxChmodUsingRwxInstructionsWrapper{
			RwxInstruction: chmodins.RwxInstruction{
				RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
					Owner: "rwx",
					Group: "r-x",
					Other: "---",
				},
				Condition: chmodins.Condition{
					IsSkipOnInvalid:   true,
					IsContinueOnError: true,
					IsRecursive:       true,
				},
			},
			Locations: chmodhelpertestwrappers.SimpleLocations,
		},
		ExpectedInput: []string{
			"Not Supported: Feature or method is not supported yet. " +
				"IsRecursive is not supported for Verify chmod. Ref(s) { " +
				"\"[" +
				"/temp/core/test-cases-2 " +
				"/temp/core/test-cases-3s " +
				"/temp/core/test-cases-3x " +
				"/temp/core/test-cases-3]\" }",
		},
	},
	{
		Title: "Missing paths + Expectation failed",
		ArrangeInput: chmodhelpertestwrappers.VerifyRwxChmodUsingRwxInstructionsWrapper{
			RwxInstruction: chmodins.RwxInstruction{
				RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
					Owner: "rwx",
					Group: "r-x",
					Other: "---",
				},
				Condition: chmodins.Condition{
					IsSkipOnInvalid:   false,
					IsContinueOnError: true,
					IsRecursive:       false,
				},
			},
			Locations: chmodhelpertestwrappers.SimpleLocations,
		},
		ExpectedInput: []string{
			"Path missing or having other access issues! Ref(s) { " +
				"\"[/temp/core/test-cases-3s /temp/core/test-cases-3x]\" " +
				"}",
			"Path:/temp/core/test-cases-2 - " +
				"Expect [\"rwxr-x---\"] != [\"rwxr-xr--\"] Actual",
			"Path:/temp/core/test-cases-3 - " +
				"Expect [\"rwxr-x---\"] != [\"rwxr-xr--\"] Actual",
		},
	},
	{
		Title: "Expectation and missing paths, isContinue false so will fail for missing paths only",
		ArrangeInput: chmodhelpertestwrappers.VerifyRwxChmodUsingRwxInstructionsWrapper{
			RwxInstruction: chmodins.RwxInstruction{
				RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
					Owner: "rwx",
					Group: "r-x",
					Other: "---",
				},
				Condition: chmodins.Condition{
					IsSkipOnInvalid:   false,
					IsContinueOnError: false,
					IsRecursive:       false,
				},
			},
			Locations: chmodhelpertestwrappers.SimpleLocations,
		},
		ExpectedInput: []string{
			"Path missing or having other access issues! Ref(s) { " +
				"\"[/temp/core/test-cases-3s /temp/core/test-cases-3x]\" }",
		},
	},
}
