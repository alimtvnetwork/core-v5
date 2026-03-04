package args

import (
	"reflect"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func (it *FuncWrap) ArgsCount() int {
	if it.IsInvalid() {
		return -1
	}

	return it.rvType.NumIn()
}

func (it *FuncWrap) InArgsCount() int {
	return it.ArgsCount()
}

func (it *FuncWrap) OutArgsCount() int {
	if it.IsInvalid() {
		return -1
	}

	return it.rvType.NumOut()
}

func (it *FuncWrap) ArgsLength() int {
	return it.ArgsCount()
}

func (it *FuncWrap) ReturnLength() int {
	if it.IsInvalid() {
		return -1
	}

	return it.rvType.NumOut()
}

func (it *FuncWrap) GetOutArgsTypes() []reflect.Type {
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

func (it *FuncWrap) GetInArgsTypes() []reflect.Type {
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

func (it *FuncWrap) InArgNames() []string {
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
			toSlice.AppendFmt("%s%s%d", inArgNamePrefix, cTypeNamePascal, i+1)
		}
	}

	it.inArgsNames = toSlice.Strings()

	return it.inArgsNames
}

func (it *FuncWrap) InArgNamesEachLine() corestr.SimpleSlice {
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

func (it *FuncWrap) OutArgNamesEachLine() corestr.SimpleSlice {
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

func (it *FuncWrap) OutArgNames() []string {
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
			toSlice.AppendFmt("%s%s%d", outArgNamePrefix, cTypeNamePascal, i)
		}
	}

	it.outArgsNames = toSlice.Strings()

	return it.outArgsNames
}

func (it *FuncWrap) GetInArgsTypesNames() []string {
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

func (it *FuncWrap) GetOutArgsTypesNames() []string {
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

func (it *FuncWrap) IsInTypeMatches(args ...any) (isOkay bool) {
	toTypes := reflectinternal.Converter.InterfacesToTypes(args)
	isOkay, _ = it.InArgsVerifyRv(toTypes)

	return isOkay
}

func (it *FuncWrap) IsOutTypeMatches(outArgs ...any) (isOkay bool) {
	toTypes := reflectinternal.Converter.InterfacesToTypes(outArgs)
	isOkay, _ = it.OutArgsVerifyRv(toTypes)

	return isOkay
}

func (it *FuncWrap) VerifyInArgs(args []any) (isOkay bool, err error) {
	toTypes := reflectinternal.Converter.InterfacesToTypes(args)

	return it.InArgsVerifyRv(toTypes)
}

func (it *FuncWrap) VerifyOutArgs(args []any) (isOkay bool, err error) {
	toTypes := reflectinternal.Converter.InterfacesToTypes(args)

	return it.OutArgsVerifyRv(toTypes)
}

func (it *FuncWrap) InArgsVerifyRv(args []reflect.Type) (isOkay bool, err error) {
	return reflectinternal.Utils.VerifyReflectTypes(
		it.Name,
		it.GetInArgsTypes(),
		args,
	)
}

func (it *FuncWrap) OutArgsVerifyRv(args []reflect.Type) (isOkay bool, err error) {
	return reflectinternal.Utils.VerifyReflectTypes(
		it.Name,
		it.GetOutArgsTypes(),
		args,
	)
}
