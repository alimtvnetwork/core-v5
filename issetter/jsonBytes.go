package issetter

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func jsonBytes(name string) []byte {
	doubleQuoted := fmt.Sprintf(
		constants.SprintDoubleQuoteFormat,
		name)

	return []byte(doubleQuoted)
}
