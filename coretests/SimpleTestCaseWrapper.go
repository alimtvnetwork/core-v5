package coretests

// SimpleTestCaseWrapper is the legacy all-any test case wrapper.
//
// For new code, prefer TypedTestCaseWrapper[TInput, TExpect]
// which provides compile-time type safety.
//
// This interface is retained for backward compatibility with CaseV1,
// BaseTestCase, SimpleTestCase, and existing test infrastructure.
type SimpleTestCaseWrapper interface {
	CaseTitle() string
	Input() any
	Expected() any
	Actual() any
	SetActual(actual any)
}
