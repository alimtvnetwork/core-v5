package simplewrap

import (
	"fmt"
	
	"github.com/alimtvnetwork/core/constants"
)

func MsgWrapNumber(name string, number any) string {
	return fmt.Sprintf(
		constants.StringWithBracketWrapNumberFormat,
		name,
		number)
}
