package args

import (
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// Dynamic is a generic map-based argument holder with a typed Expect field.
//
// Type parameter T represents the type of the Expect field.
// Use DynamicAny (= Dynamic[any]) for untyped usage.
//
// Example (typed):
//
//	d := args.Dynamic[string]{
//	    Params: args.Map{"key": "value"},
//	    Expect: "expected result",
//	}
//
// Example (untyped):
//
//	d := args.DynamicAny{
//	    Params: args.Map{"key": "value"},
//	    Expect: 42,
//	}
type Dynamic[T any] struct {
	Params        Map `json:",omitempty"`
	Expect        T   `json:",omitempty"`
	toSlice       []any
	isSliceCached bool
	toString      corestr.SimpleStringOnce
}

func (it *Dynamic[T]) ArgsCount() int {
	if it == nil {
		return 0
	}

	return it.Params.ArgsCount()
}

func (it *Dynamic[T]) GetWorkFunc() any {
	if it == nil {
		return nil
	}

	return it.Params.WorkFunc()
}

func (it *Dynamic[T]) HasFirst() bool {
	if it == nil {
		return false
	}

	return it.Params.HasFirst()
}

func (it *Dynamic[T]) GetByIndex(index int) any {
	return it.Params.GetByIndex(index)
}

func (it *Dynamic[T]) HasFunc() bool {
	return it.Params.HasFunc()
}

func (it *Dynamic[T]) GetFuncName() string {
	return it.Params.GetFuncName()
}

func (it *Dynamic[T]) Invoke(
	args ...any,
) (results []any, processingErr error) {
	return it.Params.Invoke(args...)
}

func (it *Dynamic[T]) InvokeMust(args ...any) []any {
	return it.Params.InvokeMust(args...)
}

func (it *Dynamic[T]) InvokeWithValidArgs() (
	results []any, processingErr error,
) {
	return it.Params.InvokeWithValidArgs()
}

func (it *Dynamic[T]) InvokeArgs(
	names ...string,
) (results []any, processingErr error) {
	return it.Params.InvokeArgs(names...)
}

func (it *Dynamic[T]) FuncWrap() *FuncWrapAny {
	return it.Params.FuncWrap()
}

func (it *Dynamic[T]) FirstItem() any {
	return it.Params.FirstItem()
}

func (it *Dynamic[T]) SecondItem() any {
	return it.Params.SecondItem()
}

func (it *Dynamic[T]) ThirdItem() any {
	return it.Params.ThirdItem()
}

func (it *Dynamic[T]) FourthItem() any {
	return it.Params.FourthItem()
}

func (it *Dynamic[T]) FifthItem() any {
	return it.Params.FifthItem()
}

func (it *Dynamic[T]) SixthItem() any {
	return it.Params.SixthItem()
}

func (it *Dynamic[T]) Expected() any {
	return it.Expect
}

// HasDefined confirms that key is present and defined.
func (it *Dynamic[T]) HasDefined(name string) bool {
	if it == nil {
		return false
	}

	item, has := it.Params[name]

	return has &&
		reflectinternal.Is.Defined(item)
}

// Has confirms that key is present only.
func (it *Dynamic[T]) Has(name string) bool {
	if it == nil {
		return false
	}

	_, has := it.Params[name]

	return has
}

// HasDefinedAll confirms that key is present and defined.
func (it *Dynamic[T]) HasDefinedAll(names ...string) bool {
	if it == nil || len(names) == 0 {
		return false
	}

	for _, name := range names {
		if it.IsKeyInvalid(name) {
			return false
		}
	}

	return true
}

// IsKeyInvalid confirms yes if key is missing or null.
func (it *Dynamic[T]) IsKeyInvalid(name string) bool {
	if it == nil {
		return false
	}

	item, has := it.Params[name]

	return !has ||
		reflectinternal.Is.Null(item)
}

// IsKeyMissing confirms yes if key is missing only.
func (it *Dynamic[T]) IsKeyMissing(name string) bool {
	if it == nil {
		return false
	}

	_, has := it.Params[name]

	return !has
}

func (it *Dynamic[T]) GetLowerCase(name string) (item any, isValid bool) {
	lower := strings.ToLower(name)

	return it.Get(lower)
}

func (it *Dynamic[T]) GetDirectLower(name string) any {
	x, has := it.Params[strings.ToLower(name)]

	if has {
		return x
	}

	return nil
}

func (it *Dynamic[T]) Actual() any {
	return it.GetDirectLower("actual")
}

func (it *Dynamic[T]) Arrange() any {
	return it.GetDirectLower("arrange")
}

func (it *Dynamic[T]) Get(name string) (item any, isValid bool) {
	if it == nil {
		return nil, false
	}

	item, has := it.Params[name]

	if has {
		return item, reflectinternal.Is.Defined(item)
	}

	return nil, false
}

func (it *Dynamic[T]) GetAsInt(name string) (item int, isValid bool) {
	i, isValid := it.Get(name)
	isInvalid := !isValid

	if isInvalid {
		return 0, false
	}

	conv, isValid := i.(int)

	return conv, isValid
}

func (it *Dynamic[T]) GetAsIntDefault(
	name string,
	defaultVal int,
) (item int) {
	v, isValid := it.GetAsInt(name)

	if isValid {
		return v
	}

	return defaultVal
}

func (it *Dynamic[T]) GetAsString(name string) (item string, isValid bool) {
	i, isValid := it.Get(name)
	isInvalid := !isValid

	if isInvalid {
		return "", false
	}

	conv, isValid := i.(string)

	return conv, isValid
}

func (it *Dynamic[T]) GetAsStringDefault(name string) (item string) {
	v, isValid := it.GetAsString(name)

	if isValid {
		return v
	}

	return ""
}

func (it *Dynamic[T]) GetAsStrings(name string) (items []string, isValid bool) {
	i, isValid := it.Get(name)
	isInvalid := !isValid

	if isInvalid {
		return []string{}, false
	}

	conv, isValid := i.([]string)

	return conv, isValid
}

func (it *Dynamic[T]) GetAsAnyItems(name string) (items []any, isValid bool) {
	i, isValid := it.Get(name)
	isInvalid := !isValid

	if isInvalid {
		return []any{}, false
	}

	conv, isValid := i.([]any)

	return conv, isValid
}

func (it *Dynamic[T]) HasExpect() bool {
	return it != nil &&
		reflectinternal.Is.Defined(it.Expect)
}

func (it *Dynamic[T]) ValidArgs() []any {
	var args []any

	keys := it.Params.SortedKeysMust()
	isDefined := reflectinternal.Is.Defined
	isNotFunc := reflectinternal.Is.NotFunc

	for _, key := range keys {
		val := it.Params[key]

		if isDefined(val) && isNotFunc(val) {
			args = append(args, val)
		}
	}

	return args
}

func (it *Dynamic[T]) Args(names ...string) []any {
	var args []any

	for _, key := range names {
		val := it.Params[key]
		args = append(args, val)
	}

	return args
}

func (it *Dynamic[T]) Slice() []any {
	if it.isSliceCached {
		return it.toSlice
	}

	var args []any

	keys := it.Params.SortedKeysMust()

	for i, key := range keys {
		value := it.Params[key]
		args = append(
			args,
			fmt.Sprintf("%d. %s : %s", i, key, value),
		)
	}

	if it.HasExpect() {
		args = append(args, it.Expect)
	}

	it.toSlice = args
	it.isSliceCached = true

	return it.toSlice
}

func (it *Dynamic[T]) String() string {
	if it.toString.IsInitialized() {
		return it.toString.String()
	}

	var args []string

	for _, item := range it.Slice() {
		args = append(args, toString(item))
	}

	toFinalString := fmt.Sprintf(
		selfToStringFmt,
		"Dynamic",
		strings.Join(args, constants.CommaSpace),
	)

	return it.toString.GetSetOnce(toFinalString)
}

func (it Dynamic[T]) AsArgsMapper() ArgsMapper {
	return &it
}

func (it Dynamic[T]) AsArgFuncNameContractsBinder() ArgFuncNameContractsBinder {
	return &it
}

func (it Dynamic[T]) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
