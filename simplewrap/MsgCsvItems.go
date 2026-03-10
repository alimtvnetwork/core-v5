package simplewrap

import (
	"fmt"
	
	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/internal/csvinternal"
)

func MsgCsvItems(
	msg string,
	csvItems ...any,
) string {
	csvString := csvinternal.AnyItemsToStringDefault(
		csvItems...)
	
	return fmt.Sprintf(
		constants.StringWrapValueFormat,
		msg,
		csvString)
}
