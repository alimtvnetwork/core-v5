package defaultcapacity

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/constants/percentconst"
)

func PredictiveDefaultSmall(possibleLen int) int {
	return Predictive(
		possibleLen,
		percentconst.OnePointTwoPercentIncrement,
		constants.Capacity4)
}
