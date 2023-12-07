package args

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type Dynamic struct {
	Params   Map         `json:",omitempty"`
	Expect   interface{} `json:",omitempty"`
	toSlice  *[]interface{}
	toString corestr.SimpleStringOnce
}

func (it *Dynamic) FirstItem() interface{} {
	return it.Params.FirstItem()
}

func (it *Dynamic) SecondItem() interface{} {
	return it.Params.SecondItem()
}

func (it *Dynamic) ThirdItem() interface{} {
	return it.Params.ThirdItem()
}

func (it *Dynamic) FourthItem() interface{} {
	return it.Params.FourthItem()
}

func (it *Dynamic) FifthItem() interface{} {
	return it.Params.FifthItem()
}

func (it *Dynamic) SixthItem() interface{} {
	return it.Params.SixthItem()
}

func (it *Dynamic) Expected() interface{} {
	return it.Expect
}

// HasDefined
//
// Confirms that key is present and defined.
func (it *Dynamic) HasDefined(name string) bool {
	if it == nil {
		return false
	}

	item, has := it.Params[name]

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
func (it *Dynamic) Has(name string) bool {
	if it == nil {
		return false
	}

	_, has := it.Params[name]

	return has
}

// HasDefinedAll
//
// Confirms that key is present and defined.
func (it *Dynamic) HasDefinedAll(names ...string) bool {
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
func (it *Dynamic) IsKeyInvalid(name string) bool {
	if it == nil {
		return false
	}

	item, has := it.Params[name]

	return !has ||
		reflectinternal.Is.Null(item)
}

// IsKeyMissing
//
// confirms yes if key is missing  only.
// To check either missing or null use IsKeyInvalid.
func (it *Dynamic) IsKeyMissing(name string) bool {
	if it == nil {
		return false
	}

	_, has := it.Params[name]

	return !has
}

func (it *Dynamic) GetLowerCase(name string) (item interface{}, isValid bool) {
	lower := strings.ToLower(name)

	return it.Get(lower)
}

func (it *Dynamic) GetDirectLower(name string) interface{} {
	x, has := it.Params[strings.ToLower(name)]

	if has {
		return x
	}

	return nil
}

func (it *Dynamic) Actual() interface{} {
	return it.GetDirectLower("actual")
}

func (it *Dynamic) Arrange() interface{} {
	return it.GetDirectLower("arrange")
}

func (it *Dynamic) Get(name string) (item interface{}, isValid bool) {
	if it == nil {
		return nil, false
	}

	item, has := it.Params[name]

	if has {
		return item, reflectinternal.Is.Defined(item)
	}

	return nil, false
}

func (it *Dynamic) GetAsInt(name string) (item int, isValid bool) {
	i, isValid := it.Get(name)

	if !isValid {
		return 0, false
	}

	conv, isValid := i.(int)

	return conv, isValid
}

func (it *Dynamic) GetAsIntDefault(name string, defaultVal int) (item int) {
	v, isValid := it.GetAsInt(name)

	if isValid {
		return v
	}

	return defaultVal
}

func (it *Dynamic) GetAsString(name string) (item string, isValid bool) {
	i, isValid := it.Get(name)

	if !isValid {
		return "", false
	}

	conv, isValid := i.(string)

	return conv, isValid
}

func (it *Dynamic) GetAsStringDefault(name string) (item string) {
	v, isValid := it.GetAsString(name)

	if isValid {
		return v
	}

	return ""
}

func (it *Dynamic) GetAsStrings(name string) (items []string, isValid bool) {
	i, isValid := it.Get(name)

	if !isValid {
		return []string{}, false
	}

	conv, isValid := i.([]string)

	return conv, isValid
}

func (it *Dynamic) GetAsAnyItems(name string) (items []interface{}, isValid bool) {
	i, isValid := it.Get(name)

	if !isValid {
		return []interface{}{}, false
	}

	conv, isValid := i.([]interface{})

	return conv, isValid
}

func (it *Dynamic) HasExpect() bool {
	return it != nil && reflectinternal.Is.Defined(it.Expect)
}

func (it *Dynamic) ValidArgs() []interface{} {
	var args []interface{}

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

func (it *Dynamic) Args(names ...string) []interface{} {
	var args []interface{}

	for _, key := range names {
		val := it.Params[key]
		args = append(args, val)
	}

	return args
}

func (it *Dynamic) Slice() []interface{} {
	if it.toSlice != nil {
		return *it.toSlice
	}

	var args []interface{}

	keys, err := converters.Map.SortedKeys(it.Params)

	if err != nil {
		panic(err)
	}

	for i, key := range keys {
		value := it.Params[key]
		args = append(
			args, fmt.Sprintf(
				"%d. %s : %s",
				i,
				key,
				value,
			),
		)
	}

	if it.HasExpect() {
		args = append(args, it.Expect)
	}

	it.toSlice = &args

	return *it.toSlice
}

func (it *Dynamic) String() string {
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
