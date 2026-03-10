package args

import (
	"reflect"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// ArgsCount returns the number of input arguments the wrapped function expects.
// Returns -1 if the FuncWrap is invalid.
func (it *FuncWrap[T]) ArgsCount() int {
	if it.IsInvalid() {
		return -1
	}

	return it.rvType.NumIn()
}

// InArgsCount is an alias for ArgsCount.
func (it *FuncWrap[T]) InArgsCount() int {
	return it.ArgsCount()
}

// OutArgsCount returns the number of return values from the wrapped function.
// Returns -1 if the FuncWrap is invalid.
func (it *FuncWrap[T]) OutArgsCount() int {
	if it.IsInvalid() {
		return -1
	}

	return it.rvType.NumOut()
}

// ArgsLength is an alias for ArgsCount.
func (it *FuncWrap[T]) ArgsLength() int {
	return it.ArgsCount()
}

// ReturnLength returns the number of return values.
// Returns -1 if the FuncWrap is invalid.
func (it *FuncWrap[T]) ReturnLength() int {
	if it.IsInvalid() {
		return -1
	}

	return it.rvType.NumOut()
}

// GetOutArgsTypes returns the reflect.Type for each return value.
func (it *FuncWrap[T]) GetOutArgsTypes() []reflect.Type {
	if it.IsInvalid() {
		return []reflect.Type{}
	}

	argsOutCount := it.ReturnLength()

	if argsOutCount == 0 {
		return []reflect.Type{}
	}

	if len(it.outArgsTypes) == argsOutCount {
		return it.outArgsTypes
	}

	mainType := it.rvType
	slice := make([]reflect.Type, 0, argsOutCount)

	for i := 0; i < argsOutCount; i++ {
		slice = append(slice, mainType.Out(i))
	}

	it.outArgsTypes = slice

	return slice
}

// GetInArgsTypes returns the reflect.Type for each input argument.
func (it *FuncWrap[T]) GetInArgsTypes() []reflect.Type {
	if it.IsInvalid() {
		return []reflect.Type{}
	}

	argsCount := it.ArgsCount()

	if argsCount == 0 {
		return []reflect.Type{}
	}

	if len(it.inArgsTypes) == argsCount {
		return it.inArgsTypes
	}

	mainType := it.rvType
	slice := make([]reflect.Type, 0, argsCount)

	for i := 0; i < argsCount; i++ {
		slice = append(slice, mainType.In(i))
	}

	it.inArgsTypes = slice

	return slice
}

// InArgNames returns generated variable names for each input argument,
// derived from their type names.
func (it *FuncWrap[T]) InArgNames() []string {
	if it.InArgsCount() <= 0 {
		return []string{}
	}

	count := it.InArgsCount()

	if len(it.inArgsNames) == count {
		return it.inArgsNames
	}

	allTypesNames := it.GetInArgsTypesNames()
	toSlice := corestr.New.SimpleSlice.ByLen(allTypesNames)
	convertFunc := reflectinternal.TypeNameToValidVariableName

	switch count {
	case 1:
		firstType := pascalCaseFunc(allTypesNames[0])
		toSlice.Add(inArgNamePrefix + convertFunc(firstType))
	default:
		for i, cTypeName := range allTypesNames {
			cTypeNamePascal := pascalCaseFunc(convertFunc(cTypeName))
			toSlice.AppendFmt(
				"%s%s%d",
				inArgNamePrefix,
				cTypeNamePascal,
				i+1,
			)
		}
	}

	it.inArgsNames = toSlice.Strings()

	return it.inArgsNames
}

// InArgNamesEachLine returns the input argument names formatted one per line.
func (it *FuncWrap[T]) InArgNamesEachLine() corestr.SimpleSlice {
	inArgs := it.InArgNames()

	if len(inArgs) <= 1 {
		return inArgs
	}

	toSlice := corestr.New.SimpleSlice.Cap(len(inArgs) + 2)
	toSlice.Add("\n")

	for _, arg := range inArgs {
		toSlice.Add(arg + "\n")
	}

	return toSlice.Strings()
}

// OutArgNamesEachLine returns the output argument names formatted one per line.
func (it *FuncWrap[T]) OutArgNamesEachLine() corestr.SimpleSlice {
	outArgs := it.OutArgNames()

	if len(outArgs) <= 1 {
		return outArgs
	}

	toSlice := corestr.New.SimpleSlice.Cap(len(outArgs) + 2)
	toSlice.Add("\n")

	for _, arg := range outArgs {
		toSlice.Add(arg + "\n")
	}

	return toSlice.Strings()
}

// OutArgNames returns generated variable names for each return value,
// derived from their type names.
func (it *FuncWrap[T]) OutArgNames() []string {
	if it.OutArgsCount() <= 0 {
		return []string{}
	}

	count := it.OutArgsCount()

	if len(it.outArgsNames) == count {
		return it.outArgsNames
	}

	allTypesNames := it.GetOutArgsTypesNames()
	toSlice := corestr.New.SimpleSlice.ByLen(allTypesNames)

	switch count {
	case 1:
		firstType := pascalCaseFunc(allTypesNames[0])
		toSlice.Add(outArgNamePrefix + firstType)
	default:
		for i, cTypeName := range allTypesNames {
			cTypeNamePascal := pascalCaseFunc(cTypeName)
			toSlice.AppendFmt(
				"%s%s%d",
				outArgNamePrefix,
				cTypeNamePascal,
				i,
			)
		}
	}

	it.outArgsNames = toSlice.Strings()

	return it.outArgsNames
}

// GetInArgsTypesNames returns the string representation of each input argument type.
func (it *FuncWrap[T]) GetInArgsTypesNames() []string {
	if it.IsInvalid() {
		return []string{}
	}

	argsCount := it.ArgsCount()

	if argsCount == 0 {
		return []string{}
	}

	if len(it.inArgsTypesNames) == argsCount {
		return it.inArgsTypesNames
	}

	mainType := it.rvType
	slice := make([]string, 0, argsCount)

	for i := 0; i < argsCount; i++ {
		slice = append(slice, mainType.In(i).String())
	}

	it.inArgsTypesNames = slice

	return slice
}

// GetOutArgsTypesNames returns the string representation of each return value type.
func (it *FuncWrap[T]) GetOutArgsTypesNames() []string {
	if it.IsInvalid() {
		return []string{}
	}

	argsCount := it.OutArgsCount()

	if argsCount == 0 {
		return []string{}
	}

	if len(it.outArgsTypesNames) == argsCount {
		return it.outArgsTypesNames
	}

	mainType := it.rvType
	slice := make([]string, 0, argsCount)

	for i := 0; i < argsCount; i++ {
		slice = append(slice, mainType.Out(i).String())
	}

	it.outArgsTypesNames = slice

	return slice
}

// IsInTypeMatches checks whether the given arguments match the input types
// of the wrapped function.
func (it *FuncWrap[T]) IsInTypeMatches(args ...any) (isOkay bool) {
	toTypes := reflectinternal.Converter.InterfacesToTypes(args)
	isOkay, _ = it.InArgsVerifyRv(toTypes)

	return isOkay
}

// IsOutTypeMatches checks whether the given arguments match the output types
// of the wrapped function.
func (it *FuncWrap[T]) IsOutTypeMatches(outArgs ...any) (isOkay bool) {
	toTypes := reflectinternal.Converter.InterfacesToTypes(outArgs)
	isOkay, _ = it.OutArgsVerifyRv(toTypes)

	return isOkay
}

// VerifyInArgs verifies that the given arguments match the input types.
func (it *FuncWrap[T]) VerifyInArgs(args []any) (isOkay bool, err error) {
	toTypes := reflectinternal.Converter.InterfacesToTypes(args)

	return it.InArgsVerifyRv(toTypes)
}

// VerifyOutArgs verifies that the given arguments match the output types.
func (it *FuncWrap[T]) VerifyOutArgs(args []any) (isOkay bool, err error) {
	toTypes := reflectinternal.Converter.InterfacesToTypes(args)

	return it.OutArgsVerifyRv(toTypes)
}

// InArgsVerifyRv verifies input argument types using reflect.Type slices.
func (it *FuncWrap[T]) InArgsVerifyRv(args []reflect.Type) (isOkay bool, err error) {
	return reflectinternal.Utils.VerifyReflectTypes(
		it.Name,
		it.GetInArgsTypes(),
		args,
	)
}

// OutArgsVerifyRv verifies output argument types using reflect.Type slices.
func (it *FuncWrap[T]) OutArgsVerifyRv(args []reflect.Type) (isOkay bool, err error) {
	return reflectinternal.Utils.VerifyReflectTypes(
		it.Name,
		it.GetOutArgsTypes(),
		args,
	)
}
