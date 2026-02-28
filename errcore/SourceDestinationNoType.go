package errcore

import "gitlab.com/auk-go/core/constants"

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
