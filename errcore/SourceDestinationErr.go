package errcore

import (
	"errors"

	"gitlab.com/auk-go/core/constants"
)

func SourceDestinationErr(
	isIncludeType bool,
	srcVal,
	destinationVal any,
) error {
	message := VarTwo(
		isIncludeType,
		constants.SourceLower,
		srcVal,
		constants.DestinationLower,
		destinationVal)

	return errors.New(message)
}
