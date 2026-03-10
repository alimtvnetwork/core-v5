package coretaskinfo

import (
	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corejson"
)

func (it *Info) Map() map[string]string {
	if it.IsNull() || it.IsExcludeAdditionalErrorWrap() {
		return map[string]string{}
	}

	newMap := make(
		map[string]string,
		constants.Capacity8)

	if it.IsIncludeRootName() {
		newMap[infoFieldName] = it.RootName
	}

	if it.IsIncludeDescription() {
		newMap[infoFieldDescription] = it.Description
	}

	if it.IsIncludeUrl() {
		newMap[infoFieldUrl] = it.Url
	}

	if it.IsIncludeHintUrl() {
		newMap[infoFieldHintUrl] = it.HintUrl
	}

	if it.IsIncludeErrorUrl() {
		newMap[infoFieldErrorUrl] = it.ErrorUrl
	}

	if it.IsIncludeExampleUrl() {
		newMap[infoFieldExampleUrl] = it.ExampleUrl
	}

	if it.IsIncludeSingleExample() {
		newMap[infoFieldSingleExample] = it.SingleExample
	}

	if it.IsIncludeExamples() {
		newMap[infoFieldExamples] = it.ExamplesAsString()
	}

	return newMap
}

func (it *Info) MapWithPayload(
	payloads []byte,
) map[string]string {
	compiledMap := it.Map()

	if it.IsIncludePayloads() {
		compiledMap[payloadsField] = corejson.BytesToString(payloads)
	}

	return compiledMap
}

func (it *Info) LazyMapWithPayload(
	payloads []byte,
) map[string]string {
	compiledMap := it.LazyMap()

	if it.IsIncludePayloads() {
		compiledMap[payloadsField] = corejson.BytesToString(payloads)
	}

	return compiledMap
}

func (it *Info) MapWithPayloadAsAny(
	payloadsAny any,
) map[string]string {
	compiledMap := it.Map()

	if it.IsExcludePayload() {
		return compiledMap
	}

	jsonResult := corejson.
		AnyTo.
		SerializedJsonResult(payloadsAny)

	if jsonResult.HasError() {
		compiledMap[payloadsErrField] = jsonResult.MeaningfulErrorMessage()
	}

	compiledMap[payloadsField] = jsonResult.JsonString()

	return compiledMap
}

func (it *Info) LazyMapWithPayloadAsAny(
	payloadsAny any,
) map[string]string {
	compiledMap := it.LazyMap()

	if it.IsExcludePayload() {
		return compiledMap
	}

	jsonResult := corejson.
		AnyTo.
		SerializedJsonResult(payloadsAny)

	if jsonResult.HasError() {
		compiledMap[payloadsErrField] = jsonResult.MeaningfulErrorMessage()
	}

	compiledMap[payloadsField] = jsonResult.JsonString()

	return compiledMap
}

func (it *Info) LazyMap() map[string]string {
	if it.IsNull() || it.IsExcludeAdditionalErrorWrap() {
		return map[string]string{}
	}

	if it.lazyMap != nil {
		return it.lazyMap
	}

	it.lazyMap = it.Map()

	return it.lazyMap
}
