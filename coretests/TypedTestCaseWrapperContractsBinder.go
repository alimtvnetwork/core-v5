package coretests

// TypedTestCaseWrapperContractsBinder is the generic-first contracts binder.
//
// Embeds TypedTestCaseWrapper[TInput, TExpect] and adds a self-accessor.
type TypedTestCaseWrapperContractsBinder[TInput, TExpect any] interface {
	TypedTestCaseWrapper[TInput, TExpect]
	AsTypedTestCaseWrapper() TypedTestCaseWrapper[TInput, TExpect]
}
