package codegen

import (
	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/convertinteranl"
	"gitlab.com/auk-go/core/internal/reflectinternal"
	"gitlab.com/auk-go/core/isany"
	"gitlab.com/auk-go/core/iserror"
)

type expectedLinesGenerator struct {
	baseGenerator BaseGenerator
}

func (it expectedLinesGenerator) FuncWrap() *args.FuncWrap {
	return it.baseGenerator.FuncWrap()
}

func (it expectedLinesGenerator) FuncName() string {
	return it.baseGenerator.FuncName()
}

func (it expectedLinesGenerator) Generate() *corestr.SimpleSlice {

}

func (it expectedLinesGenerator) expectedLinesUsingArrange(
	slice *corestr.SimpleSlice,
	arrangeInput interface{},
) (
	*corestr.SimpleSlice, error,
) {
	if isany.Null(arrangeInput) {
		return slice.Add("nil"), nil
	}

	funcWrap := it.FuncWrap()
	var rawErrCollection errcore.RawErrCollection

	switch casted := arrangeInput.(type) {
	case coreinterface.ValidArgsGetter:
		validArgs := casted.ValidArgs()
		results, err := funcWrap.InvokeSkip(
			codestack.Skip1,
			validArgs...,
		)

		if iserror.Defined(err) {
			return it.enhanceErrorExpectedLines(err)
		}

		return it.expectedLinesAppendToSlice(
			slice,
			validArgs,
			results,
		), nil
	case string, map[string]string, map[string]interface{}:
		results, err := funcWrap.InvokeSkip(
			codestack.Skip1,
			casted,
		)

		if iserror.Defined(err) {
			return it.enhanceErrorExpectedLines(err)
		}

		return it.expectedLinesAppendSingleInToSlice(
			slice,
			casted,
			results,
		), nil
	case []string:
		if funcWrap.IsInTypeMatches(casted) {
			results, err := funcWrap.InvokeSkip(
				codestack.Skip1,
				casted,
			)

			if iserror.Defined(err) {
				return it.enhanceErrorExpectedLines(err)
			}

			return it.expectedLinesAppendSingleInToSlice(
				slice,
				casted,
				results,
			), nil
		}

		for i, itemString := range casted {
			results, err := funcWrap.InvokeSkip(
				codestack.Skip1,
				casted,
			)

			it.expectedLinesAppendSingleInToSlice(
				slice,
				itemString,
				results,
			)
		}

	case []interface{}:
		isInterfaceTypeExpected, _ := funcWrap.VerifyInArgs(casted)

		if isInterfaceTypeExpected {
			results, err := funcWrap.InvokeSkip(
				codestack.Skip1,
				casted,
			)

			if iserror.Defined(err) {
				return it.enhanceErrorExpectedLines(err)
			}

			return it.expectedLinesAppendToSlice(
				slice,
				casted,
				results,
			), nil
		}

		for i, item := range casted {
			_, err := it.expectedLinesUsingArrange(
				slice,
				item,
			)

		}
	case interface{}:
		slice.AppendFmt(
			"%s,",
			it.writeTestCaseForProperty(casted),
		)
	}

	return slice, nil
}

func (it expectedLinesGenerator) enhanceErrorExpectedLines(err error) (*corestr.SimpleSlice, error) {
	return nil, errcore.
		ConcatMessageWithErrWithStackTrace(
			"provide args properly in the definition of Generate (to get run the func and get the expected Lines),\n",
			err,
		)
}

func (it expectedLinesGenerator) expectedLinesAppendToSlice(
	slice *corestr.SimpleSlice,
	inArgs []interface{},
	outArgs []interface{},
) *corestr.SimpleSlice {
	inArgsString := convertinteranl.AnyTo.String(inArgs)
	resultsToString := convertinteranl.AnyTo.String(outArgs)

	slice.AppendFmt(
		it.baseGenerator.FmtJoin(),
		0,
		inArgsString,
		resultsToString,
	)

	return slice
}

func (it expectedLinesGenerator) expectedLinesAppendSingleInToSlice(
	slice *corestr.SimpleSlice,
	inArgs interface{},
	outArgs []interface{},
) *corestr.SimpleSlice {
	inArgsString := convertinteranl.AnyTo.String(inArgs)
	resultsToString := convertinteranl.AnyTo.String(outArgs)

	slice.AppendFmt(
		it.baseGenerator.FmtJoin(),
		0,
		inArgsString,
		resultsToString,
	)

	return slice
}

func (it expectedLinesGenerator) expectedLinesForOther(
	caseV1 coretestcases.CaseV1,
) (*corestr.SimpleSlice, error) {

	return nil, errcore.Expected.But(
		"cannot cast caseV1.ArrangeInput to args.AsArgBaseContractsBinder",
		reflectinternal.TypeName(x),
		reflectinternal.TypeName(arrange),
	)
}
