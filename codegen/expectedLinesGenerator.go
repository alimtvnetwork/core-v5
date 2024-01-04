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

func (it expectedLinesGenerator) Generate() (*corestr.SimpleSlice, error) {
	slice := corestr.New.SimpleSlice.Cap(10)

	_, err := it.expectedLinesUsingArrange(
		slice,
		it.caseV1.ArrangeInput,
	)

	return slice, err
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
		return it.handleForArrayOrSliceArrange(slice, arrangeInput)
	}

	if slice.IsEmpty() {
		return slice, fmt.Errorf(
			"test cases only support from arg.One ... arg.Six and func versions (+ %s), given %T",
			"[]string, map[string]string, []interface{}",
			arrangeInput,
		)
	}

	return slice, nil
}

func (it expectedLinesGenerator) handleForArrayOrSliceArrange(
	slice *corestr.SimpleSlice,
	arrangeInput interface{},
) (*corestr.SimpleSlice, error) {
	funcWrap := it.FuncWrap()
	var rawErrCollection errcore.RawErrCollection

	if funcWrap.IsInTypeMatches(arrangeInput) {
		results, err := funcWrap.InvokeSkip(
			codestack.Skip1,
			arrangeInput,
		)

		if iserror.Defined(err) {
			return it.enhanceError(err)
		}

		return it.appendSingleInToSlice(
			slice,
			arrangeInput,
			results,
		), nil
	}

	_ = reflectinternal.Looper.Slice(
		arrangeInput,
		func(total int, index int, item interface{}) (err error) {
			_, expandError := it.expectedLinesUsingArrange(slice, item)

			rawErrCollection.AddFmt(
				expandError,
				"At: %d, item: %+v",
				index,
				item,
			)

			return nil
		},
	)

	return slice, rawErrCollection.CompiledError()
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

//
// func (it expectedLinesGenerator) ReplaceTemplate(
// 	format string,
// 	replacerMap map[string]string,
// ) string {
// 	if len(format) == 0 {
// 		return ""
// 	}
//
// 	return templateReplacerFunc(
// 		format,
// 		replacerMap,
// 	)
// }
