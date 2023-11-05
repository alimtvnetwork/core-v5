package coretests

type SimpleTestCaseWrapperContractsBinder interface {
	SimpleTestCaseWrapper
	asSimpleTestCaseWrapper() SimpleTestCaseWrapper
}
