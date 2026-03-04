package keymk

import (
	"encoding/json"

	"gitlab.com/auk-go/core/coredata/corejson"
)

func (it *Key) TemplateReplacer() templateReplacer {
	return templateReplacer{
		it,
	}
}

func (it *Key) JsonModel() keyModel {
	return keyModel{
		Option:        it.option,
		MainName:      it.mainName,
		KeyChains:     it.keyChains,
		CompiledChain: it.compiledChain,
	}
}

func (it *Key) JsonModelAny() any {
	return it.JsonModel()
}

func (it Key) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *Key) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *Key) UnmarshalJSON(data []byte) error {
	var deserializedModel keyModel
	err := json.Unmarshal(data, &deserializedModel)

	if err == nil {
		it.option = deserializedModel.Option
		it.mainName = deserializedModel.MainName
		it.keyChains = deserializedModel.KeyChains
		it.compiledChain = deserializedModel.CompiledChain
	}

	return err
}

func (it Key) Json() corejson.Result {
	return corejson.New(it)
}

func (it Key) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it Key) JsonString() string {
	return corejson.NewPtr(it).JsonString()
}

// ParseInjectUsingJson It will not update the self but creates a new one.
func (it *Key) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Key, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *Key) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Key {
	deserialized, err := it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return deserialized
}

func (it *Key) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Key) AsJsoner() corejson.Jsoner {
	return it
}

func (it *Key) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it Key) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return &it
}

func (it Key) AsJsonMarshaller() corejson.JsonMarshaller {
	return &it
}
