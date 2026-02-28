package corepayload

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
)

// PayloadsCollectionJson.go — JSON serialization, deserialization, and string formatting
// methods extracted from PayloadsCollection.go

func (it *PayloadsCollection) StringsUsingFmt(formatter Formatter) []string {
	list := make([]string, it.Length())

	for i := range it.Items {
		list[i] = formatter(it.Items[i])
	}

	return list
}

func (it *PayloadsCollection) JoinUsingFmt(formatter Formatter, joiner string) string {
	lines := it.StringsUsingFmt(formatter)

	return strings.Join(lines, joiner)
}

func (it *PayloadsCollection) JsonStrings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.JsonString()
	}

	return list
}

func (it *PayloadsCollection) JoinJsonStrings(joiner string) string {
	return strings.Join(it.JsonStrings(), joiner)
}

func (it *PayloadsCollection) Join(joiner string) string {
	return strings.Join(it.Strings(), joiner)
}

func (it *PayloadsCollection) JoinCsv() string {
	return strings.Join(it.CsvStrings(), constants.Comma)
}

func (it *PayloadsCollection) JoinCsvLine() string {
	return strings.Join(it.CsvStrings(), constants.CommaUnixNewLine)
}

func (it *PayloadsCollection) JsonString() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.JsonPtr().JsonString()
}

func (it *PayloadsCollection) String() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.JsonPtr().JsonString()
}

func (it *PayloadsCollection) PrettyJsonString() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.JsonPtr().PrettyJsonString()
}

func (it *PayloadsCollection) CsvStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	newSlice := make([]string, it.Length())

	for i, item := range it.Items {
		newSlice[i] = fmt.Sprintf(
			constants.SprintDoubleQuoteFormat,
			item.String())
	}

	return newSlice
}

func (it PayloadsCollection) Json() corejson.Result {
	return corejson.New(it)
}

func (it PayloadsCollection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *PayloadsCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*PayloadsCollection, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return Empty.PayloadsCollection(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *PayloadsCollection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *PayloadsCollection {
	hashSet, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return hashSet
}

func (it *PayloadsCollection) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *PayloadsCollection) AsJsoner() corejson.Jsoner {
	return it
}

func (it *PayloadsCollection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *PayloadsCollection) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}
