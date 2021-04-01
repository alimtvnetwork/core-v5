package coretests

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func ToStringNameValues(any interface{}) string {
	if any == nil {
		return constants.NilAngelBracket
	}

	return fmt.Sprintf(constants.SprintFullPropertyNameValueFormat, any)
}
