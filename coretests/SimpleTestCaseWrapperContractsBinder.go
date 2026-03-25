package coretests

// SimpleTestCaseWrapperContractsBinder is the legacy contracts binder.
//
// For new code, prefer TypedTestCaseWrapperContractsBinder[TInput, TExpect].
type SimpleTestCaseWrapperContractsBinder interface {
	SimpleTestCaseWrapper
	AsSimpleTestCaseWrapper() SimpleTestCaseWrapper
}
