package chmodhelpertests

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/tests/testwrappers/chmodhelpertestwrappers"
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
			"/tmp/core/test-cases-2 - " +
				"Expect [\"rwxrwx***\"] != [\"rwxr-xr--\"] Actual",
			"/tmp/core/test-cases-3 - " +
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
		ExpectedInput: "",
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
			"/tmp/core/test-cases-2 - " +
				"Expect [\"rwxrwx-**\"] != [\"rwxr-xr--\"] Actual",
			"/tmp/core/test-cases-3 - " +
				"Expect [\"rwxrwx-**\"] != [\"rwxr-xr--\"] Actual",
			"Path missing or having other access issues! Ref(s) { " +
				"\"[/tmp/core/test-cases-3s " +
				"/tmp/core/test-cases-3x]\" }",
		},
	},
}
