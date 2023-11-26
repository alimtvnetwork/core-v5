package args

import (
	"fmt"
	"sort"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/internal/msgcreator"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type Map map[string]interface{}

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
		reflectinternal.IsNull(item)
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

func (it Map) When() (item interface{}) {
	return it["when"]
}

func (it Map) Title() (item interface{}) {
	return it["title"]
}

func (it Map) Get(name string) (item interface{}, isValid bool) {
	if it == nil {
		return nil, false
	}

	item, has := it[name]

	if has {
		return item, reflectinternal.Is.Defined(item)
	}

	return nil, false
}

func (it Map) GetLowerCase(name string) (item interface{}, isValid bool) {
	lower := strings.ToLower(name)

	return it.Get(lower)
}

func (it Map) GetDirectLower(name string) interface{} {
	x, has := it[strings.ToLower(name)]

	if has {
		return x
	}

	return nil
}

func (it Map) Expect() interface{} {
	return it.GetDirectLower("expect")
}

func (it Map) Actual() interface{} {
	return it.GetDirectLower("actual")
}

func (it Map) Arrange() interface{} {
	return it.GetDirectLower("arrange")
}

func (it Map) First() interface{} {
	return it.GetFirstOfNames("first", "f1", "p1", "1")
}

func (it Map) Second() interface{} {
	return it.GetFirstOfNames("second", "f2", "p2", "2")
}

func (it Map) Third() interface{} {
	return it.GetFirstOfNames("third", "f3", "p3", "3")
}

func (it Map) Fourth() interface{} {
	return it.GetFirstOfNames("fourth", "f4", "p4", "4")
}

func (it Map) Fifth() interface{} {
	return it.GetFirstOfNames("fifth", "f5", "p5", "5")
}

func (it Map) Sixth() interface{} {
	return it.GetFirstOfNames("sixth", "f6", "p6", "6")
}

func (it Map) Seventh() interface{} {
	return it.GetFirstOfNames("seventh", "f7", "p7", "7")
}

func (it Map) SetActual(actual interface{}) {
	it["actual"] = actual
}

func (it Map) WorkFunc() interface{} {
	return it.GetFirstOfNames(
		"func",
		"work.func",
		"workFunc",
	)
}

func (it Map) GetFirstOfNames(names ...string) interface{} {
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

	return reflectinternal.GetFuncName(workFunc)
}

func (it Map) GetFirstFuncNameOf(names ...string) string {
	workFunc := it.GetFirstOfNames(names...)

	return reflectinternal.GetFuncName(workFunc)
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

func (it Map) GetAsAnyItems(name string) (items []interface{}, isValid bool) {
	i, isValid := it.Get(name)

	if !isValid {
		return []interface{}{}, false
	}

	conv, isValid := i.([]interface{})

	return conv, isValid
}

func (it Map) Slice() []interface{} {
	var slice []interface{}

	keys, err := converters.Map.SortedKeys(it)

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
	)

	sort.Strings(toLines)

	toFinalString := fmt.Sprintf(
		"%s {\n%s\n}\n",
		"Map",
		strings.Join(toLines, constants.NewLineUnix),
	)

	return toFinalString
}
