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

	valInt, isInt := val.(int)

	if isInt {
		return valInt
	}

	switch casted := val.(type) {
	case valueByter:
		return int(casted.Value())
	case exactValueByter:
		return int(casted.ValueByte())
	case valueInter:
		return casted.Value()
	case exactValueInter:
		return casted.ValueInt()
	case valueInt8er:
		return int(casted.Value())
	case exactValueInt8er:
		return int(casted.ValueInt8())
	}

	str := fmt.Sprintf(
		constants.SprintValueFormat,
		val)

	convValueInt, err := strconv.Atoi(str)

	if err != nil {
		return constants.MinInt
	}

	return convValueInt
}
