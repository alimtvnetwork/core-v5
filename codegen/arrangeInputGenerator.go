package codegen

import (
	"fmt"
	"log/slog"
	"reflect"
	"strings"

	"github.com/alimtvnetwork/core/codegen/coreproperty"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coreindexes"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
	"github.com/alimtvnetwork/core/isany"
)

type arrangeInputGenerator struct {
	baseGenerator BaseGenerator
}

func (it arrangeInputGenerator) Generate(
	// isSubRequest bool,
	arrangeInput any,
) (string, error) {
	slice := corestr.New.SimpleSlice.Cap(10)

	if isany.Null(arrangeInput) {
		return "nil", nil
	}

	switch casted := arrangeInput.(type) {
	case args.AsArgFuncContractsBinder:
		v := casted.AsArgFuncContractsBinder()
		argsCount := v.ArgsCount()

		for i := 0; i < argsCount; i++ {
			name := coreindexes.NameByIndex(i)

			slog.Debug("processing property", "name", name)
			slice.AppendFmt(
				argSingleTemplate,
				name,
				it.property(v, i),
			)
		}

		it.addExpect(slice, v)

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

		it.addExpect(slice, v)
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
	case map[string]any:
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
	case []any:
		for _, v := range casted {
			slice.AppendFmt(
				"%s,",
				it.writeTestCaseForProperty(v),
			)
		}
	case any:
		rt := reflect.TypeOf(arrangeInput)

		// array or slice
		if rt.Kind() == reflect.Array || rt.Kind() == reflect.Slice {
			return it.recursiveGenerateSlice(slice, arrangeInput)
		}

		if rt.Kind() == reflect.Interface {
			slice.AppendFmt(
				"%s,",
				it.writeTestCaseForProperty(casted),
			)
		}
	}

	rt := reflect.TypeOf(arrangeInput)

	// array or slice
	if rt.Kind() == reflect.Array || rt.Kind() == reflect.Slice {
		return it.recursiveGenerateSlice(slice, arrangeInput)
	}

	if slice.IsEmpty() {
		return "", fmt.Errorf(
			"test cases only support from arg.One ... arg.Six and func versions (+ %s), given %T",
			"[]string, map[string]string, []any",
			arrangeInput,
		)
	}

	return slice.Join(linerJoiner), nil
}

func (it arrangeInputGenerator) addExpect(
	slice *corestr.SimpleSlice,
	v args.ArgBaseContractsBinder,
) *corestr.SimpleSlice {
	if !v.HasExpect() {
		return slice
	}

	return slice.AppendFmt(
		argSingleTemplate,
		vars.expect,
		it.writeTestCaseForProperty(v.Expected()),
	)
}

func (it arrangeInputGenerator) recursiveGenerateSlice(
	slice *corestr.SimpleSlice,
	arrangeInput any,
) (string, error) {
	trimmedTemplate := strings.TrimSpace(curlyOutputTemplate)

	compiledErr := reflectinternal.Looper.Slice(
		arrangeInput,
		func(total int, index int, item any) (err error) {
			expand, expandError := it.Generate(item)

			slice.AppendFmtIf(
				expandError == nil,
				trimmedTemplate,
				expand,
			)

			return expandError
		},
	)

	toCompiled := slice.Join(",\n")

	return toCompiled, compiledErr
}

func (it arrangeInputGenerator) property(
	argBinder args.ArgBaseContractsBinder,
	i int,
) string {
	p := argBinder.GetByIndex(i)

	return coreproperty.Writer.Write(p)
}

func (it arrangeInputGenerator) writeTestCaseForProperty(p any) string {
	return coreproperty.Writer.Write(p)
}

func (it arrangeInputGenerator) ReplaceTemplate(
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
