package codegen

import (
	"fmt"

	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coreindexes"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/convertinteranl"
	"gitlab.com/auk-go/core/internal/reflectinternal"
	"gitlab.com/auk-go/core/isany"
	"gitlab.com/auk-go/core/iserror"
	"gitlab.com/auk-go/core/simplewrap"
)

type testCaseGenerator struct {
	baseGenerator BaseGenerator
}

func (it testCaseGenerator) GetBehaviours() corestr.SimpleSlice {
	return it.baseGenerator.GetBehaviours()
}

func (it testCaseGenerator) NewCode(codes ...string) *GoCode {
	return it.baseGenerator.NewGoCode(codes...)
}

func (it testCaseGenerator) Compile() (*GoCode, error) {
	behaviours := it.GetBehaviours()
	totalBehaviours := len(behaviours)
	testCasesSlice := corestr.New.SimpleSlice.Cap(totalBehaviours)

	for _, behaviour := range behaviours {
		caseOutput, err := it.fullTestCase(
			totalBehaviours, behaviour,
		)

		if iserror.Defined(err) {
			return it.NewCode(), err
		}

		testCasesSlice.Add(caseOutput)
	}

	if testCasesSlice.Length() == 0 {
		return it.NewCode(), errcore.InvalidEmptyValueType.Error(
			"no testcases generated for the behaviour",
			behaviours,
		)
	}

	allCompiledTestCases := testCasesSlice.Join(
		"\n\n\t",
	)

	replacerMap := map[string]string{
		vars.TestCases: allCompiledTestCases,
	}

	caseOutputWithVar := it.ReplaceTemplate(
		fullTestCaseFileTemplate,
		replacerMap,
	)

	final := it.baseGenerator.NewGoCode(caseOutputWithVar)

	return final, nil
}

func (it testCaseGenerator) fullTestCase(
	totalBehaviourCount int,
	behaviour string,
) (string, error) {
	allCases, err := it.caseItems()

	if iserror.Defined(err) {
		return "", errcore.
			ConcatMessageWithErrWithStackTrace(
				"failed for behaviour "+behaviour, err,
			)
	}

	replacerMap := map[string]string{
		vars.TestCaseName: it.testCaseName(totalBehaviourCount, behaviour),
		vars.CaseItem:     allCases.Join("\n\t\t"),
	}

	caseOutput := it.ReplaceTemplate(
		fullTestCaseTemplate,
		replacerMap,
	)

	return caseOutput, nil
}

func (it testCaseGenerator) FuncWrap() *args.FuncWrap {
	return it.baseGenerator.FuncWrap()
}

func (it testCaseGenerator) FuncName() string {
	return it.baseGenerator.FuncName()
}

func (it testCaseGenerator) testCaseName(
	totalBehaviour int,
	behaviourName string,
) string {
	return it.baseGenerator.TestCaseName(
		totalBehaviour,
		it.FuncName(),
		behaviourName,
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

func (it testCaseGenerator) ReplaceTemplate(
	format string,
	replacerMap map[string]string,
) string {
	if len(format) == 0 {
		return ""
	}

	return templateReplacerFunc(
		format,
		replacerMap,
	)
}

func (it testCaseGenerator) SingleArrange(
	_ int,
	caseV1 coretestcases.CaseV1,
) (string, error) {
	testCaseArrangeInput, err := it.testCaseArrangeInputWrite(caseV1)

	if iserror.Defined(err) {
		return "", err
	}

	expectedLines, expectedLinesErr := it.expectedLines(caseV1)

	if iserror.Defined(expectedLinesErr) {
		return "", expectedLinesErr
	}

	replacerMap := map[string]string{
		vars.Title:         simplewrap.WithDoubleQuote(caseV1.Title),
		vars.ArrangeType:   caseV1.ArrangeTypeName(),
		vars.VerifyTypeOf:  it.VerifyTypeOf(),
		vars.ArrangeSetup:  testCaseArrangeInput,
		vars.ExpectedLines: expectedLines.WrapDoubleQuote().Join(",\n\t\t\t\t"),
	}

	caseOutput := it.ReplaceTemplate(
		testCaseItemTemplate,
		replacerMap,
	)

	return caseOutput, nil
}

func (it testCaseGenerator) expectedLines(caseV1 coretestcases.CaseV1) (*corestr.SimpleSlice, error) {
	var x args.AsArgBaseContractsBinder

	arrange := caseV1.ArrangeInput
	casted, isOkay := arrange.(args.AsArgBaseContractsBinder)

	if !isOkay {
		return nil, errcore.Expected.But(
			"cannot cast caseV1.ArrangeInput to args.AsArgBaseContractsBinder",
			reflectinternal.TypeName(x),
			reflectinternal.TypeName(arrange),
		)
	}

	validArgs := casted.
		AsArgBaseContractsBinder().
		ValidArgs()
	results, err := it.
		FuncWrap().
		InvokeSkip(
			codestack.Skip1,
			validArgs...,
		)

	if iserror.Defined(err) {
		return nil, errcore.
			ConcatMessageWithErr(
				"provide args properly in the definition of Generate,\n", err,
			)
	}

	slice := corestr.New.SimpleSlice.Cap(2)
	inArgsString := convertinteranl.AnyTo.String(validArgs)
	resultsToString := convertinteranl.AnyTo.String(results)

	slice.AppendFmt(
		it.baseGenerator.FmtJoin(),
		0,
		inArgsString,
		resultsToString,
	)

	return slice, nil
}

func (it testCaseGenerator) testCaseArrangeInputWrite(caseV1 coretestcases.CaseV1) (string, error) {
	slice := corestr.New.SimpleSlice.Cap(10)

	if isany.Null(caseV1.ArrangeInput) {
		return "", nil
	}

	switch casted := caseV1.ArrangeInput.(type) {
	case args.AsArgFuncContractsBinder:
		v := casted.AsArgFuncContractsBinder()
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
	case args.AsArgBaseContractsBinder:
		v := casted.AsArgBaseContractsBinder()
		argsCount := v.ArgsCount()

		for i := 0; i < argsCount; i++ {
			name := coreindexes.NameByIndex(i)

			slice.AppendFmt(
				argSingleTemplate,
				name,
				it.property(v, i),
			)
		}

		slice.AppendFmtIf(
			v.HasExpect(),
			argSingleTemplate,
			vars.expect,
			v.Expected(),
		)
	case string:
		slice.AppendFmt(
			"\"%s\",",
			casted,
		)
	case args.String:
		slice.AppendFmt(
			"%s,",
			casted,
		)
	case []string:
		for _, item := range casted {
			slice.AppendFmt(
				"\"%s\",",
				item,
			)
		}
	case map[string]string:
		for k, v := range casted {
			slice.AppendFmt(
				"\"%s\" : \"%s\",",
				k,
				v,
			)
		}
	case map[string]interface{}:
		for k, v := range casted {
			slice.AppendFmt(
				"\"%s\" : %s,",
				k,
				it.writeTestCaseForProperty(v),
			)
		}
	case args.Map:
		for k, v := range casted {
			slice.AppendFmt(
				"\"%s\" : %s,",
				k,
				it.writeTestCaseForProperty(v),
			)
		}
	case []args.One, []args.OneFunc, []args.TwoFunc:

	case []interface{}:
		for _, v := range casted {
			slice.AppendFmt(
				"%s,",
				it.writeTestCaseForProperty(v),
			)
		}
	case interface{}:
		slice.AppendFmt(
			"%s,",
			it.writeTestCaseForProperty(casted),
		)
	}

	return "", fmt.Errorf(
		"test cases only support from arg.One ... arg.Six and func versions (+ %s), given %T",
		"[]string, map[string]string, []interface{}",
		caseV1.ArrangeInput,
	)

	return slice.Join(",\n\t\t\t\t"), nil
}

func (it testCaseGenerator) property(argBinder args.ArgBaseContractsBinder, i int) string {
	p := argBinder.GetByIndex(i)

	return it.writeTestCaseForProperty(p)
}

func (it testCaseGenerator) writeTestCaseForProperty(p interface{}) string {
	switch casted := p.(type) {
	case string:
		return simplewrap.WithDoubleQuote(casted)
	case bool, int, int32, int64,
		float64, float32, byte,
		int8, uint16, uint32,
		uint64:
		return fmt.Sprintf("%d", casted)
	case args.String:
		return fmt.Sprintf("%s", casted)
	}

	return convertinteranl.AnyTo.FullPropertyString(p)
}

func (it testCaseGenerator) VerifyTypeOf() string {
	caseV1 := it.baseGenerator.FirstTestCase()

	if caseV1 == nil {
		return "nil"
	}

	switch casted := caseV1.ActualInput.(type) {
	case string:
		return simplewrap.WithDoubleQuote("")
	case args.String:
		return casted.String() // be direct
	case bool, int, int32, int64,
		float64, float32, byte,
		int8, uint16, uint32,
		uint64:
		return convertinteranl.AnyTo.ValueString(casted)
	}

	return fmt.Sprintf(
		"%T{}",
		caseV1.ArrangeInput,
	)
}
