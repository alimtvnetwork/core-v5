package chmodhelpertestwrappers

type VerifyRwxPartialChmodLocationsWrapper struct {
	Locations                          []string
	IsContinueOnError, IsSkipOnInvalid bool
	Header,
	ExpectedPartialRwx string
	ExpectationErrorMessage string
}
