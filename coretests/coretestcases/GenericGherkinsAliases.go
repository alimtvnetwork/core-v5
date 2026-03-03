package coretestcases

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
