package args

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/convertinternal"
	"gitlab.com/auk-go/core/internal/reflectinternal"
	"gitlab.com/auk-go/core/internal/trydo"
)

type FuncWrap struct {
	Name                 string         `json:",omitempty"`
	FullName             string         `json:",omitempty"`
	Func                 any            `json:"-"`
	isInvalid            bool           `json:"IsInvalid,omitempty"`
	rvType               reflect.Type   `json:"-"`
	rv                   reflect.Value  `json:"-"`
	inArgsTypesNames     []string       `json:"-"`
	inArgsTypes          []reflect.Type `json:"-"`
	outArgsTypes         []reflect.Type `json:"-"`
	pkgNameOnly          string
	funcDirectInvokeName string
	pkgPath              string
	inArgsMap            Map
	outArgsMap           Map
	inArgsNames          []string
	outArgsTypesNames    []string
	outArgsNames         []string
}

func (it *FuncWrap) GetFuncName() string {
	if it == nil { return "" }
	return it.Name
}

func (it *FuncWrap) GetPascalCaseFuncName() string {
	if it == nil { return "" }
	return pascalCaseFunc(it.Name)
}

func (it *FuncWrap) HasValidFunc() bool {
	return it != nil && !it.isInvalid && it.rv.IsValid() && reflectinternal.Is.Func(it.Func)
}

func (it *FuncWrap) IsInvalid() bool {
	return it == nil || it.isInvalid || !it.rv.IsValid() || !it.HasValidFunc()
}

func (it *FuncWrap) IsValid() bool { return !it.IsInvalid() }

func (it *FuncWrap) PkgPath() string {
	if it.IsInvalid() { return "" }
	if len(it.pkgPath) > 0 { return it.pkgPath }
	it.pkgPath = reflectinternal.GetFunc.GetPkgPathFullName(it.FullName)
	return it.pkgPath
}

func (it *FuncWrap) PkgNameOnly() string {
	if it.IsInvalid() { return "" }
	if len(it.pkgNameOnly) > 0 { return it.pkgNameOnly }
	it.pkgNameOnly = reflectinternal.Utils.PkgNameOnly(it.Func)
	return it.pkgNameOnly
}

func (it *FuncWrap) FuncDirectInvokeName() string {
	if it.IsInvalid() { return "" }
	if len(it.funcDirectInvokeName) > 0 { return it.funcDirectInvokeName }
	it.funcDirectInvokeName = reflectinternal.GetFunc.FuncDirectInvokeNameUsingFullName(it.FullName)
	return it.funcDirectInvokeName
}

func (it *FuncWrap) ArgsCount() int {
	if it.IsInvalid() { return -1 }
	return it.rvType.NumIn()
}

func (it *FuncWrap) InArgsCount() int { return it.ArgsCount() }

func (it *FuncWrap) OutArgsCount() int {
	if it.IsInvalid() { return -1 }
	return it.rvType.NumOut()
}

func (it *FuncWrap) ArgsLength() int { return it.ArgsCount() }

func (it *FuncWrap) ReturnLength() int {
	if it.IsInvalid() { return -1 }
	return it.rvType.NumOut()
}

func (it *FuncWrap) IsPublicMethod() bool  { return it != nil && it.rvType.PkgPath() == "" }
func (it *FuncWrap) IsPrivateMethod() bool { return it != nil && it.rvType.PkgPath() != "" }

func (it *FuncWrap) GetType() reflect.Type {
	if it.IsInvalid() { return nil }
	return it.rvType
}

func (it *FuncWrap) GetOutArgsTypes() []reflect.Type {
	if it.IsInvalid() { return []reflect.Type{} }
	argsOutCount := it.ReturnLength()
	if argsOutCount == 0 { return []reflect.Type{} }
	if len(it.outArgsTypes) == argsOutCount { return it.outArgsTypes }
	mainType := it.rvType
	slice := make([]reflect.Type, 0, argsOutCount)
	for i := 0; i < argsOutCount; i++ { slice = append(slice, mainType.Out(i)) }
	it.outArgsTypes = slice
	return slice
}

func (it *FuncWrap) GetInArgsTypes() []reflect.Type {
	if it.IsInvalid() { return []reflect.Type{} }
	argsCount := it.ArgsCount()
	if argsCount == 0 { return []reflect.Type{} }
	if len(it.inArgsTypes) == argsCount { return it.inArgsTypes }
	mainType := it.rvType
	slice := make([]reflect.Type, 0, argsCount)
	for i := 0; i < argsCount; i++ { slice = append(slice, mainType.In(i)) }
	it.inArgsTypes = slice
	return slice
}

func (it *FuncWrap) InArgNames() []string {
	if it.InArgsCount() <= 0 { return []string{} }
	count := it.InArgsCount()
	if len(it.inArgsNames) == count { return it.inArgsNames }
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
	if len(inArgs) <= 1 { return inArgs }
	toSlice := corestr.New.SimpleSlice.Cap(len(inArgs) + 2)
	toSlice.Add("\n")
	for _, arg := range inArgs { toSlice.Add(arg + "\n") }
	return toSlice.Strings()
}

func (it *FuncWrap) OutArgNamesEachLine() corestr.SimpleSlice {
	outArgs := it.OutArgNames()
	if len(outArgs) <= 1 { return outArgs }
	toSlice := corestr.New.SimpleSlice.Cap(len(outArgs) + 2)
	toSlice.Add("\n")
	for _, arg := range outArgs { toSlice.Add(arg + "\n") }
	return toSlice.Strings()
}

func (it *FuncWrap) OutArgNames() []string {
	if it.OutArgsCount() <= 0 { return []string{} }
	count := it.OutArgsCount()
	if len(it.outArgsNames) == count { return it.outArgsNames }
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
	if it.IsInvalid() { return []string{} }
	argsCount := it.ArgsCount()
	if argsCount == 0 { return []string{} }
	if len(it.inArgsTypesNames) == argsCount { return it.inArgsTypesNames }
	mainType := it.rvType
	slice := make([]string, 0, argsCount)
	for i := 0; i < argsCount; i++ { slice = append(slice, mainType.In(i).String()) }
	it.inArgsTypesNames = slice
	return slice
}

func (it *FuncWrap) GetOutArgsTypesNames() []string {
	if it.IsInvalid() { return []string{} }
	argsCount := it.OutArgsCount()
	if argsCount == 0 { return []string{} }
	if len(it.outArgsTypesNames) == argsCount { return it.outArgsTypesNames }
	mainType := it.rvType
	slice := make([]string, 0, argsCount)
	for i := 0; i < argsCount; i++ { slice = append(slice, mainType.Out(i).String()) }
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
	return reflectinternal.Utils.VerifyReflectTypes(it.Name, it.GetInArgsTypes(), args)
}

func (it *FuncWrap) OutArgsVerifyRv(args []reflect.Type) (isOkay bool, err error) {
	return reflectinternal.Utils.VerifyReflectTypes(it.Name, it.GetOutArgsTypes(), args)
}

func (it *FuncWrap) VoidCallNoReturn(args ...any) (processingErr error) {
	it.MustBeValid()
	_, err := it.Invoke(args)
	return err
}

func (it *FuncWrap) MustBeValid() {
	if it == nil { panic("cannot execute on nil func-wrap") }
	if it.IsInvalid() { panic("func-wrap invalid - " + it.Name) }
}

func (it *FuncWrap) ValidationError() error {
	if it == nil { return errors.New("cannot execute on nil func-wrap") }
	if it.IsInvalid() {
		return fmt.Errorf("func-wrap is invalid:\n    given type: %T\n    name: %s", it.Func, it.Name)
	}
	return nil
}

func (it *FuncWrap) InvokeMust(args ...any) []any {
	results, err := it.Invoke(args...)
	if err != nil { panic(err) }
	return results
}

func (it *FuncWrap) Invoke(args ...any) (results []any, processingErr error) {
	return it.InvokeSkip(codestack.Skip1, args...)
}

func (it *FuncWrap) InvokeSkip(skipStack int, args ...any) (results []any, processingErr error) {
	firstErr := it.ValidationError()
	if firstErr != nil { return nil, firstErr }
	argsValidationErr := it.ValidateMethodArgs(args)
	if argsValidationErr != nil { return nil, argsValidationErr }
	rvs := argsToRvFunc(args)
	var resultsRawValues []reflect.Value
	exception := trydo.WrapPanic(func() { resultsRawValues = it.rv.Call(rvs) })
	if exception != nil {
		toMsg := convertinternal.AnyTo.SmartString(exception)
		finalError := fmt.Errorf(
			"%s - func invoke failed\nstack-trace:%s\nerr:%s",
			it.GetFuncName(),
			reflectinternal.CodeStack.StacksString(codestack.Skip1+skipStack),
			toMsg,
		)
		return rvToInterfacesFunc(resultsRawValues), finalError
	}
	return rvToInterfacesFunc(resultsRawValues), nil
}

func (it *FuncWrap) VoidCall() ([]any, error) { return it.Invoke() }

func (it *FuncWrap) ValidateMethodArgs(args []any) error {
	expectedCount := it.ArgsCount()
	given := len(args)
	if given != expectedCount {
		return errors.New(it.argsCountMismatchErrorMessage(expectedCount, given, args))
	}
	_, err := it.VerifyInArgs(args)
	return err
}

func (it *FuncWrap) argsCountMismatchErrorMessage(expectedCount int, given int, args []any) string {
	expectedTypes := it.GetInArgsTypesNames()
	expectedToNames := strings.Join(expectedTypes, newLineSpaceIndent)
	actualTypes := reflectinternal.Converter.InterfacesToTypesNamesWithValues(args)
	actualTypesName := strings.Join(actualTypes, newLineSpaceIndent)
	return fmt.Sprintf(
		"%s [Func] =>\n  arguments count doesn't match for - count:\n    expected : %d\n    given    : %d\n  expected types listed :\n    - %s\n  actual given types list :\n    - %s",
		it.Name, expectedCount, given, expectedToNames, actualTypesName,
	)
}

func (it *FuncWrap) GetFirstResponseOfInvoke(args ...any) (firstResponse any, err error) {
	result, err := it.InvokeResultOfIndex(0, args...)
	if err != nil { return nil, err }
	return result, err
}

func (it *FuncWrap) InvokeResultOfIndex(index int, args ...any) (firstResponse any, err error) {
	results, err := it.Invoke(args...)
	if err != nil { return nil, err }
	return results[index], err
}

func (it *FuncWrap) InvokeError(args ...any) (funcErr, processingErr error) {
	result, err := it.GetFirstResponseOfInvoke(args...)
	if err != nil { return nil, err }
	return result.(error), err
}

// InvokeFirstAndError useful for method which looks like ReflectMethod() (something, error)
func (it *FuncWrap) InvokeFirstAndError(args ...any) (firstResponse any, funcErr, processingErr error) {
	results, processingErr := it.Invoke(args...)
	if processingErr != nil { return nil, nil, processingErr }
	if len(results) <= 1 {
		return results, nil, errors.New(it.GetFuncName() + " doesn't return at least 2 return args")
	}
	first := results[0]
	second := results[1].(error)
	return first, second, processingErr
}

func (it *FuncWrap) IsNotEqual(another *FuncWrap) bool { return !it.IsEqual(another) }

func (it *FuncWrap) IsEqualValue(another FuncWrap) bool { return it.IsEqual(&another) }

func (it *FuncWrap) IsEqual(another *FuncWrap) bool {
	if it == nil && another == nil { return true }
	if it == nil || another == nil { return false }
	if it == another { return true }
	if it.IsInvalid() != another.IsInvalid() { return false }
	if it.Name != it.Name { return false }
	if it.IsPublicMethod() != it.IsPublicMethod() { return false }
	if it.ArgsCount() != it.ArgsCount() { return false }
	if it.ReturnLength() != it.ReturnLength() { return false }
	isInArgsOkay, _ := it.InArgsVerifyRv(another.GetInArgsTypes())
	if !isInArgsOkay { return false }
	isOutArgsOkay, _ := it.OutArgsVerifyRv(another.GetOutArgsTypes())
	if !isOutArgsOkay { return false }
	return true
}

func (it *FuncWrap) InvalidError() error {
	if it == nil { return errors.New("func-wrap is nil") }
	if !it.rv.IsValid() { return errors.New("reflect value is invalid") }
	if !it.HasValidFunc() { return errors.New("func-wrap request doesn't hold a valid func reflect") }
	return nil
}

func (it FuncWrap) AsFuncWrapper() FuncWrapper { return &it }
