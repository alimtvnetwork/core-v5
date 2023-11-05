package coretests

import "testing"

type BaseTestCaseWrapper interface {
	SimpleTestCaseWrapper
	ShouldAsserter
	TypeValidationError() error
	TypesValidationMustPasses(t *testing.T)
	ArrangeString() string
	ActualString() string
	SetActual(actual interface{})
	String(caseIndex int) string
	IsDisabled() bool
	asSimpleTestCaseWrapper() SimpleTestCaseWrapper
}
