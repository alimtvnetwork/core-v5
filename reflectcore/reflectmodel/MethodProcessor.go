package reflectmodel

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type MethodProcessor struct {
	Name             string
	Index            int
	ReflectMethod    reflect.Method
	inArgsTypesNames []string       `json:"-"`
	inArgsTypes      []reflect.Type `json:"-"`
	outArgsTypes     []reflect.Type `json:"-"`
}

var (
	util               = rvUtils{}
	newLineSpaceIndent = "\n    - "
)

func (it *MethodProcessor) HasValidFunc() bool {
	return it != nil
}

func (it *MethodProcessor) GetFuncName() string {
	return it.Name
}

func (it *MethodProcessor) IsInvalid() bool {
	return it == nil
}

func (it *MethodProcessor) Func() *reflect.Value {
	if it.IsInvalid() {
		return nil
	}
	return &it.ReflectMethod.Func
}

func (it *MethodProcessor) ArgsCount() int {
	return it.ReflectMethod.Type.NumIn()
}

func (it *MethodProcessor) ReturnLength() int {
	if it.IsInvalid() {
		return -1
	}
	return it.GetType().NumOut()
}

func (it *MethodProcessor) IsPublicMethod() bool {
	return it != nil && it.ReflectMethod.PkgPath == ""
}

func (it *MethodProcessor) IsPrivateMethod() bool {
	return it != nil && it.ReflectMethod.PkgPath != ""
}

func (it *MethodProcessor) ArgsLength() int {
	return it.ReflectMethod.Type.NumIn()
}

func (it *MethodProcessor) Invoke(args ...any) (
	responses []any,
	err error,
) {
	firstErr := it.validationError()
	if firstErr != nil {
		return nil, firstErr
	}

	argsValidationErr := it.ValidateMethodArgs(args)
	if argsValidationErr != nil {
		return nil, argsValidationErr
	}

	rvs := util.ArgsToReflectValues(args)
	resultsRawValues := it.Func().Call(rvs)

	return util.ReflectValuesToInterfaces(resultsRawValues), nil
}

func (it *MethodProcessor) GetFirstResponseOfInvoke(
	args ...any,
) (firstResponse any, err error) {
	result, err := it.InvokeResultOfIndex(0, args...)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (it *MethodProcessor) InvokeResultOfIndex(
	index int,
	args ...any,
) (firstResponse any, err error) {
	results, err := it.Invoke(args...)
	if err != nil {
		return nil, err
	}
	return results[index], err
}

func (it *MethodProcessor) InvokeError(
	args ...any,
) (funcErr, processingErr error) {
	result, err := it.GetFirstResponseOfInvoke(args...)
	if err != nil {
		return nil, err
	}
	return result.(error), err
}

func (it *MethodProcessor) InvokeFirstAndError(
	args ...any,
) (firstResponse any, funcErr, processingErr error) {
	results, processingErr := it.Invoke(args...)
	if processingErr != nil {
		return nil, nil, processingErr
	}

	if len(results) <= 1 {
		return results,
			nil,
			errors.New(it.GetFuncName() + " doesn't return at least 2 return args")
	}

	first := results[0]
	second := results[1].(error)

	return first, second, processingErr
}

func (it *MethodProcessor) IsNotEqual(another *MethodProcessor) bool {
	return !it.IsEqual(another)
}

func (it *MethodProcessor) IsEqual(another *MethodProcessor) bool {
	if it == nil && another == nil {
		return true
	}
	if it == nil || another == nil {
		return false
	}
	if it == another {
		return true
	}
	if it.IsInvalid() != another.IsInvalid() {
		return false
	}
	if it.Name != it.Name {
		return false
	}
	if it.IsPublicMethod() != it.IsPublicMethod() {
		return false
	}
	if it.ArgsCount() != it.ArgsCount() {
		return false
	}
	if it.ReturnLength() != it.ReturnLength() {
		return false
	}

	isInArgsOkay, _ := it.InArgsVerifyRv(another.GetInArgsTypes())
	if !isInArgsOkay {
		return false
	}

	isOutArgsOkay, _ := it.OutArgsVerifyRv(another.GetOutArgsTypes())
	if !isOutArgsOkay {
		return false
	}

	return true
}

func (it *MethodProcessor) GetType() reflect.Type {
	if it.IsInvalid() {
		return nil
	}
	return it.ReflectMethod.Type
}

func (it *MethodProcessor) GetOutArgsTypes() []reflect.Type {
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

	mainType := it.GetType()
	slice := make([]reflect.Type, 0, argsOutCount)
	for i := 0; i < argsOutCount; i++ {
		slice = append(slice, mainType.Out(i))
	}

	it.outArgsTypes = slice
	return slice
}

func (it *MethodProcessor) GetInArgsTypes() []reflect.Type {
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

	mainType := it.GetType()
	slice := make([]reflect.Type, 0, argsCount)
	for i := 0; i < argsCount; i++ {
		slice = append(slice, mainType.In(i))
	}

	it.inArgsTypes = slice
	return slice
}

func (it *MethodProcessor) GetInArgsTypesNames() []string {
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

	mainType := it.GetType()
	slice := make([]string, 0, argsCount)
	for i := 0; i < argsCount; i++ {
		slice = append(slice, mainType.In(i).String())
	}

	it.inArgsTypesNames = slice
	return slice
}

func (it *MethodProcessor) validationError() error {
	if it == nil {
		return errors.New("cannot execute on nil func-wrap")
	}

	if it.IsInvalid() {
		return fmt.Errorf(
			"func-wrap is invalid:\n"+
				"    given type: %T\n"+
				"    name: %s",
			it.Func(),
			it.Name,
		)
	}

	return nil
}

func (it *MethodProcessor) ValidateMethodArgs(args []any) error {
	expectedCount := it.ArgsCount()
	given := len(args)

	if given != expectedCount {
		return errors.New(it.argsCountMismatchErrorMessage(expectedCount, given, args))
	}

	_, err := it.VerifyInArgs(args)
	return err
}

func (it *MethodProcessor) argsCountMismatchErrorMessage(
	expectedCount int,
	given int,
	args []any,
) string {
	expectedTypes := it.GetInArgsTypesNames()
	expectedToNames := strings.Join(expectedTypes, newLineSpaceIndent)
	actualTypes := util.InterfacesToTypesNamesWithValues(args)
	actualTypesName := strings.Join(actualTypes, newLineSpaceIndent)

	return fmt.Sprintf(
		"%s [Func] =>\n"+
			"  arguments count doesn't match for - count:\n"+
			"    expected : %d\n"+
			"    given    : %d\n"+
			"  expected types listed :\n"+
			"    - %s\n"+
			"  actual given types list :\n"+
			"    - %s",
		it.Name,
		expectedCount,
		given,
		expectedToNames,
		actualTypesName,
	)
}

func (it *MethodProcessor) VerifyInArgs(args []any) (isOkay bool, err error) {
	toTypes := util.InterfacesToTypes(args)
	return it.InArgsVerifyRv(toTypes)
}

func (it *MethodProcessor) VerifyOutArgs(args []any) (isOkay bool, err error) {
	toTypes := util.InterfacesToTypes(args)
	return it.OutArgsVerifyRv(toTypes)
}

func (it *MethodProcessor) InArgsVerifyRv(args []reflect.Type) (isOkay bool, err error) {
	return util.VerifyReflectTypes(it.Name, it.GetInArgsTypes(), args)
}

func (it *MethodProcessor) OutArgsVerifyRv(args []reflect.Type) (isOkay bool, err error) {
	return util.VerifyReflectTypes(it.Name, it.GetOutArgsTypes(), args)
}
