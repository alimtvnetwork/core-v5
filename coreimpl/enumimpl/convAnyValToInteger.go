package enumimpl

import (
	"fmt"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
)

func convAnyValToInteger(val interface{}) int {
	_, isStr := val.(string)

	if isStr {
		return constants.MinInt
	}

	str := fmt.Sprintf(
		constants.SprintValueFormat,
		val)

	atoi, err := strconv.Atoi(str)

	if err != nil {
		return constants.MinInt
	}

	return atoi
}
