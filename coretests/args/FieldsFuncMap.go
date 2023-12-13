package args

import (
	"errors"
	"reflect"

	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type FieldsFuncMap map[string]FuncWrap

func (it FieldsFuncMap) IsEmpty() bool {
	return len(it) == 0
}

func (it FieldsFuncMap) Length() int {
	return len(it)
}

func (it FieldsFuncMap) Count() int {
	return len(it)
}

func (it FieldsFuncMap) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it FieldsFuncMap) Has(name string) bool {
	if it.IsEmpty() {
		return false
	}

	_, isFound := it[name]

	return isFound
}

func (it FieldsFuncMap) IsContains(name string) bool {
	return it.Has(name)
}

func (it FieldsFuncMap) Get(name string) *FuncWrap {
	if it.IsEmpty() {
		return nil
	}

	f, isFound := it[name]

	if isFound {
		return &f
	}

	return nil
}

func (it FieldsFuncMap) GetPascalCaseFuncName(name string) string {
	if len(it) == 0 {
		return ""
	}

	return reflectinternal.
		GetFunc.
		PascalFuncName(name)
}

func (it FieldsFuncMap) IsValidFuncOf(name string) bool {
	f := it.Get(name)

	if f == nil {
		return false
	}

	return f.HasValidFunc()
}

func (it FieldsFuncMap) IsInvalidFunc(name string) bool {
	f := it.Get(name)

	if f == nil {
		return true
	}

	return f.IsInvalid()
}

func (it FieldsFuncMap) PkgPath(name string) string {
	f := it.Get(name)

	if f == nil {
		return ""
	}

	return f.PkgPath()
}

func (it FieldsFuncMap) PkgNameOnly(name string) string {
	f := it.Get(name)

	if f == nil {
		return ""
	}

	return f.PkgNameOnly()
}

func (it FieldsFuncMap) FuncDirectInvokeName(name string) string {
	f := it.Get(name)

	if f == nil {
		return ""
	}

	return f.FuncDirectInvokeName()
}

// ArgsCount returns -1 on invalid
func (it FieldsFuncMap) ArgsCount(name string) int {
	f := it.Get(name)

	if f == nil {
		return 0
	}

	return f.ArgsCount()
}

// ArgsLength is an Alias for ArgsCount
func (it FieldsFuncMap) ArgsLength(name string) int {
	return it.ArgsCount(name)
}

// ReturnLength refers to the return arguments length
func (it FieldsFuncMap) ReturnLength(name string) int {
	f := it.Get(name)

	if f == nil {
		return 0
	}

	return f.ReturnLength()
}

func (it FieldsFuncMap) IsPublicMethod(name string) bool {
	f := it.Get(name)

	if f == nil {
		return false
	}

	return f.IsPublicMethod()
}

func (it FieldsFuncMap) IsPrivateMethod(name string) bool {
	f := it.Get(name)

	if f == nil {
		return false
	}

	return f.IsPrivateMethod()
}

func (it FieldsFuncMap) GetType(name string) reflect.Type {
	f := it.Get(name)

	if f == nil {
		return reflect.Type(nil)
	}

	return f.GetType()
}

func (it FieldsFuncMap) GetOutArgsTypes(name string) []reflect.Type {
	f := it.Get(name)

	if f == nil {
		return []reflect.Type{}
	}

	return f.GetOutArgsTypes()
}

func (it FieldsFuncMap) GetInArgsTypes(name string) []reflect.Type {
	f := it.Get(name)

	if f == nil {
		return []reflect.Type{}
	}

	return f.GetOutArgsTypes()
}

func (it FieldsFuncMap) GetInArgsTypesNames(name string) []string {
	f := it.Get(name)

	if f == nil {
		return []string{}
	}

	return f.GetInArgsTypesNames()
}

func (it FieldsFuncMap) VerifyInArgs(name string, args []interface{}) (isOkay bool, err error) {
	f := it.Get(name)

	if f == nil {
		return false, it.notFoundErr(name)
	}

	return f.VerifyInArgs(args)
}

func (it FieldsFuncMap) VerifyOutArgs(name string, args []interface{}) (isOkay bool, err error) {
	f := it.Get(name)

	if f == nil {
		return false, it.notFoundErr(name)
	}

	return f.VerifyOutArgs(args)
}

func (it FieldsFuncMap) InArgsVerifyRv(name string, args []reflect.Type) (isOkay bool, err error) {
	f := it.Get(name)

	if f == nil {
		return false, it.notFoundErr(name)
	}

	return f.InArgsVerifyRv(args)
}

func (it FieldsFuncMap) OutArgsVerifyRv(name string, args []reflect.Type) (isOkay bool, err error) {
	f := it.Get(name)

	if f == nil {
		return false, it.notFoundErr(name)
	}

	return f.OutArgsVerifyRv(args)
}

func (it FieldsFuncMap) VoidCallNoReturn(
	name string,
	args ...interface{},
) (processingErr error) {
	f := it.Get(name)

	if f == nil {
		return it.notFoundErr(name)
	}

	return f.VoidCallNoReturn(args...)
}

func (it FieldsFuncMap) MustBeValid(name string) {
	f := it.Get(name)

	if f == nil {
		panic(it.notFoundErr(name))
	}

	f.MustBeValid()
}

func (it FieldsFuncMap) ValidationError(name string) error {
	f := it.Get(name)

	if f == nil {
		return it.notFoundErr(name)
	}

	return f.ValidationError()
}

func (it FieldsFuncMap) InvokeMust(
	name string,
	args ...interface{},
) []interface{} {
	results, err := it.Invoke(name, args...)

	if err != nil {
		panic(err)
	}

	return results
}

func (it FieldsFuncMap) Invoke(
	name string,
	args ...interface{},
) (results []interface{}, processingErr error) {
	return it.InvokeSkip(codestack.Skip1, name, args...)
}

func (it FieldsFuncMap) InvokeSkip(
	skipStack int,
	name string,
	args ...interface{},
) (results []interface{}, processingErr error) {
	f := it.Get(name)

	if f == nil {
		return []interface{}{}, it.notFoundErr(name)
	}

	return f.InvokeSkip(skipStack+1, args)
}

func (it FieldsFuncMap) VoidCall(name string) ([]interface{}, error) {
	return it.Invoke(name)
}

func (it FieldsFuncMap) ValidateMethodArgs(name string, args []interface{}) error {
	expectedCount := it.Get(name)
	given := len(args)

	if given != expectedCount {
		return errors.New(it.argsCountMismatchErrorMessage(expectedCount, given, args))
	}

	_, err := it.VerifyInArgs(args)

	return err
}

func (it FieldsFuncMap) GetFirstResponseOfInvoke(
	name string,
	args ...interface{},
) (firstResponse interface{}, err error) {
	result, err := it.InvokeResultOfIndex(0, args...)

	if err != nil {
		return nil, err
	}

	return result, err
}

func (it FieldsFuncMap) InvokeResultOfIndex(
	name string,
	index int,
	args ...interface{},
) (firstResponse interface{}, err error) {
	results, err := it.Invoke(args...)

	if err != nil {
		return nil, err
	}

	return results[index], err
}

func (it FieldsFuncMap) InvokeError(
	name string,
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
func (it FieldsFuncMap) InvokeFirstAndError(
	name string,
	args ...interface{},
) (firstResponse interface{}, funcErr, processingErr error) {
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

// IsNotEqual
//
// Based on predication.
//
// Warning: it can be wrong as well
func (it FieldsFuncMap) IsNotEqual(
	another FieldsFuncMap,
) bool {
	return !it.IsEqual(another)
}

func (it FieldsFuncMap) IsEqualValue(
	another FuncWrap,
) bool {
	return it.IsEqual(&another)
}

// IsEqual
//
// Based on predication.
//
// Warning: it can be wrong as well
func (it FieldsFuncMap) IsEqual(
	another FieldsFuncMap,
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

func (it FieldsFuncMap) InvalidError() error {
	if it == nil {
		return errors.New("func-wrap is nil")
	}

	if !it.rv.IsValid() {
		return errors.New("reflect value is invalid")
	}

	if !it.HasValidFunc() {
		return errors.New("func-wrap request doesn't hold a valid func reflect")
	}

	return nil
}

func (it FieldsFuncMap) notFoundErr(name string) error {
	return errcore.NotFound.Error("func-wrap not found by the name", name)
}

func (it FuncWrap) AsFuncWrapper() FuncWrapper {
	return &it
}
