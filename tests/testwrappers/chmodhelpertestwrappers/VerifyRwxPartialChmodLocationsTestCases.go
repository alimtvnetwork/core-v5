package chmodhelpertestwrappers

var VerifyRwxPartialChmodLocationsTestCases = []VerifyRwxPartialChmodLocationsWrapper{
	{
		Header:             "Missing Paths should NOT have error with it's location!",
		Locations:          SimpleLocations,
		IsContinueOnError:  true,
		IsSkipOnInvalid:    true,
		ExpectedPartialRwx: "-rwxrwx",
		ExpectationErrorMessage: "/tmp/core/test-cases-2 - " +
			"Expect [\"rwxrwx***\"] != [\"rwxr-xr--\"] Actual\n" +
			"/tmp/core/test-cases-3 - " +
			"Expect [\"rwxrwx***\"] != [\"rwxr-xr--\"] Actual",
	},
	{
		Header:                  "Missing Paths should NOT have error with it's location and all matches with WhatIsExpected RWX!",
		Locations:               SimpleLocations,
		IsContinueOnError:       true,
		IsSkipOnInvalid:         true,
		ExpectedPartialRwx:      "-rwx",
		ExpectationErrorMessage: "",
	},
	{
		Header:             "Missing Paths should have error with it's location!",
		Locations:          SimpleLocations,
		IsContinueOnError:  true,
		IsSkipOnInvalid:    false,
		ExpectedPartialRwx: "-rwxrwx-",
		ExpectationErrorMessage: "/tmp/core/test-cases-2 - " +
			"Expect [\"rwxrwx-**\"] != [\"rwxr-xr--\"] Actual\n" +
			"/tmp/core/test-cases-3 - " +
			"Expect [\"rwxrwx-**\"] != [\"rwxr-xr--\"] Actual\n" +
			"Path missing or having other access issues! Ref(s) { \"" +
			"[/tmp/core/test-cases-3s " +
			"/tmp/core/test-cases-3x]\" }",
	},
}
