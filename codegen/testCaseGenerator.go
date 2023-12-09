package codegen

import (
	"fmt"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coreindexes"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/iserror"
)

type testCaseGenerator struct {
	baseGenerator BaseGenerator
}

func (it testCaseGenerator) Compile() (string, error) {
	it.caseItems()

	return "", nil
}

func (it testCaseGenerator) FuncWrap() *args.FuncWrap {
	return it.baseGenerator.FuncWrap()
}

func (it testCaseGenerator) caseItems() *corestr.SimpleSlice {
	testCases := it.baseGenerator.Cases()

	for i, testCase := range testCases {
		it.SingleArrange(i, testCase)
	}
}

func (it testCaseGenerator) SingleArrange(
	_ int,
	caseV1 coretestcases.CaseV1,
) (string, error) {
	arrangeSetup, err := it.arrangeSetup(caseV1)

	if iserror.Defined(err) {
		return "", err
	}

	it.expectedLines(caseV1)

	replacerMap := map[string]string{
		vars.Title:         caseV1.Title,
		vars.ArrangeType:   caseV1.ArrangeTypeName(),
		vars.ArrangeSetup:  arrangeSetup,
		vars.ExpectedLines: "",
	}

}

func (it testCaseGenerator) expectedLines(caseV1 coretestcases.CaseV1) {
	casted, isOkay := caseV1.ArrangeInput.(args.ArgBaseContractsBinder)

	if isOkay {
		args := casted.ValidArgs()
		results, err := it.FuncWrap().Invoke(args...)
	}
}

func (it testCaseGenerator) arrangeSetup(caseV1 coretestcases.CaseV1) (string, error) {
	slice := corestr.New.SimpleSlice.Cap(10)

	switch v := caseV1.ArrangeInput.(type) {
	case args.ArgFuncContractsBinder:
		argsCount := v.ArgsCount()

		for i := 0; i < argsCount; i++ {
			name := coreindexes.NameByIndex(i)

			slice.AppendFmt(
				argSingleTemplate,
				name,
				v.GetByIndex(i),
			)
		}

		slice.AppendFmtIf(
			v.HasExpect(),
			argSingleTemplate,
			vars.expect,
			v.Expected(),
		)

		slice.AppendFmtIf(
			v.HasFunc(),
			argSingleTemplate,
			vars.workFunc,
			v.GetFuncName(),
		)
	case args.ArgBaseContractsBinder:
		argsCount := v.ArgsCount()

		for i := 0; i < argsCount; i++ {
			name := coreindexes.NameByIndex(i)

			slice.AppendFmt(
				argSingleTemplate,
				name,
				v.GetByIndex(i),
			)
		}

		slice.AppendFmtIf(
			v.HasExpect(),
			argSingleTemplate,
			vars.expect,
			v.Expected(),
		)
	default:
		return "", fmt.Errorf(
			"test cases only support from arg.One ... arg.Six and func versions, given %T",
			v,
		)

	}

	return slice.JoinCsvLine(), nil
}
