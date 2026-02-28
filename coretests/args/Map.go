package args

import (
	"fmt"
	"sort"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/internal/convertinternal"
	"gitlab.com/auk-go/core/internal/msgcreator"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type Map map[string]any

func (it Map) GetWorkFunc() any {
	return it.WorkFunc()
}

func (it Map) ArgsCount() int {
	l := it.Length()

	var count int

	if it.HasExpect() {
		count++
	}

	if it.HasFunc() {
		count++
	}

	return l - count
}

func (it Map) Length() int {
	return len(it)
}

func (it Map) Expected() any {
	return it.GetFirstOfNames(
		"expected",
		"expects",
		"expect",
	)
}

func (it Map) HasFirst() bool {
	return reflectinternal.Is.Defined(it.FirstItem())
}

func (it Map) HasExpect() bool {
	return reflectinternal.Is.Defined(it.Expected())
}

func (it Map) GetByIndex(index int) any {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

func (it Map) HasFunc() bool {
	return reflectinternal.Is.Defined(it.FuncWrap())
}

func (it Map) GetFuncName() string {
	funcWrap := it.FuncWrap()

	if funcWrap != nil {
		return funcWrap.Name
	}

	return ""
}

// HasDefined
//
// Confirms that key is present and defined.
func (it Map) HasDefined(name string) bool {
	if it == nil {
		return false
	}

	item, has := it[name]

	return has &&
		reflectinternal.Is.Defined(item)
}

// Has
//
//	Confirms that key is present only.
//
//	Don't confirm not null.
//
// Use HasDefined to check not null.
func (it Map) Has(name string) bool {
	if it == nil {
		return false
	}

	_, has := it[name]

	return has
}

// HasDefinedAll
//
// Confirms that key is present and defined.
func (it Map) HasDefinedAll(names ...string) bool {
	if it == nil || len(names) == 0 {
		return false
	}

	for _, name := range names {
		if it.IsKeyInvalid(name) {
			return false
		}
	}

	// all defined

	return true
}

// IsKeyInvalid
//
// confirms yes if key is missing or null
func (it Map) IsKeyInvalid(name string) bool {
	if it == nil {
		return false
	}

	item, has := it[name]

	return !has ||
		reflectinternal.Is.Null(item)
}

// IsKeyMissing
//
// confirms yes if key is missing  only.
// To check either missing or null use IsKeyInvalid.
func (it Map) IsKeyMissing(name string) bool {
	if it == nil {
		return false
	}

	_, has := it[name]

	return !has
}

func (it Map) SortedKeys() ([]string, error) {
	if len(it) == 0 {
		return []string{}, nil
	}

	return convertinternal.
		Map.
		SortedKeys(it.Raw())
}

func (it Map) SortedKeysMust() []string {
	sortedKeys, err := it.SortedKeys()

	if err != nil {
		panic(err)
	}

	return sortedKeys
}

func (it Map) When() (item any) {
	return it["when"]
}

func (it Map) Title() (item any) {
	return it["title"]
}

func (it Map) Get(name string) (item any, isValid bool) {
	if it == nil {
		return nil, false
	}

	item, has := it[name]

	if has {
		return item, reflectinternal.Is.Defined(item)
	}

	return nil, false
}

func (it Map) GetLowerCase(name string) (item any, isValid bool) {
	lower := strings.ToLower(name)

	return it.Get(lower)
}

func (it Map) GetDirectLower(name string) any {
	x, has := it[strings.ToLower(name)]

	if has {
		return x
	}

	return nil
}

func (it Map) Expect() any {
	return it.GetDirectLower("expect")
}

func (it Map) Actual() any {
	return it.GetDirectLower("actual")
}

func (it Map) Arrange() any {
	return it.GetDirectLower("arrange")
}

func (it Map) FirstItem() any {
	return it.GetFirstOfNames("first", "f1", "p1", "1")
}

func (it Map) SecondItem() any {
	return it.GetFirstOfNames("second", "f2", "p2", "2")
}

func (it Map) ThirdItem() any {
	return it.GetFirstOfNames("third", "f3", "p3", "3")
}

func (it Map) FourthItem() any {
	return it.GetFirstOfNames("fourth", "f4", "p4", "4")
}

func (it Map) FifthItem() any {
	return it.GetFirstOfNames("fifth", "f5", "p5", "5")
}

func (it Map) SixthItem() any {
	return it.GetFirstOfNames("sixth", "f6", "p6", "6")
}

func (it Map) Seventh() any {
	return it.GetFirstOfNames("seventh", "f7", "p7", "7")
}

func (it Map) SetActual(actual any) {
	it["actual"] = actual
}

func (it Map) WorkFunc() any {
	return it.GetFirstOfNames(
		"func",
		"work.func",
		"workFunc",
	)
}

func (it Map) GetFirstOfNames(names ...string) any {
	if len(names) == 0 {
		return nil
	}

	for _, name := range names {
		v, has := it[name]

		if has && reflectinternal.Is.Defined(v) {
			return v
		}
	}

	return nil
}

func (it Map) GetAsStringSliceFirstOfNames(names ...string) []string {
	if len(names) == 0 {
		return nil
	}

	item := it.GetFirstOfNames(names...)

	if reflectinternal.Is.Defined(item) {
		return item.([]string)
	}

	return nil
}

func (it Map) WorkFuncName() string {
	workFunc := it.WorkFunc()

	return reflectinternal.GetFunc.NameOnly(workFunc)
}

func (it Map) FuncWrap() *FuncWrap {
	return NewFuncWrap.Default(it.WorkFunc())
}

func (it Map) Invoke(args ...any) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(args...)
}

func (it Map) InvokeMust(args ...any) (results []any) {
	results, err := it.FuncWrap().Invoke(args...)

	if err != nil {
		panic(err)
	}

	return results
}

func (it Map) InvokeWithValidArgs() (
	results []any, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.ValidArgs()

	return funcWrap.Invoke(validArgs...)
}

func (it Map) InvokeArgs(names ...string) (
	results []any, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.Args(names...)

	return funcWrap.Invoke(validArgs...)
}

func (it Map) ValidArgs() []any {
	var args []any

	keys, _ := it.SortedKeys()
	isDefined := reflectinternal.Is.Defined
	isNotFunc := reflectinternal.Is.NotFunc

	for _, key := range keys {
		val := it[key]

		if isDefined(val) && isNotFunc(val) {
			args = append(args, val)
		}
	}

	return args
}

func (it Map) Raw() map[string]any {
	return it
}

func (it Map) Args(names ...string) []any {
	var args []any

	for _, key := range names {
		val := it[key]
		args = append(args, val)
	}

	return args
}

func (it Map) GetFirstFuncNameOf(names ...string) string {
	workFunc := it.GetFirstOfNames(names...)

	return reflectinternal.GetFunc.NameOnly(workFunc)
}

func (it Map) GetAsInt(name string) (item int, isValid bool) {
	i, isValid := it.Get(name)

	if !isValid {
		return 0, false
	}

	conv, isValid := i.(int)

	return conv, isValid
}

func (it Map) GetAsIntDefault(name string, defaultVal int) (item int) {
	v, isValid := it.GetAsInt(name)

	if isValid {
		return v
	}

	return defaultVal
}

func (it Map) GetAsString(name string) (item string, isValid bool) {
	i, isValid := it.Get(name)

	if !isValid {
		return "", false
	}

	conv, isValid := i.(string)

	return conv, isValid
}

func (it Map) GetAsStringDefault(name string) (item string) {
	v, isValid := it.GetAsString(name)

	if isValid {
		return v
	}

	return ""
}

func (it Map) GetAsStrings(name string) (items []string, isValid bool) {
	i, isValid := it.Get(name)

	if !isValid {
		return []string{}, false
	}

	conv, isValid := i.([]string)

	return conv, isValid
}

func (it Map) GetAsAnyItems(name string) (items []any, isValid bool) {
	i, isValid := it.Get(name)

	if !isValid {
		return []any{}, false
	}

	conv, isValid := i.([]any)

	return conv, isValid
}

func (it Map) Slice() []any {
	var slice []any

	keys, err := converters.Map.SortedKeys(it.Raw())

	if err != nil {
		panic(err)
	}

	for _, key := range keys {
		value := it[key]
		slice = append(
			slice, fmt.Sprintf(
				"%s : %#v",
				key,
				value,
			),
		)
	}

	return slice
}

func (it Map) String() string {
	var args []string

	for _, item := range it.Slice() {
		args = append(args, toString(item))
	}

	toLines := msgcreator.Assert.StringsToSpaceStringUsingFunc(
		4,
		func(i int, spacePrefix, line string) string {
			return fmt.Sprintf(
				"%s%s,",
				spacePrefix,
				line,
			)
		},
		args...,
	)

	sort.Strings(toLines)

	toFinalString := fmt.Sprintf(
		"%s {\n%s\n}\n",
		"Map",
		strings.Join(toLines, constants.NewLineUnix),
	)

	return toFinalString
}

func (it Map) AsArgsMapper() ArgsMapper {
	return &it
}
