package regexnew

import (
	"bytes"
	"encoding/json"

	"github.com/alimtvnetwork/core/constants"
)

// prettyJson
//
// Warning:
//
//	swallows error
func prettyJson(anyItem any) string {
	if anyItem == nil {
		return ""
	}

	allBytes, err := json.Marshal(anyItem)

	if err != nil || len(allBytes) == 0 {
		return ""
	}

	var prettyJSON bytes.Buffer

	json.Indent(
		&prettyJSON,
		allBytes,
		constants.EmptyString,
		constants.Tab)

	return prettyJSON.String()
}
