package issetter

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func jsonBytes(name string) []byte {
	doubleQuoted := fmt.Sprintf(
		constants.SprintDoubleQuoteFormat,
		name)

	return []byte(doubleQuoted)
}
