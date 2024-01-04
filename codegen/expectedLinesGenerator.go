package codegen

import (
	"fmt"
	"reflect"

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
	caseV1        coretestcases.CaseV1
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
) (*corestr.SimpleSlice, error) {
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
			return it.enhanceError(err)
		}

		return it.appendToSlice(
			slice,
			validArgs,
			results,
		), nil
	case string, map[string]string, map[string]interface{}:
		// TODO for the Map
		results, err := funcWrap.InvokeSkip(
			codestack.Skip1,
			casted,
		)

		if iserror.Defined(err) {
			return it.enhanceError(err)
		}

		return it.appendSingleInToSlice(
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
				return it.enhanceError(err)
			}

			return it.appendSingleInToSlice(
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

			rawErrCollection.AddFmt(
				err,
				"At: %d, item: %s",
				i,
				itemString,
			)

			it.appendSingleInToSlice(
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
				return it.enhanceError(err)
			}

			return it.appendToSlice(
				slice,
				casted,
				results,
			), nil
		}

		for i, item := range casted {
			// add to slice if matches
			_, err := it.expectedLinesUsingArrange(
				slice,
				item,
			)

			rawErrCollection.AddFmt(
				err,
				"At: %d, item: %+v",
				i,
				item,
			)
		}
	case interface{}:
		results, err := funcWrap.InvokeSkip(
			codestack.Skip1,
			casted,
		)

		if iserror.Defined(err) {
			return it.enhanceError(err)
		}

		return it.appendSingleInToSlice(
			slice,
			casted,
			results,
		), nil
	}

	rt := reflect.TypeOf(arrangeInput)

	// array or slice
	if rt.Kind() == reflect.Array || rt.Kind() == reflect.Slice {
		return it.handleForArrayOrSliceArrange(arrangeInput, slice)
	}

	if slice.IsEmpty() {
		return "", fmt.Errorf(
			"test cases only support from arg.One ... arg.Six and func versions (+ %s), given %T",
			"[]string, map[string]string, []interface{}",
			arrangeInput,
		)
	}

	return slice, nil
}

func (it testCaseGenerator) handleForArrayOrSliceArrange(
	arrangeInput interface{},
	slice *corestr.SimpleSlice,
) (string, error) {
	compiledErr := reflectinternal.Looper.Slice(
		arrangeInput,
		func(total int, index int, item interface{}) (err error) {
			expand, expandError := it.testCaseArrangeInputWrite(item)

			if expandError != nil {
				return expandError
			}

			slice.Append(
				expand,
			)

			return
		},
	)

	toCompiled := slice.Join(linerJoiner)
	typeName := reflectinternal.ReflectType.NameUsingFmt(arrangeInput)
	replacerMap := map[string]string{
		vars.TypeName:   toCompiled,
		vars.ToCompiled: typeName,
	}

	finalOutput := it.ReplaceTemplate(
		typeWithCompiledItemsTemplate,
		replacerMap,
	)

	return finalOutput, compiledErr
}

func (it expectedLinesGenerator) enhanceError(err error) (*corestr.SimpleSlice, error) {
	return nil, errcore.
		ConcatMessageWithErrWithStackTrace(
			"provide args properly in the definition of Generate (to get run the func and get the expected Lines),\n",
			err,
		)
}

func (it expectedLinesGenerator) appendToSlice(
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

func (it expectedLinesGenerator) appendSingleInToSlice(
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
