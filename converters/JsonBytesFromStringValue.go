package converters

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func JsonBytesFromStringValue(name string) []byte {
	doubleQuoted := fmt.Sprintf(
		constants.SprintDoubleQuoteFormat,
		name)

	return []byte(doubleQuoted)
}
