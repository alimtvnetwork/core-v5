package codegen

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

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
		it.SingleArrange(i, testCase)
	}
}

func (it testCaseGenerator) SingleArrange(
	index int,
	caseV1 coretestcases.CaseV1,
) string {
	replacerMap := map[string]string{
		vars.Title:         "",
		vars.ArrangeType:   "",
		vars.ArrangeSetup:  "",
		vars.ExpectedLines: "",
	}

	it.arrangeSetup(caseV1)
}

func (it testCaseGenerator) arrangeSetup(caseV1 coretestcases.CaseV1) string {
	switch v := caseV1.ArrangeInput.(type) {
	case args.One:

	}
}
