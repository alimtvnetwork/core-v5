package codegen

import "gitlab.com/auk-go/core/coredata/corestr"

type testCaseGenerator struct {
	baseGenerator BaseGenerator
}

func (it testCaseGenerator) Compile() (string, error) {
	it.caseItems()

	return "", nil
}

func (it testCaseGenerator) caseItems() *corestr.SimpleSlice {
	testCases := it.baseGenerator.Cases()

	for i, testCase := range testCases {
		replacerMap := map[string]string{
			unitTestVars.Title:         "",
			unitTestVars.ArrangeType:   "",
			unitTestVars.ArrangeSetup:  "",
			unitTestVars.ExpectedLines: "",
		}
	}
}
