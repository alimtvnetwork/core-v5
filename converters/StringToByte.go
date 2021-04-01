package converters

import (
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
)

func StringToByte(input string) (byte, error) {
	if input == "" {
		return 0, msgtype.FailedToConvert.
			Error(msgtype.CannotConvertStringToByte, input)
	}

	if input == "0" {
		return 0, nil
	}

	if input == "1" {
		return 1, nil
	}

	vInt, err := strconv.Atoi(input)

	if err != nil {
		return 0, err
	}

	if vInt < 0 {
		return 0, msgtype.FailedToConvert.
			Error(msgtype.CannotConvertStringToByteForLessThanZero, input)
	}

	if vInt > constants.MaxUnit8AsInt {
		return 0, msgtype.FailedToConvert.
			Error(msgtype.CannotConvertStringToByteForMoreThan255, input)
	}

	return byte(vInt), nil
}
