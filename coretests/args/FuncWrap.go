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
	FullName         string      `json:",omitempty"`
	Func             interface{} `json:"-,omitempty"`
	isInvalid        bool
	rvType           reflect.Type
	rv               reflect.Value
	inArgsTypesNames []string
	inArgsTypes      []reflect.Type
	outArgsTypes     []reflect.Type
}

func NewFuncWrap(anyFunc interface{}) *FuncWrap {
	if reflectinternal.IsNull(anyFunc) {
		return &FuncWrap{
			Func:      anyFunc,
			isInvalid: true,
		}
	}

	typeOf := reflect.TypeOf(anyFunc)
	kind := typeOf.Kind()

	if kind != reflect.Func {
		// invalid

		return &FuncWrap{
			Func:      anyFunc,
			isInvalid: true,
			rvType:    typeOf,
		}
	}

	// valid

	return &FuncWrap{
		Name:      reflectinternal.GetFuncName(anyFunc),
		FullName:  reflectinternal.GetFuncFullName(anyFunc),
		Func:      anyFunc,
		isInvalid: false,
		rvType:    typeOf,
		rv:        reflect.ValueOf(anyFunc),
	}
}

func (it FuncWrap) FuncName() string {
	return it.Name
}

func (it *FuncWrap) HasValidFunc() bool {
	return it != nil && !it.isInvalid && reflectinternal.IsFunc(it.Func)
}

func (it *FuncWrap) IsInvalid() bool {
	return it == nil || it.isInvalid || !it.HasValidFunc()
}

// ArgsCount returns -1 on invalid
func (it *FuncWrap) ArgsCount() int {
	if it.IsInvalid() {
		return -1
	}

	// https://stackoverflow.com/a/47626214

	return it.rvType.NumIn()
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

	return it.rvType.NumOut()
}

func (it *FuncWrap) IsPublicMethod() bool {
	return it != nil && it.rvType.PkgPath() == ""
}

func (it *FuncWrap) IsPrivateMethod() bool {
	return it != nil && it.rvType.PkgPath() != ""
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

	values := it.rv.Call(argsReflectValues)

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

func (it *FuncWrap) InvokeMust(
	args ...interface{},
) []interface{} {
	results, err := it.Invoke(args...)

	if err != nil {
		panic(err)
	}

	return results
}

func (it *FuncWrap) Invoke(
	args ...interface{},
) ([]interface{}, error) {
	firstErr := it.validationError()

	if firstErr != nil {
		return nil, firstErr
	}

	argsValidationErr := it.ValidateMethodArgs(args)

	if argsValidationErr != nil {
		return nil, argsValidationErr
	}

	rvs := argsToRvFunc(args)
	resultsRawValues := it.rv.Call(rvs)

	return rvToInterfacesFunc(resultsRawValues), nil
}

func (it *FuncWrap) ValidateMethodArgs(args []interface{}) error {
	expectedCount := it.ArgsCount()
	given := len(args)

	if given != expectedCount {
		return errors.New(it.argsCountMismatchErrorMessage(expectedCount, given, args))
	}

	_, err := it.VerifyInArgs(args)

	return err
}

func (it *FuncWrap) argsCountMismatchErrorMessage(
	expectedCount int,
	given int,
	args []interface{},
) string {
	expectedTypes := it.GetInArgsTypesNames()
	expectedToNames := strings.Join(expectedTypes, "\n   - ")
	actualTypes := reflectinternal.Converter.InterfacesToTypesNamesWithValues(args)
	actualTypesName := strings.Join(actualTypes, "\n   - ")

	return fmt.Sprintf(
		"%s [Func] => "+
			"arguments count doesn't match for - count - \n"+
			"   expected : %d,\n"+
			"   given    : %d\n"+
			"expected types listed :\n"+
			"   - %s\n"+
			"actual given types list :\n"+
			"   - %s",
		it.Name,
		expectedCount,
		given,
		expectedToNames,
		actualTypesName,
	)
}

func (it *FuncWrap) GetFirstResponseOfInvoke(
	args ...interface{},
) (firstResponse interface{}, err error) {
	result, err := it.InvokeResultOfIndex(0, args...)

	if err != nil {
		return nil, err
	}

	return result, err
}

func (it *FuncWrap) InvokeResultOfIndex(
	index int,
	args ...interface{},
) (firstResponse interface{}, err error) {
	results, err := it.Invoke(args...)

	if err != nil {
		return nil, err
	}

	return results[index], err
}

func (it *FuncWrap) InvokeError(
	args ...interface{},
) (funcErr, processingErr error) {
	result, err := it.GetFirstResponseOfInvoke(args...)

	if err != nil {
		return nil, err
	}

	return result.(error), err
}

// InvokeFirstAndError
//
//	useful for method which looks like ReflectMethod() (soemthing, error)
func (it *FuncWrap) InvokeFirstAndError(
	args ...interface{},
) (firstResponse interface{}, funcErr, processingErr error) {
	results, processingErr := it.Invoke(args...)

	if processingErr != nil {
		return nil, nil, processingErr
	}

	if len(results) <= 1 {
		return nil, nil, errors.New(it.FuncName() + " doesn't return at least 2 args")
	}

	first := results[0]
	second := results[1].(error)

	return first, second, processingErr
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
