package corejson

func JsonString(anyItem any) (jsonString string, err error) {
	jsonResult := New(anyItem)

	jsonString = jsonResult.JsonString()
	err = jsonResult.MeaningfulError()
	jsonResult.Dispose()

	return jsonString, err
}
