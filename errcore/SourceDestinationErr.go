package errcore

import (
	"errors"

	"github.com/alimtvnetwork/core/constants"
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
