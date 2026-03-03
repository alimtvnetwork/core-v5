package chmodhelpertests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

var verifyPartialRwxLocationsTestCases = []coretestcases.CaseV1{
	{
		Title: "Missing Paths should NOT have error with it's location!",
		ArrangeInput: chmodhelpertestwrappers.VerifyRwxPartialChmodLocationsWrapper{
			Locations:          chmodhelpertestwrappers.SimpleLocations,
			IsContinueOnError:  true,
			IsSkipOnInvalid:    true,
			ExpectedPartialRwx: "-rwxrwx",
		},
		ExpectedInput: []string{
			"/temp/core/test-cases-2 - " +
				"Expect [\"rwxrwx***\"] != [\"rwxr-xr--\"] Actual",
			"/temp/core/test-cases-3 - " +
				"Expect [\"rwxrwx***\"] != [\"rwxr-xr--\"] Actual",
		},
	},
	{
		Title: "Missing Paths should NOT have error and all matches with expected RWX!",
		ArrangeInput: chmodhelpertestwrappers.VerifyRwxPartialChmodLocationsWrapper{
			Locations:          chmodhelpertestwrappers.SimpleLocations,
			IsContinueOnError:  true,
			IsSkipOnInvalid:    true,
			ExpectedPartialRwx: "-rwx",
		},
		ExpectedInput: []string{""},
	},
	{
		Title: "Missing Paths should have error with it's location!",
		ArrangeInput: chmodhelpertestwrappers.VerifyRwxPartialChmodLocationsWrapper{
			Locations:          chmodhelpertestwrappers.SimpleLocations,
			IsContinueOnError:  true,
			IsSkipOnInvalid:    false,
			ExpectedPartialRwx: "-rwxrwx-",
		},
		ExpectedInput: []string{
			"/temp/core/test-cases-2 - " +
				"Expect [\"rwxrwx-**\"] != [\"rwxr-xr--\"] Actual",
			"/temp/core/test-cases-3 - " +
				"Expect [\"rwxrwx-**\"] != [\"rwxr-xr--\"] Actual",
			"Path missing or having other access issues! Ref(s) { " +
				"\"[/temp/core/test-cases-3s " +
				"/temp/core/test-cases-3x]\" }",
		},
	},
}
