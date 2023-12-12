package codegen

import (
	"errors"
	"fmt"

	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coreindexes"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/coreutils/stringutil"
	"gitlab.com/auk-go/core/internal/convertinteranl"
	"gitlab.com/auk-go/core/internal/reflectinternal"
	"gitlab.com/auk-go/core/iserror"
)

type testCaseGenerator struct {
	baseGenerator BaseGenerator
}

func (it testCaseGenerator) Compile() (string, error) {
	allCases, err := it.caseItems()

	if iserror.Defined(err) {
		return "", err
	}

	replacerMap := map[string]string{
		vars.TestCaseName:  caseV1.Title,
		vars.ArrangeType:   caseV1.ArrangeTypeName(),
		vars.ArrangeSetup:  arrangeSetup,
		vars.ExpectedLines: expectedLines.WrapDoubleQuote().Join(",\n\t\t\t\t"),
	}

	caseOutput := stringutil.
		ReplaceTemplate.
		DirectKeyUsingMapTrim(
			testCaseItemTemplate,
			replacerMap,
		)

	return "", nil
}

func (it testCaseGenerator) FuncWrap() *args.FuncWrap {
	return it.baseGenerator.FuncWrap()
}

func (it testCaseGenerator) FuncName() string {
	return it.baseGenerator.FuncName()
}

func (it testCaseGenerator) testCaseName() string {
	bG := it.baseGenerator
	behaviours := bG.CurBehaviours()
	total := len(behaviours)
	it.baseGenerator.TestCaseName(
		total,
		bG.FuncName(),
	)

	return it.baseGenerator.TestCaseName(
		len(it.),
	)
}

func (it testCaseGenerator) caseItems() (*corestr.SimpleSlice, error) {
	testCases := it.baseGenerator.Cases()
	slice := corestr.New.SimpleSlice.ByLen(testCases)

	for i, testCase := range testCases {
		caseOutput, err := it.SingleArrange(i, testCase)

		if iserror.Defined(err) {
			return nil, err
		}

		slice.Add(caseOutput)
	}

	return slice, nil
}

func (it testCaseGenerator) SingleArrange(
	_ int,
	caseV1 coretestcases.CaseV1,
) (string, error) {
	arrangeSetup, err := it.arrangeSetup(caseV1)

	if iserror.Defined(err) {
		return "", err
	}

	expectedLines, expectedLinesErr := it.expectedLines(caseV1)

	if iserror.Defined(expectedLinesErr) {
		return "", expectedLinesErr
	}

	replacerMap := map[string]string{
		vars.Title:         caseV1.Title,
		vars.ArrangeType:   caseV1.ArrangeTypeName(),
		vars.ArrangeSetup:  arrangeSetup,
		vars.ExpectedLines: expectedLines.WrapDoubleQuote().Join(",\n\t\t\t\t"),
	}

	caseOutput := stringutil.
		ReplaceTemplate.
		DirectKeyUsingMapTrim(
			testCaseItemTemplate,
			replacerMap,
		)

	return caseOutput, nil
}

func (it testCaseGenerator) expectedLines(caseV1 coretestcases.CaseV1) (*corestr.SimpleSlice, error) {
	arrange := caseV1.ArrangeInput
	casted, isOkay := reflectinternal.Converter.StructToMatchInterfaceDirect(
		arrange,
		(args.ArgBaseContractsBinder)(nil),
	).(args.ArgBaseContractsBinder)

	if !isOkay {
		return nil, errors.New("cannot cast caseV1.ArrangeInput to args.ArgBaseContractsBinder")
	}

	validArgs := casted.ValidArgs()
	results, err := it.
		FuncWrap().
		InvokeSkip(
		codestack.Skip1,
		validArgs...,
		)

	if iserror.Defined(err) {
		return nil, err
	}

	slice := corestr.New.SimpleSlice.Cap(2)

	return slice.Add(convertinteranl.AnyTo.SmartString(results)), nil
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
