// aliases.go defines common result type aliases for frequently used
// return types. These are type aliases (=) to preserve %T output.
package results

// ResultAny is the untyped version of Result[T].
type ResultAny = Result[any]

// ResultBool is a Result for functions returning bool.
type ResultBool = Result[bool]

// ResultString is a Result for functions returning string.
type ResultString = Result[string]

// ResultInt is a Result for functions returning int.
type ResultInt = Result[int]

// ResultError is a Result for functions returning error.
type ResultError = Result[error]

// ResultsAny is the untyped version of Results[T1, T2].
type ResultsAny = Results[any, any]

// ResultsAnyError is for functions returning (any, error).
type ResultsAnyError = Results[any, error]

// ResultsStringError is for functions returning (string, error).
type ResultsStringError = Results[string, error]

// ResultsBoolError is for functions returning (bool, error).
type ResultsBoolError = Results[bool, error]

// ResultsIntError is for functions returning (int, error).
type ResultsIntError = Results[int, error]
