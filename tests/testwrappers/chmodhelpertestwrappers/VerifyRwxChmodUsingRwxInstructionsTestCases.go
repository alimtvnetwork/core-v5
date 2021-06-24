package chmodhelpertestwrappers

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

var VerifyRwxChmodUsingRwxInstructionsTestCases = []VerifyRwxChmodUsingRwxInstructionsWrapper{
	{
		Header: "rwx",
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
		Locations: SimpleLocations,
		ExpectedErrorMessage: "Missing or paths having other access issues! Reference(s) { \"[" +
			"/temp/core/test-cases-3s " +
			"/temp/core/test-cases-3x" +
			"]\" }",
	},
	{
		Header: "rwx",
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
		Locations: SimpleLocations,
		ExpectedErrorMessage: "Path:/temp/core/test-cases-2 - " +
			"Expect [\"rwxr-x---\"] != [\"rwxr-xr--\"] Actual\n" +
			"Path:/temp/core/test-cases-3 - " +
			"Expect [\"rwxr-x---\"] != [\"rwxr-xr--\"] Actual",
	},
	{
		Header: "rwx",
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
		Locations: SimpleLocations,
		ExpectedErrorMessage: "Not Supported: Feature / method is not supported yet. " +
			"Condition.IsRecursive is not supported for Verify chmod. " +
			"Reference(s) { \"[" +
			"/temp/core/test-cases-2 " +
			"/temp/core/test-cases-3s " +
			"/temp/core/test-cases-3x " +
			"/temp/core/test-cases-3" +
			"]\" }",
	},
	{
		Header: "rwx",
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
		Locations: SimpleLocations,
		ExpectedErrorMessage: "Missing or paths having other access issues! Reference(s) { " +
			"\"" +
			"[" +
			"/temp/core/test-cases-3s " +
			"/temp/core/test-cases-3x" +
			"]" +
			"\" }\n" +
			"Path:" +
			"/temp/core/test-cases-3 - " +
			"Expect [\"rwxr-x---\"] != [\"rwxr-xr--\"] Actual\n" +
			"Path:/temp/core/test-cases-2 - " +
			"Expect [\"rwxr-x---\"] != [\"rwxr-xr--\"] Actual",
	},
}
