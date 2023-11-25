package args

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type FuncWrap struct {
	Name             string      `json:",omitempty"`
	Func             interface{} `json:"-,omitempty"`
	rvType           reflect.Type
	method           reflect.Method
	inArgsTypesNames []string
	inArgsTypes      []reflect.Type
	outArgsTypes     []reflect.Type
}

func (it FuncWrap) FuncName() string {
	return it.Name
}

func (it *FuncWrap) HasValidFunc() bool {
	return it != nil && reflectinternal.IsFunc(it.Func)
}

func (it *FuncWrap) IsInvalid() bool {
	return it == nil || !it.HasValidFunc()
}

// ArgsCount returns -1 on invalid
func (it *FuncWrap) ArgsCount() int {
	if it.IsInvalid() {
		return -1
	}

	// https://stackoverflow.com/a/47626214

	return it.method.Type.NumIn()
}

// ArgsLength is an Alias for ArgsCount
func (it *FuncWrap) ArgsLength() int {
	return it.ArgsCount()
}

// ReturnLength refers to the return arguments length
func (it *FuncWrap) ReturnLength() int {
	if it.IsInvalid() {
		return -1
	}

	// https://stackoverflow.com/a/47626214

	return it.method.Type.NumOut()
}

func (it *FuncWrap) IsPublicMethod() bool {
	return it != nil && it.method.PkgPath == ""
}

func (it *FuncWrap) IsPrivateMethod() bool {
	return it != nil && it.method.PkgPath != ""
}

func (it *FuncWrap) GetType() reflect.Type {
	if it.IsInvalid() {
		return nil
	}

	return it.rvType
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

	// https://go.dev/play/p/dpIspUFfbu0
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

	// https://go.dev/play/p/dpIspUFfbu0
	mainType := it.rvType
	slice := make([]reflect.Type, 0, argsCount)

	for i := 0; i < argsCount; i++ {
		slice = append(slice, mainType.In(i))
	}

	it.inArgsTypes = slice

	return slice
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

	// https://go.dev/play/p/dpIspUFfbu0
	mainType := it.rvType
	slice := make([]string, 0, argsCount)

	for i := 0; i < argsCount; i++ {
		slice = append(slice, mainType.In(i).Name())
	}

	it.inArgsTypesNames = slice

	return slice
}

func (it *FuncWrap) VerifyInArgs(args []interface{}) (isOkay bool, err error) {
	toTypes := reflectinternal.Converter.InterfacesToTypes(args)

	return it.InArgsVerifyRv(toTypes)
}

func (it *FuncWrap) VerifyOutArgs(args []interface{}) (isOkay bool, err error) {
	toTypes := reflectinternal.Converter.InterfacesToTypes(args)

	return it.OutArgsVerifyRv(toTypes)
}

func (it *FuncWrap) InArgsVerifyRv(args []reflect.Type) (isOkay bool, err error) {
	return reflectinternal.Utils.VerifyReflectTypes(it.GetInArgsTypes(), args)
}

func (it *FuncWrap) OutArgsVerifyRv(args []reflect.Type) (isOkay bool, err error) {
	return reflectinternal.Utils.VerifyReflectTypes(it.GetOutArgsTypes(), args)
}

func (it *FuncWrap) InvokeDirectly(
	args ...interface{},
) (returnedValues []reflect.Value, err error) {
	it.mustBeValid()

	argsReflectValues := argsToRvFunc(
		args,
	)

	values := it.method.Func.Call(argsReflectValues)

	return values, nil
}

func (it *FuncWrap) InvokeMethodDirectlyVoid(
	args ...interface{},
) error {
	it.mustBeValid()
	it.Invoke(args)

	return nil
}

func (it *FuncWrap) mustBeValid() {
	if it == nil {
		panic("cannot execute on nil func-wrap")
	}

	if it.IsInvalid() {
		panic("func-wrap invalid - " + it.Name)
	}
}

func (it *FuncWrap) validationError() error {
	if it == nil {
		return errors.New("cannot execute on nil func-wrap")
	}

	if it.IsInvalid() {
		return errors.New("func-wrap is invalid - " + it.Name)
	}

	return nil
}

func (it *FuncWrap) InvokeVoidMethod(
	args ...interface{},
) {
	it.mustBeValid()

	argsReflectValues := argsToRvFunc(args)
	it.method.
		Func.
		Call(argsReflectValues)
}

func (it *FuncWrap) InvokeMethod(
	args ...interface{},
) []reflect.Value {
	it.mustBeValid()
	it.ValidateMethodArgs(args)

	argsReflectValues := argsToRvFunc(args)

	return it.
		method.
		Func.
		Call(argsReflectValues)
}

func (it *FuncWrap) Invoke(
	args ...interface{},
) []interface{} {
	it.ValidateMethodArgs(args)

	returnedValues := it.InvokeMethod(
		args...,
	)

	return rvToInterfacesFunc(
		returnedValues,
	)
}

func (it *FuncWrap) ValidateMethodArgs(args []interface{}) {
	expectedCount := it.ArgsCount()
	given := len(args)

	if given != expectedCount {
		panic(it.argsCountMismatchErrorMessage(expectedCount, given, args))
	}

	_, err := it.VerifyInArgs(args)

	if err != nil {
		panic(err)
	}
}

func (it *FuncWrap) argsCountMismatchErrorMessage(expectedCount int, given int, args []interface{}) string {
	expectedTypes := it.GetInArgsTypesNames()
	expectedToNames := strings.Join(expectedTypes, "\n\t -")
	actualTypes := reflectinternal.Converter.InterfacesToTypesNames(args)
	actualTypesName := strings.Join(actualTypes, "\n\t -")

	return fmt.Sprintf(
		"%s [Method] -> "+
			"arguments count doesn't match for - Count - expected : "+
			"%d, given : %d\nexpected types listed : %s\nactual given types list : %s",
		it.Name,
		expectedCount,
		given,
		expectedToNames,
		actualTypesName,
	)
}

func (it *FuncWrap) GetFirstResponseOfInvoke(
	args ...interface{},
) (
	firstResponse interface{},
) {
	return it.GetResponseOfIndexFromInvoke(0, args...)
}

func (it *FuncWrap) GetResponseOfIndexFromInvoke(
	index int,
	args ...interface{},
) (
	firstResponse interface{},
) {
	results := it.Invoke(args...)

	return results[index]
}

func (it *FuncWrap) InvokeError(
	args ...interface{},
) (err error) {
	return it.GetFirstResponseOfInvoke(args...).(error)
}

// InvokeFirstAndError
//
//	useful for method which looks like ReflectMethod() (soemthing, error)
func (it *FuncWrap) InvokeFirstAndError(
	args ...interface{},
) (
	firstResponse interface{}, err error,
) {
	responses := it.Invoke(args...)

	first := responses[0]
	second := responses[1].(error)

	return first, second
}

// IsNotEqual
//
// Based on predication.
//
// Warning: it can be wrong as well
func (it *FuncWrap) IsNotEqual(
	another *FuncWrap,
) bool {
	return !it.IsEqual(another)
}

// IsEqual
//
// Based on predication.
//
// Warning: it can be wrong as well
func (it *FuncWrap) IsEqual(
	another *FuncWrap,
) bool {
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

	// can be skipped,
	// because name also refers to public or private
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

	// most probably true,
	// but can be false as well

	return true
}
