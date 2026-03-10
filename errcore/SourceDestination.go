package errcore

import "github.com/alimtvnetwork/core/constants"

func SourceDestination(
	isIncludeType bool,
	srcVal,
	destinationVal any,
) string {
	return VarTwo(
		isIncludeType,
		constants.SourceLower,
		srcVal,
		constants.DestinationLower,
		destinationVal)
}
