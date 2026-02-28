package corejson

func JsonStringOrErrMsg(anyItem any) (jsonStringOrErr string) {
	jsonResult := New(anyItem)

	if jsonResult.HasError() {
		errMsg := jsonResult.MeaningfulErrorMessage()
		jsonResult.Dispose()

		return errMsg
	}

	jsonString := jsonResult.JsonString()
	jsonResult.Dispose()

	return jsonString
}
