package coretests

// TypedTestCaseWrapper is the generic-first test case wrapper interface.
//
// TInput  — type of the test input
// TExpect — type of the expected result
//
// Provides compile-time type safety for test case data access.
// Legacy code should use SimpleTestCaseWrapper (= TypedTestCaseWrapper[any, any]).
type TypedTestCaseWrapper[TInput, TExpect any] interface {
	CaseTitle() string
	TypedInput() TInput
	TypedExpected() TExpect
	TypedActual() TExpect
	SetTypedActual(actual TExpect)
}
