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

// GetWorkFunc
//
//	aliases:
//		Func
//		Work.Func
//		WorkFunc
func (it Map) GetWorkFunc() any {
	return it.WorkFunc()
}

// ArgsCount
//
//	Returns the number of arguments in the map, excluding "expected" and "func".
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

// Length
//
//	Returns the number of items in the map.
func (it Map) Length() int {
	return len(it)
}

// Expected
//
//	Returns the expected value from the map, checking for keys "expected", "expects", and "expect".
func (it Map) Expected() any {
	return it.GetFirstOfNames(
		"expected",
		"expects",
		"expect",
	)
}

// HasFirst
//
//	Checks if the map has a defined first item.
func (it Map) HasFirst() bool {
	return reflectinternal.Is.Defined(it.FirstItem())
}

// HasExpect
//
//	Checks if the map has a defined expected value.
func (it Map) HasExpect() bool {
	return reflectinternal.Is.Defined(it.Expected())
}

// GetByIndex
//
//	Retrieves an item from the map by its index in the sorted keys.
func (it Map) GetByIndex(index int) any {
	slice := it.Slice()

	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

// HasFunc
//
//	Checks if the map has a defined function.
func (it Map) HasFunc() bool {
	return reflectinternal.Is.Defined(it.FuncWrap())
}

// GetFuncName
//
//	Retrieves the name of the function from the map.
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
// Don't confirm not null.
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

// SortedKeys
//
//	Returns the keys of the map in sorted order.
func (it Map) SortedKeys() ([]string, error) {
	if len(it) == 0 {
		return []string{}, nil
	}

	return convertinternal.
		Map.
		SortedKeys(it.Raw())
}

// SortedKeysMust
//
//	Returns the keys of the map in sorted order, panicking if an error occurs.
func (it Map) SortedKeysMust() []string {
	sortedKeys, err := it.SortedKeys()

	if err != nil {
		panic(err)
	}

	return sortedKeys
}

// When
//
//	Returns the value associated with the key "when".
func (it Map) When() (item any) {
	return it["when"]
}

// Title
//
//	Returns the value associated with the key "title".
func (it Map) Title() (item any) {
	return it["title"]
}

// Get
//
//	Retrieves an item from the map by name, returning the item and a boolean indicating if it is valid (present and defined).
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

// GetLowerCase
//
//	Retrieves an item from the map by name, converting the name to lowercase first.
func (it Map) GetLowerCase(name string) (item any, isValid bool) {
	lower := strings.ToLower(name)

	return it.Get(lower)
}

// GetDirectLower
//
//	Retrieves an item from the map by name, converting the name to lowercase first.
func (it Map) GetDirectLower(name string) any {
	x, has := it[strings.ToLower(name)]

	if has {
		return x
	}

	return nil
}

// Expect
//
//	Retrieves the value associated with the key "expect" (case-insensitive).
func (it Map) Expect() any {
	return it.GetDirectLower("expect")
}

// Actual
//
//	Retrieves the value associated with the key "actual" (case-insensitive).
func (it Map) Actual() any {
	return it.GetDirectLower("actual")
}

// Arrange
//
//	Retrieves the value associated with the key "arrange" (case-insensitive).
func (it Map) Arrange() any {
	return it.GetDirectLower("arrange")
}

// FirstItem
//
//	Retrieves the first item from the map, checking for keys "first", "f1", "p1", and "1".
func (it Map) FirstItem() any {
	return it.GetFirstOfNames("first", "f1", "p1", "1")
}

// SecondItem
//
//	Retrieves the second item from the map, checking for keys "second", "f2", "p2", and "2".
func (it Map) SecondItem() any {
	return it.GetFirstOfNames("second", "f2", "p2", "2")
}

// ThirdItem
//
//	Retrieves the third item from the map, checking for keys "third", "f3", "p3", and "3".
func (it Map) ThirdItem() any {
	return it.GetFirstOfNames("third", "f3", "p3", "3")
}

// FourthItem
//
//	Retrieves the fourth item from the map, checking for keys "fourth", "f4", "p4", and "4".
func (it Map) FourthItem() any {
	return it.GetFirstOfNames("fourth", "f4", "p4", "4")
}

// FifthItem
//
//	Retrieves the fifth item from the map, checking for keys "fifth", "f5", "p5", and "5".
func (it Map) FifthItem() any {
	return it.GetFirstOfNames("fifth", "f5", "p5", "5")
}

// SixthItem
//
//	Retrieves the sixth item from the map, checking for keys "sixth", "f6", "p6", and "6".
func (it Map) SixthItem() any {
	return it.GetFirstOfNames("sixth", "f6", "p6", "6")
}

// Seventh
//
//	Retrieves the seventh item from the map, checking for keys "seventh", "f7", "p7", and "7".
func (it Map) Seventh() any {
	return it.GetFirstOfNames("seventh", "f7", "p7", "7")
}

// SetActual
//
//	Sets the value for the key "actual" in the map.
func (it Map) SetActual(actual any) {
	it["actual"] = actual
}

// WorkFunc
//
//	Retrieves the work function from the map, checking for keys "func", "work.func", and "workFunc".
func (it Map) WorkFunc() any {
	return it.GetFirstOfNames(
		"func",
		"work.func",
		"workFunc",
	)
}

// GetFirstOfNames
//
//	Retrieves the first defined value from the map for the given names.
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

// GetAsStringSliceFirstOfNames
//
//	Retrieves the first defined value from the map for the given names, converting it to a string slice.
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

// WorkFuncName
//
//	Retrieves the name of the work function from the map.
func (it Map) WorkFuncName() string {
	workFunc := it.WorkFunc()

	return reflectinternal.GetFunc.NameOnly(workFunc)
}

// FuncWrap
//
//	Wraps the work function in a FuncWrap struct.
func (it Map) FuncWrap() *FuncWrapAny {
	return NewFuncWrap.Default(it.WorkFunc())
}

// Invoke
//
//	Invokes the work function with the given arguments.
func (it Map) Invoke(args ...any) (
	results []any, processingErr error,
) {
	return it.FuncWrap().Invoke(args...)
}

// InvokeMust
//
//	Invokes the work function with the given arguments, panicking if an error occurs.
func (it Map) InvokeMust(args ...any) (results []any) {
	results, err := it.FuncWrap().Invoke(args...)

	if err != nil {
		panic(err)
	}

	return results
}

// InvokeWithValidArgs
//
//	Invokes the work function with the valid arguments from the map.
func (it Map) InvokeWithValidArgs() (
	results []any, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.ValidArgs()

	return funcWrap.Invoke(validArgs...)
}

// InvokeArgs
//
//	Invokes the work function with the arguments specified by the given names.
func (it Map) InvokeArgs(names ...string) (
	results []any, processingErr error,
) {
	funcWrap := it.FuncWrap()
	validArgs := it.Args(names...)

	return funcWrap.Invoke(validArgs...)
}

// ValidArgs
//
//	Returns the valid arguments from the map (defined and not functions).
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

// Raw
//
//	Returns the raw map.
func (it Map) Raw() map[string]any {
	return it
}

// Args
//
//	Returns the arguments from the map for the given names.
func (it Map) Args(names ...string) []any {
	var args []any

	for _, key := range names {
		val := it[key]
		args = append(args, val)
	}

	return args
}

// GetFirstFuncNameOf
//
//	Retrieves the name of the first function found for the given names.
func (it Map) GetFirstFuncNameOf(names ...string) string {
	workFunc := it.GetFirstOfNames(names...)

	return reflectinternal.GetFunc.NameOnly(workFunc)
}

// GetAsInt
//
//	Retrieves an item from the map by name, converting it to an integer.
func (it Map) GetAsInt(name string) (item int, isValid bool) {
	i, ok := it.Get(name)
	if !ok {
		return 0, false
	}

	conv, ok := i.(int)

	return conv, ok
}

// GetAsIntDefault
//
//	Retrieves an item from the map by name, converting it to an integer, or returning a default value if not found or not an integer.
func (it Map) GetAsIntDefault(name string, defaultVal int) (item int) {
	v, isValid := it.GetAsInt(name)

	if isValid {
		return v
	}

	return defaultVal
}

// GetAsString
//
//	Retrieves an item from the map by name, converting it to a string.
func (it Map) GetAsString(name string) (item string, isValid bool) {
	i, ok := it.Get(name)
	if !ok {
		return "", false
	}

	conv, ok := i.(string)

	return conv, ok
}

// GetAsStringDefault
//
//	Retrieves an item from the map by name, converting it to a string, or returning an empty string if not found or not a string.
func (it Map) GetAsStringDefault(name string) (item string) {
	v, isValid := it.GetAsString(name)

	if isValid {
		return v
	}

	return ""
}

// GetAsStrings
//
//	Retrieves an item from the map by name, converting it to a string slice.
func (it Map) GetAsStrings(name string) (items []string, isValid bool) {
	i, ok := it.Get(name)
	if !ok {
		return []string{}, false
	}

	conv, ok := i.([]string)

	return conv, ok
}

// GetAsAnyItems
//
//	Retrieves an item from the map by name, converting it to an any slice.
func (it Map) GetAsAnyItems(name string) (items []any, isValid bool) {
	i, ok := it.Get(name)
	if !ok {
		return []any{}, false
	}

	conv, ok := i.([]any)

	return conv, ok
}

// Slice
//
//	Returns a slice representation of the map, with each item formatted as "key : value".
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

// String
//
//	Returns a string representation of the map, with each item on a new line and sorted by key.
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
