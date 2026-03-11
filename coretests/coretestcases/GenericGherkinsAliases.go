package coretestcases

import (
	"github.com/alimtvnetwork/core/coretests/args"
)

// AnyGherkins is a GenericGherkins with all-any typed fields.
//
// Use when input and expected types are heterogeneous or unknown
// at compile time.
type AnyGherkins = GenericGherkins[any, any]

// StringGherkins is a GenericGherkins with string input and string expected.
//
// Use for text-based validation tests where both input and expected
// are plain strings.
type StringGherkins = GenericGherkins[string, string]

// StringBoolGherkins is a GenericGherkins with string input and bool expected.
//
// Use for matching/validation tests (e.g., regex, search) where input
// is a string and the expected outcome is a boolean.
type StringBoolGherkins = GenericGherkins[string, bool]

// MapGherkins is a GenericGherkins with args.Map for both input and expected.
//
// Use when test inputs and expectations are multi-field key-value pairs.
// Input holds arrange data (e.g., pattern, compareInput).
// Expected holds assertion data (e.g., isDefined, isApplicable, isMatch).
//
// This replaces opaque ExpectedLines with self-documenting semantic keys,
// making test cases readable without consulting the test runner.
//
// Example:
//
//	var testCases = []coretestcases.MapGherkins{
//	    {
//	        Title: "Lazy regex matches word pattern",
//	        When:  "given a simple word pattern",
//	        Input: args.Map{
//	            "pattern":      "hello",
//	            "compareInput": "hello world",
//	        },
//	        Expected: args.Map{
//	            "isDefined":    true,
//	            "isApplicable": true,
//	            "isMatch":      true,
//	            "isFailedMatch": false,
//	        },
//	    },
//	}
type MapGherkins = GenericGherkins[args.Map, args.Map]
