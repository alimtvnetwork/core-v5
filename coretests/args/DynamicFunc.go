package args

import (
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// DynamicFunc is a generic map-based argument holder with a typed WorkFunc field.
//
// Type parameter T represents the type of the WorkFunc field.
// Use DynamicFuncAny (= DynamicFunc[any]) for untyped usage.
//
// Example (typed):
//
//	df := args.DynamicFunc[func(string) error]{
//	    Params:   args.Map{"input": "test"},
//	    WorkFunc: myProcessor,
//	    Expect:   nil,
//	}
//
// Example (untyped):
//
//	df := args.DynamicFuncAny{
//	    Params:   args.Map{"input": "test"},
//	    WorkFunc: myFunc,
//	    Expect:   "result",
//	}
type DynamicFunc[T any] struct {
	Params        Map `json:",omitempty"`
	WorkFunc      T   `json:",omitempty"`
	Expect        any `json:",omitempty"`
	toSlice       []any
	isSliceCached bool
	toString      corestr.SimpleStringOnce
}

func (it *DynamicFunc[T]) ArgsCount() int {
	if it == nil {
		return 0
	}

	return it.Params.ArgsCount()
}

func (it *DynamicFunc[T]) GetWorkFunc() any {
	return it.WorkFunc
}

func (it *DynamicFunc[T]) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Params)
}

func (it *DynamicFunc[T]) HasFirst() bool {
	return reflectinternal.Is.Defined(it.FirstItem())
}

func (it *DynamicFunc[T]) GetByIndex(index int) any {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it *DynamicFunc[T]) FirstItem() any {
	return it.Params.FirstItem()
}

func (it *DynamicFunc[T]) SecondItem() any {
	return it.Params.SecondItem()
}

func (it *DynamicFunc[T]) ThirdItem() any {
	return it.Params.ThirdItem()
}

func (it *DynamicFunc[T]) FourthItem() any {
	return it.Params.FourthItem()
}

func (it *DynamicFunc[T]) FifthItem() any {
	return it.Params.FifthItem()
}

func (it *DynamicFunc[T]) SixthItem() any {
	return it.Params.SixthItem()
}

func (it *DynamicFunc[T]) Expected() any {
	return it.Expect
}

// HasDefined confirms that key is present and defined.
func (it *DynamicFunc[T]) HasDefined(name string) bool {
	if it == nil {
		return false
	}

	item, has := it.Params[name]

	return has &&
		reflectinternal.Is.Defined(item)
}

// Has confirms that key is present only.
//
// Don't confirm not null.
//
// Use HasDefined to check not null.
func (it *DynamicFunc[T]) Has(name string) bool {
	if it == nil {
		return false
	}

	_, has := it.Params[name]

	return has
}

// HasDefinedAll confirms that key is present and defined.
func (it *DynamicFunc[T]) HasDefinedAll(names ...string) bool {
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
func (it *DynamicFunc[T]) IsKeyInvalid(name string) bool {
	if it == nil {
		return false
	}

	item, has := it.Params[name]

	return !has ||
		reflectinternal.Is.Null(item)
}

// IsKeyMissing confirms yes if key is missing only.
// To check either missing or null use IsKeyInvalid.
func (it *DynamicFunc[T]) IsKeyMissing(name string) bool {
	if it == nil {
		return false
	}

	_, has := it.Params[name]

	return !has
}

func (it DynamicFunc[T]) When() (item any) {
	return it.Params["when"]
}

func (it DynamicFunc[T]) Title() (item any) {
	return it.Params["title"]
}

func (it DynamicFunc[T]) GetLowerCase(name string) (item any, isValid bool) {
	lower := strings.ToLower(name)

	return it.Get(lower)
}

func (it DynamicFunc[T]) GetDirectLower(name string) any {
	x, has := it.Params[strings.ToLower(name)]

	if has {
		return x
	}

	return nil
}

func (it DynamicFunc[T]) Actual() any {
	return it.GetDirectLower("actual")
}

func (it DynamicFunc[T]) Arrange() any {
	return it.GetDirectLower("arrange")
}

func (it *DynamicFunc[T]) Get(name string) (item any, isValid bool) {
	if it == nil {
		return nil, false
	}

	item, has := it.Params[name]

	if has {
		return item, reflectinternal.Is.Defined(item)
	}

	return nil, false
}

func (it *DynamicFunc[T]) GetAsInt(name string) (item int, isValid bool) {
	i, isValid := it.Get(name)
	isInvalid := !isValid

	if isInvalid {
		return 0, false
	}

	conv, isValid := i.(int)

	return conv, isValid
}

func (it *DynamicFunc[T]) GetAsString(name string) (item string, isValid bool) {
	i, isValid := it.Get(name)
	isInvalid := !isValid

	if isInvalid {
		return "", false
	}

	conv, isValid := i.(string)

	return conv, isValid
}

func (it *DynamicFunc[T]) GetAsStrings(name string) (items []string, isValid bool) {
	i, isValid := it.Get(name)
	isInvalid := !isValid

	if isInvalid {
		return []string{}, false
	}

	conv, isValid := i.([]string)

	return conv, isValid
}

func (it *DynamicFunc[T]) GetAsAnyItems(name string) (items []any, isValid bool) {
	i, isValid := it.Get(name)
	isInvalid := !isValid

	if isInvalid {
		return []any{}, false
	}

	conv, isValid := i.([]any)

	return conv, isValid
}

func (it *DynamicFunc[T]) HasFunc() bool {
	return it != nil &&
		reflectinternal.Is.Defined(it.WorkFunc)
}

func (it *DynamicFunc[T]) HasExpect() bool {
	return it != nil &&
		reflectinternal.Is.Defined(it.Expect)
}

func (it *DynamicFunc[T]) GetFuncName() string {
	return reflectinternal.GetFunc.NameOnly(it.WorkFunc)
}

func (it *DynamicFunc[T]) FuncWrap() *FuncWrapAny {
	return NewFuncWrap.Default(it.WorkFunc)
}

func (it *DynamicFunc[T]) Invoke(
	args ...any,
) (results []any, processingErr error) {
	return it.FuncWrap().Invoke(args...)
}

func (it *DynamicFunc[T]) InvokeMust(args ...any) (results []any) {
	results, err := it.FuncWrap().Invoke(args...)

	if err != nil {
		panic(err)
	}

	return results
}

func (it *DynamicFunc[T]) InvokeWithValidArgs() (
	results []any, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.ValidArgs()

	return funcWrap.Invoke(validArgs...)
}

func (it *DynamicFunc[T]) InvokeArgs(
	names ...string,
) (results []any, processingErr error) {
	funcWrap := it.FuncWrap()
	validArgs := it.Args(names...)

	return funcWrap.Invoke(validArgs...)
}

func (it *DynamicFunc[T]) ValidArgs() []any {
	return it.Params.ValidArgs()
}

func (it *DynamicFunc[T]) Args(names ...string) []any {
	return it.Params.Args(names...)
}

func (it *DynamicFunc[T]) Slice() []any {
	if it.isSliceCached {
		return it.toSlice
	}

	var args []any

	keys, err := converters.Map.SortedKeys(it.Params)

	if err != nil {
		panic(err)
	}

	for i, key := range keys {
		value := it.Params[key]
		args = append(
			args,
			fmt.Sprintf(
				"%d. %s : %s",
				i,
				key,
				value,
			),
		)
	}

	if it.HasFunc() {
		args = append(args, it.GetFuncName())
	}

	if it.HasExpect() {
		args = append(args, it.Expect)
	}

	it.toSlice = args
	it.isSliceCached = true

	return it.toSlice
}

func (it *DynamicFunc[T]) String() string {
	if it.toString.IsInitialized() {
		return it.toString.String()
	}

	var args []string

	for _, item := range it.Slice() {
		args = append(args, toString(item))
	}

	toFinalString := fmt.Sprintf(
		selfToStringFmt,
		"DynamicFunc",
		strings.Join(args, constants.CommaSpace),
	)

	return it.toString.GetSetOnce(toFinalString)
}

func (it DynamicFunc[T]) AsArgsMapper() ArgsMapper {
	return &it
}

func (it DynamicFunc[T]) AsArgFuncNameContractsBinder() ArgFuncNameContractsBinder {
	return &it
}

func (it DynamicFunc[T]) AsArgBaseContractsBinder() ArgBaseContractsBinder {
	return &it
}
