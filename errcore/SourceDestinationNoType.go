package errcore

import "github.com/alimtvnetwork/core/constants"

func SourceDestinationNoType(
	srcVal,
	destinationVal any,
) string {
	return VarTwo(
		false,
		constants.SourceLower,
		srcVal,
		constants.DestinationLower,
		destinationVal)
}
