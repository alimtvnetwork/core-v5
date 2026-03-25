package coretestcases

import "github.com/alimtvnetwork/core/coretests"

// CaseTitle returns the test case title (Title or When fallback).
//
// Implements TypedTestCaseWrapper[TInput, TExpect].CaseTitle.
func (it *GenericGherkins[TInput, TExpect]) CaseTitle() string {
	if it.Title != "" {
		return it.Title
	}

	return it.When
}

// TypedInput returns the typed input value.
//
// Implements TypedTestCaseWrapper[TInput, TExpect].TypedInput.
func (it *GenericGherkins[TInput, TExpect]) TypedInput() TInput {
	return it.Input
}

// TypedExpected returns the typed expected value.
//
// Implements TypedTestCaseWrapper[TInput, TExpect].TypedExpected.
func (it *GenericGherkins[TInput, TExpect]) TypedExpected() TExpect {
	return it.Expected
}

// TypedActual returns the typed actual value.
//
// Implements TypedTestCaseWrapper[TInput, TExpect].TypedActual.
func (it *GenericGherkins[TInput, TExpect]) TypedActual() TExpect {
	return it.Actual
}

// SetTypedActual sets the typed actual value after the Act phase.
//
// Implements TypedTestCaseWrapper[TInput, TExpect].SetTypedActual.
func (it *GenericGherkins[TInput, TExpect]) SetTypedActual(actual TExpect) {
	it.Actual = actual
}

// AsTypedTestCaseWrapper returns the GenericGherkins as a
// TypedTestCaseWrapper[TInput, TExpect] interface.
//
// Implements TypedTestCaseWrapperContractsBinder[TInput, TExpect].
func (it *GenericGherkins[TInput, TExpect]) AsTypedTestCaseWrapper() coretests.TypedTestCaseWrapper[TInput, TExpect] {
	return it
}
